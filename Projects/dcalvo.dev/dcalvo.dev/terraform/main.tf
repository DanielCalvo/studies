data "terraform_remote_state" "dcalvo-dev-zone" {
  backend = "s3"
  config = {
    bucket = "dani-terraform-states"
    key    = "dcalvo.dev/dns-domain.tfstate"
    region = "eu-west-1"
  }
}

module "s3_website" {
  source = "../../terraform_modules/s3_website"
  domain-name = var.domain-name
  bucket-name = var.bucket-name
}

resource "null_resource" "populate-s3-bucket" {
  depends_on = [module.s3_website]
  provisioner "local-exec" {
    command = "aws s3 cp ../assets s3://${var.bucket-name} --recursive"
  }
}