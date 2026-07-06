# 2026 AWS Solutions Architect Associate Refresh Curriculum

Created: 2026-07-05

Purpose: refresh 2020-era AWS Solutions Architect Associate knowledge for interviews and practical architecture work. This is not an exam-cram plan. The goal is to get current on services, patterns, and tradeoffs that became more important since the SAA-C02 period.

## How to Use This

Work through one module per week. For each module:

- Read the suggested documentation first.
- Build the lab, preferably with Terraform, CloudFormation, CDK, or repeatable CLI commands.
- Write a short architecture note afterwards: what problem it solves, when to use it, when not to use it, cost/security implications, and what you would say in an interview.
- Tear down resources after each lab.

If you only have time for the highest-value path, do modules 1, 2, 3, and 4.

## Week 1: Modern Multi-Account AWS and Identity

### Why This Matters

In 2020, a lot of study material still centered on single-account IAM and cross-account role assumptions. Current AWS architecture discussions assume multi-account environments, centralized identity, guardrails, audit accounts, security accounts, and account vending.

### Learn

- AWS Organizations: accounts, OUs, consolidated billing, SCPs, delegated administration.
- IAM Identity Center, formerly AWS SSO: permission sets, account assignments, identity sources.
- Control Tower: landing zones, Account Factory, preventive/detective/proactive controls.
- IAM Access Analyzer and resource policies.
- Where IAM policies, SCPs, permission boundaries, session policies, and resource policies fit.

### Read

- AWS Organizations overview: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html
- IAM Identity Center overview: https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html
- Control Tower overview: https://docs.aws.amazon.com/controltower/latest/userguide/what-is-control-tower.html
- AWS Well-Architected Framework: https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html

### Lab

Build a small simulated company structure:

- Create an AWS Organization if you have a safe personal/personal-lab setup.
- Create OUs such as `Security`, `Infrastructure`, `Sandbox`, and `Workloads`.
- Create or invite at least one sandbox account.
- Enable IAM Identity Center.
- Create two permission sets: `ReadOnly` and `PowerUserSandbox`.
- Assign yourself access to the sandbox account through IAM Identity Center.
- Add one SCP to a sandbox OU that denies leaving a chosen region or denies disabling CloudTrail.

If creating multiple accounts is too much overhead, do this as a paper lab: draw the org, write example SCPs, and explain which controls live at each layer.

### Interview Output

Be able to answer:

- Why use multiple AWS accounts instead of one account with many VPCs?
- What is the difference between IAM Identity Center and IAM users?
- What can an SCP do, and what can it not do?
- What does Control Tower add on top of Organizations?

## Week 2: Event-Driven and Workflow Architectures

### Why This Matters

SQS, SNS, and Kinesis are still important, but EventBridge and Step Functions are now central architectural tools. Interviews increasingly expect you to know orchestration vs choreography, retry behavior, dead-letter queues, idempotency, and event contracts.

### Learn

- EventBridge event buses, rules, targets, pipes, scheduler.
- Step Functions Standard vs Express workflows.
- Retry, Catch, Choice, Parallel, Map, and callback patterns.
- SQS vs SNS vs EventBridge vs Step Functions.
- Idempotent consumers and poison-message handling.

### Read

- EventBridge overview: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html
- Step Functions overview: https://docs.aws.amazon.com/step-functions/latest/dg/welcome.html
- Step Functions Workshop: https://catalog.workshops.aws/stepfunctions/

### Lab

Build an asynchronous image-processing or document-processing pipeline:

- Upload a file to S3.
- S3 event triggers EventBridge.
- EventBridge routes to a Step Functions workflow.
- Step Functions invokes Lambda functions for validation, metadata extraction, and result persistence.
- Store metadata in DynamoDB.
- Send a completion notification through SNS or write to an SQS queue.
- Add a failure branch and a DLQ.

Stretch goals:

- Add a human approval callback step.
- Add an EventBridge scheduled cleanup task.
- Add structured logs and X-Ray tracing.

### Interview Output

Be able to answer:

- When would you choose EventBridge instead of SNS?
- When would you choose Step Functions instead of putting orchestration logic inside Lambda?
- What is the difference between Standard and Express workflows?
- How do you make event consumers idempotent?

## Week 3: Containers on AWS: ECS, Fargate, and EKS

### Why This Matters

Older notes often treated ECS on EC2 as the main exam topic and Fargate as a side note. Current AWS architecture uses Fargate heavily for serverless containers. EKS is common in interviews even if the role is not Kubernetes-heavy.

### Learn

- ECS clusters, services, tasks, task definitions, task roles, execution roles.
- Fargate vs ECS on EC2.
- ALB integration with ECS services.
- ECR image storage and vulnerability scanning.
- EKS at the conceptual level: managed control plane, node groups, Fargate profiles, IRSA/pod identity.
- When to choose Lambda, ECS/Fargate, ECS/EC2, or EKS.

### Read

- ECS overview: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/Welcome.html
- EKS overview: https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html
- AWS containers workshops: https://catalog.workshops.aws/

### Lab

Containerize and deploy a small web app:

