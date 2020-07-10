### 272. WhitePaper Section Introduction 
- We'll see Architected Framework Whitepaper
- We'll see Architected Tool
- AWS Trusted Advisor
- Reference architectures resources (for real-world)
- Disaster Recovery on AWS Whitepaper
- Dani says: Seems noice!

### 273. Well Architected Framework + Well Architected Tool
- Stop guessing your capacity needs (use autoscaling when you can)
- Test systems at production scale
- Automate to make architectural experimentation easier
- Allow for evolutionary architectures. It's the cloud, change can happen much faster!
    - Design based on changing requirements
    - Drive architectures using data (what are your patterns? usage?)
- Improve through game days
    - Simulate applications for flash sale days
    - Simulate load!

#### Well Architected Framework 5 pillars
- 1. Operational Excellence
- 2. Security
- 3. Reliability
- 4. Performance Efficiency
- 5. Cost Optimization
- They are not something to balance, or trade-offs, they’re a synergy

#### Well Architected Framework
- It’s also questions!
- Let’s look into the Well-Architected Tool
- https://console.aws.amazon.com/wellarchitected
- Damn that's really cool

### 274. 1st pillar: Operational Excellence
- Includes the ability to run and monitor systems to deliver business value and to continually improve supporting processes and procedures
- Design Principles
    - Perform operations as code - Infrastructure as code
    - Annotate documentation - Automate the creation of annotated documentation after every build
    - Make frequent, small, reversible changes - So that in case of any failure, you can reverse it
    - Refine operations procedures frequently - And ensure that team members are familiar with it
    - Anticipate failure
    - Learn from all operational failures

#### Operational Excellence AWS Services
- Prepare
    - AWS CloudFormation
    - AWS Config
- Operate
    - AWS CloudFormation
    - AWS Config
    - AWS CloudTrail
    - Amazon CloudWatch
    - AWS X-Ray
- Evolve
    - AWS CloudFormation
    - AWS CodeBuild
    - AWS CodeCommit
    - AWS CodeDeploy
    - AWS CodePipeline

### 275. 2nd Pillar: Security 
- Includes the ability to protect information, systems, and assets while delivering business value through risk assessments and mitigation strategies
- Design Principles
    - Implement a strong identity foundation - Centralize privilege management and reduce (or even eliminate) reliance on long-term credentials - Principle of least privilege - IAM
    - Enable traceability - Integrate logs and metrics with systems to automatically respond and take action
    - Apply security at all layers - Like edge network, VPC, subnet, load balancer, every instance, operating system, and application
    - Automate security best practices
    - Protect data in transit and at rest - Encryption, tokenization, and access control
    - Keep people away from data - Reduce or eliminate the need for direct access or manual processing of data
    - Prepare for security events - Run incident response simulations and use tools with automation to increase your speed for detection, investigation, and recovery

#### Security AWS Services
- Identity and Access Management
    - IAM
    - AWS-STS
    - MFA token
    - AWS Organizations
- Detective Controls
    - AWS Config
    - AWS CloudTrail
    - Amazon CloudWatch
- Infrastructure protection
    - Amazon CloudFront
    - Amazon VPC
    - AWS Shield
    - AWS WAF
    - Amazon Inspector
- Data Protection:
    - KMS
    - S3
    - Elastic Load Balancing (ELB)
    - Amazon EBS
    - Amazon RDS
- Incident Response
    - IAM
    - AWS CloudFormation
    - Amazon CloudWatch Events

### 276. 3rd Pillar: Reliability
- Ability of a system to recover from infrastructure or service disruptions, dynamically acquire computing resources to meet demand, and mitigate disruptions such as misconfigurations or transient network issues
- Design Principles
    - Test recovery procedures - Use automation to simulate different failures or to recreate scenarios that led to failures before
    - Automatically recover from failure - Anticipate and remediate failures before they occur
    - Scale horizontally to increase aggregate system availability - Distribute requests across multiple, smaller resources to ensure that they don't share a common point of failure
    - Stop guessing capacity - Maintain the optimal level to satisfy demand without over or under provisioning - Use Auto Scaling
    - Manage change in automation - Use automation to make changes to infrastructure

#### Reliability and AWS Services
- Foundations
    - IAM
    - Amazon VPC
    - Service Limits
    - AWS Trusted Advisor
- Change management
    - AWS Auto Scaling
    - Amazon CloudWatch
    - AWS CloudTrail
    - AWS Config
- Failure management
    - Backups
    - AWS CloudFormation
    - Amazon S3
    - Amazon S3 Glacier
    - Amazon Route 53

### 277. 4rd Pillar: Performance Efficiency
- Includes the ability to use computing resources efficiently to meet system requirements, and to maintain that efficiency as demand changes and technologies evolve
- Design Principles
    - Democratize advanced technologies - Advance technologies become services and hence you can focus more on product development
    - Go global in minutes - Easy deployment in multiple regions
    - Use serverless architectures - Avoid burden of managing servers
    - Experiment more often - Easy to carry out comparative testing
    - Mechanical sympathy - Be aware of all AWS services


#### Efficiency and RDS services
- Selection
    - AWS Auto Scaling
    - AWS Lambda
    - EBS
    - S3
    - RDS
- Review 
    - CloudFormation
    - AWS News Blog
- Monitoring
    - Amazon CloudWatch
    - AWS Lambda
- Tradeoffs
    - Amazon RDS
    - Amazon ElastiCache
    - AWS Snowball
    - Amazon CloudFront

### 278. 5th Pillar: Cost Optimization
- Includes the ability to run systems to deliver business value at the lowest price point
- Design Principles
- Adopt a consumption mode - Pay only for what you use
- Measure overall efficiency - Use CloudWatch
- Stop spending money on data center operations - AWS does the infrastructure part and enables customer to focus on organization projects
- Analyze and attribute expenditure - Accurate identification of system usage and costs, helps measure return on investment (ROI) - Make sure to use tags
- Use managed and application level services to reduce cost of ownership - As managed services operate at cloud scale, they can offer a lower cost per transaction or service

#### Cost Optimization services
- Expenditure Awareness
    - AWS Budgets
    - AWS Cost and Usage Report
    - AWS Cost Explorer Reserved Instance Reporting
- Cost-Effective Resources
    - Spot instance
    - Reserved instance
    - Amazon S3 Glacier
- Matching supply and demand
    - AWS Auto Scaling
    - AWS Lambda
- Optimizing Over Time
    - AWS Trusted Advisor
    - AWS Cost and Usage Report

### 279. AWS Trusted Advisor Overview + Hands On
- No need to install anything – high level AWS account assessment
- Analyze your AWS accounts and provides recommendation:
    - Cost Optimization
    - Performance
    - Security
    - Fault Tolerance
    - Service Limits
- Popular question: How do we know if we're reaching out service limit? How can we get a general overview of that? Trusted advisor!
- Core Checks and recommendations – free for all customers, but some things are disabled unless you have a business support plan
- Can enable weekly email notification from the console
- Full Trusted Advisor – Available for Business & Enterprise support plans (for busineses)
- Ability to set CloudWatch alarms when reaching limits
- You can enable weekly email notification preferences on the UI for trusted advisor (possible exam question!)


### 280. Examples of Architecture - AWS Certified Solutions Architect Associate
- We’ve explored the most important architectural patterns:
    - Classic: EC2, ELB, RDS, ElastiCache, etc...
    - Serverless: S3, Lambda, DynamoDB, CloudFront, API Gateway, etc...
- If you want to see more AWS architectures:
    - https://aws.amazon.com/architecture/
    - https://aws.amazon.com/solutions/