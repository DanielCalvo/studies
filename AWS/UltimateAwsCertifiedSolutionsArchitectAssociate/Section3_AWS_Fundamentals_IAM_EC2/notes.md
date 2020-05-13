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
### Mid way quiz!
### 27. EC2 Instances Launch Types
### 28. [SAA-C02] Spot Instances & Spot Fleet
### 29. EC2 Instances Launch Types Hands On
### 30. EC2 Instance Types Deep Dive
### 31. EC2 AMIs
### 32. EC2 AMI Hands On
### 33. Cross Account AMI Copy
### 34. EC2 Placement Groups
### 35. Elastic Network Interfaces (ENI) with Hands On
### 36. [SAA-C02] EC2 Hibernate
### 37. [SAA-C02] EC2 Hibernate - Hands On
### 38. EC2 for the Solution Architect

### Quiz! 