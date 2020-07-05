### 197. Security Token Service (STS) Overview
- Allows you to grant limited and temporary access to AWS resources.
- A token is valid for up to one hour (if you want to reuse that token you must refresh it )
- The most important functionality of the STS API is **AssumeRole**
    - You can use AssumeRole Within your own account: for enhanced security
    - Or you can assume a role in cross Account Access: assume role in target account to perform actions there
- AssumeRoleWithSAML
    - Used to return credentials for users logged with SAML
- AssumeRoleWithWebIdentity
    - Used to return credentials for users logged with an IdP (Facebook Login, Google Login, OIDC compatible...)
    - AWS recommends against using this, and using Cognito instead
- GetSessionToken
    - For MFA, from a user or AWS account root user

#### Using STS to Assume a role
- First define an IAM Role within your account or cross-account
- Then you define which principals (which users or which other roles) can access this IAM Role
- Use AWS STS (Security Token Service) to retrieve credentials and impersonate the IAM Role you have access to (AssumeRole API)
- Temporary credentials can be valid between 15 minutes to 1 hour

- You have a user that wants to assume a role (in the same or other account)
- User > AssumeRole API > AWS STS > Check permissions > IAM
- STS Then retorns temporary security credentials to the user

#### Cross account access with STS
1. Admin creates role that grants development account read/write access to productionapp bucket (on prod account)
2. Admin grants members of the group Developers permission to assume the UpdateApp role (on dev account)
3. Users request acces to role (from dev to prod)
4. STS returns role credentials
5. User updates productionapp (on prod) by using the role credentials
- Never directly share credentials across accounts!
    - You create a role and you and make sure your user assumes that role and STS will return to you role credentials that are temporary, so even if you lose them or they leak, it's not as bad as leaking entire credentials of an account
- Can't find the diagram in the lecture, but this seems to be an equivalent: 
    - https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html
- Every time you see cross account access or assuming roles, think STS!

### 198. Identity Federation & Cognito
- Federation lets users outside of AWS to assume temporary role for accessing AWS resources
- These users assume identity provided access role
- Federations can have many flavors:
    - SAML 2.0
- Custom Identity Broker
    - Web Identity Federation with Amazon Cognito
    - Web Identity Federation without Amazon Cognito
    - Single Sign On
    - Non-SAML with AWS Microsoft AD
- Using federation, you don’t need to create IAM users (user management is outside of AWS)
- User > login > 3rd party > AWS
- AWS trust the third party and lets user access it. User management is outside AWS

#### SAML 2.0 Federation
- Great integration with Active Directory / ADFS with AWS (or any SAML 2.0)
- Using SAMK 2.0 we can get access to the AWS Console or CLI (through temporary creds)
- No need to create an IAM user for each of your employees
- https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
- Client > Identity provider (on prem) > LDAP
- Identity provider returns a SAML assertion to the client
- Client calls AssumeRoleWithSAML STS API
- STS will look at the SAML assertion and make sure it is correct, and return temporary credentials
- And then you can use these credentials to access an S3 bucket (for ex)
- Same thing happens when we want to access the console!
 
#### SAML 2.0 Federation - Active Directory FS
- User > ADFS > Identity store > returns SAML assertion to user
- User > Posts the SAML assertion to sign in > Gets temp credentials from AWS > Redirected to console
- Same process really

#### SAML 2.0 Federation, part 2
- You need to set up a trust relationship between AWS IAM and SAML (both ways)
- SAML 2.0 enables web-based, cross domain SSO
- Uses the STS API: AssumeRoleWithSAML
- Note federation through SAML is the “old way” of doing things
- Amazon Single Sign On (SSO) Federation is the new managed and simpler way (another AWS service)

#### Custom Identity Broker Application
- Use only if your identity provider is not compatible with SAML 2.0
- The identity broker must determine the appropriate IAM policy for your user 
- Uses the STS API: AssumeRole or GetFederationToken
- Overall process similar, but this time the identity broker talks to STS directly (no SAML)

#### Web Identity Federation - AssumeRoleWithWebIdentity
- For users over applications to use AWS (?)
- Not recommended by AWS – use Cognito instead (allows for anonymous users, data synchronization, MFA)

#### AWS Cognito
- Goal:
    - Provide direct access to AWS Resources from the Client Side (mobile, web app)
- Example:
    - provide (temporary) access to write to S3 bucket using Facebook Login
- Problem:
    - We don’t want to create IAM users for our app
- How:
    - Log in to federated identity provider – or remain anonymous
    - (from the cognito API) Get temporary AWS credentials back from the Federated Identity Pool
    - These credentials come with a pre-defined IAM policy stating their permissions
- App > Login to identity provider (FB, Cup, Google, etc) > Get back a token
- App > Authenticate to FIP (with Token) > Federated Identity > Verify token > Get credentials from STS > Give back to app tmp credentials
- App can now access S3 (for example)

