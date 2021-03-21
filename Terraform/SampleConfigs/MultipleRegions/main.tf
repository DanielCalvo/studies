provider "aws" {
  alias   = "eu-west-1"
  region  = "eu-west-1"
  profile = "default"

}

provider "aws" {
  alias   = "us-east-1"
  region  = "us-east-1"
  profile = "default"
}

resource "aws_instance" "hello-eu-west-1" {
  ami           = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"
  provider = aws.eu-west-1
  tags = {
    Name        = "hello-eu-west-1"
    Environment = "experimental"
  }
}

resource "aws_instance" "hello-us-east-1" {
  ami = "ami-0c94855ba95c71c99"
  instance_type = "t2.micro"
  provider = aws.us-east-1
  tags = {
    Name = "hello-us-east-1"
    Environment = "experimental"
  }
}
