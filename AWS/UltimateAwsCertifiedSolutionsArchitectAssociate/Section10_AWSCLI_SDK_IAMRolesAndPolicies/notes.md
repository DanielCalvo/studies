### 102. Developing on AWS Introduction
- AWS CLI: Exam asks questions!
- AWS Tasks can be performed in several ways
    - Using the AWS CLI from your local computer
    - Using the AWS CLI from an EC2 machine
    - Using the AWS SDK from your local computer
    - Using the AWS SDK from an EC2 machine
    - Using the AWS Instance Metadata Service for EC2

### 103. AWS CLI Set up on Windows

### 104. AWS CLI Set up on Mac

### 105. AWS CLI Set up on Linux

### 106. AWS CLI Configuration
- Do not share access keys or secret keys!
- IAM > Your User > Security Credentials > Create Access key
- Create Access key
- On your local: `aws configure`

### 107. AWS CLI on EC2

#### AWS CLI on EC2: The bad way
- You can run `aws configure` on EC2 like you did on your local
- This is super insecure. Do not do this!
- **Never put your personal credentials on EC2 instances!**
- Your personal credentials are meant to be used by you only
- If the instance is compromised, so is your personal account. If the instance is shared, other people might impersonate you.
- For EC2, there's a better way. It's called IAM roles

#### AWS CLI on EC2: The right way
- IAM Roles can be attached to EC2 instances
- IAM Roles can come with a policy authorizing exactly what the EC2 instance should be able to do
- The EC2 instances can use thes eprofiles automatically without any additional configuration
- This is the best practice on AWS and you should do this!

#### Practice time!
- Create EC2 instance with no IAM role by default
- IAM > Roles > Create Role > Attach to AWS Service > EC2 > S3 read only access
- EC2 instance > Right Click > Attach/Replace IAM Role. Attach role
- `aws s3 ls` now works on the instance. Neat!

### 108. AWS CLI Practice with S3
- You can read the docs for AWS CLI references
- `aws s3 ls`
- `aws s3 ls s3://dcalvodevbucket`
- `aws s3 cp help`
- `aws s3 cp bike.jpeg s3://dcalvodevbucket`
- `aws s3 mb s3://dcalvodevbucket2`
- `aws s3 rb s3://dcalvodevbucket2`
- Cool: https://docs.aws.amazon.com/cli/latest/reference/s3/

### 109. IAM Roles and Policies Hands on
- You can create your own policy with Service, Actions, Resources and Request conditions!
- There are inline policies but Stefan does not recommend using them
- You an also see the policy summary and JSON for a more detailed view
- AWS Policy Generator is pretty neat too
- You then attact the roles to your policy

### 110. AWS Policy Simulator
- Theres an AWS policy simulator: https://policysim.aws.amazon.com/home/index.jsp
- Helpful to test out policies!

### 111. AWS EC2 Instance Metadata
- Your EC2 instance can query metadata about itself without the need for an IAM role through a URL:
    -  http://169.254.169.254/latest/meta-data
- Dont forget:
    - Metadata: Info about the EC2 instance
    - Userdata: Launch script of the EC2 instance

### 112. AWS SDK Overview
- If you wanna use AWS from your program, you can use the SDK (software development kit )to do it, available in many languages! 
- SDK is useful when against certain AWS services such as DynamoDB
- The exam expects you to know when to use an SDK
- The SDK is used in this course when using Lambda functions
- Good to know: If you don't specify or configure a default region, us-east-1 is the default

### SDK Credentials security
- It's recommended to use the default credential provider chain. This works seamlessly with:
    - Credentials stored at ~/.aws/credentials (only on your computer or on premise)
    - Instance profile credentials using IAM Roles (like for EC2 machines)
    - Environment variables (aws_access_key_id, aws_secret_access_key)
- Never store AWS credentials in your code!

### Exponential back off
- When you make API calls that fail, these will be sent with exponential back off.