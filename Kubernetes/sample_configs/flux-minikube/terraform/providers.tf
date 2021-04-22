provider "kubernetes" {
  config_path = "~/.kube/config"
}

provider "flux" {}

provider "github" {
  owner = var.github_owner
  token = local.secret_token
}