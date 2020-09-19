//It is necessary to hardcode us-east-1 here as ACM certificates for use with S3 and Cloudfront must be in us-east-1
provider "aws" {
  alias  = "us-east-1"
  region = "us-east-1"
}
