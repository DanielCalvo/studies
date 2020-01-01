### 41: Manual Scheduling
- You can set a field named nodeName on a pod.
- This is not set by default and you usually don't set this field.
- Kubernetes managed this field automatically for you.
- Managing this field manually allows you to place a pod manually on a node.
- If there is no scheduler available for some reason, the pods will not be scheduled in any node, and stay in a pending state
- You can manually schedule pods yourself. You can only specify the nodeName entry on a pod definition at creation time.
- If the pod is already created and you want to assign it to a node, you have to create Binding object and send a POST request to the pod's binding API, thus mimicking what the scheduler actually does.


### 42: Practice test - Manual Scheduling
Q: Why is the POD in a pending state? Inspect the environment for various kubernetes control plane components.
A: kube scheduler wasn't running apparently


### 43: Labels and selectors
- Labels are way of grouping things together in Kubernetes. Labels are properties attached to objects in Kubernetes
- Labels go under
```yaml
metadata:
  labels:
    key: value
```
- Once a pod or object is created, you can select it with a label by doing:
- `kubectl get pods --selector app=myapp`
- When defining a ReplicaSet (or deployment) you may see labels defined twice.
- On a ReplicaSet, the first label will be of the ReplicaSet itself. On the spec specification is where you'll have labels for pods.
- In other to tie the ReplicaSet to pods, you configure the selector field under the ReplicaSet specification, to match the labels defined on the pod(s). A single label will do, but you can specify as many as you want.
- The same goes for services. A service uses the selector field to match a certain label against existing pods. Pods that match will be part of the service.
- Annotations are for other metadata which are not necessarily labels.
- Annotations are used to store other details for information purposes (buildversion, build information, contact details, something else)

### 44: Practice test - Labels and Selectors
Q: How many pods in the dev environment?
R: `kubectl get pods -l env=dev`

Q: How many objects are the in the prod environment?
R: `kubectl get all -l env=prod`

Q: Identify the POD which is 'prod', part of 'finance' BU and is a 'frontend' tier?
A: kubectl get pod -l env=prod -l bu=finance -l tier=frontend

Q: A ReplicaSet definition file is given 'replicaset-definition-1.yaml'. Try to create the replicaset. There is an issue with the file. Try to fix it.
A:
```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: replicaset-1
spec:
  replicas: 2
  selector:
    matchLabels:
      tier: frontend #<- these need to match
  template:
    metadata:
      labels:
        tier: frontend #<- these need to match
    spec:
      containers:
      - name: nginx
        image: nginx
```


### 46: Taints and Tolerations
- Taints and tolerations are used to set restrictions on what pods can be scheduled on a node
- Let's say we add a taint to a node, the taint "blue". No pods can be scheduled on the "blue" node, as none of them can tolerate the "blue" taint
- To schedule a pod on this node, we need to specify which pods are tolerant of this taint (the "blue" taint)
- You add a toleration to a pod -- pod must be tolerant to the "blue taint" -- so that it can be scheduled on the node tainted blue
- Taints are set on nodes, and tolerations are set on pods
- To taint a node:
- `kubectl taint nodes node-name key=value:taint-effect`
- The taint effect defines what would happen to a pod if they do not tolerate the taint
- There are 3 taint effects:
    - NoSchedule: Pods won't be scheduled on the node
    - PrefferNoSchedule: The system will attempt avoinding placing a pod on the node
    - NoExecute: New pods will not be scheduled on the node, and existing pods on the node will be evicted if they do not tolerate the taint
- Example: `kubectl taint nodes node1 app=blue:NoSchedule`
- See [./46_toleration.yaml]() for a toleration example

#### Taint - NoExecute
- Taints & Tolerations are only meant to restrict nodes from accepting certain pods. 
- Node1 may only accept pod D, but this does not guarantee that pod D will always be placed on node1, since there may be no taints applied to other nodes
- Taints & tolerations does not tell the pod to go to a particular node. It only tells the node to accept pods with certain tolerations.
- If you want to restrict a pod to certain nodes, this is achieved through another concept called "Node Affinity"
- Master nodes have a taint that prevent any nodes from being scheduled on these nodes.
- To see this taint: `kubectl describe node kubemaster | grep Taint`

### 47: Practice test - Taints and tolerations
Q: Do any taints exist on node01?
A: `kubectl describe node01`

Q: Create a taint on node01 with key of 'spray', value of 'mortein' and effect of 'NoSchedule'
A: `kubectl taint nodes node01 spray=mortein:NoSchedule`

Q: Create a taint on node01 with key of 'spray', value of 'mortein' and effect of 'NoSchedule'
A: kubectl run mosquito --image=nginx --restart=Never

