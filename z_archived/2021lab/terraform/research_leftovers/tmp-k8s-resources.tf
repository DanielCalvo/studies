resource "kubernetes_service_account" "external_dns" {
  metadata {
    name = "external-dns"
    annotations = {
      "eks.amazonaws.com/role-arn" = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/${aws_iam_role.k8s_dns_role.name}"
    }
  }
}

