### Notes from Dani:
- Amazon DynamoDB is a fully managed proprietary NoSQL database service that supports key-value and document data structures
- Man it would be really nice to develop something in serverless manner as described on the diagram on the 157 lecture

### 157. Serverless introduction

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
- Step Functions
- Fargate

### 158. Lambda Overview
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

### 159. Lambda Hands on
- Lamba from the main console!
- Hmm nice, a golang example
- Lambda can react to events which seems pretty noice. More events, more lambda functions run!
- Create application
- Use blueprint
- hello-world in python
- Create a new role with basic lambda permissions
- Created a test event and tested and it worked!
- You can see the logs on cloudwatch, neat
- You can use cloudwatch logs to debug what goes wrong and each lambda invocation will create a cloudwatch stream
- On the permissions tab you can see the role that was created that allows lambda to write to cloudwatch.

### 160. Lambda Limits 
- Exam likes to see if you know these limits

Execution:
    - Memory allocation: 128 MB (min) 3008 MB (max) (64 MB increments). Increasing memory also increases vCPU
    - Maximum execution time: 900 seconds (15 minutes)
    - Environment variables:  (4 KB)
    - Disk capacity in the “function container” (in /tmp): 512 MB
    - Concurrency executions: 1000 (can be increased)
- Deployment:
    - Max size for lambda function deployment size allowed (compressed .zip): 50 MB
    - Max size size of uncompressed deployment allowed (code + dependencies): 250 MB
    - Can use the /tmp directory to load other files at startup
    - Max size of environment variables: (4 KB)
- Exam might say something like: We need 6gb of ram or 30 minutes of run time or a big file of 3gb, then you'll know lambda is not the right way to run this workload

### 161. Lambda@Edge 
- You have deployed a CDN using CloudFront
- What if you wanted to run a global AWS Lambda alongside?
- Or how to implement request filtering before reaching your application?
- For this, you can use Lambda@Edge:
    - Deploy Lambda functions alongside each reagion around the world with your CloudFront CDN
    - This allows you to build more responsive applications
    - You don’t manage servers, Lambda is deployed globally
    - You can customize the CDN content
    - And you'll pay only for what you use

#### Lambda@Edge part 2
- You can use Lambda to change CloudFront requests and responses
- There are four types of lambda functions at the dge: 
    - After CloudFront receives a request from a viewer (viewer request)
    - Before CloudFront forwards the request to the origin (origin request)
    - After CloudFront receives the response from the origin (origin response)
    - Before CloudFront forwards the response to the viewer (viewer response)
- You can also generate responses to viewers without ever sending the request to the origin

#### Lambda@Edge: Global
- Static website on S3 bucket with JS > Dynamic API request > CloudFront > Lambda@Edge function > Amazon DynamoDB

#### Lambda@Edge: Use cases
- Website security and private
- Dynamic Web Application at the edge
- SEO
- Intelligently route across origins and data centers
- Bot Mitigation at the edge
- Real time Image Transformation
- A/B testing
- User Authentication and Authorization
- User priorization
- User tracking and analytics

- Lambda edge: Whenever you need to integrate a lambda function with your cloudfront distribution to modify the 4 types of requests seen above

### 162. DynamoDB overview
- Fully Managed, highly available with replication across 3 AZ
- NoSQL database - not a relational database
- Scales to massive workloads, distributed database
    - From an architect point of view you don't really need to know how it works, just what it is 
- Can handle millions of requests per seconds, trillions of row, 100s of TB of storage
- Fast and consistent in performance (low latency on retrieval)
- Integrated with IAM for security, authorization and administration
- Enables event driven programming with DynamoDB Streams
- Low cost and has auto scaling capabilities

#### DynamoDB Basics
- DynamoDB is made of tables
- Each table has a primary key (must be decided at creation time)
- Each table can have an infinite number of items (= rows)
- Each item has attributes (can be added over time – can be null) (atribute is like a column in the RDS world)
- Maximum size of a item is 400KB
- Data types supported are:
    - Scalar Types: String, Number, Binary, Boolean, Null
    - Document Types: List, Map
    - Set Types: String Set, Number Set, Binary Set

