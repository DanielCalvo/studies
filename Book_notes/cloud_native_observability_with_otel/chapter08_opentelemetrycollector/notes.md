# The purpose of the OpenTelemetry Collector
The OpenTelemetry Collector is a process that receives telemetry in various formats, processes it, and then exports it to one or more destinations

These are some reasons why the OpenTelemetry Collector may be helpful:

- You decouple the source of telemetry from its destination, so you can just configure a destination for the telemetry data and allow the operators of the Collector to determine where it needs to go
- You can also provide a single destination for many data types. The Collector can receive traces, metrics, and logs. Really cool!

The Collector allows users to configure pipelines for each signal separately.
You can configure any number of receivers, processors, and exporters.

Apparently, you can have multiple protocols for the receivers, so the book alleges you can receive traces, metrics, and logs from Kafka, which is really cool. There's a table on page 236.

## Processors
Processors can do some tasks like filtering unwanted telemetry or injecting additional attributes.
For instance, I imagine a processor could inject a tag for the environment that this metric is coming from, like staging or production.

So you can configure processing pipelines in the processor config, which can do things with your tracing data or metadata, but sequentially, so the ordering of the config does matter.

So it looks like there are attributes that you can change, and there's batch filtering, memory limiting, probabilistic sampling, and a bunch of other things.

Do note that some of these only apply to traces, metrics, or logs. Like, not every operation is supported for every resource type.

### Attributes processor
There is also an attributes processor that can modify telemetry data attributes.

You can delete an attribute:

- Extract an attribute and process it with a regular expression
- Calculate a hash
- Insert a given attribute for a specified key if it doesn't exist
- Update an existing attribute with a specified value
- Upsert, which means combining the functionality of insert and update: if an attribute doesn't exist, it will be inserted with a specific value

You can also match or not match spans based on match type.

This would be interesting if you want to scrub PII.

### Filter processor
- This allows you to include or exclude telemetry data based on configured criteria. You can configure names with strict or regexp matching

### Probabilistic sampling processor
- This is discussed later in chapter 12
- However, it is important to know that you have a probabilistic sampling processor that can be used to reduce the number of traces that are exported from the Collector by specifying a sampling percentage
- This will determine what percentage of traces you want to keep
- There's also a hash seed configuration that you need to pay attention to, to make sure that you sample from multiple Collectors at the same rate when this is enabled

### Resource processor
- This lets you modify attributes just like the attributes processor
- However, instead of updating attributes on spans, metrics, or logs, this updates the attributes of the resource associated with the telemetry data
- Oh, so this is what actually would allow you to upsert a deployment environment to be something like staging

### Span processor
- You can also manipulate the names of spans and or attributes of spans based on their names. This is the job of the span processor. It can extract attributes from a span and update its name based on those attributes
- Alternatively, it can take the span name and expand it to individual attributes associated with the span
- Right, apparently, you can rename a span based on, like, a store ID or something, which is handy

### Batch processor
- You can configure the Collector to send out every 10 seconds or every 10,000 records and batch things out instead of sending them synchronously
- It is recommended to configure a batch processor for all the pipelines to optimize the throughput of the Collector

### Memory limiter processor
- There is also a memory limiter processor that lets users control the amount of memory the Collector uses, so as to avoid the Collector running out of memory or something like that
- There's also a ballast extension that allows the Collector to preallocate memory to improve the stability of the heap
- The book recommends you set up the memory limiter as the first processor in the pipeline

# Exporters
- The exporter takes data in its internal Collector format, marshals it into an output format, and sends it to one or more configured destinations
- One interesting thing is that you can actually export traces, metrics, and logs to files, which is interesting
- So you can put those files anywhere for long-term storage, one would assume
- You can also export to Kafka, which I always find really cool

## Additional components
- There is a repository for the community with additional components that can be found in OpenTelemetry Collector Contrib
- People can contribute additional receivers, processors, and exporters in there, so maybe it's worth a look sometime
