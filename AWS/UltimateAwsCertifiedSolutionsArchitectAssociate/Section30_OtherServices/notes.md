- Note from Dani: Went through this one a bit quickly, revise it a bit

### 256. Other Services Section Introduction
- Basic things that might show up at the exam at a high level

### 257. CICD Introduction

#### Continuous Integration
- Developers push the code to a code repository often (GitHub / CodeCommit / Bitbucket / etc...)
- A testing / build server checks the code as soon as it’s pushed (CodeBuild / Jenkins CI / etc...)
- The developer gets feedback about the tests and checks that have passed / failed
- Find bugs early, fix bugs
- Deliver faster as the code is tested
- Deploy often
- Happier developers, as they’re unblocked

#### Continuous Deployment
- Ensure that the software can be released reliably whenever needed.
- Ensures deployments happen often and are quick
- Shift away from "one release every 3 months" to "5 releases a day"
- That usually means automated deployment
    - CodeDeploy
    - Jenkins CD
    - Spinnaker
    - Etc...

#### Technology Stack for CICD
- From AWs's perspective, very vendor-heavy this.
- Code > Build > Test > Deploy > Provision
- Code (AWS CodeCommit) > Build (AWS CodeBuild) > Test (AWS CodeBuild)  > Deploy (Beanstalk or CodeDeploy + EC2 with CloudFormation) > Provision (Beanstalk or CodeDeploy + EC2 with CloudFormation) 
- To orchestrate all of this maybe you need a pipeline Orchestrator.
- You can use AWS CodePipeline for this

### 258. CloudFormation intro

#### Infrastructure as Code
- Currently, we have been doing a lot of manual work
- All this manual work will be very tough to reproduce:
    - In another region
    - in another AWS account
    - Within the same region if everything was deleted
- Wouldn’t it be great, if all our infrastructure was... code?
- That code would be deployed and create / update / delete our infrastructure

#### What is CloudFormation
- CloudFormation is a declarative way of outlining your AWS Infrastructure, for any resources (most of them are supported).
- For example, within a CloudFormation template, you say:
    - I want a security group
    - I want two EC2 machines using this security group
    - I want two Elastic IPs for these EC2 machines
    - I want an S3 bucket
    - I want a load balancer (ELB) in front of these machines
- CloudFormation creates those for you, in the right order, with the exact configuration that you specify

#### Benefits of AWS CloudFormation (1/2)
- Infrastructure as code
    - No resources are manually created, which is excellent for control
    - The code can be version controlled for example using git
    - Changes to the infrastructure are reviewed through code
- Cost
    - Each resources within the stack is tagged with an identifier so you can easily see how much a stack costs you
    - You can estimate the costs of your resources using the CloudFormation template
    - Savings strategy: In Dev, you could automation deletion of templates at 5 PM and recreated at 8 AM, safely
    
#### Benefits of AWS CloudFormation (2/2)
- Productivity
    - Ability to destroy and re-create an infrastructure on the cloud on the fly
    - Automated generation of Diagram for your templates!
    - Declarative programming (no need to figure out ordering and orchestration)
- Separation of concern: create many stacks for many apps, and many layers. Ex:
    - VPC stacks
    - Network stacks
    - App stacks
- Don’t re-invent the wheel
    - Leverage existing templates on the web!
    - Leverage the documentation

#### How CloudFormation Works
- Templates have to be uploaded in S3 and then referenced in CloudFormation
- To update a template, we can’t edit previous ones. We have to re-upload a new version of the template to AWS
- Stacks are identified by a name
- Deleting a stack deletes every single artifact that was created by CloudFormation.

#### Deploying CloudFormation templates
- Manual way:
    - Editing templates in the CloudFormation Designer
    - Using the console to input parameters, etc
- Automated way:
    - Editing templates in a YAML file
    - Using the AWS CLI (Command Line Interface) to deploy the templates
    - Recommended way when you fully want to automate your flow

#### CloudFormation Building Blocks
- Templates components
    1. Resources: Rour AWS resources declared in the template (MANDATORY)
    2. Parameters: The dynamic inputs for your template
    3. Mappings: The static variables for your template
    4. Outputs: References to what has been created
    5. Conditionals: List of conditions to perform resource creation
    6. Metadata
