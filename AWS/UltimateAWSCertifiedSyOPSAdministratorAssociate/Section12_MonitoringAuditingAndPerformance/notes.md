### 141. Section intro
- CloudWatch: Important in the exam 

### 142. CloudWatch Metrics
- Cloudwatch provides metrics for every service in AWS
- Metric is a variable to monitor (cpu utilization, network in, etc)
- Metrics will belong to namespaces
- A dimension is an atribute of a metric (instance id, environment, etc)
- You can have up to 10 dimensions per metric
- Metrics will have timestamps
- You can create dashboards with the metrics you want

#### AWS CloudWatch EC2 Detailed monitoring
- EC2 instances have metrics every 5 minutes
- Detailed monitoring enables per minute metric for an additional cost
- Using detailed monitoring you could get more prompt scaling in your ASG!
- The AWS free tier allows you to have 10 detailed monitoring metrics
- Memory usage is not pushed by defaul on EC2, it must be pushed as a custom metric using monitoring scripts

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
- Cloud watch can collecct log from
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
- Cloudwatch doesn't test or validate the actions that it is assigned (if you assign a wrong action it won't work)
- To test alarms and notifications, set the alarm state to Alarm using the CLI
- `aws cloudwatch set-alarm-state --alarm-name "myalarm" --state-value ALARM --state-reason "testing purposes"`

#### Alarms hands on
- Continue here! Around minute 2

### 146. CloudWatch Events

### 147. CloudTrail

### 148. Config Overview

### 149. Config Hands On

### 150. CloudWatch vs CloudTrail vs Config

### 151. Section clean up

### Quiz