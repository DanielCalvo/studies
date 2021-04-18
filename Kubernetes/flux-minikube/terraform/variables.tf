variable "github_owner" {
  type    = string
  default = "DanielCalvo"
}

variable "repository_name" {
  type    = string
  default = "studies"
}

//variable "repository_visibility" {
//  type    = string
//  default = "private"
//}

variable "branch" {
  type    = string
  default = "master"
}

variable "target_path" {
  type        = string
  default     = "flux-minikube"
  description = "flux sync target path"
}