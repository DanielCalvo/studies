### Taken from here
- https://helm.sh/docs/intro/quickstart/

### Steps
- Download Helm
- Install it
- Initialize a helm chart repository
- `helm repo add stable https://charts.helm.sh/stable`
- `helm repo update`
- `helm install stable/mysql --generate-name`
- `helm show chart stable/mysql`
- `helm uninstall smiling-penguin`
- `helm ls`

### https://helm.sh/docs/intro/using_helm/
- A Chart is a Helm package
- A Repository is the place where charts can be collected and shared
- A Release is an instance of a chart running in a Kubernetes cluster

#### helm search
- `helm search hub wordpress` (searches the artifact hub)
- `helm repo ls`
- `helm search repo`
- `helm search repo wordpress`
- `helm repo add brigade https://brigadecore.github.io/charts`
- `helm install happy-panda stable/mariadb`
    - `helm install $RELEASE_NAME $CHART`

#### Customizing the Chart Before Installing
- `helm show values stable/mariadb`
- `echo '{mariadbUser: user0, mariadbDatabase: user0db}' > config.yaml`
- `helm install -f config.yaml stable/mariadb --generate-name`

#### Upgrading and rolling back
- `helm upgrade -f panda.yaml happy-panda stable/mariadb`
- `helm rollback happy-panda 1`
- `helm uninstall happy-panda`
- `helm list`

#### helm repo
- `helm repo list`
- `helm repo add dev https://example.com/dev-charts`
- `helm repo update`
- `helm repo remove`

#### Creating Your Own Charts
- `helm create dani-workflow`
- `helm package dani-workflow`
- `helm install dani-workflow ./dani-workflow-0.1.0.tgz`
- `helm inspect chart dani-workflow`
- `helm inspect all dani-workflow`

### Chart development