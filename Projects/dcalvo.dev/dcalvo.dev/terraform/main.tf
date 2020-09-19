data "terraform_remote_state" "dcalvo-dev-domain" {
  backend = "s3"
  config = {
    bucket = "dani-terraform-states"
    key    = "dcalvo.dev/dns-domain.tfstate"
    region = "eu-west-1"
  }
}

module "s3_website" {
  source = "../../terraform_modules/s3_website"
  domain-name = "dcalvo.dev"
  bucket-name = "dcalvo-dev-bucket"
  bucket-extra-provision-step = "?"
}