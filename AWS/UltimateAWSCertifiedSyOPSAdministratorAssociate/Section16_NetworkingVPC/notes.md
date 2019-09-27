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
- Each VPC has a few subnets, it appesrs the default VPC has one subnet per AZ in ireland
- Subnets have gateways, ACLs, flow logs and some other features

### 200. VPC Overview and Hands on
- Virual Private Cloud (VPC)
- You can have multiple VPCs in a region (max of 5, soft limit)
- Max CIRD per VPC is 5. For each cider:
    - Min size is /28 (16 IP Adresses)
    - Max size is /16 (65536 IP Adresses)
- Because a VPC is private, only the private IP ranges are allowed! 
- Your VPC CIDR should not overlap with your other networks!

#### Hands on
- Your VPCs > Create VPC
- VPCs by default come with a main route table and a main network ACL
- You can have multiple CIDRs on a single VPC. Cool!

### 201. Subnet overview and Hands on

### 202. Internet Gateways & Route tables

### 203. NAT Instances

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
