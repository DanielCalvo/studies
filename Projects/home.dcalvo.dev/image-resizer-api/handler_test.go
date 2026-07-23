package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestResizeJPEG(t *testing.T) {
	server := newHandler(defaultConfig(), discardLogger())
	request := multipartRequest(t, http.MethodPost, "/v1/resize?width=3", "image", makeJPEG(t, 6, 4))
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body: %s", response.Code, http.StatusOK, response.Body.String())
	}
	if contentType := response.Header().Get("Content-Type"); contentType != "image/jpeg" {
		t.Fatalf("Content-Type = %q, want image/jpeg", contentType)
	}

	got, err := jpeg.DecodeConfig(response.Body)
	if err != nil {
		t.Fatalf("decode response JPEG: %v", err)
	}
	if got.Width != 3 || got.Height != 2 {
		t.Fatalf("output dimensions = %dx%d, want 3x2", got.Width, got.Height)
	}
}

func TestHealthEndpointsReflectLifecycleState(t *testing.T) {
	routes := newHandler(defaultConfig(), discardLogger())

	assertHealthResponse(t, routes, "/livez", http.StatusOK, "alive\n")
	assertHealthResponse(t, routes, "/readyz", http.StatusOK, "ready\n")

	routes.setReady(false)

	assertHealthResponse(t, routes, "/livez", http.StatusOK, "alive\n")
	assertHealthResponse(t, routes, "/readyz", http.StatusServiceUnavailable, "not ready\n")
}

func TestResizeRejectsInvalidRequests(t *testing.T) {
	tests := []struct {
		name       string
		configure  func(*config)
		method     string
		target     string
		field      string
		body       []byte
		wantStatus int
		wantCode   string
	}{
		{
			name:       "missing width",
			method:     http.MethodPost,
			target:     "/v1/resize",
			field:      "image",
			body:       makeJPEG(t, 4, 2),
			wantStatus: http.StatusBadRequest,
			wantCode:   "invalid_width",
		},
		{
			name:       "non-numeric width",
			method:     http.MethodPost,
			target:     "/v1/resize?width=wide",
			field:      "image",
			body:       makeJPEG(t, 4, 2),
			wantStatus: http.StatusBadRequest,
			wantCode:   "invalid_width",
		},
		{
			name:       "zero width",
			method:     http.MethodPost,
			target:     "/v1/resize?width=0",
			field:      "image",
			body:       makeJPEG(t, 4, 2),
			wantStatus: http.StatusBadRequest,
			wantCode:   "invalid_width",
		},
		{
			name:       "missing image",
			method:     http.MethodPost,
			target:     "/v1/resize?width=2",
			field:      "description",
			body:       []byte("not an image field"),
			wantStatus: http.StatusBadRequest,
			wantCode:   "image_missing",
		},
		{
			name:       "PNG input",
			method:     http.MethodPost,
			target:     "/v1/resize?width=2",
			field:      "image",
			body:       makePNG(t, 4, 2),
			wantStatus: http.StatusUnsupportedMediaType,
			wantCode:   "unsupported_image_format",
		},
		{
			name: "upload too large",
			configure: func(cfg *config) {
				cfg.maxUploadBytes = 4
			},
			method:     http.MethodPost,
			target:     "/v1/resize?width=2",
			field:      "image",
			body:       []byte("five!"),
			wantStatus: http.StatusRequestEntityTooLarge,
			wantCode:   "image_too_large",
		},
		{
			name: "too many pixels",
			configure: func(cfg *config) {
				cfg.maxInputPixels = 5
			},
			method:     http.MethodPost,
			target:     "/v1/resize?width=2",
			field:      "image",
			body:       makeJPEG(t, 3, 2),
			wantStatus: http.StatusRequestEntityTooLarge,
			wantCode:   "image_dimensions_too_large",
		},
		{
			name:       "upscaling",
			method:     http.MethodPost,
			target:     "/v1/resize?width=5",
			field:      "image",
			body:       makeJPEG(t, 4, 2),
			wantStatus: http.StatusUnprocessableEntity,
			wantCode:   "upscaling_not_supported",
		},
		{
			name:       "wrong method",
			method:     http.MethodGet,
			target:     "/v1/resize?width=2",
			field:      "image",
			body:       makeJPEG(t, 4, 2),
			wantStatus: http.StatusMethodNotAllowed,
			wantCode:   "method_not_allowed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := defaultConfig()
			if test.configure != nil {
				test.configure(&cfg)
			}

			server := newHandler(cfg, discardLogger())
			request := multipartRequest(t, test.method, test.target, test.field, test.body)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)

			if response.Code != test.wantStatus {
				t.Fatalf("status = %d, want %d; body: %s", response.Code, test.wantStatus, response.Body.String())
			}
			var got struct {
				Error string `json:"error"`
			}
			if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
				t.Fatalf("decode error response: %v", err)
			}
			if got.Error != test.wantCode {
				t.Fatalf("error code = %q, want %q", got.Error, test.wantCode)
			}
		})
	}
}

