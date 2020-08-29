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
- Configuration for providers (plugins to interact with remote systems, ex: AWS) is very straight forward
- You can define provider aliases for things like multi-region terraform configs

#### 023. Purpose of Terraform State documentation
- https://www.terraform.io/docs/state/purpose.html
- To map it's configuration files to resources, terraform uses states
- The state also allows terraform to track metadata such as resource dependencies
- The state can also represent a cached state of your infrastructure
- When using terraform as part of a team, it is recommended to use remote state

#### 024. Terraform Settings documentation
- https://www.terraform.io/docs/configuration/terraform.html
- This talks about the `terraform {}` configuration block
- You can configure a backend, specify a terraform version, provider requirements and experimental language features

#### 025. Provisioners documentation
- https://www.terraform.io/docs/provisioners/
- Hashicorp recommends using provisioners only as a last resort -- investigate other solutions first (like Packer!)
- You can use the `self` object to reffer to an atribute of the current object, like `self.private_ip`
- You can have destroy time provisioners
- You can adjust adjust the failure behaviour to `continue` or `fail` when a provisioner does not succeed

### 03. Master the workflow

#### 031. The Core Terraform Workflow documentationç
- https://www.terraform.io/guides/core-workflow.html

##### Single player
1. Write - Author infrastructure as code.
    - Write code
    - Use `terraform plan` to see if it's correct
2. Plan - Preview changes before applying.
    - After planning, the web page suggests you commit your changes
    - This will show you a plan, again!
3. Apply - Provision reproducible infrastructure.
    - After further checking, you're ready to provision the real infra!
    - Push your changes to git!

##### Working as a team
1. Write
    - Create a branch
    - Do your changes and `terraform plan` to validate them
    - It can be difficult to run the plan state if you have many sensitive variables (API keys, ssl certs and so on)
    - Guide recommends a CI environment for terraform, though it's outside of the scope for this section: https://learn.hashicorp.com/terraform/development/running-terraform-in-automation 
2. Plan
    - Either put the output of you plan on a PR or host it on your CI system
    - Review the PR! Will it, for instance, cause service disruption? Anything we should watch for? Anyone to notify?
3. Apply
    - See if other changes are also included once you merge this into master
    - Give it a go!

##### Working with Terraform Cloud
- You can use terraform cloud to securely store secrets
- TF cloud can create a plan for you when a pull request happens, for instance
- The plan can be pending approval for when you want to run it
- You can even have a discussion on the plan on the TF cloud UI, neat

#### 032. Initialize a Terraform working directory with init documentation
- Initializes a directory for use with terraform
- Plugins and modules are retrieved
- Generally simple & safe to run

#### 033. Validate a Terraform configuration with validate documentation
- Validates local configuration files (does not access remote systems, like the plan command)
- `terraform plan` includes an implied validation check

#### 033. Generate and review an execution plan for Terraform with plan documentation
- Creates an execution plan
- Can also generate a `destroy` plan
- Be careful: Plan can potentially output secrets

#### 034. Execute changes to infrastructure with Terraform with apply documentation
- Applies the changes to reach the desired state

#### 035. Destroy Terraform managed infrastructure with destroy documentation
- Used to apply the changes to reach the desired state of the configuration
    - By default it'll scan the current directory and apply the configurations it finds there
    - Can use an plan file to apply changes too
    - Has a bunch of flags, all optional, nothing too crazy


### 04. Learn more subcommands
#### 041. Formatting configuration with fmt documentation
- https://www.terraform.io/docs/commands/fmt.html
- Formats terraform configuration files to a canonical format and style

#### 042. Tainting resources with taint documentation
- https://www.terraform.io/docs/commands/taint.html
- `terraform taint` marks a terraform-managed resource as tainted, forcing it to be destroyed and recreated on the next apply
- This does not modify infra structure, but does modify the state file to mark the resource as tainted
- Ex: `terraform taint aws_instance.foo`
- You can also taint a resource within a module, but the syntax looks more involved

#### 043. Managing state with state documentation
- https://www.terraform.io/docs/commands/state/index.html
- `terraform state <subcommand> [options] [args]` 
- `terraform state list`
- `terraform state mv aws_instance.example aws_instance.something` (but if you apply the state again, it will destroy renamed instance and create a new one)
- `terraform state pull`
- `terraform state rm aws_instance.example`
- `terraform state push mystatefile.tfstate`
- `terraform state show`

#### 044. Using local workspaces with workspace documentation
- https://www.terraform.io/docs/state/workspaces.html

#### 045. Importing existing resources with import documentation
- https://www.terraform.io/docs/commands/import.html

### 05. Use and create modules
#### 051. Organize Configuration with Modules (complete all tutorials)
- https://learn.hashicorp.com/collections/terraform/modules
- Modules can help you navigate your configuration files, reduce repetition and risk, as well as share reusable terraform code

