## Overal chapter status: 

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

```text

```
### 97. S3 Websites 

```text

```
### 98. S3 CORS 

```text

```
### 99. S3 Consistency Model 

```text

```

### Quiz!
