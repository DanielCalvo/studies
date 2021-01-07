output "instance_ip" {
  value = aws_instance.daniinstance.public_ip
}

output "alb_url" {
  value = aws_lb.dani_lb.dns_name
}

output "domain_name" {
  value = aws_route53_record.alb_wildcard_dns.name
}