
data "aws_region" "current" {}

output "myoutput" {
  value = data.aws_region.current.name
}