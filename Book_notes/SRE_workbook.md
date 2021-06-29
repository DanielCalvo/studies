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

### Share Ownership with Developers
- SRE bread and butter: Expertise around availability, latency, performance, efficiency, change management, monitoring, emergency response and capacity planning of the service(s) we're looking after
- You can get a lot more done if you blur the lines. SRE can instrument javascript, product people can do SRE stuff

### Use the same tooling, regardless of job function or title
- Use the same tooling to manage a service, both for SRE and product people (this one's kinda vague though)

### Narrow, rigid incentives narrow your success
- Careful with incentives that are too narrow (just launch related or reliability related)
- Early SRE engagement is good for services!

### It's better to fix it yourself, don't blame someone else
- Don't pass blame to other groups, that is the core problem with traditional engineering/ops
- Encourage engineers to change code and config required to the product. Allow these teams to be radical within the limits of their mission, eliminating incentives to proceed slowly
- Support blameless postmortems
- Allow support to move away from products that are irredeemably operational difficult. People don't like doing that, will quit 

### When can substitute for whether
- The decision to give or withdraw support to a product can be based on "comparative operational characteristics"
- A strong partnership with product development is critically important

Further reading:
- SRE book
- Effective DevOps
- Phoenix Project
- Practice of cloud and sysadmin v2
- Accelerate

# Part I - Foundations
- Basic foundations of SRE: SLOs, monitoring, alerting, toil reduction and simplicity 

## Chapter 2 - Implementing SLOs

## Chapter 3 - SLO Engineering Case Studies

## Chapter 4 - Monitoring

- Let's go all the way until here today!

## Chapter 5 - Alerting on SLOs

## Chapter 6 - Eliminating Toil

### What is Toil?
- Manual, repetitive, automatable, reactive, lacking enduring value, grows as fast as it's source (ex: disk full alerts grow with server numbers)
- Toil can deflate team morale
- Fix the root problem when possible: You can write a script to delete log files, but if you can stop generating overly verbose useless logs in the first place, that's better

### Measuring toil


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
