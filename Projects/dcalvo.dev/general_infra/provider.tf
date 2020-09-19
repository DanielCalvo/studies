provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

terraform {
  backend "s3" {
    bucket = "dani-terraform-states"
    key    = "dcalvo.dev/dns-domain.tfstate"
    region = "eu-west-1"
  }
}