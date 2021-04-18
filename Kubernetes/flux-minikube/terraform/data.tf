
# https://github.com/fluxcd/terraform-provider-flux/blob/main/docs/data-sources/install.md
data "flux_install" "main" {
  target_path = var.target_path
  target_path = "../clusters/justmemeing"
}
# https://github.com/fluxcd/terraform-provider-flux/blob/main/docs/data-sources/sync.md
data "flux_sync" "main" {
  target_path = var.target_path
  target_path = "../clusters/justmemeing"
  url         = "ssh://git@github.com/${var.github_owner}/${var.repository_name}.git"
  branch      = var.branch
}

data "kubectl_file_documents" "install" {
  content = data.flux_install.main.content
}

data "kubectl_file_documents" "sync" {
  content = data.flux_sync.main.content
}
data "github_repository" "main" {
  full_name = "${var.github_owner}/${var.repository_name}"
}