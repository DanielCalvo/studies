data "terraform_remote_state" "dcalvo-dev-domain" {
  backend = "s3"
  config = {
    bucket = "dani-terraform-states"
    key    = "dcalvo.dev/dns-domain.tfstate"
    region = "eu-west-1"
  }
}

variable "domain-name" {
    type = string
    default = ""
}

variable "bucket-name" {
    type = string
    default = ""
}