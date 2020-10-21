1. Remember that the sensitive argument only keeps sensitive data from being printed, not saved in the state

3. Can another process unlock local state files?
- No

5. integer is not a valid variable type. count is a reserved word

7. Terraform show will show human-readable output from a state or plan file

8. Modules sourced from local file paths do not support `version`
- TODO: Import some modules 

11. Remember that terraform deletes resources if you rename them

12. Once a resource is marked as tainted, terraform will recreate it on the next terraform apply

--- Continue here tomorrow ---

- Double check variable definitions (what's acceptable?)
- Go exploring the cmdline again and it's capabilities
- Review workspaces (you didn't have a look at them)

- The terraform import command currently can only import one resource at a time? I think this is true

- Sentinel is a __________ service that prevents the provisioning of out-of-policy infrastructure. Review sentinel!

Can the expression of a local value refer to other locals? No (but doublecheck)

- TODO: Oh, logfiles are a thing?

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