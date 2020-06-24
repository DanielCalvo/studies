### 136. Snowball Overview
- Physical data transport solution that helps moving TBs or PBs of data in or out of AWS
- Alternative over moving data over the network (and paying network fees)
- Good for transfering very large amounts of data from on prem to cloud!
- Secure, tamper resistant, uses KMS 256 bit encryption
- Tracking using SNS and text messages, E-ink shipping label
- Pay per data transfer job 
- Use cases: Large data cloud migrations, DC decomission, disaster recovery   
- If it takes more than a week to transfer over the network, use Snowball devices 

#### Snowball process
1. Request snowball devices from the AWS console for delivery 
2. Install the snowball client on your servers
3. Connect the snowball to your servers and copy file using the client
4. Ship back the device when you're done (goes to the right AWS facility)
5. Data will be loaded into an S3 bucket
6. Snowball is completely wiped 
7. Tracking is done using SNS, text messages and the AWS console

#### Snowball edge
- Snowball edge adds computational capability to the device
- 100TB capacity with either:
    - Storage optimized: 24 vCPU
    - Compute optimized: 52 vCPU & optional GPU
- Supports a custom EC2 AMI so you can perform processing on the go :o
- Supports custom Lambda functions
- Very useful to pre-process data while moving
- Use case: Data migration, image collation, IoT capture, machine learning

#### AWS Snowmobile
- Oh woah it's a truck
- Transfer exabytes of data 
- Each snowmobile has 100PB of capacity (use multiple in parallel)
- Petter than Snowball if you transfer more than 10PB

#### Solution Architecture: Snowball into Glacier
- Snowball cannot import to Glacier directly 
- You have to use Amazon S3 first, and then use a S3 lifecycle policy to transition that data into Glacier
- Amazon S3 is the only place that Snowball can drop it's data
- Snowball > import > S3 > S3 lifecycle policy > Amazon Glacier 

### 137. Snowball Hands on
- There a Snowball page and you can create a job!
- You can choose the Snowball type you want (compute or storage optimized) and define when you can get updates along the way

### 138. Storage Gateway Overview
- AWS is pushing for "hybrid cloud"
    - Part of your infrastructure is on the cloud
    - Part of your infrastructure is on premise
- This can be due to 
    - Long cloud migrations
    - Security requirements
    - Compliance requirements
    - IT strategy    
- S3 is a proprietary storage technology (unlike EFS/NFS) so how do you expose the S3 data to on-premise?
- That's the whole idea behind AWS Storage Gateway!

#### AWS Storage Cloud Native Options
- Block
    - Amazon EBS
    - EC2 Instance store
- File
    - Amazon EFS
- Object
    - S3
    - Glacier

#### AWS Storage Gateway
- You want to bring on premise data into S3 or bridge it
- Bridge between on-premise data and cloud data in S3
- Use cases: disaster recovery, backup & restore, tiered storage
- 3 types of storage gateway (and you need to know them for the exam!)
    - File Gateway (allows you to view files on a local filesystem but it will be backed by S3)
    - Volume Gateway (to do the exact same thing but with volumes)
    - Tape Gateway (for backups and recovery)
- These go through Storage Gateway, and then they go to EBS, S3 or Glacier behind the scenes
- You need to know when to use storage gateway, and the difference between the 3 storage gateway types

#### File Gateway
- Configured S3 buckets accessible using the NFS and SMB protocol
- Supports S3 standard, S3 IA, S3 One Zone IA
- Bucket access using IAM roles for each File Gateway
- Most recently used data is cached in the file gateway
- Since it's NFS or SMB, can be mounted on many servers
- Application server > NFS > File Gateway > HTTPS > S3/Glacier

#### Volume Gateway
- Block storage!
- What is usually mentioned on the exam: "Block storage using iSCSI protocol backed by S3" 
- Backed by EBS snapshots which can help restore on-premise volumes
- There are two options for this:
    - Cached volumes: low latency access to most data 
    - Stored volumes: entire dataset is on premises, scheduled backups to S3 
- Application server > iSCSI > Volume Gateway > Storage Gateway bucket in Amazon S3 > Amazon EBS snapshots

