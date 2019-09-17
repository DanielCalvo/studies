## Overal chapter status: Studied carefully taking all the notes. Exam is kinda tricky on multiple az vs read replicas, could use some review

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
- Can't change the VPC once db is created :o

### 128. RDS Multi AZ vs Read Replicas
Popular on the exam: Difference between multi AZ and read replicas

#### Multi AZ
- Your application reads and writes to the master database.
    - This master DB does synchronous replication to a stand by database in another AZ, which is on stand by
    - You cannot directly reach the stand by database or launch connections against it. It's just there as failover
    - Your app when it talks to the database, it talks to one DNS name
    - If a failure happens, you still talk to the same DNS name, but amazon changes the database in the background
    - Multi AZ is not to scale up reads. It's just for fault tolerance
    - It fails over only if the primarey db 

#### RDS Multi AZ in depth
- The failover happens only in the following conditions:
    - The primary db instance fails
    - An AZ has an outage
    - The DB instance server type is changed
    - The operating system of the DB instance is undergoing software patching
    - A manual failover of the DB instance was initiated using reboot with failover
- No failover for DB operations: long-running queries, deadlocks or database corruption errors
- Endpoint is the same after failover (no URL change in application needed)

##### Why use multi AZ in top of fault tolerance
- Lower maintainance impact
- You can use multi AZ to reduce maintenance time, if Amazon does maintenance on your database, it will do on the standy first
- When the standy is upgraded, amazon will then perform changes to the master
- If you have backups, those are created from the standby, not impacting performance of the master db
- Multi AZ is only within a single region. If the whole region goes down, you're still gonna have a bad time

#### RDS Read Replicas
- Improves read performance!
- Writes still happen to a single master
- Reads can happen from any db, from the master or the read replicas
- Replication from master to read replicas is asynchronous

##### RDS Read Replicas in depth
- Read replicas help scaling read traffic
- A Read replica can be promoted as a standalone database (manually)
- Read replicas can be within AZ, Cross AZ or Cross region
- Each read replica has it's own DNS endpoint. To scale, your application needs to be aware of all the DNS endpoints
- You can have read replicas of read replicas
- Read replicas can be multi AZ
- Read replicas can help with disaster recovery by using a cross region read replica
- Read replicas are not supported for Oracle
- Possible exam question: We want to run BI / Analytics reports on our production data but we don't want to impact our production application, how do we do it?
    - Create a read replica!
- Really needed for the exam: To know the difference between read replicas and a multi AZ setup

### 129. RDS Multi AZ vs Read Replicas Hands On
- Created a multi-az database
- You can reboot with failover. It takes a while!
- Actions > Create read replica. You can pick different regions! Also on another AZ
- Saw an example of a read replica working. Cool!

### 130. RDS Parameter Groups
- You can configure the DB engine using Parameter groups
- Dynamic parameters are applied immediately
- Static parameters are applied after instance reboot
- You can modify the parameter group associated with a DB (must reboot to take effect)
- Author encourages you to look at the doc if you're curious about the list of parameters for a given DB

#### Must know parameter for the exam
- PostgreSQL / SQL Server: `rds.force_ssl = 1` to force SSL connections
- For MySQL / MariaDB: `GRANT USAGE ON *.* TO 'mysqluser@'%' REQUIRE SSL;`

#### Hands on
- You can create your own parameter group. Pretty cool! It's a bunch of database settings
- Click on running db > modify > select parameter group > apply now! 
- Instance is going to be in a "modifying" state (rebooting) for a while

### 131. RDS Backup vs Snapshots

#### Backups
- Backups are "continuous" and allow point in time recovery
- Backups happen during maintenance windows
- When you delete a db instance, you can retain the automatic backups
- Backups have a retention period of anything between 0 to 35 days
- If you delete a given database, it's backups will be deleted eventually

#### Snapshots
- Snapshots take IO operations and can stop the database from seconds to minutes
- Snapshots taken on multi AZ DBs don't impact the master, just the stand by
- Snapshots are incremental after the first one (which is full)
- You can copy and share DB snapshots
- Manual snapshots never expire!
- You can take a final snapshot when you delete your DB

#### Quick snapshot hands on
- Database > Maintenance and backups > Take snapshot
- Database > Actions > Restore to point in time 
- You can also copy a snapshot in the same region or in another region!

