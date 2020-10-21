provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

resource "aws_s3_bucket" "b" {
  bucket = "test-bucket-danitest"
  acl    = "private"
  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

//Maybe this instance depends on a S3 bucket to perform it's work, but terraform can't see that dependency (IAM role not included)
resource "aws_instance" "just_testing" {
  ami = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"
  tags = {
    Name = "just testing"
  }
  depends_on = [aws_s3_bucket.b]
}
