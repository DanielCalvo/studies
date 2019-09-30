- Chapter status: Went all the way to lecture 36 and then went for the Cloudformation lecture
    - Resume on the lecture 36 with the hands on using cloudformation perhaps 

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
- Load balancers can provide HTTPs termination. Encryption/SSL is between the cliend and the ELB
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

```text
Right clicking on instance and > Launch more like this is pretty useful!
Load balancing > target group
Create target group, click on targets

When creating a load balancer you go through several steps that allow you to adjust targetting and security groups

Needless to say the healthchecks will be "unhealthy" at first as you need to install apache

You have to tweak a security group to only allow the LB to reach the instances and not the entire internet.
You can edit the inbound rules on the security group so that only the LB can reach the instances

```

### 38. Load Balancer Stickiness

```text
Stickiness: THe same client is always redirected to the same EC2 instance.
It works with cookies. Cookies are used for stickiness.
Why stickiness: To make sure the user doesn't lose his session data. If you talk to another server you lose your session information.
Enabling stickiness may bring some imbalance.
LB > Target groups > Enable Stickiness

```

### 39. ELB for SysOps

```text
Things you need to know for the exam:
Application load balancer: Layer 7, HTTP; HTTPS, websocket
URL Based routing
Does not support fixed IP, but has fixed DNS
Provides SSL termination

Network Load Balancer:
Layer 4 (TCP)
No pre-warming needed
1 static IP per subnet (no DNS name)
You cannot do SSL termination on a network LB (you need to do it at the app level)

Possible exam question: "How do you give an application load balancer a fixex IP?"
Hacky: Chain a NLB and an ALB together. You place a net load balancewr in front of your app load balancer

Load Balancer Pre Warming
ELB will gradually scale to traffic
ELB may fal in case of suddeen spike of traffic

IF you expect high traffic, open a support ticket with AWS to pre warm your ELB

LB Error codes:
Successful: 200
Unsuccessful at client side: 4xx
Unsuccessful at server side: 5xx

Supporting SSL for old browsers:
Change the ELB policy to allow fr weaker ciphers
Only a small part of the internet still uses that though

Common LB troubleshooting: 
Check security groups
Check health checks
Sticky sessions may bring "imbalance" on the LB side
For multi-az, make sure cross zone balancing is enabled
If you have have an internal LB, remember it's for private applications
If you have a prod LB, enable deletion protection! 
```

### 40. Metrics, Logging and tracing for ELBs

```text
All LB metrics are directly pushed to CloudWatch metrics
Stephane mentioned a bunch of metrics, kinda too specific.

There are LB access logs!
You only pay for the S3 storage. Helpful for compliance

```

### 41. Auto Scaling groups overview

```text
An autoscaling group can scale up or down to match a given load
And we can ensure a minimum and maximum number of machines running on an Auto Scaling Group
You can have:
Minimum size
Desired Capacity
Maximum size
Scale out = Add instances
Scale in = Remove instances

ASGs have the following attributes:
AMI + instanece type
EC2 user data
EBS volumes
Security groups
SSH keypair
minsize / maxsize / initial capacity
Network + subnets info
Load balancer + target group settings
Scaling policies (what triggers scale out or scale in)

It's possible to scale an ASG based on CloudWatch alarms
An alarm can be anything, such as average CPU usage

There are new features for autoscaling, such as rules that will scale up your ASG based on:
- Target averate CPU usage
- Number of requests on the ELB instance
- Net in, net out

You can also scale out / in based on a custom metric
ASGs are free, you only pay for the EC2 instances
Having instances under ASG means that if they get terminated the ASG will restart them or spawn them again
ASK can terminate instances marked as unhealthy by a LB

```

### 42. Auto Scaling Groups Hands on (with ELB health checks)

```text

EC2 > Autoscaling group > Create autoscaling group
```

### 43. ASG Scaling Processes Hands On

```text


```

### 44. ASG for SysOps

```text
To make sure you have high availabiliry, make you have at least 2 instances running across your AZs
Health checks are available:
EC2 status checks
ELB health checks
ASG will launch a new instance after terminating an unhealthy one
ASG will not reboot unhealthy hosts
Good to know CLI:
Set instance health
temrinate-instance-in-auto-scaling-group


Troubleshooting ASGs:
You can't launch any further instances if you're already at the desiredcapacity. If you want to add more, change desired capacity
EC2 launch can fail if SG does not exist
Key part does not exist
If the ASG can't launch instances for 24 hours it's going ot suspend the processes (a dministration suspensionm)
```

### 45. CloudWatch for ASG

```text
Metrics available, they're opt in!

```

### 46. Section Clean Up

```text
Cleaned up:
Auto scaling group
Launch configuration
Target Group
Load _Balancer

```

### Quiz!

```text

```