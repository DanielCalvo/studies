package main

import (
	"context"
	"errors"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (function roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func TestRequestInterval(t *testing.T) {
	interval, err := requestInterval(120)
	if err != nil {
		t.Fatalf("requestInterval returned an error: %v", err)
	}
	if interval != 500*time.Millisecond {
		t.Fatalf("requestInterval = %s, want 500ms", interval)
	}

	if _, err := requestInterval(0); err == nil {
		t.Fatal("requestInterval accepted zero rpm")
	}
}

func TestFindImagesReadsDimensions(t *testing.T) {
	directory := t.TempDir()
	if err := generateJPEG(filepath.Join(directory, "example.jpg"), 120, 80); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directory, "ignored.txt"), []byte("not an image"), 0o644); err != nil {
		t.Fatal(err)
	}

	images, err := findImages(directory)
	if err != nil {
		t.Fatalf("findImages returned an error: %v", err)
	}
	if len(images) != 1 {
		t.Fatalf("findImages returned %d images, want 1", len(images))
	}
	if images[0].width != 120 || images[0].height != 80 {
		t.Fatalf("dimensions = %dx%d, want 120x80", images[0].width, images[0].height)
	}
}

func TestSendResize(t *testing.T) {
	imagePath := filepath.Join(t.TempDir(), "120x80.jpg")
	if err := generateJPEG(imagePath, 120, 80); err != nil {
		t.Fatal(err)
	}

	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", request.Method)
		}
		if width := request.URL.Query().Get("width"); width != "60" {
			t.Errorf("width = %q, want 60", width)
		}
		file, _, err := request.FormFile("image")
		if err != nil {
			t.Errorf("read multipart image: %v", err)
			return testResponse(http.StatusBadRequest), nil
		}
		defer file.Close()
		config, err := jpeg.DecodeConfig(file)
		if err != nil {
			t.Errorf("decode uploaded JPEG: %v", err)
			return testResponse(http.StatusBadRequest), nil
		}
		if config.Width != 120 || config.Height != 80 {
			t.Errorf("uploaded dimensions = %dx%d, want 120x80", config.Width, config.Height)
		}
		return testResponse(http.StatusOK), nil
	})}

	endpoint, err := resizeEndpoint("http://image-resizer.test")
	if err != nil {
		t.Fatal(err)
	}
	item := sourceImage{path: imagePath, width: 120, height: 80}
	if err := sendResize(context.Background(), client, endpoint, item); err != nil {
		t.Fatalf("sendResize returned an error: %v", err)
	}
}

func TestSendResizeRejectsHTTPError(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(_ *http.Request) (*http.Response, error) {
		return testResponse(http.StatusBadRequest), nil
	})}

	imagePath := filepath.Join(t.TempDir(), "image.jpg")
	if err := generateJPEG(imagePath, 10, 10); err != nil {
		t.Fatal(err)
	}
	endpoint, err := resizeEndpoint("http://image-resizer.test")
	if err != nil {
		t.Fatal(err)
	}
	item := sourceImage{path: imagePath, width: 10, height: 10}
	if err := sendResize(context.Background(), client, endpoint, item); err == nil {
		t.Fatal("sendResize accepted an HTTP error response")
	}
}

func TestGenerateTrafficLaunchesRequestsWithoutWaitingForResponses(t *testing.T) {
	imagePath := filepath.Join(t.TempDir(), "image.jpg")
	if err := generateJPEG(imagePath, 10, 10); err != nil {
		t.Fatal(err)
	}
	endpoint, err := resizeEndpoint("http://image-resizer.test")
	if err != nil {
		t.Fatal(err)
	}

	launched := make(chan struct{}, 10)
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		launched <- struct{}{}
		<-request.Context().Done()
		return nil, request.Context().Err()
	})}

	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() {
		result <- generateTraffic(
			ctx,
			client,
			endpoint,
			[]sourceImage{{path: imagePath, width: 10, height: 10}},
			10*time.Millisecond,
		)
	}()

	for requestNumber := 1; requestNumber <= 3; requestNumber++ {
		select {
		case <-launched:
		case <-time.After(500 * time.Millisecond):
			cancel()
			t.Fatalf("request %d was not launched while earlier responses were blocked", requestNumber)
		}
	}

	cancel()
	select {
	case err := <-result:
		if err != nil {
			t.Fatalf("generateTraffic returned an error during cancellation: %v", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("generateTraffic did not wait for in-flight requests and stop")
	}
}

func TestGenerateTrafficContinuesAfterRequestFailure(t *testing.T) {
	imagePath := filepath.Join(t.TempDir(), "image.jpg")
	if err := generateJPEG(imagePath, 10, 10); err != nil {
		t.Fatal(err)
	}
	endpoint, err := resizeEndpoint("http://image-resizer.test")
	if err != nil {
		t.Fatal(err)
	}

	attempted := make(chan int64, 10)
	var attemptCount atomic.Int64
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempt := attemptCount.Add(1)
		attempted <- attempt
		if attempt == 1 {
			return nil, errors.New("EOF")
		}
		<-request.Context().Done()
		return nil, request.Context().Err()
	})}

	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() {
		result <- generateTraffic(
			ctx,
			client,
			endpoint,
			[]sourceImage{{path: imagePath, width: 10, height: 10}},
			10*time.Millisecond,
		)
	}()

	for expectedAttempt := int64(1); expectedAttempt <= 2; expectedAttempt++ {
		select {
		case attempt := <-attempted:
			if attempt != expectedAttempt {
				cancel()
				t.Fatalf("attempt = %d, want %d", attempt, expectedAttempt)
			}
		case <-time.After(500 * time.Millisecond):
			cancel()
			t.Fatalf("request %d was not attempted", expectedAttempt)
		}
	}

	select {
	case err := <-result:
		cancel()
		t.Fatalf("generateTraffic stopped after a request failure: %v", err)
	default:
	}

	cancel()
	select {
	case err := <-result:
		if err != nil {
			t.Fatalf("generateTraffic returned an error during cancellation: %v", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("generateTraffic did not stop after cancellation")
	}
}

func testResponse(statusCode int) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("response")),
	}
}
