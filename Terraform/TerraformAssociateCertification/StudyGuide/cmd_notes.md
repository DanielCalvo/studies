### terraform apply
- Builds or changes infrastructure according to terraform config files in DIR
- Backs up the state by default. Locks the state by default. State file defaults to `terraform.tfstate`
- Backup file defaults to `terraform.tfstate.backup`
- You can target a specific resource/module to be applied (as opposed to everything)
- You can pass variables on the command line with `-var foo=bar` and you can also specify a variable file.  If "terraform.tfvars" or any ".auto.tfvars" files are present, they will be automatically loaded

### terraform console
-  Starts an interactive console for experimenting with Terraform interpolations

### terraform destroy
- Takes the same arguments as terraform apply, but instead of creating infra, destroys it

### terraform env
- For workspace management. Deprecated. Use `terraform workspace` instead

### terraform fmt
- Rewrites all Terraform configuration files to a canonical format
- You can use `-diff` to display an output of formatted changes

### terraform get
- Downloads and installs modules needed for the configuration given by PATH

### terraform graph

### terraform import
- Import existing infrastructure into your Terraform state. This can only import resources into the state, it does not generate configuration

### terraform init
- Initialize a new or existing Terraform working directory by creating initial files, loading any remote state, downloading modules, etc.

### terraform login
- Retrieves an authentication token for the given hostname, if it supports automatic login, and saves it in a credentials file in your home directory.
- If no hostname is provided, the default hostname is app.terraform.io, to log in to Terraform Cloud.

### terraform logout
- Removes locally-stored credentials for specified hostname.
- If no hostname is provided, the default hostname is app.terraform.io.


### terraform output
- Reads an output variable from the terraform state and prints it
- With no arguments, prints all the terraform variables from the state

### terraform plan
- Generates an execution plan for Terraform.
- This plan can be saved to a file

### terraform providers
- Prints out a tree of modules in the referenced configuration annotated with their provider requirements

### terraform refresh
- The terraform refresh command is used to reconcile the state Terraform knows about (via its state file) with the real-world infrastructure. This can be used to detect any drift from the last-known state, and to update the state file

### terraform show
- Reads and outputs a Terraform state or plan file in a human-readable form. If no path is specified, the current state will be shown

### terraform taint
- Manually mark a resource as tainted, forcing a destroy and recreate on the next plan/apply.

### terraform untaint
- Manually unmark a resource as tainted, restoring it as the primary instance in the state
- This reverses either a manual 'terraform taint' or the result of provisioners failing on a resource

### terraform validate
- Validates the configuration files in a current directory, refeering only to the configuration and not accessing any remotes services (such as remote state, providers and APIs)

### terraform version
- Displays the version of Terraform and all installed plugins

### terraform workspace
#### terraform workspace new
- Creates a new terraform workspace

#### terraform workspace list
- Lists existing terraform workspaces

#### terraform workspace show
- Shows the name of the current workspace

#### terraform workspace select
- Selects an existing workspace

#### terraform workspace delete
- Deletes a workspace

### terraform 0.12upgrade
- Deprecated

### terraform 0.13upgrade
- This command will update the configuration files in the given directory to use the new provider source features from Terraform v0.13

### terraform debug
- Debug command with no documentation available (?)

### terraform force-unlock
- Manually unlock the state for the defined configuration.

### terraform push
- Was used with legacy terraform enterprise, no longer supported

### terraform state
#### terraform state list
- List resources in the Terraform state.

#### terraform state mv
- This command will move an item matched by the address given to the destination address. This command can also move to a destination address in a completely different state file
- This can be used for simple resource renaming, moving items to and from a module, moving entire modules, and more. And because this command can also  move data to a completely new state, it can also be used for refactoring one configuration into multiple separately managed Terraform configurations

####terraform state pull
- Pull the state from its location and output it to stdout

#### terraform state push
- This command "pushes" a local state and overwrites remote state with a local state file 

#### terraform state replace-provider
- This command will update all resources using the "from" provider, setting the provider to the specified "to" provider. This allows changing the source of a provider which currently has resources in state.
- Ex: `terraform state replace-provider hashicorp/aws registry.acme.corp/acme/aws`

#### terraform state rm
- Usage: `Usage: terraform state rm [options] ADDRESS`
- Remove one or more items from the terraform state, causing terraform to "forget" about them without destroying them in the remote system

#### terraform state show
- Usage: `terraform state show [options] ADDRESS`
- Shows the attributes of a single resource in the terraform state
- Don't confuse this with `terraform show`, which shows atributes for the entire state