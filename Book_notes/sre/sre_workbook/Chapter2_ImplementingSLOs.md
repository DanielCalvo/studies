# Implementing SLOs
Uh-oh, SLOs are key to making data driven decisions about reliability, they're at the core of SRE practices! Let's see if this chapter offers some more practical information.

- According to this chapter, SLOs are a tool to determine what engineering work to prioritize (ex: Automating rollbacks or moving to a replicated data store?)
- There's an interesting bit here as to who should own the SLO. In smaller orgs, maybe the CTO. In bigger orgs, maybe the product owner

## What to measure: Using SLIs
It's generally recommended to have SLIs as the ratio between two numbers, ex:
- Successful HTTP requests / total HTTP requests (success rate)
- Number of gRPC calls completes successfully in < 100Mms / total gRPC calls
- Some of them can be a bit more specific, like search that uses corpus, cache usage, things related to the freshness of data on the cache and more things

## SLI implementation
- For the first SLI, you can chose something that requires minimum engineering work (if you already have logs and setting up probes or instrumenting the code would be too time expensive, use the logs! You might need to configure your webserver to record this information)

Oh, also, to implement your SLI, you can use:
- Application server logs
- Load balancer monitoring
- Black box monitoring
- Client-side instrumentation
- Whatever else you find handy

## Using SLIs to calculate starter SLOs
- Proposed SLOs for the API (given historical data for it queried with Prometheus in the book!)

| SLO Type | Objective |
|----------|:-------------:|
| Availability | 97% |
| Latency | 90% of requests < 450ms |
| Latency | 99% of requests < 900ms |

## Chosing time windows
- A four week rolling window is a good general purpose interval.
- You can also have weekly summaries and quarterly summaries

## Getting stakeholder agreement
- The product managers have to agree that this threshold is good enough for users. Performance below this threshold is unacceptably low and worth spending energy to fix
- The developers need to agree that if the error budget has been exhausted, we'll take some steps to reduce risk until we're back on the budget
- The team responsible for the prod environment who are tasked with defending this SLO agree that it is defensible without Herculean effort, excessive toil and burnout.

Agreeing on all of the above is the hard part! To defend your SLO, you need to set up monitoring and alerting.


## Establishing an error policy budget
You can use your SLO to derive an error budged. To use it, you need a policy outlining what to do when your service runs out of budget . Getting the error budget policy approved by stakeholders is a good test for whether the SLOs are fit for purpose
- If the SREs feel the SLO is not defendible with too much toil, relax the SLO
- Ifht development and product people feel that fixing reliability will cause feature release velocity to drop too much, they can also argue for relaxing objectives
- If the product manager feels this SLO will result in a poor experience for users, this SLO is likely not tight enough

If the three parties do not agree on enforcing the error budget policy, you need to iterate on the SLIs and SLOs until all stakeholders are happy.

Decide how to move forward and on what you need to make a decision: More data, more resources, or a change to the SLI/SLO?

## Error budget enforcement decisions
What do you do when you exhaust your error budget?
- The development team gives top priority to bugs relating to reliability issues. This come with high-level approval to push bacn on external feature requests and mandates
- To reduce the risk of more outages, a production freeze halts certain changes to the system until there's sufficient error budget to resume changes (uh-oh)

Sometimes a service consumes the entirety of its error budget, but not all stakeholders agree that enacting the error budget is appropriate. IF this happens, you need to return to the error budget policy approval stage

## Documenting the SLO and the error budget policy
This should include, for the SLO:
- Autors of the SLO and reviewers
- When it was approved and when it should be reviewed again
- Description of the service
- Details of the SLO: Objectives and SLI implementations
- The details of how the error budget is calculated and consumed
- The rationale from behind the numbers: Whether they were derive from experimental or observational data. Even if SLOs are totally ad-hoc, document this so engineers don't take decisions based on ad-hoc data

How often you review your SLOs depends on the maturity of your SLO culture. When starting out, review them more often (maybe once a month)

And this should include, for the error budget: 
- Authores, reviewers, approvers
- When approved/When to review
- Brief service description
- Actions to be taken in response to budget exhaustion
- An escalation path to follow if there's disagreement on the calculation or whether the agreed upon actions are appropriate in the circumstances

## Dashboards and reports
- Ooohh it can be helpful to have dashboards. This is cool!
- Idea: A dashboard showing overall SLO compliance
- Ooh, you can also have a dashboard showing SLI trends, that seems interesting! (though maybe too fancy to implement)

## Improving quality of your SLOs
- Lots of aspirational advice here. Ultimately, the goal of the SLO is to keep your users happy
- If setting and following an SLO is too difficult for now, consider setting an aspirational SLO.

## Advanced topics
- You could try modeling the entire user journey in a shopping site as an SLI, though that's advanced!

## Grading interaction important
- Certain services are more important than others. Billing? Very important. Notification service for likes on a social media app? Maybe not so much.

## Modeling dependencies
- Be aware that your service might be dependent on other services. You might need to coordinate reliability requirements across the stack

## Conclusion
- SLOs are the tool by which you measure the reliability of your service
- Error budgets are a tool for balancing reliability with other engineering work
- Start using SLOs and error budgets today!