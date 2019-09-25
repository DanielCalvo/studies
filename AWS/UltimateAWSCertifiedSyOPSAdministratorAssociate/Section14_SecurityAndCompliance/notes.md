- Chapter status: Took notes attentively, got the quiz mostly right
- Follow up on better understanding JSON policies!

### 162. Intro!
- Uh-oh, one of the hardest sections, we learn about many things!

### 163. Shared responsibility model
- AWS's responsibility: Security **of** the cloud
    - Hardware, software, facilities, network, etc, everything running the AWS's services
    - Also responsible for managed services like S3, DynamoDB and many others 
- Customer's responsibility: Security **in** the cloud
    - For an EC2 instance, the customer is responsible for the management of the guest OS (including patches and updates), firewall & network config, IAM, etc

#### Example: RDS
AWS's responsibility:
    - Manage the underlying EC2 instance with no SSH access for you
    - Automated DB patching
    - Automated OS patching
    - Audit the underlying instance, disks & guarantee it functions
- Your responsibility:
    - Check the ports / IP / security groups in the DB's SG
    - In-database user creation and permissions
    - Creating a databasw with or without public access
    - Ensure parameter groups or DB is configured only to allow SSL connections    
    - Database encryption setting

#### Example: S3
- AWS:
    - Guarantee unlimited storage
    - Guarantee encryption
    - Ensure separation of the data between different customers
    - Ensure AWS employees can't access your data
- You:
    - Bucket configuration
    - Bucket policy / public setting
    - IAM user and roles
    - Enabling encryption

Also: https://aws.amazon.com/compliance/shared-responsibility-model/

### 164. DDoS, AWS Shield and AWS WAF 
- Aha, I know what a DDoS attack is :D

#### DDOS Protection on AWS
- AWS Shield Standard: Enabled by default free of charge, protects against DDoS attacks for your website and application.
- AWS Shield Advanced: 24/7 premium DDoS protection
- AWS WAF: Filters specific requests based on rules, like requests being too large (WAF = Web Application Firewall)
- Cloudfront and Route 53
    - Availability protection using global edge network
    - Combined with AWS Shield, provides attack mitigation at the edge
- An action point you can take: Be ready to scale -- leverage AWS auto scaling
- You can also separate static resources (html, css, images into S3/CloudFront) from dynamic ones (EC2/ALB). It's harder to DDoS your application this way.

Author recommends reading: `https://d1.awsstatic.com/whitepapers/Security/DDoS_White_Paper.pdf`

#### AWS Shield
- AWS Shield standard:
    - Free service that is activated for every AWS customer
    - Provides protection from attacks such as SYN/UDP floods, reflection attacks and other layer 3/4 attacks
- AWS Shield advanced:
    - The exam asks about it! 
    - Optional DDoS migration service (3000 USD per month per organization)
    - Protects against more sophisticated attack on CloudFront, Route 53, Classic (?), Application & Network Load Balancer, Elastic IP / EC2
    - 24/7 access to a AWS DDoS response team (DRP)
    - If you incur higher fees due to a DDoS attack, these fees are waived if you have AWS Shield Advanced

#### AWS WAF - Web Application Firewall
- Protects your web application from common web exploits
- You can define customizable web security rules:
    - Allow/block traffic to your application based on IP addresses, HTTP headers, HTTP body or URI strings
    - Protects from common attacks - SQL Injections and Cross-Site Scripting (XSS)
    - Protects against bots and bad user agents
    - You can also put size contraints on requests (ex: no requests larger than 2mb)
    - Geo match
- You can deploy this on CloudFront, Application Load Balancewr or API Gateway
- There is also a marketplace of rules
    
AWS Shield: DDoS protection
AWS WAF: Security rules to block certain types of traffic

### 165. Penetration testing on AWS
- Permission is required for all tests! 
- Request permission only using AWS root credentials
- You must perform the pen testing yourself. No third parties!
- Things you can pen test against: EC2, ELB, Aurora, CloudFront, API Gateway, Lambda, Lightsail.
- The rest you can't pen test again
- You also cannot pentest against nano/micro/small EC2 instances
- Pentest request takes about 2 days to be approved
- You need to specifically list the endpoints/resources you want to attack on AWS.
- You also need to specify where you traffic is coming from! (IP addresses)

