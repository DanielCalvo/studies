provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

resource "aws_key_pair" "example" {
  key_name   = "examplekey"
  public_key = file("/home/daniel/PycharmProjects/studies/Terraform/learn.hashicorp.com/GettingStarted_AWS/02_provisioner/remote-exec/terraform.pub")
}

resource "aws_instance" "example" {
  key_name      = aws_key_pair.example.key_name
  ami           = "ami-099a8245f5daa82bf"
  instance_type = "t2.micro"

  connection {
    type        = "ssh"
    user        = "ec2-user"
    private_key = file("/home/daniel/PycharmProjects/studies/Terraform/learn.hashicorp.com/GettingStarted_AWS/02_provisioner/remote-exec/terraform")
    host        = self.public_ip
  }

  provisioner "remote-exec" {
    inline = [
      "ls"
    ]
  }
}