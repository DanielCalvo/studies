output "aws_eks_auth_cmd" {
  value = "aws eks --region ${data.aws_region.current.name} update-kubeconfig --name ${var.cluster_name}"
}