## Embracing risk

- Google says: Make a service reliable, but no more reliable than it has to be.
- Maximizing reliability comes at the cost of limiting feature delivery speed and hiw quickly products can be delivered to customers
SRE seeks to balance the risk of unavailability with the goals of rapid innovation and efficient service operations, so that user overall happiness with features, service and performance is achieved.

## Managing risk
Cost does not increase linearly with reliability. After a certain point more reliability becomes really expensive
- This means the cost of machine/compute resources
- But also the opportunity cost (ex: not shipping features frequently enough)

## Measuring service risk
The most straightforward way of representing risk is in the way of unplanned downtime
- Usually expressed in nines: 99,9%, 99,99% and so on
- Google defines availability in terms of _request success rate_
- For web services this can be HTPP requests served, but for a data processing pipeline it can be a percentage of items processed

## Identifying the risk tolerance of consumer services
Different services have different availability targets depending on several factors, such as
- User expectations
- If service is tied to revenue
- If it is a paid or free service
- If there are competitors
- If the service is free or enterprise

As an example, google apps for work had a much higher availability requirement than youtube when it was adcquired

## Risk tolerance for infra services
A ton of fluff in this explanation, couldn't get to the core of it.
- Make sure to deliver services with explicitly delineated levels of service

## Motivation for error budgets
- Product dev is largely evaluated on product velocity
- SRE is largely evaluated on upon the reliability of a service
- Uh-oh, tension! The error budget settles the score and gives us something we can agree on

Hope is not a strategy!

An error budget grants a clear, objective metric that determines how reliable a service is allowed to be in a single quarter.

- Product manager sets SLO, monitoring system measures uptime.
- If you run out of error budget, you stop releases. As long as there is an error budget, new releases can be pushed

## Key insights
- Managing reliability is all about managing risk, and managing risk can be costly
- 100% is never the right availability target. It's more than what users want or can notice
- An error budget alligns incentives and emphasizes work between SRE and product dev. They make it easier to decide the rate of releases and defuse discussions. They also allow multiple teams to reach the same conclusion about risks without rancor. Nice!