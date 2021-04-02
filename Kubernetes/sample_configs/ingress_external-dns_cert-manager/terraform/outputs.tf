output "cluster_auth_command" {
  value       = "aws eks --region eu-west-1 update-kubeconfig --name ${var.cluster_name} "
  description = "Command to run to authenticate to this cluster as a human user"
}
