- Make sure you have an AWS account
- Configure some auth method for you to access your account
- Ensure Terraform is installed
- Open this on your browser: https://registry.terraform.io/modules/terraform-aws-modules/vpc/aws/2.21.0

### Actual steps to use the module:
```
git clone https://github.com/hashicorp/learn-terraform-modules.git
cd learn-terraform-modules
git checkout tags/ec2-instances -b ec2-instances
```
- Define root input variables for the module in variables.tf
- When you run `terraform init` terraform downloads the remote module into the local `.terraform` directory
- A module is indeed like a software function: You specify the input, it does something, and then you get something done and/or an output. Neat!
- Could be useful for multiple environments or projects, or DRY.