provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

resource "aws_instance" "just_testing" {
  for_each = toset(var.instance_names)
  ami = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"
  tags = {
    Name = each.key
  }
}
