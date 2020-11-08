### Exam 1

1. Remember that the sensitive argument only keeps sensitive data from being printed, not saved in the state
3. Can another process unlock local state files?
    - No
5. integer is not a valid variable type. count is a reserved word
7. Terraform show will show human-readable output from a state or plan file
8. Modules sourced from local file paths do not support `version`
11. Remember that terraform deletes resources if you rename them
12. Once a resource is marked as tainted, terraform will recreate it on the next terraform apply
13. Terraform only recommends using provisioners as a last resort -- use configuration management tools first (like Ansible)
14. To pass variables to modules, declare them in the module block
15. If a creation time provisioner fails, the resource will be marked as tainted. You can change this with the `on_failure` setting
19. Terraform import can only import once resource at a time 
22. Sentinel is a PROACTIVE service that prevents the provisioning of out-of-policy infrastructure.
    - Review sentinel: https://www.hashicorp.com/blog/using-terraform-to-improve-infrastructure-security
23. Locals can reference variables from other locals
    - DONE: Implement something with local {}
25. Arbitrary Git repositories can be used by prefixing the address with the special git:: prefix.
    - DONE: Write a module and put it on a private git repository, then import it and use it!
26. The resources defined in a module are encapsulated, so the calling module cannot access their attributes directly. 
28. The keyword list is a shorthand for list(any)
30. A connection block nested directly within a resource affects all of that resource's provisioners
30. A connection block nested in a provisioner block only affects that provisioner, and overrides any resource-level connection settings.
32. Every Terraform configuration has at least one module, known as its root module, which consists of the resources defined in the .tf
36. Terraform 0.12 recommends using the `required_providers` in the terraform block, and not using the version argument within provider blocks.
    - version is still supported for compatibility with older terraform versions
42. Providers can be passed down to descendent modules in two ways: either implicitly through inheritance
43: Setting the `TF_LOG` variable makes logs appear on stderr (not stdout!) 
44. Even if you use the sensitive variable variable option, passwords will still be stored in the state
48. If you want a different branch from a module, pass the ref option, ex: git::https://code.quizexperts.com/vpc.git?ref=v1.2.0
    - You don't use "branch" or "version" fields on terraform 
51. `TF_LOG` is set to TRACE if it's set to anything other than a supported log level
- Stopped on question 50
55. Read this for the valid module URL import syntax/sequence thing: https://www.terraform.io/docs/registry/modules/use.html#private-registry-module-sources
56. Do something with a tuple: DONE 
57.  something with a workspace: DONE


###Notes to put in the main doc later
22. 
    - Sentinel is a policy as a code framework that proactively prevents provisioning of out-of-policy infrastructure

23.
    - You can access variable definitions inside a local block
    - You can access variable from a local block inside another local block
    - You can access variables defined inside a local block in a variable definition block
    - A common pattern seems to be to define variables in a variable block somewhere and then use a local block to process them
42.
    - Provider set up is inherited by child modules
    - Providers can be passed down to descendent modules in two ways: either implicitly through inheritance, or explicitly via the providers argument within a module block
    - Defined variables on the root module (like var.something) are not inherited by child modules

--- Continue here tomorrow ---

