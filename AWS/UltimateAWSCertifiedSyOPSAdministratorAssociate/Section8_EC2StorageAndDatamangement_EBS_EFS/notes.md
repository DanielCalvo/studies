
### 75. Section intro

```text

```

### 76. EBS intro

```text
What is an EBS volume? Elastic Block Store
You need a way to store your data somewhere!
EBS is a network drive and it allows you to persist data
Can be detached from one instance and attached to another
Locked to an AZ (ex: eu-west-1a)
You get billed for the provisioned capacity (not the used capacity)

Volume scome un 4 types:
GP2: General purpose SSD
IO 1: High performance SSD
STI: Low cost HDD
SCI: Lowest cost HDD designed for less frequently accessed workloads
```

### 77. EBS Intro Hands on 

```text
Created an EC2 instance on the EC2 console and added an EBS volume to it on create time
Then formatted the partition and mounted it
```

### 78. EBS Volume Types Deep Dive
 
```text
GP2: Recommended for most workloads. Size: 1gb to 16tb. IOPS can burst.
IO1: Critical business apps. High IOPS.
ST1: Streaming workloads that require consistent, fast throughput at a low price. Max IOPS of 500.
SC1: Storage of data that is infrequently accessed. Max IOPS of 250.

```

### 79. EBS Volume Burst

```text
GP2 can burst up to 3000 IOPS. Similar concept to t2 instances, your IO can burst and you have burst credits.
You can use cloudwatch to monitor your I/O balance.

```
### 80. EBS Computing Throughput

```text
Wew a bunch of theory on what output is faster
```

### 81. EBS Operation: Volume Resizing
```text
You can resize EBS volumes! You need to repartition your drive after that though
```

### 82. EBS Operation: Snapshot
```text

```
### 83. EBS Operation: Volume Migration
```text

```

### 84. EBS Operation: Volume Encryption
```text

```

### 85. EBS vs Instance Store
```text

```

### 86. EBS for SysOps
```text

```

### 87. EBS RAID Configurations
```text

```

### 88. CloudWatch & EBS
```text

```

### 89. EFS Overview
```text

```

### 90. EFS Hands On
```text

```

### 91. Section Clean up
```text

```

### Quiz