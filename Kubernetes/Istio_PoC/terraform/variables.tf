variable "vpc_name" {
  description = "name of the vpc"
  default     = "istio-vpc"
}

variable "cluster_name" {
  description = "name of the eks cluster"
  default     = "istio-eks-cluster"
}