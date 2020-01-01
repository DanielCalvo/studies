
### 67: Monitor cluster components
- Node level metrics are interesting, as well as pod-level metrics, such as number of pods and their resource usage.
- As of 18-Oct-2018, Kubernetes did not come with a full featured monitoring solution out of the box.
- You can have a thing called "Metrics server" in your Kubernetes cluster.
- It retrieves metrics from the nodes and pods, stores them and aggregates them in memory
- Kubernetes runs a component called Kubelet on each node server, responsible for receiving instructions from the Kubernetes master.
- The Kubelet contains a subcomponent named Container Advisor, or cAdvisor.
- cAdvisor is responsible from retrieving metrics from pods and exposing them through the Kubelet API to make the metrics available for the metrics server.
- You can clone the metrics server from git and then create the pods from the yaml definitions in there.
- https://github.com/kubernetes-sigs/metrics-server
- `kubectl top node`
- `kubectl top pod`

### 68: Practice test - Monitoring 
- Installed metrics server
- `kubectl top node`
- `kubectl top pod`

### 70: Managing application logs
- You can view a stream of logs with:
- `kubectl logs -f podname`
- In case of multiple containers in a pod: `kubectl logs -f mypod mycontainer`

### 71: Practice test - Monitor application logs