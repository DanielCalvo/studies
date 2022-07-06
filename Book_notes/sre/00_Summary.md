- Create a list of chapters here, a very short summary (a line or two) and create a link to the file that contains your reading notes for that chapter

## Chapter 3: Embracing Risk
- This chapter explains error budgets and why it's important to agree on them.
- I really liked this phrase: "Hope is not a strategy"

## Chapter 4: Service Level Objectives
- You pick a bunch of SLO targets based on your SLIs and bam, you have SLOs.
- I found this chapter very vague. Examples would've been cool

## Chapter 10: Practical Alerting from Time Series Data
- This chapter was not useful. Read "Prometheus up and Running" instead

## Chapter 11: Being on call
- Shares some overall guidelines about being on call. Nice chapter.

## Chapter 13: Emergency Response
- Takeaways: Keep calm. Make sure you're familiar with the incident response process. Keep a history of outages to learn and improve upon based on what broke.
- But other than that the chapter goes over some incident response examples in Google, but more it's more anecdotal than reference material and/or actionable info. Not a particularly useful chapter

## Chapter 14: Managed incidents
- Provides a brief and high-level summary as how to handle an incident. Handy chapter.

## Chapter 15: Postmortem culture: Learning from Failure
- Briefly goes over the idea of postmortems and why they're important. Not a bad chapter, but not very hands-on. The SRE Workbook or Google might give a few concrete examples.

## Chapter 16: Tracking outages
- Google has a tool that tracks alerts and outages for aggregation. It can be handy for spotting patterns, but this seems like a "nice to have" and is nothing groundbreaking. Glancing over a sequence of alarms in slack might yield the same result.

## Chapter 18: Software Engineering in SRE
- Mostly aspirational chapter about some of the process of writing software in Google as an SRE, has a few bits of useful info, but otherwise dwells too much on Google specific details. Hard to find action on this one, unless in very specific cases (ex: building your own infrastructure tools, like Crossplane)
- Don't underestimate driving adoption and customer service. Deliver fast, launch and iterate

## Chapter 19: Load Balancing at the frontend
- Talks about load balancing using DNS and Network Load Balancers.
- Most of this is offered by Cloudflare (DNS+LB) and AWS (Global load balancer, NLBs and Route53). Nothing too ground breaking, though Google does it their own way.


## Chapter 29 - Dealing with Interrupts
- Do one thing well. Project time and interrupts at the same time don't mix
- Respect yourself and your customers. If there are too many onerous or annoying tickets -- it's okay to push some work back to them