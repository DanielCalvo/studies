## 44. Introduction to VPC private connectivity options
- If you want to communicate between two public subnets in two VPCs, you just communicate over the internet
- If its one private instante in one VPC and a public one in another VPC, the traffic goes from the private one through a NAT gateway and it reaches the public VPC through the internet
- Public to private in different VPCs: You can't. You can hack some port nat or port forwarding but directly it is not possible
- Private to private: Also not possible

### Disadvantages of having internet based traffic
- Less secure
- Internet may provide unreliable bandwidth and latency
- NAT devices have cost on AWS
- Unecessarily exposing services to the internet 

VPC Peering, VPC endpoints and VPC privatelink solve this problem!

## 45. VPC Peering
Simplest solution!

Connects 2 private subnets. VPC Peering connection through a private network owned by AWS.

- Makes them behave as if they're in the same network
- Can be in same or separate regions
- You can also do VPC peering with another AWS account!

### Caveates
- VPC CIRDS should not overlap!
- You must update route tables in both VPCs so that the traffic can go through the vpc peering connection!