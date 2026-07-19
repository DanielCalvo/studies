Quiz: Pending

### 89. Solutions Achitecture Discussions Overview
- Stephane claims: These solution architectures are the best part of this course!
- These allow you to understand how all the technologies seen until now work together
- This is a section that you need to be 100% comfortable with!

### 90. WhatsTheTime.com
- WhatIsTheTime.com allows people to know what time it is
- No DB is required -- each server knows what time it is
- For starters, we want to start small and can accept downtime
- Later, we want to fully scale veritically and horizontally with no downtime
- Let's go through the solutions architect journey for this app

#### Starting simple
- Stateless web app
- Single t2.micro instance with an elastic IP address so the IP is fixed

#### Scaling vertically
- As more requests come in, the t2.micro instance struggles to reply to all the traffic
- You can scale vertically by resizing the instance to a m5.large, for instance. This is vertical scaling
- Stop the instance, change the instance type and start it again
- Downtime was experienced when changing the instance size :(

#### Scaling Horizontally
- Add more m5.large instances with elastic IPs attached. Not ideal as the users need to be aware of the IPs
- Remove the elastic ips and add multiple route53 A entries, one for each m5.large instance (we currently have 3)
- But if you remove an instance, and have a TTL if 1 hour, users will hit a missing instance every once in a while

#### Scaling Horizontally, with a load balancer
- 3 private instances on the same AZ (ew) with a ELB with health checks
- ELB is public facing, but the instances in the back can only be reached from the ELB. This is done with security groups
- Users query whatsthetime.com on route53, and that is an alias record to the ELB name

#### Scaling horizontally, with an auto-scaling group
- Similar deal as the last session, but this time the instances will be managed by an auto scaling group
- Now you can scale on demand!
- So you have: Route53, ELB, ASG

#### Making our app multi-az
- But if a given AZ goes down, you're in trouble as all your instances are on that AZ!
- ELB + Health checks + Multi AZ
- ASG will also spawn instances across multiple AZs

#### Let's reserve some capacity!
- By reserving the minimum capacity of your ASG (like, imagine 2 instances) you can save on costs!
- Additional instances that are spawned by the ASG group can be spawned as on demand

#### In this lecture we have discussed
- Where private IPs and public IPs on EC2 instances fit on an architecture diagram
- Using Elastic IPs vs Route53 vs Load Balancers
- Because of Route53 TTLs it is recommended to use a load balancer and alias records
- The pros and cons of maintaining EC2 instances manually vs using an ASG
- Using multi AZs to survive disasters
- Enabling ELB health checks so that we only send traffic to instances that are responding correctly
- Security group rules so that the EC2 instances only receive traffic from the ELB
- Look at reservation of capacity for costs savings when possible
- We're considering 5 pillars for a well architectured application:
    - Costs, performance, reliability, security and operational excellence

### 91. MyClothes.com
- myclothes.com allows people to buy clothes online
- There's a shopping cart!
- Our website is having hundreds of users at the same time
- We need to scale, maintain horizontal scalibility and keep our web application as stateless as possible
- Users should not lose their shopping cart
- Users should have their details (account info such as address) on a database
- Let's see how we can procceed!

#### First Diagram
- Route53, ELB, ASG with EC2 instances
- Users keeps losing his shopping cart data when he reaches different instances behind the ELB
- To fix this, you can introduce stickiness (session affinity). It's an ELB feature.
- So the requests will go to the same instance for a given period of time, but if the EC2 instance gets terminated, you still lose your shopping cart
- Different approach: Introduce user cookies (ELB feature)
    - This is good as it's stateless
    - It's bad as the HTTP requests are heavier (Stephane's way of saying they're maybe larger or more expensive to process)
    - Security risk: Cookies can be altered
    - Make sure your EC2 instances validate the content of the cookies
    - But cookies are limited to less than 4kb

#### Introducing Server Session
- Instead of sending the entire shopping cart in the cookie, send a session_id
- Have the cart information stored in ElastiCache
- When a user does another request with a session_id and it goes to another instance, that instance is able, using the session_id to look up the content of the cart on ElastiCache and retrieve that session data
- ElastiCache has sub milisecond performance, so all of this happens very quickly!
- An alternative to storing session data is Amazon DynamoDB, but we haven't seen it yet

#### Storing user data in a database
- To store and retrieve user data (such as address and name) that is long term storage, we use RDS
- Each instance can talk to RDS independently

#### Scaling reads
- To scale reads you can use an RDS master that receives writes
- But you can also introduce through replication some read replicas
- Any time you read, you can read from the read replicas. You can have up to 5 read replicas in RDS, allowing you to scale up the reads

#### Write through
- There's an alternative pattern called write through
- Instead of looking into the DB directly for use info, it looks in the cache to see if that information is present there (we cache user info on the cache)
    - If it doesn't have it, it reads from RDS and puts it into cache
    - Other instances later, taking to the cache will have a cache hit when finding informatoin about this user 
    - Cache maintenance has to be done at the application side 
    
#### Multi AZ - Survive disasters
- Multi AZ ELB, Multi AZ ASG
- RDS also has a Multi AZ feature!
- ElastiCache also has a Multi AZ feature if you use Redis 

#### About security groups
- ELB: Accepts HTTP/HTTPS from 0.0.0.0/0
- For the EC2 instances: Restrict traffic coming from the LB
- For ElastiCache and RDS: Restrict traffic to the SG of the EC2 instances

#### In this lecture we've discussed 3-tier architectures for web applications
- ELB sticky sessions
- Web clients for storing cookies and making our web app stateless
- ElastiCache
    - For storing session_ids (DynamoDB is an alternative)
    - You can also use ElastiCache to cache data from RDS
    - You can also use Multi AZ with ElastiCache
- RDS
    - For storing user data
    - Read replicas can be used for scaling reads
    - Multi AZ can also be used for disaster recovery
- Tight security can be added with security groups referencing each other

### 92. MyWordPress.com
- We're trying to create a fully scalable wordpress website
- We want that website to correctly display picture uploads
- User data and blog content should be stored in a MySQL database
- Let's see how we can achieve this!

#### RDS Layer
- If you want more scalability features, you can have Aurora MySQL on the database later with Multi AZ and Read Replicas

#### Storing images with EBS
- If you only have one EC2 instance behind your load balancer, you can store your images on an EBS volume
- The problem arrives when you start scaling: Now you have 2 instances with 2 EBS volumes. If you send an image to one instance, it won't be present on the other instance!
- EBS works with 1 instance, but with multiple instances it can become problematic

#### Storing images with EFS
- You can have EFS mounted on your EC2 instances through an ENI (Elastic Network Interface) on each instance 
- This way all the EC2 instances can access your shared images

#### In this lecture we have discussed
- Aurora database to have easy Multi-AZ and read replicas
- Storing data in EBS (single instance application)
- Storing data in EFS (distributed application)
- EBS is cheaper than EFS though

### 93. Instantiating applications quickly
- When launching a full stack (EC2, EBS, RDS), it can take a long time to
    - Install applications
    - Insert initial (or recovery data)
    - Configure everything
    - Launch the application
- You can use the cloud to speed that up!

- For EC2 instances
    -  Use a Golden AMI: Install your applications, OS dependencies and so on beforehand and then launch your EC2 instance from the Golden AMI
    - So this time you don't have to install everything every time you launch an instance
    - You can also bootstrap using user data: For dynamic configuration
    - You can also have a hybrid: Mix of Golden AMI and User Data (Elastic Beanstalk uses this!)
- RDS Databases:
    - Restore from a snapshot: The database will have schemas and data ready!
- EBS Volumes:
    - Restore from snapshot: The disk will already be formatted and have data!

### 94. Beanstalk Overview
- As a dev, you don't wanna manage infra, you just wanna deploy!
- Beanstalk, will:
- Manage infrastructure
- Deploy code
- Configure all the databases, load balancers, etc
- Scale it for you
- Most web apps have the same architecture (ALB + ASG)
- All developers want is for their code to run!
- Possibly, consistently across different applications and environments

#### AWS ElasticBeanStalk Overview
- ElasticBeanStalk is a developer centric view of deploying an application on AWS
- It uses all the components we've seen before: EC2, ASG, ELB, RDS, etc
- But it's all in one view that's easy to make sense of!
- We still have full control over the configuration 
- BeanStalk itself is free, you only pay for the underlying instances

#### ElasticBeanStack
- Managed service
    - Instance configuration / OS is handled by beanstalk 
    - Deployment strategy is configurable but performed by ElasticBeanStalk
- Just the application code is the responsibility of the developer 
- Three architecture models:
    - Single instance: Good for dev
    - LB + ASG: Great for production or pre-production environments
    - ASG only: Create for non-web apps in production (workers, etc)

#### ElasticBeanStack, part deux
- ElasticBeanStalk has three components:
    - Application
    - Application version: Each deployment gets assigned a version
    - Environment name (dev, test, prod): free naming
- You deploy application versions to environments and then can promote application versions to the next environment
- Rollback feature to previous application version 
- Full control over lifecycle of environments 

#### ElasticBeanStack, part trois
- Support for many platforms
    - Go, Java SE, Java with TomCat, .NET on Windows server with IIS, Node.js, PHP, Python, Ruby, Packer Builder
    - Single/Multi Container docker, preconfigured docker
    - If not supported, you can also write your custom platform (advanced) 

### 95. Beanstalk Hands On
- Beanstalk service! Get started
- Name the app, select platform
- Sample application
- All the options you can tweak:
    - Software
    - Instances
    - Capacity
    - Load Balancer
    - Rolling updates and deployments
    - Security
    - Monitoring
    - Managed updates
    - Notifications
    - Network
    - Database
    - Tas
- In the backend, beanstalk uses cloudformation to create the environment, you can actually check the cloudformation code in the cloudformation service
- In the background, elasticbeanstalk will create all the AWS components for us (EC2 instance, ASG, ELB and maybe something else)
- By uploading a new version, we update the application
- An application can have multiple environments (you can create it in the beanstalk application windor, top right)
- App > Configuration > You can finetune a bunch of configurations, you can change the AMI for instance

### Quiz 7: Solutions Architecture Discussions - Classic
- Pending