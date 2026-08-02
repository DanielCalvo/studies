# Incremental steps taken

This file records each small change made while gradually adding OpenTelemetry
tracing to the program.

## Step 1: Create the uninstrumented program

We created a normal Go program with three sequential operations:

1. `readRandomCharacters` reads ten random bytes from `/dev/urandom` and maps
   them to printable alphanumeric characters.
2. `reverseCharacters` reverses the order of those characters.
3. `saveCharacters` writes the reversed value to
   `reversed-random-characters.txt`.

At this stage, the program had no OpenTelemetry code. Establishing the ordinary
application flow first gives us something clear to instrument gradually.

The execution flow was:

```text
main
  |
  +-- readRandomCharacters
  |
  +-- reverseCharacters
  |
  +-- saveCharacters
```

## Step 2: Pass context.Context through the program

We created one background context at the beginning of `main`:

```go
ctx := context.Background()
```

We then added `ctx context.Context` to all three function signatures and passed
the same context to each function:

```go
characters, err := readRandomCharacters(ctx, randomSource, characterCount)
reversed := reverseCharacters(ctx, characters)
err = saveCharacters(ctx, outputPath, reversed)
```

The functions do not use the context yet. Consequently, this step does not
produce tracing data or change the program's behavior.

It prepares a path through which OpenTelemetry can later carry the current span:

```text
context created in main
  |
  +-- passed to readRandomCharacters
  |
  +-- passed to reverseCharacters
  |
  +-- passed to saveCharacters
```

An important terminology distinction is:

- `context.Context` carries request-scoped information, including the current
  span, through Go code.
- OpenTelemetry `SpanContext` holds tracing identity: the trace ID, span ID,
  trace flags, and trace state.

We have added `context.Context`, but we have not created a `SpanContext` or any
spans yet.

## Step 3: Create one root span with the OpenTelemetry API

We obtained a tracer from the OpenTelemetry API:

```go
tracer := otel.Tracer("ai-incremental-tracing")
```

A `Tracer` creates spans. Its name identifies the code responsible for creating
those spans; it does not identify an individual span or trace.

We then started one span around the program's complete operation:

```go
ctx, span := tracer.Start(ctx, "process random characters")
defer span.End()
```

`tracer.Start` returns two values:

1. A new `context.Context` containing the newly created current span.
2. The `Span` representing the operation being traced.

The returned context replaces the earlier background context and is passed to
all three functions. Because the original background context did not contain a
parent span, `process random characters` is a root span.

`defer span.End()` ensures that the span is ended when `main` returns. A span's
duration is the time between starting and ending it.

The conceptual trace currently looks like this:

```text
process random characters (root span)
  |
  +-- readRandomCharacters executes inside its lifetime
  +-- reverseCharacters executes inside its lifetime
  +-- saveCharacters executes inside its lifetime
```

The three functions do not have spans of their own yet.

We have only added the OpenTelemetry API. We have not configured an
OpenTelemetry SDK, `TracerProvider`, span processor, or exporter. The API
therefore uses its safe default no-op implementation:

- The application continues to behave normally.
- The instrumentation calls are present.
- The span does not record data.
- Nothing is exported or printed.

This demonstrates the separation between instrumentation code, which calls the
OpenTelemetry API, and telemetry configuration, which supplies the SDK
implementation that records and exports those calls.

## Step 4: Configure an SDK TracerProvider

We created an SDK `TracerProvider`:

```go
provider := sdktrace.NewTracerProvider()
```

The provider is the SDK component responsible for creating SDK-backed,
recording tracers and spans. We registered it as OpenTelemetry's global
provider:

```go
otel.SetTracerProvider(provider)
```

The existing tracer creation remains:

```go
tracer := otel.Tracer("ai-incremental-tracing")
```

We do not pass `provider` directly to `otel.Tracer`. `otel.Tracer` asks the
globally registered provider for a tracer:

```text
otel.Tracer("ai-incremental-tracing")
  |
  v
global TracerProvider
  |
  v
Tracer
```

The application startup code selects and configures the SDK provider, while
instrumentation code can continue using only the OpenTelemetry API.

We also scheduled the provider to shut down when `main` finishes:

```go
defer provider.Shutdown(ctx)
```

Go executes deferred calls in last-in, first-out order. Because the provider
shutdown is deferred before `span.End`, the calls execute in this order:

```text
1. span.End()
2. provider.Shutdown()
```

