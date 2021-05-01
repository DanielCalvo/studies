provider "aws" {
  region  = "eu-west-1"
  profile = "default"
}

provider "aws" {
  alias  = "us-east-1"
  region = "us-east-1"
}

terraform {
  backend "s3" {
    bucket = "dani-terraform-states"
    key    = "dcalvo.dev/dcalvo.dev.tfstate"
    region = "eu-west-1"
  }
}
