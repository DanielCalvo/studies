### 96. Amazon S3 - Section Introduction
- Amazon S3 is one of the main building blocks of AWS
- It is advertised as "infinitely scaling" storage
- Widely popular and it deserves its own section!
- Many websites use Amazon S3 as a backbone
- Many AWS services use Amazon S3 as an integration as well

### 97. S3 buckets and objects

#### Buckets
- Amazon S3 allows you to store objects (files) in "buckets" (directories)
- Buckets must have a **globally unique name**, across all AWS accounts!
- Buckets are be defined at the region level
- There's a naming convention!
    - No uppercase
    - No underscore
    - 3-63 characters long
    - Not an IP
    - Must start with a lowercase letter or number

#### Objects
- Objects (files) have a key. The key is the **full** path, even if you have folders:
    - s3://mybucket/my_file.txt
    - s3://mybucket/another_folder/my_file.txt
- They key is composed of prefix + object name:
    - s3://mybucket/one_folder/another_folder/my_file.txt 
    - prefix: one_folder/another_folder
    - object name: my_file.txt
- There's no concept of directories within buckets, just very long key names with slashes (the UI tricks you)    
- You just have keys with very long names that contain slashes ("/")

#### Amazon S3 overview - Objects (continued)
- Object values are the contents of the body:
    - Max size is 5TB
    - If uploading more than 5GB, you must use "multi-part upload"
- Your object can have metadata (list of key/value pairs, system or user metadata)
- Your object can have tags (key/value pairs, up to 10, useful for security/lifecycle)
- There's a version ID if versioning is enabled

#### Hands on - Creating a bucket
- S3 does not require a region selection, it's a global service. But when you create a bucket it is attached to a region.
- Creating a bucket
    - Set a bucket name. Bucket names must be globally unique!
    - Select a region
    - Block public access (or not) for that bucket
    - Advanced settings will be looked at after
- Uploading a file
    - Chose file to upload
    - Next: Manage user read/write permissions and access to other AWS accounts
- Storage class
    - Deep dive later
- Encryption
    - Deep dive later
- There's two ways to open a file with S3. With default settings you can open the file through S3, but not with the public link as files are not public by default

### 98. S3 versioning - basics 
 - You can version your files in AWS S3!
 - Versioning is enabled at **bucket level** 
 - When you overwrite a file, it will increment the version of that file and keep many versions of it
 - It's considered a best practice to version your buckets
    - You can protect yourself against unintended deletes and roll back if there are any issues
    - You can easily roll back to any previous version
    - Downside is you use more space
- Any file that is not versioned prior to enabling versioning will have version "null"
- Suspending versioning does not delete the previous versions, it will just make sure that future files will not have a version assigned to it

#### Hands on
- Your bucket > Properties > versioning > enable versioning
- If you upload a file twice, you'll have two versions of it
- Hid versions, deleted a file
- When you delete a versioned file, a delete marker is placed on it
- You can go back to seeing your versions and delete your delete marker, the latest version will then be the non-deleted one and you'll be able to see your file
- When you suspend versioning, it will only affect future objects

### 99. S3 Encryption  
- There are 4 methods for encrypting objects in S3:
    - SSE-S3: Encrypts S3 objects using keys handled & managed by AWS
    - SSE-KMS: Leverages AWS Key Management Service to manage encryption keys
    - SSE-C: When you provide your own keys and Amazon encrypts the data for you
    - Client side encryption: When you encrypt your data clientside.
- It's important to understand which ones are adapted to which situation in the exam!
- Exam wants to know if you can chose the right level of encryption based on the scenario

#### SSE-S3
- Encryption using keys handled & managed by AWS S3
- Object is encrypted server side (SSE means Server Side Encryption) 
- AES-256 encryption type
- To make this work, you must set a header when you send your data to Amazon S3
    - "x-amz-server-side-encryption": "AES256"
    - Amazon S3 provides the encryption key
    - When uploading through the S3 ui, this is seen as the "Amazon S3 master key" encryption option
    
#### SSE-KMS
- Encryption using keys handled & managed by KMS on AWS
- KMS Advantages: User control + audit trail
- Object is encrypted server side
- To make this work, you must set a header when you send your data to Amazon S3
    - "x-amz-server-side-encryption": "aws:kms"
    - Key used: KMS Customer Master key (CMK)
    - When uploading through the S3 ui, this is seen as the "Amazon KMS master key" encryption option

