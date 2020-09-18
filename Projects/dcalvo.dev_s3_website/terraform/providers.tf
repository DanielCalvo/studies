provider "aws" {
  region  = "eu-west-1"
  profile = "default"
}

provider "aws" {
  alias  = "us-east-1"
  region = "us-east-1"
}
