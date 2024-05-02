## Objectives
1. Set up a VPC and EKS on it using the latest terraform version and terraform modules (when you perform this exercise)
2. Install the Prometheus operators CRDs in this (by hand, for now)
3. Install... the big prometheus manifests from kube-prometheus? Alternatively, install a trimmed down version of it
4. Configure monitoring for the following:
    - A deployment inside k8s (pod number and... readiness endpoint? Figure out how that works!)
    - Node exporter on an EC2 machine somewhere outside of the cluster
    - Some other arbitrary exporter, say, some exporter to list the number of EC2 machines on this account or some other thing.
    - Oh and a blackbox exporter!
5. Configuring the prometheus forwarder thing to send metrics to another prometheus instance would be really cool.

## Other ideas
- What does long term metric storage look like?

---
Haven't worked with Prometheus in a while, and haven't set it up inside k8s. Lets give this a go!

I wanna set this up on AWS. Stack is:
- VPC
- EKS
- Yeah nah lets use minikube. Minikube's cool

Hmmm wait but what about minikuberino?
- Minikuberino
- Prometheus helm chart install
- Figure out if you wanna add more stuff
    - some ingress
    - some gitops

## Things I'd like to know more about
- So I just k8s CRD everything? Or is it all regular yaml? I really can't remember that well

## Commands 
- `minikube start --nodes=2`

## Prometheus operator itself
- The operator is here: https://github.com/prometheus-operator/prometheus-operator
- But there was a repo with instructions on how to set it up better, right?
    - Yeah this provides examples: https://github.com/prometheus-operator/kube-prometheus

## HOLD UP I FOUND THE DOCS:
- https://prometheus-operator.dev/docs/operator/design/

## Copied over from the docs:
Before deploying kube-prometheus in a production environment, read:
- https://github.com/prometheus-operator/kube-prometheus/blob/main/docs/customizing.md
- https://github.com/prometheus-operator/kube-prometheus/tree/main/docs/customizations
- https://github.com/prometheus-operator/kube-prometheus/blob/main/docs/access-ui.md
- https://github.com/prometheus-operator/kube-prometheus/blob/main/docs/troubleshooting.md

# What you got so far
- Install CRDs: https://prometheus-operator.dev/docs/user-guides/getting-started/#installing-the-operator
    - This doesn't include any instance of prometheus, its just the CRDs! 

# A default list of manifests can be found here: 
- https://github.com/prometheus-operator/kube-prometheus/tree/main/manifests
- Looks like this is your starting point, and you add, mix and kustomize the above until you got what you want
- Your next action could be to start with a reduced set of those to set up something

# What about kube-mixing, is there anything valuable in there? 
- https://github.com/kubernetes-monitoring/kubernetes-mixin

# The custom resources managed by the Prometheus Operator are
It could be handy to go one by one in here and reckon what they do -- I think that's the key to sorting this out

## These are the CRDs that you have
- Prometheus
- Alertmanager
- ThanosRuler
- ServiceMonitor
- PodMonitor
- Probe
- PrometheusRule
- AlertmanagerConfig
- PrometheusAgent


## Prometheus
- An instance of Prometheus. You can configure the number of replicas, persistent storage, and Alertmanagers to which this prometheus instance sends alerts.
- An underlying `StatefulSet` is deployed by the operator in the same namespace to run this Prometheus instance
- This CRD defines via label and namespace selectors which `ServiceMonitor`, `PodMonitor` and `Probe` objects should be associated to this Prometheus instance.
    - It also defines which `PrometheusRules` should be reconciled 

TLDR: This represents a Prometheus instance.

## Alertmanager
- Declares an instance of AlertManager to run in this cluster. You can configure the replicas and persistent storage.
- A `StatefulSet` is configured under the hood

## ThanosRuler
- Hang on this one seems advanced, let me do this one last!

## ServiceMonitor
- Allows to declare which dynamic set of services should be monitored. Interesting
- Looks like this is meant to monitor an internal k8s `Service`, hence the `Service` in `ServiceMonitor`

## PodMonitor
- Used to monitor pods. Looks like this monitors the HTTP endpoint of a given pod

## Probe
- Lets you declare how groups of ingresses and static targets can be monitored. Neat!
- Looks like this duplicates the functionality of a blackbox exporter?

## PrometheusRule
- Lets you declare rules for Prometheus or Thanos to consume. ?? not very well explained in the docs

## AlertmanagerConfig
- Config for alertmanager, pretty much

## PrometheusAgent
- Lets you set up a Prometheus Agent. Oh looks like this is for metric forwarding. Neat!

## ScrapeConfig (Bonus!)
- https://prometheus-operator.dev/docs/user-guides/scrapeconfig/
- Used to scrape configs outside of the cluster. Interesting.
- It allows for service discovery of
    - static_config
    - file_sd
    - http_sd
    - kubernetes_sd
    - consul_sd
- Hmm, this looks like top level `scrape_configs:` config entry on a usual Prometheus instance.