#### SSE-C
- Server-side encryption using data keys fully managed by the customer outside AWS
- Amazon will not store the encryption key you provide
- **HTTPS must be used!**
- Encryption key must be provided in HTTP headers for every HTTP request made. Exam doesn't ask which header, just says "in the header"
- To get the data back from S3, you also need to provide the encryption key
- Object is encrypted into the bucket, Amazon then throws away the key

#### Client Side Encryption
- Can be achieved by using a library such as the S3 encryption client
- Clients must encrypt the data themselves before sending it to S3
- Clients must also decrypt the data
- Customer fully managed the encryption cycle

#### Encryption in transit (SSL)
- AWS S3 exposes:
    - HTTP endpoint: not encrypted
    - HTTPS endpoint: encryption in flight (also called SSL/TLS)
- You can use HTTP or HTTPS, AWS recommends HTTPS.
- HTTPS is mandatory for SSE-C
- If you set a type of bucket encryption for the entire bucket, files uploaded will have that type of encryption by default.

#### Hands on
- Through the ui
- Upload
- Add a file
- There's an encryption option!
    - None
    - Amazon S3 mater key (SSE-S3)
    - Amazon KMS master key (SSE-KMS)
- On the bucket properties, you can also change the "Default encryption" settings

### 100. S3 Security & Bucket policies 
There's 3 types of security.
- User based
    - Uses IAM policies. Which API calls should be allowed for a specific user from IAM console (Popular on the exam)
- Resource based
    - Bucket policies - bucket wide rules from the S3 console -- allows cross acount (?) (Popular on the exam)
    - Supports Object Access Control list (ACL) - finer grain. Rare at exam
    - Bucket Access Control list (ACL) - less common. Rare at exam
- Note: An IAM principal (?) can access an S3 object if:
    - The user IAM permission allows it OR the resource policy allows it
    - And there's no explicit deny!

#### S3 bucket policies
- JSON based policy 
    - Resources: buckets and objects
    - Actions: Set of API to allow or deny
    - Effect: Allow / Deny
    - Principal: The account or user to apply the policy to
- You can use the S3 bucket policy to:
    - Grant public access to the bucket
    - Force objects to be encrypted at upload
    - Grant access to another account (Cross Account access)

#### Bucket settings for Block Public Access
- Block public access to buckets and objects granted through
    - New access control lists (ACLs)
    - Any access control lists (ACLs)
    - New public public bucket or access point policies
- Block public and cross-account access to buckets and objects through any public bucket or access point policies
- What you need to remember going into the exam: There is a way to block public access to your S3 bucket through these settings
- These settings were created to prevent company data leaks
- If you know that your bucket should never be public, leave these on
- Can be set at the account level

#### S3 security - Other
- Networking
    - Supports VPC Endpoints (for instances in VPC without full internet access. These instances can talk to S3 directly)
- Logging and audit
    - S3 access logs can be stored in other S3 bucket
    - API calls can be logged in AWS CloudTrail
- User Security
    - MFA Delete: MFA can be required in versioned buckets to delete objects
    - Signed URLs: URLs that are valid only for a limited time. (ex: premium video service for logged users)

### 101. S3 Bucket Policies Hands On
- Permissions
    - Bucket policy!
    - Let's go on the policy generator

- Policy generator: https://awspolicygen.s3.amazonaws.com/policygen.html
- Yay I generated a policy that rejects unencrypted uploads!
```json
{
  "Id": "Policy1592501316863",
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Stmt1592501143584",
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Deny",
      "Resource": "arn:aws:s3:::danidanibucket/*",
      "Condition": {
        "Null": {
          "s3:x-amz-server-side-encryption": "true"
        }
      },
      "Principal": "*"
    },
    {
      "Sid": "Stmt1592501257457",
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Deny",
      "Resource": "arn:aws:s3:::danidanibucket/*",
      "Condition": {
        "StringNotEquals": {
          "s3:x-amz-server-side-encryption": "AES256"
        }
      },
      "Principal": "*"
    }
  ]
}
```
- Oh nice, I get a forbidden error trying to upload a non-encrypted file to my bucket
- You can also google "s3 bucket policy deny encryption"
- Access control lists are not at the exam any longer
- You can set it at the bucket level
- You can set it at the object level
- Remember: There's a big setting named "block public access". It blocks public access
- Blocking S3 public access can also be done at the account level!

### 102. S3 Websites 
- S3 can host static websites that are available publicly!
- The website URL will be:
    - <bucketname>.s3-website.<AWSRegion>.amazonaws.com
    - OR
    - <bucketname>.s3-website-<AWSRegion>.amazonaws.com
