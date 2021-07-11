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

## Chapter 5 - Alerting on SLOs

## Chapter 6 - Eliminating Toil

### What is Toil?
- Manual, repetitive, automatable, reactive, lacking enduring value, grows as fast as it's source (ex: disk full alerts grow with server numbers)
- Toil can deflate team morale
- Fix the root problem when possible: You can write a script to delete log files, but if you can stop generating overly verbose useless logs in the first place, that's better

### Measuring toil
- Stopped on page 94
- This chapter is really cool though, I should continue until they go with those examples at google, the principles until then are awesome!

## Chapter 7 - Simplicity

# Part II - Practices

## Chapter 8 - On-Call

## Chapter 9 - Incident Response

## Chapter 10 - Postmortem Culture: Learning from Failure

## Chapter 11 - Managing Load

## Chapter 12 - Introducing Non-Abstract Large System Design
- Google: "We consider reliability to be the most critical feature of any production system"
- The principles at the beginning of the chapter are cool
    - You should write about those!
- But then the chapter goes on to explain complex application/system design that has like 100TB of data processed by day. Not sure if it's a weird flex from the authors, but kinda hard to take action points from the example provided
    - Maybe if you really dig into it you can extract more info though


## Chapter 13 - Data Processing Pipelines

## Chapter 14 - Configuration Design and Best Practices
- Designing configuration with clarity and usability in mind is a good idea
- A good configuration change interface allows for quick, confident and testable configuration changes

### Configuration and reliability
- The quality of a human-computer interface of a system's configuration impacts the ability to run that system in production
- The more complicated this interface is, the harder it is to maintain this config
- Configs that are harder to maintain and change are less reliable
    - Early airplanes had confusing controls and this led to accidents

### Configuration Philosophy
- The ideal configuration is no configuration at all!
    - The ideal configuration can be recognized from deployment, workload, or existing pieces of configuration, or defaults can be assumed <- Careful how you interpret this
- System with a large amount of controls require a large amount of human operator training
    - Such training is no longer feasible in the majority of the IT industry
- While this reduces the amount of control you can exercise over a system, it decreases the surface area for error and cognitive load on the operator
- As the system becomes increasingly complex, this is important
- When these principles were applied at google, they resulted in easy, broad adoption and low cost for internal user support

### Configuration Asks Users Questions
- It all boils down to an interface that asks user questions, regardless if you're editing XML or using a GUI
- There are two perspectives here
    - Infrastructure centric view: Offers as many configuration knobs as possible. The more knobs the better, as the system can be tuned to perfection
    - User-centric view: Asks questions the user must answer before they can get back to working on their business goals. The fewer knobs the better, answering config questions is a bore 

- **Driven by our initial philosophy of minimizing user inputs, we favor the user-centric view**
- Focusing configuration on the user means that your software needs to be designed with a particular set of use cases for your audience. This requires user research
- Infra-centric systens requires considerable configuration from the user
- Limited configuration options can lead to better adoption ("it works out of the box")
- Some systems can begin infrastructure centric and move toward a user centric focus

### Questions should be close to user goals
- Make sure users can easily relate to the questions you ask
- The tea metaphor is here but I can't quite summarize it well... snap! - *RECHECK**

### Mandatory and Optional Questions
- A given config set up might contain mandatory and optional questions
- To remain user centric and easy to adopt, minimize mandatory questions
- The easiest path to reduce mandatory questions is to make them optional. Provide sane defaults for most use cases
- Defaults can be dynamic (ex: number of cpu cores configures == number of cpu cores on the system)
- Think carefully about your defaults: Most users will use defaults. Wrong defaults can be harmful
- Some optional questions don't have a clear use case. You might want to remove those altogether. A large number of optional params can confiuse a user. Add optional configuration input only when motivated by a real need

### Escaping Simplicity
- Configuration may need to account for power users
- Maybe have a way to support additional config, ex: Have "advanced parameters" or "advanced settings" as an optional thing that can be accessed somehow, but is not mandatory as to not induce decision paralysis, slower rate of change to lower confidence and chance of mistakes

### Mechanics of configuration
- You stopped here at end of page 308

## Chapter 15 - Configuration Specifics

## Chapter 16 - Canarying Releases

# Part III - Processes

## Chapter 17 - Identifying and Recovering from Overload

## Chapter 18 - SRE Engagement Model
- This seemed interesting with useful information on how to engage on project development and talking to other teams

## Chapter 19 - SRE: Reaching Beyond Your Walls
- Hmm, this is probably interesting too!

## Chapter 20 - SRE Team Lifecycles

## Chapter 21 - Organizational Change Management in SRE