#### DynamoDB - Provisioned Throughput
- Important to know going into the exam!
- A Table must have provisioned read and write capacity units
- Read Capacity Units (RCU): throughput for reads ($0.00013 per RCU) (the unit of provisioning for DynamoDB)
    - 1 RCU = 1 strongly consistent read of 4 KB per second
    - or
    - 1 RCU = 2 eventually consistent read of 4 KB per second
- Write Capacity Units (WCU): throughput for writes ($0.00065 per WCU)
- 1 WCU = 1 write of 1 KB per second
- Option to setup auto-scaling of throughput to meet demand
- Throughput can be exceeded temporarily using “burst credit”
- If burst credit are empty, you’ll get a “ProvisionedThroughputException”.
- It’s then advised to do an exponential back-off retry

### 163. DynamoDB Hands on
- Let's have a quick look at it, not something you need to know in depth as far as how to use it
- Console > DynamoDB > Create table
- Table name: demo
- Key: username
- Read/write capacity mode can be:
    - Provisioned (free-tier eligible)
    - On-demand
- Oh there's a capacity calculator with costs!
- You can enable or disable auto scaling too
- Can have encryption at rest
- Create table!
- On the table
    - Items
    - Create item
    - Add a few things!
- You can't add two items with the same key

### 164. [SAA-C02] DynamoDB Advanced Features
#### DynamoDB - DAX
- DAX = DynamoDB Accelerator
- Seamless cache for DynamoDB, no application re-write
- Writes go through DAX to DynamoDB
- The reads will be cached on DAX!
- Micro second latency for cached reads & queries
- Solves the Hot Key problem (too many reads on one value in DynamoDB)
- Each cache entry has a 5 minutes TTL by default
    - Every time you read something from DynamoDB it will be cached for 5 minutes in DAX
- Up to 10 nodes in the cluster
- Multi AZ (3 nodes minimum recommended for production)
- Secure (Encryption at rest with KMS, VPC, IAM, CloudTrail...)

#### DynamoDB Streams
- Changes in DynamoDB (Create, Update, Delete) can end up in a DynamoDB Stream
- This stream can be read by AWS Lambda, and we can then do:
    - React to changes in real time (welcome email to new users)
- Analytics
    - Create derivative tables / views
    - Insert into ElasticSearch
- Could implement cross region replication using Streams
- Stream has 24 hours of data retention
- Streams are mostly used with Lambda functions to take action on data changes that happen on DynamoDB

#### DynamoDB - New features
- May not be on the exam or may show up soon!
- Transactions (new from Nov 2018)
    - All or nothing type of operations
    - Either all of these things work or none of them work
    - You can coordinated Insert, Update & Delete across multiple tables
    - You can include up to 10 unique items or up to 4 MB of data
- On Demand (new from Nov 2018)
    - No capacity planning needed (WCU / RCU) – scales automatically
    - 2.5x more expensive than provisioned capacity (use with care)
    - Helpful when spikes are un-predictable or the application is very low throughput that it's actually cheaper to use on deman

#### DynamoDB - Security & Other features
- Security:
    - VPC Endpoints available to access DynamoDB without internet
    - Access fully controlled by IAM
    - Encryption at rest using KMS
    - Encryption in transit using SSL / TLS
• Backup and Restore feature available
    - Point in time restore like RDS
    - No performance impact
- Global Tables
    - Multi region, fully replicated, high performance
- Amazon DMS can be used to migrate to DynamoDB (from Mongo, Oracle, MySQL, S3, etc...)
- You can launch a local DynamoDB on your computer for development purposes

#### DynamoBD - Other features
- Global Tables: (cross region replication)
    - Active Active replication, many regions (all regions will replicate to other ones)
    - Must enable DynamoDB Streams before you get that
    - Useful for low latency, DR purposes
- Capacity planning:
    - Planned capacity: provision WCU & RCU, can enable auto scaling
    - On-demand capacity: get unlimited WCU & RCU, no throttle, more expensive

### 165. API Gateway Overview
- Amazon offers API Gateway as a way for you to create REST APIs that can, among other things, use Lambda as a backend
- Client > Rest API > API Gateway > Proxy Requests > Lambda > CRUD > DynamoDB

