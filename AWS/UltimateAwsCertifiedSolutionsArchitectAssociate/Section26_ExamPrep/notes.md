### 281. Exam Preparation - Section intro 
- Yay last section!

### 282. State of Learning Checkpoint - AWS Certified Solutions Architect Associate
- https://d1.awsstatic.com/training-and-certification/docs-sa-assoc/AWS-Certified-Solutions-Architect-Associate_Exam-Guide.pdf

### 283. Exam Tips - AWS Certified Solutions Architect Associate

#### Practice makes perfect
- If you’re new to AWS, take a bit of AWS practice thanks to this course before rushing to the exam
- The exam recommends you to have one or more years of hands-on experience on AWS
- Practice makes perfect!
- If you feel overwhelmed by the amount of knowledge you just learned, just go through it one more time

#### Proceed by elimination
- Most questions are going to be scenario based
- For all the questions, rule out answers that you know for sure are wrong
- For the remaining answers, understand which one makes the most sense
- There are very few trick questions
- Don’t over-think it
- If a solution seems feasible but highly complicated, it’s probably wrong

#### Skim the AWS Whitepapers
- You can read about some AWS White Papers here:
- Architecting for the Cloud: AWS Best Practices
- AWS Well-Architected Framework
- AWS Disaster Recovery (https://aws.amazon.com/disaster-recovery/)
- Overall we’ve explored all the most important concepts in the course
- It’s never bad to have a look at the whitepapers you think are interesting!

#### Read each service's FAQ
- FAQ = Frequently asked questions
- Example: https://aws.amazon.com/vpc/faqs/
- FAQ covers a lot of the questions asked at the exam
- They help confirm your understanding of a service

### 284. Exam Walkthrough and Sign up 
- https://www.aws.training/Certification

### 285. Save 50% on your AWS Exam Cost
- Only if you have done an exam before :(

### 286. Get an Extra 30 minutes on your AWS Exam - Non Native English Speakers Onlyfunctionality
- Don't forget to request ESL +30 MINUTES!


#### Stuff I forgot
- Amazon EMR
    - Elastic Map Reduce
    - Helps creating hadoop clusters
    - If you see hadoop, if you see spark, if you see big data processing, think EMR
- Glue: Does ETL (Extract, transform, load). If you see Glue, think ETL- Serverless, provisions apache spark
- AppSync
    - A service to store and sync data cross mobile and web apps in real time
    - Managed GraphQL
    - On the exam: Synchronizing data between mobile and web apps in real time, it's either cognito sync or app sync
- AWS ElasticSearch: Not serverless (you have to define an instance size)
- Redshift: Analytics and data warehousing
- Have another look at ElastiCache: Ok!
- Lambda limits: 15 minutes execution, up to 3gig memory, 512mb disk in /tmp (container)
- SSE-C: Key is managed by you and sent on header on S3 upload to that it is encrypted server side
- Have a look at RDS multi-az VS stand by or whatever it was: Done!
- Review the difference between pilot light and warm standby real quick 'cause I don't remember well

#### Less common network services
- VPC Endpoint: Allows you to connect to AWS services (S3, DynamoDB, RDS, etc) using a private network instead of the public www network
- Site to Site VPN: Makes your corporate network and an AWS VPC look like they're in the same network. You need to set up a customer gateway on your DC, and then on the VPC side you provision a VPN gateway, and then in between the VPN gateway and the customer gateway, you set up a Site to Site VPN Connection
- Site to Site VPN: Links VPN Gateway (AWS) and Customer Gateway (on prem)
- VPN Gateway: What you set between a VPN Gateway on AWS and a Customer Gateway on prem
- Customer gateway: Endpoint on the on prem datacenter that can be used on a Site to Site VPN
---
- Direct Connect: Provides a dedicated private connection (non-internet) from a remote network to your VPC. Used when you need better network connectivity
- Direct Connect Gateway: Direct Connect to one or more VPC in many regions 
---
AWS PrivateLink: Secure & Scalable way to expose a service to 1000s of VPC
VPN CloudHub: Provides secure communication between sites, if you have multiple VPN connections
Transit Gateway: Star gateway between all your VPC and on prem infrastructure. Hub and star

### Practice Test 1
- Double check the exotic networking classes again (vpc peering? Direct Connect Gateway? site to site vpn?)

### Practice Test 2: Extra Questions 
- Forgot what spot fleet was
- CloudHSM?
- Review AWS Shield and AWS WAF