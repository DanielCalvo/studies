## Chapter 1 - How SRE Relates to DevOps

### Background on DevOps
- Operations as a discipline is hard!
- CALMS - Culture, Automation, Lean, Measuring, Sharing
    - Improve something (often by automating it), measure the results, and share them with your colleagues
- No more silos

### Accidents are normal
- There should be safeguards to prevent things from going wrong

### Changes should be gradual
- Change is risky, split it into smaller changes. This coupled with automation leads to CI/CD

### Tooling and culture are interrelated
- A good culture can work around broken tooling, but the opposite rarely holds true

### Measurement is critical
- You need to measure what's happening so you can be sure that you're changing the reality as you expect (oof, this can be difficult)

### Background on SRE
- SRE is a implementation of DevOps philosophies

--- SRE IS DEFINED BY THE FOLLOWING PRINCIPLES ---

### Operations is a software problem
- SRE should use engineering (SWE?) approaches to solve that problem

### Manage by Service Level Objectives (SLO)
- The SRE and the product team need to select an approprivate target for the service and its use base. and manage according to that SLO
- Deciding on that target requires strong collaboration

### Work to minimize Toil
- Toil is abhorrent. If a machine can perform an operation, a machine should!
- Time spent on toil is time not spent on project work. Project work is how we make our projects more reliable and scalable!

### Automate this year's job away
- The real work here is in determining what to automate, under what conditions, and how
- SRE @ Google has a cap on how much a team member can spend on toil: 50%

# Part I - Foundations

## Chapter 2 - Implementing SLOs

## Chapter 3 - SLO Engineering Case Studies

## Chapter 4 - Monitoring

- Let's go all the way until here today!

## Chapter 5 - Alerting on SLOs

## Chapter 6 - Eliminating Toil

## Chapter 7 - Simplicity

# Part II - Practices

## Chapter 8 - On-Call

## Chapter 9 - Incident Response

## Chapter 10 - Postmortem Culture: Learning from Failure

## Chapter 11 - Managing Load

## Chapter 12 - Introducing Non-Abstract Large System Design

## Chapter 13 - Data Processing Pipelines

## Chapter 14 - Configuration Design and Best Practices

## Chapter 15 - Configuration Specifics

## Chapter 16 - Canarying Releases

# Part III - Processes

## Chapter 17 - Identifying and Recovering from Overload

## Chapter 18 - SRE Engagement Model

## Chapter 19 - SRE: Reaching Beyond Your Walls

## Chapter 20 - SRE Team Lifecycles

## Chapter 21 - Organizational Change Management in SRE
