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

### Chapter 5 - Eliminating Toil

### Chapter 7 - The Evolution of Automation at Google
### Chapter 8 - Release Engineering
Ideal release engineering characteristics
- Self service (teams can adopt existing tools/templates)
- High velocity 
- Hermetic builds (ex: The same codebase on the same hash version should build the same thing)

The book describes `rapid` which contains all the common steps you would see in a pipeline nowadays (building, branching, testing, packaging, deployment)
Nothing too groundbreaking in this chapter.


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