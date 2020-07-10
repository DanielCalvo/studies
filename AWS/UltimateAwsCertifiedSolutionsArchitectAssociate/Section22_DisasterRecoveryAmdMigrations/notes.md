### 244. Disaster Recovery in AWS 
- Any event that has a negative impact on a company’s business continuity or finances is a disaster
- Disaster recovery (DR) is about preparing for and recovering from a disaster
- What kind of disaster recovery can we do?
    - On-premise to On-premise disaster recovery: traditional DR, and very expensive
    - On-premise as a main datacenter and use AWS Cloud for disaster recovery. This is called hybrid recovery
- If all your infra is on AWS you can do disaster recovery from AWS Cloud Region A to AWS Cloud Region B
- We need to define two terms:
    - RPO: Recovery Point Objective
    - RTO: Recovery Time Objective

#### RPO and RTO
- RPO: (basically) how often you run backups. How back in time can you go to recover?
    - If you back up data every hour you can go back in time for an hour, so you would've lost one hour of data
- RTO: When a disaster strikes, the time between RPO and the disaster is going to be a data loss.
    - RTO: When you recover from a disaster. Between the disaster and the RTO is the amount of downtime you have.
- Usually the smaller you want these two to be, the higher the cost

#### Disaster Recovery Strategies
- Let's consider 4 disaster recovery strategies. These are listed from slower RTO to faster RTO:
    - Backup and Restore
    - Pilot Light
    - Warm Standby
    - Hot Site / Multi Site Approach

#### Backup and Restore (High RPO)
- Imagine you have a corporate data center and you keep your backups on S3.
    - You have an AWS Storage Gateway connecting your data center and S3.
    - You can also have lifecycle policies moving your data from S3 to Glacier.
    - Or maybe once a week you're sending a ton of data into Glacier using AWS Snowball about once a week. If you use this method, your RPO is one week!
- If you have AWS Cloud for your services and you use snapshots of EBS, Redshift and RDS
    - If you schedule regular snapshots, your RPO may be much smaller, perhaps 24h, or 1h, depends on how frequently you create your snapshots
    - Restoring snapshots can also take a while!
- Backup and restore is cheap though

#### Disaster Recovery – Pilot Light
- With Pilot Light, a small version of the app is always running in the cloud
- Useful for the critical core (pilot light)
- Very similar to Backup and Restore
- Faster than Backup and Restore as critical systems are already up
- As an example
    - Maybe on your Corporate DC you have a server running and a database
    - Maybe the database has data replication with a replica on RDS in the cloud
    - But maybe your EC2 instances are not critical just yet, so they're not running for the pilot light, but if you, route53 can allows you to failover and you can recreate those instances in the cloud
- Lower RPO, Lower RTO, not very high cost, just an RDS database is running

#### Warm Standby
- A full system is up and running, but at minimum size
- Upon disaster, we can scale this system to production load
- Maybe on your corporate datacenter you have a reverse proxy, an app server and a master DB
- And your route53 is pointing the DNS to your reverse proxy
- In the cloud
    - Data replication from on prem, running
    - EC2 ASG running at minimum capacity
    - And maybe we have an ELB as well, ready to go!
- If disaster strikes, we can use route53 to fail over to the ELB and we can also use the failover to change where the app gets the data from, so we could get the data from the RDS slave
- More costly!

#### Multi Site / Hot Site Approach
- Very low RTO (minutes or seconds) – very expensive
- Full Production Scale is running AWS and On Premise
- You pretty much have a copy of your on premise data center on the cloud at full size, running and ready to go!
- Like warm standby, but full sized!

#### All AWS Multi Region
- Multi-size all cloud would be the same thing, but it could be multi-region 
- But since this is just cloud, we can use Aurora Global with master/slave
- The exam will ask you, based on some scenarios, which DR solution would work best!

#### Disaster Recovery Tips
- Backup
    - EBS Snapshots, RDS automated backups / Snapshots, etc...
    - Regular pushes to S3 / S3 IA / Glacier, with Lifecycle Policy and ross Region Replication
    - From On-Premise: Snowball or Storage Gateway
- High Availability
    - Use Route53 to migrate DNS over from Region to Region is helpful and easy
    - RDS Multi-AZ, ElastiCache Multi-AZ, EFS, S3 are all highly available by default, if enabled
    - Site to Site VPN as a recovery from Direct Connect
- Replication
    - RDS Replication (Cross Region), AWS Aurora + Global Databases
    - Database replication from on-premise to RDS
    - Storage Gateway