- Build a tiny HTTP service.
- Push the image to ECR.
- Deploy it as an ECS service on Fargate.
- Put it behind an Application Load Balancer.
- Configure service auto scaling.
- Give the task an IAM task role that can read from one DynamoDB table or S3 bucket.
- Send logs to CloudWatch Logs.

Stretch goals:

- Add blue/green deployment with CodeDeploy or use ECS deployment circuit breakers.
- Add private subnets with NAT or VPC endpoints.
- Deploy the same container to EKS using a managed node group or Fargate profile.

### Interview Output

Be able to answer:

- What is the difference between a task role and an execution role?
- Why might you choose ECS/Fargate over Lambda?
- Why might you choose EKS over ECS?
- How does an ALB route traffic to containers with dynamic ports?

## Week 4: Cost-Aware Architecture

### Why This Matters

Cost optimization is now a first-class architecture topic, not an afterthought. For interviews, it is useful to explain tradeoffs across compute, storage, databases, networking, and operational complexity.

### Learn

- Savings Plans vs Reserved Instances vs Spot.
- Compute Optimizer recommendations.
- Cost Explorer, Budgets, Cost and Usage Reports.
- S3 Intelligent-Tiering, Glacier Instant Retrieval, Glacier Flexible Retrieval, Deep Archive, lifecycle rules.
- NAT Gateway, cross-AZ traffic, CloudFront, and data transfer cost traps.
- Right-sizing and managed-service tradeoffs.

### Read

- S3 storage classes: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html
- AWS Well-Architected Cost Optimization pillar: https://docs.aws.amazon.com/wellarchitected/latest/cost-optimization-pillar/welcome.html
- AWS Well-Architected Labs cost optimization: https://www.wellarchitectedlabs.com/cost-optimization/

### Lab

Build a cost review for a hypothetical workload:

- Use your existing notes to design a three-tier app: ALB, ECS/Fargate or EC2, RDS/Aurora, S3, CloudFront.
- Estimate cost with the AWS Pricing Calculator.
- Create three versions: cheapest dev environment, balanced production environment, and high-availability production environment.
- Identify at least five cost levers and their tradeoffs.

Optional hands-on:

- Create an S3 bucket with lifecycle rules.
- Move objects through Standard, Intelligent-Tiering, and Glacier classes.
- Configure a budget and an alert.

### Interview Output

Be able to answer:

- How would you reduce the cost of an over-provisioned EC2/RDS workload?
- When is S3 Intelligent-Tiering a good default?
- What are common hidden AWS networking costs?
- What is the difference between Savings Plans and Reserved Instances?

## Week 5: Modern Storage, Backup, and Data Movement

### Why This Matters

S3, EBS, and EFS are still core, but the practical architecture space now includes centralized backup policies, broader FSx usage, DataSync, Transfer Family, Object Lock, and more nuanced restore requirements.

### Learn

- AWS Backup plans, vaults, lifecycle, cross-account and cross-region backup concepts.
- S3 Object Lock, retention, legal hold, versioning, replication.
- FSx for Windows, Lustre, NetApp ONTAP, and OpenZFS at a high level.
- DataSync vs Storage Gateway vs Transfer Family.
- RPO/RTO tradeoffs for backup and restore, pilot light, warm standby, and active-active.

### Read

- AWS Backup overview: https://docs.aws.amazon.com/aws-backup/latest/devguide/whatisbackup.html
- S3 storage classes: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html
- AWS DataSync docs: https://docs.aws.amazon.com/datasync/latest/userguide/what-is-datasync.html
- AWS Transfer Family docs: https://docs.aws.amazon.com/transfer/latest/userguide/what-is-aws-transfer-family.html

### Lab

Build a backup and restore mini-plan:

- Create an EBS-backed EC2 instance or an RDS instance in a sandbox.
- Configure AWS Backup with a backup plan and vault.
- Create a backup.
- Restore into a new resource.
- Document what changed: endpoint, DNS, IAM, security groups, data freshness, downtime.

Optional paper lab:

- Design a ransomware-resistant S3 setup using versioning, Object Lock, replication, least privilege, and separate backup/admin roles.

### Interview Output

Be able to answer:

- Why use AWS Backup instead of configuring every service backup manually?
- What is the difference between backup, snapshot, replication, and archive?
- How would you design for accidental deletion or ransomware?
- When would you use DataSync instead of Storage Gateway?

## Week 6: Networking Refresh and Private Connectivity

### Why This Matters

Your old networking notes are still valuable, but interviews now expect stronger answers around private connectivity, centralized networking, VPC endpoints, Transit Gateway, PrivateLink, and network security controls.

### Learn

- Gateway endpoints vs interface endpoints.
- PrivateLink as a producer/consumer model.
- Transit Gateway for hub-and-spoke networking.
- VPC sharing with Resource Access Manager.
- Network Firewall vs security groups vs NACLs vs WAF.
- Direct Connect, Site-to-Site VPN, Client VPN, and hybrid routing.
- CloudFront vs Global Accelerator.

### Read

