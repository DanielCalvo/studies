output "zone_id" {
  description = "The ID of the DNS zone"
  value = aws_route53_zone.dcalvo-dev-zone.id
}

output "namesevers" {
  description = "Nameservers for this zone"
  value = aws_route53_zone.dcalvo-dev-zone.name_servers
}