### 22: Intro
```
Uh, we're gonna look at EC2 system's manager and AWS OpsWorks!
```

### 23: System Manager Overview
```text
Helps you manage your EC2 and on-premise systems at scale
Get operational insights about the state of your infrastructure, detect problems easily.
Patching automation for enhanced compliance (exam tip!)
Works for Windows and Linux
Integrated with CloudWatch metrics / dashboards
Integrated with AWS
Free!

Features:
Resource groups
Insights
Parameter Store
Automation, run commands
Session manager
Patch manager
Maintenance windows
State Manager

Most important: Run Command, Patch Manager, Run command and Parameter Store
Exam tip: How to patch systems? Use systems manager
SSM is an AWS service. You need to have the SSM agent installed on your machines. It is installed by default on AWS Linux and on some Ubuntu AMIs.
You can also have the SSM agent on on premise VMs.
If an instance can't be controlled with SSM, it's probably an issue with the SSM agent.
Also make sure the EC2 instances have a proper IAM role to allow SSM actions
```

### 24: Start EC2 instances with SSM agent
```text
New EC2 instances - Create 3
New IAM role > EC2 role > AmazonEC2RoleforSSM
New security group > Name it SSMManagerInstances > Remove all inbound rules!

After instances are launched:
AWS Systems Manager > Managed Instances > (Wait a minute or two for your EC2 instances to start up) > You see your instances!
(Author took no further action at this point)
```

### 25: AWS Tags & SSM Resource groups
```text
Tags: Key value pairs that can be attached to many resource kinds. Most commonly EC2, but anything else also goes. They're arbitrary, you can call them whatever you want (team, name, environment...)
Used for: Resource grouping, automation, cost allocation. And also used by system's manager!
Better to have too many than too few!

Resource groups on SSM are managed through tags
Went on EC2 and manually tagged 3 instances

Go to Resource Groups > Create group > Tag based > Grouping Criteria
Resource type: AWS::EC2::Instance
Tag: Environment: Dev
```

### 26: SSM Documents & SSM Run Command
```text
You can define a SSM Document.
In JSON or YAML
You define parameters
You define actions
Many documents already exist in AWS

Your documents can be applied to a bunch of things apparently (State manager, Patch Manager, Automation, Run Command and Parameter Store)
There are a lot of already existing documents :o

Let's run AWS Systems Manager to run a command. You execute a document (aka script) across multiple instances. Integrated with AIM and CloudTrail. No need for SSH.
Uuuh, let's create a document!
Name it whatever you want
Type Command Document
Let's go for YAML. The document is named document-install-apache.yml in the same directory as these notes.

I then selected the instances manually and ran it and it ran and now I have a apache installed. Cool!
```

### 27: SSM Inventory & Patches
```text
Exam question: We need to patch all our EC2 instances right meow! How do we do this?
Inventory > List software on an instance
Inventory + Run command > Patch software
Patch Manager + Maintenance window > Patch OS
Patch Manager > Gives you compliance
State manager > Ensure that instances are in a consistent state

--- FROM THIS POINT ON ---
I realized this was going to take way too long if I went detail by detail studying for the certification.
So I'm just gonna jump from lecture to lecture taking minimal notes, if any at all.
I'm more interested in an overview than to actually do everything
```

### 28: SSM Security Shell
```text
This allows you to start a secure shell on your vm but doesn't use SSH access or bastion hosts. Neat!
It uses the SSM agent to do that.
Only works for EC2
On PREM it might be supported soon
It will log all the actions on you to S3 and Cloudwatch logs
Cloud trail can intercept StartSession events

You can log the things on the security shell on cloudwatch logs. Pretty cool
```

### 29: What if I lose my EC2 SSH key?
```text
Traditional method, If EBS backed: Stop instance, detach root volume, attach it to another instance, add your other key in there as you have access to the data volume.
Once done, move volume back to stopped instance
Start instance, SSH in again

New method, if instance is EBS backed, using SSM.
Run authomation AWSSupport-ResetAccess, an automation. document in SSM.

If it's instance store backed EC2, AWS recommends termination

Author's protip: Use session manager to edit authorized_keys and give yourself access again.
```

### 30: SSM Parameter Store Overview
```text
Secure storage for configuration and secrets
Can use seamless encryption with KMS.
Serverless, scalable, durable, easy SDK, free
Allows version tracking of configs / secrets
Configuration management using path & IAM
You can be notified with CloudWatch events
Integrates with CloudFormation

Use GetParameters or GerParametersByPath
Author recommends to use a hierarchy (departments, teams, environments)
```

### 31: SSM Parameter Store Hands on (CLI)
```text
Very nice, you can create plaintext and unencrypted parameters on SSM Parameter Store

You can use the aws command line tool to get parameters from a directory or a given "file" (a single parameter):
aws ssm get-parameters-by-path --path /my-app/dev
Also: --with-decryption
````

### 32: AWS OpsWorks Overview
```text
Uses Chef & Puppet, works great with EC and on premises VM
AWS OpsWorks = Another name for Chef and Pupper.
Kinda of an alternative to AWS SSM.
No knowledge of Chef and Puppet is required to pass the exam, but author says:
Chef & Pupper = AWS OpsWorks

They leverage Recipes / Manifests

Chef / Puppet have a lot of similarities with SSM / Beanstalk / CloudFormation but they're open source tools that work cross-cloud
```

### 33: Quiz!