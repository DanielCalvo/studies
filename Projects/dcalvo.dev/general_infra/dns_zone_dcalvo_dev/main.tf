//Do not destroy -- domain names have to be manually changed on domains.google.com every time this is created
resource "aws_route53_zone" "dcalvo-dev-zone" {
  comment = "Written by Dani with Terraform"
  name    =  "dcalvo.dev"
}