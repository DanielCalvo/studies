resource "aws_s3_bucket" "b" {
  bucket = "dcalvo-bucket-fff00"
  acl    = "private"

  tags = {
    Name = "dcalvo-bucket-fff00"
  }
}

