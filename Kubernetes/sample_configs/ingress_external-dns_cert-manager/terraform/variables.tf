variable "vpc_name" {
  description = "name of the vpc"
  default     = "ing-cert-dns-vpc"
}

variable "cluster_name" {
  description = "name of the eks cluster"
  default     = "ing-cert-dns-cluster"
}