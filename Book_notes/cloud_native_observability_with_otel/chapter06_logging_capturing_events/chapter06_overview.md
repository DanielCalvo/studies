# Chapter 6: Logging – Capturing Events

## What the chapter covers

Chapter 6 presents logging as OpenTelemetry's third core signal. Its main concepts are:

- The logging pipeline: provider, logger/emitter, processor, and exporter
- Producing log records directly through OpenTelemetry
- Integrating OpenTelemetry with a language's established logging API
- Log timestamps, bodies, severity numbers and text, and attributes
- Associating Resources and instrumentation scopes with logs
- Correlating logs with traces through trace IDs and span IDs
- Capturing framework-generated logs through instrumentation

The book's logging examples use an old experimental Python API. For a current Go
exercise, we should verify the installed OpenTelemetry Go logging API and use Go's
standard `log/slog` package with the OpenTelemetry bridge where appropriate.

## Proposed incremental coding exercise

Create an `ai_incremental_logging` Go program containing a very small HTTP service:

```text
GET /hello/{name} -> "Hello, {name}"
```

The ordinary starting program should only return the greeting. It should contain no
OpenTelemetry configuration and no logging beyond errors required to start the server.

This gives later logging steps concrete events to describe:

- A request was received.
- A greeting was produced.
- A request failed.

It is small enough that logging remains the focus. As the exercise grows, the standard
Go logger can produce structured records, an OpenTelemetry logging pipeline can export
those records, and a request span can eventually demonstrate trace and log correlation.

## First decision

Confirm whether the single `/hello/{name}` endpoint is a suitable ordinary starting
program. Once confirmed, create only that program and an
`ai_incremental_steps_taken.md` file before introducing logging.
