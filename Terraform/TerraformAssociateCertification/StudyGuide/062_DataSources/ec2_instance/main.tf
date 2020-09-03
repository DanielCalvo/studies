provider "aws" {
  profile    = "default"
  region     = "eu-west-1"
}

data "aws_instance" "dcalvo_dev" {
  filter {
    name   = "tag:Name"
    values = ["dcalvo.dev"]
  }
}

output "dcalvo_dev_private_ip" {
  value = data.aws_instance.dcalvo_dev.private_ip
}

output "dcalvo_dev_public_ip" {
  value = data.aws_instance.dcalvo_dev.public_ip
}