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

---

- KRM: Kubernetes Resource Model
- XRM: Crossplane Resource Model

- XRD: Composite resource definition.
- XRC: Composite resource claim?
- XR: Composite resource

- Composition: Specifies which resources a composite resource will be composed of

In other words:
- XRD defines which infra you can access and how
- Composition: Calls that infra with simple settings?
- XRDs and compositions can be packaged and installed as a configuration

Continue here tomorrow with your coffee: https://crossplane.github.io/docs/v1.2/getting-started/install-configure.html

- Doc suggestion:
- Read everything, but then started without a composition

### Investigation TODOs
- Read the Composition API documentation and try to create one as bare bones as possible
- Read the Definition API documentation and try to create one as bare bones as possible with an S3 bucket
- Figure out how to install that on your clusters (aka a package -- but you don't need to package it and upload it, just install it for now)
- Call that package with parameters and create your S3 bucket with just "name" as parameter, nothing else
- Find out how to create a package that instantiate two or more AWS resources, say an S3 bucket and a policy or something else

### Implementation suggestion
- Create an S3 bucket. You just pass as parameter the name. The region, ACL, policy, etc are all defined by the crossplane package