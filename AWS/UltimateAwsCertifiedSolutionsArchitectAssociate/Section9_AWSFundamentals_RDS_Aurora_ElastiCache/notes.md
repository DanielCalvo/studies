- Maybe try the IAM auth short-lived credential thingy on RDS

### 65. RDS Overview
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
- Automated provisioning and OS patching
- Continuous backup and restore functionalities to a specific timestamp (point in time restore)
- Monitoring dashboards
- Read replicas for improved read performance
- Multi AZ set up for DR (disaster recovery)
- Maintenance Windows for upgrades
- Scaling capabilities (both vertical and horizontal)
- Storage is backed by EBS (gp2 or io1)

- Drawback: You can't ssh into your RDS instances :(

#### RDS Backups
- Backups are automatically enabled in RDS
- Automated backups
- Daily full snapshot of the DB (done during the maintenance window)
- Transaction logs are backed up by RDS every 5 minutes
    - These two things together give you the ability to restore to any point in time (from oldest backup to 5 minutes ago)
    - 7 days retention (can be increased to 35 days)

#### DB Snapshops
- Manually triggered by the user
- Retention is for as long as you want

### 66. RDS Read replicas vs Multi AZ

#### Read replicas for read scalability
- When you first set up RDS, your app talks to 1 db instances, all reads and writes go there
- You can create read replicas! Up to 5 read replicas, within or across AZ or cross region
- Replication is asynchronous between the master and a read replica
- Replicas can be promoted to their own DB
- Applications must update the connection to leverage read replicas

#### RDS Read Replicas - Use cases
- Imagine you have a normal production database in RDS that is taking normal load
- But then a new team comes in and says "We need to generate some reports from the data we have"
- If you run those queries against the main (and only) database, it will slow down the db and possibly degrade performance
- To fix this as a solutions architect, you create a new read replica to run the new workload there! 
- Async replication happens from the master RDS DB instance to the RDS DB read replica instance 
- Your reports can then be generated from the read replica and your production application is unaffected
- Remember: Read replicas are only used for SELECT (read) statements, **not** INSERT, UPDATE, or DELETE ones

#### RDS Read Replicas - Network Cost 
- In AWS there is a network cost when data goes from one AZ to another
- Replicating data across AZs on read replicas is going to incur costs
- If you want to reduce costs, have the read replica be on the same AZ

#### RDS Multi AZ (Disaster recovery)
- Your master database can have a synchronous replication to a RDS DB instance in stand by mode in another AZ
- You get one DNS name for these two databases, and in case there is a problem with the master, there is an automatic failover to the stand by database
    - This increases availability
    - Failover happens in case of loss of AZ, loss of network, instance or storage failure
    - No manual intervention required
    - Not used for scaling! The stand by database is just for stand by, no one can read or write to it
    - Possible question: Is there the possibility to have the read replicas set up as multi AZ disaster recovery? Yes! 

### 67. RDS Hands on
- Standard create
- MySQL Community 5.7
- Free tier
- t2.micro
- No multi AZ with stand by
- Default VPC
- Choose to create new SG. Later make sure you can connect with your internet address (ipv4 or ipv6) through this SG. 
- Added deletion protection
- All else were left as defaults
- Publicly accessible

