//variable "AWS_REGION" {
//  default = "eu-west-1"
//}

//provider "aws" {
//  access_key = var.AWS_ACCESS_KEY
//  secret_key = var.AWS_SECRET_KEY
//  region     = var.AWS_REGION
//}

provider "aws" {
  profile    = "default"
  region     = "eu-west-1"
}

//resource "aws_instance" "example" {
//  ami           = "ami-078603b469de54ad7"
//  instance_type = "t2.micro"
//}

resource "aws_instance" "example" {
  ami           = "ami-099a8245f5daa82bf"
  instance_type = "t2.micro"
}

//variable "AWS_ACCESS_KEY" {
//}
//
//variable "AWS_SECRET_KEY" {
//}


variable "AMIS" {
  type = map(string)
  default = {
    us-east-1 = "ami-13be557e"
    us-west-2 = "ami-06b94666"
    eu-west-1 = "ami-078603b469de54ad7"
  }
}


terraform {
  required_version = ">= 0.12"
}
