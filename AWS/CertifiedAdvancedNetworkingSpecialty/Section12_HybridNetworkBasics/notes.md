## 80. Introduction to Hybrid networking
Most important services
- Site to site VPN
- AWS Direct connect

To connect on prem to AWS
- Use AWS site-to-site VPN
- Use a virtual private gateway attached to the VPC (flows over theinternet with ipsec)
- Resources in both networks will then be able to acess each other using private IPs

Client to site VPN
- Client, working from home, access the AWS network (use VPC resources privately)

Both VPNs are secure but not reliable (bandwitdh issues, jitter, etc)
To have a consistent network between the datacenter and AWS, and for that, you can use AWS direct connect
With direct connect you have a physical fiber connection between your datacenter and AWS!

To understand Direct connect, you need to understand static and dynamic routing works