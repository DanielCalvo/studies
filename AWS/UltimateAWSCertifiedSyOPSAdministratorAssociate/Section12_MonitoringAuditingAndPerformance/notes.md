- Chapter status: Took notes attentively, got the quiz very right!
- Personal note: Interesting to set up on your own to track expenditures and automated things

### 141. Section intro
- CloudWatch: Important in the exam 

### 142. CloudWatch Metrics
- Cloudwatch provides metrics for every service in AWS
- Metric is a variable to monitor (cpu utilization, network in, etc)
- Metrics will belong to namespaces
- A dimension is an attribute of a metric (instance id, environment, etc)
- You can have up to 10 dimensions per metric
- Metrics will have timestamps
- You can create dashboards with the metrics you want

#### AWS CloudWatch EC2 Detailed monitoring
- EC2 instances have metrics every 5 minutes
- Detailed monitoring enables per minute metric for an additional cost
- Using detailed monitoring you could get more prompt scaling in your ASG!
- The AWS free tier allows you to have 10 detailed monitoring metrics
- Memory usage is not pushed by default on EC2, it must be pushed as a custom metric using monitoring scripts

#### AWS CloudWatch custom metrics
- You can defined and send your own metrics to cloudwatch
- You can send dimentions (attributes) to segment metrics (up to 10)
- Metric resolution
    - Standard 1 minute
    - There can be high resolution metrics, in which you send metrics to cloudwatch every 1 second changing the StorageResolution API parameter. These incur a higher cost
- To send custom metrics to cloudwatch, use an API call named PutMetricData
- Use exponential back off in case of throttle errors

#### Hands on
- Created an EC2 instance
- Went on CloudWatch > Metrics > Searched for "CPU" and graphed metrics

### 143. CloudWatch Dashboards 
One or two questions about this
- Great way to get access to key metrics and see how your application is running
- Dashboards are global
- Dashboards can include graphs from different regions (popular at the exam)
- You can change the time zone and time range of the dashboards
- You can set up automatic refresh (10s, 1m, 2m , 5m, 15m)

#### Hands on!
- Created dashboard, added a line widget for CPU usage, saved dashboard
- Remember, dashboards are global!

#### Pricing
- 3 dashboards (up to 50 metrics for free)
- 3 dollars per dashboard per month afterwards

### 144. CloudWatch Logs
- Your application can send logs to cloudwatch using the SDK
- Cloud watch can collecct logs from
    - Elastic beanstalk
    - ECS
    - AWS Lambda
    - VPC flow logs
    - API Gateway
    - CloudTrail based on filter
    - Cloudwatch log agents from EC2 machines
    - Route 53
- Cloud watch itself can send logs to whatever you want
- Batch exporter to S3
- Stream to ElasticSearch cluster for further analytics

- Logs storage architecture
    - Log groups: Arbitrary name, usually representing an application
    - Log streams: You have many of those in a log group, instances within an application / log files / containers
- You can define a log expiration policy (never expire, 30 days, etc) 
- You can use the AWS CLI to tail cloud watch logs (neat!)
- To send logs to cloudwatch, make sure you're using the correct IAM permissions!
- Security: You can encrypt logs using KMS at the group level

#### CLoudwatch logs metric filter & insights
- Cloud watch logs can use filter expressions (for search)
- You can set up a metric filter to trigger cloudwatch alarms (like looking for a specific IP)
- Cloud watch logs insight can be used to query logs and add queries to cloudwatch dashboards

#### Hands on
- You can create a log filter on a log stream and create an alarm!

### 145. CloudWatch Alarms
- Used to trigger notifications for any metrics
- Alarms can trigger autoscaling, EC2 actions, SNS notifications
- Varios options (sampling, %, max, min, etc)
- Alarm states:
    - OK
    - INSUFFICIENT_DATA (no data)
    - ALARM (when alarm is on a triggered state)
- Period:
    - Lenght of time in seconds to evaluate the mtric
    - For high resolution metrics: 10 sec or 30 swec
 
#### Cloud watch alarm targets
- Stop, terminate, reboot or recover an EC2 instance
- Trigger autoscaling action
- Send notifications to SNS (from which you can do pretty much anything)

