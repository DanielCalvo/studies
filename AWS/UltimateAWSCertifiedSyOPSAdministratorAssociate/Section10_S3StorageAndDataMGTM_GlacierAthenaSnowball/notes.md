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
### 112. S3 Storage Tiers
### 113. S3 Lifecycle rules 
### 114. S3 Analytics
### 115. Glacier Overview
### 116. Glacier vs S3 Storage Class
### 117. Glacier Vault lock - Hands on
### 118. Snowball overview
### 119. Snowball Hands on
### 120. Storage Gateway for S3
### 121. Storage Gateway for S3 - Hands on
### 122. Athena Overview
### 123. Athena Hands on
### 124. Section Clean up

### Quiz