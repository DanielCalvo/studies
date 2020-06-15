Chapter status: Redo lectures 28 and 29

### 11. AWS Regions and AZ
- AWS has regions all over the world!
- A region is a cluster of data centers
- Most AWS services are region-scoped

#### AWS Availability Zones
- Each region has many availability zones (usually 3, min 2, max 6) 
- Example: eu-west-1
    - eu-west-1a
    - eu-west-1b
    - eu-west-1c
- Each availability zone (AZ) is one or more discrete data centers with redundant power, networking and connectivity
- They're separate from each other, so that they're isolated from disasters
- But still connected with high bandwidth, ultra-low latency networking
- Some services are only offered in some regions. Check AWS's regional table for that!

### 12. IAM Introduction
- Identity and access management. Users!
- AWS security is defined in here, using:
    - Users
    - Groups
    - Roles
- Root account should never be used (only use it the first time)
- Users must be created with proper permissions
- Policies exist in IAM, they are written in JSON

#### IAM Overview part 2
- Users: Usually a physical person
- Groups: Functions or teams, usually. Contains users!
- Roles: Only for internal usage by AWS resources and services. Roles are given to machines!
- Policies: Define what users, groups and roles can do!
- IAM has a global view!
- MFA is available
- IAM has predefined "maaged policies"
- We'll see more on IAM policies in the future
- It's best to give users the minimal amount of permissions they need to do the job! Least privilege principle

#### IAM Federation
- Big enterprise can integrate their own authentication methods with IAM (like ADFS)
- This way, you can log in into AWS using company credentials
- Identity federation uses the SAML standard (AD)

#### IAM 101 Brain dump
- One user per person
- One role per application
- Never shared IAM credentials!
- Never, ever, write IAM credentials in code, EVER!
- Never commit your IAM credentials
- Never use the root account, except for the initial setup

### 13. IAM Hands on
- Added MFA 
- Added a user with the AdministratorAccess policy
- Created a group named Admin with the AdministratorAccess policy
- Applied IAM password policy

### 14. EC2 Introduction
- One of the most popular AWS offerings
- It mainly consists of:
    - Renting virtual machines (EC2)
    - Storing data on virtual drives (EBS)
    - Distributing load across machines (ELB)
    - Scaling the services using an auto-scaling group (ASG)

#### Hands on
- Launched an instance

### 15 - 19: Explanations about SSH, skipped!

### 20. EC2 Instance Connect
- Oooh, browser based ssh connection!
- Only works with Amazon Linux 2 for now
 
### 21. Introduction to Security Groups
- Security groups are how network security is set up in AWS
- They control how traffic is allowed into or out of EC2 instances
- It is the most fundamental thing to know when troubleshooting networking issues
- In this lecture, we focus on how to use the allow, inbound and outbound ports

### 22. Security Groups Deep Dive
- SGs are acting as a firewall on EC2 instances
- They regulate:
    - Access to ports
    - Authorized ranges in IPv4 and IPv6
    - Control of inbound network
    - Control of outbound network

#### Security Groups good to know
- Can be attached to multiple instances, and an instance can also have multiple security groups
- Locked down to a region / VPC combination
- It is not a setting on an instance, it lives "outside" if your EC2 instance
- It's good to maintain one separate security group just for SSH access
- If you application times out, it's probably a security group issue.
- If you receive a "connection refused", then it's an app error, not a security group error
- All inbound traffic is blocked by default
- All authorized traffic is authorized by default

#### Referencing Security Groups
- You can reference other security groups for access on a security group
- Like "On security group A, if you're coming from security group B, you're allowed!"

### 23. Private vs Public vs Elastic IP
- Stephan comments on the differentes between IPv6, IPv6, public and private IPs

