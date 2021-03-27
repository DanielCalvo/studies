resource "helm_release" "nginx-controller" {
  name       = "quickstart"
  chart      = "ingress-nginx"
  repository = "https://kubernetes.github.io/ingress-nginx"
  version    = "3.24.0"
}
# helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
# helm install quickstart ingress-nginx/ingress-nginx