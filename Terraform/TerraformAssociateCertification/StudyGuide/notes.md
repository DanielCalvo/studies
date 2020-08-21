Summary of study topics as described here: https://learn.hashicorp.com/tutorials/terraform/associate-study

### 01. Learn about IaC
#### 011. Infrastructure as Code introduction video
- https://www.youtube.com/watch?v=RO7VcUAsf-I
- Automate, script and version your infrastructure
- High level overview, pretty cool

#### 012. Infrastructure as Code in a Private or Public Cloud blog post
- https://www.hashicorp.com/blog/infrastructure-as-code-in-a-private-or-public-cloud/
- High level overview, pretty cool

#### 013. Introduction to IaC Documentation
- https://www.terraform.io/intro/index.html
- You can use Terraform for a single app of your entire datacenter Infra
- Infrastructure as a code: Infrastructure described with a high-level syntax. Allows for a blue print of your infra to be versioned and treated as code. Infrastructure can be shared and re-used
- Execution plans: Terraform has a planning step which shows you which actions it will take. This helps avoid surprises
- Resource graph: Terraform can build a graph of all your resources, and parallelize creation and modification of non-dependent resources
- Change Automation: Using planning and the resource graph feature, you should know what terraform will change and in what order

#### 014. Terraform Use Cases Documentation
- Heroku App Setup:
- Multi-Tier Applications: Can be used to describe and scale multi-tiers independently, and make sure certain tiers are ready before others (ex: db before API)
- Self Service Clusters: Allows you to delegate infrastructure to product teams, completely or as a consultant, as the knowledge on how to build a service is codified in Terraform configs
- Software Demos: Allows you to bundle the infrastructure together with your software, so users can demo it. Noice!
- Disposable Environments:  Environments can be codified and parametrized, so they can be easily spin up and disposed of
- Software Defined Networking: You can have your network as code, such as in Amazon VPCs, ACLs and SGs
- Resource Schedulers: Terraform can interact with resource providers such as Hadoop and K8s. So you can set up the underlying infrastructure of k8s, and k8s resources with Terraform. Neat!
- Multi-Cloud Deployment: Allows for multi-cloud or multi-region deployments. Terraform is cloud agnostic.

### 02. Manage infrastructure
#### 021. Get Started with Terraform (complete all tutorials)
- https://learn.hashicorp.com/collections/terraform/aws-get-started
- [./021_GetStartedAWS/02_Install](./021_GetStartedAWS/02_Install) has a docker terraform example that is pretty cool

##### [./021_GetStartedAWS/03_BuildInfra](./021_GetStartedAWS/03_BuildInfra)
- Terraform block: New for terraform 0.13, so that terraform knows which provider to download from the registry. You can also lock a provider to given version
- Providers: Configures a given provider. A provider is a plugin Terraform uses to have API interactions with a service (ex: AWS). Terraform can interact with any API, so you can represent almost any infra as a resource in TF.
- Resources: The resource block defines a piece of infrastructure.
- Initialize: Installs providers (plugins) used in the TF files
- Format: Formats files to follow syntax conventions. `terraform validate` will also tell you if your config is valid.
- Create Infra: `terraform apply` will show an execution plan and apply changes to infrastructure
- Inspect state: `terraform show` will show you the current state
- Manually manage state: `terraform state list`. You can also manually manage state with the other commands in there.
- Troubleshooting: Some AWS specific troubleshooting

##### Change Infra
- If you change the AMI ID of an instance, terraform will destroy the instance and recreate it

##### Destroy Infra
- `terraform destroy` will destroy your infrastructure

##### Provision Infra
- Provisioners let you upload files, run shell scripts, install software and trigger other other software
- [./021_GetStartedAWS/06_ProvisionInfra](./021_GetStartedAWS/06_ProvisionInfra) has both local-exec and remote-exec provisioners examples, check them out!
- Provisioners only run when a resource is created
- You can have multiple provisioner steps
- You can also have provisioners that run during a destroy operation
- Terraform recommends using provisioners only as a last resort. Consider alternatives: https://www.terraform.io/docs/provisioners/index.html
- If a resource is succesfully created but provisioning fails, terraform marks the resource as tainted. Tainted resources are marked as unsafe to use. On your next execution plan, terraform will remove any tainted resources and create new ones, attempting to provision them again
- If you want to manually destroy and re-create a resource, you can manually `taint` it with `terraform taint resource.id`

