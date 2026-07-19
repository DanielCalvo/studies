Notes from Dani: WAF seems cool to explore!

### 250. Event Processing in AWS

#### Using Lambda, SNS & SQS
- There are events on an SQS queue
- Lambda polls the SQS queue to pull messages
- If it fails to process the message, it'll put them back on the queue
- We set up a DLQ on SQS in case some messages can't be processed after many retries (say, 5)
- You can also use SQS FIFO

#### With SNS
- SNS can also asynchronously push messages into Lambda
- If Lambda can't process the message, it will retry internally (say, 3 times)
- If the message can't be processed successfuly, it can be discarded or sent into a SQS DLQ queue for later processing

#### Fan out patterns: Deliver to multiple SQS queues
Option 1: You have an app with the SDK installed and you would like to deliver these messages to 3 SQS queues
    - You can just write your app to send messages to these 3 queues
    - Not very reliable: If the app has an error after it sends to the second queue out of 3, the third queue won't get it's message
    - The content of each queue would be different :(
- Option 2: Use the fan out pattern! Combine your SQS queue with an SNS topic.
    - SDK > SNS > All your SQS queues
    - Your SQS queues are going to be subscribers of your SNS topic
    - So your SNS fans out that message into all your SQS queues

#### S3 events
- We can react to S3 events! There are many events we can react to, such as: 
    - S3:ObjectCreated, S3:ObjectRemoved, S3:ObjectRestore, S3:Replication...
- You can also have a filter to react to only certain object names in particular (such as *.jpg)
- Use case: Generate thumbnails of images uploaded to S3!
- You can react to these events in several ways!
    - You can send a notification to an SNS topic
    - You can send a message into a SQS queue
    - Or you can asynchronously invoke a Lambda function!
    - Remember that SNS can fan out to multiple SQS queues
    - If you have an SQS queue, you can create a Lambda function or an EC2 instance to read from that SQS queue (if the processing fails, the message can still be in the queue!)
    - Or you can just asynchronously call a Lambda function from your S3 upload
    - And lambda can push to a SQS DLQ if it fails to process the message
- You can create as many “S3 events” as desired
- S3 event notifications are typically delivered events in seconds but can sometimes take a minute or longer
- If two writes are made to a single non-versioned object at the same time, it is possible that only a single event notification will be sent
- If you want to ensure that an event notification is sent for every successful write,

### 251. Caching strategies in AWS
- Client > CloudFront > API Gateway > App logic / EC2 Lambda > Redis/memcached/DAX && RDS (dynamic content route)
    - Computational cost increases as you go further to the right on the above line
- Client > CloudFront (edge) > S3 (static content route)
- You can enable caching TTL on CloudFront to make sure your content is not cached for too long and goes outdated
- You can cache in cloudfront, api gateway, app logic AND redis layers. There's no right or wrong way to choose were to cache, you need to see how your app is and decide where it best makes sense to cache things

### 252. Blocking an IP address in AWS
- Suppose one of you clients is a bad actor and you want to block their IP address
- Imagine you have the following:
    - EC2 instance
    - In a SG, inside a VPC
- The first line of defence for this is to use a network ACL inside a VPC. You can create a DENY rule for this IP address.
- Remember that the SG does not have DENY rules, only ALLOW rules
- If only some IPs can access your instance, then it's a good idea to define those in your SG 
    - If your app is global, there is no way of knowing this though
- You also can run a firewall software on your EC2 to block access there

#### Blocking an IP address - with an ALB
- Client > ALB > EC2 instance (private IP)
- Both the ALB and EC2 instances have their own SGs
- On the EC2 SG, we only allow the ALB SG, so we're safe there
- Our line of defense is going to be at network ACL level

#### In case it's a NLB
- NLBs do not do connection termination
- NLBs do not have a SG
- In this case, the end EC2 instance sees the IP of the original client
- Welp, there was no conclusion here from Stephane. I think you can use NACLs on the NLB?

#### But back to ALBs
- Something we can do to deny an IP is to install WAF, aka Web Application Firewall. WAF is an additional service and it'll cost money though.
- WAF accepts rules to deny too many requests, so we have more power on the security on the ALB
- WAF is not a service in between your client and your ALB, it's a service we have installed on the ALB, you can define a bunch of rules in there.

#### Blocking an IP address - ALB, CloudFront, WAF
- If you use CloudFront in front of an ALB, CloudFront sits outside our VPC.
- Our ALB needs to allow all the CloudFront public IPs coming from the edge locations to reach it
- In this case, the NACL is not helpful, as it cannot help is block the client IP address to cloudfront
- If you want to block an IP to cloudfront:
    - You can block an entire country with cloudfront geo restriction
    - Or if it's a specific IP, you can use WAF to do some IP address filtering

### 253. High Performance Computing (HPC) on AWS
- The cloud is the perfect place to perform HPC
- You can create a very high number of resources in no time
- You can speed up time to results by adding more resources
- You can pay only for the systems you have used
- You can perform genomics, computational chemistry, financial risk modeling, weather prediction, machine learning, deep learning, autonomous driving
- Which services help perform HPC?

#### Data Management & Transfer
- AWS Direct Connect:
    - Move GB/s of data to the cloud, over a private secure network
- Snowball & Snowmobile
    - Move PB of data to the cloud
- AWS DataSync
    - Move large amount of data between on-premise and S3, EFS, FSx for Windows

#### Compute and Networking
- EC2 Instances:
    - Can be CPU optimized or GPU optimized
    - Spot Instances / Spot Fleets can be leveraged for cost savings + Auto Scaling
- EC2 Placement Groups: Cluster for good network performance
    - Same RACK
    - Same AZ

#### Further computer and Networking Improvements
- EC2 Enhanced Networking (also calledSR-IOV)
    - Has higher bandwidth, higher PPS (packet per second), lower latency
    - To enable this:
    - Option 1: Elastic Network Adapter (ENA) up to 100 Gbps
    - Option 2: Intel 82599 VF up to 10 Gbps – LEGACY
- Elastic Fabric Adapter (EFA)
    - Improved ENA for HPC, only works for Linux
    - Great for inter-node communications, tightly coupled workloads
    - Leverages Message Passing Interface (MPI) standard
    - This standard bypasses the underlying Linux OS to provide low-latency, reliable transport
- Remember what ENA means!

#### Storage
- Instance-attached storage:
    - EBS: scale up to 64000 IOPS with io1 Provisioned IOPS
    - Instance Store: scales to millions of IOPS, linked to EC2 instance, low latency
- Network storage:
    - Amazon S3: large blob, not a file system
    - Amazon EFS: Scalable IOPS based on the total size of the filesystem, or you can use provisioned IOPS
- Amazon FSx for Lustre:
    - HPC optimized distributed file system, millions of IOPS
    - Backed by S3

#### Automation and Orchestration
- AWS Batch
    - AWS Batch supports multi-node parallel jobs, which enables you to run single jobs that span multiple EC2 instances.
    - Easily schedule jobs and launch EC2 instances accordingly
- AWS ParallelCluster
    - Open source cluster management tool to deploy HPC on AWS
    - Configurable with text files
    - Automates the creation of VPC, Subnet, cluster type and instance types
- HPC is a concept is popular on the exam!

### 254. EC2 Instance High Availability

#### Creating a highly available EC2 instance
- EC2 instances are not highly available by default
- What you can do is
    - Client > EIP > EC2 instance
- You can have a second EC2 instance on standby
- If the instance fails, a cloudwatch alarm is triggered and a lambda function is launched that attached the EIP to another EC2 instance

#### Creating a highly available EC2 instance with an ASG
- ASG in 2 AZs
- ASG with
    - 1 min
    - 1 max
    - 1 desired
    - >= 2 AZ
- The EC2 instance can have an elastic IP attached to it just as in the previous example
- If the EC2 instance goes down, the ASG will create another one
- You can have some user-data script that will attach to the EIP when an instance is brought up on this ASG
- Using cloudwatch and lambda is not necessary (like above) when you do it like this
- Make sure that the EC2 instance has a role that allows it to make API calls to attach to an elastic IP

#### Creating a highly available EC2 instance with an ASG + EBS
- Maybe you have a DB
    - Maybe you're dumb enough to run a DB on EC2
- Its a EC2 instance part of an ASG with an EBS volume
- ASG has a thing called lifecycle hook, and on termination of an instance that has the EBS volume mounted, to take that EBS volume and create a snapshot from it
- You can then have a lifecycle hook on the launch event, which will create an EBS volume out of the EBS snapshot, and mount it to the instance

### 255. Bastion Host High Availability
- In this case, each AZ has a private subnet and a public subnet
- With a Bastion host you want to have a machine on the public subnet, so you can ssh into the machines that are on the private subnet
- You can create a NLB balancing 2 bastion hosts if you want
- You can also run a single bastion host across 2 AZs with an ASG with 1:1:1
- If you only have 1 bastion host, you can use an elastic IP with an ec2 user data script to access it
- If 2 bastion hosts, use an Network Load Balancer (layer 4) deployed in multiple AZ
- If NLB, the bastion hosts can live in the private subnet directly
- Note: Can’t use ALB as the ALB is layer 7 (HTTP protocol)

#### Quiz
- Pending