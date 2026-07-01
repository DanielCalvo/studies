Right, I think we should do chapters 1, 2, and 3. Those are the most useful, I feel.

One thing I found interesting from the Prometheus Up & Running book is that it briefly touches on tracing and it says that rather than doing snapshots of stack traces at points of interest, some tracing systems trace and record every function call below the function of interest. So like 1 in 100 requests might get sampled as opposed to all requests. And then for those requests, you might be able to see how much time was spent talking to backends such as databases and caches. Interesting.

So this could allow you to see how much faster you are when you have a cache hit versus a cache miss.

Right. It then says that distributed tracing takes this a step forward by attaching a unique ID to requests, and these are passed from one process to the other so you're able to string together traces from different processes and machines. You can stitch them back together by request ID. Nice. So you can troubleshoot your service-based architecture.

Page 12 has an interesting diagram of Prometheus architecture. I think what you're interested more in here is service discovery, scraping, and rules and alerts, which is actually the config of Prometheus.

There is obviously service discovery and as you know Prometheus has integrations with EC2, Kubernetes, and Consul!

Then you can use relabeling with your service discovery and target so you can see where the metrics came from.

Prometheus has recording rules, which means it runs PromQL queries on a regular basis and ingests their results into the storage engine. So it's like pre-calculated queries for vectors or something like that. Nice.

The way Alertmanager works is it receives alerts from Prometheus and turns them into notifications, which can include email or Slack messages or sending those to PagerDuty.

For long-term storage, Prometheus just says that there are other systems that take this role but doesn't elaborate.

## Chapter 2: Getting started with Prometheus

Remember that Prometheus has a targets page on the UI that allows you to see which targets you're scraping!

Metrics like `process_resident_memory_bytes` are called gauges. Gauges' current absolute value is what matters.

There's a second core type of metric called counter. Counters track how many events have happened or the total size of all events.

For counters you can use a query like `rate(prometheus_tsdb_head_samples_appended_total[1m])`.

This keeps track of the rate that the counter is increasing by minute.

Adding a Node Exporter target makes that available, so you can query Node Exporter metrics like: `process_resident_memory_bytes{job="node"}`.

(`job="node"` is called a label matcher.)

Also: `rate(node_network_receive_bytes_total[1m])`

## Alerting

So there are two parts to alerting. First, you need to add alerting rules to Prometheus defining what constitutes an alert.

And second, you need to configure Alertmanager to convert these firing alerts into notifications like email, pages, and chat messages.

## Chapter 3: Instrumentation

Going over metric types real quick:

So the counter is the type of metric you'll probably use more often. It tracks either the number or size of events. In other words, it counts how many times something has occurred, like visits to your service.

You usually use the `rate` function, for instance if you want to see how much they have increased in the last minute: `rate(hello_worlds_total[1m])`

Remember that you are not limited by increasing counters one by one. You can count anything like megabytes of recordings processed or any other thing, as long as counters go up.

Gauges tell you the current state of something. While counters count how fast something is increasing, a gauge tells you the current value of something. A gauge can go up or down.

So like number of items in a queue, memory usage of something, number of active threads, these are all things that would be well suited to be used with a gauge.

A note on metric suffix:

- Counters all end with `_total`. It is also recommended that you include the unit of your metric, such as bytes.
- In addition, the `_count`, `_sum`, and `_bucket` suffixes have other meanings so don't add them to your metrics to avoid confusion.

Then you have the histogram type of metric, which could for instance give you the quantile of something like the 0.95 quantile of request response time. You can also pass whatever bucket numbers for histograms if the default ones are not suitable.

Bucket time series are counters, so you might need to use the `rate` function.

Histograms will also create the `_count` metric, counting the number of observations, which in some cases could be the number of requests for whatever the metric is measuring.

The `_sum`, on the other hand, means the total accumulated request latency in seconds, so it's just summing everything. For instance, you could use the sum to calculate average latency by dividing it by the count.

On unit test instrumentation, the book suggests adding unit tests only to your most crucial metrics as to not add friction to introducing certain debug metrics.

### What should I instrument?

For web services you usually want request rate, latency, and error rate. Also known as RED, for rate, errors, and duration. Not mentioned in the book but there was that four golden signals thing from Google that could help looking it up but I don't have it handy, also not in the book.

The book then goes over offline serving systems, systems that do not have someone waiting on them, like log processing systems. But even log processing systems have to be connected to something to receive these logs. They're not entirely offline, so this section is a bit confusing for me.

Anyways the author says you can use the USE method for utilization, saturation, and errors.

