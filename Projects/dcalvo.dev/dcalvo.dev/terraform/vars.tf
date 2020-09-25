variable "domain-name" {
    type = string
    default = "dcalvo.dev"
}

variable "bucket-name" {
    type = string
    default = "dcalvo-dev-bucket"
}

variable "populate-s3-bucket-cmd" {
    type = string
    default = "echo banana"
}