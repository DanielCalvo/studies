variable "tuplesito" {
  type = tuple([number, bool, string])
  default = [21, false, "3"]
}

output "print" {
  value = var.tuplesito
}