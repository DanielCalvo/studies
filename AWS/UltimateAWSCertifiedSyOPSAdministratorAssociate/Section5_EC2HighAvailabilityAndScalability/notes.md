- Chapter status: Went all the way to lecture 36 and then went for the Cloudformation lecture
    - Resume on the lecture 36 with the hands on using cloudformation perhaps
    - Lecture 42 confused me a bit, maybe re-watch it

### 34. Section introduction
- Load balancers
    - Troubleshooting!
    - Advanced options and logging
    - CloudWatch Integrations
- Autoscaling groups
    - Troubleshooting!
    - Advanced options and logging
    - CloudWatch Integrations

### 35. What is high availability and Scalability?
- Scalability means that an application/system can handle greater loads by adapting
- There are two kinds of scalability
    - Vertical scalability  
    - Horizontal scalability (also called Elasticity)
- Scalability is linked but different to High Availability

#### Vertical scalability
- Vertical scalability means increasing the size of the instance
- Like increasing from t2.micro to t2.large
- Common for some non distributed systems, like some databases
- RDS and ElastiCache are services that can scale vertically
- There's usually a limit on how much you can vertically scale (hardware)

#### Horizontal scalability
- Horizontal scalability means increasing the number of instances/systems in your application
- Horizontal scaling implies distributed systems
- Very common for web applications
- It's easy to horizontally scale thanks to cloud offerings such as Amazon EC2

#### High Availability
- High Availability usually goes hand in hand with horizontal scaling
- High Availability means running your application / system in at least 2 data centers (or 2 different availability zones)
- The goal of HA is to survive a data center loss
- HA can be passive, RDS Multi AZ for example
- HA can be active, EC2 horizontal scaling for example

#### High Availability & scalability for EC2
- Vertical scaling: Increase instance size
- Horizontal scaling: Increase the number of isntances (scale out/in). This can be used with
    - Auto scaling groups
    - Load Balancers
- High Availability: Run instances for the same application across multi AZ
    - Auto Scaling group multi AZ
    - Load Balancer multi AZ

Vertical: Means increasing the size of your instance. RDS and Elasticache can scale vertically well, although there are limits to this.
Horizontal: Means increasing the number of instances/systems for your application

Horizontal scaling implies distributed systems.
Not every application can be a distributed system.
Easy to horizontally scale on EC2!

High Availability: Means running your app in at least 2 datacenters or 2 AZs in AWS.
Goal: Survive datacenter loss.

An Auto scaling group can be multi-az
A load balancer can be multi-az


### 36. Load Balancer overvieww
- Load balancing are servers that forward internet traffic to multiple EC2 instances downstream
- Users don't connect directly to an instance, they connect to a load balancer.

#### Why use a load balancer
- Spreads load across multiple downstream instances 
- Expose a single point of access (DNS) to your app
- Seamlessly handle failure of downstream instances
- Does helthcheck on the instances and redirects traffic to instances that are working
- Load balancers can provide HTTPs termination. Encryption/SSL is between the client and the ELB
- Can enforce stickiness with cookies so that the user talks to the same instance over time 
- Can provide high availability across zones
- Can separate public traffic from private traffic (?)

####  Why use an EC2 loadbalancer
- And ELB (EC2 load balancer) is a managed load balancer
    - AWS guarantees that it will be working
    - AWS takes care of updates, maintenance, high availability
    - AWS provides only a few configuration knobs
- It costs less to set up your own load balancer but it will be a lot more effort on your end
- ELB is integrated with many AWS offerings/services

#### Types of load balancers on AWS
- AWS has 3 kinds of load balancers
- Classic load balancer, or v1, from 2009
- Application load balancer, v2 - 2016
- Network load balancer, v2, 2017
- Overall it's recommended to use the newer, v2 load balancers as they provide more features
- Application load balancer and network load balancer is what will be asked mainly in the exam
- You can set up internal (internal) or external (public) ELBs

