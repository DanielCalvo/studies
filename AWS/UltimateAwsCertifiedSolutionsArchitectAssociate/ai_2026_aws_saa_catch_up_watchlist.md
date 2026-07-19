# 2026 AWS SAA Catch-Up Watch List

Created: 2026-07-14

Sources used:

- Local Udemy saved page: `Course_ Ultimate AWS Certified Solutions Architect Associate 2026 _ Udemy.html`
- Previous local research note: `ai_2026_aws_saa_refresh_curriculum.md`
- Official AWS Certified Solutions Architect - Associate page: https://aws.amazon.com/certification/certified-solutions-architect-associate/
- Official SAA-C03 exam guide: https://d1.awsstatic.com/training-and-certification/docs-sa-assoc/AWS-Certified-Solutions-Architect-Associate_Exam-Guide.pdf

## Summary

The saved Udemy page shows 425 course items, with 202 completed and 223 currently not completed. I would not treat all 223 as equally important. Some are old foundational lectures that probably reset because the course was updated, while others map directly to areas that became more important since a 2020-era SAA course.

The current AWS exam is still SAA-C03 according to the AWS certification page and exam guide. The official domain weights are:

- Design Secure Architectures: 30%
- Design Resilient Architectures: 26%
- Design High-Performing Architectures: 24%
- Design Cost-Optimized Architectures: 20%

My recommendation: watch Priority 1 and Priority 2. Skim Priority 3 only where you feel rusty, and do all section quizzes plus the practice exam.

## Priority 1: Watch Fully

These are the highest-value lectures for catching up from 2020 to the current SAA-C03-style exam and to modern AWS architecture discussions.

### Containers

Watch all of Section 18. It is completely unwatched in the saved Udemy page.

- 198. Docker Introduction
- 199. Amazon ECS
- 200. Creating ECS Cluster - Hands On
- 201. Creating ECS Service - Hands On
- 202. Amazon ECS - Auto Scaling
- 203. Amazon ECS - Solutions Architectures
- 204. Amazon ECS - Clean Up - Hands On
- 205. Amazon ECR
- 206. Amazon EKS - Overview
- 207. Amazon EKS - Hands On
- Quiz 15: Containers on AWS Quiz

What to look for:

- ECS task role vs task execution role.
- ECS on EC2 vs Fargate.
- ALB integration, service scaling, and deployment behavior.
- ECR image storage and basic vulnerability scanning concepts.
- EKS at the solution-architecture level, not Kubernetes administration depth.

### IAM, Organizations, and Multi-Account Architecture

Watch these from Section 25:

- 284. Organizations - Tag Policies
- 286. IAM - Resource-based Policies vs IAM Roles
- 289. AWS Directory Services
- 290. AWS Directory Services - Hands On
- 291. AWS Control Tower
- Quiz 22: IAM Advanced Quiz

Also refresh these from Section 4 if you do not feel fluent:

- 14. IAM Policies
- 16. IAM MFA Overview
- 18. AWS Access Keys, CLI and SDK
- 28. IAM Security Tools Hands On
- 29. IAM Best Practices
- 30. IAM Summary
- Quiz 1: IAM & AWS CLI Quiz

What to look for:

- IAM Identity Center / AWS SSO as the modern default for workforce access.
- AWS Organizations, SCPs, delegated admin, and Control Tower.
- Difference between identity policies, resource policies, permission boundaries, and SCPs.
- Why modern AWS designs use multiple accounts instead of one large account.

### Security Services and Encryption

Watch these from Section 26:

- 296. KMS - Multi-Region Keys
- 297. S3 Replication with Encryption
- 298. Encrypted AMI Sharing Process
- 302. AWS Secrets Manager - Hands On
- 303. AWS Certificate Manager (ACM)
- 307. Firewall Manager
- 308. WAF & Shield - Hands On
- 309. DDoS Protection Best Practices
- 310. Amazon GuardDuty
- 311. Amazon Inspector
- 312. Amazon Macie
- Quiz 23: AWS Security & Encryption Quiz

Optional hands-on if rusty:

