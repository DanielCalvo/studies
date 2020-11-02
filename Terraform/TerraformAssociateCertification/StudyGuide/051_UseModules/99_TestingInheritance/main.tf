terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

provider "aws" {
  region = "eu-west-1"
}

variable "banana" {
  type = string
  default = "banana"
}
module "danimodule" {
  source = "./danimodule"
}

output "uuuh" {
  value = module.danimodule.myoutput
}