#### Healthchecks:
- Health checks are crucial for load balancers
- They enable the load balancer to know which instances are healthy so that it can forward traffic to them
- The health check is done on a port and a route (/health is common)
- If the instance responds 200 (OK) it is considered healthy, if it responds with anything else it is considered unhealthy

#### Application load balancer (v2)
- Application load balancers (layer 7), these allow:
    - Load balancing to multiple HTTP applications across machines (target groups)
    - Load balancing to multiple HTTP applications on the same machine (ex: containers)
    - Common question on the exam! We need something to load balance across the same application running on the same machine (?). Application load balancer does it!
    - Load balance based on route in URL
    - Load balance based on hostname in URL
- Load balancers are a great fit for microservices & container-based applications (such as Docker & Amazon ECS)
- There is also a port mapping feature to redirect you to any dynamic port in the backend
- In comparison, you would need to create one Classic Load Balancer per application before. That was very expensive and ineficient! 

#### Diagram on application load balancer for HTTP
- A load balancer sends traffic to a target group, with a healthcheck
- A load balancer will send the http traffic based on the route (like /user)
- You can set up many target groups on the same application load balancer
- You can have as many target groups as you want behind your ELB

#### Application load balancer v2 good to know 
- Stickiness can be set at the target group level. 
    - Same request goes to the same instance
    - Stickiness is directly generated by the ALB (not the application, cookie at the ALB level)
- ALB supports HTTP/HTTPS & Websockets protocols
- The application servers don't see the IP of the client directly (common exam question!)
    - The true IP of the client is inserted in a header named `X-Forwarded-For`
    - The port and protocol are also available in `X-Forwarded-Port` and `X-Forwarded-Proto` respectively
    - ELB does the connection termination for the clients. Your EC2 sees the private IP of your load balancer

#### Network Load Balancer (v2)
- To forward TCP traffic to your instances
- Can handle millions of requests per second, very high performance
- Support static IP or elastic IP
- Less latency, about 100ms (vs 400ms for ALB)
- Network Load Balancers should be used for extreme performance and should not be the default load balancer you choose
- Overall, the creation process is the same as the applocation load balancer

#### Diagram
- It's about the same, except everything is TCP based
- You get routed to a target group based on "TCP + rules". I imagine it's a TCP port and some other rules

#### Load Balancer Good to Know
- Classic load balancers are deprecated
- Use the Application load balancer for HTTP/HTTPS and Websocket
- Use the Network Load Balancer for TCP
- CLB and ALB support SSL certificates and provide SSL termination
- All load balancers have health check capability
- ALB can route based on hostname/path
- ALB is a great fit with ECS (Docker)
- Any load balancer (CLB, ALB, NLB) has a static host name. Do not try to resolve and use the underlying IP. You get the URL, use the URL. Don't use the underlying IP. (Popular on the exam)
- LBs can scale, but not instantly, contact AWS for a "warm-up"
- NLB see directly the client IP, no X-Forwarded-For header (probably an HTTP only functionality)
- 4xx are client induced errors
- 5xx are application induced errors
    - Load balancer errors 503 means at capacity or no registered target
- If the LB cant connect to your application, check your security groups!


### 37. Load Balancer Hands On Using SSM
- Launched a single instance from the EC2 UI
- Author does this through SSM, so make sure the machines are on a SG that allows SSM access
- Right clicked on the instance and chose "Launch more like this" until we had a total of 3 instances
- Create a target group for HTTP traffic
- Once the target group is created, on the dasboard, click on the "Targets" tab, edit and select the instances you just created manually
- Go on load balancer and create an application load balancer, then create a new security group for port 80, and then configure routing to your target group
- Until your deploy your web app, the health checks on the LB will be "unhealthy" as there's nothing listening on port 80 on your server 
- Author then installs apache through SSM
- Reaching the LB we get different instances every time. Neat!


