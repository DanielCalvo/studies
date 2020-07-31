
### Practice Test 1
1. Review scaling policies again
    - https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-target-tracking.html
    - https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html
2. Review which volumes can be boot volumes
    - gp2 and io1
4. What is AWS Global Accelerator? What is S3 Transfer acceleration?
    - https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html
6. What are the invalid lifecycle transitions for S3 buckets? I was really lost on this one!
    - https://docs.aws.amazon.com/AmazonS3/latest/dev/lifecycle-transition-general-considerations.html
9. Is Multi-AZ RDS Synchronous? (not read replicas)
    - Yes. Read replicas are async
14. You need to set desired instances for a number of instances to exist, not min/max
15. What is the update strategy for multi-az RDS dbs? What gets upgraded first? Is it the primary or secondary instance?
    - https://aws.amazon.com/premiumsupport/knowledge-center/rds-required-maintenance/
    - db engine level updates have downtimes
18. What is cheaper? S3 IA or S3 Intelligent tiering?
    - https://aws.amazon.com/s3/storage-classes/
19. I forgot the placement groups again. Read on spread and other ones from AWS
    - https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html
22. Data warehousing = Redshift. Always!
    - https://aws.amazon.com/blogs/big-data/amazon-redshift-spectrum-extends-data-warehousing-out-to-exabytes-no-loading-required/
24. Damn I was stupid. Snowballs!
25. https://aws.amazon.com/elasticache/redis-vs-memcached/
27. Damn I was stupid. ASG = EC2 health checks. ALB = ALB based health checks. 
29. Hmm have another look at this, it's a bit vague
30. Can you recover a deleted key in KMS? Have a review!
    - Yes! CMK stay as pending delete anywhere between 7 to 30 days
    - https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html
31. What's up with SQS FIFO Batch mode? Is that even a thing?
    - https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html
    - SQS FIFO has a limit of 300 messages per second without batching. Batching by 3 makes this 900 (x3) and so on
32. Is ambiguous though

33. Confusing, review! <- Do last!

39. Review 39, I think lambda and kinesis data stream don't work together? <- Dunno
    - But Kinesis data stream is not for data ingestion. Use Kinesis data firehose or data stream
    - Also kinesis requires provisioning
40. Maybe it would be a good idea to do the KMS hands on
43. This one was tough to understand
    - https://aws.amazon.com/fsx/lustre/
44. What is WebSocket?
    - https://aws.amazon.com/api-gateway/
45. What is S3TA again?
    - https://aws.amazon.com/s3/transfer-acceleration/
51. Remember that with EBS you pay for the whole volume, not only for what you just use! With S3 and EFS, you pay for what you just use
59. Concurrency limits for Lambda are 1000 executions at the same time
58. https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html

- 78%, oof

#### Other notes
- AWS Glue: Fully managed, extract, transform, load (ETL). To prepare date for analytics. For batch ETL data processing.
- AWS Global Accelerator: It provides static IP addresses that acts as a fixed entry point to your application endpoints in a single or multiple AWS Regions, such as your Application Load Balancers, Network Load Balancers or Amazon EC2 instances.

### Practice Test 2
4. Review the features of WAF
    - Damn I misread the question, I'm such a dummy. You were supposed to block the IP address!
6. What's a spot fleet request again?
    - https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html
    - Messed up because I forgot an EC2 instance needs an instance role to perform API calls 
11. Can you send S3 event notifications to SQS queues and SQS FIFO queues?
    - No. S3 cannot send events to an SQS FIFO queue (not allowed/supported)
13. Have another look at that redis vs memcached feature chart
    - https://aws.amazon.com/elasticache/redis-vs-memcached/
    - Redis just has more features
16. Messed up! SNS would've sent a message to all the instances. SQS was the correct answer
20. Have a quick read on cloudwatch alarm actions real quick
    - Yeah they're a thing
21. SQS Short poling vs long poling, what are the differences?
    - Doesn't matter, your mistake was: It is the consumer application's responsibility to process the message from the queue and delete them once the processing is done. Otherwise, the message will be processed repeatedly by consumer applications. Default retention is 4 days.
