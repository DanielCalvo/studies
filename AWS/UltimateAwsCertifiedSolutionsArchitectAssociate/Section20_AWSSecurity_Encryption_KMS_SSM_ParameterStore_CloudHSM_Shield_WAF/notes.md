### 206. AWS Encryption - Section Introduction
- Exam asks a ton of questions about security!

### 207. Encryption 101

#### Encryption in flight
- Data is encrypted before sending and decrypted after receiving
- SSL certificates help with encryption (HTTPS)
- Encryption in flight ensures no MITM (man in the middle attack) can happen

#### Server side encryption at rest
- Data is encrypted after being received by the server
- Data is decrypted before being sent
- It is stored in an encrypted form thanks to a key (usually a data key)
- The encryption / decryption keys must be managed somewhere and the server must have access to it

#### Client side encryption
- Data is encrypted by the client and never decrypted by the server
- Data will be decrypted by a receiving client
- The server should not be able to decrypt the data
- Could leverage Envelope Encryption

### 208. KMS Overview
- Anytime you hear “encryption” for an AWS service, it’s most likely KMS
- Easy way to control who and what access to your data, AWS manages all the keys for us
- Fully integrated with IAM for authorization
- Seamlessly integrated into:
    - Amazon EBS: encrypt volumes
    - Amazon S3: Server side encryption of objects
    - Amazon Redshift: encryption of data
    - Amazon RDS: encryption of data
    - Amazon SSM: Parameter store
    - Etc...
- But you can also use the CLI / SDK
- A central service to AWS, Stephane claims

#### KMS - Customer Master Key (CMK) Types
- Symmetric (AES-256 keys)
    - First offering of KMS, single encryption key that is used to Encrypt and Decrypt
    - AWS services that are integrated with KMS use Symmetric CMKs (under the tood)
    - Necessary for envelope encryption
    - You never get access to the Key unencrypted (must call KMS API to use that key)
- Asymmetric (RSA & ECC key pairs)
    - Public (Encrypt) and Private Key (Decrypt) pair
    - Used for Encrypt/Decrypt, or Sign/Verify operations
    - The public key is downloadable, but you can access the Private Key unencrypted
    - Use case: encryption outside of AWS by users who can’t call the KMS API

#### AWS KMS, more details
- Able to fully manage the keys & policies:
    - Create (for eys)
    - Rotation policies (for eys)
    - Disable (for eys)
    - Enable (for eys)
- Able to audit key usage (using CloudTrail)
- Three types of Customer Master Keys (CMK):
    - AWS Managed Service Default CMK: free (like when you go on an EBS and use the AWS/EBS key)
    - User Keys created in KMS: $1 / month
    - User Keys imported (must be 256-bit symmetric key): $1 / month
- + pay for API call to KMS ($0.03 / 10000 calls)

#### AWS KMS 101 (When to use KMS?)
- Anytime you need to share sensitive information... use KMS
    - Database passwords
    - Credentials to external service
    - Private Key of SSL certificates
- The value in KMS is that the CMK used to encrypt data can never be retrieved by the user, and the CMK can be rotated for extra security

#### AWS KMS, more details
- Never ever store your secrets in plaintext, especially in your code!
- What you should do is that you should encrypt these secrets can be stored in the code / environment variables
- KMS can only help in encrypting up to 4KB of data per call
- If you want to have more data encrypted (> 4 KB) use envelope encryption
- To give access to KMS to someone:
    - Make sure the Key Policy allows the user to access the key
    - Make sure the IAM Policy allows the API calls

#### Copying snapshots across regions
- KMS Keys are bound to a specific region
- Imagine you have a EBS Volume encrypted with KMS (KMS Key A) in eu-west-2 and you want to move that to ap-southeast-2
    - This process applies to other things encrypted with KMS as well
- Create an EBS snapshot encrypted with KMS
- Then copy that snapshot over to the new region
- But then you specify a new KMS key to re-encrypt the data with!
- Then you create a volume from that snapshot on the new region and it will be encrypted with the new key
- Since KMS keys are region specific, when you migrate something encrypted to another region, this object will have to be re-encrypted