### 166. AWS Inspector
- Only for EC2 instances. The exam will try to trick you! (ex: Can you use inspector on RDS? No you can't)
- Helps you analyze against known security vulnerabilities
- Also helps analyze  against unintended network accessibility
- You need to install an inspector agent on EC2 instances
- You define a template (rules package, duration, attributes, SNS topics)
- No own custom rules possible - you can only use AWS's managed rules
- An assessment runs and after it does, you get a report with a list of vulnerabilities found

#### What does AWS inspector evaluate
- Remember: Only for EC2 instances
- For network assessments
    - Network Reachability
For Host assessments
    - Common vulnerabilities and exposures
    - Center for Internet Security (CIS) Benchmarks
    - Security Best Practices
    - Runtime Behaviour Analysis

#### Inspector Hands on
- Create an EC2 in a tag (such as Inspector: True, but can be whatever you want)
- You can install the agent either by googling it and doing it manually or automating it through SSM
- Very interesting! You get vulnerability information.
- Possible exam question: "We need to know of our instances are vulnerable to attacks or if they need to be patched." Use AWS Inspector!

### 167. Logging in AWS (for security and compliance)
- To help with compliance requirements, AWS provides many service-specific security and audit logs
- Service logs include
    - CloudTrail trails - trace all API calls
    - Config Rules - for config & compliance over time
    - CloudWatch Logs - for full data retention
    - VPC flow Logs - IP traffic within your VPC
    - ELB Access Logs - metadata of requests made to your load balancers
    - CloudFront Logs - web distribution access logs
    - WAF logs - Full logging of all requests analyzed by the service
- Logs can be analyzed using AWS Athena if they're stored in S3 (Common exam question!)
- "We have a log, how can we analyze it or check what happened to X component?" You can use Athena + ELB access logs (for example) + S3
- You should encrypt logs in S3 and controll access using IAM & Bucket Policies, MFA (multi factor auth)
- If you need to retain these logs for a very long time, remember to move them to Glacier!

### 168. GuardDuty
- A service that might be a bit difficult to understand at first
- Intelligent Threat discovery to protect an AWS account
- Uses Machine learning algorithms, anomaly detection and 3rd party data
- One click to enable (30 days trial), no need to install software. Not cheap :o
- Input data includes:
    - CloudTrail logs: unusual API calls, unauthorized deployments
    - VPC flow logs (unusual internal traffic, unusual IP addresses)
    - DNS Logs: Compromised EC2 instances sending encoded data within DNS queries
- Notifies you in case of findings
- Can be integrated with AWS Lambda

#### Hands on
- The thing provides some interesting security notices
    - Instance reached a domain associated with bitcoin
    - Suspicious DNS requests
    - RDP brute force access
    - Possible Trojans

### 169. Trusted Advisor
- High level AWS Account assessment. No install of anything needed
- Analyzes your AWS accounts and provides recommendations on
    - Cost Optimization
    - Performance
    - Security
    - Fault Tolerance
    - Service limit
- Popular exam question: "How do we know if we're reaching our service limits, how can we get a general overview of that?". TrustedAdvisor!
- Not exactly 100% free :(
- Core checks and recommendations - free to all customers, but some things are disabled unless you have a business support plan
- Can enable weekly email notifications from the console
- Full Trusted Advisor - Available for Business & Enterprise support plans
    - Ability to set CloudWatch alarms when reaching limits

#### Hands on look
- Cost Optimization: Complements billing exploring features
- Performance: Shows some alarms similar to Prometheus with node_exporter, except networks with too many rules, that's cool
- Security: Checks the config of miscelaneous services/items (SGs, S3s, public RDS snapshots, IAM use, MFA on root account...)
- Fault tolerance: Snapshot age, EC2 AZ balance, LB optimization...
- Service limits: Show's if you're getting close to a service limit

### 170. Encryption 101

#### Encryption in flight
- Data is encrypted before sending and decrypted after receiving
- SSL certificates help with encryption (HTTPS)
- No man in the middle attacks can happen

#### Server side encryption at rest
- Stored in an encrypted form
- Encryption/decryption keys must be managed somewhere and the server must have access to it
- The keys must be managed somewhere and the server must have access to it

#### Client side encryption
- Data is encrypted by the client and never decrypted by the server
- Exam asks about Envelope Encryption, but this is covered later
- Client receives data, stores it encrypted (ex: in S3). When it needs that data, it retrieves it and decrypts it

### 171. KMS Overview + Encryption in place
- Any time you hear "encryption" for an AWS service, it's most likely KMS
- Easy way to control access to your data, AWS manages the key for you
- Fully integrated with IAM for authorization 
- Seamlessly integrated into:
    - EBS: Encrypted volumes
    - S3: Server side encryption of objects
    - Redshift: Encryption of data
    - RDS: Encryption of data
    - SSM: Parameter store
    - Almost any AWS service that uses encryption is integrated with KMS
    - You can also use the CLI/SDK to perform encryption with KMS

#### AWS KMS 101
- Any time you need to share sensitive information, use KMS
    - DB passwords
    - Credentials to an external service
    - Private key of SSL certificates
- The value in KMS (Key management system) is that the CMK (Customer Master Key) can never be retrieved by the user, and the CMK can be rotated for extra security
- Never store your secrets in plaintext! (That goes without saying...)
- Secrets > KMS > used in code/environment variables
- KMS can only help in encrypting up to 4KB of data per call
- Possible exam question: "We want to encrypt 1mb of data, how do we do it?"
    - If you data is bigger than 4kb, use envelope encryption
- To give someone access to KMS:
    - Make sure they Key Policy allows the user
    - Make sure the IAM policy allows the API calls

#### AWS KMS Functionalities
- Able to manage keys / policies
    - Create
    - Rotation policies
    - Disable
    - Enable
- Audit key usage (using CloudTrail)

- There are 3 types of Customer Master Key (CMK):
    - AWS MAnaged Service Default CMK: Free
    - User keys created in KMS: 1 usd / month
    - User keys imported (must be 256bit symmetric key): 1 usd / month
- API calls to KMS are paid (0.03 usd per 10.000 calls)

#### How does KMS work?
- Two APIs: Encrypt && Decrypt
- Encrypt
    - Client does a request to KMS to encrypt something with a certain key
    - KMS checks IAM permissions to see if you can encrypt that secret
    - Will send an encrypted secret back
    - The encryption itself happened on the KMS side
- Decrypt
    - Calls decrypt API with the CMK that encrypted the data in the first place
    - IAM permissions are checked
    - KMS decrypts the secret and it sends it back
- All encryption/decryption happens on the KMS side

#### Encryption in AWS services
- Some services require a migration (through snapshot/backup)
- Snapshot first, encrypt snapshot, restore volume as encrypted
- For EBS, RDS, ElastiCache and EFS Network FS  
- S3 supports in place encryption! To encrypt a file on S3 it will happen directly/in place, you don't need to backup or restore anything
- Exam will ask which services require a migration for encryption and which dont!

### 172. CloudHSM Overview
- KMS > AWS manages the software for encryption
- CloudHSM > AWS Provisions the encryption hardware
- Dedicated hardware. HSM = Dedicated security module
- You manage your encryption keys entirely (not AWS)
- The CloudHSM hardware device is tamper resistant
- FIPS 140-2 Level 3 compliance (If you see a question on the exam asking for encryption with this compliance level, its with CloudHSM)
- CloudHSM clusters are spread across multi AZs (HA)
- Supports both symmetric and assymetric encryption (SSL/TLS keys)  
    - Assymetric encryption: HTTPS/SSL
    - KMS is only for symmetric encryption
    - Possible exam question: We're looking for an mechanism that supports both symmetric and assymetric encryption. Answer: CLoudHSM
- No free tier available
- Must use the CloudHSM client software. No easy API call to do

#### Diagram, but with words:
- CloudHSM client talks to the CloudHSM server
- You manage all the keys in CloudHSM
- AWS manages the hardware, which is dedicated to only one customer
- CloudHSM supports symetric and assymetric encryption
- CloudHSM supports crytographic acceleration (SSL/TLS and Oracle)
- CloudHSM is deployed and managed from a customer VPC. Access can be shared across vpc using vpc peering.
- Not available on the free tier

### 173. KMS + CloudHSM Hands On

#### KMS
- Under KMS, AWS managed keys are keys generated by AWS for it's own services and you don't pay for those
- Key origin can be KMS, External or Custom Key Store (CloudHSM)
- Keys have key administrators and key administrative permissions
- Key have usage permissions, like a key policy, and you can add roles to it

#### CloudHSM
- Minimum is a multi-AZ cluster
- For the exam though, you just need to know on a high level what CloudHSM is

### 174. MFA + IAM Credentials Report
- IAM can be integrated with MFA (Multi factor authentication)
- MFA adds a level of security to your account
- If your password gets compromised, a hacker won't have access to your phone or any other of your devices
- WAS MFA accepts both virtual and MFA devices
- Virtual MFA device: Google / Authy
- MFA for the root user can be configured from the IAM dashboard 
- MFA can be configured from the CLi too
- You can set up MFA for individual users
- To know if your users have MFA enabled or not, you can download a Credentials report. Popular on the exam!
    - Credentials report is a report file on al the IAM users and credentials
    - This shows who does and does not have MFA enabled

#### Hands on
- IAM > Credentials Report > Download report
- This report lets us know if our users are using MFA

### 175. IAM PassRole Action
- Popular at the exam!
- When you create an EC2 instance, you assign a role to it 
- In order to assign a role to an EC2 instance you need IAM:PassRole. Important to remember!
- IAM:PassRole can be used for any service where we assign roles (not only EC2)
- Pro tip: You should be able to read AWS Policies on JSON format

### 176. STS & Cross Account Access
- Security Token Service!
- It allows to grant limited and temporary access to AWS resources
- You'll generate tokens and these tokens will be valid for up to one hour
- You can use this for Cross Account Access
    - Allow users from one AWS account to access resources in another account
- Federation (Active Directory)
    - Provides a non-AWS user with temporary AWS access by linking users' Active Directory credentials
    - You can also use SAML (Security Assetion markup language)
    - You can also allow Single Sign On (SSO) to enable users to log in to the AWS console without assigning IAM credentials
- Any time you have security token that is temporary in nature, STS is in play
- You can also use STS for federation with third party providers, such as Cognito
    - Used mainly in web and mobile applications
    - Makes use of Facebook/Google/Amazon to federate them
    
#### Cross account access
- You define an IAM role for another account to access
- Define which accounts can access this IAM role
- Use AWS STS to retrieve credentials and impersonate the IAM role you have access to (AssumeRole API)
- Temporary credentials can be valid for between 15 minutes to 1 hour

#### Diagram:
- You're a user and you want to access a role, either in the same or in another account
- We do a call to the AssumeRole API on STS which will check against IAM permissions to see if you can do it
- And then it will send you back temporary security credentials 
- These security credentials will allow you to use the role you wanted

### 177. Identity Federation with SAML & Cognito

#### What is identity federation
- Federation lets users outside AWS to assume a temporary role for accessing AWS resources
- These users will assume an identity provided access role
- Federation: The identities of users are stored in a third party, trusted by AWS. 
- Examples for third parties
    - LDAP
    - Microsoft Active Directory (SAML)
    - Single Sign On
    - Open ID
    - Cognito
- Using federation, you don't need to create IAM users, user management is done outside of AWS
- Popular on the exam: SAML, Custom Broker, Cognito

#### SAML Federation for enterprises
- You can integrate ADFS or any SAML thing with AWS
- Your users ADFS users have access to AWS with temporary credentials

#### Custom Identity Broker Application for Enterprises
- Use only if your identity provider is not compatible with SAML 2.0
- This identity broker must determine the appopriate IAM policy
- Same principle as SAML, but you have to implement an identity broker yourself

#### AWS Cognito - Federated Identity Pools for Public Applications
Use case: You have an app and your users need to put a file on a S3 bucket
- Goal
    - Provide direct access to AWS resources from the client side
- How
    - Log in to a federated identity provider, or remain anonymous
    - Get temporary AWS credentials back from the federated identity pool
    - These credentials come with a pre-defined IAM policy stating their permissions
- Example
    - Provide temporary write access to an S3 bucket using a facebook login
- Note
    - Web Identity federation is an alternative to using Cognito but AWS recommends against it 

#### Diagram
- App > logs in to CUP > gets a token back from it
- App > sends token to federated identity 
- Federated identity talks to STS to get credentials
- Federated identity sends temporary AWS credentials to the app
- App can now access S3 bucket (example) using these temporary credentials

### 178. JSON Policies tab
- Well look at IAM and S3 buckets and try to understand them
- In the exam, there will be json policies and you'll have to analyze them :o

```text
https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_examples.html
https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_examples_aws_deny-ip.html

```

### 179. AWS Compliance frameworks
- Wew AWS has a lot of compliance certifications
- There a bunch of region specific certifications (US and Europe)
- https://aws.amazon.com/compliance/iso-certified/
- AWS being compliant does not mean that you are automatically compliant! :o
- You need to get audited as well, depending on whatever compliance you have to comply with
- Remember the AWS shared responsability model

### 180. AWS Artifact
- Not really a service
- AWS Artifact is a portal that provides customers with on-demand access to AWS compliance documentation and AWS agreements
- Artifact reports: You can download AWS security and compliance documents 
- Artifact agreements: Allows you to review, accept and track the status of AWS agreements
- Can be used to support internal audit or compliance
- Auditor: "Do you have documentation proving that AWS is PCI compliant"
- You: "Yes sir, we can find this documentation in AWS Artifact"
- Artifact is just for downloading compliance reports 

### 181. Security and Compliance Section Summary
- AWS Shield: DDoS Protection + 24/7 support for Advanced DDoS protection
- AWS WAF: Firewall to filter incoming requests based on rules
- AWS Inspector: For EC2, install agent and find vulnerabilities
- AWS GuardDuty: Find malicious behaviour with VPC, DNS and CloudTrail Logs
- AWS Trusted advisor: Will analyze AWS account and get recommendations on improving performance, cost and security
- AWS KMS: Encryption keys managed by AWS
- AWS CloudHSM: Hardware encryption, you manage the keys, supports asymetrical encryption
- AWS STS: Generates security tokens
- Identity Federations: SAML 2.0 or Custom for Enterprise, Cognito for app
- AWS Artifact: Get access to compliance reports
- AWS Config (not in chapter): Track config changes and compliance against rules
- AWS CloudTrail (not in chapter): Track API calls made by users within an account
 

### Quiz
- Remember the, the report reporting who has MFA enabled is under IAM!
- Read more JSON policies, they can be difficult to understand!