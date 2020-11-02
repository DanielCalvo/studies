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
  default = ""
}

variable "otherlist" {
  default = ["a", 15, true] //terraform converts everything to string, wew
}

resource "null_resource" "just_nulling" {
  for_each = toset(var.otherlist)
  provisioner "local-exec"{
    command = "echo Hello ${each.value}}"
  }
}