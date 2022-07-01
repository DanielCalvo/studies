
## Chapter 4 - Service Level Objectives
To manage a service correctly, you must understand which behaviours matter for the service and how to measure and evaluate these behaviours

Terminology time!
### Indicators
- A (carefully) defined measured of some aspect of the service provided.
    - Most services consider latency as a key SLI
    - Others might include error rate and throughput
    - Another key indicator is availability: The fraction of the time a service is usable
    - For storage systems, durability is important: The likelyhood that data will be retained for a long period of time

### Objectives
- A target value (or range) for a service level measured by an SLI
    - Example SLO: The average search result time for the shakespeare should be less than 100ms
- Choosing and setting SLOs set expectations as to how a service will perform

### Agreements
- An explicit or implicit contract with your users that include consequences of meeting (or missing) the SLO your service has.

### Indicators in practice
- What do you and your users care about?
    - User facing services: Availability, latency and throughput
    - Storage systems: Latency, availability and durability
    - Big data systems: Throughput and end-to-end latency
    - All systems should care about correctness! (Was the correct data delivered?)

Certain indicators should be instrumented with client side collection (ex: How long it takes for a page to become usable in a browser) (I have no idea how to instrument this -- maybe on a test?)

### Aggregation
- Be careful when aggregating/averaging metrics. It's possible for the "average" of requests to be ok, but 10% of the requests to be taking much longer than expected
- Metrics are better through of distributions rather than averages
- Consider using percentiles for indicators! High variance in the time it takes for a service to respond annoys users!

### Standardize indicators
- The book recommends standardizing indicators across all projects
- Build a set of re-usable SLI templates for each common metric (like for GET requests from blackbox monitoring!)

### Objectives in practice
- Start by thinking about (or finding out) what your users care about, not what you can measure (interesting one!)
- What users care about can be difficult to measure, so you'll end up approximating what users need in some way

### Defining objectives
- SLOs should specify how they're measured and the conditions under which they're valid!
    - 99% (averaged over one minute) of Get RPC calls will complete in less than 100ms (measured across all backend servers)
      Or, simplifying a bit and using some defaults
    - 99% of GET RPC calls will complete in less than 100ms

If the shape of the performance curves are important, you can specify multiple SLO targets, like:
- 90% of GET RPC calls will complete in less than 1ms
- 99% of GET RPC calls will complete in less than 10ms
- 99.9% of GET RPC calls will complete in less than 100ms
- Neat!

If users have heterogeneous workloads like bulk process pipelines, you can define different SLOs for each workload:
- 95% of throughput clients' SET RPC calls will complete in < 1s
- 99% of latency clients SET RPC calls with payloads < 1kB will complete in < 10ms

It is both unrealistic and undesirable to set SLOs that will be met 100% of the time. It is better to allow an error budget, otherwise you'll be stiffling innovation or require expensive set ups.  
The SLO violation rate con be compared against the error budget. Read more on budgets later! (There's a chapter for it, right?)

### Choosing targets
General advice for picking SLOs:
- Keep it simple: Complicated aggregations in SLI can obscure performance and make it difficult to figure out what's going on
- Have as few SLOs as possible: Choose just enough SLOs to provide coverage of your system. Defend the ones you pick!
- Perfection can wait: You can always refine your SLOs. Start with loose SLOs, iterate/tighten if needed
- Avoid absolutes: No system can scale indefinitely (This is vague on the book though)
- Don't pick a target based on current performance: Ok book, but then pick a target base on... I assume historical performance?

SLOs are a major driver for prioritizing SRE work, they can be large levers: Chose them wisely! Be careful not to get them wrong (overly aggressive requiring heroic efforts, or too lax resulting in a bad product)

### Control measures
So here's how the book recommends approaching SLIs and SLOs to manage your system (in other words, this seems to imply what infrastructure work you should tackle)
- Monitor your SLIs
- Compare SLIs to the SLOs and decide if action is needed
- If action is needed, figure out what you need to do to meet your targets
- Take action!

Ex: If latency on your service is increasing, it might be that your service is CPU bound and you need more CPU power to meet your SLO!

### Expectations
- Keep a safety margin: Aim for a tighter SLO internally
- Don't overachieve: This one might be a bit too far fetched, but if your system is always up (higher than SLO) users will come to rely on it as infallible. Ideally other systems should handle your system being unavaible (as your SLO isn't 100%). Google recommends to deliberately take your system offline if you're overachieving your SLO, though in practice... I'm not sure how this would play out.

### In practice
- You are in a priviledged position to help you legal team craft an SLA, by helping them understand the likelihood and difficulty of achieving your SLOs
- It is wise to be conservative in what you offer to your users. It can be hard to change SLAs.

---
Personal notes: Uuhhh so an SLO is a bunch of SLIs we keep track off, and if they're all within a certain parameter, the SLO is green?
AAAAA THIS CHAPTER IS VAGUE
LET'S WRAP THIS UP AND JUMP INTO THE WORKBOOK!
