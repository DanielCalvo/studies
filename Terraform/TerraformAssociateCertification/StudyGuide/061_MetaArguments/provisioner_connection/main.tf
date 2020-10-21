//provider "aws" {
//  profile = "default"
//  region  = "eu-west-1"
//}

resource "aws_instance" "my-instance" {
  ami = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"
  tags = {
    Name = "myinstance"
  }

  provisioner "remote-exec" {
    inline = [
      "cat /etcc/passwd",
    ]
    on_failure = continue
  }

  provisioner "remote-exec" {
    inline = [
      "cat /etcc/passwd",
    ]
    on_failure = continue
  }

  provisioner "local-exec" { //consider using an output variable for this
    when = create
    command = "echo The server IP address is ${self.private_ip} > /tmp/ip"
  }

  provisioner "local-exec" { //consider using an output variable for this
    when = destroy
    command = "echo bleep blop: Instance destroyed"
  }
}

resource "null_resource" "doing_things" {
  provisioner "local-exec" {
    command = "echo hi"
  }
}