### 38. Load Balancer Stickiness
- It's possible to implement stickiness so that the same client is always redirected to the same instance behind a load balancer
- This works for classic load balancers and application load balancers
- The cookie used for stickiness has an expiration date you control
- Use case: Make sure the user doesn't lose his session data
- Enabling stickiness may bring imbalance to the load over the backend EC2 instances
- Possible exam question: Once EC2 instance is experiencing 80% load, the other has 20% and they're both behind a load balancer, what's the reason?
- Reason: Stickiness is enabled and a lot of clients are "stickied" to an instance! 

#### Hands on
- On any given target group, stickiness is a setting. You can edit attribute and enable stickiness
- Remember that stickiness can bring imbalance as users will reach the same instance over and over again
 
### 39. ELB for SysOps
- Exam to recap on exam sysops topics
- Application load balancer
    - Layer 7, HTTP, HTTPS, websocket
    - URL Based routing (hostname or path)
    - Does not support fixed IP, but has fixed DNS name
    - Provides SSL termination
- Network Load Balancer:
    - Layer 4 (TCP, lower level)
    - No pre-warming needed
    - 1 static IP per subnet (no DNS name)
    - You cannot do SSL termination on a network LB (you need to do it at the app level)
- Possible exam question: "How do you give an application load balancer a fixed IP?"
- Hacky: Chain a NLB and an ALB together. You place a net load balancer in front of your app load balancer

#### Load Balancer Pre Warming
- ELB will gradually scale to traffic
- ELB may fail in case of sudden spike of traffic (like 10x the traffic all of sudden)
- IF you expect high traffic, open a support ticket with AWS to pre warm your ELB. Things you need to tell them:
    - Duration of traffic
    - Expected requests per second
    - Size of these requests (in KB)

#### LB Error codes:
- Successful: 200
- Unsuccessful at client side: 4xx
    - Error 400: Bad request
    - Error 401: Unauthorized 
    - Error 403: Forbidden
    - Error 460: Client closed connection
    - Error 463: X-Forwarded for header with > 30 IPs (similar to malformed requests)
- Unsuccessful at server side: 5xx
    - Error 500: Internal server error, some error on the ELB itself, rather rare
    - Error 502: Bad Gateway
    - Error 503: Service Unavailable
    - Error 504: Gateway timeout, probably an issue with the server 
    - Error 561: Unauthorized

#### Supporting SSL for old browsers
- Possible question: "How do we support legacy browsers that have use old TLS (such as TLS 1.0)?"
- Answer: Change the policy to allow for weaker ciphers (ex: DES-CBC3-SHA for TLS 1.0), old, not something AWS supports out of the box
- Only a small part of the internet still uses old browsers however
- Elastic load balancing provides security policies for application load balancers. You need to choose the oldest one, TLS-1-0-2015-04


#### Load Balancer common troubleshooting 
- Check security groups if the LB isn't working
- Check health checks, maybe some instances are unhealthy
- Sticky sessions may bring "imbalance" on the LB side
- For multi-az, make sure cross zone balancing is enabled
- Use an internal load balancer for private applications that don't need public access
- If you have a production load balancer, enabled deletion protection to prevent against accidental deletes. Deletion protection can be enabled by editing the LB attributes on the UI


### 40. Metrics, Logging and tracing for ELBs
- All load balancer metrics are directly pushed to CloudWatch metrics.
- Some examples of metrics you have avaialble:

- BackendConnectionErrors
- HealthyHostCount / Unhealthy host counts
- HTTPCode_Backend_2xx: Successful request
- HTTPCode_Backend_3xx: Redirected request
- HTTPCode_Backend_4xx: Client Error codes
- HTTPCode_Backend_5xx: Server error codes generated by the load balancer
- Latency
- RequestCount
- SurgeQueueLenght: The total number of requests (HTTP Listener) or connections (TCP listener) that are pending routing to a healthy instance. Helps to scale out ASG. Max value is 1024
- SpilloverCount: The total number of requests that were rejected because the surge queue is full

#### Hands on
- The monitoring tab on the load balancer has some interesting metrics right out of the box
- They're all pushed into CloudWatch directly, there's a link to "view in Cloudwatch"
- Target groups also have a monitoring tab. Neat!

