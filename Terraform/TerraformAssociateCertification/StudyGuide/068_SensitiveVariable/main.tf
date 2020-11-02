locals {
  secretpassword ="banana"
}

output "secret_output" {
  value = local.secretpassword
  sensitive = true
}