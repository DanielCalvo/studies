- It would be really neat to play with the patterns in this chapter (SNS and SQS fan out and ASG + SQS for processing queues)

### 143. Introduction to Messaging
- When you start deploy applications, they will need to communicate with one another
- There are two patterns for application communication:
    - Synchronous communication (app to app)
    - Asynchronous communication (application to queue to application)
- Synchronous
    - Buying service > Shipping service

- On asynchronous communication, there's a queue or middleware that connects both services
    - Buying service > Queue > Shipping service
    - Shippping service goes: "Hey is there something on the queue?" And retrieves that information
- Synchronous traffic between applications can be problematic if there are sudden spikes of traffic
- What if you suddenly need to enconde 1000 videos instead of the usual 10?

- You can decouple your applications with a decoupling layer
    - Using SQS: Queue model
    - Using SNS: Pub/sub model
    - Kinesis: real time streaming model (big data related)
- These services can scale independently from your application!

### 144. AWS SQS 
#### What's a queue?
- Very popular exam topic!
- There's an SQS queue
    - A SQS queue receives messages from producers. A producer sends a message to a SQS queue
    - You can have as many producers as you want sending messages to your queue
    - On the other hand, you will have a consumer. A consumer will pull messages from a SQS queue
    - You can have as many consumers as you want too 
    - A Queue service that receives messages from producers that can be polled by consumers

#### AWS SQS - Standard Queue
- SQS is old! About 10 years old
- Fully Managed
- Scales from 1 message per second to 10.000 to seconds
- Default message retention: 4 days, maximum of 14 days
- No limit to how many messages can be in the queue
- Low latency service (<10ms)
- You can have scaling terms of consumers and producers (or just consumers?)
- Can have duplicate messages (rare, with high throughput)
- Can have out of order messages (best effort ordering)
- Limit of 256kb per message sent

