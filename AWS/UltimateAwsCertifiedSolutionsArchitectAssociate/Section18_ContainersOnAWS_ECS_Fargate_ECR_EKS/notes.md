# Section 18: Containers on AWS: ECS, Fargate, ECR & EKS

### 198. Docker Introduction

### 199. Amazon ECS
- ecs  stands for elastic containers nervous
-  you launched docker containers on it, that is called an ess task
- The standard type is implied to be the the ec2  launch type. this means the for container runs on an easy to instance,. or various instances

However there-s a second launch type called the far gate load launch type
- You launched docker containers on adobes,  but you do not provision the infrastructure, there are no es to instances to manage
- It-s all serverless baby
- You just create the task definitions and as will run ECS tasks for you you based on the cu and ram you need, kind of cool
- To scaled up, just increase the number of tasks

As far as IAM goes there seemed to be two types of roles for ess-
 
the easy to instance profile role:
- Used by the ess agents
-  makes apr calls to these yes service
-  and this seems to be the general

- This seems to be a general role used by the EC2 instances that run your ECS workload to interact with various other amazon webservice is it may need to like,  send logs to cloudwatch, or pull images from cr, or retrieves parameter from parameter store, maybe upload something to as three? and soon iso worth

You then can have roles per cs task-
- Each task can have a specific role
- So obviously you can have different ess servicess/tasks with different roles
- These are defined in the task definition

 what wasn-t clear for me was if the instance profile role and the ess task role are like two ways of accomplishing this same thing, or like if you need to have them together or separately or how do these work it's not very clear at this time.


#### Load balancer integrations
- The default approach here seems to be to use an application load balancer
-  you can also use the network load balancer if you need high throughput or high perform permits
- The classic load balancer is also supported but not recommended since its old

#### data volumes
- If you want to have a shared  file systemFor cs tasks,  you can use efs
-  works on both EC2 and far gate launch types
-  far gate plus efs equals serverless

### 200. Creating ECS Cluster - Hands On
So we went through the user interface and created an cs cluster
### 201. Creating ECS Service - Hands On
- So you create task definition which is a service and then in this task definition you can choose the infrastructure requirements like launch type so if you wanted to run on far gate or on a managed instance
- Once you create a task definition you can launch it as a service

### 202. Amazon ECS - Auto Scaling

### 203. Amazon ECS - Solutions Architectures
- So you can have something like a client uploading something to a street bucket
-  which will send an event to amazon event bread
-  which can have a rule to run an ess task
-  and then this eases to ask can maybe save something  to  dynamo did

 so you can have this whole architecture around ecs

You can alson:
-  have an amazon event bridge running every one hour
-  with a role to run an ECS task inside a cluster
-The task can do whatever you want,Like do some thatch processing of something in S3

You can also have an sqs queue  and services on es gest polling that queue
- You can then plug autoscaling into this so the more messages you have in the queue, the more tasvs are spined up officious

There-s the last thing since if you have an cs task fail
-  you can have that reach out to event bridge
-  which can then trigger a sans notification
-  which can send an email to someone who needs to know that the event failed


### 204. Amazon ECS - Clean Up - Hands On

### 205. Amazon ECR
- Elastic contain registry
-  stores and manages docker images on ars
- You can create your own repo and store image privately or you can use the public gallery
- Fully integrated with TCS, backed by s3
- It supports image vulnerability scanning, versioning, image tags and image life cycle. these are quite handy

####Notes from Dani
-  the public gallery is particularly useful because if you're making calls from within amazon, you get charged per image poll, but you don-t get rate limited
-  so you can use it like any amazon service that you would normally use and just pay for your usage


### 206. Amazon EKS - Overview
- Elastic Kubernetes service
- Managed Kubernetes on arrs
- It-s an alternative to ecs. 
-  it supports ec2  if you want to deploy worker nose or fargate if you want to deploy serverless contaienrs
- Small note from the diagram-  you can have public and private log balancers. I actually never used the private load balancer I should use that sometime

Es supports three types of nodes
- managed note groups
  -  creates and manages nodes for you. 
  - nodes are part of an autoscaling group managed by ekes. 
  - supports on demand or spot instances

- self managed nodes
  -  notes created by you and registered to the EKS cluster managed by an ASG
  -  you can use prebuilt amis
  -  supports only mand or spot instances

-aws fargate
- No maintenance and no nodes required

You can attach data volumes to your cluster by specifying a storage class manifest on your cluster
- this will leverage a container storage interface compliant driver (csi)

You also have support for
- ebs
- efs
- fsx for lustre


### 207. Amazon EKS - Hands On

### Quiz 15: Containers on AWS Quiz


## TODO
- Terraform an ecs clusted and an ecs task and deploy it