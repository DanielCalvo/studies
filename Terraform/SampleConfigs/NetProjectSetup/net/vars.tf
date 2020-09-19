variable "environment-name" {
  description = "Name of the environment to be used in this VPC"
  type        = string
  default     = "experimental"
}

variable "region" {
  description = "Region that this environment will be on"
  type        = string
  default     = "eu-west-1"
}

//variable "tags" {
//  description = "A map of tags to add to all resources"
//  type        = map(string)
//  default     = {}
//}