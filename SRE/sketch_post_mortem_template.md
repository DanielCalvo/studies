What makes a post mortem good? What must it contain? Walk me through it.

(a rough sketch follows)

## Title section
Incident name: Customers were unable to add items to the shopping cart
Duration: 19:00 to 20:00 GMT, 01/Apr/2026
Severity: P0 (not being able to add items to the shopping cart is considered a P0, as it is a critical user journey affecting business revenue)
Impact: Users were not able to add items to their shopping cart, which meant users could not successfully complete their purchases
Authors: Joe & Bob
Brief explanation: After the release of `v123`, users were shown an error message when trying to add items to their shopping cart. This impacted 000 users and caused an estimated 000 of lost revenue during the outage.

## Description

### Affected users
- Number of affected users, usually estimated by traffic at the moment of impact

### Incident detection time
- How long did it take for this incident to be detected? How could the detection time be reduced?

### Incident resolution timeline
A table with which action was taken at a high level and at what time it happened, e.g.:

- 19:00: first responder gets paged by an automated alert for the "add something to shopping cart"
- 19:03: first responder joins incident room, starts triaging incident
- 19:05: the incident is confirmed and is escalated to the shopping-experience team (their on-call person gets paged)
- 19:07: shopping-experience on-call person joins the incident call and starts troubleshooting
- 19:10: The issue is pinpointed to a change in the shopping cart behaviour in a release that went out a few minutes ago
- 19:15: The release is rolled back and the incident is mitigated, a release freeze is put in place until the next working day so the issue can be fixed during work hours

### Why?
Ask 5 whys to try to figure out why the incident happened (https://expertprogrammanagement.com/2019/05/the-5-whys/)

## Root cause

### What went well
- Detection and escalation were fast, everyone got together on the call very quickly
- Automated alert alerted us immediately
- Faulty release was identified and rolled back quickly

### What did not go well
- The staging e2e tests did not catch this issue due to...

## Conclusion
List of Jira tickets to follow up on, e.g.:

- Have e2e tests cover adding something to the cart more explicitly
- Update `cartconfig` on staging to more closely match prod

## Observations
- Who writes the postmortem: Usually the incident commander fills it out with someone from the team that had the affected service. The commander can fill in the details for the incident response, while more in-depth information about what caused the outage is the responsibility of the team that owns the service
- The post mortem is the maximum priority task for whoever needs to fill it: Drop everything else you're doing for now
- Have a call open to the entire department / engineering team (depends how big your company/eng team is) to go through the postmortem together, review it, see how it can be improved, and have a conversation with the team about what happened (these are oftentimes productive/interesting/insightful)
    - Bonus points if you can have the postmortem ready a few hours before the call and share it on an engineering channel for early feedback
- This is beyond the scope of a postmortem, but you should have a document identifying areas of your software, and which priority level failures in these areas have
    - E.g.: the shopping cart not working is a P0, but the recommendation system not updating might be a P1-P2
- If there is time during the incident, it is usually a good idea for the incident commander to kickstart the postmortem while the owning team of the affected service gets to work