This ensures that the span is completed before its provider is shut down.

The pipeline now reaches the provider:

```text
application
  |
  v
Tracer
  |
  v
SDK TracerProvider
  |
  v
completed span
  |
  v
nothing consumes or exports it yet
```

The span is now SDK-backed and records while it is active. However, we still
have no `SpanProcessor` or `SpanExporter`, so no tracing information is printed
or sent anywhere.

## Step 5: Move the traced workflow into run

The program originally handled errors by calling `os.Exit(1)` directly from the
same function that deferred the span ending and provider shutdown.

This is a problem because `os.Exit` terminates a Go process immediately. It does
not execute deferred calls. On an error, these operations would have been
skipped:

```go
defer span.End()
defer provider.Shutdown(ctx)
```

We moved the application and tracing workflow into a `run` function that returns
an error:

```go
func run() error {
	// Configure tracing and defer cleanup.
	// Perform the application operations.
	// Return errors instead of exiting the process.
	return nil
}
```

The individual error paths now return wrapped errors:

```go
if err != nil {
	return fmt.Errorf("read random characters: %w", err)
}
```

Returning from `run` executes its deferred calls. `main` receives the final
error only after that cleanup has happened:

```go
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "program failed: %v\n", err)
		os.Exit(1)
	}
}
```

The error path is now:

```text
an operation fails
  |
  v
run returns an error
  |
  +-- span.End()
  |
  +-- provider.Shutdown()
  |
  v
main receives the error
  |
  +-- prints the error
  |
  v
os.Exit(1)
```

This preserves both requirements:

- OpenTelemetry lifecycle operations are allowed to finish.
- The process still returns a non-zero exit status to indicate failure.

## Step 6: Add a span processor and stdout exporter

Until this step, the SDK provider created a recording span, but nothing consumed
the span when it ended.

We first created an exporter:

```go
exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
if err != nil {
	return fmt.Errorf("create trace exporter: %w", err)
}
```

A `SpanExporter` determines where completed spans are sent. This exporter writes
them to standard output. `WithPrettyPrint` formats the output as indented JSON
so it is easier to study.

We then created a processor and gave it the exporter:

```go
processor := sdktrace.NewSimpleSpanProcessor(exporter)
```

A `SpanProcessor` receives lifecycle notifications about spans. When a span
ends, `SimpleSpanProcessor` immediately gives the completed span to its
exporter.

Finally, we connected the processor to the provider:

```go
provider := sdktrace.NewTracerProvider(
	sdktrace.WithSpanProcessor(processor),
)
```

The complete pipeline is now:

```text
TracerProvider
  |
  +-- creates and manages the span
  |
  v
span.End()
  |
  v
SimpleSpanProcessor
  |
  v
stdout SpanExporter
  |
  v
pretty-printed JSON
```

The data travels from the provider to the processor and then to the exporter.
The provider does not receive exported data back from the processor.

`SimpleSpanProcessor` exports synchronously inside `span.End()`. This makes its
behavior immediate and straightforward for learning. A production service would
usually use `BatchSpanProcessor` so exporting happens away from the
application's request path.

The program still creates only one root span. The difference is that this span
is now visible after it ends.

## Step 7: Add a child span to readRandomCharacters

Passing a context containing a span into a function does not automatically
create another span. Until this step, all three application functions ran during
the root span's lifetime, but none had a span of its own.

Inside `readRandomCharacters`, we obtained a tracer:

```go
tracer := otel.Tracer("ai-incremental-tracing")
```

This uses the same instrumentation scope name as the tracer in `run` and asks
the same globally registered provider for a tracer.

We then created and deferred the end of one span:

```go
_, span := tracer.Start(ctx, "read random characters")
defer span.End()
```

The incoming `ctx` carries the root span created in `run`. Starting a span from
that context automatically makes the new span a child of the current root span.

`tracer.Start` still returns a new context containing the child span. We use `_`
for that return value because `readRandomCharacters` does not currently call
another instrumented operation that needs the child context.

The trace now has this structure:

```text
process random characters
└── read random characters
```

The two exported spans have:

- The same trace ID, showing that they belong to one trace.
- Different span IDs, because they represent different operations.
- A parent ID on `read random characters` that matches the span ID of
  `process random characters`.
- A child count of one on the root span.

The child span is normally printed first. It ends when
`readRandomCharacters` returns, while the root span does not end until `run`
returns.

The other two functions still do not create spans.

