- Chapter status
- 51: Double check scaling policies
- 47: SSL lecture hands on is missing
- Further reading under this section: https://aws.amazon.com/blogs/aws/new-application-load-balancer-sni/
- Quiz: Try creating your own scaling policy. Terraform it if you're feeling brave

### 40. High availability and Scalability?
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
- Horizontal scaling: Increase the number of instances (scale out/in). This can be used with
    - Auto scaling groups
    - Load Balancers
- High Availability: Run instances for the same application across multi AZ
    - Auto Scaling group multi AZ
    - Load Balancer multi AZ

### 41. Elastic Load Balancing (ELB) Overview

#### What is load balancing?
- Load balancing are servers that forward internet traffic to multiple EC2 instances downstream
- Users don't connect directly to an instance, they connect to a load balancer.

#### Why use a load balancer?
- Spreads load across multiple downstream instances 
- Expose a single point of access (DNS) to your app
- Seamlessly handle failure of downstream instances
- Does healthcheck on the instances and redirects traffic to instances that are working
- Load balancers can provide HTTPs termination. Encryption/SSL is between the client and the ELB
- Can enforce stickiness with cookies so that the user talks to the same instance over time 
- Can provide high availability across zones
- Can separate public traffic from private traffic (?)

####  Why use an EC2 loadbalancer?
- And ELB (EC2 load balancer) is a managed load balancer
    - AWS guarantees that it will be working
    - AWS takes care of updates, maintenance, high availability
    - AWS provides only a few configuration knobs
- It costs less to set up your own load balancer but it will be a lot more effort on your end
- ELB is integrated with many AWS offerings/services

#### Healthchecks
- Health checks are crucial for load balancers
- They enable the load balancer to know which instances are healthy so that it can forward traffic to them
- The health check is done on a port and a route (/health is common)
- If the instance responds 200 (OK) it is considered healthy, if it responds with anything else it is considered unhealthy

#### Types of load balancers on AWS
- AWS has 3 kinds of load balancers
- Classic load balancer, or v1, from 2009
    - Supports HTTP, HTTPS, TCP
- Application load balancer, v2 - 2016
    - Supports HTTP, HTTPS, WebSocket
- Network load balancer, v2, 2017
    - Supports TCP, TLS & UDP
- Overall it's recommended to use the newer, v2 load balancers as they provide more features
- Application load balancer and network load balancer is what will be asked mainly in the exam
- You can set up **internal** (internal) or **external** (public) ELBs

#### Load Balancer Security Groups
- SG for the load balancer itself
    - Can receive HTTP and HTTPS traffic from 0.0.0.0/0
- SG for the instances that will receive traffic from the LB
    - Can receive HTTP traffic only from the LB

#### Load Balancer Good to Know
- LBs can scale but not instanteneously -- Contact AWs for a "warm up"
- Troubleshooting
    - 4xx: Client induced errors
    - 5xx: Application induced errors 
    - Load Balancer Error 503 means at capacity or no registered target
    - If the LB can't connect to your application, check your SG!
- Monitoring
    - ELB access logs will log all access requests (you can debug every single request)
    - CloudWatch Metrics will give you aggregate statistics

### 42. Classic Load Balancer (CLB) with Hands on
- Supports TCP (Layer 4) and HTTP & HTTPS (Layer 7)
- Health checks are TCP or HTTP based
- A classic load balancer gives you a fixed host name in the form of `xxx.region.elb.amazonaws.com`
- So the flow of traffic will be like:
    - Client > CLB > EC2

