# Section 18: Containers on AWS: ECS, Fargate, ECR & EKS

### 198. Docker Introduction

### 199. Amazon ECS
- ECS stands for Elastic Container Service
- You launch Docker containers on it; that is called an ECS task
- The standard type is implied to be the EC2 launch type. This means the container runs on an EC2 instance or various instances

However, there's a second launch type called the Fargate launch type
- You launch Docker containers on AWS, but you do not provision the infrastructure; there are no EC2 instances to manage
- It's all serverless baby
- You just create the task definitions, and AWS will run ECS tasks for you based on the CPU and RAM you need, kind of cool
- To scale up, just increase the number of tasks

As far as IAM goes, there seem to be two types of roles for ECS:

The EC2 instance profile role:
- Used by the ECS agents
- Makes API calls to the ECS service
- This seems to be a general role used by the EC2 instances that run your ECS workload to interact with various other Amazon Web Services it may need to, like sending logs to CloudWatch, pulling images from ECR, retrieving parameters from Parameter Store, maybe uploading something to S3? And so on and so forth

You then can have roles per ECS task:
- Each task can have a specific role
- So obviously you can have different ECS services/tasks with different roles
- These are defined in the task definition

What wasn't clear for me was if the instance profile role and the ECS task role are like two ways of accomplishing the same thing, or like if you need to have them together or separately, or how these work. It's not very clear at this time.

#### Load balancer integrations
- The default approach here seems to be to use an application load balancer
- You can also use the network load balancer if you need high throughput or high performance
- The classic load balancer is also supported but not recommended since it's old

#### Data volumes
- If you want to have a shared file system for ECS tasks, you can use EFS
- Works on both EC2 and Fargate launch types
- Fargate plus EFS equals serverless

### 200. Creating ECS Cluster - Hands On
So we went through the user interface and created an ECS cluster

### 201. Creating ECS Service - Hands On
- So you create a task definition, which is a service, and then in this task definition you can choose the infrastructure requirements, like launch type, so if you wanted to run on Fargate or on a managed instance
- Once you create a task definition you can launch it as a service

### 202. Amazon ECS - Auto Scaling

### 203. Amazon ECS - Solutions Architectures
- So you can have something like a client uploading something to an S3 bucket
- Which will send an event to Amazon EventBridge
- Which can have a rule to run an ECS task
- And then this ECS task can maybe save something to DynamoDB

So you can have this whole architecture around ECS

You can also:
- Have Amazon EventBridge run every hour
- With a role to run an ECS task inside a cluster
- The task can do whatever you want, like do some batch processing of something in S3

You can also have an SQS queue and services on ECS just polling that queue
- You can then plug autoscaling into this, so the more messages you have in the queue, the more tasks are spun up, obviously

There's the last thing: if you have an ECS task fail
- You can have that reach out to EventBridge
- Which can then trigger an SNS notification
- Which can send an email to someone who needs to know that the event failed

### 204. Amazon ECS - Clean Up - Hands On

### 205. Amazon ECR
- Elastic Container Registry
- Stores and manages Docker images on AWS
- You can create your own repo and store image privately or you can use the public gallery
- Fully integrated with ECS, backed by S3
- It supports image vulnerability scanning, versioning, image tags and image lifecycle. These are quite handy

#### Notes from Dani
- The public gallery is particularly useful because if you're making calls from within Amazon, you get charged per image pull, but you don't get rate limited
- So you can use it like any Amazon service that you would normally use and just pay for your usage

### 206. Amazon EKS - Overview
- Elastic Kubernetes Service
- Managed Kubernetes on AWS
- It's an alternative to ECS
- It supports EC2 if you want to deploy worker nodes or Fargate if you want to deploy serverless containers
- Small note from the diagram: you can have public and private load balancers. I actually never used the private load balancer; I should use that sometime

EKS supports three types of nodes
- Managed node groups
    - Creates and manages nodes for you
    - Nodes are part of an autoscaling group managed by EKS
    - Supports On-Demand or Spot Instances

- Self-managed nodes
    - Nodes created by you and registered to the EKS cluster, managed by an ASG
    - You can use prebuilt AMIs
    - Supports On-Demand or Spot Instances

- AWS Fargate
    - No maintenance and no nodes required

You can attach data volumes to your cluster by specifying a storage class manifest on your cluster
- This will leverage a Container Storage Interface-compliant driver (CSI)

You also have support for
- EBS
- EFS
- FSx for Lustre

### 207. Amazon EKS - Hands On

### Quiz 15: Containers on AWS Quiz

## TODO
- Terraform an ECS cluster and an ECS task and deploy it
