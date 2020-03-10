terraform {
  backend "s3" {
    bucket = "terraform-state-dcalvo"
    key = "terraform/remotestatedemo"
  }
}