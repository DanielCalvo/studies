## Building a chart
- Book starts with  `helm create chart anvil`

- Chart.yaml: Metadata and some functionality controls (?) for the chart
- charts/: Dependent charts. Empty for now
- templates/: Templates to generate the Kubernetes manifests
- NOTES.txt: A special template. When the chart is installed, NOTES.txt is templated and displayed, but not installed into the cluster
- tests/: Contain tests that are used by the helm test command
- values.yaml: Default values passed to the templated k8s objects. These values can be overriden

Installing the anvil chart above: 
- `helm install myapp anvil`
- `helm delete myapp`

On Chart.yaml:
```yaml
apiVersion: v2 #Api version of helm. v2 is for helm 3
name: anvil #Chart name
version: 0.1.0 #Version of this chart. Versions are used to order charts by Helm
appVersion: 1.0 #Version of the app. Must be unique
```

Book then goes over some of the fields in values.yaml but nothing of too much depth is discussed

### Packaging a chart
When packaging a chart, there are some useful flags: --dependency-update, --destination, --app-version and --version

### Linting charts
- `helm lint` can help you catch issues during chart development