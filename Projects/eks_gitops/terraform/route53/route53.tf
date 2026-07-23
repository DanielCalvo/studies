resource "aws_route53_zone" "dcalvo-dev" {
  name = "dcalvo.dev"
}

resource "aws_route53_zone" "k8s-gitops" {
  name = "k8s-gitops.dcalvo.dev"
}