#### KMS Key policies
- Control access to KMS keys, “similar” to S3 bucket policies
- Difference: you cannot control access without them
- If you don't specify a key policy, then no one can access your key!
- Default KMS Key Policy:
    - Created if you don’t provide a specific KMS Key Policy
    - Complete access to the key to the root user = entire AWS account has the right to use this KMS key
    - Gives access to the IAM policies to the KMS key (wtf?)
- Custom KMS Key Policy:
    - Define users, roles that can access the KMS key
    - Define who can administer the key
    - It's useful for cross-account access of your KMS key

#### Copying snapshots across accounts
1. Create a Snapshot, encrypted with your own CMK
2. Attach a KMS Key Policy to authorize cross-account access on that key
3. Share the encrypted snapshot 
4. (in target account) Create a copy of the Snapshot, encrypt it with a KMS Key in your account
5. Create a volume from the snapshot

### 209. KMS Hands on with CLI

### 210. SSM Parameter Store Overview
- Secure storage for configuration and secrets
- Optional Seamless Encryption using KMS (you can store your secrets, encrypted!)
- Serverless, scalable, durable, has easy SDK
- Has version tracking of configurations / secrets
- Configuration management is done using path & IAM
- You can get notifications with CloudWatch Events
- There is integration with CloudFormation
- Application > plaintext config > SSM Parameter Store > Checks IAM permissions (to make sure we can get the config) > SSM returns config
- Application > encrypted config > SSM Parameter Store > Checks IAM permissions (to make sure we can get the config) > Decryption service > AWS KMS > SSM returns config

#### SSM Parameter store hierarchy
- /my-department/
    - my-app/
        - dev/
            - db-url
            - db-password
        - prod/
            - db-url
            - db-password
    - other-app/
