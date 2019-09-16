### 126. RDS Overview
To start from the basics:
- RDS: Relational Database Service
- Managed DB service, uses SQL as a query language
- Allows you to create databases in the cloud that are managed by AWS
    - Postgres
    - Oracle
    - MySQL
    - MariaDB
    - MSSQL
    - Aurora

#### Advantages of RDS
- It's a managed service!
- Patching at the OS level
- All sorts of backup functionalities
- Monitoring dashboards
- Read replicas
- Multi AZ set  ups
- Maintenance Windows
- Scaling capabilities

Drawback: You can't ssh into your RDS instances :(

#### Read replicas for read scalability
- When you first set up RDS, your app talks to 1 db instances, all reads and writes go there
- You can create read replicas! Up to 5 read replicas, within or across AZ or cross region
- Replication is asynchronous
- Replicas can be promoted to their own DB
- Applications must update the connection to leverage read replicas

#### RDS Multi AZ (Disaster recovery)
- Your master DB can be synchronized to another DB in another DB called standby
- Only one DNS name is exposed to your app, and that name has an automatically failover from the master to the standby DB.
- No manual intervention required
- Multi AZ is to be used only for disaster recovery! Not used for scaling

#### RDS Backups
- Backups are automatically enabled in RDS
- Automated backups
    - Daily full snapshot of the DB
    - All transaction logs are captured in real time  (you have the ability to restore to any point in time)
    - 7 days retention (can be increased to 35 days)

#### DB Snapshops
- Manually triggered by the user
- Retention is for as long as you want

#### RDS Encryption
Popular topic at the exam!
- Encryption at rest capability with AWS KMS - AES-256 Encryption
- You can use SSL certificates to encrypt data to RDS in flight (over the network)
- To enforce SSL (popular question)
    - PostgreSQL: `rds.force_ssl = 1` in the AWS RDS Console (Parameter groups)
    - MySQL: Within the DB: `GRANT USAGE ON *.* TO 'mysqluser@'%' REQUIRE SSL;` <- AFter you run this statement, no user will be able to connect to MySQL unless on SSL
- To connect using SSL
    - Provide the SSL Trust certificate (can be downloaded from AWS)
    - Provide SSL options when connecting to database
    - Remember: Connecting using SSL doesn't mean SSL is enforced, only that it is available
    - But down the road you should know how to enforce SSL

#### RDS Security
- RDS DAtabases are usually deployed within a public subnet, not a public one
- RDS Security works by leverasging security groups (the same concept for EC2 instances), it controls who can communicate with RDS
- IAM Policies help control who can manage AWS RDS
- Traditional username and password can be used to login to the database
- IAM users can now be used too! (New! For MySQL / Aurora)


#### RDS vs Aurora
- Aurora is proprietary from AWS
- Postgres and MySQL are both aupposrted as the Aurora DBs
- Aurora is "AWS cloud optimized" and AWS claims a 5x performance improvement over MySQL on RDS and 3x over on Postgres on RDS 
- Aurora storage automatically grows in increments of 10GB, up tp 64TB
- Aurora can have 15 replicas while MySQL has 5 and the replication process is faster (sub 10ms replica lag)
- Failover in Aurora is instantenous. HA Native
- Aurora costs more than RDS (20% more) but it's more efficient
- Not many questions in the exam about Aurora, but you should know about it!

### 127. RDS Hands on
- Aurora is not available as part of the free tier :(
- Most popular at the exam: MySQL
)
### 128. RDS Multi AZ vs Read Replicas

### 129. RDS Multi AZ vs Read Replicas Hands On

### 130. RDS Parameter Groups

### 131. RDS Backup vs Snapshots

### 132. RDS Security

### 133. RDS API & Hands on

### 134. RDS & CloudWatch

### 135. RDS Performance insights

### 136. Aurora overview

### 137. Aurora Hands On

### 138. ElastiCache

### 139. ElastiCache Hands On

### 140. Section Clean up

##Quiz!