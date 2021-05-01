terraform {
  required_version = ">= 0.14"
  required_providers {
    aws = {
      version = ">= 2.28.1"
    }
    helm = {
      version = "1.3.1"
    }
    random = {
      version = "~> 2.1"
    }
    local = {
      version = "~> 1.2"
    }
    null = {
      version = "~> 2.1"
    }

  }
}

provider "kubernetes" {
  load_config_file       = "false"
  host                   = data.aws_eks_cluster.cluster.endpoint
  token                  = data.aws_eks_cluster_auth.cluster.token
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
}

provider "aws" {
  region = "eu-west-1"
}

//
//provider "template" {
//  version = "~> 2.1"
//}
