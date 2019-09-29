- Chapter status: Took notes attentively, got the right!

### 22: Intro
- Amazon EC2 Systems Manager
- AWS OpsWorks

### 23: System Manager Overview
- Helps you manage your EC2 and on-premise systems at scale
- Get operational insights about the state of your infrastructure
- Easily detect problems easily.
- Patching automation for enhanced compliance (exam tip: there are questions around patching)
- Works for both Windows and Linux
- Integrated with CloudWatch metrics / dashboards
- Integrated with AWS Config
- Free!
- Possible exam question: We need to patch all the OS of our EC2 instances, how do we do this? Use AWS Systems manager!

#### Features
- Resource groups
- Insights
    - Insights Dashboard
    - Inventory: Discover and audit the installed software
    - Compliance
- Parameter Store
- Automation, run commands
- Session manager
- Patch manager
- Maintenance windows
- State Manager to define how your OS and applications should be configured
- Most important: Run Command, Patch Manager, Run command and Parameter Store

#### How System Manager works
- SSM is an AWS service. You need to have the SSM agent installed on your machines.
- It is installed by default on AWS Linux and on some Ubuntu AMIs.
- If an instance can't be controlled with SSM, it's probably an issue with the SSM agent or IAM (you need proper permissions to talk to SSM)
- Exam tip: How to patch systems? Use systems manager
- You can also have the SSM agent on on premise VMs.
- Make sure the EC2 instances have a proper IAM role to allow SSM actions

### 24: Start EC2 instances with SSM agent
- New EC2 instances - Create 3, all default settings with Amazon Linux
- New IAM role > EC2 role > AmazonEC2RoleForSSM > AmazonEC2RoleforSSM
- New security group > Name it SSMManagerInstances > Remove all inbound rules!
- After instances are launched:
- AWS Systems Manager > Managed Instances > (Wait a minute or two for your EC2 instances to start up) > You see your instances!

### 25: AWS Tags & SSM Resource groups

- Tags: Key value pairs that can be attached to many resource kinds.
- Most commonly EC2, but they can be attached to many services
- They're arbitrary, you can call them whatever you want (team, name, environment...)
- Used for
    - Resource grouping
    - Automation
    - Cost allocation
    - And also used by system's manager!
    - Better to have too many than too few!

#### AWS Systems Manager Resource groups
- Create, view or manage local groups of resources with tags
- Allows createion of logical groups such as
    - Applications
    - Layers of app stack
    - Prod / Dev environment
- Resource groups are a regional service 
- Resource groups work with EC2, S3, DynamoDB, Lambda and a few others

#### Hands on
- Tagged EC2 instances with `Name` and `Environment`
- Go to Resource Groups > Create group > Tag based > Grouping Criteria
- Resource type: AWS::EC2::Instance
- Tag: Environment: Dev


### 26: SSM Documents & SSM Run Command
- You can define a SSM Document to perform an action
- Documents can be in JSON or YAML
- You define parameters
- You define actions
- Many documents already exist in AWS

Your documents can be applied to many different things
- State manager
- Patch Manager
- Automation
- Run Command
- It can also reference the Parameter Store
- There are a lot of already existing documents :o
- If you use SSM at scale, it's adviseble to write your own documents

#### AWS Systems Manager Run Command
- You execute a document (aka script) or just run a command
- Run command across multiple instances (using resource groups)
- Rate Control / Error Control
- Integrated with IAM & CloudTrail
- No need for SSH
- Results in the console

#### Hands on
- AWS System Manager > Documents > Create Document
- Created a document by copying and pasting a document that installs apache
- Run commands > Run a command > Select your document
- Manually select instances
- You can have the output of your commands in CloudWatch!

### 27: SSM Inventory & Patches
- Possible exam question: We need to patch all our EC2 instances right meow! How do we do this? Using System's manager! 
- Inventory > List software on an instance
- Inventory + Run command > Patch software
- Patch Manager + Maintenance window > Patch OS
- Patch Manager > Gives you compliance
- State manager > Ensure that instances are in a consistent state

#### Hands on
- AWS Systems Manager > Inventory 
- Enable inventory on all instances
- You can see inventory information on every single instance. Neat

--- FROM THIS POINT ON ---
I realized this was going to take way too long if I went detail by detail studying for the certification.
So I'm just gonna jump from lecture to lecture taking minimal notes, if any at all.
I'm more interested in an overview than to actually do everything

