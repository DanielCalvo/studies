### 171. Mobile Application: MyTodoList
- We want to create a mobile application with the following requirements
    - Expose as REST API with HTTPS
    - Serverless architecture
    - Users should be able to directly interact with their own folder in S3
    - Users should authenticate through a managed serverless service
    - The users can write and read to-dos, but they mostly read them
    - The database should scale, and have some high read throughput

#### Mobile app: REST API layer
- Mobile client > Rest HTTPS > Invoke > AWS Lambda > Query > DynamoDB
- We can use Amazon Cognito to authenticate, and API Gateway will along the way verify the authentication with cognito
- Classic serverless architecture!

#### Mobile app: Giving users access to S3
- Mobile client > Authenticate > Amazon Cognito > Generate temp credentials > AWS STS
- Return these credentials to the mobile client
- These credentials can be used to store/retrieve files in S3
- Common question: How to do the above?
    - Wrong answer: To store AWS user credentials on your mobile client
    - Right answer: Use Amazon Cognito, STS and Amazon S3 with temporary credentials 

#### Mobile app: High read throughput, static data
- Turns out the TODO lists get read a lot more than they are modified
- How can we change the above architecture to improve the read throughput?
- We can use DAX as a cache layer between lambda and Dynamo DB!
- Reads will be cached in DAX 
- Will improve read output and reduce costs
- You can also start caching the responses at the Amazon Api GW level :o
    - If you think that the answers don't change as often

#### In this lecture
- We've seen the Serverless REST API architecture: HTTPS, API Gateway, Lambda, DynamoDB
- Using Cognito to generate temporary credentials with STS to access S3 bucket with restricted policy. App users can directly access AWS resources this way. Pattern can be applied to DynamoDB, Lambda...
- We are caching the reads on DynamoDB using DAX
- Or maybe we are caching the REST requests at the API Gateway level (if the answers are mostly static)
- Security for authentication and authorization using Cognito and STS

### 172. Serverless Website: MyBlog.com
- Our website should scale globally
- Blogs are rarely written, but often read
- Some  (most) of the website is purely static files, the rest is a dynamic REST API
- Caching must be implement where possible
- Any new users that subscribes should receive a welcome email
- Any photo uploaded to the blog should have a thumbnail generated
- Well, how do we go around implementing all of this?

#### Serving static content, globally
- Client > interacts with edge locations > Amazon CloudFront Global Distribution > Amazon S3

#### Serving static content, globally, securely
- Client > interacts with edge locations > Amazon CloudFront Global Distribution > Amazon S3
    - Amazon S3 has a OAI (origin access identity) from CloudFront to S3 (only cloudfront can read from S3)
- To add the public serverless API:
- Mobile client > Rest HTTPS > API Gateway > Invoke > AWS Lambda > Query/read > DAX Caching Layer > DynamoDB
- If you're going global, maybe you could leverage DynamoDB Global databases to reduce the latency

#### User welcome email flow
- But what about the user welcome email flow?
- Maybe on DynamoDB you want to enable DynamoDB Streams
- Flow is like this:
    - DynamoDB > Stream changes > DynamoDB Stream > Invoke lambda > Lamba > SDK to send email > Simple Email Service (SES)

#### Thumbnail Generation flow
- Client might upload to S3 bucket directly
- Or if there is OAI and a CloudFrount distribution
- Client > Upload photos > CloudFront > S3 > Lambda > S3 (thumbnail)

#### AWS Hosted Website Summary
- We’ve seen static content being distributed using CloudFront with S3
- The REST API was serverless, didn’t need Cognito because it was a public REST API
- We leveraged a Global DynamoDB table to serve the data globally (we could have used Aurora Global Tables)
- We enabled DynamoDB streams to trigger a Lambda function
- The lambda function had an IAM role which could use SES
- SES (Simple Email Service) was used to send emails in a serverless way
- S3 can trigger SQS / SNS / Lambda to notify of events

### 173.Microservices Architecture
- Not exactly serverless but still a very interesting discussion to have!
- We want to switch to a micro service architecture
- Many services interact with each other directly using a REST API
- Each architecture for each micro service may vary in form and shape
- We want a micro-service architecture so we can have a leaner development lifecycle for each service

#### Microservices Environment
- Stephane shows that microservices can have different architectures and don't have to follow the same rules (neat!)
- First microservice:
    - Users > HTTPS > ELB > ECS > DynamoDB
- Second microservice. Uses serverless architecture!
    - Users > HTTPS > Amazon API Gateway > AWS Lambda > ElastiCache
- Third service: Old school
    - Users > HTTPS > ELB > EC2 Auto Scaling > RDS

