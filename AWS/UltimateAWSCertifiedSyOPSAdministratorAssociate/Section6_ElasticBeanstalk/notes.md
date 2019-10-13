- Chapter status: Reviewd the theory carefully, missing exam

### 47. Beanstalk intro 
- Usually a developer exam topic, but for the sysops one you need to know a few things

### 48. Beanstalk overview
- As a dev, you don't wanna manage infra, you just wanna deploy!
- Beanstalk, will:
- Manage infrastructure
- Deploy code
- Configure all the databases, load balancers, etc
- Scale it for you
- Most web apps have the same architecture (ALB + ASG)
- All developers want is for their code to run!
- Possibly, consistently across different applications and environments

#### AWS ElasticBeanStalk Overview
- ElasticBeanStalk is a developer centric view of deploying an application on AWS
- It uses all the components we've seen before: EC2, ASG, ELB, RDS, etc
- But it's all in one view that's easy to make sense of!
- We still have full control over the configuration 
- BeanStalk itself is free, you only pay for the underlying instances

#### ElasticBeanStack
- Managed service
    - Instance configuration / OS is handled by beanstalk 
    - Deployment strategy is configurable but performed by ElasticBeanStalk
- Just the application code is the responsibility of the developer 
- Three architecture models:
    - Single instance: Good for dev
    - LB + ASG: Great for production or pre-production environments
    - ASG only: Create for non-web apps in production (workers, etc)

#### ElasticBeanStack, part deux
- ElasticBeanStalk has three components:
    - Application
    - Application version: Each deployment gets assigned a version
    - Environment name (dev, test, prod): free naming
- You deploy application versions to environments and then can promote application versions to the next environment
- Rollback feature to previous application version 
- Full control over lifecycle of environments 

#### ElasticBeanStack, part trois
- Support for many platforms
    - Go, Java SE, Java with TomCat, .NET on Windows server with IIS, Node.js, PHP, Python, Ruby, Packer Builder
    - Single/Multi Container docker, preconfigured docker
    - If not supported, you can also write your custom platform (advanced) 


### 49. Beanstalk Hands on
- Beanstalk service! Get started
- Name the app, select platform
- Sample application
- In the backend, beanstalk uses cloudformation to create the environment, you can actually check the cloudformation code in the cloudformation service
- In the background, elasticbeanstalk will create all the AWS components for us (EC2 instance, ASG, ELB and maybe something else)
- By uploading a new version, we update the application
- An application can have multiple environments (you can create it in the beanstalk application windor, top right)
- App > Configuration > You can finetune a bunch of configurations, you can change the AMI for instance


### 50. Beanstalk Deployment Modes
- Very popular on the exam! 
- The exam will ask you about which deployment modes beanstalk has and which are better for which situation
- Single instance deployment: Great for dev
    - One EC2 instance, one elastic IP, one ASG, maybe one RDS instance. In one AZ.
- High availability with load balancer: Great for prod
    - ASG, multiple AZs, one or several EC2 instances per AZ,  


#### Beanstalk Deployment Options for Updates
- All at once (deploy all in one go) - fastest, but instances aren't available to serve traffic for a bit (downtime)
- Rolling: Update a few instances at a time (bucket), and then move onto the next bucket once the first bucket is healthy
- Rolling with additional batches: Like rolling, but spins up new instances to move the batch (so that the old application is still available)
- Immutable: Spins up new instances in a new ASG, deploys version to these instances, then swaps all the instances when everything is healthy

#### Elastic Beanstalk Deployment All at once
- All the instances that contain v1 will stop the application (?), and then we'll be running v2. All instances at the same time.
- Fastest deployment
- Application has downtime
- Great for quick iterations development environment
- No additional cost

#### Elastic Beanstalk Deployment Rolling
- Application will be running below capacity
- We can set the bucket size
- 4 instances running v1
- On 2 instances, the application will be stopped, but the other 2 instances will continue running v1
- Then the first 2 instances will be updated, which will then be running v2
- Then the other 2 instances running v1 will also be updated to run v2
- At some point during the deployment, the application will be running both versions simultaneously
- No additional cost
- If you set a very small bucket size and you have hundreds of instances, it can be a long deployment

