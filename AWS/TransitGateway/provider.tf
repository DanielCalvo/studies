terraform {
  required_version = ">= 1.8.2"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "eu-west-1"
  default_tags {
    tags = {
      Terraform = "true"
      lab = "transitgateway"
    }
  }
}