- AWS PrivateLink overview: https://docs.aws.amazon.com/vpc/latest/privatelink/what-is-privatelink.html
- Transit Gateway docs: https://docs.aws.amazon.com/vpc/latest/tgw/what-is-transit-gateway.html
- VPC endpoints docs: https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html
- Network Firewall docs: https://docs.aws.amazon.com/network-firewall/latest/developerguide/what-is-aws-network-firewall.html

### Lab

Build private access to AWS services:

- Create a VPC with private subnets.
- Deploy a private EC2 instance or ECS task.
- Add an S3 gateway endpoint.
- Add interface endpoints for Systems Manager, ECR, CloudWatch Logs, or Secrets Manager.
- Remove public internet dependence where practical.
- Connect to the instance through Systems Manager Session Manager.

Stretch goals:

- Build two VPCs and connect them with Transit Gateway.
- Expose a private service through PrivateLink.
- Add Route 53 private hosted zones.

### Interview Output

Be able to answer:

- Why use VPC endpoints instead of NAT Gateway?
- When is PrivateLink better than VPC peering?
- What does Transit Gateway solve?
- Where do WAF, Network Firewall, security groups, and NACLs each apply?

## Capstone Project: Current AWS Reference Architecture

Build or design one complete architecture that touches most of the refresh areas.

Suggested architecture:

- Static frontend on S3 and CloudFront.
- API through API Gateway or ALB.
- Compute with Lambda for event handlers and ECS/Fargate for a long-running API.
- DynamoDB for metadata.
- S3 for object storage with lifecycle rules and Object Lock if appropriate.
- EventBridge and Step Functions for asynchronous workflows.
- IAM Identity Center and least-privilege roles.
- VPC endpoints for private service access.
- CloudWatch logs, metrics, alarms, X-Ray, and CloudTrail.
- AWS Backup or documented backup/restore process.
- Cost estimate and cost reduction plan.

Deliverables:

- Architecture diagram.
- Short README explaining tradeoffs.
- Threat model: identity, network, data, logging, backup.
- Cost estimate.
- Failure-mode notes: what happens if one AZ fails, a Lambda fails, SQS backs up, RDS becomes unavailable, or an operator deletes data.
- Interview script: a 5-minute explanation of the architecture.

## Topics to Deprioritize

These are still useful to recognize, but they should not dominate the refresh:

- Classic Load Balancer.
- NAT instances.
- EC2 Classic / ClassicLink.
- SWF, except as historical context for why Step Functions is preferred.
- OpsWorks.
- Elastic Transcoder.
- Deep exam trivia around exact service quotas, unless you run into them during labs.

## Quick Mapping to Your Existing Notes

- Security and multi-account: `Section19_IAM_Advanced`, `Section20_AWSSecurity_Encryption_KMS_SSM_ParameterStore_CloudHSM_Shield_WAF`
- Event-driven and workflows: `Section14_SQS_SNS_Kinesis_ActiveMQ`, `Section15_ServerlessOverviewFromASolutionArchitectPerspective`, `Section24_OtherServices`
- Containers: `Section24_OtherServices`
- Cost: `Section25_WhitepapersAndArchitectures_AWS_CSAA`, `Section26_ExamPrep`
- Storage and backup: `Section11_AdvancedAmazonS3AndAthena`, `Section13_AWSStorageExtras`, `Section22_DisasterRecoveryAmdMigrations`
- Networking: `Section21_Networking_VPC`, `Section12_CloudFrontAndAWSGlobalGenerator`

## Suggested Final Review Questions

- Design a multi-account AWS environment for a 30-person company. What accounts do you create and why?
- A team wants to process uploaded documents asynchronously. Which services do you choose and why?
- A containerized API needs predictable latency and private access to S3 and Secrets Manager. How do you deploy it?
- A company is surprised by AWS costs. Where do you look first?
- A workload must survive accidental deletion and ransomware. What controls do you add?
- A private SaaS provider needs to expose an API to hundreds of customer VPCs. Do you use peering, Transit Gateway, or PrivateLink?
- A monolith runs on EC2 today. When would you move it to ECS/Fargate, Lambda, EKS, or leave it on EC2?

## Official Sources Checked

- Current SAA-C03 exam guide: https://d1.awsstatic.com/training-and-certification/docs-sa-assoc/AWS-Certified-Solutions-Architect-Associate_Exam-Guide.pdf
- AWS Certified Solutions Architect Associate page: https://aws.amazon.com/certification/certified-solutions-architect-associate/
- AWS Well-Architected Framework: https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html
- AWS Well-Architected Labs: https://www.wellarchitectedlabs.com/
- AWS Organizations: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html
- IAM Identity Center: https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html
- AWS Control Tower: https://docs.aws.amazon.com/controltower/latest/userguide/what-is-control-tower.html
- Amazon EventBridge: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html
- AWS Step Functions: https://docs.aws.amazon.com/step-functions/latest/dg/welcome.html
- Amazon ECS: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/Welcome.html
- Amazon EKS: https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html
- Amazon S3 storage classes: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html
- AWS Backup: https://docs.aws.amazon.com/aws-backup/latest/devguide/whatisbackup.html
- AWS PrivateLink: https://docs.aws.amazon.com/vpc/latest/privatelink/what-is-privatelink.html