- Automation, or how to recover:
    - CloudFormation / Elastic Beanstalk to re-create a whole new environment
    - Recover / Reboot EC2 instances with CloudWatch if alarms fail
    - AWS Lambda functions for customized automations, you can apparently automate many DR features with Lambda
- Chaos
    - Netflix has a “simian-army” randomly terminating EC2

### 245. Database Migration Service (DMS)
- Quickly and securely migrate databases to AWS, resilient, self healing
- The source database remains available during the migration
- Supports:
    - Homogeneous migrations: ex Oracle to Oracle
    - Heterogeneous migrations: ex Microsoft SQL Server to Aurora
- Supports continuous Data Replication using CDC
- To use DMS, you must create an EC2 instance to perform the replication tasks
- Source DB > EC2 instance running DMS > Target DB

#### DMS Sources and Targers
- You don't need to remember all of them but it's good to see them once
- SOURCES:
    - On-Premise and EC2 instances databases: Oracle, MS SQL Server, MySQL, MariaDB, PostgreSQL, MongoDB, SAP, DB2
    - Azure: Azure SQL Database
    - Amazon RDS: all including Aurora
    - Amazon S3
- TARGETS:
    - On-Premise and EC2 instances databases: Oracle, MS SQL Server, MySQL, MariaDB, PostgreSQL, SAP
    - Amazon RDS
    - Amazon Redshift
    - Amazon DynamoDB
    - Amazon S3
    - ElasticSearch Service
    - Kinesis Data Streams
    - DocumentDB

#### AWS Schema Conversion Tool (SCT)
- Convert your Database’s Schema from one engine to another
- Example OLTP: Can migrate from (SQL Server or Oracle) to MySQL, PostgreSQL, Aurora
- Example OLAP: Can migrate from (Teradata or Oracle) to Amazon Redshift
    - Source DB > DMT + SCT > Target DB (different engine)
- You do not need to use SCT if you are migrating the same DB engine
    - Ex: On-Premise PostgreSQL => RDS PostgreSQL (no SCT)
    - The DB engine is still PostgreSQL (RDS is the platform)

### 246. Database Migration Service (DMS) - Hands on

### 247. On-Premises Strategies with AWS
- We have the ability to download Amazon Linux 2 AMI as a VM (.iso format) (that's cool!)
    - You can use this .iso with VMWare, KVM, VirtualBox (Oracle VM), Microsoft Hyper-V, etc
- VM Import / Export
    - Allows you to migrate existing applications into EC2
    - You can also create a DR repository strategy for your on-premise VMs
    - You also can export back the VMs from EC2 to on-premise
- AWS Application Discovery Service
    - A service that allows you to gather information about your on-premise servers to plan a migration
    - Gives you server utilization and dependency mappings
    - You can track your migration with AWS Migration Hub
- AWS Database Migration Service (DMS)
    - Allows you to replicate from On-premise to AWS , AWS to AWS, and AWS On-premise
    - Works with various database technologies (Oracle, MySQL, DynamoDB, etc..)
- AWS Server Migration Service (SMS)
    - An incremental replication of on-premise live servers to AWS (so you can replicate the volumes on AWS)
- Just remember the services at a high level for the exam

### 248. AWS DataSync - Overview
- Service used to move large amount of data from on-premise to AWS
- Can synchronize to: Amazon S3,Amazon EFS, Amazon FSx for Windows
- Can move data from your on-premise NAS or file system via NFS or SMB
- Replication tasks can be scheduled hourly, daily, weekly
- You need to install the DataSync agent on premise to make this work (to connect to your systems)
- It also works from EFS to EFS replication

### 249. Transfering Large Datasets into AWS
- Example: transfer 200 TB of data in the cloud. We have a 100 Mbps internet connection.
- Over the internet / Site-to-Site VPN:
    - Good: Immediate to setup
    - Bad: Will take 200(TB)*1000(GB)*1000(MB)*8(Mb)/100 Mbps = 16,000,000s = 185d
- Over direct connect 1Gbps (we provisioned a 1gb line)
    - Bad: Long for the one-time setup (over a month)
    - Kinda of OK: Will take 200(TB)*1000(GB)*8(Gb)/1 Gbps = 1,600,000s = 18.5d
- Over Snowball:
    - Will take 2 to 3 snowballs in parallel
    - Takes about 1 week for the end-to-end transfer (arrival, loading, arriving at AWS, their setup, etc)
    - If it's a database it can be combined with DMS to transfer the rest of the data afterwards
- For on-going replication / transfers: Site-to-Site VPN or Directo Connect (DX) with DMS or DataSync
- Exam will probably ask you what is the easiest, fastest or most reliable way to send data into AWS if it's a small or large dataset

### Quiz
- Pending