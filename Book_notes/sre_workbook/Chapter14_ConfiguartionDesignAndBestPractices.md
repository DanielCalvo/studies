
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