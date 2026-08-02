package main

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"runtime"
	"testing"
	"time"
)

func TestRunStopsWhenContextIsCanceled(t *testing.T) {
	t.Setenv("PORT", "0")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	finished := make(chan error, 1)
	go func() {
		logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
		finished <- run(ctx, logger)
	}()

	select {
	case err := <-finished:
		if err != nil {
			t.Fatalf("run returned an error during graceful shutdown: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("run did not finish after its context was canceled")
	}
}

func TestConfigFromEnvironmentUsesAvailableCPUByDefault(t *testing.T) {
	t.Setenv(maxConcurrentResizesEnv, "")

	cfg, err := configFromEnvironment()
	if err != nil {
		t.Fatalf("configFromEnvironment returned an error: %v", err)
	}
	if cfg.maxConcurrentResizes != runtime.GOMAXPROCS(0) {
		t.Fatalf("maxConcurrentResizes = %d, want GOMAXPROCS %d", cfg.maxConcurrentResizes, runtime.GOMAXPROCS(0))
	}
}

func TestConfigFromEnvironmentOverridesConcurrency(t *testing.T) {
	t.Setenv(maxConcurrentResizesEnv, "7")

	cfg, err := configFromEnvironment()
	if err != nil {
		t.Fatalf("configFromEnvironment returned an error: %v", err)
	}
	if cfg.maxConcurrentResizes != 7 {
		t.Fatalf("maxConcurrentResizes = %d, want 7", cfg.maxConcurrentResizes)
	}
}

func TestConfigFromEnvironmentRejectsInvalidConcurrency(t *testing.T) {
	for _, value := range []string{"not-a-number", "0", "-1"} {
		t.Run(value, func(t *testing.T) {
			t.Setenv(maxConcurrentResizesEnv, value)
			if _, err := configFromEnvironment(); err == nil {
				t.Fatalf("configFromEnvironment accepted %q", value)
			}
		})
	}
}

func TestHTTPServerSetsSlowClientTimeouts(t *testing.T) {
	server := newHTTPServer(":0", http.NotFoundHandler())

	if server.ReadHeaderTimeout != readHeaderTimeout {
		t.Fatalf("ReadHeaderTimeout = %v, want %v", server.ReadHeaderTimeout, readHeaderTimeout)
	}
	if server.ReadTimeout != readTimeout {
		t.Fatalf("ReadTimeout = %v, want %v", server.ReadTimeout, readTimeout)
	}
	if server.WriteTimeout != writeTimeout {
		t.Fatalf("WriteTimeout = %v, want %v", server.WriteTimeout, writeTimeout)
	}
	if server.IdleTimeout != idleTimeout {
		t.Fatalf("IdleTimeout = %v, want %v", server.IdleTimeout, idleTimeout)
	}
	if server.MaxHeaderBytes != maxHeaderBytes {
		t.Fatalf("MaxHeaderBytes = %d, want %d", server.MaxHeaderBytes, maxHeaderBytes)
	}
}
