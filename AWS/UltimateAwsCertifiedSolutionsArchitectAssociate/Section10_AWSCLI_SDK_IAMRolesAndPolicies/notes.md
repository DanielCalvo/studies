### 106. Developing on AWS Introduction
- So far, we've interacted with services manually and they exposed standard information for clients:
    - EC2 exposes a standard Linux machine we can use any way we want
    - RDS exposes a standard database we can connect to using a URL
    - ElastiCache exposes a cache URL we can connect using a URL
    - ASG / ELBs are automated and we don't have to program them
- Developing against AWS has two components:
    - How to perform interactions with AWS without using the online console?
    - How to interact with AWS proprietary services? (S3, DynamoDB, etc...)
- Developing and performing AWS tasks against AWS can be done in several ways:
    - Using the AWS CLI from your local computer
    - Using the AWS CLI from an EC2 machine
    - Using the AWS SDK from your local computer
    - Using the AWS SDK from an EC2 machine
    - Using the AWS Instance Metadata Service for EC2
- In this section, we'll learn
    - How to do all of those
    - The right & most secure way, adhering to best practices

### 107. AWS CLI Set up on Windows

### 108. AWS CLI Set up on Mac

### 109. AWS CLI Set up on Linux

### 110. AWS CLI Configuration
- Let's learn how to properly configure the CLI
- Computer > CLI access over WWW > AWS Account (checks credentials and permissions)
- We'll learn how to get our access credentials and how to protect them
- Do not share access keys or secret keys!

#### Hands on
- IAM > Your User > Security Credentials > Create Access key
- Create Access key
- On your local: `aws configure`

### 111. AWS CLI on EC2

#### AWS CLI on EC2: The bad way
- You can run `aws configure` on EC2 like you did on your local
- This is super insecure. **Do not do this!**
- **Never put your personal credentials on EC2 instances!**
- Your personal credentials are meant to be used by you only
- If the instance is compromised, so is your personal account. If the instance is shared, other people might impersonate you.
- For EC2, there's a better way. It's called IAM roles

#### AWS CLI on EC2: The right way
- IAM Roles can be attached to EC2 instances
- IAM Roles can come with a policy authorizing exactly what the EC2 instance should be able to do
- The EC2 instances can use these profiles automatically without any additional configuration
- This is the best practice on AWS and you should do this!

#### Practice time!
- Create EC2 instance with no IAM role by default
- IAM > Roles > Create Role > Attach to AWS Service > EC2 > S3 read only access
- EC2 instance > Right Click > Attach/Replace IAM Role. Attach role
- `aws s3 ls` now works on the instance. Neat!
- Remember that when you apply a policy things can take a minute or two to be effective (AWS has it's own internal mechanisms)

### 112. AWS CLI Practice with S3
- You can read the docs for AWS CLI references
- `aws s3 ls`
- `aws s3 ls s3://dcalvodevbucket`
- `aws s3 cp help`
- `aws s3 cp bike.jpeg s3://dcalvodevbucket`
- `aws s3 mb s3://dcalvodevbucket2`
- `aws s3 rb s3://dcalvodevbucket2`
- Cool: https://docs.aws.amazon.com/cli/latest/reference/s3/

### 113. IAM Roles and Policies Hands on
- You can attach an existing policy
- You can also create your own policy with Service, Actions, Resources and Request conditions!
- There are inline policies but Stefan does not recommend using them
- You an also see the policy summary and JSON for a more detailed view
- AWS Policy Generator is pretty neat too
- You then attact the roles to your policy

#### Hands on - creating your own policy
- You can also create your own IAM policy
- IAM > Access Management > Create Policy
- Just created a policy that allows you to only GetObject on danidanibucket
- AWS Policy generator is nice too: https://awspolicygen.s3.amazonaws.com/policygen.html 

### 114. AWS Policy Simulator
- Theres an AWS policy simulator: https://policysim.aws.amazon.com/home/index.jsp
- Helpful to test out policies!

### 115. AWS EC2 Instance Metadata
- AWS EC2 Instance Metadata is powerful but one of the least known features to developers
- It allows AWS EC2 instances to "learn about themselves" without using an IAM role for that purpose
-  The URL is: http://169.254.169.254/latest/meta-data
- This URL is internal and will only work from your EC2 instance. You can retrieve the IAM role name from the metadata but you cannot retrieve the IAM policy
- Dont forget:
    - Metadata: Info about the EC2 instance
    - Userdata: Launch script of the EC2 instance
    
### Hands on
- curl http://169.254.169.254/latest/meta-data/iam/info
- curl http://169.254.169.254/latest/meta-data/iam/security-credentials
- curl http://169.254.169.254/latest/meta-data/instance-id
- curl http://169.254.169.254/latest/meta-data/local-ipv4
- curl http://169.254.169.254/latest/meta-data/iam/security-credentials/MySecondEC2Role
    - That gives an AccessKeyId, a SecretAccessKey and a Token. Behind the scenes when you attach an IAM role to an EC2 instances, the instance queries this URL to get temporary credentials (on the background) 


### 116. AWS SDK Overview
- What if you want to perform actions on AWS directly from your applications code? (without using the CLI)
- You can use an SDK! (software development kit)
- There are official SDKs for Java, .NET, Node.js, PHP, Python (boto3/botocore, Go, Ruby, C++ and maybe some other languages
- The AWS CLI uses the Python SDK (boto3)
- We have to use the AWS SDK when coding against AWS services such as DynamoDB
- The exam expects you to know when to use an SDK, but doesn't expect you to know how to code or to actually use it
- The SDK is used in this course when using Lambda functions
- Good to know: If you don't specify or configure a default region, us-east-1 is the default

### AWS SDK Credentials security
- It's recommended to use the default credential provider chain. This works seamlessly with:
    - Credentials stored at ~/.aws/credentials (only on your computer or on premise)
    - Instance profile credentials using IAM Roles (like for EC2 machines)
    - Environment variables (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)
- Never store AWS credentials in your code!
- The best practive is for credentials to be inherited from the mechanisms above, and IAM roles if working from within AWS services

### Exponential back off
- Any API that fails because of too many calls needs to be retried with exponential backoff
- These apply only to rate limited APIs
- The SDK usually implements a retry mechanism with exponential backoff
- Exponential backoff: At each failed API call, the client waits will wait a double amount of time before attempting another call
- This ensures you don't overload the API by making calls every milisecond

### Quiz
- Pending