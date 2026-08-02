package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	randomSource           = "/dev/urandom"
	outputPath             = "reversed-random-characters.txt"
	characterCount         = 10
	instrumentationName    = "ai-incremental-tracing"
	instrumentationVersion = "1.0.0"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "program failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// This is the starting context for the whole operation. There is no
	// OpenTelemetry data in it yet. Later, it will carry the current span from
	// run into each application function.
	ctx := context.Background()

	// The exporter determines where completed spans are sent. This learning
	// exporter writes them to standard output as readable JSON.
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("create trace exporter: %w", err)
	}

	// The processor receives completed spans and queues them for asynchronous
	// export in batches. This keeps exporter work away from the application
	// path. Provider shutdown flushes spans that are still waiting in the queue.
	processor := sdktrace.NewBatchSpanProcessor(exporter)

	// A Resource describes the entity producing telemetry. These attributes
	// apply to every span created by this provider, rather than to one
	// particular operation.
	res, err := resource.New(
		ctx,
		// WithHost uses a resource detector to discover host information from
		// the runtime environment instead of requiring a hardcoded value.
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceName("random-character-processor"),
			semconv.ServiceVersion("1.0.0"), //Hardcoding 1.0.0 is fine for this study program, but production services normally obtain the version from the build or deployment process.
		),
	)
	if err != nil {
		return fmt.Errorf("create resource: %w", err)
	}

	// The SDK TracerProvider is the implementation that creates recording
	// spans. The provider notifies its processor as spans start and end.
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(processor),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(provider)
	defer provider.Shutdown(ctx)

	// A Tracer creates spans. The name identifies the code doing the
	// instrumentation. otel.Tracer asks the globally registered provider for
	// this tracer; the provider is not passed as a direct argument.
	tracer := newTracer()

	// Start creates one span around the complete operation and returns a new
	// context containing that span. Because the background context has no
	// parent span, "process random characters" is a root span.
	//
	// When this span ends, the processor will give it to the stdout exporter.
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

	fmt.Printf("Original: %s\n", characters)
	fmt.Printf("Reversed: %s\n", reversed)
	fmt.Printf("Saved the reversed characters to %s\n", outputPath)

	return nil
}

func readRandomCharacters(ctx context.Context, path string, count int) (string, error) {
	// The incoming context carries the root span created in run. Starting a
	// span from that context makes this new span a child of the root span.
	tracer := newTracer()
	//   Start returns a new context containing the read span. We discard it because this function does not currently call another operation that needs to become a child of the read span.
	_, span := tracer.Start(ctx, "read random characters")
	defer span.End()

	// Attributes describe properties of this particular operation. The count
	// is useful for filtering spans without recording the generated characters.
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
	// This function receives the root span's context from run, just like
	// readRandomCharacters. Its span is therefore another child of the root
	// span, making the two function spans siblings.
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
	// This function also receives the root span's context, so its span is the
	// third sibling underneath "process random characters".
	tracer := newTracer()
	_, span := tracer.Start(ctx, "save characters")
	defer span.End()

	// These attributes describe this particular file-writing operation -- they are optional
	span.SetAttributes(
		attribute.String("file.path", path),
		attribute.Int("file.size_bytes", len(characters)),
	)

	if err := os.WriteFile(path, []byte(characters), 0o600); err != nil {
		// RecordError adds a timestamped exception event to the span. It does
		// not handle the Go error, so we still return it to the caller.
		span.RecordError(err)

		// Status describes the final outcome of this operation. RecordError
		// does not set it automatically, so a failed save is marked explicitly.
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	// An event records that something notable happened at a particular time
	// during the span. This event is added only after the file write succeeds.
	//
	// Events are optional and are stored inside their enclosing span; adding an
	// event does not create a separate span.
	span.AddEvent("file write completed")

	return nil
}

func newTracer() trace.Tracer {
	return otel.Tracer(
		instrumentationName,
		trace.WithInstrumentationVersion(instrumentationVersion),
	)
}