- Double check variable definitions (what's acceptable?)
- Go exploring the cmdline again and it's capabilities
- Review workspaces (you didn't have a look at them)

- The terraform import command currently can only import one resource at a time? I think this is true

- Sentinel is a __________ service that prevents the provisioning of out-of-policy infrastructure. Review sentinel!

Can the expression of a local value refer to other locals? No (but doublecheck)

Remote state allows teams to share infrastructure resources in a read-only way without relying on any additional configuration store. ???
    - Ambiguous
    
- Do a remote provisioner with a connection once and pay attention to the syntax

Remote state allows teams to share infrastructure resources in a read-only way without relying on any additional configuration store.
- ??

The parent module has to provide input variables and providers to a child module explicitly in the module block as this information can not be passed down to descendent modules implicitly through inheritance.- I think this is true

- didn't know there were rules as far as naming modules went :o
To publish a module to Terraform Public Registry. A module must meet the following requirements.
Choose TWO correct answers.

- How do you get a module on a given branch?

How to add provider to Terraform v0.12 configuration?
- I thought you had to add the provider configureatoin block, but can you just add the resources and do terraform init? Hmmm

Does terrasform have more environment variables like TF_LOG_PATH?

- How do you use and import a module from a private registry again?
- Where is workspace state infromation stored again? (directory structure)

- Which of the following configuration block has Terraform 0.12 recommended way of provider version constraint?
    - How to specify provider requirements in terraform 0.12 again? No idea :o
    
- I think you messed up on a question about providers being "required" or something like that too.
- I didn't even know modules had an obligatory naming scheme?


---

### Exam 2
- Really review the command line flags again :sisi:
3. `terraform state mv` can be used to rename resources, move items from and to a module and move data to a completely new state
    - This command will output a backup copy of the state prior to saving any changes
4. Terraform provisioner blocks only work within resource definitions. How did you make this mistake man?
7. Terraform init downloads "the source code of modules" (huh?) to a directory on local disk
    - Note from Dani: Strictly speaking this command doesn't download source code, it downloads binary files
10. Audit logging and SAML SSO are exclusive features of terraform enterprise. Sentinel isn't
    - Terraform enterprise is a self hosted distribution of terraform cloud
14. Terraform does not support user definined functions
18. The same variable can only be asigned once in the same source
    - If assigned more than once in different sources, it'll use the last assigned value
19. Terraform import can import resources into modules! It can also import into resources configured with count and foreach
20. Removing the remote state will just cause things to be recreated, right?
    - On S3 if you delete the state file, yes! If you delete the entire bucket, it fails. But this exam question is apparently wrong
21. Additional provider configurations are never inherited automatically by child modules
    - You must pass them explicitly using the providers map
28. A destroy-time provisioner within a resource that is tainted will not run
29. Workspaces in terraform cloud and the CLI are not the same. Terraform keeps previous copies of state files and it also keeps a run history
30. The element function wraps around. FML 
32. `terraform init -upgrade` upgrades the terraform provider to the latest acceptable version
35. Use `terraform force-unlock` to unlock state!
37. Terraform automatically converts value from string to number to boolean when required. So you can assign "true" to a boolean variable :o
44. Cost estimation is a paid feature on terraform cloud
45. `terraform init` can install third party providers
48. The use of the version constraint within `required_providers` is recommended in terraform 0.13 or later
49. To ensure that workspace names are stored correctly and safely in all backends, the name must be valid to use in a URL path segment without escaping

### Exam 3
6. The wording on this was too long and I was just too tired
10. Each Terraform configuration has an associated PROVIDER that defines how operations are executed and where persistent data such as the Terraform state are stored.
13. The `min` function takes arguments like this:
    - `min(1,55,2)`
    - But if the arguments are in a list, it takes them like this: `min([1,454,2]...)`
16. Dynamic Blocks allow us to dynamically construct repeatable nested blocks.
    - TODO: A good use case would try to be using it when defining a security group with repeating ingress rules
20. Read this real quick: https://www.terrafins
orm.io/docs/commands/environment-variables.html
    - `export TF_LOG=TRACE`
    - `export TF_LOG=`
    - `export TF_LOG_PATH=./terraform.log`
    - `export TF_INPUT=0`
    - `export TF_VAR_region=us-west-1`
    - `export TF_VAR_ami=ami-049d8641`
    - `export TF_VAR_alist='[1,2,3]'`
    - `export TF_VAR_amap='{ foo = "bar", baz = "qux" }'`
    - `export TF_CLI_ARGS="-input=false"`
    - `export TF_CLI_ARGS_plan="-refresh=false"`
21. Your definition of "data source" is not clear
28. For local state, if using workspaces terraform store the state under `terraform.tfstate.d`. If you're not using workspaces, the state is in `terraform.tfstate`
32. Modules on S3 do *not* have to be public for you to be able to use them. Terraform looks for AWS credentials in the same place as the AWS CLI tool to access the bucket
34. Every terraform command that modifies the state generates a backup file, by default it looks like it's named `terraform.tfstate.backup`
35. Oh boy! Looks like terraform does not accept single quotes when passing arguments to string functions! 
40.Check variable definition precedence here: https://www.terraform.io/docs/configuration/variables.html#variable-definition-precedence
    - Alrite, they're loaded in this order:
    - The terraform.tfvars file, if present.
    - The terraform.tfvars.json file, if present.
    - Any *.auto.tfvars or *.auto.tfvars.json files, processed in lexical order of their filenames.
    - Any -var and -var-file options on the command line, in the order they are provided. (This includes variables set by a Terraform Cloud workspace.)

### Overall
- Just read the docs man. ACtually read them. No big deal, just read