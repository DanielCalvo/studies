package main

import (
	"context"
	"io"
	"log/slog"
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
