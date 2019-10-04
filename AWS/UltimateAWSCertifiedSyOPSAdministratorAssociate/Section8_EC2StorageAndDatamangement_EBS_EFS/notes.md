- Chapter status: Studied carefully, still need to do the quiz
- Lecture 80 was a bit confusing

### 75. Section intro
- Amazon EFS
- Amazon EBS
- Let's kick ass!

### 76. EBS intro
- An EC2 machine loses its root volume (main drive) when it's manually terminated
- Unexpected terminations might happen from time to time (AWS would email you)
- You need a way to store your data somewhere! 
- An EBS (Elastic Block Store) volume is a network drive you can attach to your instances while they run
- It allows your instances to persist data
- So if you have a DB, it would be smart to place your db data on the EBS volume

#### EBS Volume
- It's a network drive (not a physical drive)
    - It uses the network to communicate to the instance, so there might be a bit of latency
    - It can be dettached from an EC2 instance and attached to another one quickly
- It's locked to an AZ
    - An EBS volume in us-east-1a cannot be attached to us-east-1b
    - To move a volume across, you need to snapshot it
- EBS volumes have a provisioned capacity
    - You get billed for all the provisioned capacity
    - You can increase the capacity of the drive over time

#### Diagram
- It appears you can attach many EBS volumes to an instance, but an EBS volume can only be attached to one instance

#### EBS Volume types
- GP2: General purpose SSD that balances prive and performance for a wide variety of workloads
- IO1: High performance SSD for mission critical, low latency or high throughput workloads
- STI: Low cost HDD designed for frequently acessed, throughput intensive workloads
- SCI: Lowest cost HDD designed for less frequently accessed workloads

- EBS volumes are characterized in size, throughput, IOPS (I/O Ops Per Sec)
- When in doubt, always check the AWS documentation!
- Only GP2 and IO1 can be used for boot volumes

### 77. EBS Intro Hands on 
- When creating an EC2 instance, on `Step 4: Storage`, you can add an EBS volume
- Created an EC2 instance on the EC2 console and added an EBS volume to it on create time
- Then did `lsklb`, `mkfs`, editec fstab and all that and mounted the partition

### 78. EBS Volume Types Deep Dive
#### GP2
- Recommended for most workloads
- System boot volumes
- Virtual Desktops
- Low-latency interactive apps
. Development and test environments
- 1 GiB - 16 TiB
- Small GP2 volumes can burst to 3000
- Max IOPS is 16.000
- 3 IOPS per GB, means at 5334GB we are at the max IOPS

#### IO1
- Critial business applications that require sustained IOPS performance, or more than 16.000 IOPS per volume (the GP2 limit)
- You can use this for large database workloads (Mongo, Cassandra, MySQL, etc)
- 4GiB - 16 TiB
- IOPS is provisioned (PIOPS). Mix 100, Max 64.000 (Nitro instances), else MAX 32.000 (other instances)
- The maximum ratio of provisioned IOPS to requested volume size in GB is 50:1

#### ST1
- Streaming workloads that require consistent, fast throughput at a low price. Max IOPS of 500.
- Big data, data warehousing, log processing, apache kafka
- Cannot be boot volume
- Sizes: 500gb to 16tB
- Max IOPS of 500
- Max throughput of 500MiB/s - can burst
- Not too popular on the exam

#### SC1
- Storage of data that is infrequently accessed
- Scenarios where the lowest storage cost is important
- Cannot be a boot volume
- Sizes: 500gb to 16tB
- Max IOPS of 250
- Max throughput of 250MiB/s - can burst


#### Hands on
- EC2 > Volumes > Create Volume
- You can see all the options, numbers and settings there
- Exam tip: As soon as you reach the I/O limit on GP2, then IO1 is the answer for your performance issues as you can go to 32k IOPS or 64k IOPS for Nitro instances

### 79. EBS Volume Burst

#### GP2 Volumes burst
- If your GP2 volume is lower than 1000GB (IOPS less than 3000) it can burst to 3000 IOPS performance
- It is a similar concept to t2 instances with their CPU 
- You accumulate "burst credit over time", which allows your volume to have good performance when needed
- The bigger the volume, the faster you fill up your "burst credit balance"
- What happens if I empty my I/O credit balance?
    - The maximum I/O you get becomes the baseline you paid for
    - If you see the balance as being 0 all the time, increase the GP2 volume or switch to IO1
    - You can use cloudwatch to monitor your I/O credit balance (popular exam question!)
