## Table of Contents
## Foreword
## Preface

## Part I - Introduction
## Chapter 1 - Introduction
- Errors budgets are meant to solve the conflict between developers and ops
- We can then aim to spend our error budget getting maximum feature velocity. An outtage is no longer a bad thing, it's an expected feature of innovation
- Hmm theres some interesting basic insights on alarming here
- playbooks > winging it

## Chapter 2 - The Production Environment at Google, from the Viewpoint of an SRE


## Part II - Principles
## Chapter 3 - Embracing Risk

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

### Chapter 5 - Eliminating Toil
### Chapter 6 - Monitoring Distributed Systems
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

### Chapter 7 - The Evolution of Automation at Google
### Chapter 8 - Release Engineering
Ideal release engineering characteristics
- Self service (teams can adopt existing tools/templates)
- High velocity 
- Hermetic builds (ex: The same codebase on the same hash version should build the same thing)

The book describes `rapid` which contains all the common steps you would see in a pipeline nowadays (building, branching, testing, packaging, deployment)
Nothing too groundbreaking in this chapter.

### Chapter 9 - Simplicity
- The book stresses: Software simplicity is a prerequisite for reliability
- Keep your releases small and simple
- Keep your APIs simple
- Boring is a positive attribute when it comes to software.
- Be careful of bloat. Less can be more.

## Part III - Practices
There is a site reliability hierarchy! From the base to the top:
1. Monitoring
2. Incident response
3. Post mortem / root cause analysis
4. Testing + release procedures
5. Capacity planning
6. Development
7. Product
---
1. Monitoring: Without monitoring, you're flying blind. You need to tell if the service is working
2. Incident response: Once a service breaks, you need to know how to respond to it
3. Post mortem: Build a culture of blameless post mortems to learn what went wrong and how to fix it
4. Testing: Once we know what went wrong, we should be able to test it to keep it from happening again
5. Capacity planning: ??? No further info is given here
6. Development ??? - Also vague. Book seems to hint at doing large scale software design as part of the dev process
7. Product: No further info

### Chapter 10 - Practical Alerting
### Chapter 11 - Being On-Call
### Chapter 12 - Effective Troubleshooting
- There's no substitute for knowledge of the system you're trying to troubleshoot!
- Theory: Look at the logs
- AAAAA this chapter is boring I can't deal with it

### Chapter 13 - Emergency Response

### Chapter 14 - Managing Incidents
- This chapter is a bit aspirational. It may not be possible to ask yourself all of the big, improbable questions people at google ask themselves here.

I feel the key takeaways are, when an incident happens:
- Focus on what went well
- And focus on what you learned

Google encourages proactive testing of your systems.

### Chapter 15 - Postmortem Culture: Learning from Failure
### Chapter 16 - Tracking Outages 
### Chapter 17 - Testing for Reliability
### Chapter 18 - Software Engineering in SRE
### Chapter 19 - Load Balancing at the Frontend
### Chapter 20 - Load Balancing in the Datacenter
### Chapter 21 - Handling Overload
### Chapter 22 - Addressing Cascading Failures
### Chapter 23 - Managing Critical State: Distributed Consensus for Reliability
### Chapter 24 - Distributed Periodic Scheduling with Cron

### Chapter 25 - Data Processing Pipelines


### Chapter 26 - Data Integrity: What You Read Is What You Wrote
### Chapter 27 - Reliable Product Launches at Scale

## Part IV - Management
### Chapter 28 - Accelerating SREs to On-Call and Beyond
### Chapter 29 - Dealing with Interrupts
- Google recommends that if you're going to be handling interrupts, do that only (pages, tickets, slack, etc) 
- You can't be in the zone (aka flow) if you're being interrupted all the time. However, if your job is 100% interrups, you can get in the zone too!
- Book says: Do one thing well. Project time and interrupts at the same time don't mix

- Limit context switches. Do just project work or just interrupts for a week, a day or half a day

- The book encourages you to attempt to find a root cause to some of your tickets to see if you can fix the problem, instead of doing tickets dealing with the symptoms.

Respect yourself and your customers
- If tickets are onerous or annoying, it's okay to push back some of the effort onto your customers. Strike a balance between respecting your customer and yourself.
- Maybe you need to support a flaky tool that's annoying. If you don't get a lot of help, maybe the tool isn't important.
- If dealing with an interrupt doesn't require privileges, consider using policy to push the request back to the requestor. Instruct the customer to perform the step and send you back for review. If a customer wants a certain task accomplished, it's okay for them to spend some effort getting what they want.
- I like this part: "Your guiding principle in constructing a strategy for dealing with customer requests is that the request should be meaningful, be rational, and provide all the information and legwork you need in order to fulfill the request. In return, your response should be helpful and timely"

### Chapter 30 - Embedding an SRE to Recover from Operational Overload


### Chapter 31 - Communication and Collaboration in SRE
### Chapter 32 - The Evolving SRE Engagement Model

## Part V - Conclusions
### Chapter 33 - Lessons Learned from Other Industries
### Chapter 34 - Conclusion
### Appendix A - Availability Table
### Appendix B - A Collection of Best Practices for Production Services
### Appendix C - Example Incident State Document
### Appendix D - Example Postmortem
### Appendix E - Launch Coordination Checklist
### Appendix F - Example Production Meeting Minutes
### Bibliography