24. I really think this question is right but forgot what the wrong options mean!
    - VPC Peering: Network connection between 2 vpcs
    - VPC Endpoint: Connect to AWS services on a private subnet
    - Internet Gateway: Network object to route traffic to the internet on a VPC
    - Transit Gateway: Connects VPC and on-premises networks through a central hub.
27. I forgot all the details for account migration (like from org a to org b)
    1. Remove member account from old org
    2. Send invite to member account to new org
    3. Accept invite to new org from member account
29. This question was cool, revise the answer, maybe you got it right!
30. S3 is eventually consistent, so if you overwrite an object, you MIGHT (not WILL) get the old one until replication is complete
32. Completely misread the question. You dummy! Use hibernate to store the memory stage of an instance
33. Tough one! But I got it right! :D
    - Kinesis Data Streams requires manual administration of shards, so requires provisioning
35. OLTP means what service again? I suspect aurora
    - Yep you were right
36. Man I have no idea about his kinesis wizardry :joy:. I'm gonna YOLO this one
    - I YOLOed correctly! ahaha
    - Kinesis Agent cannot write to a Kinesis Firehose for which the delivery stream source is already set as Kinesis Data
37. Spot fleet can be cheaper than spot instances right?
    - Maybe, but this was a use-case type thing apparently
40. This one seems a bit ambiguous!
    - But I got it right!
41. What the difference between vpc peering and vpc private link again?
    - VPC Peering: Network connection between 2 vpcs
    - VPC PrivateLink: AWS PrivateLink provides private connectivity between VPCs, AWS services, and on-premises applications
45. ASG Launch template is newer and has more features than a launch configuration, right?
    - Correct. Launch templates also allow multiple versions of a template
50. Oh this one was tricky. Remember to read the question well!
52. Interesting, hard to understand well. Have a read on the explanation
    - https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html
    - Partition: Maximum of 7 partitions per AZ, only limited by the limits on your account
    - Spread: Maximum of 7 instances per AZ
    - Cluster: Only limited by the limits on your account
55. Should've used permission boundaries to avoid developers attaching the AdmnistratorAccess policy to themselves
    - Permission Boundaries can only be attached to roles or users, not IAM groups
    - Service Control Policies (SCP) can only by used in the context of AWS Organizations, which were not mentioned in the question
    - https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html <- Yeah have another read at this one, it's cool
63. What the naming convention for S3 static website DNS entries again?
    - http://bucket-name.s3-website.Region.amazonaws.com
    - http://bucket-name.s3-website-Region.amazonaws.com
    - I lucked out so much on this one it was hilarious

### Practice Test 3
1. Which of the following S3 storage classes supports encryption by default for both data at rest as well as in-transit?
    - What the fuck was this question, jesus lord
    - S3 Glacier is the correct answer
2.
    - Message timers can be used to delay the delivery of _certain_ (not all) messages
    - Delay queues postpone the delivery of _all_ new messages.
6. ALB geo match is a thing? Can you block IPs with WAF?
    - ALB cannot block or allow traffic based on geographical matches
    - WAF Can do geo matching and IP based rules for allow/deny of traffic
9. Remember that for data ingestion you want to use kinesis data stream or firehose, not analytics
12. Forgot some of these services, give them a review
    - AWS Config: - Helps with auditing and recording compliance of your AWS resources. Records configurations and changes over time
13. No idea with this Kinesis stuff, maybe give it a review again. Was kinesis fan-out even covered?
    - You can't have application consuming out of kinesis data firehose
    - Enhanced fan out allows stream consumers to receive their own 2mb/s pipe of read throughput per shard
14. Turns out VPC Sharing is a thing and it's part of RAM. Tricky question!
15. AWS Route 53 does not charge for ALIAS queries
    - You cannot create CNAME record for an apex zone
19. Have of review of cloudformation real quick
    - Yeah really review cloudformation nomenclature
22. Is ELB idle timeout a thing?
    - Yeah
24. No idea man
25. No idea, review
    - Kinesis Firehose does not support EMR as a target for delivering the streaming data
    - Supported targets: Amazon S3, Amazon Redshift, Amazon Elasticsearch Service, and Splunk
27. Can you recover EC2 instances?
    - Yes. They are identical to the original instance, retaining their public IPv4 if applicable
    - https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-recover.html
