variable "settings" {
  type = list(map(string))
  default = [
    {
      name    = "100"
      owner   = "allow"
      purpose = "0"
    },
    {
      name    = "100"
      owner   = "allow"
      purpose = "0"
    },
  ]
}

