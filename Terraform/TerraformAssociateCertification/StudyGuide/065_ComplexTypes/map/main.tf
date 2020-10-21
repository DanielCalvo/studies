provider "aws" {
  profile    = "default"
  region     = "eu-west-1"
}

variable "ami_id" {
  type = map
  default = {
    ubuntu   = "ami-06fd8a495a537da8b"
    amazon-linux2 = "ami-0bb3fad3c0286ebd5"
    redhat = "ami-08f4717d06813bf00"
  }
}

resource "aws_instance" "my-instance" {
  ami = var.ami_id["ubuntu"]
  instance_type = "t2.micro"
  tags = {
    Name = "myinstance"
  }
}