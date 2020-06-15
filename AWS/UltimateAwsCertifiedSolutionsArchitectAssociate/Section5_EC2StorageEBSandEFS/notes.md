- Chapter status: Studied carefully, still need to do the quiz
- Lecture 80 was a bit confusing
- Docs appear to be here: See https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/raid-config.html
- Terraforming an EFS FS with two EC2 machines that mount it would be cool


### 53. EBS intro
- An EC2 machine loses its root volume (main drive) when it's manually terminated
- Unexpected terminations might happen from time to time (AWS would email you)
- You need a way to store your data somewhere! 
- An EBS (Elastic Block Store) volume is a network drive you can attach to your instances while they run
- It allows your instances to persist data
- So if you have a DB, it would be smart to place your db data on the EBS volume

#### EBS Volume
- It's a network drive (not a physical drive)
    - It uses the network to communicate to the instance, so there might be a bit of latency
    - It can be dettached from an EC2 instance and attached to another one quickly (as long as they're in the same AZ)
- It's locked to an AZ
    - An EBS volume in us-east-1a cannot be attached to us-east-1b
    - To move a volume across, you need to snapshot it
- EBS volumes have a provisioned capacity
    - You get billed for all the provisioned capacity (not for what you actually use!)
    - You can increase the capacity of the drive over time

#### Diagram
- It appears you can attach many EBS volumes to an instance, but an EBS volume can only be attached to one instance

#### EBS Volume types
- GP2: General purpose SSD that balances price and performance for a wide variety of workloads
- IO1: High performance SSD for mission critical, low latency or high throughput workloads
- STI: Low cost HDD designed for frequently acessed, throughput intensive workloads
- SCI: Lowest cost HDD designed for less frequently accessed workloads

- EBS volumes are characterized in size, throughput, IOPS (I/O Ops Per Sec)
- When in doubt, always check the AWS documentation!
- Only GP2 and IO1 can be used for boot volumes

### 54. EBS Intro Hands on 
- When creating an EC2 instance, on `Step 4: Storage`, you can add an EBS volume
- Created an EC2 instance on the EC2 console and added an EBS volume to it on create time
- Then did `lsklb`, `mkfs`, editec fstab and all that and mounted the partition
- Volume shows up in "Volumes" under "Elastic Block Store" on the left menu of the EC2 dashboard

### 55. EBS Volume Types Deep Dive
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

#### EBS - Volume Types Summary
- gp2: General Purpose Volumes (cheap)
    - 3 IOPS / GiB, minimum 100 IOPS, burst to 3000 IOPS, max 16000 IOPS
    - Size range: 1 GiB – 16 TiB, adding a TB gives you 3000 more IOPS
- io1: Provisioned IOPS (expensive)
    - Min 100 IOPS, Max 64000 IOPS (Nitro instances) or 32000 (other)
    - Size range: 4 GiB - 16 TiB. Size of volume and IOPS are independent
- st1: Throughput Optimized HDD
    - 500 GiB – 16 TiB , 500 MiB/s throughput
- sc1: Cold HDD, Infrequently accessed data
    - 500 GiB – 16 TiB , 250 MiB/s throughput

### 56. EBS Operation: Snapshots
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

### 57. EBS Operation: Volume Migration
- EBS volumes are only locked to a specific AZ
- To migrate it to a different AZ (or region)
    - Snapshot the volume
    - (optional) Copy the volume to a different region
    - Create a volume from the snapshot in the AZ of your choice

### 58. EBS Operation: Volume Encryption
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

### 59. EBS vs Instance Store
- Some instances do not come with root EBS volumes
- Instead, they come with "Instance Store" (ephemeral storage)
- Instance store is physically attached to the machine (EBS is a network drive)
- Pros:
    - Better I/O performance
    - Good for buffer / cache / scratch data / temporary content
    - Data survives reboots
- Cons:
    - On stop or termination, the instance store data is lost
    - You can't resize the instance store
    - Backups operated by the user (can't right click and backup through the UI)
- Popular exam question: Should we use Elastic block store or Instance Store? Is your data ephemeral? If so, then yes

#### Local EC2 Instance Store
- Physical disk attached to the physical server where your EC2 is
- Very high IOPS (it's a physical connection!) much higher than EBS
- Disks can be up to 7.5 TiB (can change over time) disks can be stripped (?) to reach 30 TiB (can also change over time)
- It's block storage, just like EBS. From the instance standpoint, it's just a disk.
- Cannot be increased in size
- There is risk of data loss if there is a hardware failure

#### Hands on
- EC2 > Launch instance
- On the instance type (Step 2) you need to choose one that supports instance store (larger ones, t2.micro doesn't)
- Volume type will probably show as "Ephemeral0"

### 60. EBS RAID Configurations
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

### 61. EFS Overview
- Elastic File system!
- It's a managed NFS (Network file system) that can be mounted on many EC2 instances
- EFS works with EC2 instances in Multi-AZ
- High available, scalable, expensive (3x GP2), pay per data used
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
- You can also backup EFS-to-EFS (incremental, you can choose frequency)
- Supports encryption at rest using KMS

#### EFS - Performance & Storage Classes
- EFS Scale
    - Meant to support 1000s of concurrent NFS clients with more than 10GB/s throughput
    - Advertised as "Grow to Petabyte-scale NFS, automatically"
- Performance modes (set at EFS creation time)
    - General purpose (default): latency-sensitive use cases (web server, CMS, etc)
    - Max I/O: Higher latency, throughput, parellelism used (many many files, big data, media processing)

- There are different storage tiers on EFS. There is a lifecycle management feature that can more files automatically for you
    - Standard tier: For frequently accessed files
    - Infrequent access (EFS-IA) tier: Cheaper to store files, but with a retrieval fee

### 62. EFS Hands On
- First of all, create an SG, no inbound rules, all outbound rules, default SG
- Create some EC2 instances, make sure they're in the same SG as the EFS
- Used the EC2 mounting instructions to mount on the same VPC
- You have to add the NFS port as an inbound rule for this to work
- Exam might ask if your connection from EC2 to EFS times out, why? SG issue


### 63. Section clean up
- Delete everything fam

### 64. EBS vs EFS

#### EBS volumes:
- Can be attached to only one instance at a time
- Are locked to a specify AZ
- gp2: IO Increases if the disk size inreases
- ioI: Can increase IO independently
- To migrate an EBS volume across AZs:
    - Take a snapshot
    - Restore this snapshot to another AZ
    - EBS backups use IO and you shouldn't run them while your application is handling a lot of traffic
- Root EBS volumes of instances get terminated by default if the EC2 instances gets terminated (you can disable this behaviour)

#### EFS
- Can be mounted to hundreds or thousands of instances across AZs
- EFS shares can be used to share website files, such as wordpress
- Only for Linux (POSIX filesystem)
- EFS is more expensive than EBS (about 3x)
- EFS-IA can be leveraged for cost savings

- Remember the differences between EFS, EBS and Instance Store!

### Quiz
- 6/7