### Load Balancer access logs
- Access logs from load balancers can be stored in S3 and contain
    - Time
    - Client IP address
    - Latencies
    - Request paths
    - Server response
    - Trace ID
- The logging service is free, you only pay for the S3 storage
- Helpful for compliance
- Helpful for keeping access data even after the ELB or EC2 instances are terminated
- Access logs are automatically encrypted

#### Hands on
- On the load balancer's attributes you can enable the acccess logs, and then you specify an S3 location

### Application Load Balancer Request Tracing
- Request tracing - Each HTTP request has an added custom header: "X-Amzn-Trace-id"
- This is useful in logs/distributed tracing platform to track a single request
- Application load balancer is not yet integrated with X-Ray (?)

### Load Balancer troubleshooting using metrics
- HTTP 400: BAD_REQUEST: The client send a malformed request that does not meet HTTP specifications
- HTTP 503: Service Unavailable: Ensure that you have healthy instances in every AZ that your load balancer is configured to reach. Look for the HealthyHostCount metric in CloudWatch
- HTTP 504: Gateway Timeout: Check if keep-alive settings on your EC2 instances are enabled and make sure that the keep-alive timtout is greater than the idle timeout settings for the load balancer
- There's a page with useful information on troubleshooting here: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/ts-elb-error-message.html

### 41. Auto Scaling groups overview
- In real life, the load on your websites and applications will change
- In the cloud, you can create and get rid of servers very quickly
- The goal of an Auto Scaling group (ASG) is to:
    - Scale out (add EC2 instances) to match an increased load
    - Scale in (remove EC2 instances) to match a decreased load
    - Ensure we have a minimum and maximum number of machines running 
    - Automatically register new instances to a load balancer

#### Diagram
- Things to know for the exam: Minimum size, Maximum size, scaling out, scaling in, Actual Size / Desired Capacity

#### ASGs have the following attributes
- A launch configuration
    - AMI + Instance type
    - EC2 user data
    - EBS volumes
    - Security Groups
    - SSH key pairs
- Min size, Max Size, Initial Capacity
- Network + Subnets information
- Load balancer information or Target group information
- Scaling policies

#### Auto Scaling Alarms
- It is possible to scale an AGS based on CloudWatch alarms
- When a given metric in cloud watch reaches a certain level, an alarm condition is reached and the ASG can scale up.
- Metrics are computed for the overall ASG instances!
- Based on the alarm
    - We can create scale out policies
    - We can create scale in policies

#### Auto Scaling New Rules
- It is now possible to define "better" auto scaling rules that are directly managed by EC2:
    - Target Averate CPU Usage
    - Number of requests on the ELB per instance
    - Avg net in / net out
- These rules are easier to set up than the previous ones

#### Auto Scaling Custom Metric
- We can auto scale based on a custom metric (ex: number of connected users)
1. Send custom metric from application on EC2 to CloudWatch (PutMetric API)
2. Create CloudWatch alarm to react to low/high values
3. Use the CloudWatch alarm as the scaling policy for ASG

#### ASG Brain dump
- Scaling policies can be on CPU, Network... and can even be custom metrics or based on a schedule (if you know your visitor/workload patterns)
- ASGs use launch configurations and you update an ASG by providing a new launch configuration
- IAM roles attached to an ASG will get assigned to EC2 instances
- ASGs are free. You pay for the underlying resources being launched
- Having instances under an ASG means that if they get terminated for whatever reason, the ASG will restart them. Extra safety!
- ASGs can terminate instances market as unhealthy by an LB (and hence replace them)

### 42. Auto Scaling Groups Hands on (with ELB health checks)
- Create ASG > Create launch configuration
- User data on instances is a script to install apache
- Then create the ASG itself, select 3 subnets
- Advanced detail > Enable Load balancers, select a target group
- You can then define a scaling policy or keep the group at its initial size 
- It can scale based on cpu usage, or you can set the scaling based on cloudwatch alarms
- Popular exam question: The ASG is healthy but the ELB is not. Why? Well dunno but you have to go from an EC2 health check type to an ELB health check type

