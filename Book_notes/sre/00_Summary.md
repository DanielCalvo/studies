Read the book intro!

## Chapter 1: Introduction
- PENDING

## Chapter 2: The Production Environment at Google
- Google used k8s, prometheus and Global load balancers before it was cool
- Most of the relevant technologies described in this chapter are available by public cloud providers

## Chapter 3: Embracing Risk
- This chapter explains error budgets and why it's important to agree on them
- I really liked this phrase: "Hope is not a strategy"

## Chapter 4: Service Level Objectives
- You pick a bunch of SLO targets based on your SLIs and bam, you have SLOs
- I found this chapter very vague. Examples would've been cool

## Chapter 5: Eliminating Toil
- Beware of toil: Manual, repetitive work that is of little value. Be careful with the amount of toil that you take, and work to automate it away 

## Chapter 6 - Monitoring Distributed Systems
- What's broken, and why? The 4 golden signals are a good place to start. Histograms are better than averages. Keep it simple.

## Chapter 7 - The Evolution of Automation at Google
- PENDING summary! Though you read this.
- THIS CHAPTER IS SO BORING AAAAAAAAA

## Chapter 8 - Release Engineering
- PENDING

## Chapter 9 - Simplicity
- Keep it simple whenever possible. Boring can be a positive attribute

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

## Chapter 24 - Distributed Periodic Scheduling with Cron
- Talks about some characteristics of a distributed Cron system.
- This one is a bit too specific to Google scale. Us mere mortals can use k8s, nomad, mesos or whatever other schedulers that solve this on a more common scale

## Chapter 25 - Data Processing Pipelines
- SKIPPED FOR NOW: Seems to case specific for data processing, will skip for now

## Chapter 29 - Dealing with Interrupts
- Do one thing well. Project time and interrupts at the same time don't mix
- Respect yourself and your customers. If there are too many onerous or annoying tickets -- it's okay to push some work back to them

## Chapter 31 - Communication and Collaboration in SRE
- Strange chapter. Goes over weekly "production update" meetings (going over the state of currently managed services. Could work outside google, but could also be wasteful and have low attendance)
- Then goes over a tool they developed which... doesn't seem to relate much to the topic at hand. Not a very good chapter

## CHAPTER 32 - Evolving The SRE Engagement Model MODELLLELELEL
- Do a proper summary!
- Great chapter with some great ideas. 

## Chapter 33 - Lessons Learned from Other Industries
- Do a proper summary!
- Interesting, varied chapter. Goes over decision making processes, postmortems and when runbooks make sense (only if your systems don't change too fast)