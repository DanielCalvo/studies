How would you set up incident response?

(a rough sketch follows)

## Roles: incident commander and bodyguards
- Okay, so I would begin with having a person who is the incident commander. This is the person who is usually on-call for a period, usually a 24-hour period, and any engineer on the team can be an incident commander. You don't need special technical know-how
- You then have the bodyguards for the teams. The bodyguard is the person on a team who has technical knowledge in a given domain (e.g., infra or the checkout part of an e-commerce platform) that issues can be escalated to. You have one bodyguard per team, and bodyguard duty is equally rotated among the team members

How you want to do on-call for the bodyguards depends. You can always have an on-call person but then also do follow-the-sun, in which case this on-call person will not be paged very frequently because there's almost always someone online. Or you can actually have someone on-call with a document that says you have to be 15 minutes away from a computer at all times during your on-call. Like, it depends

## Duties of the incident commander
- Escalate the issue to appropriate parties
- Communicate and provide updates (e.g., declare an incident, write on channels every X minutes with updates, handle any questions anyone might have)
- Get the postmortem started with an incident response timeline (if possible during the incident)

# You get paged runbook
To start off, if you get paged and you can't respond properly, like you're at the supermarket or something, let the page go through and let the second-in-line person get it

## Part 1: Triaging & declaring an incident
- Hop into the incident response room (like a Google Meet call)
- Triage the alert to confirm it's real (maybe it's a false positive, maybe it's 1 customer having local issues)
- If it is real, see the product areas and their priority for issues
    - E.g., if the recommendation engine is not working, it can wait until tomorrow (it's a P2). If someone can't add things to the shopping cart, we must act immediately
- Escalate the issue to the corresponding team (it's a P0)
    - If you are not sure if the issue is real or you don't have enough know-how to determine it, escalate it anyway. If in doubt, escalate and get more people on board
- Formally declare an incident on whatever incident management tool you might have (PagerDuty, FireHydrant, and so on)

## Part 2: During the incident
- Keep all parties informed of progress by providing periodic updates
- Escalate and bring more people in as necessary
- Sketch out any useful information, context, or timeline in the sketch of a postmortem just to keep the information somewhere so you don't lose it (as you're able to)

## Part 3: Post-incident: contribute to the postmortem
- Contribute to the postmortem with any information that you had at the time as incident commander, with things like:
    - Timelines
    - Dashboards used to troubleshoot
    - Context discussed during the incident response call
    - Any other information that might be relevant

## Part 4: Postmortem call
- After the postmortem is complete, schedule a call with the rest of the engineering team so that the postmortem can be reviewed and discussed together
- It can be difficult to get a postmortem perfect the first try
    - Reviewing it with the team always brings up new ideas and things that were initially missed

## Part 5: Follow-ups
- After the postmortem review call, make sure that any feedback given on the call is addressed and documented
- Also make sure that all the action items are on the radar of their respective teams to be followed up with the appropriate priority
