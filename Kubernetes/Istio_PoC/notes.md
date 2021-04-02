

---
- I was unable to finish this in time and the external-dns part does not work. Needs redoing
---


1. `cd terraform && terraform init && terraform apply`

2. `aws eks --region eu-west-1 update-kubeconfig --name sample-eks-cluster`
	- Or whatever the name of the cluster is

3. `istioctl install --set profile=default -y`

4. `kubectl label namespace default istio-injection=enabled`

5. `kubectl apply -f k8s`

### Istio docs
- https://istio.io/latest/docs/concepts/traffic-management/
- https://istio.io/latest/docs/tasks/traffic-management/ingress/ingress-control/

### cert-manager setup
- https://cert-manager.io/docs/installation/kubernetes/
1. `kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.2.0/cert-manager.yaml`

- https://cert-manager.io/docs/configuration/
- https://cert-manager.io/docs/configuration/acme/dns01/route53/
- cert-manager's documentation is difficult -- you really gotta read between the lines and infer things!

### End goal
- Configure istio + cert-manager + external-dns together so you can have DNS and TLS for the domain names you configure on virtualservices/ingress-gateways
- I got Istio and cert-manager to work, but external-dns bamboozled me with it's RBAC permissions and unclear error messages

### External DNS
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install external-dns bitnami/external-dns -f external-dns-helm/args.yaml
```

### Out of place nodes
- Could it be that the ClusterRole for helm for external-dns is breaking things somehow?