- /other-department/
- /aws/reference/secretsmanager/secret_ID_in_Secrets_Manager (using parameters store, you can also reference secrets from the secrets manager)
- /aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2 (you're also able to reference parameters directly from AWS, like to retrieve the lastes AMI ID (?)
- A lambda function would retrieve these parameters using GetParameters or GetParametersByPath

#### Standard and advanced parameter tiers
- Standard is free, advanced is paid
- Standard doesn't allow you to have parameter policies  :(
- The exam probably won't ask you the difference between the standard and advanced tiers

#### Parameter Policies (for advanced parameters)
- Allows you to assign a TTL to a parameter (expiration date) to force updating or deleting sensitive data such as passwords
- Can assign multiple policies at a time
- Shown as examples
    - Expiration policy (parameter expires at)
    - Expiration notification (sends you a notification before the parameter expires)
    - NoChangeNotificaton (if parameter hasn't been changed in 20 days, then send me a notification)


### 211. SSM Parameter Store Hands on (CLI)
### 212. SSM Parameter Store Hands on (AWS Lambda)

### 213. AWS Secrets Manager - Overview
- Newer service, came after SSM meant for storing secrets
    - More secret oriented than SSM
- Has capability to force rotation of secrets every X days
- Has the capability to automate generation of secrets on rotation (uses Lambda)
- Integrates with Amazon RDS (MySQL, PostgreSQL, Aurora)
- Secrets are encrypted using KMS
- Any time the exam mentions: Secret storage, rotation of secrets, integration with RDS, think secrets manager
- Mostly meant for RDS integration
- Seems like a simple service

### 214. AWS Secrets Manager - Hands on

### 215. CloudHSM
- With KMS, AWS manages the software for encryption
- With CloudHSM, AWS provisions the encryption hardware, you have to use your own client
- Dedicated Hardware (HSM = Hardware Security Module)
    - They give you the hardware and you manage all the rest
- You manage your own encryption keys entirely (not AWS)
- The HSM device is tamper resistant, FIPS 140-2 Level 3 compliance
- CloudHSM clusters can be spread across Multi AZ (HA) – must setup 
- Supports both symmetric and asymmetric encryption (SSL/TLS keys)
- No free tier available
- Must use your own CloudHSM Client Software (?))
- Redshift supports CloudHSM for database encryption and key management
- Good option to use with SSE-C encryption

#### Diagram
- AWS only manages the hardware, you cannot recover your keys if you lose your credentials
- CloudHSM Client > SSL Connection > AWS CloudHSM
- IAM permissions will allow you to
    - CRUD an HSM cluster
- But to manage the keys in the cluster itself you must use the CloudHSM software which will allow you to
    - Manage the keys
    - Manage the users

### 216. Shield - DDoS Protection
- AWS Shield Standard
    - Free service that is activated for every AWS customer
    - Provides protection from attacks such as SYN/UDP Floods, Reflection attacks and other layer 3/ layer 4 attacks
- AWS Shield Advanced
    - Optional DDoS mitigation service ($3,000 per month per organization)
    - Protect against more sophisticated attack on Amazon EC2, Elastic Load Balancing (ELB), Amazon CloudFront, AWS Global Accelerator, and Route 53
    - 24/7 access to AWS DDoS response team (DRP)
    - Protect against higher fees during usage spikes due to DDoS

### 217. Web application firewall (WAF)
#### AWS WAF
- Protects your web applications from common web exploits (Layer 7)
- Layer 7 is HTTP (vs Layer 4 is TCP)
- WAF can be deployed on 3 things, remember them on the exam: Application Load Balancer, API Gateway and CloudFront
- To use WAF you need to define a Web ACL (Web Access Control List):
    - Rules can include: IP addresses, HTTP headers, HTTP body, or URI strings
    - Protects from common attacks at the HTTP level - SQL injection and Cross-Site Scripting (XSS)
    - You can put size constraints on queries, You can also have geo-match rules to block specific countries
    - Rate-based rules (to count occurrences of events based on IP for example) – for DDoS protection

#### AWS Firewall manager
- Manages rules of all the WAFs on all accounts of an AWS Organization
- Can define a common set of security rules
- WAF rules apply to the same things (Application Load Balancer, API Gateways, CloudFront)
- But also can apply to AWS Shield Advanced (ALB, CLB, Elastic IP, CloudFront)
- Security Groups for EC2 and ENI resources in VPC
- It's a centralized way to manage all of these things together

#### Sample architecture for DDoS protection
- Users > Route 53 (with shield) > CloudFront Distribution (with shield) > Load Balancer (with shield) > ASG (with shield)

### 218. WAF & Shield - Hands on

### 219. Shared Responsability Model
- AWS responsibility - Security of the Cloud
    - They're responsible for protecting infrastructure (hardware, software, facilities, and networking) that runs all of the AWS services
    - They're responsible for managed services like S3, DynamoDB, RDS etc
- Customer responsibility - Security in the Cloud
    - For example, for an EC2 instance, the customer is responsible for management of the guest OS (including security patches and updates), firewall & network configuration, IAM etc

#### For example, RDS
- AWS responsibility:
    - Manage the underlying EC2 instance, disable SSH access
    - Automated DB patching
    - Automated OS patching
    - Audit the underlying instance and disks & guarantee it functions
- Your responsibility:
    - Check the ports / IP / security group inbound rules in DB’s SG
    - In-database user creation and permissions
    - Creating a database with or without public access
    - Ensure parameter groups or DB is configured to only allow SSL connections
    - Database encryption setting

#### For example, S3
- AWS responsibility:
    - Guarantee you get unlimited storage
    - Guarantee you get encryption
    - Ensure separation of the data between different customers
    - Ensure AWS employees can’t access your data
- Your responsibility:
    - Bucket configuration
    - Bucket policy / public setting
    - IAM user and roles
    - Enabling encryption
- Check this out for more details: https://aws.amazon.com/compliance/shared-responsibility-model/

### Quiz!