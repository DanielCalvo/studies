# Chapter 8 incremental Collector study plan

This is an adaptable plan for exploring the configuration and concepts from
Chapter 8 locally. It is a guide rather than a rigid commitment: we can revise,
split, combine, or omit steps when our observations suggest a better learning
path.

We will not initially introduce Kubernetes, Alloy, Tempo, Loki, Prometheus, or
another telemetry backend.

The local learning environment will follow this general data flow:

```text
Test telemetry or a small Go application
                    |
                    | OTLP
                    v
        OpenTelemetry Collector
                    |
                    v
             terminal or file
```

Each implementation step should be introduced separately, commented, tested,
and recorded in `ai_incremental_steps_taken.md`.

## Step 1: The smallest complete Collector pipeline

Configure:

```text
OTLP receiver -> traces pipeline -> debug exporter
```

Run the Collector locally through Docker and send it a small test trace.

This demonstrates:

- The structure of a Collector configuration
- Receivers as pipeline inputs
- Exporters as pipeline outputs
- `service.pipelines` as the wiring that enables components
- The fact that merely declaring a component does not activate it

Expected result: one trace appears in the Collector's terminal output.

Receivers, pipelines, and exporters must appear together because none of them
can produce a useful result alone.

## Step 2: OTLP gRPC versus OTLP HTTP

Enable both standard OTLP transports:

```text
OTLP/gRPC -> port 4317
OTLP/HTTP -> port 4318
```

Send equivalent telemetry through each transport.

This demonstrates that OTLP defines the telemetry data model and transport
behavior, while gRPC and HTTP are different ways of carrying it.

Expected result: the Collector accepts telemetry through either endpoint and
produces equivalent output.

## Step 3: Connect a real Go application

Replace the stdout trace exporter in a small copy of one of the previous Go
examples with an OTLP exporter:

```text
Before:

Go SDK -> stdout exporter

After:

Go SDK -> OTLP exporter -> Collector -> debug exporter
```

This is the chapter's central architectural change: the application no longer
needs to know about the final backend.

Expected result: the same spans previously seen in the application terminal now
appear in the Collector terminal.

## Step 4: Add the other signal pipelines

Add metrics and logs incrementally:

```text
OTLP receiver
    |-- traces pipeline -> debug
    |-- metrics pipeline -> debug
    `-- logs pipeline -> debug
```

This demonstrates:

- Pipelines are signal-specific.
- A receiver can be reused by multiple pipelines.
- An exporter can be reused by multiple pipelines.
- Data does not cross automatically from one signal pipeline into another.

Expected result: the debug exporter identifies trace, metric, and log records
separately.

## Step 5: Add the batch processor

Change a pipeline from:

```text
receiver -> exporter
```

to:

```text
receiver -> batch -> exporter
```

Initially give the batch processor a noticeable timeout so its behavior is easy
to observe.

This demonstrates:

- Processors sit between receivers and exporters.
- Batching combines telemetry for more efficient export.
- Processor configuration and pipeline activation are separate.
- The same processor configuration can be used in several pipelines.

Expected result: multiple telemetry items appear together after a short delay
rather than being exported individually.

## Step 6: Add a telemetry attribute centrally

Use an attributes processor to add something such as:

```text
study.chapter = "8"
```

The application will not set this attribute; the Collector will.

This demonstrates that a Collector can enrich, normalize, hash, remove, or
replace attributes without changing application code.

Expected result: the new attribute appears on telemetry printed by the debug
exporter.

## Step 7: Modify resource attributes

Add a resource processor that sets something such as:

```text
deployment.environment = "local-study"
```

This demonstrates the distinction between:

```text
Span, log, or data-point attribute
    Describes an individual operation or record

Resource attribute
    Describes the entity that produced the telemetry
```

Expected result: `deployment.environment` appears on the resource rather than
among an individual span's attributes.

## Step 8: Transform a span and demonstrate processor ordering

The book uses its old `span` processor to construct a span name from attributes.
Reproduce the idea using the current transform processor and OpenTelemetry
Transformation Language.

The pipeline will resemble:

```text
add attribute -> build span name from attribute -> export
```

Then reason about what would happen if those processors were reversed.

This demonstrates:

- Processors execute serially.
- Processor order can change the result.
- Collector transformations can centrally standardize telemetry.
- Excessively dynamic span names can create cardinality problems, so renaming
  must be done carefully.

Expected result: the exported span has a modified name derived from an
attribute.

Current reference:
[Transform processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/processor/transformprocessor/README.md)

## Step 9: Send the same telemetry to multiple exporters

Add a local file exporter alongside the debug exporter:

```text
                     |-- debug exporter -> terminal
