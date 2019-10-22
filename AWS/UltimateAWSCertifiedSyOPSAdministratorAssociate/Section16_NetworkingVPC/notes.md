- Review/redo lecture 203 on nat instances

### 197. Section intro
- Amazon VPC
- Amazon Direct Connect!

### 198. CIDR, Private vs Public IPs 
- Classless inter-domain routing
- Used for security group rules or AWS networking in general
- They help define an ip range
- /32 = one IP
- /0 = All ips
- /24, /16, /8, all the ranges you know
- There was a bunch of network theory in this chapter that I didn't write down, I'm rather familiar with network concept from my Cisco training
- https://ipaddressguide.com/cidr

### 199. Default VPC overview
- All accounts have a default VPC!
- New instances are launched in the default VPC if no subnet is specified
- Default VPC has internet connectivity and all instances have a public IP
- You also get a public and a private DNS name for each instance 

### Hands on
- There's the default VPC!
- Each VPC has a few subnets, it appears the default VPC has one subnet per AZ in ireland
- Subnets have gateways, ACLs, flow logs and some other features

### 200. VPC Overview and Hands on
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
- You can have multiple CIDRs on a single VPC. Cool!

### 201. Subnet overview and Hands on
- Subnets are tied to specific AZs
- Well create a public subnet and a private subnet, yay!

#### Hands on!
- Created some subnets in different AZs with different IP ranges
- AWS reserves 5 IP addresses per subnet (first 4 and last 1 IP addresses)
- These IP addresses are not available and cannot be assigned to any instances
- Possible exam question:
    - I need 29 IP addresses in a network, which network size do I use?
    - Don't use /27 (32 IPs) as 5 of these will be reserved and you can't use them. You would need /26

### 202. Internet Gateways & Route tables

#### Hands on
- Enabled auto assing public IPv4 on public subnet A
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
- Create IGW, attach to VPC
- Created two routing tables
    - PublicRouteTable, PrivateRouteTable
- Associated the Private and Public route tables to private and public subnets
- Edited routes on the public route table and added an Internet Gateway (IGW, the one you created before)

### 203. NAT Instances
- But our instances in our private subnet would have no internet connection! :(
- We can use NAT for this. NAT in AWS comes in two flavours
    - NAT instances (outdated but still on the exam) 
    - NAT gateways (which we'll see in the next lecture)

- About NAT instances
    - NAT instancess allow instances in private subnets to connect to the Internet
    - They must be launched in a public subnet
    - Must disable EC2 flag: Source/Destination check
    - Must have an Elastic IP attached to it
    - Route table must be configured to route traffic from private subnets to NAT instance

--- STOPPED HERE --- as the hands on was about to start. It was kinda boring 
    
### 204. NAT Gateways

### 205. DNS Resolution options

### 206. NACL & Security groups

### 207. VPC Peering

### 208. VPC Endpoints

### 209. VPC Flow logs + Athena

### 210. VPC Flow Logs troubleshooting for NACL and SG

### 211. Bastion Hosts

### 212. Site to site VPN, Virtual Private Gateway & Customer Gateway

### 213. Direct Connect & Direct Connect Gateway

### 214. Egress only Internet Gateway

### 215. 

### Quiz
