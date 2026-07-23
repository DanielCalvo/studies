data "aws_availability_zones" "available" {
  # Exclude local zones
  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

data "aws_ecrpublic_authorization_token" "token" {
  region = "us-east-1"
}


data "terraform_remote_state" "vpc_remote_state" {
  backend = "s3"
  config = {
    bucket = "dani-terraform-states"
    key    = "labs/k8s_gitops/terraform/vpc/vpc.tfstate"
    region = "eu-west-1"
  }
}
