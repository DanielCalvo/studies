### 152. Section intro
- As a SysOps in real life, these are useful

### 153. AWS Health dashboards
- Shows all regions, all services health
- Shows historical information for each day (you can go back in time and see if a service was reliable or not)
- Has an RSS feed that you can subscribe to
- `https://status.aws.amazon.com`
- A global status for all Amazon services on all regions. Cool!

#### Personal health dashboard!
- Global service
- Shows how AWS outages directly impact you
- Shows impact on your resources
- List issues and actions you can take to remediate them
 - Possible exam question: Which dashboard is right for you you? Which one will give you information about all services in all regions, and which one dashboard can show you the issues that impact you directly?
 - `https://phd.aws.amazon.com`

### 154. AWS Organizations Overview 
- Global service
- Allows you to manage multiple AWS accounts from one root account
- The main account is the master account - you can't change it
- All other accounts are member accounts
- One member can only be part of one organization
- Consolidated billing across all accounts - single payment method
- You get pricing benefits from aggregated usage (volume discount for EC2, S3)
- There's an API available to automate AWS account creation!

#### OU & Service Control Policies (SCPs)
- Organize accounts in a Organizational Unit (OU)
    - OU can be anything: dev/test/prod or finance/HR/IT 
    - You can nest OUs within OUs
- You can apply Service Control Policies to OUs
    - These permit or deny access to AWS services
    - SCP has a similar syntax to IAM
    - It's a filter to IAM
- Very helpful for sandbox account creation (possible exam question)
- Helpful to separate dev and prod resources
- Helpful to only allow approved services

### 155. AWS Organizations Hands-on

### 156. AWS Service Catalog Overview

### 157. AWS Service Catalog Hands-on

### 158. AWS Billing Alarms

### 159. AWS Cost Exporter

### 160. AWS Budgets

### 161. AWS Cost Allocation tags

### Quiz