### 132. RDS Security
Security is a popular exam topic on RDS
- Encryption at rest:
    - Can only be enabled when you create the database instance
- To go from an unencrypted database to an encrypted one
    - unencrypted db > snapshot > copy snapshot as encrypted > create DB from snapshot

#### Your security settings regarding DBs
- Check the ports / IP / Security group inbound rules in the DB's SG to have the connectivity that you want (public/private, what network, etc)
- In database users/permissions are handled by you, not AWS
- Creating a database with or without public access
- Ensure parameters groups or the DB is configured to only allow SSL connections

#### AWS security settings regarding DBs
- No SSH access
- No manual DB patching
- No manual OS patching
- No way to audit the underlying instance
- RDS is a managed service. The exam will gauge your knowledge on what can be done by you and what is done by Amazon

### 133. RDS API & Hands on
There are some RDS API parts we need to know for the exam

- DescribeBDInstances API
    - Helps to get a list of all the DB instances you have deployed including Read Replicas
    - Helps to get the DB version (possible exam question: How do you get the DB version)
- CreateDBSnapshot API
    - Makes a snapshot of a DB (possible exam question: How to automate snapshot?)
    - DescribeEvents API
- DescribeEvents API
    - Returns information about events related to your DB instance
- RebootDBInstanceAPI
    - Helps to initiate a 'forced' failover by rebooting the DB instance
    
`aws rds describe-db-instances`
`aws rds describe-events`
`aws rds reboot-db-instance --db-instance-identifier mydb`

### 134. RDS & CloudWatch
- Cloudwatch metrics associated the RDS, gathered from the hypervisor:
    - DatabaseConnections
    - SwapUsage
    - ReadIOPS / WriteIOPS
    - ReadLatecy / Write Latency
    - ReadThroughPut / WriteThroughPut
    - DiskQueueDepth
    - FreeStorageSpace
- Enhanced monitoring is also available! (gathered from an agent on the DB instance)
    - In depth CPU / Thread / process metrics
    - Over 50 new cpu/mem/IO/sys metrics

#### Hands on!
- RDS > mydb > Modify > Enable enhanced monitoring
- RDS > mydb > monitoring > Drop down "Enhanced Monitoring"

### 135. RDS Performance insights
- Visualize your database performance and analyze any issues that affect it
- With the performance insights dashboard, you can visualize the database load and filter the load:
    - By waits: Find the resource that is the bottleneck (CPI, IO, lock, etc)
    - By SQL statements: Find the SQL statement that is the problem
    - By hosts: Find the server that is using most the DB
    - By Users: Find the user that is using the most the DB
- DBLoad: The number of active sessions for the DB engine
- You can view the SQL queries that are putting the most load in the database

https://aws.amazon.com/blogs/database/tuning-amazon-rds-for-mysql-with-performance-insights/

### 136. Aurora overview
No deep knowledge necessary, but a high level view of Aurora is required
- Proprietary technology from AWS (not open source)
- Postgres and MySQL are both supported as the Aurora database type (that means your DB drivers will work as if Aurora was a Postgres or MySQL database) 
- Aurora is "AWS cloud optimized" and AWS claims a 5x performance improvement over MySQL on RDS and 3x over on Postgres on RDS 
- Aurora storage automatically grows in increments of 10GB, up tp 64TB
- Aurora can have 15 replicas while MySQL has 5 and the replication process is faster (sub 10ms replica lag)
- Failover in Aurora is instantenous. HA Native
- Aurora costs more than RDS (20% more) but it's more efficient

#### Aurora high availability and read scaling
- Aurora stores 6 copies of your data across 3 AZs
    - 4 copies out of 6 are needed for writes
    - 3 copies out of 6 needed for reads
    - Aurora has peer-to-peer data replication to fix corrupted data
    - Storage stored as a"Shared storage volume" which is striped across hundreds of volumes (not manageable by the end user)
    - The shared storage volume has replication + self healing + auto expanding features
- One Aurora instance takes writes (master)
- If something happens with the master, failover happens in last than 30 seconds
- You can have the master + 15 aurora read replicas to serve reads
- Any of these read replicas can become the master in case of failure
- One master, multiple replicas, storage: replicated, self healing, auto expanding

