
### 187. AWS Monitoring - Section Introduction
- Monitoring is very important!
- Stephane says: "I never deploy anything without monitoring". Wise Stephane

### 188. AWS CloudWatch Metrics
- CloudWatch provides metrics for every services in AWS
- A metric is a variable to monitor (CPUUtilization, NetworkIn...)
- Metrics belong to namespaces (they're grouped!)
- A dimension is an attribute of a metric (instance id, environment, etc...).
- You can have up to 10 dimensions per metric
- Metrics have timestamps (it's a time series!)
- Can create CloudWatch dashboards of metrics

#### AWS CloudWatch EC2 Detailed monitoring
- EC2 instance metrics have metrics “every 5 minutes”
- With detailed monitoring (for a cost), you get data “every 1 minute”
- (you can) Use detailed monitoring if you want to more prompt scale your ASG!
- The AWS Free Tier allows us to have 10 detailed monitoring metrics
- Note: EC2 Memory usage is by default not pushed (must be pushed from inside the instance as a custom metric)

#### AWS CloudWatch Custom Metrics
- You can defined and send your own custom metrics to CloudWatch (using the CLI for instance)
- (you have the) Ability to use dimensions (attributes) to segment metrics
- You can send, for instance
    - Instance.id
    - Environment.name
- Metric resolution (StorageResolution API parameter – two possible value):
    - Standard: 1 minute (60 seconds)
    - High Resolution: 1 second - Higher cost (if you need very high detail)
- To just send a metric to CloudWatch use API call PutMetricData
- Use exponential back off in case of throttle errors

#### Hands on!
- Console > CloudWatch!
- Metrics on the left hand side menu
- Oh, there are metric groups
- Launch an EC2 instance
- Oh you can see EC2 per instance metrics!
- Damn there are logs of metrics, neat

### 189. AWS CloudWatch Dashboards
- Popular at the exam!
- Great way to setup dashboards for quick access to keys metrics
- Dashboards are global! You can access them in any region!
- Dashboards can include graphs from different regions
- You can change the time zone & time range of the dashboards
- You can setup automatic refresh (10s, 1m, 2m, 5m, 15m)
- Pricing:
    - 3 dashboards (up to 50 metrics) for free
    - $3/dashboard/month afterwards (they can get pricy!)

#### Hands on!
- On cloudwatch, go on dashboards, and then create dashboard!
- Add a line graph for cpu utilization
- Ooh we have a widget! Don't forget to save the dashboard!
- Nice time range possibilities on the UI
- There's a refresh button! With auto refresh! And you can choose timezones! 
- Dashboards are global! When you create a dashboard in Ireland you can see it in North Virginia
    - So you can have a dashboard with metrics from instances in several regions
    - At the exam they'll ask you how to build a global dashboard
    - You can go to another region and add metrics from that region to your dashboard!


### 190. AWS CloudWatch Logs
- Applications can send logs to CloudWatch using the SDK
- But CloudWatch can also collect logs from
    - Elastic Beanstalk: Collection of logs from application
    - ECS: collection from containers logs
    - AWS Lambda: collection from function logs
    - VPC Flow Logs: VPC specific logs
    - API Gateway logs
    - CloudTrail based on filter (if you set it up)
    - CloudWatch log agents: for example on EC2 machines
    - Route53: Log DNS queries
- CloudWatch Logs can go to:
    - Batch exporter to S3 for archival
    - Stream to ElasticSearch cluster for further analytics

#### CloudWatch logs, part two
- Logs storage architecture: You need to store logs in 2 things (?)
    - Log groups: arbitrary name, usually representing an application
    - Log stream: instances within application / log files / containers
- You can define log expiration policies (never expire, 30 days, etc..)
- You can use the AWS CLI we can tail CloudWatch logs
- To send logs to CloudWatch, make sure IAM permissions are correct! (it's a common mistake)
- Security: encryption of logs using KMS at the Group Level. Logs can be encrypted

#### CloudWatch Logs Metric Filter & Insight
- New feature that might not appear on the exam, but good to know!
- CloudWatch Logs can use filter expressions
    - For example, find a specific IP inside of a log
    - Metric filters can be used to trigger alarms
- CloudWatch Logs Insights (new – Nov 2018) can be used to query logs and add queries to CloudWatch Dashboards

#### Hands on!
- EC2 > Launch instance
- On the instance details, click on create IAM role
- Ec2 role, next, add the policy `CloudWatchFullAccess` (this policy seems a bit excessive with the permissions though)
- Named it `DemoCloudWatchEC2Role`. Refresh on the IAM thingy on the EC2 creation page and add that role!
- And launch it with the defaults!
- SSH into the instance
- We can now configure the log agent as described here: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/QuickStartEC2Instance.html
    - `yum update`
    - `yum install awslogs`
- You then need to change `/etc/awslogs/awscli.conf` to send logs to the correct region
- `systemctl start awslogsd`
- Check `cat /var/log/awslogs.log` to see if there's anything wrong with the set up. Aaaay I got permission denied, I forgot to attach a policy to my IAM role
- Yay, now there's a log group for /var/log/messages and you can filter by instance, and my instance is there! 
- You can set an expiration policy on the logs! You also an export the data to S3, Stream to Lambda or Elasticsearch. Neat


### 191. AWS CloudWatch Alarms
- Alarms are used to trigger notifications for any metric
- Alarms can be attached to Auto Scaling, EC2 Actions, SNS notifications
- There are various options (sampling, %, max, min, etc...)
- Alarm States:
    - OK
    - INSUFFICIENT_DATA
    - ALARM
- Period:
    - You have to specify a length of time in seconds to evaluate the metric
    - High resolution custom metrics: can only choose 10 sec or 30 sec as the period for the evaluation of your alarm

#### Hands on
- Cloudwatch > alarms
- Stephane then goes on to show alarms linked to an ASG
- Let's create an alarm
    - Select a metric
    - Select the statistic (live avg over 5 minutes)
    - Threshold type, I'll go for static
    - And I'll put higher than 0.01 for CPU load which will trigger!
- Next, set up the notification
    - AWS as asking me to specify a SNS topic for this, I think I'll leave it blank
    - You can also add an EC2 action or an autoscaling action, neat!

### 192. AWS CloudWatch Events
- Schedule: Cron jobs (you can define a schedule)
- Event Pattern: Event rules to react to a service doing something
    - Ex: CodePipeline state changes!
- Triggers to Lambda functions, SQS/SNS/Kinesis Messages
- CloudWatch Event creates a small JSON document to give information about the change

#### Hands on
- Can select either an event pattern or a schedule to trigger the event
    - So you could trigger and event every 5 minutes for instance
    - You can then send this event to a target (SQS, SNS, Lambda)
- Example: You generate an event every time an EC2 instance is shut down, or a CodePipeline fails

### 193. AWS CloudTrail
- Provides governance, compliance and audit for your AWS Account
- CloudTrail is enabled by default!
- "Any time you do anything in your AWS account it should show up in cloud trail"
- Get an history of events / API calls made within your AWS Account by:
    - Console
    - SDK
    - CLI
- AWS Services
- Can put logs from CloudTrail into CloudWatch Logs
- If a resource is deleted in AWS, look into CloudTrail first! (possible exam question, most common one)

#### Hands on!
- Console > CloudTrail
- Look at all those events! That's cool
- I can look at when I deleted an instance or an S3 bucket
- You can filter by events and you can also create a trail, but creating a trail seems to be out of scope

### 194. [SAA-C02] AWS Config - Overview
- Helps with auditing and recording compliance of your AWS resources
- Records configurations and changes over time
- Can store all of this configuration data into S3 (you could analyze it with Athena!)
- Questions that can be solved by AWS Config:
    - Is there unrestricted SSH access to my security groups?
    - Do my buckets have any public access?
    - How has my ALB configuration changed over time?
- You can receive alerts (SNS notifications) for any changes
- AWS Config is a per-region service
- Can be aggregated across regions and accounts

#### Example
- You can view the compliance of a resource over time
- View the configuration of a resource over time
- It integrates with cloud trail, so you can see who did the changes

#### AWS Config Rules
- Can use AWS managed config rules (there are over 75)
- Or you can make custom config rules (must define that rule using AWS Lambda)
- Example rules: 
    - Evaluate if each EBS disk is of type gp2
    - Evaluate if each EC2 instance is t2.micro
- Rules can be evaluated / triggered:
    - For each config change
    - And / or: at regular time intervals
    - Can trigger CloudWatch Events if the rule is non-compliant (and chain it with Lambda)
- Rules can have auto remediations:
    - If a resource is not compliant, you can trigger an auto remediation
    - Ex: stop instances with non-approved tags
- AWS Config Rules does not prevent actions from happening (no deny)
- But it does help evaluate the compliance of your resources over time
- Pricing: no free tier, $2 per active rule per region per month
- No free tier :(

### 195. [SAA-C02] AWS Config - Hands on
- Choose the resource you want to record
- You can log that change to a S3 bucket
- You can also stream changes to an SNS topic
- Then you can choose the config rules to check
    - There's a rule to only run approved AMIs, neat!
    - Also another one for incoming ssh disabled
- You can set a rule to be triggered every time there's a configuration change or every x time
- You can also choose the scope of the rule (ex: the resources)
- Rule was applied to all SG, now we can see non compliant SGs! (the ones with ssh allowed)
- And you can have a dashboard with which resources are compliant and which ones are non-compliant. Neat!

### 196. [SAA-C02] CloudTrail vs CloudWatch vs Config
- Popular question at the exam!
- CloudWatch
    - Performance monitoring (metrics, CPU, network, etc...) & dashboards
    - Events & Alerting
    - Log Aggregation & Analysis
- CloudTrail
    - Record API calls made within your Account by everyone
    - Can define trails for specific resources
    - Global Service
- Config
    - Record configuration changes
    - Evaluate resources against compliance rules
    - Get timeline of changes and compliance

#### For an Elastic Load Balancer, here's an example:
 - CloudWatch
    - Monitoring Incoming connections as metrics
    - Visualize error codes as a % over time
    - Make a dashboard to get an idea of your load balancer performance
    - You can even make it a global dashboard if you have multiple load balancers on your application
- Config
    - Tracking the security group rules for the LB, making sure everything is compliant to security standards
    - Track configuration changes for the LB
    - Ensure that an SSL certificate is always assigned to the LB (compliance)
- CloudTrail
    - Who made any changes to the load balancer woth API calls (if someone changes SG rules or the SSL certificate or w/e)
    - CloudTrail will let you know who made those changes!
- Stephan recommends thinking about the context of a load balancer for questions regarding these services (what service does what!)

#### Quiz
- Pending!