#### AWS SQS -  Delay Queue
- Delay a message (consumers don't see it immediately) up to 15 minutes
- Default is 0 seconds for standard queues
- Can set a default at queue level
- Or you can override the default using the DelaySeconds parameter when you sent data to SQS
- Producer > Send message > SQS (stays there on hold for 5 minutes) > Poll messages > Consumer

#### SQS - Producing Messages
- Define body (up to 256kb, has to be string)
- Add message attributes (metadada - optional)
- Provide Delay Delivery (optional)
- Message is sent to SQS!
- What we get back from SQS
    - Message identifier (aka an ID)
    - MD5 has of the body

#### SQS - Consuming Messages
- Consumers
- Poll SQS for messages (receive up to 10 messages at a time)
    - Consumers talk to SQS and go "Hey, do you have any messages for me?"
    - SQS goes: "Yeah I have x messages" (up to 10 messages at a time)
    - Or maybe SQS goes: "No new messages"
    - Consumers have the duty to process the message within the visibility timeout
    - Once consumers are done processin a message, they can delete it using the message ID & the receipt handle 

#### SQS - Visibility timeout
- When a consumer polls a message from a queue, the message is "invisible" to other consumers for a defined period. This is referred to as the "Visibility timeout"
- The visibility timeout can be set
    - Between 0 seconds and 12 hours (default 30 seconds)
    - If you set this visibility time out too high (say 15 minutes) and the consumer fails to process the message, you'll have to wait a long time before processing the message again
    - If too low (say 30 seconds) and the consumer needs time to process the message (2 minutes) another consumer will receive the message and the message will be processed more than once
- A consumer can change the timeout using the ChangeMessageVisibility API
    - ChangeMessageVisibility: API to change the visibility while processing a message
- DeleteMessage: API to tell SQS the message was successfully processed

#### AWS SQS - Dead Letter Queue
- If a consumer fails to process a message within the VisibilityTime out, the message goes back to the queue
- We can set a threshold of how many times a message can go back to the queue - it's called a "redrive policy"
- After the threshold is exceeded (there's something wrong with this message) it goes into a dead letter queue (DQL)
- We have to create a DLQ first and then designate it a dead letter queue
- Make sure to process the messages in the DLQ before they expire!

#### AWS SQS - Long Polling
- Popular at the exam!
- When a consumer requests messages from a queue, it can optinally "wait" for messages to arrive if there are none in the queue
- This is called Long Polling
- LongPolling decreases the number of API calls made to SQS while increasing the efficiency and latency of your application
- The wait time can be configured between 1 sec and 20 sec (20 sec preferable)
- Documentation says: Long polling os preferable to Short Polling
- Long polling can be enabled at the queue level or at the API level using WaitTimeSeconds

### 145. AWS SQS Console Hands On

#### Creating the queue 
- SQS Service
- Standard queue
- Configure queue
    - Default Visibility Timeout - The length of time (in seconds) that a message received from a queue will be invisible to other receiving components.
    - Message Retention Period - The amount of time that Amazon SQS will retain a message if it does not get deleted
    - Maximum Message Size - Maximum message size (in bytes) accepted by Amazon SQS
    - Delivery Delay - The amount of time to delay the first delivery of all messages added to this queue
    - Receive Message Wait Time - The maximum amount of time that a long polling receive call will wait for a message to become available before returning an empty response
- Dead Letter Queue settings (which queue is DLQ, maximum receives)
- Encryption settings

#### Sending and receiving messages
- Clicked on send message and sent myself a message. Neat!
- Right clicked the queue and clicked on "View/Delete messages"
    - Messages displayed in the console will not be available to other applications until the console stops polling for messages!
- You can see the message details on the UI

### 146. AWS SQS FIFO Queues
- Newer offering (First In - First out) - not available in all regions!
- The name of the queue must end in .fifo to make it clear that it is a FIFO queue
- Has lower throughput (up to 3,000 per second with batching, 300/s without)
- Messages are processed in order by the consumer
- Messages are sent exactly once
- You can't set a per message delay (only per queue delay)
- Ability to do content based de-duplication
- 5-minute interval de-duplication using “Duplication ID”
- Message Groups:
    - Possibility to group messages for FIFO ordering using “Message GroupID”
    - Only one worker can be assigned per message group so that messages are processed in order
    - Message group is just an extra tag on the message!
- Possible exam question: "We want to process messages in order, how do we set up a SQS queue for this?"
    - But exam doesn't go much into detail about FIFO queues

#### Hands on
- Select FIFO Queue when creating a queue
- Set up options are the same except now you can select "Content based deduplication"
    - SQS will create a hash of a given message to try to deduplacate it

### 147. SQS + Auto Scaling Group (ASG)
- There is an ASG with a bunch of instances that are pulling messages from an SQS queue
- What you want to do is to scale the number of EC2 instances based on how many messages you have on the SQS queue
- If more load on the queue, more instances! If there's less load, less instances
- EC2 instances are going to push a cloudwatch custom metric which is going to represent the queue length (how many messages are waiting to be processed) divided by the number of instances on your ASG
- This metric, if it goes above a certain threshhold, it means we have either too many messages or not enough instances
- We set a threshold and create a cloudwatch alarm, and the cloudwatch alarm will be automatically be assigned to a scaling policy on your autoscaling group and scale your ASG accordingly
- Maybe you need to create 2 alarms: One for going above the threshold and one for going below it
- No out of the box metric :(, you have to create your own

#### Decoupled scaling application tiers
- This is cool
- Requests > EC2 instances (ASG) > PUT > SQS Queue > Pull > EC2 instances on ASG
- The first group of EC2 instances on the ASG can just put messages on the SQS queue 
- Ths second group can process those messages and can scale up or down based on the load 

### 148. AWS SNS
- What if you want to send a message to many receivers?
- One way of doing this would be "Direct Integration"
    - Buying service sends a message to
    - Email notification
    - Fraud Service
    - Shipping Service
    - SQS Queue
- Another way of doing this is to use the Pub / Sub model
- The buying service publishes the message to a SNS topic
- The SNS topic has many subscribers: email notification service, fraud service, shipping service and a SQS queue
- The "even producer" only sends a message to one SNS topic
- As many "event receivers" (subscriptions) as we want listen to the SNS topic for notifications
- Each subscriber to the topic will get all the messages
- You can have up to 10,000,000 subscription per topic
. 100,000 topics limit
- Subscribers can be:
    - SQS queues
    - HTTP/HTTPS endpoints (with delivery retries!)
    - Lamda functions
    - Emails
    - SMS Messages
    - Mobile notifications

#### SNS integrates with a lot of Amazon Products
- Some services can send data directly to SNS for notifications
    - CloudWatch (for alarms
    - Auto Scaling Group notifications
    - Amazon S3 (on bucket events)
    - CloudFormation (upon stage changes > Failed to build, etc)
    - Etc (I suppose this means "and others")

#### AWS SNS - How to Publish
- Topic Publish (within your AWS server - Using the SDK)
    - Create a topic
    - Create a subscription
    - Publish to the topic
- Direct Publish (for mobile apps SDK) (probably not in the scope of the exam)
    - Create a platform application
    - Create platform endpoint
    - Publish to platform endpoint
    - Works with Google GCM, Apple APNS, Amazon ADM (?)

#### SNS + SQS: Fan out
- Common pattern, popular at exam!
- Push once in SNS, receive in many SQS
- Buying service > SNS topic > 2 or more SQS queues
- Fully decoupled
- No data loss
- Ability to add receivers of data later
- SQS allows for delayed processing
- SQS allows for retries to work
- You can have as many workers as you want on one queue or just one worker per queue

### 149. AWS SNS Hands on
- AWS > SNS > Create Topic
    - Name
    - Encryption - optional
    - Access policy - optional
    - Delivery retry policy (HTTP/S) - optional
    - Delivery status logging - optional
    - Tags (optional)
- Create!
- Create a susbscription
    - Email
    - Hey let's add our email!
    - You get an email, confirm your subscription 
- Publish a message, an email shows up! 
- Neat!

### 150. Kinesis Data Streams Overview
- Increasingly popular exam topic!
- You don't need to know it in-and-out but you need to know what it is, how it works and a few details
- Kinesis is a managed alternative to Apache Kafka
- Great for application logs, metrios, IoT, clickstreams
- Great for "real-time" big data
- Great/compatible with streaming processing frameworks (Spark, NiFi, etc)
- Data is automatically replicated to 3 AZ
- Kinesis Streams: low Latency streaming ingest at scale
- Kinesis Analytics: perform real-tie analytics on streams using SQL
- Kinesis Firehose: load streams into S3, Redshift, ElasticSearch
- Main focus of the exam os on Kinesis streams

#### Kinesis
- Click streams/IoT devices/Metrics & logs > Amazon Kinesis Streams > Amazon Kinesis Analytics > Amazon Kinesis Firehose > S3 Bucket / Amazon Redshift
- Kinesis allows you to onboard data in real time from your big data use cases all the way into the analytics and the final storage

#### Kinesis Streams Overview
- The important part for the exam
- Streams are divided in ordered shards / partitions
- producers > Shard 1 / Shard 2 / Shard 3 > consumers
- Think of shards like lanes in a highway
    - More shards, more throughput
- Data is on a sharrd by 1 day by default, can go up to 7 days
- Kinesis is just a massive highway, it's a big pipe. You just want to process that data as soon as possible and put it somewhere else
- Kinesis also allows you to reprocess / replay data
- Multiple applications can consume the same stream
- Real time processing with scale of throughput
- Once data is inserted in Kinesis, it can't be deleted (immutability). Kinesis is not a database

#### Kinesis Stream Shards
- One stream is made of many different shards
- One shard represents 1MB/s or 1000 messages/s at write per shard
    - A produce can write up to 1MB/s or 1000 messages per second per shard
- On the read side, you have 2MB/s throughput per shard
- Billing is per shard provisioned, you can have as many shards as you want
    - If you overprovision your shards and don't use them to their full capacity you're going to overpay
- Batching available or per message calls(?)
- The number of shards can evolve over time (reshard/merge)
- Records are ordered per shard

### Diagram
- Producers > Shard 1 / Shard 2 / Shard 3 > Consumers
- If you get more throughput:
    - Producers > Shard 1 / Shard 2 / Shard 3 / Shard 4 > Consumers
    - (adding a shard: resharding)
- If you get less throughput: 
    - Producers > Shard 1 / Shard 2 > Consumers
    - (removing a shard: merging)

#### AWS Kinesis API - Put Records
- Producer side, there's something called Put records. A way to send data to Kinesis
- You send:
    - Data
    - Message Key
- Message key is whatever you want as a string
- This key will get hashed to determine the shard id
- The key is a way for your to route the data to a specific shard
    - The same key always goes to the same partition (helps with ordering for a specific key)
- Messages when they're sent to a shard they get a sequence number (always increasing)
- Remember to choose a partition key that is highly distributed! (helps prevent "hot partition")
- If your key isn't distributed, all your data will go to the same shard and this shard will be overwhelmed
- If you have an app:
    - user_id is a great key. You have many many users
    - country_id is a poor key. You chould have all your users (or most of them) in a single country
- You can use batching with PutRecords to reduce costs and increase throughput
- If you get an exception named **ProvisionedThroughputExceeded* it meant you went over the message limit
    - For this you can use retries and exponential backoff

#### AWS Kinesis API - Exceptions
- ProvisionedThroughputExceeded Exceptions
    - Happens when sending more data whan what was provisioned (exceeding MB/s or TPS for any shard)
    - Make sure you don't have a hot shard (such as your partition key being bad with too much data going to a single partition)
Solution:
    - Retries with backoff
    - Increase shards (scaling)
    - Ensure your partition key is a good one
        
#### AWS Kinesis API - Consumers
- Can use a normal consumer (CLI, SKD, etc)
- Can use the Kinesis client library (in Java, Node, Python, Ruby, .Net)
    - KCL uses DynamoDB to checkpoint offsets (?) 
    - KCL uses DynamoDB to track other workers and share work amongst shards
- "KCL enables to consume from Kinesis efficiently"    

#### Kinesis Security
- Control access / authorization using IAM policies
- Encryption in flight using HTTPS endpoints
- Encryption at rest using KMS
- Possibility to encrypt / decrypt data client side (harder to implement, you need to write your own code)
- VPC Endpoints available for Kinesis to access within VPC

### 151. Kinesis Data Streams Hands on
- Create data stream
- There's a shard calculator
- I'll go for just 1 shard!
- Creatr data stream
- `aws kinesis help
- `aws kinesis help list-streams`
- `aws kinesis help describe-stream`
- `aws kinesis describe-stream --stream-name myfirststream`
- `aws kinesis help put-record`
- https://www.base64encode.org/
- `aws kinesis put-record --stream-name myfirststream --data "aGVsbG8gd29ybGQ=" --partition-key user-123`
- `aws kinesis get-shard-iterator help`
- `aws kinesis get-shard-iterator --stream-name myfirststream --shard-id shardId-000000000000 --shard-iterator-type TRIM_HORIZON`
- `aws kinesis get-records --shard-iterator "AAAAAAAAAAEpHLP88sp9CFBzkO2q//KBWtqnX4k8qKCmY/WAcMXHBzRMqgI4vmfM8n54Spb6G9+uZIfGmF+edg6Y12C4ZivnFWWgxgSz1wGd3DczitUpdDz10QKC6+hBn/18fBrjtl17p+LOGkeyTp1Xuy1Jkc8Km+oiHUteNi7IoojoTT+l3Kf+9/ln91nPM7bS5YEAzXGFUkwplWzcLG8ajx748Rqy"`

### 152. Kinesis Firehose & Kinesis Data Analytics
- Fully Managed Service, no administration, automatic scaling, serverless
- Loads data into Redshift / Amazon S3 / ElasticSearch / Splunk (remember those!)
- Near Real Time (kinesis data stream was real time)
    - 60 seconds latency minimum for non full batches
    - Or minimum 32 MB of data at a time to load into these stores above
- Supports many data formats, conversions, transformations, compression (csv, json and so on )
- Pay for the amount of data going through Firehose (you don't pay anything for provisioning)

#### Diagram
Sources:
    - SDK Kinesis Producer Library - (KPL)
    - Kinesis Agent
    - Kinesis Data Streams
    - CloudWatch logs & events
- Go into
    - Amazon Kinesis Data Firehoses
    - You can do transformations on the data using Lamda functions
- That will save the data to
    - S3
    - Redshift
    - ElasticSearch
    - Splunk
- Great service to transport data and do data ingestion in near real time

#### Kinesis Data Stream vs Firehose
- Streams
    - For when you're going to write custom code (producer / consumer)
    - Real time (about 200ms)
    - You must manage scaling yourself (shard splitting / merging)
    - Can store data from 1 to 7 days, replay capability, multi consumers
- Firehose
    - Fully managed, sends to S3, Splunk, RedShift and ElasticSearch
    - Serverless data transformations with Lambda
    - Near real time (lowest buffer time is 1 minute)
    - Automated Scaling
    - No data storage. You can't replay from firehose

#### Kinesis Analytics, conceptually
- Kinesis analytics can take data from Kinesis Datastreams and Firehose and perform queries on it
- And the output of these queries can be analyzed by your own analytics too (?)
- Perform real time analytics on kinesis streams using SQL
- Kinesis Data Analytics
    - Auto Scaling
    - Managed: no servers to provision
    - Continous: real time
- Pay for actual consumption rate
- Can create streams out of the real-time queries

### 153. [SAA-C02] Data Odering for Kinesis vs SQS FIFO
#### Ordering data into Kinesis
- Imagine you have 100 trucks (truck_1, truck_2, ... truck_100) on the road sending their GPS positions regularly into AWS
- You want to consume the data in order for each truck, so that you can track their movement accurately
- How should you send that data into Kinesis?
    - Answer: send using a "Partition Key" value of the "truck_id"
    - The same key will always go to the same shard

#### Ordering data into SQS
- For SQS standard, there is no ordering
- For SQS FIFO, if you don’t use a Group ID, messages are consumed in the order they are sent, with only one consumer
- If you had trucks like in the previous example, all the trucks would be sending data into a FIFO queue, but there can only be one consumer
- Sometimes uou want to scale the number of consumers, but you want messages to be "grouped" when they are related to each other
• For this you use a Group ID (similar to a Partition Key in Kinesis)
- Using a group ID your FIFO queue for instance can have two groups processed in a FIFO manner
- So you can have group1 and group2 and you can have consumers independently reading group1 and group2
- The more group ids you have the more consumers you can have

#### Kinesis vs SQS ordering
- Let’s assume 100 trucks, 5 kinesis shards, 1 SQS FIFO
- Kinesis Data Streams:
    - On average you’ll have 20 trucks per shard
    - Trucks will have their data ordered within each shard
    - The maximum amount of consumers in parallel we can have is 5 (we have 5 shards and we need 1 consumer per shard)
    - Can receive up to 5 MB/s of data
- SQS FIFO
    - You only have one SQS FIFO queue
    - You will have 100 Group IDs
    - You can have up to 100 Consumers (due to the 100 Group ID)
    - You can have up to 300 messages per second (or 3000 if using batching)

### 154. SQS vs SNS vs Kinesis
- Common question at the exam!

#### SQS
- Consumers "pull data"
- Data is deleted after being consumed
- Can have as many workers (consumers) as we want
- No need to provision throughput
- No ordering guarantee (except FIFO queues, but then you do get limited throughput)
- Individual message delay capability

#### SNS
- More of a pub/sub
- Push data to many subscribers
- Up to 10,000,000 subscribers
- Data is not persisted (lost if not delivered)
- Up to 100,000 topics
- No need to provision throughput
- Integrates with SQS for fan-out architecture architecture pattern

#### Kinesis
- Consumers "pull data"
- As many consumers as we want, but can only have 1 consumer per shard
- There is the possibility to replay data
- Meant for real-time big data, analytics and ETL
- Ordering at the shard level
- Data expires after X days
- Must provision throughput

### 155. Amazon MQ
- SQS, SNS are “cloud-native” services, and they’re using proprietary protocols from AWS (they're just available on AWS)
- Traditional applications running from on-premise may use open protocols such as: MQTT, AMQP, STOMP, Openwire, WSS and others
- When migrating to the cloud, instead of re-engineering the application to use SQS and SNS, we can use Amazon MQ!
- Amazon MQ = managed Apache ActiveMQ
- Amazon MQ doesn’t "scale" as much as SQS / SNS, You have to provision it.
- Amazon MQ runs on a dedicated machine, can run in HA with failover and Multi AZ
- Amazon MQ has both queue feature (~SQS) and topic features (~SNS)
- If the exam says "There's an app using MQTT or AMQP and we want to move it to the cloud, what do we use?" AmazonMQ!

#### Hands on
- Console > AmazonMQ
- "Managed message broker for ApacheMQ"
- Select deployment and storage type
    - You can use single instance, great for development and testing
    - You can have a standy by broker option, similar to Multi-AZ RDS
- Configure settings
    - Name
    - Instance size
    - Web console credentials
    - Deployment time is about 20 minutes
- On advanced options
    - Configuration
    - Logs
    - VPC
    - Encryption
    - Maintenance Windows
    - Tags

### Quiz:
- Pending!