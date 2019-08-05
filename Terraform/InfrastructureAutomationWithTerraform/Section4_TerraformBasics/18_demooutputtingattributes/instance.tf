resource "aws_instance" "daniinstance"{
  ami = "${lookup(var.AMIS, var.AWS_REGION)}"
  instance_type = "t2.micro"
  key_name = ""

  provisioner "local-exec" {
    command = "echo ${aws_instance.daniinstance.private_ip} >> private_ips.txt"
  }
}

output "ip" {
  value = "${aws_instance.daniinstance.public_ip}"
}