- 295. KMS Hands On w/ CLI
- 300. SSM Parameter Store Hands On (CLI)

What to look for:

- KMS key policy vs IAM policy.
- Secrets Manager vs SSM Parameter Store.
- WAF vs Shield vs Firewall Manager vs Network Firewall.
- GuardDuty, Inspector, and Macie use cases.
- Certificate Manager and TLS termination patterns.

### Networking Refresh

Watch these from Section 27:

- 334. VPC Endpoints Hands On
- 335. VPC Flow Logs
- 340. Direct Connect + Site to Site VPN
- 342. VPC Traffic Mirroring
- 343. IPv6 for VPC
- 344. IPv6 for VPC - Hands On
- 345. Egress Only Internet Gateway
- 350. AWS Network Firewall
- Quiz 24: VPC Quiz

Also watch or rewatch if you are not comfortable:

- 327. NAT Gateways Hands On
- 328. Regional NAT Gateway
- 330. NACL & Security Groups Hands On
- 332. VPC Peering Hands On
- 338. Site to Site VPN, Virtual Private Gateway & Customer Gateway Hands On

What to look for:

- Gateway endpoints vs interface endpoints.
- Why VPC endpoints can reduce NAT Gateway cost and improve private access.
- Security group vs NACL vs WAF vs Network Firewall.
- Direct Connect vs VPN.
- IPv6 egress-only internet gateway.
- Private connectivity tradeoffs: peering, Transit Gateway, PrivateLink, VPN, Direct Connect.

### Data, Analytics, and Streaming

Watch all unwatched lectures from Section 22:

- 245. QuickSight
- 246. Glue
- 247. Lake Formation
- 248. Amazon Managed Service for Apache Flink
- 249. Amazon Managed Service for Apache Flink - Hands On
- 250. MSK - Managed Streaming for Apache Kafka
- Quiz 19: Data & Analytics Quiz

Also watch these from Section 17:

- 194. Amazon Data Firehose
- 195. Amazon Data Firehose - Hands On
- Quiz 14: Messaging & Integration Quiz

What to look for:

- Athena, Glue, Lake Formation, QuickSight at a use-case level.
- Kinesis vs Data Firehose vs MSK.
- Streaming vs batch ingestion.
- Data lake access control and transformation concepts.

### Event-Driven and Serverless Architecture

Watch these from Section 24:

- 273. EventBridge Overview (formerly CloudWatch Events)
- 274. Amazon EventBridge - Hands On
- 278. CloudTrail - EventBridge Integration
- Quiz 21: Monitoring & Auditing Quiz

Watch these from Section 19:

- 213. Lambda Concurrency
- 214. Lambda Concurrency - Hands On
- 215. Lambda SnapStart
- 217. Lambda in VPC
- 218. RDS - Invoking Lambda & Event Notifications
- Quiz 16: Serverless Overview Quiz

What to look for:

- EventBridge vs SNS vs SQS.
- Lambda concurrency, throttling, cold starts, and SnapStart.
- Lambda in VPC networking implications.
- Event-driven integration with CloudTrail, S3, RDS, and Step Functions.
- Idempotency and dead-letter handling.

### Modern Storage, S3, and Backup

Watch these from Section 12:

- 136. S3 Replication
- 137. S3 Replication Notes
- 138. S3 Replication - Hands On
- 141. S3 Express One Zone
- Quiz 9: Amazon S3 Quiz

Watch these from Section 13:

- 143. S3 Lifecycle Rules - Hands On
- 144. S3 Requester Pays
- 146. S3 Event Notifications - Hands On
- 148. S3 Batch Operations
- 149. S3 Storage Lens
- Quiz 10: Amazon S3 Advanced Quiz

Watch these from Section 14:

- 151. About DSSE-KMS
- 152. S3 Encryption - Hands On
- 163. S3 Access Points
- 164. S3 Object Lambda
- Quiz 11: Amazon S3 Security Quiz

Watch these from Section 16:

- 179. AWS Transfer Family
- Quiz 13: AWS Storage Extras Quiz

Watch these from Section 28:

