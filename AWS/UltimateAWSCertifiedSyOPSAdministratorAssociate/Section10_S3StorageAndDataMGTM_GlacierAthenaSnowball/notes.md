- Chapter notes: Re-write the description on storage gateway, it's a bit confusing still (maybe consul the AWS docs and copy and paste something from there)
- Maybe plan with athena a bit more later
- Quiz status: Pending

### 100. Section Intro
- Session is intense
- Popular on the exam, some questions might be tricky

### 101. S3 Versioning Advances - For SysOps
- S3 versioning will create a new version each time you change a file 
- That also includes when you encrypt a file
- It's a nice way to get protected against hackers attempting to encrypt our data
- Deleting a file in the S3 bucket just adds a delete marker on the versioning 
- To delete a bucket, you need to remove all the file versions within it

#### Hands on
- S3 > New bucket > Tick versioning
- Upload file > Encrypt the file with AWS-256
- Encryption a file on a bucket that is versioned makes a new file appear
- If you delete a file on a versioned bucket and you look at the version, you will still be able to access older copies of the file
- So you can do `aws s3 ls s3://mybucket` on a versioned bucket which has delete markers on the files in the bucket, so you'll get no files.
- But if you try to `aws s3 ls s3://mybucket` you'll get a "bucket not empty message"
- If you delete all the versions of the file on the UI, then it's actually deleted and then you can remove the bucket on the cmd line

### 102. MFA Delete
- MFA (multi factor authentication) forces a user to generate a code ona device (usually a phone or some hardware) before doing important operations on S3
- To use MFA-Delete, enable versioning on the S3 bucket
- You will need MFA to
    - Permanently delete an object version
    - Suspend versioning on the bucket
- You won't need MFA to
    - Enable versioning
    - List deleted versions
- Only the root account can enable/disable MFA delete (admin account not enough)
- MFA delete can only be enabled using the CLI
- The idea appears to be that people can't delete any files

#### Hands on

### 103. S3 Default encryption
- The old way to enable default encryption was to use a bucket policy and refuse any HTTP command without the proper headers
- The new way is to use the "default encryption" option in S3 
- Note: Bucket policies are evaluated before "default encryption"

#### Hands on
- Create bucket > Automatically Encrypt > AES-256
- You can either enable encryption by using the "default encryption" setting or by doing a bucket policy with HTTP headers checks

### 104. S3 Access Logs
- Possible exam question
- For audit purposes, you want to log all access to S3 buckets
- Any request made to S3, from any account, authorized or denied, will be logged into another S3 bucket
- That data can be analyzed using data analysis tools such as Hive or Amazon Athena
- S3 log format: https://docs.aws.amazon.com/AmazonS3/latest/dev/LogFormat.html

#### Hands on
- Create a bucket with default settings and another bucket with default settings
- On the bucket that you want to enable logs, under properties turn on "Server access logging" and choose a bucket to send logs to
- It takes a while for the logs to show up after you do the requests (took about 2h for the author)

### 105. S3 Policies Hands on
- Let's analyze the policies here: https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html
- Sometimes in the exam they'll show a bucket policy and you have to understand what it does
- Read them and re-read them to be able to understand them! Exam might ask about them

