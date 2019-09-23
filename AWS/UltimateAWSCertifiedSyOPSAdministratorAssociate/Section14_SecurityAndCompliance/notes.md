### 162. Intro!
- Uh-oh, one of the hardest sections, we learn about many things!

### 163. Shared responsibility model
- AWS's responsibility: Security **of** the cloud
    - Hardware, software, facilities, network, etc, everything running the AWS's services
    - Also responsible for managed services like S3, DynamoDB and many others 
- Customer's responsibility: Security **in** the cloud
    - For an EC2 instance, the customer is responsible for the management of the gues OS (including patches and updates), firewall & network config, IAM, etc

#### Example: RDS
AWS's responsbility:
    - Manage the underlying EC2 instance with no SSH access for you
    - Automated DB patching
    - Automated OS patching
    - Audit the underlying instance, disks & guarantee it functions
- Your responsability:
    - Check the ports / IP / security groups in the DB's SG
    - In-database user creation and permissions
    - Creating a databasw with or without public access
    - Ensure parameter groups or DB is configured only to allow SSL connections    
    - Database encryption setting

#### Example: S3
- AWS:
    - Guarantee unlimited storage
    - Guarantee encryption
    - Ensure separation of the data between different customers
    - Ensure AWS employees can't access your data
- You:
    - Bucket configuration
    - Bucket policy / public setting
    - IAM user and roles
    - Enabling encryption

Also: https://aws.amazon.com/compliance/shared-responsibility-model/

### 164. DDoS, AWS Shield and AWS WAF 
- Aha, I know what a DDoS attack is :D

#### DDOS Protection on AWS
- AWS Shield Standard: Enabled by default free of charge, protects against DDoS attacks for your website and application.
- AWS Shield Advanced: 24/7 premium DDoS protection
- AWS WAF: Filters specific requests based on rules, like requests being too large (WAF = Web Application Firewall)
- Cloudfront and Route 53
    - Availability protection using global edge network
    - Combined with AWS Shield, provides attack mitigation at the edge
- An action point you can take: Be ready to scale -- leverage AWS auto scaling
- You can also separate static resources (html, css, images into S3/CloudFront) from dynamic ones (EC2/ALB). It's harder to DDoS your application this way.

Author recommends reading: `https://d1.awsstatic.com/whitepapers/Security/DDoS_White_Paper.pdf`

#### AWS Shield
- AWS Shield standard:
    - Free service that is activated for every AWS customer
    - Provides protection from attacks such as SYN/UDP floods, reflection attacks and other layer 3/4 attacks
- AWS Shield advanced:
    - The exam asks about it! 
    - Optional DDoS migration service (3000 USD per month per organization)
    - Protects against more sophisticated attack on CloudFront, Route 53, Classic (?), Application & Network Load Balancer, Elastic IP / EC2
    - 24/7 access to a AWS DDoS response team (DRP)
    - If you incur higher fees due to a DDoS attack, these fees are waived if you have AWS Shield Advanced

#### AWS WAF - Web Application Firewall
- Protects your web application from common web exploits
- You can define customizable web security rules:
    - Allow/block traffic to your application based on IP addresses, HTTP headers, HTTP body or URI strings
    - Protects from common attacks - SQL Injections and Cross-Site Scripting (XSS)
    - Protects against bots and bad user agents
    - You can also put size contraints on requests (ex: no requests larger than 2mb)
    - Geo match
- You can deploy this on CloudFront, Application Load Balancewr or API Gateway
- There is also a marketplace of rules
    
AWS Shield: DDoS protection
AWS WAF: Security rules to block certain types of traffic

### 165. Penetration testing on AWS
- Permission is required for all tests! 
- Request permission only using AWS root credentials
- You must perform the pen testing yourself. No third parties!
- Things you can pen test against: EC2, ELB, Aurora, CloudFront, API Gateway, Lambda, Lightsail.
- The rest you can't pen test again
- You also cannot pentest against nano/micro/small EC2 instances
- Pentest request takes about 2 days to be approved
- You need to specifically list the endpoints/resources you want to attack on AWS.
- You also need to specify where you traffic is coming from! (IP addresses)

### 166. AWS Inspector
- Only for EC2 instances. The exam will try to trick you! (ex: Can you use inspector on RDS? No you can't)
- Helps you analyze against known security vulnerabilities
- Also helps analyze  against unintended network accessibility
- You need to install an inspector agent on EC2 instances
- You define a template (rules package, duration, attributes, SNS topics)
- No own custom rules possible - you can only use AWS's managed rules
- An assessment runs and after it does, you get a report with a list of vulnerabilities found

#### What does AWS inspector evaluate
- Remember: Only for EC2 instances
- For network assessments
    - Network Reachability
For Host assessments
    - Common vulnerabilities and exposures
    - Center for Internet Security (CIS) Benchmarks
    - Security Best Practices
    - Runtime Behaviour Analysis

#### Inspector Hands on
- Create an EC2 in a tag (such as Inspector: True, but can be whatever you want)
- You can install the agent either by googling it and doing it manually or automating it through SSM
- Very interesting! You get vulnerability information.
- Possible exam question: "We need to know of our instances are vulnerable to attacks or if they need to be patched." Use AWS Inspector!

### 167. Logging in AWS

### 168. GuardDuty

### 169. Trusted Advisor

### 170. Encryption 101

### 171. KMS Overview + Encryption in place

### 172. CloudHSM Overview

### 173. KMS + CloudHSM Hands On

### 174. MFA + IAM Credentials Report

### 175. IAM PassRole Action

### 176. STS & Cross Account Access

### 177. Identity Federation with SAML & Cognito

### 178. JSON Policies tab

### 179. AWS Compliance frameworks

### 180. AWS Artifact

### 181. Security and Compliance Section Summary

### Quiz