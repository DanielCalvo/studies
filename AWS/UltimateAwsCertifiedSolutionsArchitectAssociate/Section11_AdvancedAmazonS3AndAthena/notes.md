
### 117. MFA Delete
- MFA (multi factor authentication) forces a user to generate a code on a device (usually a phone or some hardware) before doing important operations on S3
- To use MFA-Delete, enable versioning on the S3 bucket
- You will need MFA to
    - Permanently delete an object version
    - Suspend versioning on the bucket
- You won't need MFA to
    - Enable versioning
    - List deleted versions
- Only the root account can enable/disable MFA delete (admin account not enough)
- MFA delete can only be enabled using the CLI (no way to do it through the console as of this writing)
- The idea appears to be that people can't delete any files
- You will only need MFA to permanently delete an object version or suspend versioning on the buckets

#### Hands on
- Create a bucket with versioning
- Add an MFA device on your account and generate access keys (bleh)
- `aws configure root-vioseven`
- `aws s3 ls --profile root-vioseven`
- To enable MFA delete: 
- `aws s3api put-bucket-versioning --bucket mfa-demo-stephane --versioning-configuration Status=Enabled,MFADelete=Enabled --mfa "arn-of-mfa-device mfa-code" --profile root-datacumulus`
- The idea is that no one will be able to delete any files
- To use MFA delete to actually delete a file you need to do it from the command line

### 118. S3 Default encryption
- The old way to enable default encryption was to use a bucket policy and refuse any HTTP command without the proper headers
    - We did this some lectures ago!
- The new way is to use the "default encryption" option in S3 
- Note: Bucket policies are evaluated before "default encryption"

#### Hands on
- Create bucket > Automatically Encrypt > AES-256
- You can either enable encryption by using the "default encryption" setting or by doing a bucket policy with HTTP headers checks
- Possible exam question: How to enable default encryption?
    - Enable default encryption on the bucket properties
    - Or have a bucket policy that reject empty encryption headers

### 119. S3 Access Logs
- For audit purposes, you want to log all access to S3 buckets
- Any request made to S3, from any account, authorized or denied, will be logged into another S3 bucket
- That data can be analyzed using data analysis tools
- Or something that we'll see later in this section called Amazon Athena
 - S3 log format: https://docs.aws.amazon.com/AmazonS3/latest/dev/LogFormat.html

#### S3 Access Kogs: Warning
- Do not set your logging bucket to be the monitored bucket!
- This will create a logging loop and your bucket will grow in size exponentially

#### Hands on
- Create a bucket with default settings and another bucket with default settings
    - danidani-bucket
    - danidani-access-logs
- On the bucket that you want to enable logs, under properties turn on "Server access logging" and choose a bucket to send logs to
- It takes a while for the logs to show up after you do the requests (took about 2h for the author)