Q: Create another pod named 'bee' with the NGINX image, which has a toleration set to the taint Mortein
A:
```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: bee
  name: bee
spec:
  containers:
  - image: nginx
    name: bee
  tolerations:
    - key: "spray"
      operator: "Equal"
      value: "mortein"
      effect: "NoSchedule"
```
Q: Remove the taint on master, which currently has the taint effect of NoSchedule
A: `kubectl taint nodes master node-role.kubernetes.io/master:NoSchedule-`

### 48: Node Selectors
- There are two ways to make a pod run on a particular node:
    - NodeSelectors, making pods run on nodes with a specific label
    - 
- To label a node: `kubectl label nodes $nodename $labelkey=$labelvalue`
- Ex: `kubectl label nodes node-1 size=Large`
- NodeSelector has its uses, but it also has limitations: What if you want to run a pod on a node that is MEDIUM or LARGE (or not SMALL?)

### 49: Node Affinity
- The main feature of node affinity is to make sure certain pods are hosted on certain nodes.
- NodeAffinity rules: `requiredDuringSchedulingIgnoredDuringExecution` and `prefferedDuringSchedulingIgnoredDuringExecution`  
- If you have the `required` rule, the scheduler will mandate that the pod is scheduled on a node with the proper affinity rules. If you can't find one, the pod will not be scheduled
- If you have the `preffered` rule, if a node cannot be found, the scheduler will simply ignore node affinity rules and place the pod on any available node
- `IgnoredDuringExecution` if a label is changed on a node while a pod is already scheduled there, something may happen. In this case nothing. Pods will continue to run
- `RequiredDuringExecution` will evict pods as soon as a label change happens (planned, not yet implemented)

### 50: Practice test - Node Affinity
Q: How many Labels exist on node node01?
A: `kubectl describe nodes node01`

Q: Apply a label color=blue to node node01
A: `kubectl label node node01 color=blue`

Q: Create a new deployment named 'blue' with the NGINX image and 6 replicas
A: `kubectl create deployment blue --image=nginx`
A: `kubectl edit deployment blue` <- make it have 6 replicas

Q: Schedule the nginx deployment on a pod that has the blue label
A:
```yaml
apiVersion: v1
items:
- apiVersion: extensions/v1beta1
  kind: Deployment
  metadata:
    labels:
      app: blue
    name: blue
  spec:
    selector:
      matchLabels:
        app: blue
    template:
      metadata:
        labels:
          app: blue
      spec:
        containers:
        - image: nginx
          imagePullPolicy: Always
          name: nginx
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:
              nodeSelectorTerms:
              - matchExpressions:
                - key: color
                  operator: In
                  values:
                  - blue
```

### 51: Taints and Tolerations vs Node Affinity
- If you want certain pods to run on certain nodes, you can taint the nodes with a value, and then set the pods to tolerate that value
- But taints and tolerations does not guarantee that these pods will always run on these nodes
- Using node affinity you can label the nodes with something, and then set node selector on the pods to tie the pods to the nodes
- But this does not guarantee that other pods are not places on these nodes
- You can use a combination of this to solve your problem:
    - First you use taints and tolerations to prevent other pods from being placed on your nodes
    - Then you use node affinity to make sure your pods are not placed on other nodes

### 52: Resource requirements and limits
- Kubernetes places pods on the nodes with most available resources by default.
- If there are no resources available on a node to place a pod, kubernetes holds back scheduling the pod.
- You will see the pod in a pending state. If you look at the events you'll see the reason.
- By default, Kubernetes defines that a pod or a container within the pod requires 0.5 CPU units and 256 mebibytes of memory. These are known as the "resource request" -- The minimum amount of CPU and memory requested by a container.
- When the scheduler tries to place your application on a node, it will look for a node that has these resources available, minimally.
- You can specify the resources requirements for your pod under the pod specs. See [52_pod_resourcedef.yaml](./52_pod_resourcedef.yaml)

- What does 1 count of CPU mean?
- "0.1" CPU can also be expressed as "100m". m stands for "mili". You can go as low as 1m. A count of 1 CPU is equal to 1 vcpu on whatever undelying platform you have (gcloud, aws, on prem)

- By default, Kubernetes sets a limit of 1 vcpu per container, so if you do not specify it explicitly, a container will be limited to using 1 vcpu when under load.
- The same goes with memory. By default the limit is 512 Mi. These can also be changed, take a look at See 41_pod_resourcedef.yaml.

- Remember that the limits are set for each container within the pod!

