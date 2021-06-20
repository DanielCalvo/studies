- Crossplane.yaml: The metadata about the configuration
- Definition.yaml: The XRD
- Composition.yaml: The composition

### Install crossplane
```shell
minikube delete && minikube start
kubectl create namespace crossplane-system
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane --version 1.2.2
```

### Configure the AWS provider
```shell
kubectl create secret generic aws-creds -n crossplane-system --from-file=creds=/home/daniel/.aws/credentials
kubectl apply -f Provider.yaml
kubectl apply -f ProviderConfig.yaml
kubectl describe providerconfig.aws.crossplane.io/default
kubectl describe providerconfig default
```

### Verifying the above
```shell
helm list -n crossplane-system
kubectl get all -n crossplane-system
```

### Uh I still need to process these

--- Before all of this, the provider config needs to be ready!

CustomResourceDefinition:
- kubectl describe CompositeResourceDefinition compositepostgresqlinstances.database.example.org

Composition:
- kubectl describe composition compositepostgresqlinstances.aws.database.example.org

Claim:
- kubectl describe postgresqlinstance my-db