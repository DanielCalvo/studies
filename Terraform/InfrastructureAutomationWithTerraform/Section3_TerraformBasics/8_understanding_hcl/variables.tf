variable "myvar" {
  type = string
  default = "hello terraform"
}

variable "mymap"{
  type = map(string)
  default = {
    mykey = "myvalue"
  }
}

variable "mylist"{
  type = list
  default = [1,2,3]
}