func TestResizeLogsCompletedRequest(t *testing.T) {
	var logs bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logs, nil))
	server := newHandler(defaultConfig(), logger)
	request := multipartRequest(t, http.MethodPost, "/v1/resize?width=3", "image", makeJPEG(t, 6, 4))
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	entry := decodeLogEntry(t, &logs)
	assertLogValue(t, entry, "msg", "request completed")
	assertLogValue(t, entry, "method", http.MethodPost)
	assertLogValue(t, entry, "route", "/v1/resize")
	assertLogNumber(t, entry, "status", http.StatusOK)
	assertLogNumber(t, entry, "input_width", 6)
	assertLogNumber(t, entry, "input_height", 4)
	assertLogNumber(t, entry, "target_width", 3)
	assertLogNumber(t, entry, "output_height", 2)

	requestID, ok := entry["request_id"].(string)
	if !ok || requestID == "" {
		t.Fatalf("request_id = %#v, want a non-empty string", entry["request_id"])
	}
	if got := response.Header().Get("X-Request-ID"); got != requestID {
		t.Fatalf("X-Request-ID = %q, want %q", got, requestID)
	}
	if _, ok := entry["duration_ms"].(float64); !ok {
		t.Fatalf("duration_ms = %#v, want a JSON number", entry["duration_ms"])
	}
	if value, ok := entry["input_bytes"].(float64); !ok || value <= 0 {
		t.Fatalf("input_bytes = %#v, want a positive number", entry["input_bytes"])
	}
	if value, ok := entry["output_bytes"].(float64); !ok || value <= 0 {
		t.Fatalf("output_bytes = %#v, want a positive number", entry["output_bytes"])
	}
}

func TestResizeLogsStableErrorCode(t *testing.T) {
	var logs bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logs, nil))
	server := newHandler(defaultConfig(), logger)
	request := multipartRequest(t, http.MethodPost, "/v1/resize?width=0", "image", makeJPEG(t, 4, 2))
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	entry := decodeLogEntry(t, &logs)
	assertLogNumber(t, entry, "status", http.StatusBadRequest)
	assertLogValue(t, entry, "error", "invalid_width")
	if _, exists := entry["input_bytes"]; exists {
		t.Fatal("input_bytes was logged even though the image was not read")
	}
}

func TestMetricsDescribeRequestsAndImageProcessing(t *testing.T) {
	server := newHandler(defaultConfig(), discardLogger())

	success := multipartRequest(t, http.MethodPost, "/v1/resize?width=3", "image", makeJPEG(t, 6, 4))
	server.ServeHTTP(httptest.NewRecorder(), success)

	rejected := multipartRequest(t, http.MethodPost, "/v1/resize?width=0", "image", makeJPEG(t, 4, 2))
	server.ServeHTTP(httptest.NewRecorder(), rejected)

	request := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("metrics status = %d, want %d", response.Code, http.StatusOK)
	}
	metrics := response.Body.String()
	wantSamples := []string{
		`image_resizer_http_requests_total{method="POST",outcome="succeeded",route="/v1/resize",status="200"} 1`,
		`image_resizer_http_requests_total{method="POST",outcome="rejected",route="/v1/resize",status="400"} 1`,
		`image_resizer_processing_stage_duration_seconds_count{stage="decode"} 1`,
		`image_resizer_processing_stage_duration_seconds_count{stage="resize"} 1`,
		`image_resizer_processing_stage_duration_seconds_count{stage="encode"} 1`,
		`image_resizer_input_size_bytes_count 1`,
		`image_resizer_output_size_bytes_count 1`,
		`image_resizer_input_pixels_count 1`,
		`image_resizer_output_pixels_count 1`,
		`image_resizer_rejections_total{reason="invalid_width"} 1`,
		`go_goroutines `,
		`process_resident_memory_bytes `,
	}
	for _, sample := range wantSamples {
		if !strings.Contains(metrics, sample) {
			t.Errorf("metrics output does not contain %q", sample)
		}
	}
	if strings.Contains(metrics, "request_id") {
		t.Fatal("metrics output contains the high-cardinality request_id field")
	}
}