#### Tape Gateway
- Some companies still have backup processes using physical tapes
- With Tape Gateway, companies use the same process but in the cloud
- Virtual Tape library (VTL) backed by Amazon S3 and Glacier
- Back up data using existing tape-based processes (and SCSI interface)
- Works with leading backup software vendors
- Backup server > iSCSI > Tape Gateway > HTTPS > Virtual Tapes stores in Amazon S3 > Archived tapes stored in Amazon Glacier 

#### AWS Storage Gateway Summary
- Exam tip: Read the question well, it will hint at which gateway to use
- On premise data to the cloud > Think Storage Gateway
- More detailed question aobout "File Access / NFS" > Think File Gateway (backed by S3)
- Volume / Block Storage / iSCSI > Volume Gateway (backed by S3 with EBS snapshots)
- VTL Tape solution / Backup with iSCSI > Tape Gateway (backed by S3 and Glacier)

### 139. Storage Gateway Hands on
- Get started
- Choose gateway type
- Select host platform (ESXi, MS Hyper V, Amazon EC2)

### 140. [SAA-C02] Amazon FSx - Overview

#### Amazon FSx for Windows (File Server)
- New file storage!
- EFS is a shared POSIX system for Linux systems
- You cannot use EFS with your Windows Servers :(
- FSx for Windows is a fully managed Windows file system share drive
- Supports SMB and Windows NTFS
- Supports MS Active Directory integration, ACLs and user quotas
- Built on SSD, can scale up to 10s of GB/s, millions of IOPS, 100s PB of data
- Can also be accessed from your on-premise infrastructure
- Can be configured to be Multi-AZ (high availability)
- Data is backed up daily to S3

#### Amazon FSx for Lustre
- Lustre is a type of parallel distributed filesystem for large-scale computing
- The term Lustre is derived from "Linux" and "Cluster"
- Lustre is used for Machine learning and High Performance Computing (HPC)
    - Anytime HPC and filesystem are on the same phrase on the exam it usually means Lustre
- Also used with Video Processing, Financial Modeling, Electronic Design Automation
- Scales up to 100s of GB/s, millions of IOPS, sub-ms latencies, Meant for high performance (aka HPC)
- Seamless integration with S3
    - You can "read your S3 as a filesystem" through FSx 
    - Can write the output of the computations back to S3 (through FSx)
    - So in other ways this is also a way to expose your S3 as a filesystem to your Linux instances
- Can also be used from on-premise servers

#### Recap
- FSx for Windows: Distributed filesystem for your windows instances
- FSx for Lustre: For linux, it's a cluster, for high performance, high IOPS, high throughput, low latency, that can be used with HPC, and with integration with S3 on the backend

### 141. [SAA-C02] Amazon FSx - Hands On

#### For Windows
- FSx in the console
- Chose the Windows option
- Can select Multi AZ or Single AZ
- Specify capacity
- Specify the throughput you want
- Need to specify a VPC and a SG
- Prefferred subnet and standy subnet (Multi AZ so you can have failover to the standard subnet)
- Windows Authentication can be through AWS Managed MS AD or Self managed AD

#### For Lustre
- Have to chose the size of the FS, but the throughput capacity is based on the size for this one
- You can have petabyte-sized filesystems in there :o
- Network settings: Same as before, VPC and SG, but not multi AZ
- Automatically encrypted at rest
- You can integrate with S3
- You can chose the when you want the maintenance window to happen too

### 142. [SAA-C02] All AWS Storage Options Compared
- S3: Object storage. Serverless, you don't have to provision capacity ahead of time. Integrates with a bunch of services
- Glacier: Object Archival. When you want to store objects for a long period of time and retrieve them rarely
- EFS: Network File System for Linux instances, POSIX filesystem. Acessible across AZs
- FSx for Windows: NFS for Windows Servers (very similar to EFS but for Windows)
- FSx for Lustre: High Performance Computing Linux file system. HPC, very high IOPS, integrated with S3
- EBS volumes: Network storage for one EC2 instance at at time. Bound to a specific AZ
- Instance Storage: Physical storage for your EC2 instance (high IOPS)
- Storage Gateway: File Gateway, Volume Gateway (cache & stored), Tape Gateway
    - Transporting files from on premise to AWS 
- Snowball / Snowmobile: To move large ammounts of data to the cloud, physically
- Database: For specific workloads, usually with indexing and queryin

### Quiz!
- Pending!