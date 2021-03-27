resource "helm_release" "cert-manager" {
  name       = "cert-manager"
  chart      = "cert-manager"
  repository = "https://charts.jetstack.io"
  version    = "1.2.0"

  set {
    name  = "namespace"
    value = "cert-manager"
  }
//  set {
//    name  = "version"
//    value = "v1.2.0"
//  }
  set {
    name  = "create-namespace"
    value = "true"
  }
  set {
    name  = "installCRDs"
    value = "true"
  }

}
