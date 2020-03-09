
## Compute
### App Engine
- Software as a Service. Equivalent to Heroku.
### Compute Engine
- Virtual machines. Equivalent to Amazon EC2
#### VM Instances
- Those are... VMs. `¯\_(ツ)_/¯`
#### Managed instance groups
- MIGs allow you autoscale, autoheal, have regional deployments and automatic updates.
- Good for
    - Stateless workloads (such as web frontend)
    - Stateless batch high performance jobs
#### Instance Templates
- A resource that you can use to create VM instances and manage instance groups. Useful for instances in MIGs
#### Sole Tenant nodes
- Physical servers that run only the VMs on your project, so you don't share the server with anyone else. Sounds expensive.
#### Disks
- Persistent disks you can add to your machines, mostly
#### Snapshots
- Snapshots of disks of Compute Engine machines
#### Images
- Equivalent of AWS AMI. Operating system images to create boot disks for your instances
#### TPUs
- Tensorflow processing units
#### Commited use discounts
- You can purchase commitments to earn discounts. Think of long term VM rental, for lower prices
#### Metadata
- Key/Value pairs you can attach to instances
#### Health checks
- Determines if a VM is healthy. Used with Load Balancers
#### Zones
- Summary of how many instances you have and in which zones
#### Network Endpoint groups
- (BETA) Can be used as a backend for some load balancers
#### Operations
- A log of things you did in Compute Engine. Neat!
#### Security Scans
- Automatically scan your App Engine, Compute Engine, and Google Kubernetes Engine apps for common vulnerabilities.
#### Settings
- For now only lets you export reports of Compute Engine usage
### Google Cloud Functions
- Google Cloud Functions is a lightweight, event-based, asynchronous compute solution that allows you to create small, single-purpose functions that respond to cloud events without the need to manage a server or a runtime environment
### Cloud Run
- Cloud Run is a fully managed compute platform that enables you to run stateless containers that are invocable via HTTP requests.

## Storage
### Big Table
- Cloud Bigtable is a fully managed, wide-column NoSQL database that offers low latency and replication for high availability (Proprietary)
### Datastore
- Older version of Firestore
### Firestore
- Proprietary document storage database. Similar to CouchDB and MongoDB. Newer version of Datastore
### Filestore
- An instance is a fully managed network-attached storage system you can use with your Google Compute Engine and Kubernetes Engine instances. Neat!
### Storage
- File storage. Equivalent to Amazon S3
### SQL
- Fully managed, relational MySQL, PostgreSQL, and SQL Server databases.
### Spanner
- Google's proprietary SQL DB. Equivalent to AWS Aurora
### Memorystore
- Creates and manages Redis instances on the Google Cloud Platform.

## Networking
### VPC networks
#### VPC networks
#### External IP Addresses
- An external IP address makes an attached project resource available on another Google Cloud Platform network or the Internet.
#### Firewall rules
- Firewall rules control incoming or outgoing traffic to an instance. 
#### Routes
- Default Internet routes & inter-network routing
#### VPC network peering
- Cloud VPC Network Peering lets you privately connect two VPC networks, which can reduce latency, cost, and increase security. 
#### Shared VPC
- Unavailable :(
#### Serverless VPC access
- Unavailable :(
#### Packet Mirroring
- Beta. 

### Network Services
#### Load balancing
- Load balancers distribute incoming network traffic across multiple VM instances to help your application scale.
- There are Internal/External HTTP/S and TCP/UDP load balancers.
#### Cloud DNS
- DNS zones let you define your namespace. You can create public or private zones. 
#### Cloud CDN
- Google Cloud CDN uses Google's globally distributed edge points to cache your HTTP(S) content closer to users, resulting in faster delivery and reduced serving costs
#### Cloud NAT
- Cloud NAT lets your Compute Engine instances and Kubernetes Engine container pods communicate with the internet using a shared, public IP address.
#### Traffic Director


### Hybrid Connectivity
### Network Service Tiers
- Network Service Tiers enable you to optimize network quality and performance vs. cost for your resources and projects
### Network Security
### Network Intelligence

## Stack driver
### Monitoring
### Debug
### Trace
### Logging
### Error Reporting
### Profiler

## Tools
### Cloud Build
- Cloud Build, Google Cloud’s continuous integration (CI) and continuous delivery (CD) platform, lets you build software quickly across all languages.
### Cloud Tasks
- Cloud Tasks lets you run background tasks outside of a user request, or process a batch of tasks at once.
### Container Registry
### Deployment Manager
### Endpoints
### Identity Platform
### Source Repositories
### Private Catalog