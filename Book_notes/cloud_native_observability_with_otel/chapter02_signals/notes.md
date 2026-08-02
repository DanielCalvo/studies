# 2 OpenTelemetry signals - Traces, metrics, and logs

### Traces
A distributed trace is a series of event data generated at various points throughout a system. This is tied together via a unique identifier!

Each trace represents a unique request through a system that can be either synchronous or asynchronous.

Each operation recorded in a trace is represented by a span. A span is a single unit of work done in the system.

#### Anatomy of a trace
- A unique identifier referred to as a trace ID identifies the request through the system
- There is a span ID associated with the span that last interacted with the context. This may also be referred to as the parent identifier
- Trace flags include additional information about the trace, such as the sampling decision and trace level
- There is also a trace state field. This is for individual vendors to propagate information necessary for their systems to interpret their tracing data

A span can represent a method call or subset of the code being called within a method

The first span in a trace is called the root span and is identified because it does not have a parent span identifier

A span has a unique identifier, a parent span identifier, a name describing the work being recorded, and a start and end time
Additionally, spans can contain metadata in the form of key-value pairs.

Exemplars enable a metric to contain information about an active span. This would be a really cool thing to trial!

Logs recorded via OpenTelemetry contain the trace ID and span ID for any span active at the time of the event. This is really cool so you can correlate the logs with the traces!

There is also a schema URL that allows the producers and consumers of telemetry to understand how to interpret the data -- Maybe this part you need to look it up a little bit to see some examples