#### Aurora DB cluster
- Aurora provides a writer endpoint, a DNS name, always pointing to the master. The master can fail and failover can take place, but the endpoint remains the same
- You can have a read replicas. The read replicas can have auto scaling!
- There is a reader endpoint! Very similar to the writer endpoint. It helps with connection load balancing. Load balancing happens and connection level, not statement level (popular on the exam)
- Writer endpoint! Reader endpoint! Autoscaling! Shared storage volume that autoexpands!

#### Features of Aurora (more like buzzwords)
- Automatic fail over
- Backup and recovery
- Isolation and security
- Industry compliance
- Push-button scaling
- Automated patching with zero downtime
- Advanced monitoring
- Routine maintenance
- Backtrack: Restore data at any point without using backups

#### Aurora security
- Encryption at rest using KMS
- Automated backups, snapshots and replicas are also encrypted
- Encryption in flight using SSL (same process as MySQL or Postgress)
- You can authenticate using IAM
- You're responsible for protecting the instance with SGs
- You can't SSH

#### Aurora serverless
- No need to choose instance size
- Only supports MySQL and Postgres
- Helpful when you can't predict your workload
- DB cluster will start and autoscale based on CPU usage / connections
- Can migrate from Aurora cluster to Aurora serverless and vice versa (cool!)
- Aurora serverless usage is measured in ACU (Aurora Capacity units)
- Some features of Aurora aren't supported in serverless mode, so make sure to check the documentation if you plan on using it

### 137. Aurora Hands On
- Cool! Multiple instances by default! Creation looks a lot like normal RDS, except with more options
- Remember to always connect to the writer and the reader endpoints, never connect to a single instance of aurora directly!

### 138. ElastiCache overview
- In the same way that RDS is for managed relational databases...
- ElastiCache is for managed Redis or Memcached! 
- Caches are in-memory databases with high performance and low latency
- Helps reduce load off databases by caching data
- Helps make your application stateless
- Write scaling using sharding
- Read scaling using read replicas
- Multi AZ with failover capabilities
- AWS takes care of maintenance, pathces, optimizations, set up, configuratio, monitoring, failure recover and backups, just like with RDS

#### ElastiCache solution architecture
- Application queries ElastiCache, if data is not available there, get data from RDS and store it in ElastiCache
- When you look for cached data in ElastiCache and the data is there, it's called a cache hit
- The opposite of that is a cache miss, and in that case our app goes ahead and queries RDS for the data, and then write to cache
- Caches help relieve the read load from RDS
- Cache has to have some validation strategy to make sure only the most current data is there (to be done at application level though)

#### Elasticache Solution architecture -- User Session store
- User logs in into the application
- Application writes the session data into ElastiCache
- User hits another instance running the application
- Application retrieves the session from ElastiCache
- The user session store is stored in ElastiCache and shared across instances. Cool!

#### Redis overview
- Redis is an in memory key-value store
- Low latency (sub ms)
- Cache can survive reboots by default (called persistence)
- Great to host:
    - User sessions
    - Leaderboards (there's a sorting capability)
    - Distributed states
    - Relieve pressure on databases (such as RDS)
    - Pub/sub capability for messaging (?)
- Multi AZ with automatic failove for disaster recovery if you don't want to lose your cache data
- Supports read replicas

#### Memcached overview
- In memory object store
- Cache doesn't survive reboots
- Use cases
    - Quick retrieval of objects from memory
    - Cache often access objects
- Overall, Redis has largely grown in popularity and has better feature sets than memcached
- Author recommends Redis, though the exam won't ask anything as to which one is better / most popular
    

### 139. ElastiCache Hands On
- You can encrypt elasticache at rest, and in transit! Cool!
- Redis supports cluster mode
- There are also automated backups

### 140. Section Clean up
- Make sure delete snapshots and backups

##Quiz!
- Queries against the main database are causing performance problems. Answer: Set up a read replica
- To enable SSL connections to your Postgre database, use parameter groups with the SSL option enabled
- To have automatic recovery in case of your db becoming irresponsive, use Multi AZ. (One DNS name, automatic failover)
- To have your backups run faster and not impact the production database, enable Multi AZ
- To obtain the RDS database engine version, use the DescribeDBInstances api call
- To find out which SQL statements are slowing down your db, slice the performance by SQL statements on the insights dashboard
- Remember to use the Aurora read endpoint to use load balancing for read replicas and reach the appropriate replicas

