terraform {
  required_version = ">= 0.15"
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "= 2.0.3"
    }
    github = {
      source  = "integrations/github"
      version = ">= 4.5.2"
    }
    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.10.0"
    }
    flux = {
      source  = "fluxcd/flux"
      version = ">= 0.0.13"
    }
    tls = {
      source  = "hashicorp/tls"
      version = "3.1.0"
    }
  }
}