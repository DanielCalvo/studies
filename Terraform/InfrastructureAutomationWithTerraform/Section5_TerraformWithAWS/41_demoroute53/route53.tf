resource "aws_route53_zone" "dcalvo-eu" {
  name = "dcalvo.eu"
}

resource "aws_route53_record" "server1-record" {
  zone_id = aws_route53_zone.dcalvo-eu.zone_id
  name    = "test.dcalvo.eu"
  type    = "A"
  ttl     = "300"
  records = ["34.247.47.245"]
}

//resource "aws_route53_record" "www-record" {
//  zone_id = aws_route53_zone.dcalvo-eu.zone_id
//  name    = "www.newtech.academy"
//  type    = "A"
//  ttl     = "300"
//  records = ["104.236.247.8"]
//}

output "ns-servers" {
  value = aws_route53_zone.dcalvo-eu.name_servers
}

