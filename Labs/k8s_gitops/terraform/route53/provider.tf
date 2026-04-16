terraform {
  required_version = "v1.14.8"
  backend "s3" {
    bucket = "dani-terraform-states"
    key    = "labs/k8s_gitops/terraform/route53/vpc.tfstate"
    region = "eu-west-1"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.40"
    }
  }
}

provider "aws" {
  region = "eu-west-1"

  default_tags {
    tags = {
      lab = "k8s-gitops-route53"
    }
  }
}

