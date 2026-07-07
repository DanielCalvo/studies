## Chapter 6 - Monitoring Distributed Systems
- Your monitoring system should address two questions: What's broken, and why?
- Whitebox monitoring is essential for debugging

- Four golden signals:
    - Latency: How long it takes to serve a request
    - Traffic: How much demand is being placed on your system
    - Errors: The rate of request that fails
    - Saturation: How much of your system resources are being used
- Measuring all 4 signals and alerting a human when one of them is problematic is a good start
- Worrying about your tail: Careful with averages. Some of your requests might take very long, but not show on an average. Historigrams per exponential boundaries are useful to track these

- There are some really really good rules on paging and how to page here.
- Pages (as in, being paged) should be actionable, require intelligence, and possibly be about a novel problem

As a general rule, the chapter advises to keep it simple. Trying to get too fancy with monitoring might make it too fragile, or too time and money expensive.