### 120. S3 Replication (Cross Region and same Region) (CRR & SRR)
- Bucket in eu-west-1 > Asynchronous replication > us-east-1
- Must enable versioning in the source and destination
- It's Cross Region Replication (CRR) if they're on different regions
- It's Same Region Replication (SRR) if they're on the same region
- Buckets can be in different accounts
- Copying is asynchronous (Stephane claims: It's quick)
- Must give proper IAM permissions to S3
- CRR Use cases: compliance, lower latency access, replication across accounts
- SRR Use cases: Log aggregation, live replication between production and test accounts

#### S3 Replications - Notes
- After activating, only new objects are replicated (not retroactive)
- For DELETE operations:
    - If you delete without a version ID, it will add a delete marker, will not be  replicated
    - If you delete with a version ID, it deletes in the source, not replicated
    - Any delete operation is not replicated
- There is no "chaning" of replication
    - If a bucket 1 has replication into bucket 2, which has replication into bucket 3
    - Then objects created in bucket 1 are not replicated to bucket 3

#### Hands on
- Create the first bucket (Ireland), enable versioning
- Create another bucket in another region and enable versioning
- Go to first bucket > management > and then replication > add rule > entire bucket 
- Choose destination bucket
- IAM Role: Create a new role and name it "Replication Demo"
- If you actually go on IAM you can find the policy. It allows get and list on the source bucket and replicate actions on the destination bucket
- Neat, replication works!

### 121. S3 Pre signed URLs
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
- `aws s3 presign s3://danidanibucket2/Tracer700.jpg --expires-in 300 --region eu-west-1`
- Neat, my private file is public for 5 minutes!


### 122. S3 Storage Tiers
- Amazon S3 Standard - General purpose
- Amazon S3 Standard-Infrequent access (IA)
    - For when your files are infrequently accessed 
- Amazon S3 One Zone-Infrequent access
- Amazon S3 Reduced Redundancy Storage (deprecated)
- Amazon S3 Intelligent Tiering
    - Moves data between your storage classes intelligently
- Amazon Glacier
    - For archives
- Amazon Glacier Deep Archive
    - For the archives you don't need right away
- Amazon S3 Reduced Redundancy Storage (deprecated - omitted)

#### S3 Standard - General Purpose
- High durability (99.999999999%) (eleven 9s) of objects across multiple AZ
- If you store 10.000.000 (ten million) objects in Amazon S3, you can on average expect to incur a loss of a single object once every 10.000 years
- 99.99% availability over a given year
- S3 is designed to be hosted across multiple AZs so it can sustain 2 concurrent facility failures
- Use cases for S3: Big data analytics, mobile & gaming, content distribution...

#### S3 Standard - Infrequent Access (IA)
- Suitable for data that is less frequently accesses but requires rapid access when needed
- High durability (99.999999999%) of objects across multiple AZs
- 99.9% availability
- Can sustain 2 concurrent facility failures
- Use cases: Store data for disaster recovery, backups...

#### S3 One Zone - Infrequent access (IA)
- Same as IA but data is stored in a single AZ
- Same durability (11 9's) of objects in a single AZ, but data is lost when AZ is destroyed
- Destroyed != Being down. Destroyed means there's a flood, fire, volcano, nuclear explosion, alien invasion,  etc
- 99.95% availability 
- Low latency and high throughput performance
- Supports SSL for data at transit and encryption at rest
- Low cost compared to IA (by 20%)
- Use cases: Storing secondary backup copies of on-premise data, or storing data you can recreate 

#### S3 Intelligent tiering
- Same low latency and high throughput performance of S3 standard
- Small monthly monitoring and auto-tiering fee
- Automatically moves objects between two access tiers (S3 Gen Purpose and S3 IA) based on changing access patterns
- Designed for durability of 99.999999999% of objects across multiple AZs
- Resilient against events that impact an entire AZ
- Designed for 99.9% AZ over a given year

#### Amazon Glacier
- Low cost object storage meant for archiving / backup
- Data is retained for the longer term (10s of years)
- Alternative to on-premise magnetic tape storage
- Average anual durability is still 11 nines
- Cost per storage per month is low (0.004 / GB) + retrieval cost
- Each item in Glacier is called an "Archive" and it can be up to 40TB
- Archives are stored in vaults

#### Amazon Glacier & Glacier Deep Archive
- Amazon Glacier - 3 retrival options:
    - Expedited: 1 to 5 minutes interval
    - Standard: 3 to 5 hours
    - Bulk: 5 to 12 hours
    - Minimum storage duration for Glacier: 90 days
- Amazon Glacier Deep Archive - for long term storage - even cheaper:
    - Standard: 12 hours
    - Bulk: 48 hours
    - Minumum storage duration of 180 days
- Remember the numbers (hours to retrieve and minimum storage durations)
- https://aws.amazon.com/s3/storage-classes/

#### Hands on
- Create a bucket with default settings
- Upload a file
- On the file > properties > storage class. You can change it!
- You can also set the storage class on upload time
- Looks like you set permissions when you upload a file, it's a per file permission thing :thinking:

### 123. [SAA-C02] S3 Lifecycle Policies

#### S3 - Moving objects between storage classes
- You can transition objects between storage classes
- For infrequently acessed objects, move them to STANDARD_IA
- For archive objects you don't need in real time, GLACIER or DEEP_ARCHIVE
- Moving objects can be done manually but it can also be automated using a lifecycle configuration

#### S3 Lifecycle Rules
- You can define transiction actions: They define when objects are transitioned to another storage class
    - Move objects to Standard IA 60 days after creation
    - Move to Glacier for archiving after 6 months
- Expiration actions: Configura objects to expire (delete) after some time
    - Access log files can be set to delete after 365 days 
    - Can also be used to delete old versions of files (if versioning is enabled)
    - Can be used to delete incomplete multi-part uploads
- Rules can be created for a certain prefix (ex s3://mybucket/mp3/*)
- Rules can also be created for certain object tags (ex: Department:Finance)

#### S3 Lifecycle Rules - Scenario I
- The exam is likely to ask you some scenario questions
- Your application on EC2 creates images thumbnails after profile photos are uploaded to Amazon S3. These thumbnails can be easily recreated, and only need to be kept for 45 days. The source images should be able to be immediately retrieved for these 45 days, and afterwards, the user can wait up to 6 hours. How would you design this?
    - S3 source images can be on STANDARD, with a lifecycle configuration to transition them to GLACIER after 45 days
    - S3 thumbnails can be on ONEZONE_IA, with a lifecycle configuration to expire them (delete them) after 45 days

#### S3 Lifecycle Rules - Scenario I
- A rule in your company states that you should be able to recover your deleted S3 objects immediately for 15 days, although this may happen rarely. After this time, and for up to 365 days, deleted objects should be recoverable within 48 hours.
    - You need to enable S3 versioning in order to have object versions, so that “deleted objects” are in fact hidden by a “delete marker” and can be recovered
    - You can transition these “noncurrent versions” of the object to S3_IA
    - You can transition afterwards these “noncurrent versions” to DEEP_ARCHIVE

#### Hands on
- Your bucket > Management > Lifecycle > Add lifecycle rule
    - Which objects: Whole bucket!
    - Transition to Standard IA after 30 days
    - Transition to One-Zone-IA after 60 days
    - Neat!
    - Expire current version of object after 425 days
    - Permanently delete previous versions after 545 days

--- Stopped here!

### 124. [SAA-C02] S3 Performance
#### S3 - Baseline Performance
- Amazon S3 automatically scales to high request rates, with latencies between 100-200ms
- Your application can achieve at least 3.500 PUT/COPY/PASTE/DELETE and 5.500 GET/HEAD requests per second per prefix in a bucket
- There are no limits to the number of prefixes in a bucket
- Example (object path > prefix)
    - bucket/folder1/sub1/file > /folder1/sub1/
    - The prefix is anything between the bucket and the file
    - bucket/folder2/sub2/file > /folder2/sub2/
    - You also get the GET and PUTS for this other prefix, separately from the other one listed above
    - bucket/1/file > /1/
    - bucket/2/file > /2/
- If you spread reads across prefixes evenly, you can achieve various times the 5.5000 GET/HEAD performance limitation

#### S3 - KMS Limitation
- If you use SSE-KMS, you may be impacted by the KMS limits
- When you upload, it calls the GenerateDataKey KMS API
- When you download, it calls the Decrypt KMS API
- KMS has a quota of number of requests per second and those S3 requests count toward it (5500, 10000, 30000 req/s based on region)
- As of today, you cannot request a quota increase for KMS

#### S3 Performance
- Multi Part Upload
    - Recommended for files > 100MB
    - Required for files > 5 GB
- Can help parellelize uploads (speeds up transfers)

#### S3 Transfer Acceleration (upload only)
- Increases transfer speed by transferring a file to an AWS edge location which will forward the data to the S3 bucket in the target region
- Edge locations are more than regions, there are about 200 edge locations today
- Compatible with multi part upload

#### S3 Performance - S3 Byte-range fetches
- Parallelize GETs by requesting specific byte ranges
- Better resilience in case of failures
- Can be used to speed up downloads

- File in S3
    - Part 1
    - Part 2
    - Part N
- All these requests can be made in parallel
- Can be used to retrieve only partial data (for example the head of a file)
- You can then issue a byte-range request for a header (first XX bytes)

### 125. [SAA-C02] S3 Select & Glacier Select
- Retrieve less data using SQL by performing server side filtering
- Can filter by rows & columns (simple SQL statements)
- Less network transfer, less CPU cost client-side

#### Diagram
- So if you have a CSV on S3 you can get it with S3 select and a SQL query and it will return a subset of that CSV

### 126. S3 Event Notifications
- Some events happen in your S3 bucket. This could be object created, object removed, object restored, etc...
- You can create rules and you can apply these only to certain objects (for instance you want to react only to *.jpg)
- Use case: Generate thumbnails of images uploaded to S3
- Possible targets for S3 events notifications:
    - SNS (simple notification service, to send an email for instance)
    - SQS (simple queue service, to add messages into a queue) 
    - Lambda functions (to run some custom code)
- Can create as many S3 events as desired
- S3 event notifications typically deliver events in seconds but sometimes can take a minute or longer
- Caveat: If two writes are made to a single non-versioned object at the same time, it is possible that only a single vent notification will be sent
- If you want to ensure that en event notification is sent for every successful write, you can enable versioning on your bucket
 
#### Hands on
- Create bucket
- Enable versioning
- Scroll down, events, add notification, send to SQS queue
- Let's go and create a SQS queue! 
- Queue:
    - demo-s3-event
    - standard-queue
    - quick-create
    - Copy the ARN and put it on the S3 notification thingy
    - Go back to SQS and under permissions
    - Allow everybody to send message
- And now upload something to S3. Neat, a message was sent!

### 127. Athena Overview
- Serverless service to perform analytics directly against S3 files
- Use SQL language to query the files
- Has a JDBC/ODBC driver
- Charged per query and amount of data scanned
- Supports CSV, JSON, ORC, Avro and Parquet (built on Presto)
- Use cases for Athena: Business Intelligence/analytics/reporting, analyse & query VPC flow logs, ELB logs, CloudTrail Trails, etc
- Exam tip: Analyze data directly on S3, or how can we analyze our ELB logs or ou VPC flog logs > Use Athena

### 128. Athena Hands on
- Ah, Stephane has a bucket full of S3 access logs.
- Stephan creates a database first
- Creates a table
- There's a query already done for you if you google "aws athena s3 access log query"
- You can query the S3 access logs as if they were a database, you can then see how many 404s you got and any other information available in the logs, but using S3

### 129. [SAA-C002] S3 Lock Policies & Glacier Vault Lock
- S3 Object Lock
    - Adopt a WORM (write once read many) model
    - Block an object version deletion for a specific amount of time
    - You have the guarantee that the file was only written once and you don't have deletion or modifications happening
- Glacier Vault Lock
    - Adopt a WORM (write once read many) model
    - You can lock the policy for future edits (can no longer be changed)
    - Helpful for compliance and data retention (ex: So no one will delete that object and you can retrieve it in 7 years for an audit)
- Even admins can't delete files with lock policies!

### Quiz
- Pending