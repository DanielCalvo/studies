- Chapter status: This seems fun! I want to practice! I want to launch things!
- Lecture 64: Hey I'd like a find in map example
- Fn::Sub was not very well explained on lecture 64

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
- Go to North Virginia, Cloud Formation, Create a stack. We're going to upload this template in `0-just-ec2.yaml`

- Neat, it created an instance! With some tags indicating it was created by cloud formation
- Most important tabs on cloudformation: Template

### 57. CloudFormation Update and Delete stack
- Let's update the template with the code on `1-ec2-with-sh-eip.yaml`
- Uuuuh it has parameters!
- You can preview your changes! :o
- Replacement: True < The old EC2 instance will be deleted and a new one will be created
- Deleting the stack deletes all resources. In the right order! Neat!

### 58. YAML Crash Course
- CloudFormation supports YAML and JSON 
- Author says: JSON is horrible for CloudFormation :joy:, it's hard to read, it's hard to write

### 59. CloudFormation Parameters
- Parameters are a way to provide inputs to your AWS CloudFormation template
- They're support important to know about if:
    - You want to reuse your templates across the company
    - Some inputs cannot be determined ahead of time
- Parameters are extremely powerful, controlled, and can prevent errors from happening in your templates thanks to types

Example parameter:
```yaml
Parameters:
  SecurityGroupDescription:
    Description: Security Group Description
    Type: String
```

- When to user a parameter, ask yourself this:
    - Is this CloudFormation resource likely to change in the future? If so, make it a parameter
    - If using a parameter, you won't have to re-upload a template to change its content

#### Parameters Settings
- Parametesr can be controlled by all these settings

- Type
    - String
    - Number
    - CommaDelimitedList
    - List<Type>
    - AWS Parameter (to hlep catch invalid values - match against existing values in the AWS Account)
- Description
- Constraints
- ConstraintDescription (string)
- Min/MaxLenght
- Min/MaxValue
- Defaults
- AllowedValues (array)
- AllowedPatterm (regexp)
- NoEcho (Boolean)

- Exam probably only covers strings as parameters

#### How to Reference a Parameter
- The Fn::Ref function can be leveraged to reference parameters
- Parameters can be used anywhere in a template
- The shorthand for a reference in YAML is !Ref (!Ref == Fn::Ref)
- The function can also reference other elements within the template

```yaml
DbSubnet1:
  Type: AWS::EC2::Subnet
  Properties:
    VpcId: !Ref MyVPC
```

- The !Ref function can be used to reference paramenters or another resource

#### Concept: Pseudo Parameters
- AWS offers us pseudo parameters in any CloudFormation template
- These can be used at any time and are enabled by default
- More like AWS Built in variables, like AWS::AccountId and AWS::Region, AWS::StackName and AWS::NotificationARNs

