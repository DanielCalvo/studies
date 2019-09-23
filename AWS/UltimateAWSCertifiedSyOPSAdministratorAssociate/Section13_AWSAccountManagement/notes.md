- Chapter status: Took notes attentively, got the quiz very right!

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
- You invite/create accounts
- You then associate an account with an organizational unit
- You can create an access policy that can allow/deny access to certain services
- You then enable policy types and apply a policy(ies) to an OU
- Not only you have to check IAM for access denied errors, but if your account is part of an organization, there's a chance there's a service control policy attached to it
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
- Graphical tool to view and analyze your costs and service usage
- You can review charges and costs associated with your AWS account or Organization
- It can forecast spending for the next 3 months!
- Can give recommendation for which EC2 reserved instances to purchase
- Access to default reports
- There's an API to build custom cost management applications

#### Hands on cost explorer
- Services > Billing > Cost Explorer
- Shows month to date cost, and a forecast month end costs
- You can see graphs, per day!
- "Reservations" on the Cost explorer is popular on the exam!
- "Which service allows you to see the reservation summary and also to get recommendation on reserved instances? Protip: Cost exporter"
- You can also get reserved instances recommendations for RDS, ElastiCache, Redshift and Elasticsearch. 
- You can also see expenses per service and see what is costing you money. You can also filter by service!
- You can set a bunch of filters and save a report! 

### 160. AWS Budgets
- Similar to cost explorer!
- Create a budget and send alarms when costs exceed the budget
- 3 types of budgets: Usage, Cost, Reservation
- For reserved instances (RI)
    - Track utilization
    - Supports EC2, ElastiCache, RDS, Redshift
- Up to 5 SNS notifications per budget
- You can filter by Service, Linked account, tatg, purchase option, instance type, region, AZ, API operation...
- Same options as AWS Cost explorer!
- The first 2 budgets are free, then it's 0.02/day/budget

#### Hands on budgets
- The graphs in the Budget part come from cost explorer
- You have a lot of options on the budget section, you can also create a budget for t2.micro instances on Ireland for instance
- You can create an alert on actual costs or forecasted costs
- You can create an alert when you are at x% of your budgeted ammount (like 80%)
- You can also have a usage budget! Like how much % of 100gb of traffic you used on S3 data transfer. Neat!
    - Usage budgets are not about money/costs. They're about a metric and you can budget against that metric
- You can also track your reserved instance budget through this tool, though there was no example :(

### 161. AWS Cost Allocation tags
- We can use tags to track resources that relate to one another
- For tags to show up in a cost report, we have to define these tags as Cost Allocation Tags
- Just like tags, but they show up as columns in reports
- There are two types of cost allocation tags
- AWS Generated Cost Allocation Tags
    - Starts with prefix `aws:` (ex: aws: createdBy)
    - They're not applied to resources created before the activation (?)
- User tags
    - Defined by the user
    - Starts with the prefix `user:`
- Cost Allocation Tags just appear in the billing console
- It can take up to 24 hours for the tags to show up in the report

#### Hands on
- Billing > Cost Allocation Tags
- Remember that AWS-Generated Cost Allocation Tags are not seen on the resource, only on the Billing > Cost Allocation tag console!
- You can have any tag on your service become a cost allocation tag (ex: environment, team, app, etc)
- Possible question exam: "How can we separate our cost by tags or environments?". Use cost allocation tags!

### Quiz