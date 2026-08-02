package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	randomSource           = "/dev/urandom"
	outputPath             = "reversed-random-characters.txt"
	characterCount         = 10
	instrumentationName    = "ai-incremental-collector-trace-sender"
	instrumentationVersion = "1.0.0"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "program failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// This is the starting context for the whole operation. Later, it carries
	// the current span into each application function.
	ctx := context.Background()

	// This is the only conceptual tracing change from the Chapter 4 program.
	// The old stdout exporter printed spans inside the application process.
	// This OTLP exporter sends them to our local Collector over gRPC instead.
	//
	// WithInsecure is acceptable for this local exercise because the sender and
	// Collector communicate only through localhost. Production connections
	// should normally use transport security.
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint("localhost:4317"),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return fmt.Errorf("create OTLP trace exporter: %w", err)
	}

	// The processor still batches ended spans inside the application. Provider
	// shutdown flushes any remaining batch to the OTLP exporter before exit.
	processor := sdktrace.NewBatchSpanProcessor(exporter)

	// The Resource still describes the entity producing all these spans. The
	// Collector receives and preserves this information through OTLP.
	res, err := resource.New(
		ctx,
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceName("random-character-processor"),
			semconv.ServiceVersion("1.0.0"),
		),
	)
	if err != nil {
		return fmt.Errorf("create resource: %w", err)
	}

	// The provider, processor, tracers, spans, and context propagation are
	// unchanged. Only the exporter destination changed.
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(processor),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(provider)
	defer provider.Shutdown(ctx)

	tracer := newTracer()

	// Because the starting context has no parent span, this is the root span.
	ctx, span := tracer.Start(ctx, "process random characters")
	defer span.End()

	characters, err := readRandomCharacters(ctx, randomSource, characterCount)
	if err != nil {
		operationErr := fmt.Errorf("read random characters: %w", err)
		span.SetStatus(codes.Error, operationErr.Error())
		return operationErr
	}

	reversed := reverseCharacters(ctx, characters)

	if err := saveCharacters(ctx, outputPath, reversed); err != nil {
		operationErr := fmt.Errorf("save reversed characters: %w", err)
		span.SetStatus(codes.Error, operationErr.Error())
		return operationErr
	}

	// These are ordinary application messages. Trace JSON no longer appears
	// here because completed spans now travel to the Collector via OTLP.
	fmt.Printf("Original: %s\n", characters)
	fmt.Printf("Reversed: %s\n", reversed)
	fmt.Printf("Saved the reversed characters to %s\n", outputPath)

	return nil
}

func readRandomCharacters(ctx context.Context, path string, count int) (string, error) {
	// Starting from the incoming root-span context makes this a child span.
	tracer := newTracer()
	_, span := tracer.Start(ctx, "read random characters")
	defer span.End()

	span.SetAttributes(attribute.Int("random.character.count", count))

	file, err := os.Open(path)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return "", err
	}
	defer file.Close()

	randomBytes := make([]byte, count)
	if _, err := io.ReadFull(file, randomBytes); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return "", err
	}

	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i, randomByte := range randomBytes {
		randomBytes[i] = alphabet[int(randomByte)%len(alphabet)]
	}

	return string(randomBytes), nil
}

func reverseCharacters(ctx context.Context, characters string) string {
	// This function also starts from the root span, making its span a sibling
	// of the read and save spans.
	tracer := newTracer()
	_, span := tracer.Start(ctx, "reverse characters")
	defer span.End()

	reversed := []rune(characters)
	for left, right := 0, len(reversed)-1; left < right; left, right = left+1, right-1 {
		reversed[left], reversed[right] = reversed[right], reversed[left]
	}

	return string(reversed)
}

func saveCharacters(ctx context.Context, path, characters string) error {
	tracer := newTracer()
	_, span := tracer.Start(ctx, "save characters")
	defer span.End()

	span.SetAttributes(
		attribute.String("file.path", path),
		attribute.Int("file.size_bytes", len(characters)),
	)

	if err := os.WriteFile(path, []byte(characters), 0o600); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	span.AddEvent("file write completed")

	return nil
}

func newTracer() trace.Tracer {
	return otel.Tracer(
		instrumentationName,
		trace.WithInstrumentationVersion(instrumentationVersion),
	)
}