And then you have batch jobs which are not always running, so scraping them isn't the best fit. For those you can use Pushgateway, discussed later.

Ideally at the end of a batch job you should record how long it took to run, how long each stage took, and the time at which the job succeeded. You can then add alerts if a job hasn't succeeded recently enough, allowing you to perhaps tolerate individual batch job run failures if you want to.

The book then says you should probably instrument your libraries. So if you have a cache library, you would instrument cache hits and cache misses, and if you have a thread and worker pool library then you can instrument those as well.

### How much should I instrument?

Well, the book says that adding metrics by hand is usually fine.

The problem is when you automate the process because things can spiral to very high levels of information very quickly. Like for instance if you were to have metrics broken out by request type and HTTP path, then all the possible combinations could take a significant chunk of your resources, and that will not be free.

### What should I name my metrics?

Well, the whole naming thing is a bit of an art. Remember to add units to whatever your metric is, like bytes, seconds, and ratios.

Remember that oftentimes the user will only have the metric name to go on, so a metric name that is indicative like `http_requests_authenticated` is better than `http_requests` and just `requests`.

Avoid putting labels in metric names. Do not procedurally generate metrics or metric names; it seems that's what labels are for.

Remember that metric names are effectively a global namespace, so it's important to try to avoid collisions and indicate where a metric is coming from. Be careful also because some library names are already established, like the process library, so be cautious not to create stuff with names that already exist.

### Chapter 4: Exposition

The process of making metrics available to Prometheus is known as exposition. I wonder why this warrants an entire chapter...

Alright, well this chapter just goes over instrumenting multiple languages, which is very language dependent and now with AI it's something you can do much faster, so I'm gonna skip notes on this part.

### Pushgateway

So Pushgateway is a metric cache for service-level batch jobs and it remembers only the last push that you make for each batch job. Prometheus then scrapes these metrics from your Pushgateway and you can alert and graph them. Usually you run a Pushgateway right beside a Prometheus, so it's like an add-on.

### Metric types

You can also have a help together with your metric type to describe the metric, and you can also have a type for the metric explicitly indicating which type of metric this is.

And then the book describes labels very briefly and it's okay to have a trailing comma before closing brace with a label.

It's also possible to have a metric with no time series, like you just have the help and the type lines. Maybe no children were initialized.

The book advises against having timestamps in the exposition format.

Oh this is really cool, there is a `promtool` included with Prometheus that can check the metrics for things like mistakes or misformatting or anything like that.

## Chapter 5: Labels

Labels are key-value pairs associated with the time series that, in addition to the metric, identify them.

You have two types of labels: instrumentation and target labels.

Instrumentation labels are things that come from your instrumentation, like a type of request your application received.

Target labels identify a specific monitoring target. A target label relates to your architecture and might include something like which application this is, which data center it is on, and if it's in a development or production environment.

The book then says that different Prometheus servers run by different teams might have different views of what a team, region, or service is, so an instrumented application shouldn't try to expose these labels itself. These are more like target labels that come from service discovery and relabeling that are discussed further later in the book.

So accordingly you're not gonna find features in client libraries to add labels to themselves.

You can use PromQL to aggregate labels and then just see for instance all the GET and POST requests for job `my_job`, which is cool. So you drop labels and then you get broad indications of what is happening across systems.

The book then says that you can abuse labels for something like an enum, so you can have a label like resource state and then you can have starting, running, stopping, or terminated, which is interesting.

You can also have info metrics for annotations such as build version or other information like that, which would be useful to query on but it doesn't make sense to use as target labels.

For instance there's a Python info example that has implementation, major, minor, patch level, and version set things in it, and the actual value of the gauge is always one but you're interested in the labels and not the actual value of the gauge.

The book then says that one hint that a label is not useful is that every time you want to use the metric you find yourself needing to use the label, in which case you probably should move the label to the metric name instead, which makes sense if you think of it.

Another thing to avoid is having a time series that is a total sum of the rest of the metric, as this will break aggregation and PromQL already provides you with the ability to calculate an aggregate.

### Cardinality

Don't go too far when adding labels.

If you just add a single metric across a few thousand hosts, that's most likely fine.

But if you add a metric with a label with ten values and in addition to that it's a histogram that by default has twelve time series, that is one hundred twenty series, which is a lot. So be careful with how everything multiplies several times and how you might be using a lot of resources.

The book says that as a rule of thumb the cardinality of an arbitrary metric on one application instance should be kept below ten.

It's okay to have some metrics that have a cardinality of around a hundred, but you should be prepared to reduce metric cardinality and rely on logs as cardinality grows, which makes sense.