- If you get a 403 forbidden error, make sure the bucket policy allows public reads!

#### Hands on
- Remove the mandatory encrypted file bucket policy!
- Uploaded an index.html
- Properties > Static Website Hosting > Use this bucket to host a website
- Uploaded index.html and error.html
- Oh noes, got a forbidden! 
- Go back to permissions and uncheck "Block public access"
- But you still need to create a policy for the website to be publicly access
```json
{
  "Id": "Policy1592502734656",
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Stmt1592502733881",
      "Action": [
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::danidanibucket/*",
      "Principal": "*"
    }
  ]
}
```
- After successfully implementing this policy, you'll see a message saying "This bucket has public access"
- You can access your website!

### 103. S3 CORS 
- Cross Origin Resource Sharing
- Stephane claims: Complicated topic but does come up at the exam in very simple use cases
- An origin is a scheme (protocol), host (domain) and port
    - Ex: https://example.com (implied port is 443 for HTTPS)
- (CORS?) Is a web browser mechanism to allow requests to other origins while visiting the main origin 
- Same origin: http://example.com/app1 and http://example.com/app2
- Different origins: http://www.example.com and http://other.example.com 
- The request won't be fulfilled unless the other origin allows for the requests, using CORS Headers
    - ex: Access-Control-Allow-Origin

#### CORS - Diagram
- Web browser visits site at https://www.example.com
- There is a second webserver (denominated Cross Origin) which is https://www.other.com
- The browser asks other.com if it can make requests from example.com with an OPTIONS http request

#### S3 CORS
- If a client does a cross-origin request on our S3 bucket, we need to enable the correct CORS headers
- It's a popular exam question!
- You can allow for a specific origin or for * (all origins)

#### Quick Diagram
- A web browser is getting files from our static website bucket
- But there is a second bucket that is going to be our cross origin bucket, also enabled as a website that contains some files that we want.
- You get index.html from the first bucket, but then index.html says you need to perform a GET to coffe.jpg, which is on the other bucket (other origin)
- If the other bucket is configured with the correct CORS headers, the website will be able to make the request
- The CORS headers have to be defined on the Cross origin bucket

### 104. CORS Hands On
- If you do a same origin request (like a page requesting another page on your bucket) things work!
- Create new bucket, like `bucket-assets`, and make it public. Add the public bucket policy!
- Enable this bucket for static website hosting
- Add the extra page on that bucket!
- On your first index.html on the first bucket, add the extra page with the entire second bucket URL in there
- On the chrome developer console you can see that you'll get an error! Blocked by CORS policy!
- For this to work, you need to change the CORS on the second bucket that has the extra page 
- Permissions > CORS Configuration. Allow your first bucket to make requests here!
- Going into the exam: If one website makes a request to another website, that other website needs to have the correct CORS headers for that request to work!

### 105. S3 Consistency Model 
- Amazon S3 is an eventually consistent system
- Read after write consistency for PUTS of new objects
    - As soon as a new object is written, we can retrieve it
    - ex: (PUT 200 -> GET 200) 
    - This is true, except if you did a GET before to see if the object existed. (GET 404 -> PUT 200 -> GET 404). It will eventually be consistent
- There is eventual consistency for DELETES and PUTS of existing objects
    - If you read an object after updating, you might get an older version. PUT 200 -> PUT 200 -> GET 200 (might get an older version)
    - The way to get the new version is just to wait a little bit
    - If you delete an object, you might still be able to retrieve it for a short time
    - ex: DELETE 200 -> GET 200 (it's eventually consistent, if you wait a few seconds, it'll be gone)
- Note: There is no way to request "strong consistency" in S3





When something is done in S3, it can take a bit of time to replicate and to be active. There are two types of consistency you get on S3: 
#### Read after write consistency for PUTS of new objects
- When you use the put object API and the object did not exist before, as soon as the object is written, we can retrieve it

- This is true only if you did not do a GET before the object was uploaded to see if it existed, you may get a 404 because you are eventually consistent. The previous result gets cached. You need to wait a bit.
- ex: 

#### Eventual consistency for DELETES and PUTS of existing objects
- If you read an object after updating it, you might get the older version, ex: 

In other words:
- As soon as you write, you can retrieve it
- If you overwrite an object or delete an object you get eventual consistency, you might get an older version or might be able to get the object after it was deleted just for a short time.

### Quiz!
