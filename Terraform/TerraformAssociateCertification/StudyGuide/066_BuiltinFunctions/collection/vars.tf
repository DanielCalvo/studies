variable "ami_id" {
  type = map
  default = {
    ubuntu   = "ami-06fd8a495a537da8b"
    amazon-linux2 = "ami-0bb3fad3c0286ebd5"
    redhat = "ami-08f4717d06813bf00"
  }
}

variable "mylist" {
  type = list
  default = ["asd", "bsd"]
}
variable "myotherlist" {
  type = list
  default = ["apple", "banana", "orange"]
}
