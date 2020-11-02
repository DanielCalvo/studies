
variable "banana" {
  type =  string
  default = "asdasd"
}

output "outputting_my_stuff" {
  value = var.banana
}