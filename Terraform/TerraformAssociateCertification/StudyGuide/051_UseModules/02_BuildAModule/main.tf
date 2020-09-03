provider "aws" {
  region = "eu-west-1"
}

module "website_s3_bucket" {
  source = "./s3_website"

  bucket_name = "daniveryuniquename"

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}