trace pipeline ------|
                     `-- file exporter  -> file
```

This demonstrates exporter fan-out: a single pipeline may send the same
telemetry to more than one destination.

Expected result: the same trace is visible in both the Collector terminal and a
local file.

This represents what a real deployment could do with multiple backends without
requiring us to install any backend.

## Step 10: Add the host metrics receiver

Add an active receiver:

```text
host_metrics receiver -> metrics pipeline -> debug exporter
```

This demonstrates an important receiver distinction:

```text
OTLP receiver
    Passively listens for telemetry pushed by applications

Host metrics receiver
    Actively gathers system measurements
```

Begin with memory and network metrics.

Expected result: metrics appear periodically even when no application sends
anything.

Current reference:
[Host metrics receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/receiver/hostmetricsreceiver/README.md)

## Step 11: Limit data at the receiver

Configure the network scraper to include only a selected interface.

This demonstrates filtering at collection time:

```text
Do not collect unwanted data in the first place
```

Expected result: network metrics are produced only for the selected interface.

The exact interface will be selected after inspecting the local environment
rather than copying the book's macOS-specific `lo0` value.

## Step 12: Filter data with a processor

Add a filter processor to remove a noisy metric after collection:

```text
host metrics receiver -> filter -> debug exporter
```

This demonstrates a different kind of reduction:

```text
Receiver selection
    Controls what gets collected

Filter processor
    Receives telemetry and decides what continues downstream
```

Expected result: the selected metric disappears from exported output while the
other host metrics remain.

Use the current OTTL-based filter syntax rather than the book's older matching
syntax.

## Step 13: Add the memory limiter safely

Place the memory limiter first:

```text
receiver -> memory limiter -> batch -> other processors -> exporter
```

This demonstrates:

- The Collector itself has finite resources.
- A memory limiter applies backpressure when memory pressure becomes excessive.
- Its position matters because errors should propagate toward receivers.
- Batching improves export efficiency, while memory limiting protects the
  Collector.

Do not deliberately exhaust the computer's memory. This step should demonstrate
valid configuration and ordering without manufacturing an unsafe out-of-memory
condition.

The book's recommendation to use the old `memory_ballast` extension is dated.
Modern Go and Collector deployments use mechanisms such as `GOMEMLIMIT`.
Document that historical change rather than adding the removed ballast
component.

## Step 14: Configure an extension

Enable a health-check extension:

```text
Collector pipelines
        +
health-check HTTP endpoint
```

This demonstrates that extensions are not inserted into telemetry pipelines.
They add operational capabilities to the Collector itself.

Expected result: a local HTTP request returns the Collector's health.

The important distinction is:

```text
Receiver, processor, or exporter
    Participates in telemetry flow

Extension
    Adds Collector-level functionality
```

## Step 15: Final configuration and Alloy mapping

Once each piece is understood, assemble the final local configuration and
create a concept map:

```text
sources
   |
   v
receivers
   |
   v
memory limiter
   |
   v
batching
   |
   v
enrichment, transformation, and filtering
   |
   v
exporters
```

Then show how each vanilla Collector component maps to Alloy:

```text
Collector YAML                 Alloy

otlp receiver                  otelcol.receiver.otlp
batch processor                otelcol.processor.batch
attributes processor           otelcol.processor.attributes
debug exporter                 otelcol.exporter.debug
component pipeline             component references and forwarding
```

This final step will be explanatory. Do not deploy Alloy or Kubernetes yet.

## Deliberately deferred topics

Leave these for later chapters or a future Kubernetes exercise:

- Tempo, Loki, Prometheus, and Mimir
- Kubernetes deployment patterns
- Sidecar, DaemonSet, agent, and gateway placement
- TLS and authentication
- Persistent queues and production retry tuning
- Kafka
- Horizontal scaling
- Probabilistic and tail sampling, which Chapter 12 covers properly
- Advanced OTTL transformations
- Collector-to-Alloy migration

## First implementation step

The first implementation step will be only:

```text
OTLP receiver -> traces pipeline -> debug exporter
```

This provides the Collector's essential mental model before anything else is
added.
