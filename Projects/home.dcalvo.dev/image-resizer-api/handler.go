package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

const multipartOverheadAllowance = 1 << 20

var (
	errImageMissing   = errors.New("image field is missing")
	errImageDuplicate = errors.New("image field appears more than once")
	errImageTooLarge  = errors.New("image exceeds the upload limit")
)

type config struct {
	maxUploadBytes int64
	maxInputPixels int64
	maxOutputWidth int
	jpegQuality    int
}

func defaultConfig() config {
	return config{
		maxUploadBytes: 10 << 20,
		maxInputPixels: 40_000_000,
		maxOutputWidth: 4_000,
		jpegQuality:    85,
	}
}

type handler struct {
	config  config
	logger  *slog.Logger
	metrics *metrics
}

type routes struct {
	handler http.Handler
	ready   atomic.Bool
}

func newHandler(cfg config, logger *slog.Logger) *routes {
	h := handler{config: cfg, logger: logger, metrics: newMetrics()}
	routes := &routes{}
	routes.ready.Store(true)
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/resize", h.resize)
	mux.Handle("/metrics", h.metrics.handler())
	mux.HandleFunc("/livez", routes.livez)
	mux.HandleFunc("/readyz", routes.readyz)
	routes.handler = mux
	return routes
}

func (r *routes) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	r.handler.ServeHTTP(w, request)
}

func (r *routes) setReady(ready bool) {
	r.ready.Store(ready)
}

func (r *routes) livez(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Use GET for this endpoint")
		return
	}
	writeHealth(w, http.StatusOK, "alive")
}

func (r *routes) readyz(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Use GET for this endpoint")
		return
	}
	if !r.ready.Load() {
		writeHealth(w, http.StatusServiceUnavailable, "not ready")
		return
	}
	writeHealth(w, http.StatusOK, "ready")
}

func writeHealth(w http.ResponseWriter, status int, state string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	_, _ = fmt.Fprintln(w, state)
}

func (h handler) resize(w http.ResponseWriter, r *http.Request) {
	request := newRequestLog(h.logger, h.metrics, r)
	w.Header().Set("X-Request-ID", request.id)
	defer request.write()

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		request.writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Use POST for this endpoint")
		return
	}

	targetWidth, err := h.targetWidth(r)
	if err != nil {
		request.writeError(w, http.StatusBadRequest, "invalid_width", err.Error())
		return
	}
	request.targetWidth = targetWidth

	r.Body = http.MaxBytesReader(w, r.Body, h.config.maxUploadBytes+multipartOverheadAllowance)
	data, err := readImagePart(r, h.config.maxUploadBytes)
	if err != nil {
		h.writeUploadError(w, request, err)
		return
	}
	request.inputBytes = len(data)

	imageConfig, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil || format != "jpeg" {
		request.writeError(w, http.StatusUnsupportedMediaType, "unsupported_image_format", "Only valid JPEG images are supported")
		return
	}
	request.inputWidth = imageConfig.Width
	request.inputHeight = imageConfig.Height

	inputPixels := int64(imageConfig.Width) * int64(imageConfig.Height)
	if imageConfig.Width <= 0 || imageConfig.Height <= 0 || inputPixels > h.config.maxInputPixels {
		request.writeError(w, http.StatusRequestEntityTooLarge, "image_dimensions_too_large", fmt.Sprintf("JPEG images may contain at most %d pixels", h.config.maxInputPixels))
		return
	}

	if targetWidth > imageConfig.Width {
		request.writeError(w, http.StatusUnprocessableEntity, "upscaling_not_supported", "The target width cannot exceed the original image width")
		return
	}
	h.metrics.inputBytes.Observe(float64(len(data)))
	h.metrics.inputPixels.Observe(float64(inputPixels))
	h.metrics.resizeInFlight.Inc()
	defer h.metrics.resizeInFlight.Dec()

	stageStarted := time.Now()
	source, err := jpeg.Decode(bytes.NewReader(data))
	h.metrics.observeStage("decode", time.Since(stageStarted))
	if err != nil {
		request.writeError(w, http.StatusUnsupportedMediaType, "invalid_jpeg", "The uploaded JPEG could not be decoded")
		return
	}

	stageStarted = time.Now()
	resized := resizeWidth(source, targetWidth)
	h.metrics.observeStage("resize", time.Since(stageStarted))
	request.outputHeight = resized.Bounds().Dy()
	var output bytes.Buffer
	stageStarted = time.Now()
	if err := jpeg.Encode(&output, resized, &jpeg.Options{Quality: h.config.jpegQuality}); err != nil {
		h.metrics.observeStage("encode", time.Since(stageStarted))
		request.writeError(w, http.StatusInternalServerError, "encoding_failed", "The resized image could not be encoded")
		return
	}
	h.metrics.observeStage("encode", time.Since(stageStarted))
	request.outputBytes = output.Len()
	h.metrics.outputBytes.Observe(float64(output.Len()))
	h.metrics.outputPixels.Observe(float64(targetWidth * request.outputHeight))

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(output.Len()))
	w.WriteHeader(http.StatusOK)
	request.status = http.StatusOK
	_, _ = output.WriteTo(w)
}

