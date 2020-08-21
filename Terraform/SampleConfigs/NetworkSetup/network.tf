resource "aws_vpc" "experimental" {
  cidr_block           = "10.0.0.0/24"
  instance_tenancy     = "default"
  enable_dns_support   = "true"
  enable_dns_hostnames = "true"
  enable_classiclink   = "false"
  tags = {
    Name = "experimental"
  }
}

resource "aws_subnet" "experimental_eu-west-1a" {
  vpc_id                  = aws_vpc.experimental.id
  cidr_block              = "10.0.0.0/25"
  map_public_ip_on_launch = "true"
  availability_zone       = "eu-west-1a"
  tags = {
    Name = "experimental_eu-west-1a"
  }
}

#Security group
resource "aws_security_group" "ec2_http_ssh" {
  name        = "ec2_http_ssh"
  description = "Allow HTTP and SSH traffic"
  vpc_id      = aws_vpc.experimental.id

  ingress {
    description = "Allowing HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allowing SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "ec2_http_ssh"
  }
}

#Internet Gateway
resource "aws_internet_gateway" "experimental_gw" {
  vpc_id = aws_vpc.experimental.id
  tags = {
    Name = "experimental_gw"
  }
}

#Route table
#II think you need to add the route table to the subnet on subnet creation
resource "aws_route_table" "experimental_rt" {
  vpc_id = aws_vpc.experimental.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.experimental_gw.id
  }
  tags = {
    Name = "experimental_rt"
  }
}

resource "aws_route_table_association" "experimental_route_association" {
  subnet_id      = aws_subnet.experimental_eu-west-1a.id
  route_table_id = aws_route_table.experimental_rt.id
}