### 60. CloudFormation Resources
- Resources are the core of your CloudFormation template (they're mandatory!)
- They represent the different AWS components that will be created and configured
- Resources are declared and can reference each other
- AWS figures out creation, updates and deletes of resources for us
- There are over 224 types of resources, woah
- Resource types identifiers have the following form:
    - AWS::aws-product-name::data-type-name

#### How to find resources documentation
- `https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html`
- Read the docs! 

#### FAQ for resources
- Can I create a dynamic amount of resources?
    - No you can't. Everything in the CloudFormation template has to be declared. You can't perform code generation there.
- Is every AWS service supported?
    - Almost. A few niche things are not there yet
    - You can work around that using AWS Lambda Custom Resources

### 61. CloudFormation Mappings
- Mappings are fixed variables within your CloudFormation template
- They're very handy to differentiate between different environments (dev vs prod) regions, AMI types and other things
- All values are hardcoded within the template

#### When would you use mappings vs parameters?
- Mappings are great when you know in advance all the values that can be taken and that they can be deduced from variables
    - Region
    - AZ
    - AWS Account
    - Environment
- They allow safer control over the template
- Use parameters when the values are really user specific

#### Fn::FindInMap - Accessing Mapping Values
- You use Fn::FindInMap to return a named value from a specific key
- `!FindInMap [MapName, TopLevelKey, SecondLevelKey]` (syntax is important for the exam)

### 62. CloudFormation Outputs
- The output section declares optional output values that you can import into other stacks (if you export them first) Popular exam question!
- You can also view the outputs in the AWS console or using the AWS CLI
- They're very useful for example if you define a network CloudFormation, and output the variables such as VPC ID and your Subnet IDs 
- It's the best way to perform some types of collaboration, as you can let different people handle different parts of the stack 
- You can't delete a CloudFormation stack if its outputs are being referenced by another CloudFormation stack

#### Outputs Example
- Creating a SSH Security Group as part of one template
- We create an output that references that security group
```yaml
Outputs:
  StackSSHSecurityGroup:
    Description: The SSH Security Group for our company
    Value: !Ref MyCompanyWideSSHSecurityGroup
    Export:
      Name: SSHSecurityGroup
```

##### Cross stack reference
- We then create a second template that leverages that security group
- For this, you use the `Fn::ImportValue` function
- You can't delete the underlying stack untill all the references are deleted too

```yaml
Resources:
  MySecureInstance:
    Type: AWS::EC2::Instance
    Properties:
      SecurityGroups:
        - !ImportValue SSHSecurityGroup

```
- Outputs and exports are popular on the exam, with questions like "How do you link CloudFormation templates?" or "pass the value from one template to another"

### 63. CloudFormation Conditions
- Conditions are used to control the creation of resources or outputs based on a condition
- Conditions can be whatever you want, but common ones are
    - Environment (dev/test/prod)
    - AWS Region
    - Any parameter value
- Each condition can reference another condition, parameter value or mapping

#### How to define a condition
```yaml
Conditions:
  CreateProdResources: !Equals [ !Ref EnvType, prod ]

```
- On the above: The condition is true if EnvType is prod
- The logical ID is for you to choose. It's how you name the condition
- The intrinsic function (logical) can be any of the following:
    - Fn::And
    - Fn::Equals
    - Fn::If
    - Fn::Not
    - Fn::Or

#### Using a condition
- Conditions can be applied to resources/outputs/etc
```yaml
Resources:
  MountPoint:
    Type: "AWS::EC2::VolumeAttachment"
    Condition: CreateProdResources
```
- The resource above will only get created if the `CreateProdResources` (defined separately) is true 
- Not sure if asked at the exam, but still good to know!

### 64. CloudFormation Instrinsic Functions
- List of functions you must absolutely know for the exam!
- Ref
- Fn::GetAtt
- Fn::FindInMap
- Fn::ImportValue
- Fn::Join
- Fn::Sub
- Condition functions (Fn::if, Fn::Not, Fn::Equals, etc, see above)

#### Fn::Ref
- The Fn::Ref function can be leveraged to reference
    - Paramaters > returns the value of the parameter
    - Resources > returns the physical ID of the underlying resource (ex: EC2 ID)
- The shorthand for this in YAML is !Ref
```yaml
DbSubnet1:
  Type: AWS::EC2::Subnet
  Properties:
  VpcId: !Ref MyVPC
```

#### Fn::GetAtt
- Atributes are attachd to any resources you create
- To know the attributew of your resources, the best place to look is at the documentation
- `https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#aws-properties-ec2-instance-return-values`
- Popular exam question: "How do we get this attribute from this resource?"

```yaml
Resources:
  EC2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      ImageId: ami—1234567
      InstanceType: t2.micro
```
```yaml
NewVolume:
  Type: "AWS::EC2::Volume"
  Condition: CreateProdResources
  Properties:
    Size: 100
    AvailabilityZone:
      !GetAtt EC2Instance.AvailabilityZone   
```

#### Fn::FindInMap - Accessing Map Values
- You use Fn::FindInMap to return a named value from a specific key
- `!FindInMap [MapName, TopLevelKey, SecondLevelKey]` (syntax is important for the exam)

#### Fn::ImportValue
- To import values that are exported in other templates
- For this, we use the Fn::ImportValue function

```yaml
Resources:
  MySecureInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us—east—la
      ImageId: ami—a4c7edb2
      InstanceType: t2.micro
      SecurityGroups:
        - !ImportVa1ue SSHSecurityGroup
```

#### Fn::Join
- Join values with a delimiter
- `!Join [ delimiter, [ comma-delimited-list-of-values ] ]`
- This creates "a:b:c"
- `!Join [ ":", [ a, b, c ] ]`

#### Fn::Sub
- Stephane didn't go a good job on this one so I'll have to find my own explanation for this

#### Conditions:
- These were explained above and this explanation is just a copy and paste. See above!

### 65. ==== SysOps Specific Lectures ====
- The following lectures focus on topics that the exam brings up

### 66. CloudFormation User Data
- We can have user data at EC2 instance when launching one through the console
- We can also include this data in CloudFormation
- The important thing to pass is the entire script through the function Fn::Base64
- Good to know: user data script log is in /var/log/cloud-init-output.log
- Let's see how to do this in CloudFormation
- Look at `3-user-data.yaml`
- Remember to look at /var/log/cloud-init-output.log for debug information

### 67. CloudFormation cfn-init 
- User data can get a bit clunky when you have something a bit more elaborate
- AWS::CloudFormation::Init must be in the Metadata of a resource
- With the cfn-init script, the complex EC2 configurations are made more readable
- The EC2 instance will query the CloudFormation service to get the init data
- Logs go to /var/logs/cfn-init.log

#### Diagram
- CloudFormation service launches an EC2 instance
- As part of the user data script, cfn-init runs on the EC2 instance
- As part of this script, the instance will query the CloudFormation service to retreive init data, and then it will apply it

### 68. CloudFormation cfn-signal and wait conditions
- We still don't know how to tell CloudFormation that the EC2 instance got properly configured after a cfn-init 
- For this, we can use the cfn-signal script!
    - We'll run the cfn-signal right after cfn-init
    - Tell CloudFormation service to keep going or fail
- We need to define a WaitCondition
    - This will block the template until it receives a signal from cfn-signal
    - We attach a CreationPolicy (also works for EC2 and ASG)

#### Diagram
- CloudFormation service launches an EC2 instance
- As part of the user data script, cfn-init runs on the EC2 instance, CloudFormation will wait in the meantime
- As part of this script, the instance will query the CloudFormation service to retreive init data, and then it will apply it
- After it's done retrieving (and probably applying?) the init data, the EC2 instante will send a signal from cfn-signal to CloudFormation

### 69. CloudFormation cfn-signal failures troubleshooting

#### Wait condition didn't receive the required number of signals from an Amazon EC2 instance
- Important exam question!
- Ensure that the AMI you're using has the AWS CloudFormation helmet scripts installed. If the AMI doesn't include the helper scripts, you can also download them to you instance
- Verify that the cfn-init and cfn-signal commands were successfully run on the instance. You can view logs, such as /var/log/cloud-init.log or /var/log/cfn-init.log t to help you debug the instance launch
- You can retrieve the logs by logging in to your instance, but you must disable rollback on failure or else AWS CloudFormation deletes the instance after your stack fails to create
- Verify that the instance has a connection to the Internet. If the instance is in a VPC, the instance should be able to connect to the Internet through a NAT device if it is in a private subnet or through an Internet gateway if it's in a public subnet. If an instance can't talk to the internet, it can't talk to the CloudFormation service

- Instance shuts down in this scenario, no way to debug it :()

### 70. CloudFormation Rollbacks
- If something fails creating the stack (using the CreateStack API), everything rolls back
    - **Default**: Everything rolls back (gets deleted). But you can look at the log on the CloudFormation console!
    - OnFailure=ROLLBACK 
    - **Troubleshoot**: There's an option to disable rollback and manually troubleshoot
    - OnFailure = DO_NOTHING
    - **Delete**: Get rid of the stack entirely, do not keep anything
    - OnFailure=DELETE
- If something fails updating the stack (UpdateStack API)
    - The stack automatically rolls back to the previous known working state
    - You can only see in the log what happened and the error messages

### 71. CloudFormation Nested Stacks
- Can be asked about in the exam!
- Nested stacks are stacks as part of other stacks
- They allow you to isolate repeated patterns/common components in separate stacks and call them from other staacks
- Example
    - Load Balancer configuration that is re-used
    - Security group that is re-used
- Nested stacks are considered best practive
- To update a nested stack, you always update the parent stack first (root stack)

### 72. CloudFormaton ChangeSets
- When you update a stack, you need to know what changes before it happens for greater confidence
- A ChangeSet will tell you what happens if you apply a given update (it won't tell you if it will work 100% though)
- https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-changesets.html

### 73. CloudFormation DeletionPolicy
- You can put a DeletionPolicy on any resource to control what happens when the CloudFormation template is deleted
- DeletionPolicy=Retain
    - Specify on resources to preserve/backup in case of CloudFormation deletes
    - To keep a resource, specify Retain (works for any resource / nested stack)
- DeletionPolicy=Snapshot
    - EBS Volume, ElastiCache Cluster, ElastiCache Cluster, ElastiCache ReplicationGroup
    - RDS, Redshift
    - Instance gets deleted, but it makes a snapshot right before deleting it
- DeletionPolicy=Delete (default behaviour)
    - Note for AWS::RDS::DbCluster resources, the default policy is snapshot
    - To delete an S3 bucket, you need to first empty the bucket of its content

### 74. CloudFormation TerminationProtection
- To prevent accidental deletes of CloudFormation templates, use TerminationProtection
- Create a stack, edit termination protection, enable termination protection
- To delete it, you must delete the termination protection first

#### Quiz