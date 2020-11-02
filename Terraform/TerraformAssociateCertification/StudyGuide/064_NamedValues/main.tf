terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

variable something {
  type = string
  default = local.banana
}

variable somethingelse {
  type = list
  default = [ local.banana ]
}

locals {  banana = "banana" }
locals {  apple = "apple" }

locals {
  something = "1"
  something = "2"
}

locals {
  fruits = [local.banana, local.apple]
  asd = var.something
}

output "myoutput" {
  value = local.banana
}

output "myoutputsomething" {
  value = local.something
}

//See if you can do an if statement too!