#### Elastic IP
- When you stop and start an EC2 instance, it's public IP can change
- If you need a public IP for your instance, you need an Elastic IP
- An Elastic IP is a public IPv4 IP you own as long as you don't delete it
- You can attach it to one instance at a time
- With an Elastic IP address, you can mask the failure of an instance or software by rapidly remapping the address to another instance in your account, but that's not a common pattern
- You can have up to 5 Elastic IPs on your account, or ask AWS to increase that
- Stephan recommends attempting to avoid using Elastic IPs
    - They often reflect poor architectural decisions
    - Instead, use a random public IP and register a DNS name to it
    - Or later, use a Load Balancer, which is a much better pattern!
- By default, an EC2 instance:
    - Comes with a public IP to be reached from the internet
    - Comes with a private IP to be reached by other AWS services internally
    - When the machine is restarted, the public IP can change 

### 24. Private vs Public vs Elastic IP Hands On
- Created Elastic IP address
- Associated it with instance
- The instance's public IP will change to be the allocated Elastic IP (if it has a public IP)
- Released public IP address

### 25. Install Apache on EC2
- Launched instance
- `yup update`, `yum install httpd.x86_64`, `systemctl start httpd.server`
- Modified the security group to allow port 80 and saw the sample page

### 26. EC2 User Data
- It is possible to bootstrap our instances using an EC2 User data script.
- bootstrapping means launching commands when a machine starts
- That script is only run once at the instance first start
- EC2 user data is used to automate boot tasks such as:
    - Installing updates
    - Installing software
    - Downloading common files from the internet
    - Anything you can think of
- The EC2 User Data Script runs with the root user

#### EC2 User data hands on
- We want to make sure that this EC2 instance has an Apache HTTP server installed on it – to display a simple web page
- For it, we are going to write a user-data script
- This script will be executed at the first boot of the instance
- Let’s get hands on!

#### Hands on notes
- Under instance details and advanced options, there's "user data" text box
- Copied the contents of `ec2-basics-user-data.sh` in there!

### Mid way quiz!
- Got everything right, yay!

### 27. EC2 Instances Launch Types
- On Demand Instances: short workload, predictable pricing
- Reserved: (MINIMUM 1 year)
- Reserved Instances: long workloads
- Convertible Reserved Instances: long workloads with flexible instances
- Scheduled Reserved Instances: example – every Thursday between 3 and 6 pm
- Spot Instances: short workloads, for cheap, can lose instances (less reliable)
- Dedicated Instances: no other customers will share your hardware
- Dedicated Hosts: book an entire physical server, control instance placement

#### EC2 On Demand
- Pay for what you use (billing per second, after the first minute)
    - This is the default instance behaviour when you go on the EC2 dashboard and launch an instance 
- Has the highest cost but no upfront payment
- No long term commitment
- Recommended for short-term and un-interrupted workloads where you can't predict how the application will behave

#### EC2 reserved instances:
- Up to 75% discount compared to on-demand
- Pay upfront for what you use with long term commitment
- Reservation period can be 1 year or 3 years
- Reserve a specific instance type
- Recommended for steady usage applications (like a database for instance)

There's a variation of this called Convertible Reserved Instance
- You can change the EC2 instance type
- But the discount is only up to 54%

There are also scheduled reserved instances
- Launch a given instance within a time window that you reserve

#### EC2 spot instances
- Can get a discount of up to 90% compared to on demand
- You bid a price and get the instance as long as it is under the price
- Price varies based on offer and demand
- Spot instances are reclaimed with a 2 minute notification warning when the spot price goes above your bid
- Useful for batch jobs, big data analysis and work loads that are resilient to failure

#### EC2 dedicated hosts
- Physical dedicated EC2 server for you to use!
- Full control of EC2 instance placement
- Visibility into the underlying physical cores of the hardware
- Allocated for 3 years
- Expensive! But you control the hardware in some ways.
- Useful for software that bills you based on number of cores or companies that face certain regulatory / compliance needs

#### EC2 dedicated instances
- Instances running on hardware that's dedicated to you
- May share hardware with other instances in the same account
- No control over instance placement