- The burst concept also applies to ST1 or SC2 (for increased throughput)

#### Hands on
- EC2 Dashboard > EBS > Click on a Volume > Monitoring > Burst Balance (percent)

### 80. EBS Computing Throughput

- GP2
    - Throughput MiB/S = (Volume size in GiB) x (IOPS per GiB) x (I/O size in KiB)
    - Ex: 300 I/O operations per second * 256 KiB per I/O Operations = 75 MiB/s
    - 100 gb volume x 3 IOPS per GiB x Max I/O you can ever get on EBS GP2 is 256KiB 
    - Limited to a max of 250MiB/s (means volume >= 334GiB won't increase throughput)
- IO1
    - Throughput MiB/S = (Volume size in GiB) x (I/O size in KiB)
    - The throughput limit of IO1 is 256KiB/s for each IOPS provisioned
    - Limit to a max of 500MiB/s (at 32.000 IOPS) and 1000MiB/s (at 64.000 IOPS)

### 81. EBS Operation: Volume Resizing
- You can resize EBS volumes.
- You can increase the size and the IOPS (only in IO1)
- After resizing your EBS volume you'll need to repartition your drive
- After increasing the size, it's possible for the volume to still be in the "optimization" phase for a while. The volume is still usable!

#### Hands on
- EBS > Volumes > Right click > Modify volume

### 82. EBS Operation: Snapshot
- Popular on the exam!
- Snapshots are incremental - only backup change blocks
- EBS backups use IO and you shouldn't run them while your application is handling a lot of traffic 
- Snapshots are stored in S3 but you won't directly see them
- Not necessary to detach a volume do to do a snapshot, but recommended
- Max 100.000 snapshots per account
- You can copy your snapshot across AZ or region
- Can create image (AMI) from snapshot
- EBS volumes restored by snapshots need to be pre-warmed (using fio or dd command to read the entire volume)
- Snapshots can be automated using Amazon Data Lifecycle manager

#### Hands on
- EBS Volumes > Right click > Create snapshot
- You can create a volume or an image from the snapshot once it's completed
- You can also copy it to another region
- If you go to lifecycle manager just below snapshots on the left side menu, you can create a Snapshot Lifecycle Policy!
- You can automate snapshot creation based on tags or schedules

### 83. EBS Operation: Volume Migration
- EBS volumes are only locked to a specific AZ
- To migrate it to a different AZ (or region)
    - Snapshot the volume
    - (optional) Copy the volume to a different region
    - Create a volume from the snapshot in the AZ of your choice

### 84. EBS Operation: Volume Encryption
- When you create an encrypted EBS volume, you get the following
    - Data at rest is encrypted inside the volume
    - All the data in flight between the instance and the volume is encrypted
    - All snapshots are encrypted
    - All volumes created from the snapshots are encrypted
- Encryption and decryption are handled transparently (you have nothing to do)
- Encryption has a minimal impact on latency
- EBS encryption leverages keys from KMS (AES-256)
- Copying an unencrypted snapshot allows encryption
- Snapshots of encrypted volumes are encrypted

#### Encryption: Encrypt and unencrypted EBS volume
- Create an EBS snapshot of the volume
- Encrypt the EBS snapshot (using copy)
- Create a new EBS volume from the snapshot (the volume will also be encrypted)
- Now you can attach the encrypted volume to the original instance

### 85. EBS vs Instance Store
- Some instances do not come with root EBS volumes
- Instead, they come with "Instance Store" (ephemeral storage)
- Instance store is physically attached to the machine (EBS is a network drive)
- Pros:
    - Better I/O performance
    - Good for buffer / cache / scratch data / temporary content
    - Data survives reboots
- Cons:
    - On stop or termination, the instance is lot
    - You can't resize the instance store
    - Backups operated by the user
- Popular exam question: Should we use Elastic block store or instance store? Is your data ephemeral? If so, then yes

#### Hands on
- EC2 > Launch instance
- On the instance type (Step 2) you need to choose one that supports instance store (t2.micro doesn't)
- Volume type will probably show as "Ephemeral0"

### 86. EBS for SysOps
If you plan to use the root volume of an instance after it's terminated:
    - Set the delete termination flag to "No"
    - You can see this option when creating the EC2 instance
- If you use EBS for high performance, use EBS-optimized instance types (AWS has a list of EBS optimized instances)
- If an EBS volume is unused, you still pay for it
- For cost saving over a long period, it can be cheaper to snapshot a volume and restore it later if it's unused

#### EBS Troubleshooting
- High wait time or slow reponse for SSH > Increase IOPS
- EC2 won't start with EBS volume as root: Make sure volume names are properly mapped (/dev/xvdb instead of /dev/xvda for example)
- After increasing a partition size, you still need to repartition to use the incremented storage (xfs_growfs for example)

### 87. EBS RAID Configurations
- EBS is already redundant storage (replicated within AZ)
- But what if you want to increase IOPS to say, 100.000 IOPS?
- What if you want to mirror your EBS volumes?
- You would mount volumes in parallel in RAID settings!
- RAID is possible as long as your OS supports it
- Some raid options are: (0 & 1 are asked at the exam)
    - RAID 0 
    - RAID 1
    - RAID 5 (not recommended for EBS)
    - RAID 6 (not recommended for EBS)
- We'll explore RAID 0 and RAID 1

#### RAID 0
- Increased performance
- Combining 2 or more volumes and getting the total disk space and I/O
- But if one disk fails, all the data is lost
- Use cases would be:
    - An application that needs a lot of IOPS and doesn't need fault-tolerance
    - A database that has replication already built—in
- Using this, we can have a large disk with a lot of IOPS
    - Two 500 GB Amazon EBS IO1 volumes with 4.000 provisioned IOPS each will create a  IOOO GiB RAID 0 array with an available bandwidth of 8.000 IOPS and 1.OOO MB/s of throughput   

#### RAID 1
- Mirrored volumes
- If one disk fails, our logical volume is still working
- We have to send the data to two EBS volumes at the same time (2x the network usage)
- Use case
    - Application that needs increased fault tolerance
    - Application were you need to service disks
- For example
    - Two 500GiB Amazon EBS IO1 volumes with 4.000 provisioned IOPS a 500 GiB RAID 1 with an available bandwidth of 4.000 IOPS and 500MB/s of throughput
- You have to do these on your own, but on the exam they don't expect you to know how to do it, they just want to make sure you remember the concepts

### 88. CloudWatch & EBS
- Important EBS Volume CloudWath metrics
    - VolumeIdleTime: Number of seconds when no read/write is submitted
    - VolumeQueueLenght: Number of operations waiting to be executed. Maybe you don't have enough IOPS of maybe there's an application issue
    - BurstBalance: If it becomes 0, you need a volume with more IOPS
- GP2: Metrics reported every 5 minutes
- IO2: Metrics reported every 1 minutes
- EBS volumes have status checks
    - Ok: The volume is performing well
    - Warning: Performance below expected
    - Impaired: Stalled, performance severaly degraded
    - Insufficient data: Metric data collection in progress

### 89. EFS Overview
- Elastic File system!
- It's a managed NFS (Network file system) that can be mounted on many EC2 instances
- EFS works with EC2 instances in Multi-AZ
- High available, scalable, expensive (3x GP2), pay per use
- The difference though is that you don't provision capacity, you pay for what you use. So it could end up being cheaper than GP2

#### EFS, part two
- Use cases: Content management, web serving, data sharing, Wordpress
- Uses NFS v4.1 protocol
- Uses security group to control access to EFS  
- Compatible with Linux based AMIs (not Windows)
- Performance modes:
    - General purpose (default)
    - Max I/O: Used for when you have thousands of EC2 instances accessing EFS
- EFS has a feture to sync filesystems on premise to EFS
- You can also backup EFS-to-EFS (incremental, you can choose frequency)-
- Supports encryption at rest using KMS

### 90. EFS Hands On
- First of all, create an SG, no inbound rules, all outbound rules, default SG
- Create some EC2 instances, make sure they're in the same SG as the EFS
- Used the EC2 mounting instructions to mount on the same VPC
- You have to add the NFS port as an inbound rule for this to work
- Exam might ask if your connection from EC2 to EFS times out, why? SG issue

### 91. Section Clean up


### Quiz