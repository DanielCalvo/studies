
### Explaining this shit on terraform
- there's the data source
- there's the data sync
- there's the github stuff

minikube + flux install!

Let's start a chonky minikube
- minikube start --cpus=8 --memory=16384

1. Install methods
    - 1: Use terraform
    - 2: Use Kustomize? 
2

### Install methods

#### Terraform
```shell
cd ./terraform
terraform015 init && terraform015 apply
```

#### Command line bootstrap
```shell
flux bootstrap github \
  --owner=DanielCalvo \
  --repository=github-actions-shenanigans \
  --branch=main \
  --path=./clusters/my-cluster \
  --personal
```

#### Without git
```shell
flux install #one way to install it
kubectl apply -f https://github.com/fluxcd/flux2/releases/latest/download/install.yaml #another way to install it
flux create whateveryouwant
kubectl apply -f whateveryouwant.yaml
```

### Resetting minikube
minikube delete && minikube start --cpus=8 --memory=16384