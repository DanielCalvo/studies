1. `cd terraform && terraform init && terraform apply`

2. `aws eks --region eu-west-1 update-kubeconfig --name ing-cert-dns-cluster`

3.
```shell
kubectl apply -f external-dns/external-dns.yaml
```

4.
- https://kubernetes.github.io/ingress-nginx/deploy/#using-helm
```shell
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx
```

5. 
- https://cert-manager.io/docs/installation/kubernetes/
```shell
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.2.0/cert-manager.yaml
kubectl apply -f cert-manager
```

6. `kubectl apply -f hello-kubernetes`