####  Elastic Beanstalk Deployment Rolling with additional batches
- Application is running at capacity
- Can set the bucket size
- Application will run both versions simultaneously
- Small additional cost
- Additional batch is removed at the end of second deployment
- Longer deployment, good for prod
- We have 4 v1 instances
- First thing that happens is we get 2 additional instances, now running v2, so we have a total of 6 instances
- Then 2 instances running v1 are removed 
- Then 2 more instances running v2 are added
- Repeat until all instances are of type v2
- We're always running at capacity, sometimes over capacity. Possible exam question: Is there an additional cost to this? Answer: Yes

#### Elastic Beanstalk Deployment Immutable
- Zero downtime
- New code is deployed to new applications on a temporary ASG
- High cost, double capacity
- Longest deployment
- Quick rollback in case of failures (just terminate the ASG)
- Great for prod
- Current ASG: 3 applications running v1
- A new, temp ASG will be created runnign only 1 v2 instance at first
- If this instance works and it passes the healthchecks, it will launch all the remaining ones (in this case 3 v2 applications) on the new ASG
- After a while (?) all v2 applications will be moved from the temporaty ASGs to the current ASGs, so you'll have a total of 6 applications (3 v1, 3 v2)
- After that, v1 applications will be terminated

#### Elastic Beanstalk deployment blue/green
- Not a direct feature of elastic beanstalk
- Zero downtime and release facility 
- Create a new "stage" environment and deploy v2 there
- The new environment (green) can be validated independently and rolled back if issues arise
- Route 53 can be set up using weighted policies to redirect a little bit of traffic to the stage environment  
- Once all is ready, you can swap URLs on the beanstalk console, so you'll go live
- Not a direct feature, very manual. Some doc talks about it, some doesn't. But it's there.

#### Elastic Beanstalk Deployment Summary from AWS DOC
- `https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.deploy-existing-version.html`
- Author recommends having a read

### 51. Beanstalk Deployment Modes Hands-on
- Sample app is single instances, we need to edit that on the UI to be "load balancer" type
- LB and ASG will be created by beanstalk, neat
- If degraded, you can click on causes to see why
- To update, click on update and deploy, you upload the app then
- On the update and deploy tab, you can select the deployment policy (All at once, rolling, rolling with additional batch, immutable)
- You can also create new environments: Action > New environment
- For blue/green: Actions > Swap environment URLs

### 52. Beanstalk for SysOps
- Beanstalk can put your app logs directly into CloudWatch logs
- You manage the application, AWS will manage the underlying infrastructure
- You must know the different deployment modes for your application (see above)
- Route 53 question: You can put an ALIAS or CNAME on top of your beanstalk URL
- You're not responsible for patching the runtimes (AWS does it)
- Questions are quite basic compared to the dev ones

#### How beanstalk deploys applications
- Trick question time!
1. EC2 instance has base AMI (can be configured)
2. EC2 gets the new code of the app
3. EC2 will resolve the application dependencies (can talk a while)
4. Apps get swapped on the EC2 instance
- Resolving dependencies can take a long time.
- Popular question: How do you make the time resolving dependencie shorter? How do you get an app deployed as fast as possible?
- Answer: User the Golden AMI to solve the problem

#### Golden AMI
- Gloried term 
- If your application has a lot of application or OS dependencies and you want to deploy as quickly as possible, you should create a Golden AMI
- Golden AMI = Standardized company-specific AMI with:
    - OS dependencies
    - App dependencies
    - Maybe some company-wide software
    - It's just an AMI with a bunch of dependencies installed
- By using a Golden AMI to deploy to beanstalk (in combination with blue/green new ASG deployments) or application won't need to resolve dependencies or take a long time to configure!

#### Troubleshooting Beanstalk
- If the health of your environment changes to red, try the following
    - Review environment events
    - Pull logs to view recent log entries
    - Roll back to a previous, working version of the application
- When accessing external resources, make sure the security groups are correctly configured (if you have an RDS db, make sure you can reach it and that it can be reached)
- In case of command timeouts, you can increase the deployment timeout

### 53. Beanstalk clean up

### Quiz!
