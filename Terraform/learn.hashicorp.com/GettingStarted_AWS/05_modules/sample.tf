#terraform {
#  required_version = "0.11.11"
#}

provider "aws" {
  region     = "eu-west-1"
}

module "consul" {
  source      = "hashicorp/consul/aws"
  num_servers = "3"
}