#### Discussions on Micro Services
- You are free to design each micro-service the way you want
- There are two patterns for microservices:
    - Synchronous patterns: API Gateway, Load Balancers
    - Asynchronous patterns: SQS, Kinesis, SNS, Lambda triggers (S3)
- Challenges with micro-services:
    - Repeated overhead for creating each new microservice,
    - Issues with optimizing server density/utilization
    - Complexity of running multiple versions of multiple microservices simultaneously
    - Proliferation of client-side code requirements to integrate with many separate services
- Some of the challenges are solved by Serverless patterns:
    - API Gateway, Lambda scale automatically and you pay per usage
    - You can easily clone API, reproduce environments
    - Generated client SDK through Swagger integration for the API Gateway

### 174. Distributing Paid Content
- We sell videos online and users have to paid to buy videos
- Each videos can be bought by many different customers
- We only want to distribute videos to users who are premium users
- We have a database of premium users
- We want to send links to the premium users that should be short lived
- Our application is global
- We want to be fully serverless

#### Staring simple, the premium user service
- Client > Rest HTTPS > API Gateway > invoke > AWS Lambda > Query/read > DynamoDB (has table of users and who is premium)

### Adding authentication
- Client > Amazon Cognito > Verifies authentication > Amazon API Gateway 

### Adding a video storage service (but globally and securely)
- Cloudfront access S3 with an OAI and a bucket policy (only cloudfront can reach the bucket!)
- We're going to use signed URLs so that only premium users can access the premium content
- Client > Authenticate > Cognito > Verify auth > API Gateway > Invoke Lambda to get signed url > Lambda > Verifies premium user > DynamoDB > If premium user, using the AWS SDK will generate a signed URL with cloudfront (on lambda step) > When done, returns response to API Gateway > Client get signed url 

#### Premium user video service
- We have implemented a fully serverless solution:
    - Cognito for authentication
    - DynamoDB for storing users that are premium
    - 2 serverless applications
        - Premium User registration
        - CloudFront Signed URL generator
- Content is stored in S3 (serverless and scalable)
- Integrated with CloudFront with OAI for security (users can’t bypass)
- CloudFront can only be used using Signed URLs to prevent unauthorized users
- What about S3 Signed URL? They’re not efficient for global access (cloudfront is more efficient to distribute content globally!)

### 175. Software update distribution
- We have an application running on EC2, that distributes software updates once in a while
- When a new software update is out, we get a lot of request and the content is distributed in mass over the network. It’s very costly
- We don’t want to change our application, but want to optimize our cost and CPU, how can we do it?

#### Current state
- Users > ALB > ASG > EFS

#### An easy way to fix things
- Super easy, just put CloudFront in front of the ASG!

#### Why CloudFront?
- No changes to architecture
- Will cache software update files at the edge
• Software update files are not dynamic, they’re static (never changing)
- Our EC2 instances aren’t serverless
- But CloudFront is, and will scale for us
- Our ASG will not scale as much, and we’ll save tremendously in EC2 and network costs and efs costs 
- We’ll also save in availability, network bandwidth cost, etc
- Easy way to make an existing application more scalable and cheaper! (if it's mostly static content)

### 176. Big Data Ingestion Pipeline
- We want the ingestion pipeline to be fully serverless
- We want to collect data in real time
- We want to transform the data
- We want to query the transformed data using SQL
- The reports created using the queries should be in S3
- We want to load that data into a data warehouse and create dashboards
- (usual big data problem of ingestion, collection, transformation, querying and analysis)

#### Diagram
- IoT Devices (managed with IoT Core) > Real time to Kinesis Data Streams > Kinesis Data Firehose > S3 ingestion bucket (offload every 1 minute)
    - It's possible to cleanse or transform the data using lambda linked to firehose
- S3 ingestion bucket > SQS (optional, can trigger lambda directly) > Lambda > Amazon Athena (SQL Query) > S3 p(ull data from the ingestion bucket) > S3 reporting bucket
- S3 reporting bucket > Amazon QuickSight (for visualization!) 
    - Or you can load the data in redshift for data warehousing (not serverless though)

#### Big Data Ingestion Pipeline Discussion
- IoT Core allows you to harvest data from IoT devices
- Kinesis is great for real-time data collection
- Firehose helps with data delivery to S3 in near real-time (1 minute)
- Lambda can help Firehose with data transformations
- Amazon S3 can trigger notifications to SQS
- Lambda can subscribe to SQS (we could have connected S3 to Lambda directly)
- Athena is a serverless SQL service and results are stored in S3
- The reporting bucket contains analyzed data and can be used by reporting tool such as AWS QuickSight, Redshift, etc...

### Quiz!
- Pending!