### Introduction
- Page 32 has an ideal release process (seems to be of the continuous delivery persuasion)
- This entire section is excellent! Many best practices are summarized briefly

### Chapter 1: Designing in a Distributed World

### Chapter 4: Application Architectures
- Single machine webserver: Just a single webserver, serving things. Appropriate for small web sites.
- Three tier web service: Load Balancer > Web Server > Data server (such as a DB)
- Four tier web service: Load Balancer > Frontend web > App server > Data server (such as a DB)