#### Hotel Metaphor for instance types
- On demand: Coming in and staying in the hotel whenever you like and paying the full price
- Reserved: Planning ahead and if you're going to stay for a long time, the hotel might give you a discount
- Spot instances: The hotel allows people to bid for empty rooms and the highest bidder keeps the rooms. You can be kicked out any time if you get outbid! 
- Dedicated hosts: You book and entire building of the resort

### 28. [SAA-C02] Spot Instances & Spot Fleet
- This whole lecture was a bit confusing. Maybe review this outside Stephan's explanation
- Follow up from Dani: Review

### 29. EC2 Instances Launch Types Hands On
- Hands on was OK, but do it yourself and take your own notes in here as this whole chapter as a lot of info
- Follow up from Dani: Review

### 30. EC2 Instance Types Deep Dive
- R: Applications that need a lot of ram, such as in memory caches
- C: Apps that use CPU, compute/databases
- M: Balanced, think general/web app
- I: I/O focused (dbs)
- G: GPU. Video rendering / machine learning
- T2/T3 Burstable: Can burst CPU usage (up to a capacity)
- T2/T3 unlimited: Unlimited burst.
- Suggested by the author: https://ec2instances.info/

#### Burstable instances (T2/T3)
- Burst mean that the instance can burst CPU usage for extra processing, but only for a limited amount of time. During this time you use "burst credit". You regain "burst credit" over time when not using the extra processing power
- When the burst credits are over, expect worse CPU performance
- Burstable instances can be good to handle unexpected traffic and to know you can handle bursts well
- If your instance constantly runs low on credits, you need to move to a different, non-burstable instance 
- In CloudWatch you can see your burst usage / credits
- On T2/T3 unlimited instances you have "unlimited burst", you can pay more for these and you never have performance penalties. It can be expensive however.

### 31. EC2 AMIs
- AWS comes with a bunch of base images such as Ubuntu, Fedora, Redhat, Windows etc
- These images can be customized at runtime using EC2 user data
- But what if you wanted to create your own image, ready to go?
- That's an AMI: An image used to launch instances

#### Why use custom AMI?
- Using a custon built AMI can provide the following advantages
    - Pre installed packages can come installed
    - Faster boot time (no need for ec2 user data)
    - Machines can come configured with monitoring/enterprise software
    - Control of maintenance and updates of AMIs over time
    - AD integration out of the box
    - Installing your app ahead of time (for faster deploys when auto-scaling)
    - Or you can use someone elses AMI that's optimized for a given purpose, such as an app or a DB

- Remember: AMIs are build for a specific AWS region!
- You can create a custom AMI! With pre-installed packages and any other thing you might need

#### Using Public AMIs
- You can leverasge AMIs from other people
- You can also pay for other people's AMI by the hour
    - These people have optimized the software
    - The machine is easy to run and configure
    - You "rent expertise" from the AMI creator
- AMIs can be found and published on the Amazon Marketplace

#### Warning
- Don't use an AMI you don't trust! 
- Some AMIs might containr malware or might not be secure!

