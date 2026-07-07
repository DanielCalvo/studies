- You can onboard an existing service into SRE support 
- SRE can be involved in the design of a service
- Or SRE can provide product development of... SRE validated infrastructure? More info is required

## The PRR Model
- The most typical initial step of SRE engagement is the Production Readiness Review!
- This is to identify the reliability needs of a service based on its specific details

## The SRE Engagement Model
SRE seeks production responsability for services for which it can make contributions to reliability. Aspects of a service SRE is concerned with are:
- System architecture and interservice dependencies
- Insturmentation, metrics and monitoring
- Emergency response
- Capacity planning
- Change management
- Peformance: Availability, latency and efficiency

When SRE enganges with a service, we aim to improve it along all of these axes.

## Alternative support
Not all Google services receive SRE engagement though. Some don't need high reliability and availability, and sometimes the number of dev teams might exceed the bandwidth the SRE team has

### Documentation
- Google's production guide documents production best practices for services as determined by experiences of SRE and development teams alike
- Developers can implement the solutions and recommendations in such documentation to improve their services

### Consultation
- Developers may also seek SRE consulting to discuss specific services or problem areas to improve the service in production
- Services can grow too much, or other services can be become dependent on this service (?)

## Simple PRR
The objective of the PRR are as follows:
- Verify that the service meets the accepted standards of the productions set up and that the owners are ready to work with the SRE team
- Improve the reliability of the service in production and minimize the severity of incidents that might be expected (by taking action on certain things... that are not listed here?)


There are three different engagement models
- Simple PRR Model
- Early Engagement Model
- Frameworks and SRE platform

## Simple PRR

### Engagement
The discussion with the development team goes over
- Establishing SLO/SLAs 
- Planning for potentially disruptive design changes to improve reliability (this one seems very optimistic from google)
- Planning and training schedules (training for what?)

### Analysis
Actual work! Analyze the service, gauge it's maturity for the things SREs are interested on (see above for a list). Here's a few things to go over:
Oh man this list is too google
- A web service should not be dependable on a service that does batch processing
- Does the service sends its logs to a loggin service?
- Is the service instrumented and monitored?

### Improvements and Refactoring
After the analysis phase, you can improve things!
1. Improvements are prioritized based upon importance for service reliability
2. The priorities are discussed and negotiated with the development team
3. Both the SRE and dev teams assist each other

### Training
- Ah, so the SREs that were involved in the PRR above train the rest of the SRE team to be able to manage the service in production

### Onboarding
- Then the responsability of this service in production is gradually transfered over to the SRE team

### Continuous Improvement
- Learn how the service behaves in prod and share that with the dev team


## Evolving the PRR Model: Early Engagement
TLDR: It is positive if SRE is involved in the initial development process so they can help a given service be reliable 
- There's some description in which services qualify for early engagement but to be honest I think almost every service could benefit from some SRE review

Alrite let's go over the benefits!

### Design phase
- SREs can help developers make trade off during the design phase, and understand their decisions!

### Build and implementation
- SRE can help with instrumentation, metrics, emergency controls, resource usage and efficiency

### Launch
- Oh you can do dark launches to gain insight on possible errors. 

There are also sections for "Post Launch" and "Disengaging from a service" here but they don't contain anything particularly meaningful

## Evolving Services Development: Frameworks and SRE Platform

### Lessons learned from Early Engagement and Simple PRR
- Onboarding required 2-3 SREs for 2-3 quarters. That's a lot of manpower.
- Software was done differently, certain production features were implemented differently
- Certain patterns were also solved differently
- SRE software contributions were often local to the service, building a generic solution to be reused was difficult. There was no way to implement learned lessons across various services

## Towards a structural solution: Frameworks!
To solve the above, a model was developed that:
- Codified best practices of what works well in production in code, so services can use this code and become production ready!
- Reusable solutions: Commond and shareable implementations of techniques used to mitigate scalability and reliability issues
- A common production platform with a common control surface: Uniform: Monitoring, logging and configuration for all services. Also... uniform operational controls?
- Easier automation and smarter systems: A common control surface makes it easy to automate things.

Framweork modules developed by the SRE team typically address:
- Instrumentation and metrics (Standard dimensions for monitoring instrumentation)
- Request Logging (Standard format for request debugging logs)
- Control systems involving traffic and load management
Plus some other details about load shedding, determination of "overload", and business logic that I couldn't fully grasp at this time.

Frameworks drive a single re-usable solution for production concerns. Devs don't have to reinvent all the things in slightly incompatible ways (ex: logging systems)

## New Service and Management Benefits
Using frameworks provided a ton of benefits!

### Significantly lower operational overhead
- It supports strong conformance tests for coding structure, dependencies, tests, coding style guides and so on
- It features built in service deployment, monitoring and automation
- It facilitates management of a large number of services, especially microservices
- It enables much faster deployment! You can go from idea to full SRE-level production service in just a few days!

### Universal support by design
Not al services warrant SRE support or can maintained by SREs. But they can be build using production ready features that are maintained and developed by SREs

### Faster, lower overhead engagements
The framework approach results in faster PRR as we can rely upon
- Built in service features as part of the framework
- Faster service onboarding
- Less cognitive burden for SRE teams! (I like this one)

### A new engagement model based on share responsibility
- The original SRE model only presented the options of either full SRE support or approximately no SRE engagement
- A production platform with SRE teams providing support to the platform infra, while the developers handled on call support
- Under this structure, SREs assume responsibility for areas like infrastructure, such as load shedding, overload, autoamtion, traffic management, logging and monitoring



---
Notes from self:
- If you want to have a PRR, you have to develop your own checklist. Google's one is too specific for Google.
- This chapter seems to imply SREs need to have application architecture knowledge
- You'll also have to figure out what your "Production Ready Framework" looks like. Google's examples are very Google centric!