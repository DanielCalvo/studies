## 84. Introduction to AWS Site-to-Site VPN
- Allows you communicate over the internet, in an encrypted form!
- AWS VPN supports layer 3 VPN (not 2)
- Site-to-site vpn connects two different networks
- Client to site vpn connects something like a laptop to a private network
- IPSec is the only VPN supported by AWS (GRE and DMVPN are not supported)

### How site-to-site works in AWS
- On the AWS site, site-to-site connection terminates on a virtual private gateway (VGW)
- On the customer site, you need something called "Customer gateway"
- AWS creates 2 tunnels for high-availability on the VPN connection 

### Virtual private gateway (VGW)
- Managed gateway endpoint for the VPC
- Only one VGW can be attached to a VPC at at time! :o
- VGW supports both static routing and dynamic routing using BGP
- Something something some bgp settings
- VGW supports AWS-256 and SHA-2 for encryption and data integrity