
### Get your management cluster & crossplane going

```shell
minikube start
kubectl create namespace crossplane-system
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane --version 1.2.2
```
- kubectl create secret generic aws-creds -n crossplane-system --from-file=creds=/home/daniel/.aws/credentials

### Specific to this 

- kubectl apply -f provider.yaml
- kubectl apply -f providerconfig.yaml
- kubectl apply -f s3.yaml