### 106. S3 Cross Region Replication
- Bucket in eu-west-1 > Asynchronous replication > us-east-1
- Must enable versioning in the source and destination
- Buckets must be in different AWS regions
- Buckets can be in different accounts
- Copying is asynchronous (won't happen right away)
- Must give proper IAM permissions to S3
- Use cases: Compliance, lower latency access, replication across accounts

#### Hands on
- Create the first bucket (Ireland), enable versioning
- Create another bucket in another region with versioning
- Go to first bucket > management > and then replication > add rule > entire bucket 
- Choose destination bucket > Create new role on step 3
 
### 107. S3 Pre signed URLs
- You can generate pre-signed URLs using the SDK or CLI
    - For downloads (easy, can use the CLI)
    - For uploads (harder, must use SDK)
- Valid for a default of 3600 seconds, can change timeout with `--expires-in- [TIME_BY_SECONDS]` argument
- Users given a pre-signed URL inherit the permissions of the person who generated the URL for /GET/PUT
- Use cases:
    - Allow only logged-in users to download a premium video on your S3 bucket
    - Allow an ever changing list of users to download files by generating URLs dynamically
    - Allow temporarely a user to upload a file to a precise location in our bucket

#### Hands on
- On the command line: `aws s3 presign help`
- Description
    - Generate a pre-signed URL for an Amazon S3 object. This allows anyone who receives the pre-signed URL to retrieve the S3 object with an HTTP GET request.
    - For sigv4 requests the region needs to be configured explicitly
- `aws configure set default.s3.signature_version s3v4`
- `aws s3 presign s3://yourbucket/yourfile.jpeg --expires-in 300 --region eu-west-1`

### 108. CloudFront Overview
- Content Delivery network (CDN)
- Improves read performance, content is cached at the edge
- About 136 points of presence globally (edge locations)
- Popular with S3, but works with EC2, Load balancing. Exam might ask about this! Remember that it not only applies to S3, but EC2 and load balancing too
- Can protect against network attacks
- Can provide SSL encryption (HTTPS) at the edge using ACM
- CloudFront can use SSL encryption (HTTPS) to talk to your applications
- CloudFront supports RTMP protocol (videos/media)
- Possible exam question: Every time you see "caching of files, efficiently around the world" think CloudFront

### 109. CloudFront with S3 - Hands on
- We'll create an S3 bucket
- We'll create a CloudFront distribution to distribute the content of that bucket 
- We'll create an Origin Access Identity (user of CloudFront acessing the S3 bucket)
- We'll limit the S3 bucket to be accesses only using this identity

#### Hands on
- Create bucket, everything default
- CloudFront > Create distribution > Web 
- Origin Domain name: The bucket you just created, Create new Origin Access Identity, Yes update bucket policy
- Redirect HTTP to HTTPS, Methods: GET, HEAD.
- Leave everything else as default and create distribution

- Uh-oh, we have a Origin Access Identity created
- If you go back to the S3 bucket and check it's policy, you can see that it was updated to include the CloudWatch permissions
- Using the CloudFront URL I got redirected to the S3 bucket :(
- Author says it takes about 3 hours for it to be fixed! :O, says we need to make files public in the meanwhile, it's a DNS propagation thing

### 110. CloudFront Monitoring
- CloudFront access logs: Logs every request made to CLoudFront into a logging S3 bucket

#### CloudFront Reports
- It's possible to generate reports on
    - Cache statistics reports
    - Popular Objects report
    - Top Referrers Report
    - Usage Reports
    - Viewers Report
- These reports are based on the data from the Access logs

#### CloudFront Troubleshooting
- CloudFront caches HTTP 4xx and 5xx status codes returned by S3 (or the origin server)
- 4xx error code means that the user doesn't have access to the underlying bucket (403) or the object the user is requesting is not found (404)
- 5xx error codes indicates Gateway issues

#### Hands on
- CloudFront distributions > The one you created previously > Edit > Bucket for logs
- Go back to S3 and create a bucket for logs, default settings

#### Monitoring
- Tons of reports on cache statistics, really cool  

### 111. S3 Inventory
- Amazon S3 inventory helps manage your storage
- Audit and report on the application and encryption status of your objects
- Use cases
    - Business
    - Compliance
    - Regulatory needs
- You can query all the data using Amazon Athena, Redshift, Presto, Hive, Spark
- You can set up multiple inventories
- Data goes from a source bucket to a target bucket (need to set up policy)

#### Hands on
- Create a bucket with default settings
- Bucket > Management > Inventory
- Inventory takes up to 48h to generate :(
- A bucket policy is applied to the destination bucket

### 112. S3 Storage Tiers
- Not asked in detail at the exam, but you should know how it works
- Amazon S3 Standard - General purpose
- Amazon S3 Standard-Infrequent access (IA)
- Amazon S3 One Zone-Infrequent access
- Amazon S3 Reduced Redundancy Storage (deprecated)
- Amazon S3 Intelligent Tiering (new!)
- Amazon Glacier (there's a dedicated lecture for it)

#### S3 Standard
- High durability  (99.999999999%) of objects across multiple AZ
- If you store 10.000.000 (ten million) objects in Amazon S3, you can on average expect to incur a loss of a single object once every 10.000 years
- 99.99% availability over a given year
- S3 is designed to be hosted across multiple AZs so it can sustain 2 concurrent facility failures
- Use cases for S3: Big data analytics, mobile & gaming, content distribution...

#### S3 Reduced Redundancy Storage (DEPRECATED)
- Designed to provide 99.99% durability
- 99.99% availability of objects over a given year
- Designed to sustein the loss of data in a single facility
- Use cases: Noncritical, reproducible data at lower levels of redundancy than Amazon's S3 standard storage (thumbnails, transcoded media, processed data that can be reproduced)

#### S3 Standard - Infrequent Access (IA)
- Suitable for data that is less frequently accesses but requires rapid access when needed
- High durability (99.999999999%) of objects across multiple AZs
- 99.99% availability
- Low cost compared to Amazon S3 standard, but you get a "retrieval fee" every time you access a file. Maybe use this only when you know that you're going to access your files, say, once a month
- Can sustain 2 concurrent facility failures
- Use cases: Store data for disaster recovery, backups...

#### S3 One Zone - Infrequent access (IA)
- Same as IA but data is stored in a single AZ
- Same durability (11 9's) of objects in a single AZ, but data is lost when AZ is destroyed
- Destroyed != Being down. Destroyed means there's a flood, fire, volcano, nuclear explosion, alien invasion,  etc
- 99.95% availability 
- Low cost compared to IA (by 20%)
- Use cases: Storing secondary backup copies of on-premise data, or storing data you can recreate 

#### S3 Intelligent tiering (new!)
- Probably not at the exam
- Same low latency and high throughput performance of S3 standard
- Small monthly monitoring and auto-tiering fee
- Automatically moves objects between two access tiers based on changing access patterns
- Designed for durability of 99.999999999% of objects across multiple AZs
- Resilient against events that impact an entire AZ
- Designed for 99.9% AZ over a given year

#### Hands on
- Create a bucket with default settings
- Looks like you set permissions when you upload a file, it's a per file permission thing :thinking:

### 113. S3 Lifecycle rules 
- Set of rules to move data between different tiers to save storage cost
    - Example: General purpose > Infrequent Access > Glacier
- Transition actions: Defines when objects are transitioned to another storage class
    - Example: We can choose to move objects to Stndard IA class 60 days after you created them or can move to Glacier for archiving after 6 months
- Expiration actions: Helps to configure objects to expire after a certain time period. S3 deletes expired objects on our behalf
    - Example: Access log files cna be set to delete after a specified period of time
- Can be used to delete incomplete multi-part uploads!

#### Hands on
- Some bucket > Management > Lifecycle rule
- Created a rule to move objects to IA after 30 days and to delete them after 365 days. Neat!

### 114. S3 Analytics - Storage class analysis
- You can setup analytics to help determine when to transition objects from Standard to Standard_IA 
- Does not work for ONEZONE_IA or Glacier
- Report is updated on a daily basis
- Takes about 24h to 48h for it to first start
- This report helps you put together efficient lifecycle rules

#### Hands on
- To enable it: Your bucket > Management > Analytics

### 115. Glacier Overview
- Low cost storage meant for archiving / backup
- Data is retained for the longer term (10s of years)
- Alternative to on-premise tape storage
- Average anual durability is 99.99999999% (11 9's, same as S3)
- Cost per storage per month is low (USD 0.004 per GB) + retrieval cost
- Naming: In S3 we have buckets and objects. In Glacier we have archives. Each item in glacier is called an archive. An item can be as big as 40TB
- Archives are stored in vaults
- Exam tip: "We want to archive data form S3 after xxx days, what do we use?". Use Glacier!

#### Glacier operations
- There are 3 Glacier operations:
    - Upload: Single operation of by parts (MultiPart upload) for larger archives
    - Download: First initiate a retrieval job for the particular archive, Glacier then prepares it for download. User then has a limited time to download the data from a staging server
    - Delete: Use Glacier rest API or AWS SDK by specifying archive ID
- Restore links have an expiry date
- 3 retrieval options (possible exam question):
    - Expedited: (1 to 5 minutes interval): 0.03 per GB and 0.01 per request
    - Standard: (3 to 5 hours): 0.01 per GB and 0.05 per 1000 requests
    - Bulk: (5 to 12 hours): 0.0025 per GB and 0.025 pero 1000 requests 

#### Vaunt policies & Vault lock
- Vault is a collection of archives
- Each Vault has:
    - One vault access policy 
    - One vault lock policy
- Vault policies are written in JSON 
- Vault access policy is similar to bucket policy (restrict user / account permissions)
- Vault lock policy is a policy you lock, for regulatory and compliance requirements (popular exam question)
    - The policy is immutable, it can never be changed (which is why it's called lock).
    - Lock policy example: Forbid deleting an archive if it's older than 1 year old
    - Other example: Implement WORM (write once read many), to prevent any archive once created to be tampered with (popular exam question, How do we implement a WORM policy for glacier? Using a vault lock policy)   

### 116. Glacier vs S3 Storage Class
- On any bucket > right click file > change storage class
- Different retrieval methods have different prices and different restore times

### 117. Glacier Vault lock - Hands on
- Go to the Glacier console, create vault
- You then create a vault policy and lock it with a vault lock, meaning your policy won't be able to be changed
- You're not able to upload files to Glacier through the UI, you need to use the SDK or CLI

### 118. Snowball overview
- Physical data transport solution that helps moving TBs or PBs of data in or out of AWS
- Alternative over moving data over the network (and paying network fees)
- Secure, tamper resistant, uses KMS 256 bit encryption
- Tracking using SNS and text messages, E-ink shipping label
- Pay per data transfer job 
- Use cases: Large data cloud migrations, DC decomission, disaster recovery   
- If it takes more than a week to transfer over the network, use Snowball devices 

#### Snowball process
1. Request snowball devices from the AWS console for delivery 
2. Install the snowball client on your servers
3. Connect the snowball to your servers and copy file using the client
4. Ship back the device when you're done (goes to the righ AWS facility)
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

### 119. Snowball Hands on
- There a Snowball page and you can create a job!

### 120. Storage Gateway for S3
- AWS is pushing for "hybrid cloud"
    - Part of your infrastructure is on the cloud
    - Part of your infrastructure is on premise
- This can be due to 
    - Long cloud migrations
    - Security requirements
    - Compliance requirements
    - IT strategy    
- S3 is a proprietary storage technology (unlike EFS/NFS) so how do you expose the S3 data to on-premise?
- AWS Storage Gateway!

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
- Bridge between on-premise data and cloud data in S3
- Use cases: disaster recovery, backup & restore, tiered storage
- 3 types of storage gateway (and you need to know them for the exam!)
    - File Gateway
    - Volume Gateway
    - Tape Gateway
- These go through Storage Gateway, and then they go to EBS, S3 or Glacier behind the scenes

#### File Gateway
- Configured S3 buckets accessible using the NFS and SMB protocol
- Supports S3 standard, S3 IA, S3 One Zone IA
- Bucket access using IAM roles for each File Gateway
- Most recently used data is cached in the file gateway
- Since it's NFS or SMB, can be mounted on many servers

##### Diagram
- Application server > NFS > File Gateway > HTTPS > S3/Glacier

#### Volume Gateway
- Block storage using iSCSI protocol backed by S3 
- Backed by EBS snapshots which can help restore on-premise volumes
- Cached volumes: low latency access to most data 
- Stored volumes: entire dataset is on premises, scheduled backups to S3 

##### Diagram
- Application server > iSCSI > Volume Gateway > Storage Gateway bucket in Amazon S3 > Amazon EBS snapshots

#### Tape Gateway
- Some companies still have backup processes using physical tapes
- With Tape Gateway, companies use the same process but in the cloud
- Virtual Tape library (VTL) backed by Amazon S3 and Glacier
- Back up data using existing tape-based processes (and SCSI interface)
- Works with leading backup software vendors

##### Diagram
- Backup server > iSCSI > Tape Gateway > HTTPS > Virtual Tapes stores in Amazon S3 > Archived tapes stored in Amazon Glacier 

#### AWS Storage Gateway Summary
- Exam tip: Read the question well, it will hint at which gateway to use
- On premise data to the cloud > Storage Gateway
- File Access / NFS > File Gateway (backed by S3)
- Volume / Block Storage / iSCSI > Volume Gateway (backed by S3 with EBS snapshots)
- VTL Tape solution / Backup with iSCSI > Tape Gateway (backed by S3 and Glacier)

### 121. Storage Gateway for S3 - Hands on
- Storage gateway console!

### 122. Athena Overview
- Serverless service to perform analytics directly against S3 files
- Use SQL language to query the files
- Has a JDBC/ODBC driver
- Charged per query and amount of data scanned
- Supports CSV, JSON, ORC, Avro and Parquet (built on Presto)
- Use cases for Athena: Business Intelligence/analytics/reporting, analyse & query VPC flow logs, ELB logs, CloudTrail Trails, etc
- Exam tip: Analyze data directly on S3 > Use Athena

### 123. Athena Hands on
### 124. Section Clean up

### Quiz