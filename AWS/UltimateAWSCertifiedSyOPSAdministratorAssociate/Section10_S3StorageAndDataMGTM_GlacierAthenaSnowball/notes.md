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
### 106. S3 Cross Region Replication
### 107. S3 Pre signed URLs
### 108. CloudFront Overview
### 109. CloudFront with S3 - Hands on
### 110. CloudFront Monitoring
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