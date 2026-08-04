package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	shutdownTimeout         = 30 * time.Second
	telemetryStartupTimeout = 5 * time.Second
	telemetryFlushTimeout   = 5 * time.Second
	readHeaderTimeout       = 5 * time.Second
	readTimeout             = 15 * time.Second
	writeTimeout            = 30 * time.Second
	idleTimeout             = 60 * time.Second
	maxHeaderBytes          = 1 << 20
	maxConcurrentResizesEnv = "MAX_CONCURRENT_RESIZES"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := runApplication(logger); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

func runApplication(logger *slog.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	telemetryCtx, cancelTelemetry := context.WithTimeout(context.Background(), telemetryStartupTimeout)
	provider, err := configureOpenTelemetry(telemetryCtx, logger)
	cancelTelemetry()
	if err != nil {
		return err
	}

	runErr := run(ctx, logger)
	flushCtx, cancelFlush := context.WithTimeout(context.Background(), telemetryFlushTimeout)
	defer cancelFlush()
	flushErr := provider.Shutdown(flushCtx)
	return errors.Join(runErr, flushErr)
}

func run(ctx context.Context, logger *slog.Logger) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg, err := configFromEnvironment()
	if err != nil {
		return err
	}
	routes := newHandler(cfg, logger)
	server := newHTTPServer(":"+port, routes)

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", server.Addr, err)
	}

	logger.Info("server starting",
		"address", listener.Addr().String(),
		"max_upload_bytes", cfg.maxUploadBytes,
		"max_input_pixels", cfg.maxInputPixels,
		"max_output_width", cfg.maxOutputWidth,
		"jpeg_quality", cfg.jpegQuality,
		"max_concurrent_resizes", cfg.maxConcurrentResizes,
	)

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.Serve(listener)
	}()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		routes.setReady(false)
		logger.Info("server shutting down", "timeout", shutdownTimeout)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("graceful shutdown: %w", err)
		}

		if err := <-serverErrors; err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		logger.Info("server shutdown complete")
		return nil
	}
}

func configFromEnvironment() (config, error) {
	cfg := defaultConfig()
	value := os.Getenv(maxConcurrentResizesEnv)
	if value == "" {
		return cfg, nil
	}

	limit, err := strconv.Atoi(value)
	if err != nil || limit <= 0 {
		return config{}, fmt.Errorf("%s must be a positive integer", maxConcurrentResizesEnv)
	}
	cfg.maxConcurrentResizes = limit
	return cfg, nil
}

func newHTTPServer(address string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              address,
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}
}