30. Difference between AWS Step Functions and Simple Workflow Service?
    - Use step functions for new applications
    - https://aws.amazon.com/swf/faqs/
32. Oh wow I got this right.
    - If cross az load balancing is enabled, traffic is distributed evenly across all instances regardless of az
    - it if is disabled, traffic is split equally between azs, regardless how many instances you have
33. Well whadyaknow, NAT instances support port forwarding
    - https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-comparison.html
35. You can't convert an existing standard queue into a FIFO queue.
39. Interesting question
    - Use a VPC endpoint to access Amazon SQS (if on a private subnet)
43. I was just a dummy on this one
44. If you set the Launch Configuration Tenancy to default and the VPC Tenancy is set to dedicated, then the instances have dedicated tenancy. If you set the Launch Configuration Tenancy to dedicated and the VPC Tenancy is set to default, then again the instances have dedicated tenancy.
    - https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-vpc-tenancy
45. 
    - If a spot request is persistent, then it is opened again after your Spot Instance is interrupted
    - Spot blocks are designed not to be interrupted
    - When you cancel an active spot request, it does not terminate the associated instance
47. This one felt kinda ambiguous to me
    - Server bound software (licensing and stuff) should go on EC2 dedicated hosts, according to AWS theory
48. I got this one really wrong. Valid tenancy changes:
    - You can change the tenancy of an instance from dedicated to host
    - You can change the tenancy of an instance from host to dedicated
    - https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html
49. I have no rememberance of route 53 endpoints :o
    - Create an inbound endpoint on Route 53 Resolver and then DNS resolvers on the on-premises network can forward DNS queries to Route 53 Resolver via this endpoint
    - Create an outbound endpoint on Route 53 Resolver and then Route 53 Resolver can conditionally forward queries to resolvers on the on-premises network via this endpoint
    - https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver.html
53. With EBS volumes:
    - Data at rest inside the volume is encrypted
    - Any snapshot created from the volume is encrypted
    - Data moving between the volume and the instance is encrypted
55. SQS Long polling can reduce the cost of SQS as it reduces the number of empty receives
56. Remember to only use snowball with only up to 10PB of data
58. Turns out AWS DataSync is capable of saturating a 10Gbps data link
    - Cant use snowballs as they're for offline data transfer only
59. Aww snap I messed up.
    - Correct: With AWS Managed Microsoft AD, you can run directory-aware workloads in the AWS Cloud such as SQL Server-based applications
    - Wrong: AD Connector - Use AD Connector if you only need to allow your on-premises users to log in to AWS applications and services with their Active Directory credentials
62. Dear lord, AWS DAtabase migration service can target Redshift
    - https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html
    - Turns out it can target almost everything :o
65. Oh wow
    - By default, basic monitoring is enabled when you use the AWS management console to create a launch configuration. Detailed monitoring is enabled by default when you create a launch configuration using the AWS CLI
    - https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-monitoring.html

### Practice test 0
29. Is Elasticbeanstalk deployment caching feature a thing?
31. Tricky one, should have 2 answers?
41. I keep fogetting about these big data things. It's either EMR or Redshift
    - Google: "Difference between EMR and Redshift"

### Extra questions not in the course
1. B OK
2. A OK
3. C OK
4. B OK
5. A FAIL (review the answer though)
6. C & D (1/2) 
7. C OK
8. B & E (1/2)
9. C OK
10. D OK
11. C OK
12. B OK
13. C OK
14. B or A ("Database requires block storage")
15. D OK
16. B & C OK
17. C OK
18. B OK
19. C OK
20. B OK
21. A OK
22. A & D OK
23. B & D (??)
24. A & D (1/2)
25. C OK
26. A OK
27. A OK
28. C OK
29. D OK
30. C & E OK
31. ??
32. C OK
33. A & B OK
34. B OK
35. B OK
36. C (??)
37. C OK 
38. D OK
39. B OK
40. A OK
41. D OK
42. B OK
43. B OK
44. A OK
45. A OK
46. B OK
47. B OK
48. C OK
49. D OK
50. B OK
51. C OK
52. D OK
53. D OK
54. A & D OK
55. D OK
56. E (??)
57. B & D OK
58. C OK
59. A & B OK
60. D OK
61. B OK
62. D OK
63. C OK
64. A OK
65. C OK