### Introduction
- Page 32 has an ideal release process (seems to be of the continuous delivery persuasion)
- This entire section is excellent! Many best practices are summarized briefly

#### Business Objectives
- The end result of our ideal environment is that business objectives are met
- To understand that, understand the business objectives and work backwards to arrive at the system you should build
- Well defined business objectives are measurable. Here are some example ones:
    - Sell products via a website
    - Provide service 99.99% of the time
    - Process x purchases a month, growing 10 percent monthly
    - Introduce major features twice a week
    - Fix major bugs within 24h

#### Ideal System Architecture
- Meets the requirements today and provides a path for growth
- Resilient to failure. The architecture includes redundancy and resiliency features that work around failures. (failure is expected/taken into account/planned around) 
- Each subsystem take makes up our service is a service. All subsystems are programmable through an API. The entire system is an ecosystem of interconnected services
- Each subservice is loosely coupled to others, and they can be independently scaled, upgraded or replaced
- The infrastructure is described as code
    - The production environment can be build without human intervention
- Software engineers use this to create micro versions of the envirtonment for personal use
- QA & test engineers use this automation to create environments for system tests

#### Ideal Release Process
- Strive to eliminate all manual steps in the software build, test, release and deployment processes
    - The automation performs tests that prevent defects from being passed to the next step
- Rather than large releases, do small releases. This process could look like:
    1. When you push code, tests are ran that verify basic test functionality
    2. If these pass, an image is built
    3. A test environment is created
    4. More tests are ran against this test environment
    5. Once this is complete, these packages and be pushed to production automatically, with order and caution
    6. The system watches for failures
    7. If there are no failures, the new packages are rolled out to more and more systems until the entire environment is upgraded
- Ideally all problems are caught before they reach production
- If an issue does happen, it is considered serious and new releases are stopped until root cause analysis is complete. Tests are added to prevent future occurrences of this failure
- Due to automation, the roles of release engineering and QA blur together

#### Ideal Operations
- Software is instrumented so it can be monitored
- Data is collected on how long it takes to process transactions, both for internal and external users
- This data is important so that operational decisions can be taken based on data, not guesses, luck or hope
- Measurements are used to detect problems while they are small
- Automated systems detect problems and alert whoever is on call
- There is a playbook of instructions on how to handle every alert that can be generated. The playbook is continuosly improved
- Alerts are documented with a technical description, business impact, and how to fix the issue
- All failures have an active countermeasure. Countermeasures that are activated frequently are automated
- Infrequently activated countermeasures are periodically exercised by causing failures. Like fire drills in school. It's better to learn about a not working db failover on a Monday than it is on a Sunday. Practice makes perfect
- The ideal environment scales up and down automatically
- Dashboards can help understand when it is better to come up with a new architecture other than just add cpu and ram
- When the system is overloaded or degraded, low bandwidth or read only versions of the website can be displayed
- Features of a service can be enabled or disabled without doing an extra deployment
- We strive for excellence: Maintain clear and updated documentation to handle every countermeasure, process and alert. Overactive alerts are tuned and not ignored. Open bug counts are kept to a minimum. Outages are followed by a postmortem with recommendations on how to improve the system
- Developers and operations do not think of themselves as two different teams, they are simply specializations within one large team. All share responsability for maintaining high uptime. All members participate in the on call rotation. Developers are motivated to improve the code when they feel the pain of operations too. Operations must understand the development process if they are going to be able to constructively collaborate

### Chapter 1: Designing in a Distributed World
#### 1.1 Visibility at scale
#### 1.2 The Importance of Simplicity
#### 1.3 Composition
#### 1.4 Distributed State
#### 1.5 The CAP Principle
#### 1.6 Loosely Coupled Systems
#### 1.7 Speed
#### 1.8 Summary

### Chapter 4: Application Architectures

#### 4.1 Single machine webserver
- Just a single webserver, serving things. Appropriate for small web sites.

#### 4.2 Three tier web service
- Load Balancer > Web Server > Data server (such as a DB)

#### 4.3 Four tier web service
- Load Balancer > Frontend web > App server > Data server (such as a DB)


### Chapter 11: Upgrading Live Services

#### 11.1 Taking the Service Down for Upgrading
- Take it down, push new code, bring it back up
- This requires downtime, which is very often not an option
- It can work with replicated services though (like VMs as work replicas behind a load balancer)

#### 11.2 Rolling Upgrades
- Individual machines are removed from service, upgraded, and put back into service

#### 11.3 Canary
- You upgrade a small number of replicas, and wait to see if problems develop
- If no problems are found, you can gradually upgrade more and more machines
- If problems are found, you can roll back
- Canarying is not meant to be a testing process

#### 11.4 Phased Roll-outs
- Certain groups of users receive the upgrade first
- Ex: Imagine you're upgrading reddit, certain subreddits can receive the upgrade first

#### 11.5 Proportional Shedding
- The new service is build on new machines in parallel to the old service
- The load balancer sends a percentage of traffic to the new service
- If all is fine, a larger percentage is sent
- Requires as twice as much capacity though

#### 11.6 Blue-Green Deployment
- There are two environments on the same machine (blue and green)
- Green is live, blue is dormant
- When it's time to go live, traffic is redirected to the blue environment
- When the process is finished, the names of the environments are swapped

#### 11.7 Toggling Features
- You can tie new features to a software flag or configuration setting
- This way you can decouple deployment and releases
- There are many ways to implement this, flags can be command line flags, for instance

#### 11.8 Live Schema Changes
- Sometimes you expect database schema changes
- One way to deal with this is add new required fields to the database on one release. You then have both new and old fields simultaneously. You can then remove old fields on further releases

#### 11.9 Live Code Changes
- Generally frowned upon although some languages make this possible (Erlang)
- Very case specific

#### 11.10 Continuous Deployment
- Every release that passes tests is deployed to production automatically
- Author encourages you to have this as a goal
- This requires: Continuous integration, continuous delivery, continuous deployment and automating any other testing, approval and code push processes
- Continuous delivery results in packages that are production ready, but the decision to actually deploy these packages in production is a business one
- It can be risky to always push to production, and it is! One precaution is to use CD only with noncritical environments, such as beta/staging
- Or you can use this in production, but only with subsystems that have high confidence in their releases, such as a certain API. Web UIs usually benefit from manual testing
- Book makes a very good point: Make metrics of the things that you think are stopping you from achieving CD, such as build health, test comprehensiveness, schedules (such as financial deadlines) and so on. The book goes in great detail and has a comprehensive list, very cool
- Remember: Operations should be based in data and science. Quantify your gut feelings, don't just have them be feelings
- CD is non trivial and easier if done at the beginning of new projects when they're small

#### 11.11 Dealing with Failed Code Pushes
- You can roll back, or you can fix what's wrong and roll forward. Author is critical of bugfixes

### 11.12 Release Atomicity
- If certain systems are tightly coupled, test them and release them as a set (certain versions for system A and B, for instance)
- If components are loosely coupled, they can be tested independently and pushed at their own velocity
- Ideally you want your components to be loosely coupled