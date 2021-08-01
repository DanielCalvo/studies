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
- Let's read this next! 
- Your monitoring system should address two questions: What's broken, and why?
- Whitebox monitoring is essential for debugging

- Four golden signals:
    - Latency: How long it takes to serve a request
    - Traffic: How much demand is being placed on your system
    - Errors: The rate of request that fails
    - Saturation: How much of your system resources are being used
- Measuring all 4 signals and alerting a human when one of them is problematic is a good start
- Worrying about your tail: Careful with averages. Some of your requests might take very long, but not show on an average. Historigrams per exponential boundaries are useful to track these

### Chapter 7 - The Evolution of Automation at Google
### Chapter 8 - Release Engineering
### Chapter 9 - Simplicity

## Part III - Practices
### Chapter 10 - Practical Alerting
### Chapter 11 - Being On-Call
### Chapter 12 - Effective Troubleshooting
### Chapter 13 - Emergency Response
### Chapter 14 - Managing Incidents
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