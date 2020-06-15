### 151. Serverless introduction

#### What's serverless
- Serverless is a new paradigm in which the developers don't have to manage servers anymore
- You just deploy code
- You just deploy... functions!
- Initially, Serverless == FaaS (functions as a service)
- This was pioneered by AWS Lambda but now also includes anything that is managed (databases, messaging, storage, etc)
- Serverless in AWS usually means Lambda
- Serverless doesn't mean there are no servers! You just don't manage, see or provision them

#### Serverless in AWS
- AWS Lambda & Step functions
- Dynamo DB
- AWS Cognito
- AWS API Gateway
- Amazon S3
- AWS SNS & SQS
- AWS Kinesis
- Aurora Serverless

### 152. Lambda Overview
- Amazon EC2
    - Virtual servers in the cloud
    - Limited by RAM and CPU
    - Continuosly running
    - Scaling means you have to add / remove servers
- Amazon Lambda
    - Virtual functions - no servers to manage!
    - Limited by time - short executions
    - Run on-demand
    - Scaling is automated

#### Benefits of AWS Lambda
- Easy pricing!
    - Pay per request and compute time
    - Free tier of 1,000,000 AWS Lambda requests and 400,000 GBs of compute time
- Integrated with the whole AWS stack
- Integrated with many programming languages
- Easy to monitor with AWS CloudWatch
- Easy to get more resources per function (up to 3GB of ram)
- Increasing RAM will also improve CPU and network!

#### AWS Lambda Language support
- Node.js, Python, Java, C#, Golang and many others

#### AWS Lambda Integrations (Most Used / Main Ones)
- API Gateway, Kinesis, DynamoDB, S3, IoT, CloudWatch Events & Logs, SNS, Cognito and SQS

#### Example: Serverless Thumbnail creation
1. New image is uploaded to S3
2. This triggers an AWS Lambda function to create a thumbnail
3. This function can then push the thumbnail to S3
3. This function can also push image metadata to DynamoDB (image size, name, creation date, etc)

#### Example: Serverless cronjob
- CloudWatch events can emit events based on a cron schedule. 
- You can use this to every 1 hour trigger an AWS Lambda Function to perform a task.

#### Pricing
- Details here: https://aws.amazon.com/lambda/pricing/
- Pay per calls
    - First 1,000,000 requests are free
    - 0.20 per 1 million requests afterwards
- Pay per duration (in 100ms increments)
    - 400,000 GB-Seconds of compute time per month
    - This is 400,000 seconds if your function is using 1GB of ram
    - This is 3,200,000 seconds if your function is using 128MB of ram
    - After that you pay 1.00 dollar for 600,000 GB- seconds
- It is usually cheap to run AWS Lambda, so it's very popular

### 153. Lambda Hands on
- Read the docs here and create and understand a golang app: https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html 