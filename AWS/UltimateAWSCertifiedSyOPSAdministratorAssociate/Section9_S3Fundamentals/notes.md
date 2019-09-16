## Overal chapter status: Studied carefully taking all the notes

### 92. S3 Fundamentals 
Huge exam topic!

### 93. S3 buckets and objects

#### Buckets
- Amazon allows you to store objects (files) in "buckets" (directories)
- Buckets must have a **globally unique name**, across all AWS accounts!
- Buckets are be defined at the region level
- There's a naming convention!
    - No uppercase
    - No underscore
    - 3-63 characters long
    - Not an IP
    - Must start with a lowercase letter or number

#### Objects
- Objects (files) have a key. The key is the **full** path:
    - mybucket/my_file.txt
    - mybucket/another_folder/my_file.txt
- There's no concept of directories within buckets, just very long key names with slashes (the UI tricks you)
- ObjectValues are the contents of the body:
    - Max size is 5TB
    - If uploading more than 5GB, you must use "multi-part upload"
- Your object can have metadata (list of key/value pairs, system or user metadata)
- Your object can have tags (key/value pairs, up to 10, useful for security/lifecycle)
- There's a version ID if versioning is enabled

#### Creating a bucket
- S3 does not require a region selection, but when you create a bucket it is attached to a region
- There's two ways to open a file with S3. With default settings you can open the file through S3, but not with the public link as files are not public by default

### 94. S3 versioning - basics 
 - You can version your files in AWS S3!
 - Versioning is enabled at **bucket level** 
 - When you overwrite a file, it will increment the version of that file and keep many versions of it
 - It's considered a best practice to version your buckets
    - You can protect yourself against unintended deletes and roll back if there are any issues
    - Downside is you use more space
- Any file that is not versioned prior to enabling versioning will have version "null"
- A versioned file, upon deletion, can be restored from an older version (latest version will appear as "Delete Marker")
    - The file not really deleted, it just has a Delete Market on it

### 95. S3 Encryption  
This section is popular on the exam!

- There are 4 methods for encrypting objects in S3:
    - SSE-S3: Encrypts S3 objects using keys handled & managed by AWS
    - SSE-KMS: Leverages AWS Key Management Service to manage encryption keys
    - SSE-C: When you provide your own keys and Amazon encrypts the data for you
    - Client side encryption: When you encrypt your data clientside.

It's important to understand which ones are adapted to which situation in the exam!

#### SS3-S3
- Encryption using keys handled & managed by AWS S3
- Object is encrypted server side
- AES-256 encryption type
- To make this work, you must set a header when you send your data to Amazon S3
    - "x-amz-server-side-encryption": "AES256"
    - Amazon S3 provides the encryption key
    - When uploading through the S3 ui, this is seen as the "Amazon S3 master key" encryption option
    
#### SSE-KMS
- Encryption using keys handled & managed by KMS
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

If you set a type of bucket encryption for the entire bucket, files uploaded will have that type of encryption by default.


### 96. S3 Security & Bucket policies 
There's 3 types of security.
- User based
    - Uses IAM policies. Which API calls should be allowed for a specific user from IAM console (Popular on the exam)
- Resource based
    - Bucket policies - bucket wide rules from the S3 console -- allows cross acount (?) (Popular on the exam)
    - Supports Object Access Control list (ACL) - finer grain
    - Bucket Access Control list (ACL) - less common 

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

#### S3 security - Other
- Networking
    - Supports VPC Endpoints (for instances in VPC without full internet access. These instances can talk to S3 directly)
- Logging and audit
    - You can have S3 generate access logs and store those in another bucket. Do not store logs in the same bucket as this will be recursive and generate infinite logs!
    - API calls can be logged in AWS CloudTrail
- User Security
    - You can enable MFA in versioned buckets to delete objects
    - Signed URLs: URLs that are valid only for a limited time. (popular at the exam!)

#### Hands on
- Access control list: Not popular
- Bucket policy: Very popular

