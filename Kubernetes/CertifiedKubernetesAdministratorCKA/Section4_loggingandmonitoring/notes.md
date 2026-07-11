
### 67: Monitor cluster components
- Node level metrics are interesting, as well as pod-level metrics, such as number of pods and their resource usage.
- As of the recording, Kubernetes did not come with a full featured monitoring solution out of the box.
- You can have a thing called "Metrics server" in your Kubernetes cluster.
- It retrieves metrics from the nodes and pods, stores them and aggregates them in memory
- Kubernetes runs a component called Kubelet on each node server, responsible for receiving instructions from the Kubernetes master.
- The Kubelet contains a subcomponent named Container Advisor, or cAdvisor.
- cAdvisor is responsible from retrieving metrics from pods and exposing them through the Kubelet API to make the metrics available for the metrics server.
- You can clone the metrics server from git and then create the pods from the yaml definitions in there.
- https://github.com/kubernetes-sigs/metrics-server
- `kubectl top node`
- `kubectl top pod`

#### Notes from Dani
- Here's an interesting bit of trivia: `kubectl top node` will report usage for the entire node, not only Kubernetes pods
- So if you launch a CPU or memory intensive process in the host outside of Kubernetes, the command will still report the usage of resources including that process, even though it's outside the cluster
- However the `kubectl top pod` command will only show usage that can be attributed to pods

### 68: Practice test - Monitoring 
- Installed metrics server
- `kubectl top node`
- `kubectl top pod`

### 70: Managing application logs
- You can view a stream of logs with:
- `kubectl logs -f podname`
- In case of multiple containers in a pod: `kubectl logs -f mypod mycontainer`

#### Notes from Dani
- Previously if you wanted to get logs from all the pods in a Deployment you had to do something like: `kubectl logs -l app=my-app`, so you would target by label
- As of Kubernetes 1.31 you can do: `kubectl logs deployment/my-app --all-pods` which is really cool!

### 71: Practice test - Monitor application logs
