### 45. Introducing istioctl
- You can generate your own yaml files and customize your own istio installation using istioctl
- The most common use of istioctl is to install istio on your cluster
- Make sure you have istioctl installed and on your $PATH
- Easiest way to install istio, not recommended: `istioctl install`
- OoOOoo there's an `istioctl experimental` which has many interesting experimental features, and of of those features is "uninstall"

### 46. Istio profiles
- There are a few built in configuration profiles that can be used when installing istio
- The most used/important ones are 
	- default: a "normal" istio install, recommended for production clusters. This is the profile chosen when you do `istioctl install`
	- demo: Designed to showcase all of istio's functionality, comes with a lot of tracing and access logigng out of the box!
	- 
### 47. Installing addons
- The user interface components are now named "integrations" or "addons"
- They're no longer a first party components, but addons you can have by just using some yaml
- So for the control plane dashboard the docs tell you to `kubectl apply -f https://github.com (...) grafana.yaml`

### 48. Tuning profiles
- You can do a `istioctl profile dump demo` and this will dump to yout terminal all the yaml for this profile, which you can then apply or change by hand if you want to
- After tuning this profile by hand, you can: `istioctl install -f myistioconfig.yaml`
- There can be an issue with jwt auth in minikube as it's not properly supported, have a look at the docs if your istio pod is stuck while trying to be created

### 49. (optional) Enabling Third Party Tokens on Kops 
- Eh, some kops specific stuff for third party tokens
- EKS supports this out of the box

### 50. Default vs Demo Profiles - CPU and Memory
- You can adjust if you want HPA on istio, and if you don't you can set how many replicas you want
- You can also adjust the sampling rate for pilot!
- You can also monitor with grafana how many resources istiod is using and adjust the requests accordingly. as these are set to what are mostly sensible defaults, but if your cluster is large, you might need more requests!
- Remember that there are also limits/requests for the sidecar containers that istio uses
- The requests for the demo and default profiles are different: Demo has very low requests as it's mean to run on demo systems (like minukube) and default has much larger requests, as these are expected to run in production clusters
- Try not to make guesses with requests: Look at grafana to see how much CPU and memory your pods are using

### 51. Generating YAML manifests
- When you do `istioctl profile dump demo` this will output yaml that is meant to be used by istioctl, this is not vanilla kubernetes yaml and kubectl will not accept it
- You can run `istioctl manifest generate` to generate normal kubernetes yaml from your istio config, or from a yaml that contains istio config, for example:
- `istioctl manifest generate -f my-istio.yaml > plain-k8s-istio.yaml`
- There are a few downsides to this and the doc doesn't recommend you do it. The istio namespace won't be included in the plain yaml and this plain yaml will also contain a bunch of crds that need to be applied first!
- Advantage of having plain yamls: All the yaml in the cluster can be dumped with kubectl and reapplied as part of a backup thing, you can also have your cluster configs in version control, although there might be better ways to do this (like using flux) 