#### Cloud watch alarms, good to know for the exam:
- Alarms can be created based on cloudwatch logs metrics filters
- Cloudwatch doesn't test or validate the actions that it is assigned (if you assign a wrong action cloudwatch won't tell you)
- To test alarms and notifications, set the alarm state to Alarm using the CLI
- `aws cloudwatch set-alarm-state --alarm-name "myalarm" --state-value ALARM --state-reason "testing purposes"`

#### Alarms hands on
- Created an alarm based on the EC2 CPU load and set the action to instance reboot. Cool!
- Author then went on the CLI and set the alarm to ALARM state manually (I was lazy)

### 146. CloudWatch Events
You don't need to know a lot of detail, just know them at a high level

- You choose a source, apply a rule to it and give it a target (Source + Rule => Target)
- Schedule:
    - Cronjobs
    - Event patterns: React to a service doing something, such as an EC2 instance being created
- Can trigger lambda functions, SQS, SNS, Kinesis Messages
- ClougWatch Events creates a small JSON document to give information about the change

#### Hands on
- Create an Alarm for every time there is an EC2 instance in a pending state (created instances are always pending at first)
- This will notify us whenever someone creates an instance! 

### 147. CloudTrail
Important for the exam!
- Provides governance, compliance and audit for your AWS account. It will track every API call to your account
- CloudTrail is enabled by default
- Gets a history of events / API calls made within your AWS account by:
    - Console
    - SDK
    - CLI
    - AWS Services
- You can see who did that and when
- You can put logs from Cloudtrail into Cloudwatch logs
- If a resource is deleted in AWS, look into Cloudtrail first! (popular exam question)
- Cloudtrail shows the past 90 days of activity
- The default UI will only show create, modify of delete events

But you can create cloudtrail trails
- Get a detailed list of all the events you choose
- You have the ability to store these events in S3 for further analysis
- Can be region specific or global
- Cloudtrail logs have SSE-S3 encryption when placed into S3
- You can control S3 access using IAM, bucket policies and etc

#### Hands on
- Hey this is pretty cool, you can see everything in cloud trail!
- Created a new cloud trail to log all EC2 events. Stores them in a S3 bucket. Neat!
- In case anything goes wrong or in case you need to comply with some policy, this is very useful

### 148. Config Overview
Popular exam topic!
- Helps with auditing and compliance of your AWS resources
- Helps record configuration changes over time
- Helps record compliance over time
- It's possible to store your AWS config data into S3 (can be queried by Athena)
- Questions that can be answered by AWS config:
    - Is there unrestricted SSH access to my security groups?
    - Do my buckets have public access?
    - How has my ALB configuration changed over time? Who changed it, and how did it get changed?
- You can receive alerts (SNS notifications) for any changes
- AWS Config is a per-region service
- But you can use it to aggregate across regions and accounts
 
#### Config rules
- Can use AWS managed config rules (over 75)
- You can also make custom rules (must be defined in AWS Lambda). Examples:
    - Evaluate if each EBS disk is of type gp2
    - Evaluate if each EC2 instance is of type t2.micro
- Rules can be evaluated based on two triggers:
    - For each config change
    - And/or at regular time intervals
- Pricing:
    - No free tier `:(`
    - 2 USD per active rule per region per month (less after 10 rules)

#### AWS config resource (as seen on the UI)
- You can see the compliance of a resource over time
- View configuration of a resource over time 

### 149. Config Hands On
- You create a config
- You add rules to it
- You can check if you're only running a certain type of AMI, which is tagged as "compliant". That's so cool!
- You can also check if certain security groups have open ssh, or a lot of other things. So cool!

### 150. CloudWatch vs CloudTrail vs Config
- CloudWatch
    - Performance monitoring (metrics, cpu, network, etc) & dashboards
    - Events & Alerts
    - Log aggregation & analysis
- CloudTrail
    - Records API calls made within your account by everyone
    - You can define some trails for specific resources (such as EC2 only)
    - It's a global service
- Config
    - Record configuration changes
    - Evaluate resources against compliance rules
    - Get a timeline of changes and compliance

#### Elastic load balancer example
- Cloudwatch
    - Monitoring incoming connections as a metric
    - Visualize error codes as % over time
    - Make a dashboard to get an idea of your load balancer performance
- Config
    - Track security groups for the load balancer
    - Track configuration changes for the load balancer
    - Ensure an SSL certificate is always assigned to the load balancer (compliance)
- CloudTrail
    - Track who made changes to the load balancer with API calls

### 151. Section clean up

### Quiz