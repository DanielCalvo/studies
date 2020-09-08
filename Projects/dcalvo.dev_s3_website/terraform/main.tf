provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

resource "aws_s3_bucket" "dcalvo-dev-bucket" {
  bucket = "dcalvo-dev-bucket"
  acl    = "public-read"

  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}
