## 48. Introduction to VPC endpoints
Provides connectivity between your VPC and other AWS services!

### Described scenario
- Normally, something in a private subnet accesses AWS services through the internet, going through a nat gateway and an internet gateway
- Using VPC endpoints you can reach AWS services through a private network

### VPC Endpoints
- Endpoints allow you to connect to AWS services without using the Internet
- Removes the need for IGW and NAT GW to access AWS Services
- Endpoint devices scale horizontally, are redundant, etc

- **Gateway endpoint**: Providers a target and must be used in a route table to access S3 and DynamoDB.
- **Interface Endpoint**: Provides an ENI (private IP) as an entry point -- for most other AWS services