#### AWS API Gateway
- AWS Lambda + API Gateway: No infrastructure to manage
- API Gateway support for the WebSocket Protocol (real time streaming)
- API Gateway can handle API versioning (v1, v2...)
- API Gateway handles different environments (dev, test, prod...)
- Several security features (such as with Authentication and Authorization)
- You can create API keys, handle request throttling
- Common standards such as Swagger / Open API can be used to import to quickly define APIs
- You can transform and validate requests and responses at the API Gateway level
- Can generate SDK and API specifications
- Can cache API responses
- An application doesn't have all these features, neat 

#### API Gateway - Integrations at a high level
- Lambda Function
    - Can invoke Lambda function
    - Easy way to expose REST API backed by AWS Lambda to have a serverless application
- HTTP
    - Expose HTTP any endpoints in the backend 
    - Example: internal HTTP API on premise, Application Load Balancer...
    - Why? Add rate limiting, caching, user authentications, API keys, etc... (any API gateway feature)
- AWS Service
    - You can Expose any AWS API through the API Gateway
    - Example: start an AWS Step Function workflow, post a message to SQS directly through an API gateway (neat)
    - Why? Add authentication, deploy publicly, rate control... any other API gateway feature

#### API Gateway - Endpoint types
- There are 3 ways to deploy your API gateway, these are called endpoint types
- Edge-Optimized (default): For global clients
    - Requests are routed through the CloudFront Edge locations (improves latency)
    - The API Gateway still lives in only one region but it's acessible efficiently through every cloudfront location 
- Regional:
    - For clients within the same region as the API Gateway (no cloudfront edge locations)
    - Could manually combine with CloudFront (more control over the caching strategies and the distribution)
- Private:
    - Can only be accessed from your VPC using an interface VPC endpoint (ENI)
    - Use a resource policy to define access


### 166. API Gateway Basics Hands-On
- Console > API Gateway
- API Types:
    - HTTP API: Build low-latency and cost-effective REST APIs with built-in features such as OIDC and OAuth2, and native CORS support.
    - WebSocket API: Build a WebSocket API using persistent connections for real-time use cases such as chat applications or dashboards.
    - REST API: Develop a REST API where you gain complete control over the request and response along with API management capabilities.
    - REST API (Private): Create a REST API that is only accessible from within a VPC.

- Let's go for REST API
- Instead of the example API, let's go for a new API
    - REST
    - New API
    - Regional endpoint
- Create method
    - GET at /
- For this get:
    - Integration type: Lambda Function
    - Use Lambda function: yes
- But before we save we need to create a lambda function!
- Create a function from scratch with python 3.8 as the runtime. Add this as code:
```python
import json

def lambda_handler(event, context):
    body = "Hello from Lambda!"
    statusCode = 200
    return {
        "statusCode": statusCode,
        "body": json.dumps(body),
        "headers": {
            "Content-Type": "application/json"
        }
    }
``` 
- Go back to API Gateway and click on save! (a)
    - This creates in the background a resource policy for your API gateway to invoke your lambda function
- Let's create a second method!
    - Action
    - Create resource
    - /houses
    - Now let's create a new lambda function to associate with this endpoint
    - Runtime is python 3.8
    - Just linked it together and now I have another endpoint that responds to a get

- Now lets make this public!
    - Action
    - Deploy API
    - Created a stage with dev on all fields

### 167. API Gateway Security
#### IAM Permissions
- I think Stephane is talking about this in the context of allowing peole to use the API gateway
- API Gateways are attached to an IAM policy that provides authorization 
- API Gateway verifies IAM permissions when you call your REST API
- Good to provide access within your own infrastructure
- IAM credentials are in a header (?). Sig v4 is a capability that is leveraged where IAM credentials are in headers
- Everytime you see Sig v4 in the exam think IAM permissions for API gateway
- Client > Calls API Gateway with Sig v4 > API Gateway calls IAM to verify this policy and if it's happy, makes the request to the backend

#### Lambda Authorizer (formerly Custom Authorizers)
- Uses AWS Lambda to validate the token that is being passed in the header of your request
- You can cache the result of your authentication, you don't have to call lambda authorizer every time
- Helps to use OAuth / SAML / 3rd party type of authentication
- Lambda must return an IAM policy for the user

