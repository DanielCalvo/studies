
- Helm is not a configuration management tool
    - Those are: Helmfile, flux, Reckoner

- A common theme with (good) helm charts: 
- You can take the same chart and release it with minimal configuration input on your local development
- You can then release that same chart with a sophisticated configuration set up in a production environment

- YAML is a superset of JSON, I didn't know that! :o

## Chapter 2: Using Helm

Alrite let's go

```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo list
helm search repo drupal
helm search repo whatever
helm search repo drupal --versions
---
helm install mysite bitnami/drupal #You cannot install multiple charts with the same name
kubectl create namespace one
helm install --namespace one mysite bitnami/drupal
kubectl create namespace two
helm install --namespace two mysite bitnami/drupal
```

Lets get some values going
```shell
cat << EOF > /tmp/values.yaml #It is suggested to check this file into version control
drupalUsername: admin
drupalEmail: admin@example.com
EOF

helm uninstall mysite
helm install mysite bitnami/drupal --values /tmp/values.yaml

helm uninstall mysite
helm install mysitea bitnami/drupal --set drupalUsername=admin --set drupalEmail=admin@example.com
```

You can use a dotted notation with `--set` but this gets very verbose very quickly (ex: `--set mariadb.db.name=mydb`)

Helm maintainers suggest using storing everything in a file (usually values.yaml, but name it whatever you want) as opposed to using --set

### Useful commands

```shell
helm list
helm list --all-namespaces
```

A helm installation is an instance of a chart in your cluster.
A release is a combination of a particular configuration and chart version
When you upgrade a chart, you're creating a new release of the same installation

To upgrade to the latest version of a helm chart

```shell
helm repo update
helm upgrade mysite bitnami/drupal --version 6.2.22
helm search repo drupal --versions
helm upgrade mysite bitnami/drupal --version 7.0.0
```

Configuration gets applied freshly on each release! 
In other words: When upgrading, always pass values.yaml
Helm stores some state, so you can do:

```shell
helm upgrade mysite bitnami/drupal --reuse-values
```


Although it is really recommended that you store your state elsewhere (ex: values.yaml)

Uninstalling a chart
```shell
helm uninstall mychart
helm uninstall 
```

Deletion can take several minutes on large applications

### How helm stores release information
Helm seems to store (all of?) it's state as Kubernetes secrets
- `kubectl get secrets`
When you uninstall a chart, it will also delete the secrets related to keeping state related to that chart. You can keep the metadata related to the chart by doing:
- `helm uninstall --keep-history
  
Don't forget: Charts, installations and releases!`