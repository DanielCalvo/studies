- Notes from dani: It would be _DANK_ to terraform all of this. Also networking is fun!

### 220. Section Introduction
- Amazon VPC
- Amazon Direct Connect!

### 221. CIDR, Private vs Public IPs 
- Classless inter-domain routing
- Used for security group rules or AWS networking in general
- They help define an ip range
- /32 = one IP
- /0 = All ips
- /24, /16, /8, all the ranges you know
- There was a bunch of network theory in this chapter that I didn't write down, I'm rather familiar with network concept from my Cisco training
- https://ipaddressguide.com/cidr

#### Understanding CIDR
- A CIDR has two components:
- The base IP (XX.XX.XX.XX)
    - The Subnet Mask (/26)
    - The base IP represents an IP contained in the range
- The subnet masks defines how many bits can change in the IP
- The subnet mask can take two forms. Examples:
- 255.255.255.0 < less common
- /24 < more common

#### Understanding CIDR subnet masks
- /32 allows for 1 IP = 2^0
- /31 allows for 2 IP = 2^1
- /30 allows for 4 IP = 2^2
- /29 allows for 8 IP = 2^3
- /28 allows for 16 IP = 2^4
- /27 allows for 32 IP = 2^5
- /26 allows for 64 IP = 2^6
- /25 allows for 128 IP = 2^7
- /24 allows for 256 IP = 2^8
- /16 allows for 65,536 IP = 2^16
- /0 allows for all IPs = 2^32
- Cool I never thought about it this way

- Quick memo
    - /32 – no IP number can change
    - /24 - last IP number can change
    - /16 – last IP two numbers can change
    - /8 – last IP three numbers can change
    - /0 – all IP numbers can change

#### Private vs Public IP allowed ranges
- The Internet Assigned Numbers Authority (IANA) established certain blocks of IPV4 addresses for the use of private (LAN) and public (Internet) addresses.
- Private IP can only allow certain values
    - 10.0.0.0 – 10.255.255.255 (10.0.0.0/8) <= in big networks
    - 172.16.0.0 – 172.31.255.255 (172.16.0.0/12) <= default AWS one
    - 192.168.0.0 – 192.168.255.255 (192.168.0.0/16) <= example: home networks
- All the rest of the IP on the internet are public IP

### 222. Default VPC overview
- All accounts have a default VPC!
- New instances are launched in the default VPC if no subnet is specified
- Default VPC has internet connectivity and all instances have a public IP
- You also get a public and a private DNS name for each instance (this was configured in the default VPC)

### Hands on
- There's the default VPC!
- Each VPC has a few subnets, it appears the default VPC has one subnet per AZ in ireland
- Subnets have gateways, ACLs, flow logs and some other features

### 223. VPC Overview and Hands on
- Virual Private Cloud (VPC)
- You can have multiple VPCs in a region (max of 5, soft limit)
- Max CIDR per VPC is 5. For each cider:
    - Min size is /28 (16 IP Adresses)
    - Max size is /16 (65536 IP Adresses)
- Because a VPC is private, only the private IP ranges are allowed! 
- Your VPC CIDR should not overlap with your other networks!

#### Hands on
- Your VPCs > Create VPC
- VPCs by default come with a main route table and a main network ACL
- You add multiple CIDRs on a single VPC. Cool!

### 224. Subnet overview and Hands on
- Subnets are tied to specific AZs
- We'll create a public subnet and a private subnet, yay!
- Public subnets are usually smaller than private subnets as not everything needs to have a private address (but everything on AWS has a private address at least, well, most of the time!)

#### Hands on!
- PublicSubnetA
    - MyVPC
    - eu-west-1a
    - VPC CIDR: 10.0.0.0/16
    - IPv4 CIDR block: 10.0.0.0/24
- PublicSubnetB
    - MyVPC
    - eu-west-1b
    - VPC CIDR: 10.0.0.0/16
    - IPv4 CIDR block: 10.0.0.0/24
