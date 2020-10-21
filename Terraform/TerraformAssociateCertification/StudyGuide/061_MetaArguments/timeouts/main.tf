provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

resource "aws_instance" "timeouts" {
  ami = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"
  tags = {
    Name = "asdasd"
  }
  timeouts {
    create = "60m"
    delete = "2h"
    update = "3h"
  }
}