- 357. AWS Backup
- 358. AWS Backup - Hands On
- Quiz 25: Disaster Recovery & Migration Quiz

What to look for:

- S3 lifecycle rules, replication, Object Lock, access points, and encryption choices.
- S3 Express One Zone use case and tradeoffs.
- Storage Lens for visibility.
- Transfer Family vs DataSync vs Storage Gateway.
- AWS Backup plans, vaults, restore expectations, cross-account/cross-region concepts.

### Disaster Recovery and Migration

Watch these from Section 28:

- 352. Elastic Disaster Recovery (DRS)
- 354. Database Migration Service (DMS) - Hands On
- 355. RDS & Aurora Migrations
- 359. Application Migration Service (MGN)
- 361. VMware Cloud on AWS

What to look for:

- Backup and restore, pilot light, warm standby, and active-active.
- RPO/RTO-driven design.
- DMS vs Application Migration Service vs Elastic Disaster Recovery.
- Aurora/RDS migration and replication patterns.

### Cost Optimization and Operations

Watch these from Section 30:

- 373. SSM Session Manager
- 374. SSM Other Services
- 375. AWS Cost Explorer
- 376. AWS Cost Anomaly Detection
- 377. AWS Outposts
- 378. AWS Batch
- 381. Instance Scheduler on AWS
- Quiz 27: Other Services Quiz

Also watch:

- 31. AWS Budget Setup
- 44. EC2 Instance Purchasing Options
- 46. EC2 Instances Launch Types Hands On

What to look for:

- Budgets, Cost Explorer, Cost Anomaly Detection, tags, and CUR at a use-case level.
- Savings Plans vs Reserved Instances vs Spot.
- NAT Gateway and cross-AZ/cross-region data transfer cost traps.
- SSM Session Manager as the modern alternative to SSH bastion-first access.
- Outposts and hybrid compute at a conceptual level.

## Priority 2: Watch or Rewatch for Exam Readiness

These are not necessarily new since 2020, but they are important enough that you should not leave them weak if you might sit the exam again.

### Load Balancing and Scaling

Watch:

- 73. Application Load Balancer (ALB) - Hands On - Part 1
- 74. Application Load Balancer (ALB) - Hands On - Part 2
- 76. Network Load Balancer (NLB) - Hands On
- 77. Gateway Load Balancer (GWLB)
- 81. Elastic Load Balancer - SSL Certificates - Hands On
- 85. Auto Scaling Groups - Scaling Policies
- Quiz 5: High Availability & Scalability Quiz

Focus:

- ALB vs NLB vs GWLB.
- Listener rules, target groups, health checks.
- TLS certs with ACM.
- Scaling policy types and metrics.

### RDS, Aurora, and Database Architecture

Watch:

- 90. RDS Custom for Oracle and Microsoft SQL Server
- 93. Amazon Aurora - Advanced Concepts
- 94. RDS & Aurora - Backup and Monitoring
- 96. RDS Proxy
- Quiz 6: RDS, Aurora, & ElastiCache Quiz

Focus:

- Aurora replicas, global database, serverless, backup, restore, and failover.
- RDS Proxy for connection pooling and Lambda/database integration.
- When RDS Custom matters.

### Route 53 and Hybrid DNS

Watch:

- 112. Route 53 - Health Checks Hands On
- 115. Routing Policy - Geoproximity
- 116. Routing Policy - IP-based
- 119. Route 53 Resolvers & Hybrid DNS
- Quiz 7: Route 53 Quiz

Only skim if rusty:

- 101. What is a DNS ?
- 103. Route 53 - Registering a domain
- 104. Route 53 - Creating our first records

Focus:

- Latency, weighted, failover, geolocation, geoproximity, and IP-based routing.
- Resolver endpoints and hybrid DNS.
- Health checks and DNS failover limitations.

### EC2 and EBS Basics That Still Matter

Watch or skim depending on memory:

- 34. EC2 Instance Types Basics
- 35. Security Groups & Classic Ports Overview
- 36. Security Groups Hands On
- 43. EC2 Instance Roles Demo
- 56. EBS Overview
- 58. EBS Snapshots
- 60. AMI Overview
- 63. EBS Volume Types
- 64. EBS Multi-Attach
- Quiz 2: EC2 Fundamentals Quiz
- Quiz 3: EC2 SAA Level Quiz
- Quiz 4: EC2 Data Management Quiz

