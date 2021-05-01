//resource "kubernetes_deployment" "hello_kubernetes" {
//  metadata {
//    name = "hello-kubernetes"
//  }
//
//  spec {
//    selector {
//      match_labels = {
//        name = "hello-kubernetes"
//      }
//    }
//    template {
//      metadata {
//        labels = {
//          name = "hello-kubernetes"
//        }
//      }
//      spec {
//        container {
//          name  = "app"
//          image = "paulbouwer/hello-kubernetes:1.8"
//          port {
//            container_port = 8080
//          }
//        }
//      }
//    }
//  }
//}
//
//resource "kubernetes_service" "hello_kubernetes" {
//  metadata {
//    name = "hello-kubernetes"
//  }
//
//  spec {
//    port {
//      port        = 80
//      target_port = "8080"
//    }
//
//    selector = {
//      name = "hello-kubernetes"
//    }
//
//    type = "NodePort"
//  }
//}
//resource "kubernetes_ingress" "hello_kubernetes" {
//  metadata {
//    name = "hello-kubernetes"
//    annotations = {
//      "alb.ingress.kubernetes.io/scheme" = "internet-facing"
//      "kubernetes.io/ingress.class" = "alb"
//    }
//  }
//  spec {
//    rule {
//      http {
//        path {
//          path = "/"
//          backend {
//            service_name = "hello-kubernetes"
//            service_port = "80"
//          }
//        }
//      }
//    }
//  }
//}