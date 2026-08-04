# Forms of Service-Level Indicators

A service-level indicator is a carefully defined quantitative measurement of
some behavior that matters to users. It does not have to be a percentage.

The unit should follow the behavior being measured.

## Examples

| SLI form | Example |
| --- | --- |
| Proportion | 99.95% of eligible requests succeed |
| Threshold proportion | 99% of eligible requests complete within 5 seconds |
| Duration | The oldest queued job is 45 seconds old |
| Freshness | Published data is 8 minutes old |
| Throughput | The pipeline processes 1,000 records per second |
| Latency percentile | p99 request latency is 4.2 seconds |
| Queue depth | 750 jobs are waiting |
| Durability proportion | One object is lost per billion objects stored |
| Correctness proportion | 99.999% of results match the expected output |

The SLO places a target on the chosen SLI:

```text
SLI: oldest queued job age
SLO: oldest queued job age must remain below 10 minutes
```

```text
SLI: data freshness
SLO: published data must be no more than 30 minutes old
```

```text
SLI: successful-request percentage
SLO: at least 99.9% of eligible requests must succeed
```

## Why percentages are common

Request-oriented services frequently classify individual events as good or bad:

```text
good eligible events
--------------------
total eligible events
```

This naturally produces a percentage and maps cleanly to an error budget:

```text
error budget = eligible events - required good events
```

For latency, compare these formulations:

```text
p99 latency must remain below 5 seconds
```

```text
99% of eligible requests must complete within 5 seconds
```

Both can be valid. The second is often easier to use for an error budget because
every request is explicitly classified as good or bad.

## Image Resizer examples

The Image Resizer's important user-facing behaviors can all be represented as
event proportions.

Availability and correctness:

```text
eligible requests returned correctly
------------------------------------
      total eligible requests
```

Latency among correct responses:

```text
correct responses completed within 5 seconds
--------------------------------------------
         total correct responses
```

Primary composite good-request SLI:

```text
eligible requests returned correctly within 5 seconds
-----------------------------------------------------
               total eligible requests
```

These measurements answer different questions:

- Availability and correctness: did eligible requests work?
- Latency: were correct responses fast?
- Composite good-request SLI: did users receive results that were both correct
  and fast?

## Internal signals versus user-facing SLIs

A measurement can be operationally useful without being a good user-facing SLI.
CPU utilization, heap size, pod count, and queue depth can help explain service
behavior, but users normally care about successful, correct, and timely results.

Queue depth illustrates the distinction. A target such as fewer than 500 queued
jobs is measurable, but it may be an internal saturation target rather than a
direct user promise. Five hundred small jobs could clear quickly, while ten
expensive jobs could wait a long time. Oldest-job age may therefore represent the
user experience more directly.

The practical rule is:

> SLIs do not have to be percentages. Request-oriented SLIs are frequently
> expressed as proportions because they are understandable, aggregatable, and
> easy to connect to error budgets. The best unit is the one that most directly
> represents the behavior promised to users.
