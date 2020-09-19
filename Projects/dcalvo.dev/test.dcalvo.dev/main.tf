module "s3_website" {
  source = "../terraform_modules/s3_website"
  domain-name = "test.dcalvo.dev"
  bucket-name = "dcalvo-mctestbucketson"
}