# Chapter 11: Being on call

## Life of an on call engineer
- Kinda obvious now that we think about it, but the quickness of the on call response is directly related to how available a service must be
- If it's highly available, we have to respond quickly. If not, then not so much
- Nonpaging production alerts can be handled during business hours

## Balanced on call
- Google insists on having half ot the time being engineering time
- Up to 25% of the time can be on call time
- Therefore, the minimum number of engineers to run an on-call schedule is 8, with a primary and a secondary
    - If you only have a primary, then 4 people
- Multi-side teams (geographically distributed) are advantageous! Follow the sun is good, night shifts are stresfull

## Feeling safe
The most important on-call resources are:
- Clear escalation paths
- Well-defined incident-management procedures
- A blameless post mortem culture

## Avoiding operational overload
- All paging alerts should be actionable (yeah!)
- If a given software is too unreliable, re-negotiate on call (ex: you can give back the pager). Have a conversation with them!