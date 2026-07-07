This chapter was fun and entertaining to read. It's well written, though some of it may not relate directly to SRE work. Cool things talked about in this chapter:

Hope is not a strategy.
- I really like this saying

## Preparedness and Disaster Testing
- Attention to detail: A small mistake can have big effects. Pay attention!
- Swing capacity: Have means to scale up when unforeseen load shows up
- Training and certification: Training for what you need to do is important
- Focus on detailed requirements gathering and design: Other industries have a much lower apetite for risk and spend longer designing their systems. These are different from the "launch and iterate" culture present in IT
- Defense in depth: Certain industries (nuclear) have several layers of protection for safety. I wonder how this could apply to IT... Disaster recovery? High availability architecture? High availability humans? All of the above?

## Post mortem culture!
- Other industries use this as well, but with other names. When something goes wrong, it's important to evaluate:
- What happened
- The effectiveness of the response
- What we would do differently next time
- What actions will be taken to make sure a particular incident won't happen again

Other industries have other motivations for post mortems, mostly safety related.

## Automating away repetitive work and operational overhead
- SRE as a discipline is really keen on automation
- Other areas of work, like nuclear reactor maintenance, are not so keen on automating things.
    - A slow and steady, methodical approach might be better than trying to finish something quickly
- However, when a response to a given incident must be quick, automation can be your best bet.
- Automation to test input (test that your changes will work) can eliminate errors. Book references LASIK surgery here

## Structured and rational decision making
In SRE at Google, the teams strive to take decisions ensuring that
- The basis for the decision is agreed upon in advance
- The inputs to the decisions are clear (Hmm, does this mean the "why"?)
- Any assumptions are explicitly stated
- Data driven decisions win over decisions based on feelings, hunches, or the opinion of the most senior employee in the room

Google SRE also operates under the assumption that everyone on the team:
- Has the best interests of a service's user at heart
- Can figure out how to proceed with the data available

### Many industries heavily focus on playbooks and procedures rather than open ended problem solving
This is common in
- Industries that evolve and develop relatively slowly (in other words, definitely not startups)
- Also in industries in which the skill level of workers may be limited

Interesting take here. The slower your service changes, the more runbooks make sense. If you're iterating quickly, runbooks are not as useful and they'll get outdated very quickly.

## Conclusion
- In software/web in general, you can move a lot faster than in other industries as errors don't cause people to die or millions to be lost.
- You can use tools like error budgets to fund your culture of innovation and calculated risk taking