## Step 8: Add a sibling span to reverseCharacters

We added the same basic manual instrumentation pattern inside
`reverseCharacters`:

```go
tracer := otel.Tracer("ai-incremental-tracing")
_, span := tracer.Start(ctx, "reverse characters")
defer span.End()
```

Both `readRandomCharacters` and `reverseCharacters` receive the context passed
directly from `run`. That context carries the root span, so both new function
spans use the root as their parent:

```text
process random characters
├── read random characters
└── reverse characters
```

The two function spans are siblings. Although `reverseCharacters` executes after
`readRandomCharacters`, sequential execution does not make the read span its
parent. The context supplied to `tracer.Start` determines the parent.

The output should now show:

- Three spans in total.
- One trace ID shared by all three spans.
- A unique span ID for each span.
- The same parent ID on both function spans.
- That parent ID matching the root span's span ID.
- `ChildSpanCount: 2` on the root span.

`saveCharacters` is now the only application function without its own span.

## Step 9: Add a sibling span to saveCharacters

We instrumented the final application function using the same pattern:

```go
tracer := otel.Tracer("ai-incremental-tracing")
_, span := tracer.Start(ctx, "save characters")
defer span.End()
```

`saveCharacters` also receives the context passed directly from `run`.
Consequently, its span has the root span as its parent and is a sibling of the
read and reverse spans.

The complete trace structure is now:

```text
process random characters
├── read random characters
├── reverse characters
└── save characters
```

The output should now show:

- Four spans in total.
- One trace ID shared by all four spans.
- A unique span ID for every span.
- The same parent ID on all three function spans.
- That parent ID matching the root span's span ID.
- `ChildSpanCount: 3` on the root span.

This step only creates and ends the save span. Although `os.WriteFile` can
return an error, we have not recorded errors or set span status yet.

## Step 10: Record a save error as an exception event

Previously, `saveCharacters` returned the result of `os.WriteFile` directly:

```go
return os.WriteFile(path, []byte(characters), 0o600)
```

We expanded the error path so the save span can record the error before returning
it:

```go
if err := os.WriteFile(path, []byte(characters), 0o600); err != nil {
	span.RecordError(err)
	return err
}

return nil
```

`span.RecordError(err)` adds a timestamped event named `exception` to the span.
The event contains standardized exception information such as the Go error type
and error message.

Recording the error does not handle it. The function must still return the Go
error so the existing application error flow continues:

```text
os.WriteFile fails
  |
  +-- record the error on the save span
  |
  +-- return the Go error
  |
  +-- defer ends and exports the save span
  |
  v
run receives and returns the error
```

An important OpenTelemetry distinction is that `RecordError` does not
automatically set the span status to `Error`. After this step, a failed save span
contains an exception event while its status remains `Unset`.

This separates two pieces of information:

- The exception event records what happened at a particular time.
- Span status expresses the final interpreted outcome of the operation.

We will set error status explicitly in a later step so the two effects can be
compared.

## Step 11: Set error status on a failed save

The save span already recorded a failed file write as an exception event. We now
also mark the operation's final outcome as an error:

```go
span.RecordError(err)
span.SetStatus(codes.Error, err.Error())
return err
```

These calls communicate different information:

```text
RecordError
  |
  +-- adds a timestamped exception event
  +-- records the error type and message

SetStatus
  |
  +-- marks the span's final outcome as Error
  +-- includes a description of the failure
```

`RecordError` does not set status automatically, which is why both calls are
present.

When saving succeeds, we do not explicitly set `OK`. The default `Unset` status
does not mean failure; it means that no explicit status was assigned.

On a failed save, the exported `save characters` span should now contain both:

```text
Events:
  exception

Status:
  Code: Error
  Description: the file-writing error
```

Only the save span is marked as `Error` in this step. The root
`process random characters` span still has `Unset` status even though `run`
returns an error. Span status belongs to each individual span and is not
automatically copied to its parent.

## Step 12: Set root status when the workflow fails

A child span's status is not automatically copied to its parent. After the
previous step, a failed save produced this result:

```text
process random characters — Status: Unset
└── save characters      — Status: Error
```

However, `run` stops and returns an error when either reading or saving fails.
That means the overall operation represented by the root span has also failed.

For each error path in `run`, we now build the error that will be returned and
set the root span's status:

```go
operationErr := fmt.Errorf("save reversed characters: %w", err)
span.SetStatus(codes.Error, operationErr.Error())
return operationErr
```