- PrivateSubnetA
    - MyVPC
    - eu-west-1a
    - VPC CIDR: 10.0.0.0/16
    - IPv4 CIDR block: 10.0.16.0/20
- PrivateSubnetB
    - MyVPC
    - eu-west-1b
    - VPC CIDR: 10.0.0.0/16
    - IPv4 CIDR block: 10.0.32.0/20

- AWS reserves 5 IP addresses per subnet (first 4 and last 1 IP addresses)
- These IP addresses are not available and cannot be assigned to any instances
- Ex, if CIDR block 10.0.0.0/24, reserved IP are:
    - 10.0.0.0: Network address
    - 10.0.0.1: Reserved by AWS for the VPC router
    - 10.0.0.2: Reserved by AWS for mapping to Amazon-provided DNS
    - 10.0.0.3: Reserved by AWS for future use
    - 10.0.0.255: Network broadcast address. AWS does not support broadcast in a VPC, therefore the address is reserved
- Possible exam question:
    - I need 29 IP addresses in a network, which network size do I use?
    - Don't use /27 (32 IPs) as 5 of these will be reserved and you can't use them. You would need /26

### 225. Internet Gateways & Route tables

#### Hands on
- PublicSubnet(A|B) > Actions > Modify Auto Assign IP settings > Enable auto-assign public IPv4 address (Yes)
- Launching an EC2 instance on the MyVPC and on the public subnet will not work as we set up previously
- Launched an instance in one of the public subnets
- I can't ssh into it! Turns out the instance doesn't have an internet connection, it needs an Internet gateway

#### Internet Gateways
- Internet gateways help our VPC instance connect with the Internet
- They scale horizontally, have high availability (HA) and are redundant
- They must be created separately from the VPC
- One VPC can only be attached to one IGW and vice versa
- IGW is also a NAT device for the instances that have a public IPv4
- On their own, Internet Gateways do not allow internet access.
- Route tables must also be edited!

