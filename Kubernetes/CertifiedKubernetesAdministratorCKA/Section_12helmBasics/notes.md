### 250. Helm - Introduction

### 251. Installation and Configuration

### 252. Lab: Installing Helm

### 253. A Quick Note on Helm2 vs Helm3

### 254. Helm Components
- When a Helm chart is installed in your cluster, a release is created. An installed chart can have multiple revisions
- You can search and find and download publicly available charts and you can search for them in a repository
- Chart metadata is saved in a cluster as a secret
- So Helm will always know what it did in the cluster because it has all its metadata saved in the cluster itself and not on your local for instance

Well here the course then goes over templating and how the templates look inside the chart, and then they show us the values file, old stuff I'm already familiar with

When a chart is installed in your cluster, a release is created

Most charts are available and described in Artifact Hub, apparently this is a Helm chart repository site listing thing

### 255. Helm Charts
The templating on the chart is templated using the values from the `values.yaml` file for the Helm chart
Apart from the values file, every chart has a `Chart.yaml` file, which contains information about the chart itself
Charts can have dependencies too, like the WordPress chart depends on the MariaDB chart

So a chart has the following directory structure:
- `templates/`: folder containing the YAML to be templated
- `values.yaml`: configurable values for this chart
- `Chart.yaml`: information about the chart itself
- `charts/`: a folder containing dependency charts
- And then you can also have a license file and a README file

### 256. Working With Helm - Basics
- The Helm command line has a bunch of help content using `--help`
- You can then ask something like `helm repo --help` to see help for the repo part

You can use the `helm search` command to search in a hub or repo for a given package, sorry chart

- To install a chart, first you add the repo with `helm repo add`, and then to install the chart you usually do `helm install my-release` followed by the chart name
- `helm list` will list installed charts
- `helm uninstall chart-name` will uninstall the chart for you

There's also the `helm repo` family of commands
- `helm repo list` lists repositories
- `helm repo update`: updates locally saved repository information

### 257. Customizing Chart Parameters
- You probably don't want to install a Helm chart with all its default values
- You can set values in the command line by doing `helm install --set` and then pass the arguments you want. You can then pass this multiple times to customize multiple parameters
- Now of course if you set too many values on the command line it's going to get unwieldy very fast, so what you can do is you can move these to a `values.yaml` file

Now if you want to modify the `values.yaml` file of the chart itself, and I mean not pass a parameter but actually modify it, you need to pull the chart with a `helm pull` command and then you need to decompress it. And then you can run `helm install` with that folder on your file system containing the Helm chart as an argument

### 258. Lab: Using Helm to Deploy a chart

### 259. Lifecycle Management With Helm
- Helm releases can be managed independently, even if you have multiple instances of the same chart installed
- You use the `helm upgrade` command to upgrade an installed chart to a newer version
- `helm list` will list the current releases for you
- You can run the `helm history` command with a specific installed chart to see its release history
- You can also roll back a Helm release by using the `helm rollback` command
- However under the hood it doesn't actually go back to revision one, it goes to a revision three, with a similar configuration as to what revision one had
- Some charts when you upgrade you can do backups using a thing called chart hooks so for instance if your chart has a sub chart dependency on a database, but that's not covered in the course

### 260. Lab: Upgrading a Helm Chart
