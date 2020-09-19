provider "aws" {
  region  = "eu-west-1"
  profile = "default"
}

terraform {
  backend "s3" {
    bucket = "dani-terraform-states"
    key    = "dcalvo.dev/test.dcalvo.dev.tfstate"
    region = "eu-west-1"
  }
}