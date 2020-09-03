provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

resource "aws_instance" "server" {
  count = 2

  ami           = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"
  tags = {
    Name = "Server ${count.index}"
  }


}

resource "aws_instance" "other_server" {
  ami           = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"
  for_each      = toset(["Server one", "Server two"])
  tags = {
    Name = each.key
  }
}

output "count_ip" {
  value = aws_instance.server[0].public_ip
}

output "foreach_ip" {
  value = aws_instance.other_server["Server two"].public_ip
}