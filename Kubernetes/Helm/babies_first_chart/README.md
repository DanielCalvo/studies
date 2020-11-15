- Bitnami's docs are excellent: https://docs.bitnami.com/tutorials/create-your-first-helm-chart/

### Getting to it
- `helm create danichart`
    - Oh wow it creates everything
- `helm install danichart --dry-run --debug ./danichart`
- Chart.yaml contains metadata information for this chart
- Values.yaml contains the values that get templated into the chart
- You can override values in the command line though: `helm install -danichart -dry-run --debug ./danichart --set service.internalPort=8080`
- `helm package ./danichart`
- `helm install example3 danichart-0.1.0.tgz --set service.type=NodePort`

```shell script
cat > ./danichart/requirements.yaml <<EOF
dependencies:
- name: mariadb
version: 0.6.0
repository: https://charts.helm.sh/stable
EOF
```

- Manually added a dependency to Chars.yaml and ran:
- `helm dep update ./danichart`