The same pattern is used for a failure returned by `readRandomCharacters`.

On a save failure, the trace now communicates two levels of meaning:

```text
process random characters — Status: Error
└── save characters      — Status: Error
    └── exception event containing the file-writing error
```

We do not call `RecordError` again on the root span in this step. The detailed
exception event remains on the child span where the error occurred. The root
status communicates that the child failure caused the complete workflow to
fail.

This status decision depends on the application's outcome. If a child operation
failed but the workflow recovered and completed successfully, the root would not
necessarily be marked as `Error`.

## Step 13: Record and mark errors in readRandomCharacters

`readRandomCharacters` has two operations that can fail:

```go
file, err := os.Open(path)
```

and:

```go
_, err := io.ReadFull(file, randomBytes)
```

For both failure paths, we now record the exception and set the read span's
status before returning the Go error:

```go
span.RecordError(err)
span.SetStatus(codes.Error, err.Error())
return "", err
```

A read failure now produces tracing information at two levels:

```text
process random characters — Status: Error
└── read random characters — Status: Error
    └── exception event containing the read error
```

The read span records where the error originated. After the function returns the
error, `run` marks the root span to show that this failure also caused the
complete workflow to fail.

The two functions that naturally return errors now handle tracing consistently:

```text
read random characters
  +-- records its own exception
  +-- marks its own status as Error
  +-- returns the Go error

save characters
  +-- records its own exception
  +-- marks its own status as Error
  +-- returns the Go error

process random characters
  +-- marks the overall workflow as Error when either error is returned
```

OpenTelemetry records the failure, but normal Go error returns still control the
application's behavior.

## Step 14: Add an attribute to the read span

We added one attribute describing the number of random characters requested:

```go
span.SetAttributes(attribute.Int("random.character.count", count))
```

An attribute is a key-value property of a span. The exported
`read random characters` span now contains information resembling:

```text
Attributes:
  random.character.count = 10
```

The telemetry concepts used so far have different purposes:

```text
Span name
  +-- identifies the operation
  +-- "read random characters"

Attribute
  +-- describes a property of the operation
  +-- random.character.count = 10

Event
  +-- records something that happened at a specific time
  +-- exception

Status
  +-- describes the final outcome
  +-- Unset or Error
```

Attributes allow tracing backends to filter and group spans. For example, an
operator could search for read spans requesting a particular number of
characters.

We record the requested count, but not the generated random string itself.
Telemetry should avoid unnecessary content, secrets, personally identifiable
information, and highly variable values.

## Step 15: Configure the service Resource

Until this step, the SDK generated a default resource containing a fallback
service identity:

```text
service.name = unknown_service:ai-incremental-tracing
```

We created a resource with an intentional service name and version:

```go
res, err := resource.New(
	ctx,
	resource.WithAttributes(
		semconv.ServiceName("random-character-processor"),
		semconv.ServiceVersion("1.0.0"),
	),
)
if err != nil {
	return fmt.Errorf("create resource: %w", err)
}
```

We used OpenTelemetry semantic-convention helpers for the attribute names.
Semantic conventions give common telemetry concepts consistent names across
applications, languages, and backends.

We then attached the resource to the provider:

```go
provider := sdktrace.NewTracerProvider(
	sdktrace.WithSpanProcessor(processor),
	sdktrace.WithResource(res),
)
```

A resource describes the entity producing telemetry and is shared by every span
created by that provider:

```text
Resource:
  service.name = random-character-processor
  service.version = 1.0.0

Spans:
  process random characters
  read random characters
  reverse characters
  save characters
```

This differs from a span attribute:

```text
Resource attribute
  +-- describes who produced the telemetry
  +-- service.name = random-character-processor

Span attribute
  +-- describes one particular operation
  +-- random.character.count = 10
```

The tracer's instrumentation scope name remains
`ai-incremental-tracing`. The instrumentation scope identifies the code creating
the spans, while the resource identifies the running service.

## Step 16: Add attributes to the save span

We added two attributes describing the file-writing operation:

```go
span.SetAttributes(
	attribute.String("file.path", path),
	attribute.Int("file.size_bytes", len(characters)),
)
```

The exported `save characters` span now includes information resembling:

```text
Attributes:
  file.path = reversed-random-characters.txt
  file.size_bytes = 10
```

These attributes describe one particular operation, unlike the service resource
attributes shared by every span:

```text
Resource attributes:
  service.name
  service.version

Save span attributes:
  file.path
  file.size_bytes
```

