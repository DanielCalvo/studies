# OpenTelemetry Auto-Instrumentation in Go

## Question

How does OpenTelemetry auto-instrumentation work in Go? Does initializing or importing OpenTelemetry in `main()` automatically instrument the program, or is automatic instrumentation limited to specific supported packages? For example, can a simple program using only the standard library automatically trace writing `"Hello, world!"` to a file?

## Answer

Automatic instrumentation is limited to operations that the instrumentation tool knows about. OpenTelemetry as a whole is not limited to those operations, however, because any Go code can be instrumented manually.

There are three related concepts that can easily be confused.

### 1. Initializing the OpenTelemetry SDK

OpenTelemetry can be imported and initialized in `main()` to configure things such as:

- A `TracerProvider`
- An exporter
- The service name
- Sampling
- Shutdown and flushing

This creates the telemetry pipeline, but it does not automatically discover functions or create spans around them.

```text
SDK initialization = somewhere for telemetry to go
instrumentation     = code that creates the telemetry
```

### 2. Go instrumentation libraries

An instrumentation package understands the API of a particular library and creates telemetry around its operations. For example, `otelhttp` understands `net/http` and can wrap an HTTP handler:

```go
handler := otelhttp.NewHandler(myHandler, "my-handler")
```

The original library does not necessarily need built-in OpenTelemetry support. A separate instrumentation package can understand and instrument it.

Because Go does not normally support monkey-patching functions at runtime, merely importing an instrumentation package does not modify all calls to the corresponding library. Handlers, transports, database drivers, or other components generally need to be wrapped or configured explicitly.

These packages instrument the supported library operations, such as HTTP requests or database calls. They do not automatically instrument the application's own functions or business logic.

See [Using instrumentation libraries in Go](https://opentelemetry.io/docs/languages/go/libraries/).

### 3. Zero-code instrumentation

Go also has more automatic approaches that do not require changing the application's source code.

#### eBPF instrumentation

An external eBPF agent observes supported functions, protocols, and operating-system activity while the program runs. It can recognize operations such as HTTP, gRPC, and certain database calls, but it cannot infer arbitrary application behavior or business meaning.

See [OpenTelemetry eBPF Instrumentation](https://opentelemetry.io/docs/zero-code/obi/).

#### Compile-time instrumentation

The `otelc` tool wraps `go build` and injects instrumentation hooks into recognized functions as the program is compiled.

It supports a defined set of packages and frameworks, including examples such as:

- `net/http`
- `database/sql`
- gRPC
- Gin
- Redis
- MongoDB
- Kafka
- `log/slog`

It does not automatically instrument every function in the program.

See [Go compile-time instrumentation](https://opentelemetry.io/docs/zero-code/go/compile-time/) and its [supported libraries](https://opentelemetry.io/docs/zero-code/go/compile-time/supported-libraries/).

## The `Hello, world!` file example

Consider this program:

```go
package main

import "os"

func main() {
    _ = os.WriteFile("hello.txt", []byte("Hello, world!\n"), 0o644)
}
```

Current Go auto-instrumentation should not be expected to produce an application span such as:

```text
write hello.txt
```

The standard-library file operation is not among the currently supported compile-time instrumentation targets. A local file write is also not the kind of network or protocol operation that eBPF instrumentation would ordinarily turn into an application span.

This does not mean `os.WriteFile` needs to be replaced with a special third-party package. The operation can be described with a manual span:

```go
ctx, span := tracer.Start(context.Background(), "write hello file")
defer span.End()

err := os.WriteFile("hello.txt", []byte("Hello, world!\n"), 0o644)
if err != nil {
    span.RecordError(err)
    span.SetStatus(codes.Error, "failed to write file")
}
```

The application would also need an initialized telemetry pipeline and must flush or shut down its `TracerProvider` before exiting.

## Conclusion

It is accurate to say:

> Go automatic instrumentation can only create detailed spans for packages, functions, and protocols that its instrumentation implementation explicitly supports.

It would not be accurate to say:

> OpenTelemetry can only be used with supported libraries.

Any application operation can be instrumented manually. A typical Go application therefore combines:

- Automatic or library instrumentation for HTTP, gRPC, SQL, messaging, and similar operations
- Manual spans for application-specific behavior
- Manual attributes and events where business context matters

Automatic instrumentation provides broad technical coverage, while manual instrumentation describes what the application is actually trying to accomplish.
