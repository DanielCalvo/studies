terraform {
  required_version = ">= 0.14"
  required_providers {
    aws = {
      version = ">= 3.22.0"
    }
  }
}

provider "aws" {
  region = "eu-west-1"
}