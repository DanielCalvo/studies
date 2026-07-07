## Foreword
- The chapters are not rigorous academic papers, they're personal accounts, written with pride

## Preface
- Read the book and draw useful conclusions about your own environment!

The SRE Way:
- Throughness and dedication
- Belief in the value of preparation and documentation
- An awareness of what could go wrong, and a strong desire to prevent it

The book is a series of essays. Some are good, others not so much.

## Chapter 1: Introduction
- Hope is not a strategy
- SRE is what happens when you ask a software engineer to design an operations team.
    - Dani notes: I see some issues in SRE (ex: over engineered systems) that seem to stem from a software POV (sysadmins usually keep it simple) 
- 50-60% SREs at Google were hired as SWEs, the rest are close to SWEs but in addition know unix and networking
- Common to SREs is the belief in and aptitude to develop software to solve problems
- SREs have to code to manage their services or they'll drown in manual work. Google places a 50% cap on ops work (on call, manual tasks, etc)

### Tenets of SRE
In general, a SRE team is reponsible for, for their services:
- Availability
- Latency
- Performance
- Efficiency
- Chamge management
- Monitoring
- Emergency Response
- Capacity planning

### Other notes
- The conflict between pursuing maximum change velocity without breaking a SLO is solved by using an error budget! More on error budgets in other chapters.
- 100% reliability is the wrong target!

### Monitoring
There are three kinds of valid monitoring output
- Alerts signify a human needs to take action immediately
- Tickets means a human should take action eventually, but not immediately
- Logging means its recorded for investigation later in case you need it

### Emergency response
- Playbooks improve MTTR by about 3x. Playbooks are great if you can have them

### Change Management
70% of outages are due to changes in a live system. Use automation to:
- Implement progressive rollouts
- Quickly and accurately detect problems
- Roll back changes safely when problems arise

Do capacity planning based on the demand you're expecting!

## Provisioning and Efficiency and performance
- Make sure to only provision what you need
- SREs provision to meet a capacity target at a certain response speed, and thus have an interest in the service peformance
- SREs and developers should monitor and modify a service to improve its performance, "adding capacity and improving efficiency"
- I would replace the above by "Making sure enough capacity is allocated to a service so that it remains efficient", as most places care a bit more about money than Google does!