### 199. [SAA-C02] Directory Services - Overview
- Found on any Windows Server with AD Domain Services
- It's a database of objects: User Accounts, Computers, Printers, File Shares, Security Groups
- Centralized security management, create account, assign permissions
- Objects are organized in trees
- A group of trees is a forest
- The idea is that you have a machine that is the domain controller, and you configure the user john in there
- Then all the other machines that have this machine as a domain controller, can login as john (after the credentials are verified on the controler)

#### AWS Directory Services
- 3 flavours, no deep dive needed, but understand the differences!
- AWS Managed Microsoft AD
    - Create your own AD in AWS, manage users locally, supports MFA
    - You can establish “trust” connections with your on-premise AD
- AD Connector
    - Directory Gateway (proxy) to redirect to on-premise AD
    - Users are managed on the on-premise AD
- Simple AD
    - AD-compatible managed directory on AWS (not MS)
    - Cannot be joined with on-premise AD
    - Can be useful if you don't have an on-premise AD
- Exam might ask very high level questions (we want to do x, what service do we use?)

### 200. Organizations - Overview
- Multi accounts!
- It's global service
- Allows you to manage multiple AWS accounts
- The main account is the master account – you can’t change it
- Other accounts within the org are member accounts
- Member accounts can only be part of one organization
- Consolidated Billing across all accounts - single payment method
- Pricing benefits from aggregated usage (volume discount for EC2, S3...)
- API is available to automate AWS account creation

#### Multi Account strategies
- Create accounts per department, per cost center, per dev / test / prod, based on regulatory restrictions (using SCP), for better resource isolation (ex: VPC), to have separate per-account service limits, isolated account for logging
- There's a difference between multi Account vs One Account Multi VPC
    - There's a chance on one big account that users can accidentally have access to VPCs that they shouldn't
- Use tagging standards for billing purposes
- You can enable CloudTrail on all accounts, send logs to central S3 account
- You also can send CloudWatch Logs to central logging account
- And you can establish Cross Account Roles for Admin purposes

#### Organizational units (OU) - Examples
- Business Unit (Sales, Retail, Finance)
- Environment lifecycle (Dev, prod, stage)
- Project based (project1, 2, 3)
- You can have OUs as children/inside other OUs

#### Service Control Policies
- Allows you Whitelist or blacklist IAM actions
- Applied at the Root, OU or Account level
- Does not apply to the Master Account
- SCP is applied to all the Users and Roles of the Account, including Root
    - So if you restrict EC2 usage on a OU, even root can't use EC2 there
- The SCP does not affect service-linked roles
- Service-linked roles enable other AWS services to integrate with AWS Organizations and can't be restricted by SCPs.
- SCP must have an explicit Allow (does not allow anything by defa1ult)
- Use cases:
    - Restrict access to certain services (for example: can’t use EMR)
    - Enforce PCI compliance by explicitly disabling services

#### SCP Hierarchy
- SCPs do not apply to the master account
- Denies take precedences over authorizes
- SCPs are inherited by accounts down the organization tree

#### SCP Examples blacklist and whitelist strategies
- They look very much like json policies 

#### AWS Organizations - Moving accounts
To migrate accounts from one organization to another
1. Remove the member account from the old organization
2. Send an invite to the new organization
3. Accept the invite to the new organization from the member account

If you want the master account of the old organization to
also join the new organization, do the following:
1. Remove the member accounts from the organizations using procedure above
2. Delete the old organization
3. Repeat the process above to invite the old master account to the new org

### 201. Organizations - Hands on
- AWS Organizations
- Stephane created a new account named aws account master
- Create organization
- You can then invite aws accounts to join your organization
    - Actually you can invite an account and create an account
- Inviting account: By email or account ID
- On the child account, go on organizations and you should have an invitation there
- On the master account, you now should be able to see both accounts on the same organization
- You can then go to the organize accounts and create OUs!
    - By default you only have root
- You can create new OUs inside root, like dev and test
- And inside those OUs you can add more OUs! 
- You can then select an account and move it to the OU you want!
- It's good pratice to leave your master account on the root OU but you can also move it if you want
- On root, you can enable service control policies
- Once they're enabled, you can attach policies to your accounts inside your OUs
- FullAWSAccess is inherited from root
- Created a policy to deny access to Athena, all actions
- Neat, you can't access athena

### 202. IAM Advanced

#### IAM Conditions
- A way to make your IAM policies a bit more restrictive based on a condition
- aws:SourceIP: restrict the client IP from which the API calls are being made
    - Ex: Deny anything that doesn't come from these IPs
- Aws:RequestedRegion: restrict the region the API calls are made to
    Ex: You can only allow certain services on certain regions
- You can also restrict on tags!
    - Ex: You can only stop and start instances that have "project:banana" on it
    - ec2:ResourceTag 
- You can also force MFA for certan actions, like stopping instances

#### IAM for S3
- ListBucket permission applies to
    - arn:aws:s3:::test
    - bucket level permission
