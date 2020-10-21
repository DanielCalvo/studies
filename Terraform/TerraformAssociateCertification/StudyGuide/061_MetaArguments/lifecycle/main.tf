provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

resource "aws_instance" "just_testing" {
  ami = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro" //If you change instance type and apply, terraform does not recreate the resource
  tags = {
    Name = "asdasd"
  }
  lifecycle {
    create_before_destroy = true
    ignore_changes = [instance_type]
  }
}

resource "aws_instance" "instance_that_doesnt_go_away" {
  ami = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"
  tags = {
    Name = "asdasd"
  }
  lifecycle {
    prevent_destroy = false //Ha, can't run terraform destroy with this set to true
  }
}

