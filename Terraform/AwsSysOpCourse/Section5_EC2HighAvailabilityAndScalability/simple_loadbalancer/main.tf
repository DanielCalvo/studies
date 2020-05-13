provider "aws" {
  profile    = "default"
  region     = "eu-west-1"
}

//variable "vm_names" {
//  description = "Create EC2 instances with these names"
//  type        = list(string)
//  default     = ["InstanceOne", "InstanceTwo"]
//}
//
//resource "aws_instance" "example" {
//  ami           = "ami-099a8245f5daa82bf"
//  instance_type = "t2.micro"
//  count = length(var.vm_names)
//  tags = {
//    Name  = var.vm_names[count.index]
//  }
//}

resource "aws_instance" "InstanceOne" {
  ami           = "ami-099a8245f5daa82bf"
  instance_type = "t2.micro"
  tags = {
    Name  = "InstanceOne"
  }
}

resource "aws_instance" "InstanceTwo" {
  ami           = "ami-099a8245f5daa82bf"
  instance_type = "t2.micro"
  tags = {
    Name  = "InstanceTwo"
  }
}

resource "aws_lb_target_group" "test" {
  name     = "tf-example-lb-tg"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id
}