## To install

```shell
helm repo add metallb https://metallb.github.io/metallb
helm repo update

helm --kube-context hp \
  upgrade --install metallb metallb/metallb \
  --namespace metallb-system \
  --create-namespace \
  --version 0.16.1
```

## Then apply the address pool

```bash
kubectl --context hp apply -f hp-cluster/metallb/metallb-config.yaml
```

## Then apply the example
```bash
kubectl --context hp apply -f hp-cluster/metallb/example.yaml
```