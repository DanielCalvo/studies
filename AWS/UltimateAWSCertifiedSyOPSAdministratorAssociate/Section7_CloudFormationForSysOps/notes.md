
### 54. CloudFormation intro
- Infrastructure as a code :D
- Important exam topic!

### 55. CloudFormation Overview
- Currently we've been doing a lot of manual work
- Manual work is tough to reproduce! Imagine if you had to do it in another region, or even in another AWS account!
- Or imagine something was deleted, it would be tough to recreate it
- Wouldn't it be great if all our infrastructure was... code? :D
- The code that we write will be able to create/update delete our infrastructure

#### What is CloudFormation
- CloudFormation is a declarative way of outlining your AWS infrastructure, for any resources (most of them are supported)
- In a CloudFormation template, you could specify
    - I want a security group
    - I want two EC2 machines using this Security group
    - I want two elastic IPs for these EC2 machines
    - I want a S3 bucket
    - I want a load balancer (ELB) in front of these machines
- Cloudformation creates those for you, in the right order, with the exact configuration that you specify 

#### Benefits of CloudFormation (1/2)
- Infrastructure as Code
    - No resources will be manually created, which is excellent for control
    - The code can be versioned (using git for example)
    - Changes to the infrastructure are reviewed as code
- Cost
    - Each resource within a stack is stagged with an identified so you can see how much a stack costs you
    - You can estimate the costs of your resources using the CloudFormation template
    - Savings strategy: In dev, you could automate deletion of templates at 5PM and recreate them at 8am safely. (that's so cool)

#### Benefits of CloudFormation (2/2)
- Productivity
    - Ability to destroy and re-create infrastructure on the cloud on the fly 
    - Automated generation of diagrams for your templates
    - Declarative programming (no need to figure out ordering and orchestration)
- Separation of concern: Create many stacks for many apps and many layers, ex:
    - VPC stacks
    - Network stacks
    - App stacks
- Don't reinvent the wheel
    - Leverage existing templates on the web
    - Leverate the documentation

#### How CloudFormation works
- Templates have to be uploaded in S3 and then referenced in CloudFormation
- TO update a template, you can't edit the previous ones. You have to re-upload a new version of the template to AWS
- Stacks are identified by name
- Deleting a stack deletes every single artifact that was created by CloudFormation

#### Deploying CloudFormation templates
- Manual way
    - Editing templates in the CLoudFormation Designer
    - Using the console to input parameters
Automated way
    - Editing templates in a YAML file
    - Using the AWS CLI to deploy the templates
    - Recommended way when you fully want to automate your flow
    
#### CloudFormation Building blocks
- Template components
1. Resources: Your AWS resources declated in the template (Mandatory, you need to have a resource)
2. Parameters: The dynamic input for your templates
3. Mappings: The static variables for your template
4. Outputs: References to what has been created
5. Conditionals: List of conditions to perform resource creation 
6. Metadata

- Template helpers
1. References
2. Functions

### Note on introduction to CloudFormation
- It can take over 3 hours to properly learn and master CloudFormation
- This section is meant to get you a good idea of how it works
- We'll be slightly less hands-on than in other sections
- We'll learn anything we need to answer questions for the exam
- The exam does not require you to write CloudFormation (wew). It will ask about what feature should you use in CloudFormation to perform something
- Exam expects you to understand how to read CloudFormation

### 56. CloudFormation Create Stack Hands on
- We're going to create a simple EC2 instance
- Then we're going to create and add an Elastic IP to it
- And we're going to add two security groups to it!
- For now, forget about the syntax, we'll at that later on, this is just to get started
- Uh-oh, CloudFormation designer is very fancy 

#### Hands on
- Go to North Virginia, Cloud Formation, Create a stack. We're going to upload this template:

```yaml
Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us-east-1a
      ImageId: ami-009d6802948d06e52
      InstanceType: t2.micro
```

### 57. CloudFormation Update and Delete stack
### 58. YAML Crash Course
### 59. CloudFormation Parameters
### 60. CloudFormation Resources

### 60. CloudFormation Resources

### 61. CloudFormation Mappings

### 62. CloudFormation Outputs

### 63. CloudFormation Conditions

### 64. CloudFormation Instrinsic Functions
### 65. ==== SysOps Specific Lectures ====

### 66. CloudFormation User Data
### 67. CloudFormation cfn-init 
### 68. CloufFormation cfn-signal and wait conditions


### 69. CloudFormation cfn-signal failures troubleshooting

### 70. CloudFormation Rollbacks

### 71. CloudFormation Nested STacks
### 72. CloudFormaton ChangeSets
### 73. CloudFormation DeletionPolicy
### 74. CloudFormation TerminationProtection

#### Quiz