### 43. ASG Scaling Processes Hands On
- Behind the scenes, an ASG has scaling processes.
- Launch: Add a new EC2 instance to the group, increasing capacity
- Terminate: Removes an EC2 instance from the group, decreasing its capacity
- HealthCheck: Checks the health of the instances
- ReplaceUnhealthy: Terminates unhealthy instances and re-creates them
- AZRebalance: Balance the number of EC2 instances a cross an AZ
- AlarmNotification: Accept notifications from CloudWatch
- ScheduledActions: Performs scheduled actions that you create
- AddToLoadBalancer: Adds instances to the load balancer or target group
- We can suspend these processes! The first five are the most common

#### Hands on
- Edit ASG
- Suspend processes checkbox, you can choose which processes you want to suspend.
- By suspending the terminate process, it cannot terminate any instances, even if you lower the desired capacity
- You can also suspend ReplaceUnhealthy
- You can use the CLI to set instance health: `aws autoscaling set-instance-health --instance-id <ID> --health-status Unhealthy`

```shell script
#!/bin/bash
yum update -y
yum install -y httpd
systemctl start httpd.service
systemctl enable httpd.service
echo "Hello from host: $(hostname -f)" >> /var/www/html/index_old.html
```

##### Note on AZ rebalance
- When there is an imbalance on the AZ spread
- AZRebalance = launch new instance then terminate old instance
- If you suspend the Launch process
    - AZRebalance won't launch instances
    - AZRebalance won't terminate instances
- If you suspend the terminate process
    - It will still launch instances, the ASG can grow to up 10% of it's size (it's allowed during rebalanced)
    - In this case, the ASG will remain at the increased capacity because it can't terminate instances (popular exam question). Remember, 10%
    - All of the scaling processes can affect the ASG scaling processes

### 44. ASG for SysOps
- Exam tips!
- To make sure you have high availability, you need to have at least 2 instances running across 2 AZs (therefore 2 subnets) in your ASG (must configure multi AZ ASG)
- Health checks available: 
    - EC2 status checks (knowing if the VM/Hardware is OK, no app info)
    - ELB health checks (can check the app health)
- ASG will lauch a new instance after terminating an unhealthy one
- ASG will not reboot unhealthy hosts for you (there's no such thing as an ASG reboot)
- Two good to know CLIs:
    - set-instance-health
    - terminate-instance-in-auto-scaling-group

#### Troubleshooting ASG issues
- <number of instances> instances(s) are already running. Launching EC2 instance failed 
    - The Auto Scaling Group has reached the limit set by the DesiredCapacity parameter. Update your Auto Scaling group by providing a new value for the desired capacity
- Launching EC2 instances is failling
    - The security group may not exist. SG might've been deleted
    - The key pair does not exist. The key pair might've been deleted.
- If the ASG fails to launch an instance for over 24 hours, it will automatically suspend the process (administration suspension)


### 45. CloudWatch for ASG
- The following metrics are available for ASG, but they're opt-in:
    - GroupMinSize
    - GroupMaxSize
    - GroupDesiredCapacity
    - GroupInServiceInstances
    - GroupPendingInstances
    - GroupStandyInstances
    - GroupTerminatingInstances
    - GroupTotalInstances
- You should enable metric collection to see these metrics
- Metrics are collected every 1 minute
- You can also monitor the underlying EC2 instances 
- Basic monitoring with 5 minute granularity
- Detailed monitoring: 1 minute granularity

#### Hands on
- ASG > Monitoring > Enable Group Metrics Collection
- You also have EC2 instance metrics
- CloudWatch > Metrics > AutoScaling > GroupMetrics

### 46. Section Clean Up
- Cleaned up:
    - Auto scaling group
    - Launch configuration
    - Target Group
    - Load Balancer

### Quiz!