#### Further notes
- Aurora is not available as part of the free tier :(
- Most popular at the exam: MySQL
- Can't change the VPC once db is created :o

### 68. RDS Encryption + Security
- Popular topic at the exam!
- At rest encryption
    - It's possible to encrypt the master & read replicas with AWS KMS AES-256 encryption
    - Encryption has to be defined at launch time
    - If the master is not encrypted, the read replicas cannot be encrypted! (common exam question)
    - You can also enable Transparent Data Encryption (TDE) for Oracle and SQL Server
- In flight encryption 
    - SSL certificates to encrypt data to RDS in flight
    - For this you provide SSL options with a trust certificate when connecting to the database
    - To enforce SSL:
        - PostgreSQL: `rds.force_ssl = 1` in the AWS RDS Console (Parameter groups)
        - MySQL: Within the DB: `GRANT USAGE ON *.* TO 'mysqluser@'%' REQUIRE SSL;` <- After you run this statement, no user will be able to connect to MySQL unless on SSL
        
#### RDS Encryption Operations
- Encrypting RDS backups
    - Snapshots of un-encrypted RDS backups are un-encrypted
    - Snapshots of encrypted RDS databases are encrypted
    - You can however, copy a snapshot that is unencrypted into an encrypted one 
- To encrypt and un-encrypted RDS database
    - Create a snapshot of the un-encrypted database (which will be un-encrypted)
    - Copy the snapshot and enable encryption for the snapshot
    - Restore the database from the encrypted snapshot 
    - And now we have an encrypted database!
    - Now migrate applications to the new database and delete the old one
    
- To connect using SSL
    - Provide the SSL Trust certificate (can be downloaded from AWS)
    - Provide SSL options when connecting to database
    - Remember: Connecting using SSL doesn't mean SSL is enforced, only that it is available
    - But down the road you should know how to enforce SSL

#### RDS Security - Network & IAM
- Network Security
    - RDS Databases are usually deployed within a public subnet, not a public one
    - RDS Security works by leveraging security groups (the same concept for EC2 instances), it controls which IP/SG can communicate with RDS
- Access Management
    - IAM Policies help control who can manage AWS RDS (through the RDS API)
    - Traditional username and password can be used to login to the database
    - IAM-based authentication can be used to login into RDS MySQL & PostgreSQL

#### RDS - IAM Authentication
- IAM database authentication works with MySQ Land PostgreSQL
- You don't need a password, just an authentication token obtained through IAM & RDS API calls
- Auth token has a lifetime of 15 minutes
- Benefits:
    - Network in/out must be encrypted using SSL
    - IAM is used to centrally manage users instead of the DB
    - You can leverage IAM roles and EC2 instance profiles for easy integration (?) 

#### RDS Securuty - Summary
- Encryption at rest:
    - Is done only when you first create the DB instance
    - Or unencrypted DB > snapshot > copy snapshot as encrypted > create DB from snapshot 
- Your responsibility
    - Check the ports/IPs/SG inbound rules in DB's SGs 
    - In-database user creation and permissions (or management of those through IAM)
    - Creating a database with or without public access
    - Ensure parameter groups or DB is configured to only allow SSL connections
- AWS Responsability
    - Make sure you don't have SSH access
    - Make sure you don't have to perform manual db patching
    - Make sure you don't have to do OS patching
    - You don't have a way to audit the underlying instance

### 69. Aurora overview
- No deep knowledge necessary, but a high level view of Aurora is required
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
- If something happens with the master, failover happens in less than 30 seconds
- You can have the master + 15 aurora read replicas to serve reads
- Any of these read replicas can become the master in case of failure
- Diagram: One master, multiple replicas, storage: replicated, self healing, auto expanding
- Cross region replication is supported

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
- There's the possibility to authenticate using IAM tokens (same method as RDS)
- You're responsible for protecting the instance with SGs
- You can't SSH into your instance

#### Aurora serverless
- Automated database instantiation and auto-scaling based un actual usage
- Good for infrequent, intermittent or unpredictable workloads
- No capacity planning required
- Pay per second, can be more cost-effective

- Diagram:
    - There's a proxy fleet managed by aurora that the client will connect to
    - If load increases, more DBs are created on the backend, and if there's less load, DBs are scaled in

#### Global Aurora
- There are two ways to have Global Aurora across multiple regions 
- Aurora Cross Region Read Replicas
    - Useful to disaster recovery 
    - Very simple to put in place
- Aurora Global Database (recommended)
    - 1 primary region where all the reads and writes happen
    - Up to 5 secondary (read-only) regions, with replication lag of less than 1 second
    - Up to 16 read replicas per secondary region can exist
    - Helps decreasing latency
    - Promoting another region (for disaster recovery) has a RTO (Recovery Time Objective) of < 1 minute
    - Example:
        - us-east-1 > Primary region > apps can read and write
        - eu-west-1 > Secondary region > apps can only read

### 70. Aurora Hands On
- You choose Aurora with either MySQL or PostgreSQL compatibility 
- You can chose regional or global database locations. Neat
- You can set it up with
    - One writer and multiple readers (+1 popularity at exam)
    - One writer and multiple readers with parallel query
    - Multiple writers
    - Serverless (+1 popularity at exam)
- Instance size can be selected, all cool with that
- By default Multi-AZ deployment is enabled with the option: "Create an Aurora Replica or Reader node in a different AZ (recommended for scaled availability)" set to yes
- You can choose VPC, subnet, if public accessible etc
- There are "db cluster parameter groups" and "db parameter group" settings under Advanced which look cool but are not in the scope of the exam :(
- There are also settings for: Backup, Encryption, Backtrack (allows you to go back in time) Monitoring (with high granularity), Log exports and Deletion protection
- Important from an exam perspective: Multi-AZ, and one writer and multiple readers or serverless

#### Hands on continued
- Awesome! Upon creation I can see the reader instance and the writer instance on different AZs
- Remember to always use the writer and reader endpoints on the parent DB on the UI, not the children database themselves
- Oh dayum you can add read replica autoscaling with a CPU percentage usage just like an ASG, with a min and max capacity too

### 71. ElastiCache overview
- In the same way that RDS is for managed relational databases...
- ElastiCache is for managed Redis or Memcached! 
- Caches are in-memory databases with high performance and low latency
- Helps reduce load off databases by caching data
- Helps make your application stateless
- Write scaling using sharding
- Read scaling using read replicas
- Multi AZ with failover capabilities
- AWS takes care of maintenance, pathces, optimizations, set up, configuration, monitoring, failure recover and backups, just like with RDS

#### ElastiCache solution architecture - DB Cache
- Application queries ElastiCache, if data is not available there, get data from RDS and store it in ElastiCache
- When you look for cached data in ElastiCache and the data is there, it's called a cache hit
- The opposite of that is a cache miss, and in that case our app goes ahead and queries RDS for the data, and then write to cache
- Caches help relieve the read load from RDS
- Cache has to have some validation strategy to make sure only the most current data is in your cache (to be done at application level though)

#### Elasticache Solution architecture - User Session store
- User logs in into the application
- Application writes the session data into ElastiCache
- User hits another instance running the application
- Application retrieves the session from ElastiCache
- The user session store is stored in ElastiCache and shared across instances. Cool!
- Common architecture pattern with ElastiCache. Certainly beats using files over EFS...

#### ElastiCache - Redis vs Memcached
- Redis
    - Multi AZ feature with Auto-Failover
    - You can also have read replicas to scale reads and have high availability
    - You can have Data Durability using AOF persistence (?)
    - There are backup and restore features
    - Redis can have a primary instance with a replica?
- Memcached
    - Uses multiple nodes for partitioning of data (sharding)
    - Non persistent
    - No backup and restore
    - Multi threaded architecture
    - Memcaching is all about sharding. A part of the cache will be in one shard, another part will be in another shard

### 72. ElastiCache Hands On
- Cluster engine: Can be either redis or memcached
    - Let's go with Redis!
- Some settings you can set up are:
    - Name, description, port, node size, number of replicas...
    - Subnet, VPC, SG, encryption at rest, encryption in transit
    - Can have multi-az with failover if you have multiple replicas

### 73. ElastiCache for the Solution Architect

#### ElastiCache - Cache Security
- All caches in ElastiCache:
    - Support SSL in flight encryption
    - Do not support IAM authentication
    - IAM policies on ElastiCache are only used for AWS API-level security
- Redis has authentication available and it's called **Redis Auth*
    - You can set a password/token when you create a Redis cluster
    - This is an extra level of security for your cache (on top of security groups)
    - By default Redis doesn't have any sort of authentication
- Memcached:
    - Supports SASL-based authentication (advanced, not in the exam)

#### ElastiCache for Solutions Architects
- Patterns for ElastiCache
    - Lazy loading: All the read data is cached, data can become stale in cache
    - Write Through: Adds or updates data in the cache when written to a DB (no stale data)
    - Session Store: Store temporary session data in a cache (using TTL features)
- Quote: There are only two hard things in computer science: Cache invalidation and naming things. (Caching is a hard problem!)

- Quick explanation of Lazy loading:
    - If cache hit, great!
    - If cache miss, read data from DB, then write it to cache

### Quiz
14/15
- Don't forget IAM auth is done through the UI

### 74. [SAA-C02] List of ports to be familiar with
- Important ports:
    - FTP: 21
    - SSH: 22
    - SFTP: 22 (same as SSH)
    - HTTP: 80
    - HTTPS: 443
- RDS Databases ports:
    - PostgreSQL: 5432
    - MySQL: 3306
    - Oracle RDS: 1521
    - MSSQL Server: 1433
    - MariaDB: 3306 (same as MySQL)
    - Aurora: 5432 (if PostgreSQL compatible) or 3306 (if MySQL compatible)