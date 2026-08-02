package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math/rand/v2"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultBaseURL = "http://192.168.1.222"
	defaultRPM     = 120
)

type fixture struct {
	width  int
	height int
}

type sourceImage struct {
	path   string
	width  int
	height int
}

func main() {
	baseURL := flag.String("base-url", defaultBaseURL, "base URL of the image resizer service")
	rpm := flag.Int("rpm", defaultRPM, "resize requests to start per minute")
	imageDirectory := flag.String("images", "test-data", "directory containing source JPEGs")
	flag.Parse()

	if err := run(*baseURL, *imageDirectory, *rpm); err != nil {
		fmt.Fprintf(os.Stderr, "[in_flight=0] %v\n", err)
		os.Exit(1)
	}
}

func run(baseURL, imageDirectory string, rpm int) error {
	endpoint, err := resizeEndpoint(baseURL)
	if err != nil {
		return err
	}
	interval, err := requestInterval(rpm)
	if err != nil {
		return err
	}
	if err := generateFixtures(imageDirectory); err != nil {
		return err
	}
	images, err := findImages(imageDirectory)
	if err != nil {
		return err
	}

	fmt.Printf("[in_flight=0] Generating open-loop traffic against %s at %d requests per minute; press Ctrl+C to stop\n", baseURL, rpm)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	client := &http.Client{Timeout: 30 * time.Second}
	if err := generateTraffic(ctx, client, endpoint, images, interval); err != nil {
		return err
	}
	fmt.Println("[in_flight=0] Traffic generation stopped")
	return nil
}

func generateTraffic(
	ctx context.Context,
	client *http.Client,
	endpoint *url.URL,
	images []sourceImage,
	interval time.Duration,
) error {
	requestContext, cancelRequests := context.WithCancel(ctx)
	var requests sync.WaitGroup
	var inFlight atomic.Int64
	defer func() {
		cancelRequests()
		requests.Wait()
	}()

	launchRequest := func() {
		item := images[rand.IntN(len(images))]
		inFlight.Add(1)
		logTraffic(&inFlight, "Starting resize of %s from %dx%d to %dx%d",
			filepath.Base(item.path),
			item.width,
			item.height,
			item.width/2,
			item.height/2,
		)
		requests.Add(1)
		go func() {
			defer requests.Done()
			err := sendResize(requestContext, client, endpoint, item)
			inFlight.Add(-1)
			if err == nil || requestContext.Err() != nil {
				return
			}
			logTraffic(&inFlight, "Resize of %s failed; continuing: %v", filepath.Base(item.path), err)
		}()
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	launchRequest()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			launchRequest()
		}
	}
}

func logTraffic(inFlight *atomic.Int64, format string, arguments ...any) {
	message := fmt.Sprintf(format, arguments...)
	fmt.Printf("[in_flight=%d] %s\n", inFlight.Load(), message)
}

func resizeEndpoint(baseURL string) (*url.URL, error) {
	endpoint, err := url.Parse(strings.TrimRight(baseURL, "/") + "/v1/resize")
	if err != nil {
		return nil, fmt.Errorf("parse base URL: %w", err)
	}
	if endpoint.Scheme != "http" && endpoint.Scheme != "https" {
		return nil, fmt.Errorf("base URL must use http or https")
	}
	if endpoint.Host == "" {
		return nil, fmt.Errorf("base URL must include a host")
	}
	return endpoint, nil
}

func requestInterval(rpm int) (time.Duration, error) {
	if rpm <= 0 {
		return 0, fmt.Errorf("rpm must be greater than zero")
	}
	interval := time.Minute / time.Duration(rpm)
	if interval == 0 {
		return 0, fmt.Errorf("rpm is too large")
	}
	return interval, nil
}

func findImages(directory string) ([]sourceImage, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("read image directory: %w", err)
	}

	var images []sourceImage
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		extension := strings.ToLower(filepath.Ext(entry.Name()))
		if extension != ".jpg" && extension != ".jpeg" {
			continue
		}

		path := filepath.Join(directory, entry.Name())
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("open %s: %w", path, err)
		}
		config, err := jpeg.DecodeConfig(file)
		closeErr := file.Close()
		if err != nil {
			return nil, fmt.Errorf("read JPEG dimensions from %s: %w", path, err)
		}
		if closeErr != nil {
			return nil, fmt.Errorf("close %s: %w", path, closeErr)
		}
		images = append(images, sourceImage{
			path:   path,
			width:  config.Width,
			height: config.Height,
		})
	}
	if len(images) == 0 {
		return nil, fmt.Errorf("no JPEG images found in %s", directory)
	}
	sort.Slice(images, func(i, j int) bool {
		return images[i].path < images[j].path
	})
	return images, nil
}

func sendResize(ctx context.Context, client *http.Client, endpoint *url.URL, item sourceImage) error {
	imageData, err := os.ReadFile(item.path)
	if err != nil {
		return fmt.Errorf("read %s: %w", item.path, err)
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("image", filepath.Base(item.path))
	if err != nil {
		return fmt.Errorf("create multipart image field: %w", err)
	}
	if _, err := part.Write(imageData); err != nil {
		return fmt.Errorf("write multipart image field: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("finish multipart request: %w", err)
	}

	outputWidth := item.width / 2
	requestURL := *endpoint
	query := requestURL.Query()
	query.Set("width", fmt.Sprintf("%d", outputWidth))
	requestURL.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL.String(), &body)
	if err != nil {
		return fmt.Errorf("create resize request: %w", err)
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("send resize request: %w", err)
	}
	defer response.Body.Close()
	if _, err := io.Copy(io.Discard, response.Body); err != nil {
		return fmt.Errorf("read resize response: %w", err)
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("resize %s: server returned %s", filepath.Base(item.path), response.Status)
	}
	return nil
}

func generateFixtures(directory string) error {
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create image directory: %w", err)
	}

	fixtures := []fixture{
		{width: 640, height: 480},
		{width: 800, height: 1_200},
		{width: 1_200, height: 1_200},
		{width: 1_600, height: 1_200},
		{width: 2_000, height: 1_500},
		{width: 2_400, height: 1_800},
		{width: 2_000, height: 3_000},
		{width: 3_200, height: 2_400},
		{width: 4_000, height: 3_000},
		{width: 5_000, height: 5_000},
	}
	for _, item := range fixtures {
		name := fmt.Sprintf("%dx%d.jpg", item.width, item.height)
		path := filepath.Join(directory, name)
		if _, err := os.Stat(path); err == nil {
			continue
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("inspect %s: %w", path, err)
		}
		if err := generateJPEG(path, item.width, item.height); err != nil {
			return err
		}
	}
	return nil
}

func generateJPEG(path string, width, height int) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			offset := y*img.Stride + x*4
			img.Pix[offset] = uint8(x * 255 / width)
			img.Pix[offset+1] = uint8(y * 255 / height)
			img.Pix[offset+2] = uint8((x/64 + y/64) * 17)
			img.Pix[offset+3] = 255
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 85}); err != nil {
		_ = file.Close()
		return fmt.Errorf("encode %s: %w", path, err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close %s: %w", path, err)
	}
	return nil
}