The attributes are set before attempting the write, so they are available on
both successful and failed save spans.

The path in this study program is fixed and contains no sensitive information.
Production telemetry should avoid recording paths containing usernames,
credentials, identifiers, or highly variable values.

## Step 17: Add a successful file-write event

After `os.WriteFile` succeeds, we add an event to the save span:

```go
span.AddEvent("file write completed")
```

An event records that something notable happened at a particular moment during
a span. The exported event includes its name and timestamp.

The event is placed after the error path, so the two outcomes differ:

```text
Successful save:
  save characters span
  └── event: file write completed

Failed save:
  save characters span
  └── event: exception
```

Events are optional and are stored inside their enclosing span. An event does
not create a separate span and does not appear as another operation in the trace
hierarchy.

The concepts can now be compared as:

```text
Span
  +-- represents an operation with a duration

Attribute
  +-- describes a property of that operation

Event
  +-- records something that happened at a specific time during the operation
```

## Step 18: Replace the simple processor with a batch processor

We replaced:

```go
processor := sdktrace.NewSimpleSpanProcessor(exporter)
```

with:

```go
processor := sdktrace.NewBatchSpanProcessor(exporter)
```

The exporter and provider wiring remain unchanged:

```text
Before:
  TracerProvider
  └── SimpleSpanProcessor
      └── stdout exporter

After:
  TracerProvider
  └── BatchSpanProcessor
      └── stdout exporter
```

With the simple processor, ending a span synchronously called the exporter.
With the batch processor, ending a span normally places it into an in-process
queue. A background worker exports queued spans in batches, keeping exporter
work away from the application path.

This makes the existing provider shutdown especially important:

```go
defer provider.Shutdown(ctx)
```

During normal shutdown, the provider asks the batch processor to export spans
that are still waiting in its queue. Without that shutdown, this short-lived
program could terminate before its spans were exported.

The stdout exporter is still used for learning. Only the method used to process
and deliver completed spans to it has changed.

## Step 19: Name and version the instrumentation scope

The same tracer name was repeated everywhere a tracer was obtained. We defined
the instrumentation scope identity once:

```go
const (
	instrumentationName    = "ai-incremental-tracing"
	instrumentationVersion = "1.0.0"
)
```

We added a small helper that obtains a tracer using that identity:

```go
func newTracer() trace.Tracer {
	return otel.Tracer(
		instrumentationName,
		trace.WithInstrumentationVersion(instrumentationVersion),
	)
}
```

Every operation now calls:

```go
tracer := newTracer()
```

This avoids repeated strings and ensures that all spans created by this
instrumentation have a consistent scope name and version.

The exported spans now distinguish two identities:

```text
Resource:
  service.name = random-character-processor
  service.version = 1.0.0

Instrumentation scope:
  name = ai-incremental-tracing
  version = 1.0.0
```

The resource identifies the running service. The instrumentation scope
identifies the code or library responsible for creating its spans. Their
versions may happen to match in this small program, but they represent different
concepts.

## Step 20: Detect host resource information

The service resource already contained manually supplied attributes. We added a
resource detector:

```go
res, err := resource.New(
	ctx,
	resource.WithHost(),
	resource.WithAttributes(
		semconv.ServiceName("random-character-processor"),
		semconv.ServiceVersion("1.0.0"),
	),
)
```

`resource.WithHost()` discovers host information from the runtime environment.
This avoids hardcoding the host name into the application.

The resource now combines detected and manually supplied information:

```text
Detected:
  host.name

Manually supplied:
  service.name
  service.version
```

Because the resource belongs to the provider, the resulting host information is
attached to every span produced by that provider.

Resource detectors can also be used to discover process, operating-system,
container, Kubernetes, and cloud-platform information. The exact detectors
available depend on the SDK and additional OpenTelemetry packages being used.

As with all telemetry, detected metadata should be reviewed for privacy and
operational requirements before being enabled in production.

### Addendum: Other resource detectors

OpenTelemetry provides many other resource detectors, including detectors for
process and runtime information, operating systems, containers, environment
variables, Kubernetes, and cloud platforms such as AWS, Azure, and GCP.

It is up to each team to select the detectors relevant to its runtime and
operational needs. Avoid enabling every detector automatically: first consider
whether the resulting metadata is useful and whether it could expose sensitive
information such as command-line arguments, usernames, identifiers, or
filesystem paths.
