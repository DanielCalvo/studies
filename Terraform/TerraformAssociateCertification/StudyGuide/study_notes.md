Summary of study topics as described here: https://learn.hashicorp.com/tutorials/terraform/associate-study

- [01. Learn about IaC](#01-learn-about-iac)
  * [011. Infrastructure as Code introduction video](#011-infrastructure-as-code-introduction-video)
  * [012. Infrastructure as Code in a Private or Public Cloud blog post](#012-infrastructure-as-code-in-a-private-or-public-cloud-blog-post)
  * [013. Introduction to IaC Documentation](#013-introduction-to-iac-documentation)
  * [014. Terraform Use Cases Documentation](#014-terraform-use-cases-documentation)
- [02. Manage infrastructure](#02-manage-infrastructure)
  * [021. Get Started with Terraform (complete all tutorials)](#021-get-started-with-terraform--complete-all-tutorials-)
    + [[./021_GetStartedAWS/03_BuildInfra](./021_GetStartedAWS/03_BuildInfra)](#--021-getstartedaws-03-buildinfra---021-getstartedaws-03-buildinfra-)
    + [Change Infra](#change-infra)
    + [Destroy Infra](#destroy-infra)
    + [Provision Infra](#provision-infra)
    + [Input Variables](#input-variables)
    + [Output Variables](#output-variables)
    + [Store Remote State](#store-remote-state)
  * [022. Providers documentation](#022-providers-documentation)
  * [023. Purpose of Terraform State documentation](#023-purpose-of-terraform-state-documentation)
  * [024. Terraform Settings documentation](#024-terraform-settings-documentation)
  * [025. Provisioners documentation](#025-provisioners-documentation)
- [03. Master the workflow](#03-master-the-workflow)
  * [031. The Core Terraform Workflow documentationç](#031-the-core-terraform-workflow-documentation-)
    + [Single player](#single-player)
    + [Working as a team](#working-as-a-team)
    + [Working with Terraform Cloud](#working-with-terraform-cloud)
  * [032. Initialize a Terraform working directory with init documentation](#032-initialize-a-terraform-working-directory-with-init-documentation)
  * [033. Validate a Terraform configuration with validate documentation](#033-validate-a-terraform-configuration-with-validate-documentation)
  * [033. Generate and review an execution plan for Terraform with plan documentation](#033-generate-and-review-an-execution-plan-for-terraform-with-plan-documentation)
  * [034. Execute changes to infrastructure with Terraform with apply documentation](#034-execute-changes-to-infrastructure-with-terraform-with-apply-documentation)
  * [035. Destroy Terraform managed infrastructure with destroy documentation](#035-destroy-terraform-managed-infrastructure-with-destroy-documentation)
- [04. Learn more subcommands](#04-learn-more-subcommands)
  * [041. Formatting configuration with fmt documentation](#041-formatting-configuration-with-fmt-documentation)
  * [042. Tainting resources with taint documentation](#042-tainting-resources-with-taint-documentation)
  * [043. Managing state with state documentation](#043-managing-state-with-state-documentation)
  * [044. Using local workspaces with workspace documentation](#044-using-local-workspaces-with-workspace-documentation)
  * [045. Importing existing resources with import documentation](#045-importing-existing-resources-with-import-documentation)
- [05. Use and create modules](#05-use-and-create-modules)
  * [051. Organize Configuration with Modules (complete all tutorials)](#051-organize-configuration-with-modules--complete-all-tutorials-)
    + [Modules Overview](#modules-overview)
    + [Use Modules From the Registry](#use-modules-from-the-registry)
    + [Build a Module](#build-a-module)
    + [Share Modules in the Private Module Registry](#share-modules-in-the-private-module-registry)
    + [Separate Development and Production Environments](#separate-development-and-production-environments)
  * [052. Finding and using modules documentation](#052-finding-and-using-modules-documentation)
  * [053. Module versioning documentation](#053-module-versioning-documentation)
  * [054. Input Variables documentation](#054-input-variables-documentation)
  * [055. Calling a child module documentation](#055-calling-a-child-module-documentation)
- [06. Read and write configuration](#06-read-and-write-configuration)
  * [061. Resources describe infrastructure objects](#061-resources-describe-infrastructure-objects)
  * [062. Data sources let Terraform fetch and compute data](#062-data-sources-let-terraform-fetch-and-compute-data)
  * [063. Resource addressing lets you refer to specific resources](#063-resource-addressing-lets-you-refer-to-specific-resources)
  * [064. Named values let you reference values](#064-named-values-let-you-reference-values)
  * [065. Complex types let you validate user-provided values](#065-complex-types-let-you-validate-user-provided-values)
  * [066. Built in functions help transform and combine values](#066-built-in-functions-help-transform-and-combine-values)
  * [067. Dynamic blocks allow you to construct nested expressions within certain configuration blocks](#067-dynamic-blocks-allow-you-to-construct-nested-expressions-within-certain-configuration-blocks)
  * [068. Vault Provider for Terraform](#068-vault-provider-for-terraform)
- [07. Manage state](#07-manage-state)
  * [071. State locking documentation](#071-state-locking-documentation)
  * [072. Sensitive data in state documentation](#072-sensitive-data-in-state-documentation)
  * [073. Reconcile state and real resources with refresh documentation](#073-reconcile-state-and-real-resources-with-refresh-documentation)
  * [074. Backends overview documentation](#074-backends-overview-documentation)
  * [075. Local backend documentation](#075-local-backend-documentation)
  * [076. Backend types documentation](#076-backend-types-documentation)
  * [077. How to configure a backend documentation](#077-how-to-configure-a-backend-documentation)
- [08. Debug in Terraform](#08-debug-in-terraform)
  * [081. Review the debugging documentation](#081-review-the-debugging-documentation)
- [09. Understand Terraform Cloud and Enterprise](#09-understand-terraform-cloud-and-enterprise)
  * [091.Terraform Cloud overview documentation](#091terraform-cloud-overview-documentation)
  * [092.Understanding Workspaces and Modules resource](#092understanding-workspaces-and-modules-resource)
  * [093.CLI workspaces documentation](#093cli-workspaces-documentation)
  * [094.Terraform Cloud workspaces documentation documentation](#094terraform-cloud-workspaces-documentation-documentation)
  * [095.Module registry blog post](#095module-registry-blog-post)
  * [096.Module registry documentation](#096module-registry-documentation)
  * [097.Why Policy as Code? blog post](#097why-policy-as-code--blog-post)
  * [098.Sentinel Policy as Code documentation](#098sentinel-policy-as-code-documentation)
  * [099.Feature comparison pricing page (scroll down for feature matrix)](#099feature-comparison-pricing-page--scroll-down-for-feature-matrix-)

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
- See LINK 

##### Build a Module
- https://learn.hashicorp.com/tutorials/terraform/module-create
- See LINK

##### Share Modules in the Private Module Registry
- https://learn.hashicorp.com/tutorials/terraform/module-private-registry
- Terraform allows you to create and share modules privately through their private module registry
- It is linked to terraform cloud
- Not too hyped on these commercial offerings...

##### Separate Development and Production Environments
- https://learn.hashicorp.com/tutorials/terraform/organize-configuration
- Cool, check [./051_UseModules/03_SeparateEnvironments](./051_UseModules/03_SeparateEnvironments) for an example with a module for different environments

#### 052. Finding and using modules documentation
- https://www.terraform.io/docs/registry/modules/use.html
- Open this: https://registry.terraform.io/
- Search!
- Many cool ones for AWS here: https://registry.terraform.io/providers/hashicorp/aws/latest
    - EKS
    - VPC
    - RDS
- And so on...

#### 053. Module versioning documentation
- https://www.terraform.io/docs/configuration/modules.html#module-versions
- When using modules from a module registry, it is recommended to explicitly specify the version
- Much like you would not use `latest` on a production Dockerfile
- Versions are available for module shosted on the terraform registry or terraform cloud's private module registry
- Modules sourced from local file paths do not support versioning since they're loaded from the same source repository

#### 054. Input Variables documentation
- https://www.terraform.io/docs/configuration/variables.html
- Input variables serve as parameters for a Terraform module
- Each variable accepted by a module must be declarted in a variable block
- Simple variable types:
    - string
    - number
    - bool
- Complex variables:
    - list(<TYPE>)
    - set(<TYPE>)
    - map(<TYPE>)
    - object({<ATTR NAME> = <TYPE>, ... })
    - tuple([<TYPE>, ...])
- You can also document the purpose of your variables:
```hcl-terraform
variable "image_id" {
  type        = string
  description = "The id of the machine image (AMI) to use for the server."
}
```

#### 055. Calling a child module documentation
- https://www.terraform.io/docs/configuration/modules.html#calling-a-child-module
- Welp, to call a module, add the `module $MODULE_NAME { }` block
- Inside the module block, you indicate the source of the module, as well as variable names
- All modules require a source argument

### 06. Read and write configuration
#### 061. Resources describe infrastructure objects
- https://www.terraform.io/docs/configuration/resources.html
- The most important element in terraform!
- Resources are things: EC2 instances, VPCs, security groups, RDS databases and so on
- Syntax is: resource "resource_type" resource_name
    ex: `resource "aws_instance" "web"`
- Each resource type is implemented by a provider (ex: AWS)
- Terraform has the documentation for each resource type
- You can use expressions to access information about resources in the same module, using the `RESOURCE_TYPE.NAME.ATTRIBUTE` syntax
- About dependencies: Some resources must be processed before others, and most of these dependencies are handled automatically by terraform
- However, when you want to be explicit about this dependency, you can use the `depends_on` meta-argument
- Meta arguments:
    - depends_on, for specifying hidden dependencies
    - count, for creating multiple resource instances according to a count
    - for_each, to create multiple instances according to a map, or set of strings
    - provider, for selecting a non-default provider configuration
    - lifecycle, for lifecycle customizations
    - provisioner and connection, for taking extra actions after resource creation
- Meta arguments are pretty cool!

#### 062. Data sources let Terraform fetch and compute data
- https://www.terraform.io/docs/configuration/data-sources.html
- Data resources allow data to be fetched or computed to be used elsewhere in terraform
- Use of data resources allows terraform to make use of informaton defined outside terraform, or defined by other terraform configs
- There are some examples in [./062_DataSources](./062_DataSources)

#### 063. Resource addressing lets you refer to specific resources
- https://www.terraform.io/docs/internals/resource-addressing.html
- You can refer to certain resources either by number: `module.module_name[resource_index]`
- Or by key: `module.foo[0].module.bar["a"]`
- Check [./062_DataSources](./062_DataSources)

#### 064. Named values let you reference values
- https://www.terraform.io/docs/configuration/expressions.html#references-to-named-values
- You can refer named values in a variety of ways
- Most of them are too specific to be done one by one, so here's a list:
    - <RESOURCE TYPE>.<NAME>
    - data.<DATA TYPE>.<NAME>
    - var.<NAME>
    - local.<NAME>
    - module.<MODULE NAME>.<OUTPUT NAME>
    - path.module
    - path.root
    - path.cwd
    - terraform.workspace

#### 065. Complex types let you validate user-provided values
- https://www.terraform.io/docs/configuration/types.html#complex-types
- Collection types
    - list(...)
    - map(...)
    - set(...)
- Structural types
    - object(...)
    - tuple(...)
- TODO: Write some config with these if you're feeling like it

#### 066. Built in functions help transform and combine values
- https://www.terraform.io/docs/configuration/functions.html
- <FUNCTION NAME>(<ARGUMENT 1>, <ARGUMENT 2>)
- There are a lot of them: numerical, string, collection, encoding, filesystem, data & time, hash & cryptop, IP Network and Type conversions
- TODO: Have a look at the most interesting ones

#### 067. Dynamic blocks allow you to construct nested expressions within certain configuration blocks
- https://www.terraform.io/docs/configuration/expressions.html#dynamic-blocks
- A dynamic block acts much like a for expression, but produces nested blocks for each element of that complex value
- In theory it looks nice, but I couldn't make it work, the docs were confusing apparently :(
- TODO: Actually try again at a later date

#### 068. Vault Provider for Terraform
- https://www.terraform.io/docs/providers/vault/index.html
- Allows you to read, write from and configure hashicorp vault
- Don't put your secrets in Terraform!
- You can use (sensitive) terraform configs to put secrets into vault
- You can also integrate terraform to read secrets from vault


### 07. Manage state
#### 071. State locking documentation
- https://www.terraform.io/docs/state/locking.html
- If your backend supports it, terraform will lock your state for all operations that could write state
- This happens automatically
- Terraform has a force unlock command to unlock the state if this failed.
- This commmand is dangerous! If you unlock the state when someone esle is holding the lock, it could cause multiple writes. This should only be used when you own the lock

#### 072. Sensitive data in state documentation
- https://www.terraform.io/docs/state/sensitive-data.html
- Terraform state can contain sensitive data, depending on your resources
- If you manage any form of sensitive data with terraform, treat the state itself as sensitive data.
- Storing the state remotely can provide better seucrity
    - Terraform cloud always encrypts the state at rest
    - The S3 backend supports encryption at rest when the encrypt option is enabled

#### 073. Reconcile state and real resources with refresh documentation
-` https://www.terraform.io/docs/commands/refresh.html`
- Used to refresh your terraform state with the state of your infrastructure in the cloud
- Here's an example of how this works: 
    - You write a terraform config with 1 ec2 instance
    - You apply the config and now have 1 ec2 instance running on AWS and registed on your state
    - You go on AWS and manually delete the instance
    - You run `terraform refresh`
    - This will see that the remote instance created by the apply previously is missing, and will refresh your local state to reflect the missing instance

#### 074. Backends overview documentation
- https://www.terraform.io/docs/backends/index.html
- A backend in terraform defines how the state is stored and how an operation such as `apply` is executed.
- You can have non-local file storage, and remote execution
- Benefits are:
    - Store state remotely
    - Keep sensitive information off disk
    - Remote operations

#### 075. Local backend documentation
- https://www.terraform.io/docs/backends/types/local.html
- The local backend stores state on the local filesystem, locks that state using system APIs, and performs operations locally.
- Every single config you wrote so far uses the local backend

#### 076. Backend types documentation
- https://www.terraform.io/docs/backends/types/index.html
- There are various backend types: local, remote, s3, artifactory, etcd, http...
- local and standard are enchanced backends (support remote operations) all others are standard

#### 077. How to configure a backend documentation
- https://www.terraform.io/docs/backends/config.html
- To configure a backend, use the `backend {}` configuration section
- When configuring a new backend for the first time, terraform gives you the option to migrate to your new backend
- There's an example of an S3 backend at [./077_S3Backend](./077_S3Backend)

### 08. Debug in Terraform
#### 081. Review the debugging documentation
- https://www.terraform.io/docs/internals/debugging.html
- Terraform has logs which can be enabled/adjusted by setting the `TF_LOG` environment variable to any value.
- Can be set to: TRACE, DEBUG, INFO, WARN or ERROR

### 09. Understand Terraform Cloud and Enterprise
#### 091.Terraform Cloud overview documentation
- https://www.terraform.io/docs/cloud/index.html
- Terraform Cloud is an application that helps teams use Terraform together
- Large enterprises can purchase Terraform Enterprise, a self-hosted distribution of Terraform Cloud 

#### 092.Understanding Workspaces and Modules resource
- https://www.hashicorp.com/resources/terraform-enterprise-understanding-workspaces-and-modules/
- Video is about terraform enterprise, bleh
- A workspace is associated with a set of terraform configs
- Sentinel: Hashicorp's policy as code framework, creates guardrails/sandboxes around thigns you can or cannot do with terraform
- A workspace consists of:
    - A terraform configuration
    - Values for variables used by the configuration
    - Persistent stored state for the resources managed by the config
    - Historical state and run logs
- Your app has its infrastructure: VPC, subnets, sg, iam, alb, asg, and so on
- You can have different teams managing different layers of the infrastructure (net team = manages vpc and subnets)
- You probably have more than one app, and more apps are likely to use the same structure/patterns!
- Each app could become a workspace apparently
- Benefits of terraform Modules:
    - Reusable infrastructure code
    - Better infrastructure organization
    - Standardize infrastructure templates across org

#### 093.CLI workspaces documentation
- https://www.terraform.io/docs/state/workspaces.html
- Certain backends support multiple named workspaces, allowing multiple states to be associated with a single configuration.

#### 094.Terraform Cloud workspaces documentation documentation
- https://www.terraform.io/docs/cloud/workspaces/index.html
- Terraform Cloud manages infrastructure collections with workspaces instead of directories. A workspace contains everything Terraform needs to manage a given collection of infrastructure, and separate workspaces function like completely separate working directories

#### 095.Module registry blog post
- https://www.hashicorp.com/blog/hashicorp-terraform-module-registry/
- The HashiCorp Terraform Module Registry gives Terraform users easy access to templates for setting up and running their infrastructure with verified and community modules
- There's an example of this in: [./051_UseModules/01_FromTheRegistry](./051_UseModules/01_FromTheRegistry)

#### 096.Module registry documentation
- https://www.terraform.io/docs/cloud/registry/index.html
- Terraform Cloud's private module registry helps you share Terraform modules across your organization.

#### 097.Why Policy as Code? blog post
- https://www.hashicorp.com/blog/why-policy-as-code/
- Shameless plug to Sentinel, which allows you to define policies in a fashion similar to Terraform

#### 098.Sentinel Policy as Code documentation
- https://www.terraform.io/docs/cloud/sentinel/index.html
- Sentinel is an embedded policy-as-code framework integrated with the HashiCorp Enterprise products. It enables fine-grained, logic-based policy decisions, and can be extended to use information from external sources.

#### 099.Feature comparison pricing page (scroll down for feature matrix)
- https://www.hashicorp.com/products/terraform/pricing/
- Looks like sentinel is only available for paid users? :(