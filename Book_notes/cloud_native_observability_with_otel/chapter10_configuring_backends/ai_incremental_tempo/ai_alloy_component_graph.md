# Alloy's component graph

As a rule of thumb, Alloy configurations connect components by passing one
component's output to another component's compatible input.

Our trace pipeline is linear:

```text
receiver -> batch processor -> exporter
```

In the Alloy configuration, each reference creates the next connection:

```alloy
otelcol.receiver.otlp "applications" {
  output {
    traces = [otelcol.processor.batch.traces.input]
  }
}

otelcol.processor.batch "traces" {
  output {
    traces = [otelcol.exporter.otlp.tempo.input]
  }
}
```

Conceptually:

```text
receiver output -> batch input
batch output    -> exporter input
```

## It is a graph, not necessarily one chain

An Alloy pipeline does not always have to be linear. One output can branch to
multiple destinations:

```text
                 +-> Tempo exporter
receiver -> batch
                 +-> debug exporter
```

Multiple upstream components can also feed a shared downstream component:

```text
receiver A --+
             +-> batch -> exporter
receiver B --+
```

This allows Alloy configurations to represent branching, merging, and more
complex telemetry flows.

## Useful rules of thumb

- Receivers and sources usually begin a flow.
- Processors optionally transform, filter, sample, enrich, or batch data.
- Exporters and write components usually terminate a flow.
- Connections must carry compatible types: traces connect to trace inputs,
  metrics to metric inputs, and so forth.
- Components that are not connected into a useful graph do not participate in
  that telemetry path.
- Not every Alloy component carries telemetry directly. Some components
  produce discovered targets, authentication handlers, configuration values,
  or other typed outputs.

The broader Alloy model is therefore:

```text
component output -> compatible component input
```

This component-graph model is the major structural difference from the
OpenTelemetry Collector's centralized `service.pipelines` configuration.

## Reference

- [How Grafana Alloy works](https://grafana.com/docs/alloy/latest/introduction/how-alloy-works/)