- GetObject, PutObject, DeleteObject applies to
    - arn:awn:s3:::test/*
    - object level permission

#### IAM Roles vc Resource Based Policies
- Attach a policy to a resource (example: S3 bucket policy) versus attaching of a using a role as a proxy
- What's the difference of using an S3 policy vs an IAM role?
- When you assume a role (user, application or service), you give up your original permissions and take the permissions assigned to the role
- When using a resource based policy, the principal doesn’t have to give up his permissions
- Example: User in account A needs to scan a DynamoDB table in Account A and dump it in an S3 bucket in Account B.
- Supported by: Amazon S3 buckets, SNS topics, SQS queues (can have resource based policies)

### 203. [SAA-C02] IAM - Policy Evaluation Logic

#### IAM Permission Boundaries
- IAM Permission Boundaries are supported for users and roles (not available for groups)
- Advanced feature to use a managed policy to set the maximum permissions an IAM entity can get
- IAM Permission Boundary: It's an IAM policu that you attach to a use or role and it can only do those things.
    - You can attach more roles to that user but it doesn't matter, it will still be bound by the IAM permission boundary

#### Hands on!
- IAM > users > Add user
- Programmatic access and defaults for all the rest
- On the user: Set boundary
- Gave him S3 admin access and now that's all he can do!

#### IAM Permission Boundaries, part 2
- Can be used in combinations of AWS Organizations SCP
- Use cases for permission boundaries:
    - Delegate responsibilities to non administrators within their permission boundaries, for example create new IAM users
    - Allow developers to self-assign policies and manage their own permissions, while making sure they can’t “escalate” their privileges (= make themselves admin)
    - Useful to restrict one specific user (instead of a whole account using Organizations & SCP)

#### IAM Policy Evaluation Logic
- That's a big diagram
- Big diagram is here: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html
- Deny evaluation > Organizations SCPs > Resource-based policies > IAM permissions boundaries > Session Policies (STS) > Identity based policies

#### Example IAM Policy
- Has a deny for all SQS resources
- And then an allow for sqs:Delete queue
- You can't create a queue because there's a deny for sqs:*
- As soon as you have an explicit deny, then the decision is going to be denied
    - You also can't delete the queue, because the deny prevails over the allow
- You also can't perform ec2:Describe instances as there is no allow on your policy, and a deny is implied for everything that is not allowed

### 204. [SAA-C02] Resource Access Manager (RAM)
- Allows you to share AWS resources that you own with other AWS accounts
- You can share with any account or within your Organization
- The idea is that you avoid resource duplication!
- You can share:
- VPC Subnets:
    - allow to have all the resources launched in the same subnets
    - Must be from the same AWS Organizations.
    - Cannot share security groups and default VPC
    - Participants can manage their own resources in there
    - Participants can't view, modify, delete resources that belong to other participants or the owner
- You also can share:
    - AWS Transit Gateway
    - Route53 Resolver Rules
    - License Manager Configurations

#### Resource Access Manager - VPC example
- Each account...
    - Is responsible for its own resources
    - Cannot view, modify or delete other resources in other accounts
- Network is shared so...
    - Anything deployed in the VPC can talk to other resources in the VPC
    - Applications are accessed easily across accounts, using private IP!
    - Security groups from other accounts can be referenced for maximum security
- Dani: You can have two accounts deploying things on the same private subnet. Noice!

#### Hands on!
- Resource access manager
- And go on share something and man there's a lot of stuff you can share, subnets too!
    - Subnets can't be on the default VPC though
- Then you can share that with another account or OU
- Exam tip: If you see a VPC subnet, think Resource Access Manager

### 205. [SAA-C02] AWS Single Sign On (SSO) - Overview
- This is to centrally manage Single Sign-On to access multiple accounts and 3rd party business applications
    - You go to a portal and once you're signed on to that single sign on portal, you can login to any of your AWS accounts and dropbox and office365 and slack and so on. 
- Integrated with AWS Organizations
    - So if you have a ton of accounts on your org, you just set up SSO and you will have access to login into all the accounts within the org. One login for all the accounts!
- Supports SAML 2.0 markup
- Integration with on-premise Active Directory
- Centralized permission management
- Centralized auditing with CloudTrail
- Exam tip: When you see a use case talking about doing a sign on to multiple AWS accounts, or to business applications that require SAML 2.0, think single sign on.

#### AWS Single Sign on (SSO) - Setup with AD
- AWS SSO can connect to on premise AD to manage user 
- But you can also use MS Managed AD on AWS
- SSO gets the users from AD
- Once logged in with SSO you can give access to AWS consoles, business cloud apps (slack, dropbox) and custom SAML applications

#### SSO vs AssumeRoleWithSAML
- AssumeRoleWithSaml
    - Browser > 3rd parrty IDP login portal > identity store > returns SAML
    - Browser with SAML > STS > And with STS creds > AWS 
- AWS SSO
    - Browser > Login on AWS SSO Login Portal > Identity Store SAML 2.0 compatible
    - Browser > AWS
- SSO is better for multiple accounts 

### 206. [SAA-C02] AWS Single Sign On (SSO) - Hands on
- AWS Sigle sign on
- Can enable SSO, and the steps after that are
1. Choose your identity source
2. Manage SSO access to your AWS accounts
3. Manage SSO access to your cloud applications (anything SAML 2.0 based)
- You get a user portal, neat!
- You can have SSO as an identity source, AD, or an External Identity provider

### Quiz!
- Pending