#### Hands on
- Create an EC2 instance on a SG that has both HTTP and SSH inboud traffic allowed
- Under EC2 dashboards > Load Balancers > Create Load Balancer > CLB
- Default config then next to assign SG
- New SG: Any one on port 80 to access the LB
- Health check can be left default with `/index.html` or just `/`
- Next and add the EC2 instance you just created as a target
- Review and create and then create!
- Now go back to your EC2 instance and install httpd in there!
- You can also edit the security group that the EC2 instance is on to only receive traffic from the CLB (so you can't access the instance directly)


### 43. Application Load Balancer (ALB) with Hands On
- The application load balancer is layer 7 (HTTP)
- It allows you to load balance multiple HTTP applications across machines
- Machines are grouped in Target Groupes
- It also allows you to route to multiple applications on the same machine
- Supports HTTP/2 and WebSockets. Supports redirects (HTTP to HTTPS for example)  
- It supports routing to different target groups:
    - Routing based on path url (like example.com/users and example.com/posts)
    - Routing based on hostname (one.example.com, other.example.com)
    - Routing based on query string (example.com/users?id=123&order=false)
- ALB are a great fit for microservices and container based applications
- ALB has a port mapping feature that allows you to redirect on a dynamic port on ECS (more on that later on the ECS section)
- To achieve what an ALB does, you would need multiple CLBs per application

#### Target Groups
- EC2 instances (can be managed by an auto scaling group, aka ASG)
- ECS tasks
- Lambda Functions (HTTP Request is translated into a JSON even)
- IP Addresses (must be private IPs)
- ALB can route to multiple target groups
- Health checks are at the target group level

#### ALB Good to know
- You get a fixed hostname for your ALB (xxx.region.elb.amazonaws.com)
- The application servers don't see the IP of the client directly
    - The IP of the client is inserted in the X-Forwarded-For header
    - You can also get the port (X-Forwarded-Port) and the protocol (X-Forwarded-Proto)

#### Hands on
- Go into LB, select Application load balancer
- Name it, all defaults will do, choose all AZs
- Reuse the myfirstloadbalancer SG which allows port 80 from anywhere
- Create a new target group, name it and defaults and go to register targets
- Welp I was supposed to have instances running webservers for targets but I'm just gonna add dcalvo.dev as target for now
- Add a webserver to a target group with a passing health check. Make sure TG is associated with the ALB and you're good!
- Under the Listener tab on the load balancer page you can set up target groups with different rules (like url PATH or header or source IP and so on)
- You can also add custom responses. Neat!

### 44. Network Load Balancer (NLB) with Hands On
- A Network load balancer (layer 4) allows you to:
    - Forward TCP & UDP traffic to your instances
    - Handle millions of requests per second. It's high performance!
    - Handle requests with a lower latency compared to an ALB. 100ms latency for NLB vs 400ms for ALB (in avg)
- A NLB has one static IP per AZ and supports assigning Elastic IPs (helpful for whitelisting a certain IP)
- NLBs are used for very high performance, or when TCP or UDP load balancing is required
- Not included in the free tier

#### Hands on
- Load balancer tab > Create Load Balancer > Network Load Balancer
- Name it, select all AZs
- Name target group whatever you want
- Add instances to target group!
- Uh-oh, NLBs do not have a security group attached to it
- Interesting: When instances are behind and NLB, they see traffic coming from the outside (Internet) as the NLB doesn't have a security group
- When creating the NLB, you can also choose an elastic IP under the AZ section on Step 1

### 45. Load Balancer Stickiness
- Stickiness makes possible to make the same client always to be redirected to the same instance behind the load balancer
- This works for classic load balancers and application load balancers
- A cookie is used for stickiness with an expiration date. As long as the cookie remains the same, the load balancer will redirect to that instance
- Use case: Make sure the user does not lose his session data
- Enabling stickiness may bring some imbalance to the load over the backend EC2 instances

#### Hands on
- On a ALB, the stickiness is set on the target group
- Once you enable stickiness, you can set a Stickiness duration

### 46. Elastic Load Balancer - Cross Zone Load Balancing
- With cross zone load balancing, each load balancer instance distributes traffic evenly across all registered instances in all AZs
- With this feature enabled, a load balancer with distribute load to all instances behind it, independently of what AZ the LB and instances are on
- Otherwise each load balancer node distributes requests evenly across instances in it's availability zone only

- Classic load balancer
    - Disabled by default
    - No charges for inter AZ traffic if enabled
- Application load balancer
    - Always on, can't be disabled
    - No charges for inter AZ traffic
- Network load balancer
    - Disabled by default
    - You pay charges ($) for inter AZ data if enabled

#### Hands on 
- CLB: On the description tab of the load balancer, Cross zone load balancing is a setting all the way down
- ALB: Enabled by default, no setting you can change
- NLB: Disabled by default, regional data transfer charges may apply

### 47. Elastic Load Balancer - SSL Certificates

#### SSL/TLS - Basics
- An SSL Certificate allows traffic between your clients and your load balancer to be encrypted while in transit (this is called in-flight encryption)
- SSL refers to Secure Sockets Layer, used to encrypt connections
- TLS refers to Transport Layer Security, which is a newer version
- Nowadays, TLS certificates are mainly used, but people still refer to this as SSL
- Public SSL certificates are issued by Certificate Authorities (CA) such as Comodo, Symantec, GoDaddy, GlobalSign, Digicert, Letsencrypt, etc...
- SSL certificates have an expiration date (you set) and must be renewed

#### More on Load balancers and SSL
- Diagram of how communication happens
    - Users > HTTPS > Load Balancer > HTTP > EC2 instance
- The load balancer performs SSL termination and talks to your EC2 instances over HTTP
- The load balancer uses an X.509 certificate (called SSL/TLS server certificate)
- You can manage certificates in AWS using ACM (AWS Certificate Manager)
- You can create and upload your own certificates alternatively
- When you set an HTTPS listener:
    - You must specify a default certificate
    - You can add an optional list of certs to support multiple domains
    - Clients can use SNI (Server Name Indication) to specify the hostname they want to reach
    - You can also specify a security policy to support older versions of SSL / TLS (legacy clients)

#### SSL - Server Name Indication (SNI)
- SNI solves the problem of loading multiple SSL certificates onto one web server (to serve multiple websites)
- It is a newer protocol and requires the client to indicate the hostname of the target server in the initial ssl handshake
- The server will then find the correct certificate or return the default one
- Only works for ALB & NLB (newer generation) and CloudFront
- Does not work with CLB

#### Diagram
- So you can have an ALB receiving traffic for 2 domains: www.mycorp.com and domain1.example.com
- The ALB then has 2 certificates configured on it: One for mycorp and other for example.com
- The ALB also has 2 target groups: One for mycorp and other for example.com
- If a client connects requesting www.mycorp.com, the ALB then serves the certificate for mycorp.com and forwards the traffic to that target group

### Elastic Load Balancers - SSL Certificates
- Classic Load Balancer (v1)
    - Supports only one SSL certificate
    - You need to use multiple CLBs for multiple hostnames with multiple SSL certificates
- Application load balancer (v2)
    - Supports multiple listeners with multiple SSL certificates
    - Uses Server Name Indication (SNI) to make it work
- Network Load Balancer (v2)
    - Supports multiple listeners with multiple SSL certificates
    - Uses Server Name Indication (SNI) to make it work
    
#### Hands on
- Follow up from Dani: Actually set up a certificate!

### 48. [SAA-C02] Elastic Load Balancer - Connection Draining
- This has 2 different names depending on which load balancer you're considering
    - CLB: Connection draining
    - Target Group: Deregistration Delay (for ALB &  NLB)
- Connection draining: It's the time to complete "in-flight requests" while the instance is de-registering or unhealthy
- The ELB, as soon as the instance is in draining mode, will stop sending requests to the instance
- But it will wait for connections to finish to the instance that is draining before it is completely deregistered.
- New connections will go to normal instances
- The deregistration delay is between 1 to 3600 seconds, the default is 300 seconds
- Can also be disabled (set value to 0)
- Set a low value if your requests are short (like 10 or 20 seconds)
- But if your EC2 instances are slow to respond, give it a higher value

### 49. Auto Scaling Groups (ASG) - Overview
- The load of your website and applications can vary
- In the cloud, you can create and get rid of servers very quickly

- The goal of an Auto Scaling Group (ASG) is to:
    - Scale out (add EC2 instances) to match an increased load
    - Scale in (remove EC2 instances) to match a decreased load
    - Ensure you have a minimum and a maximum of instances running at any given time
    - Automatically register new instances to a load balancer

#### Auto Scaling Groups In AWS
An ASG has:
    - Minimum size: The minimum size of instances to be running on an ASG at any given time
    - Actual size / Desired capacity: The number of instances that should be currently running on your ASG
    - Maximum size: The maximum number of instances that can exist for scaling out when the load goes up

#### Auto Scaling Groups in AWS With Load Balancer
- A Load balancer can load traffic directly to instances managed by a load balancer. Instances added on scale out will be added as targets and health checks will be performed on them

#### ASGs have the following attributes
- A launch configuration
    - AMI + Instance type
    - EC2 User Data
    - EBS Volumes
    - Security groups
    - SSH Key Pair
- Min Size / Max size / Initial capacity
- Network and Subnets in which the ASG will be able to create instances
- Load Balancer Information / Target Group Information
- Scaling policies (so what will trigger a scale out or scale in)

#### Auto Scaling Alarms
 - It's possible to scale your ASG based on a CloudWatch alarm
 - An Alarm monitors a metric (such as average CPU)
 - Metrics are computed for the overall ASG instances (not from just a single instance!)
 - Based on the alarm
    - You can create scale-out policies to increase the number of instances
    - You can create scale-in policies to decrease the number of instances

#### Auto Scaling new rules
- It is now possible to define auto scaling rules that are directly defined by AWS, such as:
    - Target Average CPU usage (ex: I want my CPU usage to be 40% on avg)
    - Number of requests on the ELB per instance
    - AVG Net In
    - AVG Net Out
- These rules are easier to set up than the previous ones and they can make more sense

#### Auto Scaling
- You can auto scale based on a custom metric (ex: number of connected users)
- To do this, you
    1. Send a custom metric from your application on EC2 to CloudWatch through the PutMetric API
    2. Create a CloudWatch alarm to react to low / high values
    3. Use the CloudWatch alarm as the scaling policy for the ASG

#### ASG Brain Dump
- Scaling policies can be things like CPU, Network or can can be custom metrics or based on a schedule (if you know your load patterns)
- ASG can use Launch Configurations or Launch Templates (newer)
- To update an ASG, you must provide a new launch configuration / launch template. EC2 instances will be replaced over time
- IAM roles attached to an ASG will get assigned to EC2 instances
- ASG are free. You pay for the underlying resources being launched
- Having instances under an ASG means that if they get terminated for whatever reason, the ASG will automatically create new ones as a replacement. Extra safety!
- (for example) An ASG can terminate instances marked unhealthy by an LB (and replace them)

### 50. Auto Scaling Groups - Hands On
- Make sure you have a LB for HTTP on / and a TG with no targets. Just create empty ones for now for this hands on

#### Step 1: Choose launch template or configuration
- EC2 > Auto Scaling Group > Create Auto Scaling Group
    - Launch template and launch configuration are pretty much the same thing
- Click on Create Launch template
    - Launch template describes how to create EC2 instances
- On the launch template creation page:
    - AMI: Amazon Linux 2
    - Instance Type: t2.micro
    - Keypair: Yours!
    - SG: Something that allows SSH and HTTP
    - On the advanced settings at the very bottom, you can add user data!

#### Step 2: Configure settings
- Now go back to the ASG console to create an ASG and select your launch template
- Adhere to launch template
- Select 3 subnets (I used the default VPC)

#### Step 4: Load Balancing and health checks
- Enable load balancing, choose your TG (which is now empty)
- Chose ELB health check

#### Step 4: Configure group size and scaling policies
- Des/Min/max: 1,1,3
- Scaling policy: None (we'll do that later)

#### Other notes
- Hey cool everything works
- You can add more instances to the ASG by increasing the target size

### 51. Auto Scaling Groups - Scaling Policies

#### The scaling policies available are
- Target Tracking scaling
    - Simplest and easiest to set up
    - Example 1: I want the average ASG CPU usage to stay at around 40%
    - If you're over that CPU usage, more instances will be provisioned
    - If you're under that CPU usage, scaling in will kick in and instances will be terminated
- Simple / Step Scaling policy
    - When a CloudWatch alarm is triggered (ex: > 70% CPU usage) then add 2 units
    - When a CloudWatch alarm is triggered (ex: < 30% CPU usage) then remove 1 unit/instance (I think)
- Scheduled actions
    - Anticipate scaling based on known usage patterns
    - Example: Increase the min capacity to 10 at 5pm on Fridays

#### Auto Scaling Groups - Scaling Cooldowns
- The cooldown period helps to ensure that your Auto Scaling group doesn't launch or terminate additional instances before the previous scaling activity takes effect.
- In addition to default cooldown for Auto Scaling group, we can create cooldowns that apply to a specific simple scaling policy
- A scaling-specific cooldown period overrides the default cooldown period
- One common use for scaling-specific cooldowns is with a scale-in policy—a policy that terminates instances based on a specific criteria or metric. Because this policy terminates instances, Amazon EC2 Auto Scaling needs less time to determine whether to terminate additional instances.
- If the default cooldown period of 300 seconds is too long—you can reduce costs by applying a scaling-specific cooldown period of 180 seconds to the scale-in policy
- If your application is scaling up and down multiple times each hour, modify the Auto Scaling Groups cool-down timers and the CloudWatch Alarm Period that triggers the scale in

#### Hands on
- Go your ASG and then "Automatic Scaling"
- Scaling policy > Add Policy
- Easy!
    - Policy type: Target tracking scaling
    - Scaling policy name: Whatever
    - Metric Type: AVG CPU Utilization
    - Target Value: 40
    - Cooldown: Left default (300)
    - Did not disable scale in
- This creates CloudWatch alarms in the background. Neat!

### 52. Auto Scaling Groups - For Solutions Architects

#### ASG Default Termination Policy
- When terminating an instance on an ASG, AWS will attempt to delete the instance with the following criteria
    1. Find the AZ which has the most number of instances
    2. If there are multiple instances in the AZ to chose from to delete, the one with the oldest launch configuration will be eleted
- ASG tries to balance the number of instances across AZs by default

#### Lifecycle hooks
- By default, as soon as an instance is launched in an ASG it's in service
- You have the ability to perform extra steps before the instance goes in service. You can do extra things (like install extra software) before your instance is considered in service. Neat!
- You can also define the lifecycle hook on the terminating state. To extract files or logs from the instance, for instance
- https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html

#### Launch Template vs Launch Configuration
- Both:
    - Allow you specify the ID of the Amazon Machine Image (AMI), the instance type, keypair, security groups and other parameters you use to launch an EC2 instance (such as user data, tags, etc)

- Launch configuration (legacy):
    - Must be recreated every time you want to change a parameter

- Launch Template (newer):
    - Can have multiple versions
    - Can have parameter subsets (partial configuration for re-use and inheritance)
    - Allows you to provision a mix of on-demand and spot instances
    - Can use the T2 unlimited burst feature
    - Recommended by AWS going forward

### Quiz!
- 17/20, could improve!

