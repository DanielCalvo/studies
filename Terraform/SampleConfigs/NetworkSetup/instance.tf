resource "aws_instance" "SingleInstanceWebHello" {
  ami                    = "ami-07d9160fa81ccffb5"
  instance_type          = "t2.micro"
  user_data              = "${file("install_apache.sh")}"
  subnet_id              = aws_subnet.experimental_eu-west-1a.id
  vpc_security_group_ids = [aws_security_group.ec2_http_ssh.id]
  tags = {
    Name        = "SingleInstanceWebHello"
    Environment = "experimental"
  }
}