# Prometheus Operator CRDs

Source: https://prometheus-operator.dev/docs/api-reference/api/

## Alertmanager (`alertmanagers.monitoring.coreos.com`)
Deploys an alert manager in your Kubernetes cluster.
- Docs: https://prometheus-operator.dev/docs/api-reference/api/#alertmanager

## AlertmanagerConfig (`alertmanagerconfigs.monitoring.coreos.com`)
Configures an alert manager specifying how alerts should be grouped, inhibited and otherwise notified to external systems.

- Docs: https://prometheus-operator.dev/docs/api-reference/api/#alertmanagerconfig

## PodMonitor (`podmonitors.monitoring.coreos.com`)
Defines how Prometheus and Prometheus agent can scrape metrics from a group of pods. 
 you can configure the label selectors for scraping pods, the container ports to scrape, authentication credentials to use, and target and metric relabeling 
Docs: https://prometheus-operator.dev/docs/api-reference/api/#podmonitor

If you have something like a deployment with a service in front of it that maybe serves traffic to an ingress then you probably want a service monitor. you would use pod monitor for something like a daemon set or a cron job

## Probe (`probes.monitoring.coreos.com`)
A probe defines how to scrape metrics from prober exporters such as the blackbox exporter.

- Docs: https://prometheus-operator.dev/docs/api-reference/api/#probe

## Prometheus (`prometheuses.monitoring.coreos.com`)
Defines a Prometheus instance to run inside the cluster, you can configure a number of replicas, persistent storage, alert managers where firing alerts should be sent, and more settings. so this is your actual Prometheus
- Docs: https://prometheus-operator.dev/docs/api-reference/api/#prometheus

## PrometheusAgent (`prometheusagents.monitoring.coreos.com`)
The Prometheus agent is mainly for scraping targets and sending samples somewhere else with remote write, think like sending metrics to Thanos.
 when using a Prometheus agent you don't store metrics locally you just forward it to Thanos so you don't alert or query anything with the agent

- Docs: https://prometheus-operator.dev/docs/api-reference/api/#prometheusagent

## PrometheusRule (`prometheusrules.monitoring.coreos.com`)
This resource defines alerting and recording rules to be evaluated by Prometheus or Thanos ruler objects.

 so Prometheus and Thanos ruler objects select Prometheus rule objects using labels and namespace selectors. Neat, so it seems that you could have a Prometheus rule that is selected by multiple Prometheus instances. interesting!
- Docs: https://prometheus-operator.dev/docs/api-reference/api/#prometheusrule

## ScrapeConfig (`scrapeconfigs.monitoring.coreos.com`)
Scrape config is when you need to scrape something that doesn't fit cleanly into service monitor, pod monitor or probe.
 like maybe you need to scrape an external target, like a  virtual machine
 or maybe use some service discovery not represented by service monitor like DNS service discovery or Consul service discovery
- Docs: https://prometheus-operator.dev/docs/api-reference/api/#scrapeconfig

## ServiceMonitor (`servicemonitors.monitoring.coreos.com`)
Defines how Prometheus and Prometheus agent can scrape metrics from a group of services.  seems similar to a pod monitor but in this case the target is a service.

If your deployment has three pods a service monitor will discover and scrape these three pod endpoints not just a service virtual IP

- Docs: https://prometheus-operator.dev/docs/api-reference/api/#servicemonitor

## ThanosRuler (`thanosrulers.monitoring.coreos.com`)
Runs rule evaluation against a Thanos query layer. that layer may combine data from multiple Prometheus instances or clusters or object storage or deduplicated high availability replicas.

you use this if Thanos is your global metrics layer or if you want to query metrics from Thanos or if a single Prometheus instance on your cluster doesn't have all the data you want

- Docs: https://prometheus-operator.dev/docs/api-reference/api/#thanosruler