func TestScaledHeight(t *testing.T) {
	tests := []struct {
		name                        string
		sourceWidth, sourceHeight   int
		targetWidth, expectedHeight int
	}{
		{name: "landscape", sourceWidth: 6, sourceHeight: 4, targetWidth: 3, expectedHeight: 2},
		{name: "portrait", sourceWidth: 4, sourceHeight: 8, targetWidth: 2, expectedHeight: 4},
		{name: "minimum one pixel", sourceWidth: 100, sourceHeight: 1, targetWidth: 1, expectedHeight: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := scaledHeight(test.sourceWidth, test.sourceHeight, test.targetWidth)
			if got != test.expectedHeight {
				t.Fatalf("scaledHeight(%d, %d, %d) = %d, want %d", test.sourceWidth, test.sourceHeight, test.targetWidth, got, test.expectedHeight)
			}
		})
	}
}

func multipartRequest(t *testing.T, method, target, field string, data []byte) *http.Request {
	t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile(field, "image")
	if err != nil {
		t.Fatalf("create multipart field: %v", err)
	}
	if _, err := part.Write(data); err != nil {
		t.Fatalf("write multipart field: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close multipart writer: %v", err)
	}

	request := httptest.NewRequest(method, target, &body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request
}

func makeJPEG(t *testing.T, width, height int) []byte {
	t.Helper()
	return encodeImage(t, width, height, func(writer io.Writer, source image.Image) error {
		return jpeg.Encode(writer, source, &jpeg.Options{Quality: 90})
	})
}

func makePNG(t *testing.T, width, height int) []byte {
	t.Helper()
	return encodeImage(t, width, height, png.Encode)
}

func encodeImage(t *testing.T, width, height int, encode func(io.Writer, image.Image) error) []byte {
	t.Helper()

	source := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			source.Set(x, y, color.RGBA{R: uint8(x * 10), G: uint8(y * 10), B: 100, A: 255})
		}
	}

	var buffer bytes.Buffer
	if err := encode(&buffer, source); err != nil {
		t.Fatalf("encode test image: %v", err)
	}
	return buffer.Bytes()
}

func discardLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(io.Discard, nil))
}

func decodeLogEntry(t *testing.T, logs io.Reader) map[string]any {
	t.Helper()
	var entry map[string]any
	if err := json.NewDecoder(logs).Decode(&entry); err != nil {
		t.Fatalf("decode log entry: %v", err)
	}
	return entry
}

func assertLogValue(t *testing.T, entry map[string]any, key string, want any) {
	t.Helper()
	if got := entry[key]; got != want {
		t.Fatalf("%s = %#v, want %#v", key, got, want)
	}
}

func assertLogNumber(t *testing.T, entry map[string]any, key string, want int) {
	t.Helper()
	if got := entry[key]; got != float64(want) {
		t.Fatalf("%s = %#v, want %d", key, got, want)
	}
}

func assertHealthResponse(t *testing.T, handler http.Handler, path string, wantStatus int, wantBody string) {
	t.Helper()
	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != wantStatus {
		t.Fatalf("GET %s status = %d, want %d", path, response.Code, wantStatus)
	}
	if got := response.Body.String(); got != wantBody {
		t.Fatalf("GET %s body = %q, want %q", path, got, wantBody)
	}
}
