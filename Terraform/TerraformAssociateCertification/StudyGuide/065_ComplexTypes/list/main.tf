variable "names" {
  type = list(string)
  default = ["Alice", "Bob", "Claire"]
}

// valid
variable "mylist" {
  default = []
}

//also valid! :o
variable "something" {

}

//oh lord
variable "anything" {
  type = any
}

resource "null_resource" "just_nulling" {
  for_each = toset(var.names)
  provisioner "local-exec"{
    command = "echo Hello ${each.value}}"
  }
}