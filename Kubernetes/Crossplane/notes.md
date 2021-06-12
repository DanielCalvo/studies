Let's start here: https://crossplane.io/docs/v1.2/getting-started/install-configure.html
- This seems fun!

```shell
minikube start
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane --version 1.2.2
```

### Random notes
- Looks like there are composite resources (XRs) and compositions
- crossplane.yaml: Metadata about the configuration
- definition.yaml: The XRD (Composite resource definition)
- What you wanna accomplish today:
- EC2 instance with keypair and security group with 22 open to the internet so you can ssh to it

- Managed resources: Things like S3 buckets and EC2 instances

- XR: Composite resources
- A composite resource is a special kind of custom resource that is composed of other resources

- A composition specified how crossplane should reconcile a composite infrastructure resource