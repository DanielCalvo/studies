
locals {
  myvar = chomp(file("/home/daniel/Projects/gh_token.secret"))
}

output "myvar" {
  value = local.myvar
}