##### Input Variables
- Variables allow us to parametrize terraform 
- Check [./021_GetStartedAWS/07_InputVariables](./021_GetStartedAWS/07_InputVariables) for an example. Note that you can name the var file anything, terraform loads all the .tf files in the directory it runs
- You can also: `terraform apply -var 'region=us-west-2`
- You can create a local file named `secrets.tfvars` and then save it
- One way to deploy applications with multiple environments would be like this: `terraform apply -var-file="secret.tfvars" -var-file="production.tfvars"` 
- You can assign variables from a file, from environment variables (Terraform will read environment variables in the form of TF_VAR_name), with UI input if you leave variables unspecified, or with defaults variables
- There are also rich variable types: Lists [], maps (aka key/value) {}.
- You can also pass maps through the command line: `terraform apply -var 'amis={ us-east-1 = "foo", us-west-2 = "bar" }'`

##### Output Variables
- Useful to output some values of your `terraform apply`
- Check [./021_GetStartedAWS/08_OutputVariables](./021_GetStartedAWS/08_OutputVariables) for an example

##### Store Remote State
- In a production environment it's considered a best practice to store state anywhere other than your local machine
- Terraform supports remote remote backends to remotely store state
- Created a workspace in app.terraform.io
- Generated a new token
- Saved token at `~/.terraformrc`
- You then need to export your AWS Credentials on app.terraform.io, as this will run your terraform state (remotely, on the cloud), as well as store the terraform backend remotely
- Alternatively, you can use S3 bucket for your S3 state. It is highly recommended to version this bucket
- Check [./021_GetStartedAWS/09_RemoteState](./021_GetStartedAWS/09_RemoteState)
- Neat!

#### 022. Providers documentation
- https://www.terraform.io/docs/configuration/providers.html
#### 023. Purpose of Terraform State documentation
#### 024. Terraform Settings documentation
#### 025. Provisioners documentation

### Master the workflow
#### The Core Terraform Workflow documentation
#### Initialize a Terraform working directory with init documentation
#### Validate a Terraform configuration with validate documentation
#### Generate and review an execution plan for Terraform with plan documentation
#### Execute changes to infrastructure with Terraform with apply documentation
#### Destroy Terraform managed infrastructure with destroy documentation

### Learn more subcommands
#### Formatting configuration with fmt documentation
#### Tainting resources with taint documentation
#### Managing state with state documentation
#### Using local workspaces with workspace documentation
#### Importing existing resources with import documentation

### Use and create modules
#### Organize Configuration with Modules (complete all tutorials)
#### Finding and using modules documentation
#### Module versioning documentation
#### Input Variables documentation
#### Calling a child module documentation

### Read and write configuration
#### Resources describe infrastructure objects
#### Data sources let Terraform fetch and compute data
#### Resource addressing lets you refer to specific resources
#### Named values let you reference values
#### Complex types let you validate user-provided values
#### Built in functions help transform and combine values
#### Dynamic blocks allow you to construct nested expressions within certain configuration blocks
#### Vault Provider for Terraform

### Manage state
#### State locking documentation
#### Sensitive data in state documentation
#### Reconcile state and real resources with refresh documentation
#### Backends overview documentation
#### Local backend documentation
#### Backend types documentation
#### How to configure a backend documentation

### Debug in Terraform
#### Review the debugging documentation

### Understand Terraform Cloud and Enterprise
#### Terraform Cloud overview documentation
#### Understanding Workspaces and Modules resource
#### CLI workspaces documentation
#### Terraform Cloud workspaces documentation documentation
#### Module registry blog post
#### Module registry documentation
#### Why Policy as Code? blog post
#### Sentinel Policy as Code documentation
#### Feature comparison pricing page (scroll down for feature matrix)