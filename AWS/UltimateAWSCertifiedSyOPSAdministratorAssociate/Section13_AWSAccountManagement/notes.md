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
- You invite create accounts
- You then associate an account with an organizational unit
- You can create an access policy that can allow/deny access to certain services
- You then enable policy types and apply a policy(ies) to an OU
- Not only you have to check IAM for acess denied errors, but if your account is part of an organization, there's a chance there's a service control policy attached to it
- You can only see from the root account which policies are assigned to which accounts

### 156. AWS Service Catalog Overview
- Not very popular on the exam, but important to know
- Users that are new to AWS have too many options and may create stacks that are not compliant or in line with the rest of the organization
- Some users just want a quick self-service portal to launch a set of authorized products pre-definer by admins
- As an admin, you create a product, which is a cloudformation template
- And then you create a portfolio, which is a collection of products 
- As a user, you're presented with a product list

#### AWS service catalog, more info
- Creates and manages catalogs of IT services that are approved on AWS
- The "products" are CloudFormation templates (VMs, servers, software, databases, regions, ip ranges, you can do a bunch of stuff)
- CloudFormation helps ensure consistency and standardization by Admins
- Products are assigned to portfolios. Teams are presented a self-service portal where they can launch the products
- All the deployed products are centrally managed deployed services
- Helps with governance, compliance and consistency
- Can give user access to launching products without requiring deep AWS knowledge
- Integrations with "self-service portals" such as serviceNow 

### 157. AWS Service Catalog Hands-on
- Created a stack from a cloudformation config
- Created a portfolio
- Added a product to a portfolio
- You add users/groups to a portfolio

### 158. AWS Billing Alarms
- Billing data metric is stored in CLoudWatch us-east-1
- Billing data is for overall WorldWide AWS costs
- This is for global costs, not for project costs
- Let's create a billing alarm 
-  Top right, your account name > My Billing Dashboard > Billing preferences > Receive Billing Alets > Manage Billing Alerts
- On the left pannel: Billing > Alarms

### 159. AWS Cost Exporter

### 160. AWS Budgets

### 161. AWS Cost Allocation tags

### Quiz