- Templates helpers:
    1. References
    2. Functions

#### Note: This is an introduction to CloudFormation
- It can take over 3 hours to properly learn and master CloudFormation
- This lecture is meant so you get a good idea of how it works
- The exam expects you to understand how to read CloudFormation

#### CloudFormation - StackSets
- Create, update, or delete stacks across multiple accounts and regions with a single operation
- Administrator account to create StackSets
- Trusted accounts to create, update, delete stack instances from StackSets
- When you update a stack set, all associated stack instances areupdated throughout all accounts and regions

### 259. CloudFormation Hands-On

### 260. CloudFormation - Extras

### 261. ECS Introduction 
- ECS is a container orchestration service
- ECS helps you run Docker containers on EC2 machines
- ECS is complicated, and made of:
    - "ECS Core": Running ECS on user-provisioned EC2 instances
    - Fargate: Running ECS tasks on AWS-provisioned compute (serverless) (probably not on the exam)
    - EKS: Running ECS on AWS-powered Kubernetes (running on EC2)
    - ECR: Docker Container Registry hosted by AWS
    - ECS & Docker are very popular for microservices
- For now, for the exam, only “ECS Core” & ECR is in scope
- IAM security and roles at the ECS task level

#### What is Docker
- Docker is a “container technology”
- Run a containerized application on any machine with Docker installed
- Containers allows our application to work the same way anywhere
- Containers are isolated from each other
- Control how much memory / CPU is allocated to your container
- Ability to restrict network rules
- More efficient than Virtual machines
- Scale containers up and down very quickly (seconds)

#### AWS ECS - Use cases
- Run microservices
    - Ability to run multiple docker containers on the same machine
    - Easy service discovery features to enhance communication
    - Direct integration with Application Load Balancers
    - Auto scaling capability
- Run batch processing / scheduled tasks
- Schedule ECS containers to run on On-demand / Reserved / Spot instances
- Migrate applications to the cloud
- Dockerize legacy applications running on premise
- Move Docker containers to run on ECS

#### AWS ECS Concepts
- ECS cluster: Set of EC2 instances
- ECS service: Applications definitions running on ECS cluster
- ECS tasks + definition: Containers running to create the application
- ECS IAM roles: Roles assigned to tasks to interact with AWS

#### AWS ECS - ALB integration
- Popular at the exam!
- Application Load Balancer (ALB) has a direct integration feature with ECS called “port mapping”
- This allows you to run multiple instances of the same application on the same EC2 machine
- Use cases:
    - Increased resiliency even if running on one EC2 instance
    - Maximize utilization of CPU / cores
    - Ability to perform rolling upgrades without impacting application uptime

#### AWS ECS - ECS Setup & Config File
- Run an EC2 instance, install the ECS agent with ECS config file
- Or use an ECS-ready Linux AMI (still need to modify config file)
- ECS Config file is at /etc/ecs/ecs.config
- Remember that you have to change this config file to assign instances to your cluster, pull images from private registries, enable cloudwatch logs integration and finally to enable IAM roles

#### AWS ECR - Elastic Container Registry
- Store, managed and deploy your containers on AWS
- Fully integrated with IAM & ECS
- Sent over HTTPS (encryption in flight) and encrypted at rest


### 262. ECS - Extras
#### ECS - IAM Task Roles
- The EC2 instance should have an IAM role allowing it to access the ECS service (for the ECS agent)
- Each ECS task should have an ECS IAM task role to perform their API calls
- Use the "taskRoleArn" parameter in a task definition

#### Fargate
- When launching an ECS Cluster, we have to create our EC2 instances
- If we need to scale, we need to add EC2 instances
- So we manage infrastructure...
- With Fargate, it’s all Serverless!
- We don’t provision EC2 instances
- We just create task definitions, and AWS will run our containers for us
- To scale, just increase the task number. Simple! No more EC2

