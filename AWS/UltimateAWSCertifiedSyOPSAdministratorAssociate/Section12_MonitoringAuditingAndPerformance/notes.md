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

### 144. CloudWatch Logs

### 145. CloudWatch Alarms

### 146. CloudWatch Events

### 147. CloudTrail

### 148. Config Overview

### 149. Config Hands On

### 150. CloudWatch vs CloudTrail vs Config

### 151. Section clean up

### Quiz