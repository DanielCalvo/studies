provider "aws" {
  profile = "default"
  region  = "eu-west-1"
}

variable "settings" {
  type = list(map(string))
  default = [
    {
      name    = "1001"
      owner   = "allow"
      purpose = "0"
    },
    {
      name    = "100"
      owner   = "allowwwww"
      purpose = "0"
    },
  ]
}

resource "aws_instance" "example" {
  ami           = "ami-07d9160fa81ccffb5"
  instance_type = "t2.micro"


  dynamic "tags" {
    for_each = var.settings
    content {
      tags = {
        Environment = var.settings.value["name"]
      }
    }
  }
}



