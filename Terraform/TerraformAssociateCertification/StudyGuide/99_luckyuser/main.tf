variable "users" {
  type    = list(string)
  default = ["Peter", "Lois", "Brian", "Meg"]
}

output "lucky_user" {
  value = element(var.users, 4)
}