#### To force policies
- How to ensure that the objects can be encrypted at upload (popular question at the exam)
- Bucket > Permissions > Bucket policy > Generate a bucket policy (you can use the generator) like this: https://aws.amazon.com/blogs/security/how-to-prevent-uploads-of-unencrypted-objects-to-amazon-s3/
- Bucket policies are cool.
- You can only upload from the UI if you select "Amazon S3 master key" upon upload time.

### 97. S3 Websites 
- S3 can host static websites that are available publicly!
- The website URL will be:
    - <bucketname>.s3-website.<AWSRegion>.amazonaws.com
    - OR
    - <bucketname>.s3-website-<AWSRegion>.amazonaws.com
- If you get a 403 forbidden error, make sure the bucket policy allows public reads!
- Sometimes when you change a bucket policy the result isn't instantaneous, it can take a minute or two
- Properties > Static website hosting > Use this bucket to host a website
- But after doing that I get a 403 as my bucket isn't public and I don't have the appropriate policy for it set up

#### Bucket permissions for S3 websites:
- Block public access to buckets and objects granted through new public bucket policies: Off
- Block public and cross-account access to buckets and objects through any public bucket policies: Off

Created an S3 bucket policy with Policy generator to allow GetObject to any object:

```json
{
    "Version": "2012-10-17",
    "Id": "Policy1568644565619",
    "Statement": [
        {
            "Sid": "Stmt1568644552111",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::testmctest/*"
        }
    ]
}
```

### 98. S3 CORS 
Popular on the exam!
- Cross Origin Resource Sharing allows you to limit the number of websites that can request your files in S3 (and limit your costs)
- It's from when you have a static page on a S3 bucket (bucket A) that gets images from another bucket (bucket B). You can set images on bucket B to only be acessible to your static site on bucket A, not the entire internet.
- Here's a snipped of documentation:
```text
Scenario 1: Suppose that you are hosting a website in an Amazon S3 bucket named website as described in Hosting a Static Website on Amazon S3. Your users load the website endpoint http://website.s3-website-us-east-1.amazonaws.com. Now you want to use JavaScript on the webpages that are stored in this bucket to be able to make authenticated GET and PUT requests against the same bucket by using the Amazon S3 API endpoint for the bucket, website.s3.amazonaws.com. A browser would normally block JavaScript from allowing those requests, but with CORS you can configure your bucket to explicitly enable cross-origin requests from website.s3-website-us-east-1.amazonaws.com.
Scenario 2: Suppose that you want to host a web font from your S3 bucket. Again, browsers require a CORS check (also called a preflight check) for loading web fonts. You would configure the bucket that is hosting the web font to allow any origin to make these requests.
From: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
```

### 99. S3 Consistency Model 

When something is done in S3, it can take a bit of time to replicate and to be active. There are two types of consistency you get on S3: 
#### Read after write consistency for PUTS of new objects
- When you use the put object API and the object did not exist before, as soon as the object is written, we can retrieve it
- ex: (PUT 200 -> GET 200)
- This is true only if you did not do a GET before the object was uploaded to see if it existed, you may get a 404 because you are eventually consistent. The previous result gets cached. You need to wait a bit.
- ex: (GET 404 -> PUT 200 -> GET 404)

#### Eventual consistency for DELETES and PUTS of existing objects
- If you read an object after updating it, you might get the older version, ex: PUT 200 -> PUT 200 -> GET 200.
- If you delete an object, you might still be able to retrieve it for a short time, ex: DELETE 200 -> GET 200

In other words:
- As soon as you write, you can retrieve it
- If you overwrite an object or delete an object you get eventual consistency, you might get an older version or might be able to get the object after it was deleted just for a short time.

### Quiz!
- Buckets must have a unique name
- Files that are already on the bucket when you start versioning will have a version of null
- To have encryption happen when you upload the files to S3, but to manage the key yourself and not have it on Amazon, use the SSE-C encryption method
- SSE-C requires HTTPS
- 