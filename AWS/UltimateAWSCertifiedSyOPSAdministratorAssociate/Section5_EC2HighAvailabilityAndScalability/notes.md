### 34. Section introduction

```text
Load balancers and auto scaling groups!
Troubleshooting!
Advanced options!
CloudWatch Integrations!
```

### 35. What is high availability and Scalability?

```text
Your system can handle greater loads by scaling
There's two types of scalability: Vertical and Horizontal

Vertical: Means increasing the size of your instance. RDS and Elasticache can scale vertically well, although there are limits to this.
Horizontal: Means increasing the number of instances/systems for your application

Horizontal scaling implies distributed systems.
Not every application can be a distributed system.
Easy to horizontally scale on EC2!

High Availability: Means running your app in at least 2 datacenters or 2 AZs in AWS.
Goal: Survive datacenter loss.

An Auto scaling group can be multi-az
A load balancer can be multi-az
```

### 36. Load Balancer overvieww

```text
Load balancing are servers that foward internet traffic to multiple EC2 instances in this case
- It spreads load across multiple instances downstream
- Expose a single point of access to your app
- Seamlessly handle failure of downstream instances
- Does helthcheck on the instances and redirects traffic to instances that are working
- Load balancers can provide HTTPs termination
- Can enforce stickiness with cookies
- Can provide high availability across zones
- Can separate public traffic from private traffic (?)

There's an EC2 load balancewr named ELB. It's a managed load balancer
EC2 ensures it's working, take care of updates, provide some configurations settings
It is integrated with many AWS services

There are different types of LBs:
- Classic load balancer, or V1, from 2009.
- Application loadbalancer, v2 - 2016
- Network load balancewr, v2, 2017
Overall it's recommended to use the newer, v2 load balancers.

You can set up internal (internal) or external (public) load balancer

Healthchecks: 
These allow the LB to know which instances to forward traffic to.
Healthcheck is done to a port and route (like /health)

Application load balancers (layer 7)
These are called L7 as they allow you to work at the HTTP level. You group applications in target groups.
You can load balancer to multiple applications on the same machine (common exam question)
You can load balance based on the route or hostname (?) on the URL.

Good fit for micro services and container based apps.

There's a port mapping feature to redirect to a dynamic port.

Good to know:
- Stickiness can be set at the target group level.
- The application servsrs don't see the IP of the client direcetly. The true IP of the client is inserted in the header "X-Forwaded-For"
 
 
Network Load Balancers (v2) Layer 4
TCP traffic.
Handles high throughput
Support for static IP or elastic IP

Load Balancer Good to know:
- Classic ones are deprecated
- Use application load balancers for HTTP / HTTPS & healthsocket
- Network Load balancer for TCP
- Load balancers support SSL termination
- All load balancers have health check capabilities
- Any LB has a static host name. Do not try to resolve and use the underlying IP.
- LB can scale but not instantly.

There are 3 types of LB: CLB, NLB and ALB

If the LB can't connect to you app, check your security group!

 
```

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