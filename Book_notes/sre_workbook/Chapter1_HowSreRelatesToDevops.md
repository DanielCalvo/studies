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
- The SRE and the product team need to select an appropriate target for the service and its use base. and manage according to that SLO
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