func (h handler) targetWidth(r *http.Request) (int, error) {
	values, ok := r.URL.Query()["width"]
	if !ok || len(values) != 1 || values[0] == "" {
		return 0, errors.New("width must be provided exactly once")
	}

	width, err := strconv.Atoi(values[0])
	if err != nil || width <= 0 {
		return 0, errors.New("width must be a positive integer")
	}
	if width > h.config.maxOutputWidth {
		return 0, fmt.Errorf("width may not exceed %d pixels", h.config.maxOutputWidth)
	}
	return width, nil
}

func readImagePart(r *http.Request, maxBytes int64) ([]byte, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, fmt.Errorf("invalid multipart request: %w", err)
	}

	var imageData []byte
	for {
		part, err := reader.NextPart()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		if part.FormName() != "image" {
			err = discardPart(part)
			if err != nil {
				return nil, err
			}
			continue
		}
		if imageData != nil {
			_ = part.Close()
			return nil, errImageDuplicate
		}

		imageData, err = readLimitedPart(part, maxBytes)
		if err != nil {
			return nil, err
		}
	}

	if imageData == nil {
		return nil, errImageMissing
	}
	return imageData, nil
}

func readLimitedPart(part *multipart.Part, maxBytes int64) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(part, maxBytes+1))
	closeErr := part.Close()
	if err != nil {
		return nil, err
	}
	if closeErr != nil {
		return nil, closeErr
	}
	if int64(len(data)) > maxBytes {
		return nil, errImageTooLarge
	}
	return data, nil
}

func discardPart(part *multipart.Part) error {
	_, err := io.Copy(io.Discard, part)
	closeErr := part.Close()
	if err != nil {
		return err
	}
	return closeErr
}

func (h handler) writeUploadError(w http.ResponseWriter, request *requestLog, err error) {
	var maxBytesError *http.MaxBytesError
	switch {
	case errors.Is(err, errImageTooLarge), errors.As(err, &maxBytesError):
		request.writeError(w, http.StatusRequestEntityTooLarge, "image_too_large", fmt.Sprintf("The JPEG may not exceed %d bytes", h.config.maxUploadBytes))
	case errors.Is(err, errImageMissing):
		request.writeError(w, http.StatusBadRequest, "image_missing", "Provide one JPEG in the multipart field named image")
	case errors.Is(err, errImageDuplicate):
		request.writeError(w, http.StatusBadRequest, "multiple_images", "Provide exactly one multipart field named image")
	default:
		request.writeError(w, http.StatusBadRequest, "invalid_multipart_request", "The request must contain one multipart image field")
	}
}

type requestLog struct {
	logger       *slog.Logger
	metrics      *metrics
	context      context.Context
	started      time.Time
	id           string
	method       string
	route        string
	status       int
	errorCode    string
	targetWidth  int
	inputWidth   int
	inputHeight  int
	outputHeight int
	inputBytes   int
	outputBytes  int
}

func newRequestLog(logger *slog.Logger, metrics *metrics, r *http.Request) *requestLog {
	request := &requestLog{
		logger:  logger,
		metrics: metrics,
		context: r.Context(),
		started: time.Now(),
		id:      newRequestID(),
		method:  r.Method,
		route:   "/v1/resize",
	}
	metrics.startRequest(request.method, request.route)
	return request
}

func (r *requestLog) writeError(w http.ResponseWriter, status int, code, message string) {
	r.status = status
	r.errorCode = code
	if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
		r.metrics.rejections.WithLabelValues(code).Inc()
	}
	writeError(w, status, code, message)
}

func (r *requestLog) write() {
	duration := time.Since(r.started)
	r.metrics.finishRequest(r.method, r.route, r.status, duration)
	attributes := []any{
		"request_id", r.id,
		"method", r.method,
		"route", r.route,
		"status", r.status,
		"duration_ms", duration.Milliseconds(),
	}
	if r.errorCode != "" {
		attributes = append(attributes, "error", r.errorCode)
	}
	if r.targetWidth != 0 {
		attributes = append(attributes, "target_width", r.targetWidth)
	}
	if r.inputWidth != 0 {
		attributes = append(attributes,
			"input_width", r.inputWidth,
			"input_height", r.inputHeight,
		)
	}
	if r.outputHeight != 0 {
		attributes = append(attributes, "output_height", r.outputHeight)
	}
	if r.inputBytes != 0 {
		attributes = append(attributes, "input_bytes", r.inputBytes)
	}
	if r.outputBytes != 0 {
		attributes = append(attributes, "output_bytes", r.outputBytes)
	}

	level := slog.LevelInfo
	if r.status >= http.StatusInternalServerError {
		level = slog.LevelError
	}
	r.logger.Log(r.context, level, "request completed", attributes...)
}

func newRequestID() string {
	var value [12]byte
	if _, err := rand.Read(value[:]); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(value[:])
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{
		Error:   code,
		Message: message,
	})
}
