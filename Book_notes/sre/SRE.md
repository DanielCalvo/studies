## Table of Contents
## Foreword
## Preface

## Part I - Introduction
### Chapter 1 - Introduction
- Errors budgets are meant to solve the conflict between developers and ops
- We can then aim to spend our error budget getting maximum feature velocity. An outtage is no longer a bad thing, it's an expected feature of innovation
- Hmm theres some interesting basic insights on alarming here
- playbooks > winging it

### Chapter 2 - The Production Environment at Google, from the Viewpoint of an SRE


## Part II - Principles
### Chapter 3 - Embracing Risk
### Chapter 4 - Service Level Objectives
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