#### AMI Storage
- When you create an AMI it will reside and take space in S3, but you won't see it in the console.
- Amazon S3 is a durable, cheap and resilient storage where most of your backups will live (but you won't see them in the S3 console)
- By default all AMIs are private and locked for your account and region 
- Make sure you remove AMIs you don't use as you'll be charged for them
- You can also make your AMIs public and share them with other AWS accounts or sell them on the AWS Marketplace

#### AMI Pricing
- Since AMIs live in S3, you get charged for the actual space it takes in Amazon S3
- S3 Pricing in us-east-1
    - First 50TB/Month: 0.023 per GB
    - Next 450TB/Month: 0.022 per GB
- Overall it's quite inexpensive to store private AMIs
- Make sure to remove AMIs you don't use!

### 32. EC2 AMI Hands On
- Launch an EC2 instance, do whatever you want on it (like install apache with a hello world through web)
- Right click on your running instance > image > create image
- Once created, you can go on the AMI menu and "Copy AMI" to a different region. A given AMI is locked to a given region.
- You an also modify the image permissions, de-register it or remove it.
- Follow up from Dani: Can you script AMI creation?

### 33. Cross Account AMI Copy
- Popular at the exam!
- You can share an AMI with another AWS account
- Sharing an AMI does not affect the ownership of the AMI
- If you copy an AMI that has been shared, you then become the owner of that AMI.
- To copy an AMI that was shared with you from another account, the owner of the source AMI must grant you read permission for the storage that backs the AMI
    - Either the associated EBS snapshot (for an Amazon EBS backed AMI)
    - Or the associated S3 bucket (for an instance store-backed AMI)
- Limits
    - You can't copy an encrypted AMI that was shared with you from another account. Instead, If the underlying snapshot and encryption key were shared with you, you can copy the snapshop and re-encrypt it with a key of your own. You own the copied snapshot and can register it as a new AMI
    - You can't copy an AMI with an associated **billingProduct** code that was shared with you from another account. This includes Windows AMIs and AMIs from the AWS marketplace. To copy a shared AMI with a billingProduct code, launch an EC2 instance in your account using the shared AMI and then create an AMI from the instance. Popular at the exam!
- Reemphasizing: To copy an AMI with a billingProduct code, first launch an EC2 instance from that AMI and then you make an AMI from that EC2 instance 

#### Quick hands on
- Create an AMI from a running instance
- To share AMI
    - Right click > Modify Image Permissions
    - You can make it public, or share it with an AWS Account number
    - If you tick the box for Create volume permissions, it means those accounts will be able to create a copy of your AMI
    - If you don't tick it, they can still launch an instance from your AMI and from that instance, create their own AMI. It just blocks "Copy AMI" directly

### 34. EC2 Placement Groups
- Sometimes you want control over the EC2 placement strategy
- This strategy can be defined using placement groups
- When you create a placement group, you specify one of the following strategies for the group
    - Cluster: Clusters instances into a low-latency group in a single AZ
    - Spread: Spreads instances across underlying hardware (max 7 instances per AZ)
    - Partition: Spreads instances across many different partitions (which rely on different racks) within an AZ. Scales to hundreds of instances per group

#### Cluster
- Clusters instances in a low-latency group inside a single availability zone (fast, but risky).
- Same rack, same AZ.
- Pro: Great network (10gbps)
- Con: If rack fails, all instances fail at the same time
- Use case: Application that needs low latency and high network throughput (like a big data job)

#### Spread
- Spreads instances across underlying hardware (max 7 instances per group per AZ)
- Pros: Can span across Availability zones (AZ), reduced risk of simultaneous failure, EC2 instances are on separate hardware.
- Cons: Limited to 7 instances per AZ per placement group
- Use case: Application that needs to maximize high availability, critical applications in which each instance must be isolated from failure

#### Partition
- Partition: Spreads instances across many different partitions (different sets of racks) within an AZ. Scales to 100s of EC2 instances per group.
- Partitions are a set of racks, on each partition you're going to have sets of EC2 instances.
- The instances in a partition do not share racks with instances in other partitions
- EC2 instances can get access to the partition information metadata
- Use cases: Distributed big data applications

#### Hands on
- EC2 dashboard > Placement groups > Create placement groups (with the strategy you want)
- When creating an EC2 instance, you can select the placement group on on Step 3: Instance details
- Cluster is not available for t2.micro instances :(

### 35. Elastic Network Interfaces (ENI) with Hands On
- Logical component in a VPC that represents a virtual network card
- Looks like it's a virtual ethernet interface you can attach to a given instance
- An ENI can have the following attributes:
    - Primary private IPv4, one or more secondary IPv4s
    - Each ENI can have one Elastic IPv4 per private IPv4 or per public IPv4
    - You can have one or more security groups attached to your ENI 
    - You can have a MAC address attached to it too
    - You can create ENIs independently from your EC2 instances
    - ENIs are bound to specific AZs
- You can detach a ENI from a given instance and attach it to another instance

#### Hands on
- Created 2 EC2 instances 
- Went to the ENI tab under Network and Security. Neat, every instance EC2 instance has an ENI when launched that is listed here.
- Go ahead and create an ENI independent of an instance for now
- You can then right click this ENI and attach it to an instance. Remember, they have to be on the same AZ!
- Your instance will then have 2 private IPs. Neat!
- You can then detach this interface from this machine and attach it to another machine

### 36. [SAA-C02] EC2 Hibernate
- OooOooOoh new feature!
- You can stop and terminate instances
    - When you stop: The data on disk (EBS) is kept intact in the next start
    - When you terminate: EBS volumes (root) also set up to be destroyed are lost
- On instance start, the following happens:
    - On first start: OS boots & the EC2 user data script is run
    - On subsequent starts: The OS boots up
    - Then your application starts, caches get warmed up and that can take time!

#### Then comes EC2 hibernate!
- When you stop an instance with EC2 hibernate enabled, the in-memory (RAM) state is preserved
- The instance will then boot up much faster as the operating system has not been restarted
- Under the hood: The RAM state is written to a file on the root EBS volume
- The root EBS volume must be encrypted

### 37. [SAA-C02] EC2 Hibernate - Hands On
- Go into launching and EC2 instances and choose m5.large
- On Step 3: Instance details, the "Stop - Hibernate" behaviour is available. You can then enable hibernation as an additional stop behaviour
- Click next, make sure the volume is encrypted.
- Launch instance!
- Right click > instance state > stop - hibernate
- Once it stops, you can then right click it and start it again. Notice that the `uptime` of the instance will be the same as in-memory data is preserved
- Use cases:
    - Long running processes
    - Saving the RAM state
    - Services that take time to initialize

#### Good to know about EC2 Hibernate
- Doesn't support all instances, only larger ones apparently
- Instance RAM size must be less than 150gb (ram goes to disk after all)
- Bare metal instances are not supported
- AMIs supported: Amazon Linux 2, Linux AMIs and Windows
- RootVolume: Must be EBS, encrypted, not instance store, and large enough to support dumping the whole ram in it
- Only available for on-demand and reserved instances
- An instance cannot be hibernated for more than 60 days
- Exam will probably ask about just what can you use to save the ram state and hibernate the instance

### 38. EC2 for the Solution Architect
- EC2 instances are billed by the second, t2.micro is free tier
- To connect to an instance, on Linux / Mac we use SSH, on Windows we use Putty
- SSH runs  on port 22. It is recommended to lock down the security group to your IP so that you can only SSH into your instances from your IP
- If you get time out issues, those are considered security groups issues as far as the exam goes
- Permission issues on the SSH key: Run “chmod 0400”
- Security Groups can have IPs as rules but they can also reference other Security Groups instead of IP ranges (very popular exam question)
- Know the difference between Private IP, Public I^and Elastic IP
- You can customize an EC2 instance at boot time using EC2 User Data

### Part 2:
- Know the 4 EC2 lanch modes:
    - On demand, billed by the minute
    - Reserved (you buy an instance for a year or more)
    - Spot instances: Instances that you can bid to spawn them, you lose them if someone bids higher than you
    - Dedicated hosts: Gives access to an entire host, usually when you need access to sockets (?) or have licensing issues 
- Know the basic instances types: R, C, M, I, G, T2/T3
- You can create AMIs to pre-install software for your ec2 instances, resulting in faster boot times (you don't have to install stuff at startup)
- AMIs can be copied accross regions and accounts. AMIs by default are region scoped

- EC2 instances can be started in placement groups
    - Cluster: For high performance computing
    - Spread: Spread over AZs for high availability
    - Partition: Not mentioned? :(

### Quiz!
6/7
- Remember than an AMI created for a region can only be seen in that region!