##### Patch manager hands on
- SSM > Patch Manager > Configure Patching
- Selected instances manually
- Set a patching schedule (or patch right now)
- And it "patches" instances. Uh-oh

##### Compliance
- The outcome of the patch scan went into the compliance tab on SSM
- Exam is gonna be very high level, as in "We need to patch instances, how do we do it?" Answer is, using systems manager! :D


### 28: SSM Secure Shell
- Not in the exam yet! 
- Session Manager allows you to start a secure shell on your VM 
- Does not use SSH access or bastion hosts. Neat!
- It uses the SSM agent to do that
- Only works for EC2, On premise  it might be supported soon 
- Logs all actions done through secure shell to S3 and CloudWatch Logs
- For this to work we need IAM Permissions: Access to SSM + Write to S3 + Write to CloudWatch 
- Cloud trail can intercept StartSession events, so you can have audits to see who's starting sessions on your infrastructure
- AWS Secure Shell vs SSH
    - No need to open the port 22 at all
    - No need for bastion hosts
    - All commands are logged to S3 / CloudWatch
    - Access to secure shell is done through User IAM, not SSH keys

#### Hands on 
- SSM > Session Manager > Preferences > Set up S3 storage or CloudWatch storage
- Created CloudWatch Log group
- On Session Manager: Send logs to CloudWatch

### 29: What if I lose my EC2 SSH key?
- Common exam question: What if I lost my SSH key for EC2?
- Traditional method If the instance is EBS backed
    - Stop instance
    - Detach root volume
    - attach it to new EC2 instance as a data volume
    - Modify ~/.ssh/authorized_keys with your new key
    - Move the volume back to the stopped instance
    - Start the instance and you can SSH in there again, yay!
- New method, using SSM, only if the instance is EBS backed
    - Run the AWSSupport-ResetAccess automation document in SSM
- If your EC2 instance is Instance Store backed 
    - You can't stop the instance (otherwise data is lost) - AWS recommends termination
    - Author's protip: Use session manager to edit authorized_keys and give yourself access again.

### 30: SSM Parameter Store Overview
- Author claims it's revolutionary and underutilized!!1!one
- Popular on the exam
- Secure storage for configuration and secrets
- Optional seamless encryption with KMS.
- A good place to put your database passwords, for instance
- Serverless, scalable (tens of thousands of parameters!), durable, easy SDK, free
- Allows version tracking of configs / secrets
- Configuration management using path & IAM
- You can be notified with CloudWatch events
- Integrates with CloudFormation

#### Diagram
- Your application on your EC2 instance talks to SSM Parameter store and goes
- "Hey I'd like to have this password in plaintext"
- The parameter store checks the IAM permissions of the instance and returns the password (or not)
- If your instance wants an encrypted configuration, SSM Parameter store will also call AWS KMS to decrypt the encrypted configuration, and then return it to the instance


#### AWS Parameter Store Hierarchy
- my department
    - my app
        - dev
            - db url
            - db password
        - prod
            - db url
            - db password
    - other app
- other department

Author recommends to use a hierarchy (departments, teams, environments). You can organize this tree in any way you want (dept, environment, app, etc)

### 31: SSM Parameter Store Hands on (CLI)
- Systems Manager > Parameter Store > Create parameter
- Very nice, you can create plaintext and unencrypted parameters on SSM Parameter Store
- You can use the aws command line tool to get parameters from a directory or a given "file" (a single parameter):
- `aws ssm get-parameters-by-path --path /my-app/dev`
- Also: --with-decryption


### 32: AWS OpsWorks Overview
- Chef & Pupper help you perform server configuration automatically or repetitive actions
- They work great with EC and on premises VM
- AWS OpsWorks = Managed Chef and Pupper.
- It's an alternative to AWS SSM.
- No knowledge of Chef and Puppet is required to pass the exam, but author says:
- Any time in the exam you see something like "we need to apply/use chef/puppet, which service do you use?"
- AWS OpsWorks
- Chef & Puppet = AWS OpsWorks

#### Quick word on Chef and Puppet
- They help managing configuration as a code
- Helps in having consistent deployments
- Works with Linux / Windows
- Can automate: user accounts, cron, ntp, packages, services...
- They leverage Recipes / Manifests
- Chef / Puppet have a lot of similarities with SSM / Beanstalk / CloudFormation but they're open source tools that work cross-cloud

### 33: Quiz!