##### Modules Overview
- https://learn.hashicorp.com/tutorials/terraform/module
- Modules are for:
    - Organize configuration: By using modules you can organize your configuration into logical components
    - Encapsulate configuration: Allows you to organize your config into distinct logical components
    - Re-use configuration: Using modules allows you to easily re-use terraform code written by yourself or others
    - Provide consistency & ensure best practices: Modules can help you provide consistency in your configurations
- What's a terraform module?
    - A set of tf configuration files in a single directory
- Local and remote modules
    - Modules can be loaded from the local filesystem or a remote source, such as HTTP, git or Terraform Cloud or Terraform Enterprise Module Registries
- Module best practices
    - The concept of terraform modules are similar to the concepts of libraries, packages or modules found in most programming languages
    - Start writing your config with modules in mind
    - Use local modules to organize and encapsulate your code
    - Use the public Terraform Registry to find useful modules
    - Publish and share modules with your team

##### Use Modules From the Registry
- https://learn.hashicorp.com/tutorials/terraform/module-use
- 

##### Build a Module
- https://learn.hashicorp.com/tutorials/terraform/module-create
- 

##### Share Modules in the Private Module Registry
- https://learn.hashicorp.com/tutorials/terraform/module-private-registry
- 

##### Separate Development and Production Environments
- https://learn.hashicorp.com/tutorials/terraform/organize-configuration
- 

#### 052. Finding and using modules documentation
- https://www.terraform.io/docs/registry/modules/use.html

#### 053. Module versioning documentation
- https://www.terraform.io/docs/configuration/modules.html#module-versions

#### 054. Input Variables documentation
- https://www.terraform.io/docs/configuration/variables.html

#### 055. Calling a child module documentation
- https://www.terraform.io/docs/configuration/modules.html#calling-a-child-module

### 06. Read and write configuration
#### 061. Resources describe infrastructure objects
- https://www.terraform.io/docs/configuration/resources.html

#### 062. Data sources let Terraform fetch and compute data
- https://www.terraform.io/docs/configuration/data-sources.html

#### 063. Resource addressing lets you refer to specific resources
- https://www.terraform.io/docs/internals/resource-addressing.html

#### 064. Named values let you reference values
- https://www.terraform.io/docs/configuration/expressions.html#references-to-named-values

#### 065. Complex types let you validate user-provided values
- https://www.terraform.io/docs/configuration/types.html#complex-types

#### 066. Built in functions help transform and combine values
- https://www.terraform.io/docs/configuration/functions.html

#### 067. Dynamic blocks allow you to construct nested expressions within certain configuration blocks
- https://www.terraform.io/docs/configuration/expressions.html#dynamic-blocks

#### 068. Vault Provider for Terraform
- https://www.terraform.io/docs/providers/vault/index.html


### 07. Manage state
#### 071. State locking documentation
- https://www.terraform.io/docs/state/locking.html

#### 072. Sensitive data in state documentation
- https://www.terraform.io/docs/state/sensitive-data.html

#### 073. Reconcile state and real resources with refresh documentation
- https://www.terraform.io/docs/commands/refresh.html

#### 074. Backends overview documentation
- https://www.terraform.io/docs/backends/index.html

#### 075. Local backend documentation
- https://www.terraform.io/docs/backends/types/local.html

#### 076. Backend types documentation
- https://www.terraform.io/docs/backends/types/index.html

#### 077. How to configure a backend documentation
- https://www.terraform.io/docs/backends/config.html


### 08. Debug in Terraform
#### 081. Review the debugging documentation
- https://www.terraform.io/docs/internals/debugging.html


### 09. Understand Terraform Cloud and Enterprise
#### 091.Terraform Cloud overview documentation
- https://www.terraform.io/docs/cloud/index.html

#### 092.Understanding Workspaces and Modules resource
- https://www.hashicorp.com/resources/terraform-enterprise-understanding-workspaces-and-modules/

#### 093.CLI workspaces documentation
- https://www.terraform.io/docs/state/workspaces.html

#### 094.Terraform Cloud workspaces documentation documentation
- https://www.terraform.io/docs/cloud/workspaces/index.html

#### 095.Module registry blog post
- https://www.hashicorp.com/blog/hashicorp-terraform-module-registry/

#### 096.Module registry documentation
- https://www.terraform.io/docs/cloud/registry/index.html

#### 097.Why Policy as Code? blog post
- https://www.hashicorp.com/blog/why-policy-as-code/

#### 098.Sentinel Policy as Code documentation
- https://www.terraform.io/docs/cloud/sentinel/index.html

#### 099.Feature comparison pricing page (scroll down for feature matrix)
- https://www.hashicorp.com/products/terraform/pricing/