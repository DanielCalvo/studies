data "terraform_remote_state" "network" {
  backend = "s3"
  config = {
    bucket = "dani-terraform-states"
    key    = "network/terraform.tfstate"
    region = "eu-west-1"
  }
}


output "vpc-id" {
  value = data.terraform_remote_state.network.vpc-id
}