- If a pod tries to use resources beyond it's defined limit:
    - In case of CPU: Kubernetes throttles it so it doesn't go beyond the specified limit
    - In case of memory: If a pod tries to consume more memory than it's limit constantly, it will terminate.

### 53: A note on editing pods:
- These are the things you CAN edit on an already existing pod:
    - spec.containers[*].image
    - spec.initContainers[*].image
    - spec.activeDeadlineSeconds
    - spec.tolerations

- You cannot edit resource limits, service accounts and environment variables of a running pod.

- But there are workarounds to this!
- You can do: `kubectl edit pod <podname>`
- Then change whatever you want in vim, save it.
Delete the pod: `kubectl delete pod <podname>`
And then apply the file you saved with whatever modifications you made: `kubectl create -f /tmp/whateverpathtofile.yaml`

- The second workaround is to extract the pod definition: `kubectl get pod mypod -o yaml > mynewpod.yaml`
- `vim mynewpod.yaml` <- edit resource limitations or whaver else you want
- Then delete the existing pod: `kubectl delete pod mypod`
- Then create a new pod with the yaml: `kubectl create -f mynewpod.yaml`

- With deployments, you can edit any field of the pod template, since the deployment will just remove the old pods and spawn new ones with the desired specifications.
- `kubectl edit deployment mydeployment`

### 54: Practice test - Resource requirements and limits
Q: The elephant runs a process that consume 15Mi of memory. Increase the limit of the elephant pod to 20Mi.
A: Save the pod to yaml, delete pod, apply yaml again


### 56: DaemonSets
- DaemonSets are similar to ReplicaSets, the difference being that they run one instance of your pod on each node in your cluster
- Whenever a node is added to the cluster, a copy of your pod on the DaemonSet is added to that node
- When a node is removed, that pod is automatically removed too
- The DaemonSet ensures that one copy of a given pod is always present on all nodes on your cluster
- Use cases of DaemonSets: Monitoring agent (Prometheus Exporter) or Log exporter (Filebeat)
- Kube-proxy actually runs as a DaemonSet in the cluster
- Some networking solutions such as weave-net require an agent to be deployed in each node in the cluster 
- Creating a DaemonSet is very similar to a Replicaset, the only difference is that instead of having Kind: ReplicaSet you have Kind: DaemonSet
- You could emulate the functionality of a DaemonSet by having several copies of your pod definition, each with a different nodeName entry.
- That's how it actually used to be until Kubernetes v1.12 ahahah
- Currently DaemonSet uses nodeAffinity rules to make sure the pods all fall onto different nodes.

### 57: Practive test - Daemonsets
Q: Create a daemonset named fluentd for logging
A: You can generate a deployment, remove the replica number and change the kind to DaemonSet and:
```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    app: fluentd
  name: elasticsearch
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: fluentd
  template:
    metadata:
      labels:
        app: fluentd
    spec:
      containers:
      - image: k8s.gcr.io/fluentd-elasticsearch:1.20
        name: fluentd-elasticsearch
```

### 58: Static Pods
- You can configure the kubelet service to read pod definitions by itself (without etc and kube-api) on /etc/kubernetes/manifests
- Pods created this way are known as Static pods. Pods are the only object that you can create by interacting with the kubelet directly.
- You can list the pods by doing "docker ps", hah! kubectl is unavailable as it connects to the kube-api and that is unavailable too.
- You can still define static pods for Kubelet even when it's part of a cluster.
- If you do kubectl get pods and you have pods created manually through the kubelet with files on the node, they will still show up when you do kubectl get pods.
- The Kubelet creates a mirror object in the kube-api server. They will be read-only through the kube-api
- You can use static pods to deploy the control plane itself (ooooh!)
- Install kubelet on the master nodes and then put the correct pod definitions in there, BAM, Kubernetes.

- The Kubelet doesn't run as a pod. To see where the static pods config are coming from, you have to do a:
- `ps aux | grep kubelet`
- You then need to check inside the config file for kubelet for that definition:
- `vim /var/lib/kubelet/config.yaml`
- Look for staticPodPath, by default it should be /etc/kubernetes/manifests


### 59: Practive Test - Static pods
Q: How many static pods exist in this cluster in all namespaces?
A: Run the command `kubectl get pods --all-namespaces` and look for those with -master appended in the name

### 60: Multiple schedulers
- You can write your own scheduler program if you want, package and deploy it as the default scheduler or as an additional scheduler.
- A K8s cluster can have multiple schedulers simultaneously
- To view events: `kubectl get events`

### 63: Configuring Kubernetes Scheduler
There's further reading material out there for advanced scheduling. Could be an interesting read!