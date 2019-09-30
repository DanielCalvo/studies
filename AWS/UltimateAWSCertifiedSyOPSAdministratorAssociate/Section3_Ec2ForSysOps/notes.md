Follow ups
- What is an EBS backed instance?
- Make sure you remember the difference between an EC2 dedicated host and a EC2 dedicated instance 

### 5. Launching an EC2 instance
- Chapter is all hands on
- Created an instance with a new SG and an existing keypair

### 6. Changing EC2 instance type
- Only works for EBS backed instances
- Stop the instance. You can then right click on the instance and instance settings > change instance type.
- You can log in into the instance and see that your files are there and that you have different resources (more memory for instance, if you resized to a bigger instance)

### 7. EC2 placement groups
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
- When creating an EC2 instance, you can select the placement group
- Cluster is not available for t2.micro instances :(

### 8. EC2 Shutdown behaviour and termination protection

#### Shutdown behaviour
- How should the instance react when shutdown is done using the OS?
- Stopped: Default
- But you can also set it to `terminate`
- Not applicable when you use the UI or the API. This is only from the OS!
- CLI Attribute: `InstanceInitiatedShutdownBehaviour`

#### Termination protection
- To protect against accidental termination in the AWS Console or CLI
- Exam tip: You have an instance where shutdown behaviour is set to terminate and termination protection is enabled. What happens if you shutdown the instance from the OS?
    - It still terminates.

### Hands on
- You can enable termination protection through the UI
- Once you enable it, you can't terminate from the console, you have to disable termination first

### 9. Troubleshooting EC2 launch issues
- Exam tip: Popular on the exam!

- InstanceLimitExceeded error
    - This means you've reached your maximum number of instances per region.
    - Resolution: Either launch the instance in a different region or request AWS to raise the limit for you.
    - Default number of instances in each region: 20
    - On the EC2 console, on the left side there's a limit section. You can see your limits there.

InsufficientInstanceCapacity
    - AWS does not have that much on demand capacity in that particular AZ to launch the instance(s) you want.
    - Resolution: Wait a few minutes before trying again.
    - Or maybe try to break down the requests (ask for 1 instance 5 times instead of 5 instances at once)
    - If urgent, maybe create an instance of a smaller type and try to increase it's size later.

If instances terminates immediately (goes from pending to terminated). Popular question at the exam! Possible reasons:
    - You reached your EBS volume limit
    - An EBS snapshot is corrupt
    - The root EBS volume is encrypted and you do not have permissions to access the key for decryption
    - The instance store-backed AMI that you used to launch the instance is missing a required part

#### Hands on error tips
To find the exact reason, there's a description tab on EC2 instances with a "state transition reason" and "State transition reason message"
You can also check these when you click over an EC2 instance and it shows up on screen on the bottom part.

### 10. Troubleshooting EC2 SSH issues
- Make sure that the private key as 400 permissions on your machine otherwise you'll get the "unprotected private key file" error
- Make sure the user name for the OS is given correctly when logging in via SSH, else you'll get the "Host key not found" error

- If you get a timeout:
    - Usually a security group not configured correctly (por 22 not reachable) or
    - CPU usage is so high that you can't log in

### 11. EC2 Instance launch types

#### EC2 on demand:
- Pay for what you use
- Highest cost, but no upfront payment
- No long term commitment.
- Recommended for short term, uninterrupted workloads in which you can't predict what's gonna happen in the future (if the instance is gonna stay there forever or not)

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

### 12. EC2 launch modes hands on:
- You can search for reserved instances on Amazon's EC2 panel on the left
- Same for allocating hosts
- Same for scheduled instances
- Spot instances and tenancy options can be found when creating a normal EC2 instance as well
- For spot instances your request can be persistent and you can establish a date range for it. Pretty cool
- Exam will ask which instance type is the right type for some usages (like a database)
 
### 13. EC2 instances deep dive
- R: Applications that need a lot of ram, such as in memory caches
- C: Apps that use CPU, compute/databases
- M: Balanced, think general/web app
- I: I/O focused (dbs)
- G: GPU. Video rendering / machine learning
- T2/T3 Burstable: Can burst CPU usage (up to a capacity)
- T2/T3 unlimited: Unlimited burst.
- Suggested by the author: https://ec2instances.info/

#### Burstable instances (T2/T3)
- Burst mean that the instance can burst CPU usage for extra processing, but only for a limited amount of time. You regain "burst credit" over time.
- In CloudWatch you can see your burst usage / credits
- On T2/T3 unlimited instances you have "unlimited burst", you can pay more for these and you never have performance penalties. It can be expensive however.

### 14. EC2 AMIs
- AWS comes with a bunch of base images such as Ubuntu, Fedora, Redhat, Windows etc
- These images can be customized at runtime using EC2 user data.

#### Why use custom AMI?
- Pre installed packages can come installed
- Faster boot time
- Control of maintenance and updates of AMIs over time
- AD integration out of the box
- Installing your app ahead of time
- Or you can use someone elses AMI that's optimized for a given purpose
- Remember: AMIs are build for a specific AWS region!

- You can create a custom AMI! With pre-installed packages and any other thing you might need.

#### Using Public AMIs
- You can leverasge AMIs from other people
- You can also pay for other people's AMI by the hour
    - These people have optimized the software
    - The machine is easy to run and configure
    - You "rent expertise" from the AMI creator
- AMIs can be found and published on the Amazon Marketplace

#### Warning
- Don't use an AMI you don't trust! 

#### AMI Storage
- When you create an AMI it resides in S3, but you won't see it in the console.
- By default all AMIs are private. 
- Make sure you remove AMIs you don't use as you'll be charged for them.

### 15. EC2 AMIs hands on:
- Launch an EC2 instance, do whatever you want on it (like install apache with a hello world through web)
- Right click on your running instance > image > create image
- Once created, you can go on the AMI menu and "Copy AMI" to a different region. A given AMI is locked to a given region.
- You an also modify the image permissions, de-register it or remove it.


### 16. Cross account AMI Copy
- Popular at the exam!
- You can share an AMI with another AWS account, this does not affect the ownership of the AMI
- If you copy an AMI that has been shared, you then become the owner of that AMI.
- To copy an AMI that was shared with you from another account, the owner of the source AMI must grant you read permission for the storage that backs the AMI
    - Either the associated EBS snapshot (for an Amazon EBS backed AMI)
    - Or the associated S3 bucket (for an instance store-backed AMI)

- You can't copy an encrypted AMI that was shared with you from another account. Instead, If the underlying snapshot and encryption key were shared with you, you can copy the snapshop and re-encrypt it with a key of your own. You own the copied snapshot and can register it as a new AMI

- You can't copy an AMI with an associated **billingProduct** code that was shared with you from another account. This includes Windows AMIs and AMIs from the AWS marketplace. To copy a shared AMI with a billingProduct code, launch an EC2 instance in your account using the shared AMI and then create an AMI from the instance. Popular at the exam!

- Reemphasizing: To copy an AMI with a billingProduct code, first launch an EC2 instance from that AMI and then you make an AMI from that EC2 instance 

### 17. Elastic IPs
- When you stop and you start an EC2 instance, it changes its public IP.
- If you need to have a fixed public IP, you need an elastic IP.
- An Elastic IP is a public IPv4 IP you own as long as you don't delete it 
- You can only attach it to one instance at a time
- You can remap it across instances
- You don't pay for the elastic IP if it's attached to a server, but you do pay it it's not attached.

#### Elastic IP, part deux
- Exam tip:  With an elastic IP address, you can mask the failure of an instance or software by rapidly remapping the address to another instance in your account
- You can only have 5 elastic IPs by default but you can ask AWS for an increase.
- Overall, author recommends trying to avoid using elastic IPs.
    - Think of the alternatives
    - Maybe use a random public IP (as assigned to EC2 instances normally) and register a DNS name to it
    - Or use a Load Balancer with a static hostname

#### Hands on
- You can go on the Elastic IP tab and easily create an elastic IP. You can then, right click on an elastic IP and associate it with an instance. This instance will have the elastic IP.
- You can disassociate the IP from one machine and associate it to another one in the elastic IP menu.

### 18. CloudWatch metrics for EC2 
- Exam tip: Important for the exam to know how CloudWatch is linked to EC2
- Basic monitoring: Default, metrics that are collected at 5 minutes intervals. These are free.
- Detailed monitoring: Paid, metrics that are collected at 1 minute intervals
- Metrics include: CPU, Network, Disk and Status Check metrics

#### Custom metrics
- Yours to push!
- Basic resolution: 1 minute resolution 
- You can also have high resolution, all the way up to 1 second!
- Custom metrics you could have: RAM, application level metrics
- If you do enable custom metrics, make sure the IAM permissions on the EC2 instance roles are correct! 

#### EC2 included metrics
- CPU: Utilization + credit usage / balance for burst instances
- Network: in/out
- Status check
    - Instance status, checks the EC2 VM.
    - System status: Checks the underlying hardware.
- Disk: Read/Write for Ops / Bytes (only for instance store backed instances)
- Ram: Not included in the AWS EC2 metrics! AWS doesn't push it into CloudWatch :o
- Common exam question: Can you get RAM from cloudwatch? No, you have to push it yourself!

#### Hands on
- There's a bunch of stuff available on the EC2 instance monitoring tab out of the box! Neat
- You can also enable detailed monitoring on the EC2 dashboard (metrics every 1 minute). It incurs costs!

### 19. EC2 custom metrics
- Sample custom metrics for EC2
    - RAM usage
    - Swap usage
    - Any custom metric from your application (like requests per second)

#### Hands on to enable ram
- Scripts exist from AWS themselves
- We'll push RAM as a custom metric
- Logged in into the instance, installed some dependencies, downloanded and unzipped a scripts folder
- IAM > Roles > Create role > EC2 instance role > CloudWatchFullAccess
- On EC2, right click on instance > Instance settings > Attach IAM role


###20. CloudWatch Logs for EC2
- By default, no logs from your EC2 instance will go to CloudWatch
- You need to run a CloudWatch agent on EC2 to push the logfiles you want to CloudWatch

#### Hands on
- Create and apply to your EC2 instance an IAM role that has CloudWatchFullAccess
- On your EC2 instance
```bash
sudo yum update -y
sudo yum install -y awslogs
sudo vim /etc/awslogs/awslogs.conf
sudo systemctl start awslogsd
sudo systemctl enable awslogsd.service
```
Rerefence: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/QuickStartEC2Instance.html

- You need to make sure you have the correct permissions to push logs to cloudwatch!
- The cloudwatch log agent can be set up on premise too

You install awslogsd and you can then check log metrics in cloudwatch. Cool!

### Quiz
1. To change the EC2 instance type, you need to stop the instance first, then change type, then start it again
2. To make sure your instances have the maximum network performance, launch them in cluster mode
3. Even if shutdown protection is enabled, if you shut down an instance from the OS using shutdown and the shutdown behaviour is set to terminate, the instance will terminate!
4. The InsufficientInstanceCapacity error means AWS does not have enough on-demand capacity regarding your particular AZ
5. On start up, if an instance goes from pending to terminating, your issue is NOT the per region limit on your account (could be EBS volume limit, corrupted EBS snapshot or an encrypted EBS volume that you don't have the key for)
6. If you're going to run an instance for a year, consider running a reserved instance, it will be cheaper!
7. If you're running out of CPU credits on a t2.small, you cannot purchase CPU credits manually. You can turn to t2 unlimited, upgrade to t2.medium or higher, or upgrade to a non t* instance.
8. AMIs are by default region locked!
9. If an AMI has a billingProduct code, you can only clone it by launching an EC2 instance and then creating an AMI from that instance.
10. RAM needs to be passed as a custom metric for cloudwatch