- Client > REST API with Token > Amazon API Gateway > Evaluates request Token on Lamda Authorizer
- If all good, API gateway reaches the backend for the response and all good

#### Cognito User Pools
- Cognito fully manages user lifecycle
- API gateway verifies identity automatically from AWS Cognito
- No custom implementation required
- Cognito only helps with authentication, not authorization

- Client first talk to Cognito User pools to authenticate and receive a token
- Client then talks to the REST API from Api Gateway with the Pass token it received from Congnito
- API GW will make sure the token is correct by talking to cognito
- If all is well, API GW reaches for the backend to perform the operation 

#### API Gateway - Security - Summary
- IAM:
    - Great for users / roles already within your AWS account
    - Handles authentication + authorization though IAM policies
    - Leverages Sig v4
- Custom Authorizer:
    - Great for 3rd party tokens
    - Very flexible in terms of what IAM policy is returned
    - Handles Authentication + Authorization
    - Pay per Lambda invocation (but you can use caching)
- Cognito User Pool:
    - You manage your own user pool (can be backed by Facebook, Google login etc...)
    - No need to write any custom code
    - Must implement authorization in the backend

### 168. AWS Cognito Overview
- Stephane claims: The most confusing AWS out there :joy:

- Cognito is used when te want to give our users an identity so that they can interact with our application
- Cognito has a few subproducts
- Cognito User Pools:
    - Sign in functionality for app users
    - Integrates with API Gateway
- Cognito Identity Pools (Federated Identity):
    - Provide AWS credentials to users so they can access AWS resources directly
    - Integrates with Cognito User Pools as an identity provider
- Cognito Sync:
    - Synchronize data from a device to Cognito
    - May be deprecated and replaced by AppSync

#### AWS Cognito User Pools (CUP)
- Create a serverless database of user for your mobile apps
- Simple login: Username (or email) / password combination
- Possibility to verify emails / phone numbers and add MFA
- Can enable Federated Identities (Facebook, Google, SAML...)
- Sends back a JSON Web Tokens (JWT)
- Can be integrated with API Gateway for authentication

#### AWS Cognito User Pools (CUP)
- A serverless database of user for your mobile apps
- A Simple login: Username (or email) / password combination
- Possibility to verify emails / phone numbers and add MFA
- Can enable Federated Identities (Facebook, Google, SAML...)
- Sends back a JSON Web Tokens (JWT)
- Can be integrated with API Gateway for authentication

- App > Register/Login > CUP
- Cup returns JWT

#### AWS Cognito - Federated Identity Pools
Goal:
    - To provide direct access to AWS Resources from the Client Side (no proxy, no API, just straight access)
- How:
- Log in into a federated identity provider – or remain anonymous
- From this we get temporary AWS credentials back from the Federated Identity Pool
    - These credentials come with a pre-defined IAM policy stating their permissions
- Example:
    - Provide (temporary) access to write to S3 bucket using Facebook Login

- App > Logins (FB/Google/CUP) > Gets a token
- App > Authenticate to FIP (with Token) > Federated identity
- Once the identity is verified the Federated identity service will talk to the STS to get temporary AWS credentials
- Once it has that it will pass those temporary credentials to the application
- With these credentials the app can for instance interact with a S3 bucket (and there's an IAM policy that can allow you to do only certain things)
- This is more to allow you to interact with AWS being an app user

#### AWS Cognito Sync
- Deprecated - Use AWS Appsync Now
- Store user preferences, configuration, state of app
- Has cross device synchronization capability (any platform – iOS, Android, etc...)
- Offline capability (synchronization when back online)
- Requires Federated Identity Pool in Cognito (not User Pool)
- Data is stored in datasets (up to 1MB)
- You can have up to 20 datasets to synchronise

### 169. [SAA-C02] Serverless Application Model (SAM) Overview
- SAM = Serverless Application Model
- It's a framework for developing and deploying serverless applications
- All the configuration of your serverless application is YAML code
- Through the SAM framework, you can configure:
    - Lambda Functions
    - DynamoDB tables
    - API Gateway
    - Cognito User Pools
- SAM will help you deploy all of that
- SAM can help you to run Lambda, API Gateway, DynamoDB locally
- SAM can use CodeDeploy to deploy Lambda functions

### Quiz!
- Pending