resource "tls_private_key" "main" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "kubernetes_namespace" "flux_system" {
  metadata {
    name = "flux-system"
  }

  lifecycle {
    ignore_changes = [
      metadata[0].labels,
    ]
  }
}

resource "kubectl_manifest" "install" {
  for_each   = { for v in local.install : lower(join("/", compact([v.data.apiVersion, v.data.kind, lookup(v.data.metadata, "namespace", ""), v.data.metadata.name]))) => v.content }
  depends_on = [kubernetes_namespace.flux_system]
  yaml_body  = each.value
}

resource "kubectl_manifest" "sync" {
  for_each   = { for v in local.sync : lower(join("/", compact([v.data.apiVersion, v.data.kind, lookup(v.data.metadata, "namespace", ""), v.data.metadata.name]))) => v.content }
  depends_on = [kubernetes_namespace.flux_system]
  yaml_body  = each.value
}

resource "kubernetes_secret" "main" {
  depends_on = [kubectl_manifest.install]

  metadata {
    name      = data.flux_sync.main.secret
    namespace = data.flux_sync.main.namespace
  }

  data = {
    identity       = tls_private_key.main.private_key_pem
    "identity.pub" = tls_private_key.main.public_key_pem
    known_hosts    = local.known_hosts
  }
}

# GitHub
//resource "github_repository_deploy_key" "main" {
//  title      = "flux-minikube"
//  repository = "github-actions-shenanigans"
//  key        = tls_private_key.main.public_key_openssh
//  read_only  = true
//}
//
//resource "github_repository_file" "install" {
//  repository = "github-actions-shenanigans"
//  file       = data.flux_install.main.path
//  content    = data.flux_install.main.content
//  branch     = var.branch
//}
//
//resource "github_repository_file" "sync" {
//  repository = "github-actions-shenanigans"
//  file       = data.flux_sync.main.path
//  content    = data.flux_sync.main.content
//  branch     = var.branch
//}
//
//resource "github_repository_file" "kustomize" {
//  repository = "github-actions-shenanigans"
//  file       = data.flux_sync.main.kustomize_path
//  content    = data.flux_sync.main.kustomize_content
//  branch     = var.branch
//}