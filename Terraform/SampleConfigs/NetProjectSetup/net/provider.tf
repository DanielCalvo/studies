provider "aws" {
  profile = "default"
  region  = var.region
}

terraform {
  backend "s3" {
    bucket = "dani-terraform-states"
    key    = "network/terraform.tfstate"
    region = "eu-west-1"
  }
}