#### Hands on
- Create IGW
- Attach to MyVPC
- Tried to SSH to instance and... nothing :(
- We also need to change the route tables!
- Created two routing tables
    - PublicRouteTable, PrivateRouteTable
    - Associated the PublicRouteTable to the public subnets
    - Associated the PrivateRouteTable to the private subnets
- Added a route on the PublicRouteTable to 0.0.0.0/0 to go through the IGW we just configured
- Yay now we can ssh into our EC2 instance!

### 226. NAT Instances
- Our instances in our private subnet would have no internet connection! :(
- We can use NAT for this. NAT in AWS comes in two flavours
    - NAT instances (outdated but still on the exam) 
    - NAT gateways (which we'll see in the next lecture)
- NAT instances allow instances in private subnets to connect to the Internet
- NAT instances must be launched in a public subnet
- Must disable EC2 flag on the NAT instance: Source/Destination check
- Must have an Elastic IP attached to the NAT instance
- Route table must be configured to route traffic from private subnets to NAT instance
- Diagram:
    - Create a NAT instance in the public subnet, which will use the route table to talk to the internet
    - We'll use a router + route table on the private subnet to point to the nat instance for internet access

#### Hands on
- EC2
- Look for "NAT" in the AMI list
- Choose the first one provided as a community AMI provided by Amazon
- t2.micro
- On the MyVPC
- PublicSubnetB
- New SG: NAT-SG
    - ssh from everywhere
    - HTTP from 10.0.0.0/16
    - HTTPS from 10.0.0.0/16
- Launch a private instance too
    - Put it in PrivateSubnetA
    - SG: SSH only from 10.0.0.0/16
- On NAT instance on the EC2 dashboard:
    - Right click > Networking > Change Source/Dest. Check > Yes, disable
- Now I can ssh to my public instance and I can reach my private instance
- But the private instance still does not have internet
- Go to route tables and edit PrivateRouteTable, we're going to create a new route
    - 0.0.0.0/0 > To an instance, the nat gateway
- Neat! Now I can connect to my public instance, then the private one from there, and run `sudo yum update`
- Very hard to manage, many moving parts. I agree with Stephane.

#### NAT Instances - Comments- Amazon Linux AMI pre-configured are available
- Amazon Linux AMIs pre-configured for NAT Gateways are available
- Not a highly available / resilient setup out of the box
    - Would need to create ASG in multi AZ + resilient user-data script. Awful
- Internet traffic bandwidth depends on EC2 instance performance. t2.micro wouldn't have a lot of throughput
- Must manage security groups & rules:
    - Inbound:
        - Allow HTTP / HTTPS Traffic coming from Private Subnets
        - Allow SSH from your home network (access is provided through Internet Gateway)
    - Outbound:
        - Allow HTTP / HTTPS traffic to the internet
        
- NAT Instances are old. A much better way of doing things is a NAT Gateway!

### 227. NAT Gateways
- AWS managed NAT, higher bandwidth, better availability, no administration required
- Pay by the hour for usage and bandwidth
- NAT is created in a specific AZ, uses an EIP
- Cannot be used by an instance from the same subnet that we created the NAT Gateway in (only from other subnets)
- Requires an IGW (Private Subnet => NAT => IGW)
- 5 Gbps of bandwidth with automatic scaling up to 45 Gbps
- No security group to manage / required

#### NAT Gateway with High Availability
- NAT Gateway is resilient within a single-AZ
- If you want high availability, you must create multiple NAT Gateways in multiple AZs for fault-tolerance
- There is no cross AZ failover needed because if an AZ goes down it doesn't need NAT

#### Hands on
- Any time a route does not lead anywhere, it's status is "blackhole" (just deleted the NAT Gateway with an instance pointing to it) 
- NAT Gateway > Create NAT Gateway
    - **Remember to create the NAT Gateway on PublicSubnetA**. NAT GWs cannot be used by instances on the same Subnet
    - Allocate Elastic IP Address
    - Create NAT Gateway!
- Route tables > PrivateRouteTable > Routes > Edit rules
    - 0.0.0.0/0 through the NAT Gateway you just created

### Other notes
- This was a cool lecture, check this for more details: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-comparison.html

### 228. DNS Resolution options
- For DNS resolution in a VPC there are two important settings the exam may ask you about
- enableDnsSupport: (DNS Resolution setting)
    - Default True
    - Helps decide if DNS resolution is supported for the VPC
    - If True, queries the AWS DNS server at 169.254.169.253
- enableDnsHostname: (DNS Hostname setting)
    - False by default for newly created VPC, True by default for Default VPC
    - Won’t do anything unless enableDnsSupport=true
    - If True, Assign public hostname to EC2 instance if it has a public IP
    - If you use custom DNS domain names in a private zone in Route 53, you must set both these attributes to true for your instances to resolve that private zone

#### Hands on
- VPC > DemoVPC > Right Click > Edit DNS resolution > Enabled by default
- VPC > DemoVPC > Right Click > Edit DNS hostname > Disabled by default
    - By enabling this feature, instances on PublicSubnets will now have a public DNS entry! (ec2-xxx.eu-west-1.compute.amazonaws.com)
- Route53 > Create Hosted Zone
    - Domain name: dani.internal
    - Comment: Internal domain
    - Type: Private Hosted Zone for Amazon VPC
    - VPC ID: Select the ID of MyVPC
- Then on that zone
    - Create a CNAME with demo.dani.internal pointing to google.com
- This whole thing can be a common exam question!

### 229. NACL & Security groups
#### Network ACLS & Security Group Incoming Request
- Common at the exam: What's the differente between a network ACL and a security group?
- Diagram: You have an EC2 instance on a network and a hypothetical HTTP request is made to this instance
    - First the NACL rules are evaluated. If the request is accepted through the NACL rules, then
    - The SG group rules are evaluated. If the request passes through those rules, then the EC2 instances receives the request
    - The outbound response for this request will be allowed. SGs are stateful, so If an inbound request passes, the outbound request will pass as well, even if there is a rule to deny any traffic from the EC2 instance. If an inbound rule passes, then the outbound rule passes as well.
    - NACL outbound rules are stateless, so the NACL outbound rule **will** get evaluated
    - Both NACL inbound and outbound rules will get evaluated on a request
    - But on a SG if they inbound rule allowed the traffic to go in, then the SG will allow the traffic to go out
    
#### Network ACLS & Security Group Outgoing Request
- Imagine we do `wget google.com`
    - SG outbound rules are evaluated
    - NACL outbound rules are evaluated
- As for the response:
    - NACL inbound rules are evaluated (stateless)
    - SG inbound rules (stateful). Inbound will be allowed on matter what on the SG as it is stateful. The request will come back even if you haven't opened any ports.

#### Network ACLs
- NACL are like a firewall which control traffic from and to a subnet
- Default NACL allows everything outbound and everything inbound
- You define one NACL per Subnet, new Subnets are assigned the Default NACL
- To define NACL rules:
    - Rules have a number (1-32766) and higher precedence with a lower number
    - In other words, rules with a lower number have a higher precedence
    - E.g. If you define #100 ALLOW <IP> and #200 DENY <IP> , the IP will be allowed as 100 is less than 200
    - Last rule is an asterisk (*) and denies a request in case of no rule match
    - AWS recommends adding rules by increment of 100
- A newly created NACL will deny everything
- NACL are a great way of blocking a specific IP at the subnet level

#### Hands on
- VPC > Network ACLs > MyVPC
    - If you look at Inbound Rules, all traffic is allowed
    - If you look at Outbound Rules, all traffic is allowed
    - Create an instance on the public ec2 instance, modify the SG so that rule 80 inbound is allowed and start apache on that machine, so that you can access it through the internets
    - SSH into the instance
- Back to Network ACLs
    - Default NACL
    - Edit inbound rules
    - DENY port 80 with weight 80 from 0.0.0.0/0
    - Aaand now I can't access apache. It works! Even though port 80 is still allowed on the SG, it's blocked with the Network ACL
    - If you change the rule weight to 200, and you still have the allow everything rule with the weight 100, then you will be able to access apache, as the allow everything rule will take precedence
- Now go into the security group for that instance and delete all outbound rules
- You can still do HTTP requests to the instance! This is because since HTTP inbound on port 80 is allowed, all outbound traffic to respond these requests is also allowed
- This works due to the security group being stateful
- Let's add the back the rule to allow all outbound traffic
- Now let's go to network ACLs
    - Deny all outbound traffic
    - Aaand now the site doesn't work anymore
- The difference is discussed here: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Security.html
- Further information is available here: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html
- One last thing: When you create a brand new NACL, everything is denied. You need to add the allow rules to it

### 230. VPC Peering
- VPC peering allows you to connect two VPCs, privately, using AWS's network
- You can make them behave as if they were in the same network
- These VPCs must not have overlapping CIDRs
    - To make this happen, you must create a VPC peering connection between the two VPCs
    - And update some route tables
    - VPC A <-> VPC B <-> VPC C
    - VPC A and VPC C are not connected. You need to connect them directly
- VPC Peering connection is not transitive (must be established for each VPC that needs to communicate with one another)
- You can do VPC peering with another AWS account
- You must update route tables in each VPC’s subnets to ensure instances can communicate

#### VPC Peering - Good to know
- VPC peering can work inter-region, cross-account
- You can reference a security group of a peered VPC (works cross account)

#### Hands on
- Let's connect our default VPC to our MyVPC
- And well, obviously they can't communicate as they are right now (VPCs are isolated by default, not peered)
- VPC dashboard > Peering connections > Create peering connection
    - DemoPeering
    - My account
    - This reagion
    - VPC Requester: MyVPC
    - VPC Accepter: DefaultVPC
    - Once created, right click on the peering connection and click "accept request"
- You need to add route now!
    - Route tables
    - PublicRouteTable
    - Edit route
    - 172.31.0.0/16 > Target is the peering connection
    - DefaultVPCRouteTable
    - Edit route
    - 172.31.0.0/16 > Target is the peering connection

### 231. VPC Endpoints
- VPC endpoints are for you to access amazon services within a private network. So you don't do it through the Internet!
    - Ex: you can have a private endpoint for S3
- Endpoints allow you to connect to AWS Services using a private network instead of the public internet network
- They scale horizontally and are redundant
- They remove the need of IGW, NAT, etc... to access AWS Services
- There are two kinds of VPC endpoints:
    - Interface: provisions an ENI (private IP address) as an entry point (must attach security group) – most AWS services
    - Gateway: provisions a target and must be used in a route table – S3 and DynamoDB
- In case of issues:
    - Check DNS Setting Resolution in your VPC
    - Check Route Tables

#### Hands on
- Let's set up S3 as a gateway!
- EC2 instances > Private instance > Assign an IAM role that has full access to S3
- Assign the role to the private instance you have (or create one)
- SSH into a public instance and then into your private instance
- `aws s3 ls` works on the private instance. Neat!
- Go to end points on the VPC dash > Create Endpoint > Select S3 as the service
- You can see that most services are of interface type, but some of them are of Gateway type. S3 is of Gateway type
- Select MyVPC and add it to the private route table further down on the page 
- Oh now there's a route for S3 on my route table
- `aws s3 ls --region eu-west-1` (remember to include the region as by default s3 will try to talk to another region and won't try to reach your internal gateway, odd)
    - Could be a trick question at the exam too

### 232. VPC Flow logs + Athena
- Flow logs help you capture information about IP traffic going into your interfaces
- There are 3 kinds of flow logs:
    - VPC Flow Logs (everything within VPC)
    - Subnet Flow Logs (everything within subnet)
    - Elastic Network Interface Flow Logs (just for one ENI)
- This helps to monitor & troubleshoot connectivity issues
- Flow logs data can go to S3 / CloudWatch Logs
- Captures network information from AWS managed interfaces too: ELB, RDS, ElastiCache, Redshift, WorkSpaces

#### Flow Log Syntax
- <version> <account-id> <interface-id> <srcaddr> <dstaddr> <srcport> <dstport> <protocol> <packets> <bytes> <start> <end> <action> <log-status>
- Srcaddr, dstaddr helps identify problematic IP (or to filter by them)
- Srcport, dstport helps identity problematic ports
- Action : Success or failure of the request due to Security Group / NACL
- Can be used for analytics on usage patterns, or malicious behavior
- Flow logs example: https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records
- You can query VPC flow logs using Athena on S3 or CloudWatch Logs Insights

#### Hands on
- VPC > DemoVPC > Flow Logs > Create flog low
- Let's send these logs to cloud watch, so before continuing, let's hop back to cloudwatch
- CloudWatch > Log groups > Action > Create log group > vpc-flow-logs
- Now back to the VPC flow log creation screen
    - Filter: all
    - Destination: Cloudwatch logs
    - Log group: vpc-flow-logs
    - IAM role: flowlogsRole (IAM can create them automatically for you)
- Create some more flog logs with the same settings as above (or as similar as possible) and send them to S3 this time
- A bucket policy is added automatically to the bucket
- Popular exam question: "How do you analyze VPC flow logs?" Athena!
- googled `athena vpc flow logs example`
- Ended up here and did all of this: https://docs.aws.amazon.com/athena/latest/ug/vpc-flow-logs.html

```
CREATE EXTERNAL TABLE IF NOT EXISTS vpc_flow_logs (
  version int,
  account string,
  interfaceid string,
  sourceaddress string,
  destinationaddress string,
  sourceport int,
  destinationport int,
  protocol int,
  numpackets int,
  numbytes bigint,
  starttime int,
  endtime int,
  action string,
  logstatus string
)
PARTITIONED BY (`date` date)
ROW FORMAT DELIMITED
FIELDS TERMINATED BY ' '
LOCATION 's3://daniflowlogs/AWSLogs/452788477576/vpcflowlogs/eu-west-1/'
TBLPROPERTIES ("skip.header.line.count"="1");
```

```
ALTER TABLE vpc_flow_logs
ADD PARTITION (`date`='2020-07-05')
location 's3://daniflowlogs/AWSLogs/452788477576/vpcflowlogs/eu-west-1/2020-07-05';
```

```
SELECT *  FROM vpc_flow_logs;
```

### 233. Bastion Hosts
- We can use a Bastion Host to SSH into our private instances
- The bastion is in the public subnet which is then connected to all other private subnets
- Bastion Host security group must be tightened
- Exam Tip: Make sure the bastion host only has port 22 traffic from the IP you need, not from the security groups of your other instances
- We've done this before using the PublicInstance to SSH into the PrivateInstance

### 234. Site to site VPN, Virtual Private Gateway & Customer Gateway
- Virtual Private Gateway:
    - A VPN concentrator is set up on the AWS side of the VPN connection
    - Virtual Private Gateway is created and attached to the VPC from which you want to create the Site-to-Site VPN connection
    - Possibility to customize the ASN (?)
- Customer Gateway:
    - Software application or physical device on customer side of the VPN connection
    - https://docs.aws.amazon.com/vpc/latest/adminguide/Introduction.html#DevicesTested (list all the devices AWS has tested)
- IP Address:
    - Use static, internet-routable IP address for your customer gateway device.
    - If behind a CGW behind NAT (with NAT-T), use the public IP address of the NAT

#### Theoretical hands on 
- Create Customer Gateway
- Set up a Virtual Private Gateway
- Then oyou set up a Site-to-Site VPN connection

### 235. Direct Connect & Direct Connect Gateway
- It's an alternative to using the public internet!
- Provides a dedicated private connection from a remote network (your private datacenter) to your VPC
- The dedicated connection must be setup between your DC and AWS Direct Connect locations
- For this on the AWS side you need to setup a Virtual Private Gateway on your VPC
- Direct connect will allow you to access both public resources (S3) and private resources (EC2) on same connection
- Use Cases:
    - Increase bandwidth throughput - working with large data sets – lower cost
    - More consistent network experience - applications using real-time data feeds
    - Hybrid Environments (on prem + cloud)
- Supports both IPv4 and IPv6

#### Direct Connect Gateway
- If you want to setup a Direct Connect to one or more VPC in many different regions (same account), you must use a Direct Connect Gateway
- Direct connect does not peer VPCs if connected to multiple VPCs. CIDRs must also not overlap
- Common exam question: We need to set up a direct connect to many VPCs from our datacenter. Direct Connect is the answer!

#### Direct Connect - Connection Types
- Dedicated Connections: 1Gbps and 10 Gbps capacity
    - Physical ethernet port dedicated to a customer
    - To get that port you must first make a request to AWS, then completed by AWS Direct Connect Partners
- Hosted Connections: 50Mbps, 500 Mbps, to 10 Gbps
    - Connection requests are made via AWS Direct Connect Partners
    - Capacity can be added or removed on demand (so if you have a big migration you can just increase connection speed for a week)
    - 1, 2, 5, 10 Gbps available at select AWS Direct Connect Partners
- Lead times are often longer than 1 month to establish a new connection

#### Theoretical hands on 
- Direct Connect is its own service (not in the VPC)
- You can then choose a location and request a direct connect

#### Direct Connect - Encryption
- Data in transit in direct connect is not encrypted but is private
- If you need encryption, you can combine AWS Direct Connect + VPN provides an IPsec-encrypted private connection
- Good for an extra level of security, but slightly more complex to put in place

### 236. Egress only Internet Gateway
- Egress means outgoing!
- Egress only Internet Gateway is for IPv6 only
- Similar function as a NAT, but a NAT is for IPv4
    - It's the same as a NAT, but for IPv6
    - This service allows our IPv6 instances to access the Internet but not be accessible
- Why do we need this though? Well, good to know: IPv6 are all public addresses
- Therefore all our instances with IPv6 are publicly accessibly
- Egress Only Internet Gateway gives our IPv6 instances access to the internet, but they won’t be directly reachable by the internet
    - This is useful if you don't want your IPv6 instances to be publicly accessible
- After creating an Egress Only Internet Gateway, edit the route tables

#### Hands on
- VPC > Egress only Internet Gateway > Create Egress only Internet Gateway
    - Select MyVPC, create
- Go on a route table, and add an IPv6 rule and add the egress only internet gateway and save the route

### 237. AWS PrivateLink - VPC Endpoint Services
### Exposing services in your VPC to Other VPC
- Problem: We have an application in our account and we want to expose that application to other VPCs in other accounts
- Option 1: make it public
    - Goes through the public www
    - Tough to manage access
- Option 2: VPC peering
    - Must create many peering relations (doesn't scale that well)
    - Opens the whole network

#### AWS PrivateLink (aka VPC Endpoint Services)
- Most secure & scalable way to expose a service to 1000s of VPC (own or other accounts)
- Does not require VPC peering, internet gateway, NAT, route tables...
- Requires a network load balancer (Service VPC) and ENI (Customer VPC)
- If the NLB is in multiple AZ, and the ENI in multiple AZ, the solution is fault tolerant!
- Private link links a NLB in the service VPC to a ENI on the customer VPC
- The ENI on the customer VPC makes it look like the application on the customer VPC is in their network

#### Hands on
- VPC > Create Endpoint service
    - Select a NLB
    - And then you would create it
- On the target VPC
    - VPC > Endpoints
    - Create endpoints
    - Find service by name
    - And they would be connected!
- Possible exam question: How to expose services from one VPC to hundreds of other VPCs? Think PrivateLink!

### 238. AWS ClassicLink 
#### EC2-Classic & AWS ClassicLink (deprecated)
- How it started: EC2-Classic: Instances run in a single network shared with other customers
- Then Amazon came up with Amazon VPC: your instances run logically isolated to your AWS account
- ClassicLink allows you to link EC2-Classic instances to a VPC in your account
    - Must associate a security group
    - Enables communication using private IPv4 addresses
    - Removes the need to make use of public IPv4 addresses or Elastic IP addresses
- These options/questions are likely to be distractors at the exam (and not a correct answer?)

### 239. VPN CloudHub
- Provide secure communication between sites, if you have multiple VPN connections
- Low cost hub-and-spoke model for primary or secondary network connectivity between locations
- It’s a VPN connection so it goes over the public internet
    - Notes from Dani: Appears to be a way to connect multiple on-prem installations/DCs to each other, not to AWS services
    - Low cost hub and spoke model to connect multiple physical networks together securely over the internet

### 240. Transit Gateway
- It's to have a star gateway! Apparent function seems to be to simplify network topologies
    - Cool diagram here: https://aws.amazon.com/transit-gateway/
- For having transitive peering between thousands of VPC and on-premises, hub-and-spoke (star) connection
- Regional resource, but it can work cross-region
- Can be shared cross-account using Resource Access Manager (RAM)
- You can peer Transit Gateways across regions
- Using route Tables: You can limit which VPC can talk with other VPC
- Works with Direct Connect Gateway, VPN connections
- Supports IP Multicast (not supported by any other AWS service)
- Possible exam question: We need to support IP multicast in our AWS infra, what service do we use? Transit gateway!

#### Theoretical hands on
- Not free :(
- You pay per attachment and per GB of data processed
- VPC > Create Transito Gateway
    - You don't need to know it for the exam in depth
- You can manage attachments to the transit gateway (VPC, VPN, Peering connections)
- You would define your route tables to define how each of these attachments can talk to each other
- You can use the Network Manager to centrally manage your network and visualize your global network in a dashboard
- If you go to Resource Access Manager (RAM)
    - You can see how you're able to share your transit gateway accross multiple accounts
- Remember that it simplifies a network topology!

### 241. VPC Section Summary

#### VPC Section Summary (1/3)
- CIDR: IP Range
- VPC: Virtual Private Cloud > We define a list of IPv4 & IPv6 CIDR
- Subnets: Tied to an AZ, we define a CIDR for the subnet
- Internet Gateway: at the VPC level, provides IPv4 & IPv6 Internet Access
- Route Tables: Must be edited to add routes from Subnets to the IGW, VPC Peering, Connections, VPC Endpoints, etc...
- NAT Instances: Gives internet access to instances in private subnets. Old, must be setup in a public subnet, disable Source / Destination check flag
- NAT Gateway: Managed by AWS, provides scalable internet access to private instances, IPv4 only
- Private DNS + Route 53: enable DNS Resolution + DNS hostnames (VPC)
- NACL: Stateless, subnet level rules for inbound and outbound, don’t forget ephemeral ports otherwise some traffic may not go through
- Security Groups: Stateful, operate at the EC2 instance level

#### VPC Section Summary (2/3)
- VPC Peering: Connects two VPC with non overlapping CIDR, VPC Peering is non transitive, if you peer A <> B <> C, A and C can't communicate, you need to peer them too
- VPC Endpoints: Provides private access to AWS Services (S3, DynamoDB, CloudFormation, SSM) within VPC (without an Internet connection if desired)
- VPC Flow Logs: Can be setup at the VPC / Subnet / ENI Level, for ACCEPT and REJECT traffic, helps identifying attacks, analyze using Athena or CloudWatch Log Insights
- Bastion Host: Public instance to SSH into, that has SSH connectivity to instances in private subnets
- Site to Site VPN: Setup a Customer Gateway on DC, a Virtual Private Gateway on VPC, and site-to-site VPN over public internet
- Direct Connect: Setup a Virtual Private Gateway on VPC, and establish a direct private connection to an AWS Direct Connect Location
- Direct Connect Gateway: Setup a Direct Connect to many VPC in different regions
- Internet Gateway Egress: like a NAT Gateway, but for IPv6

#### VPC Section Summary (3/3)
- Private Link / VPC Endpoint Services:
    - Connects services privately from your service VPC to customers VPC
    - Doesn’t need VPC peering, public internet, NAT gateway, or route tables
    - Must be used with Network Load Balancer & ENI
- ClassicLink: connect EC2-Classic instances privately to your VPC (deprecated!)
- VPN CloudHub: hub-and-spoke VPN model to connect your sites
- Transit Gateway: transitive peering connections for VPC, VPN & DX

### 242. Section Clean up
- Done!

### Quiz!
- Pending!

### 243. Networking Costs in AWS
- Traffic going into EC2 instances is free
- Any private traffic between 2 instances in the same AZ (going through their private IP) is free
- Cross-AZ traffic 
    - $0.02 per GB if using Public IP / Elastic IP 
    - $0.01 per GB if using Private IP 
- Cross region traffic
    - $0.02 per GB
- Use a private IP instead of a public IP for savings and better network performance
- Use same AZ for maximum savings (at the cost of high availability)
- Possible exam question: We have an RDS database and we want to create a read-replica and do some analytics on top of this read replica, how do we create a read read replica the cheapest way / using the least money possible?
    - Create that read replica in the same AZ!
    - You're not going to be charged anything to replicate one DB to another in terms of network cost
    - But if you create it in another AZ, then you'll pay 1 cent per GB of data transfered between the databases