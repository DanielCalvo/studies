## Monitoring notes
- it would be interesting to alert on a batch job failure, like say if the job for some reason wasn-t triggered I would suppose metrics would be absent as they would be normally pushed to push gateway.  this is a bit loosh here but it would be an idea for an implementation leader
- Well moving forward I think you would be gold to create an app to test and learn and dashboard everything so we would have a practical implementation of everything we read on the Prometheus book and elsewhere
- You should also try out that pre calculated metric thing from Prometheus to avoid generating it in real time.
- I know this question is really out there, but how do you delete a metric from Prometheus? like let-s assume you ingested something with high cardinality by accident and it's causing performance issues and now you want to get rid of it
- It would be cool to use the text filed collector thing two fetch some metrics for a given machine being just for fun so I could be familiar with it
-  you should give the Prometheus metric checking two ago it's call prom tool,  it-s something useful to have on a ci process if you are right and exporter actually
- What are the differences between the open metrics format and the Prometheus format? they look very similar but alleged leather are differences

Also, don"t forget to install grafana!



## From 06-jul session with gepeto
- Check which other metrics you wanna import like kubelet/cAdvisor/metrics-server metrics, also dashboard them and have a brief look! You wrote about this on [sources](ai_kubernetes-metrics-sources.md)
- Explore longer term metrics retention, maybe you can leverage some storage operator or somehthing reading from your other server, or maybe s3, though both things are detours
danb how about synthetic web tests?
- Dont forget to look up what to monitor in a cluster by defaunlt (pod restars? pod unabailability? failing pods? cluster capacity? do some reseach!
- How can i have a hosts dashboard like i have in datadog? Try to moditfy the node exporter dashboard real quick for that so you can select multiple nodes!
- Have a read at your DD/open source feature equivalents and see if theres anything else you wanna explore (grafana k6 seemed interesting)
- Explore some basic cluster health alerts, like etcd/api replica warnings or something so you can familiarize yourself with alerting! maybe send the alerts somewhere!
- Also you have 2 nodes, i think you need some prometheus redundancy in case you wanna alert if one of them ever goaes down, add some alerting too!
- Rememer to writo your dummy app that resizes images or something and export some metrics from it, and maybe make it have a batch job
- how does prometheus ingest and work with traces again? thats a puzzling one!