### 263. EKS - Overview
- Amazon EKS standards for Amazon Elastic Kubernetes Service
- It is a way to launch managed Kubernetes clusters on AWS
- Kubernetes is an open-source system for automatic deployment, scaling and management of containerized (usually Docker) application
- It’s an alternative to ECS, similar goal but different API
- EKS supports EC2 if you want to to deploy worker nodes or Fargate to deploy serverless containers
- Use case: if your company is already using Kubernetes on-premises or in another cloud, and wants to migrate to AWS using Kubernetes

### 264. Step Functions & SWF
#### Step functions
- Build serverless visual workflow to orchestrate your Lambda functions
- Represent flow as a JSON state machine (any time you see state machine, it's probably step functions)
- Features: sequence, parallel, conditions, timeouts, error handling...
- Can also integrate with EC2, ECS, On premise servers, API Gateway
- Maximum execution time of 1 year
- Possibility to implement human approval feature
- Use cases:
    - Order fulfillment
    - Data processing
    - Web applications
    - Any workflow

#### AWS SWF - Simple Workflow Service
- Coordinate work amongst applications
- Code runs on EC2 (not serverless)
- 1 year max runtime
- Concept of “activity step” and “decision step”
- Has built-in “human intervention” step
- Example: order fulfilment from web to warehouse to delivery
- **Step Functions is recommended to be used for new applications, except:** 
    - If you need external signals to intervene in the processes
    - If you need child processes that return values to parent processes
- SWF is old, you'll probably use step functions (that's how the exam is likely to ask things)  

### 265. EMR
- EMR stands for “Elastic MapReduce”
- EMR helps creating Hadoop clusters (Big Data) to analyze and process vast amount of data
- The clusters can be made of hundreds of EC2 instances
- Also supports Apache Spark, HBase, Presto, Flink...
- EMR takes care of all the provisioning and configuration
- Auto-scaling and integrated with Spot instances
- Use cases: data processing, machine learning, web indexing, big data...

### 266. AWS Glue
- Fully-managed ETL (Extract, Transform & Load) service
- Automating time consuming steps of data preparation for analytics
- Serverless, pay as you go, fully managed, provisions Apache Spark
- Crawls data sources and identifies data formats (schema inference)
- Automated Code Generation
- Sources: Amazon Aurora, RDS, Redshift & S3
- Sinks: S3, Redshift, etc...
- Glue Data Catalog: Metadata (definition & schema) of the Source Tables

### 267. OpsWorks
- Chef & Puppet help you perform server configuration automatically, or repetitive actions
- They work great with EC2 & On Premise VM
- AWS Opsworks = Managed Chef & Puppet
- It’s an alternative to AWS SSM
- No hands on here, no knowledge of chef and puppet needed
- In the exam: Chef & Puppet needed => AWS Opsworks

#### Quick word on Chef / Puppet
- They help with managing configuration as code
- Helps in having consistent deployments
- Works with Linux / Windows
- Can automate: user accounts, cron, ntp, packages, services...
- They leverage “Recipes” or ”Manifests”
- Chef / Puppet have similarities with SSM / Beanstalk / CloudFormation but they’re open-source tools that work cross-cloud

### 268. ElasticTranscoder
- Convert media files (video + music) stored in S3 into various formats for tablets, PC, Smartphone, TV, etc
- Features: bit rate optimization, thumbnail, watermarks, captions, DRM, progressive download, encryption
- 4 Components:
    - Jobs: what does the work of the transcoder
    - Pipeline: Queue that manages the transcoding job
    - Presets: Template for converting media from one format to another
    - Notifications: SNS for example
- Pay for what you use, scales automatically, fully managed

### 269. AWS Workspaces
- Managed, Secure Cloud Desktop
- Great to eliminate management of on-premise VDI (Virtual Desktop Infrastructure)
- On Demand, pay per by usage
- Secure, Encrypted, Network Isolation
- Integrated with Microsoft Active Directory

### 270. AppSync
- Store and sync data across mobile and web apps in real-time
- Makes use of GraphQL (mobile technology from Facebook)
- Client Code can be generated automatically
- Integrations with DynamoDB / Lambda
- Real-time subscriptions
- Offline data synchronization (replaces Cognito Sync)
- Fine Grained Security

### 271. Other Services: Cheat Sheet
  