Focus:

- Instance families and purchasing options.
- Instance roles instead of long-lived credentials.
- EBS gp3/io2/st1/sc1, snapshots, AMIs, and Multi-Attach.
- Placement groups and ENIs if you see exam questions on low latency or failover patterns.

## Priority 3: Skim Only If Rusty

These are either basic, mostly hands-on refreshes, or low-yield compared with the items above.

- Section 3: AWS Cloud overview, console UI update, console tour.
- Basic IAM hands-on lectures such as users/groups, simultaneous sign-in, and MFA hands-on if you already know them.
- Basic EC2 launch, SSH, EC2 Instance Connect, user data, and security group hands-on if you can already do them.
- Basic S3 bucket, website hosting, versioning, and storage class hands-on if you already know them.
- Basic Route 53 registration and first records if DNS fundamentals are still fresh.
- Snow Family overview and hands-on unless you want to refresh migration appliance use cases.
- Machine Learning Section 23: watch at 1.5x or skim for service recognition unless practice questions show weakness.
- Congratulations/certification path lectures unless you want the administrative context.

## Machine Learning Recommendation

Section 23 is completely unwatched:

- 252. Rekognition Overview
- 253. Transcribe Overview
- 254. Polly Overview
- 255. Translate Overview
- 256. Lex + Connect Overview
- 257. Comprehend Overview
- 258. Comprehend Medical Overview
- 259. SageMaker AI Overview
- 260. Kendra Overview
- 261. Personalize Overview
- 262. Textract Overview
- 263. Machine Learning Summary
- Quiz 20: Machine Learning Quiz

Do not go deep here for SAA unless the course or practice exam is highlighting gaps. For SAA, the main requirement is usually recognizing the right managed service for a scenario. Watch this section at 1.5x, take the quiz, and move on.

## Practice and Validation

Do every section quiz for sections where you watch lectures. Also do:

- Quiz 8: Classic Solutions Architecture Discussions Quiz
- Quiz 17: Serverless Solutions Architecture Discussions Quiz
- Quiz 18: Databases in AWS Quiz
- Quiz 26: More Solution Architectures Quiz
- Quiz 28: WhitePapers & Architectures Quiz
- Practice Test 1: Practice Exam - AWS Certified Solutions Architect Associate

Use practice results to decide whether to go back into Priority 2 basics. If you score weakly in a domain, do not just rewatch videos. Write one architecture note per missed topic:

- What problem does this service solve?
- When would I choose it?
- When would I not choose it?
- What are the cost, security, scaling, and failure-mode tradeoffs?

## Suggested Study Order

1. IAM advanced, Organizations, Control Tower, security services.
2. Networking refresh: endpoints, hybrid connectivity, Network Firewall, IPv6.
3. Containers: ECS, Fargate, ECR, EKS.
4. Serverless and event-driven: Lambda concurrency, EventBridge, integration patterns.
5. S3, storage, backup, and DR/migration.
6. Data analytics and streaming.
7. Cost optimization and operational services.
8. RDS/Aurora and Route 53 refresh.
9. Machine learning service recognition.
10. Quizzes, whitepapers/architecture review, and the full practice exam.

## Extra Notes

- If your goal is practical architecture and interviews, the most important missing areas are multi-account identity, private networking, containers, event-driven design, cost optimization, backup/restore, and service tradeoffs.
- If your goal is passing the exam, do not skip quizzes. The exam is scenario-driven, and the fastest way to find gaps is practice questions.
- Keep your 2020 notes, but treat them as a base layer. The main change is not that EC2, S3, RDS, VPC, and IAM disappeared; it is that current questions expect more managed-service composition, more cost awareness, more multi-account/security design, and more serverless/container/event-driven patterns.
- For hands-on labs, prefer small disposable labs and tear resources down immediately. The cost-sensitive topics are now part of the knowledge you are trying to refresh.
