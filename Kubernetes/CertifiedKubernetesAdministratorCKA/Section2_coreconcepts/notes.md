
### Before moving on
- Recheck lecture 8 with "Resources for this lecture"
- There might be something useful there, the taints & tolerations pdf appeared interesting!

### 9: Cluster architecture
- Cluster consists of a set of nodes. There are worker nodes and master nodes.
- Worker node: Hosts applications as containers
- Master node: Manages, plans, executes and monitors nodes

Cluster components:
- **etcd**: Database that stores information in a key-value format about worker nodes and what containers they have running with additional info is stored in etcd in master
- **kube-scheduler**: Identifies the right node to place a container on based on assorted rules/resources, policies, taints & tolerations, node affinity rules and so on. More lectures on this later 
- Controller managers
    - **node-controller**: Responsible for onboarding new nodes to the cluster, handling unavailable nodes, 
    - **replication-controller**: Ensures that the designed number of containers are running at all times in a replication group
- **kube-apiserver**: Responsible for orchestrating all operations within the cluster, exposes the kubernetes API
- **Container run time engine**: What runes the containers, usually docker, but kubernetes supports others
- **kubelet**: Agent that runs on every node in the cluster, listens from instructions from the kube-api server and handles containers on this node as required. The kube-api server periodically gets status reports from the kubelet to monitor the status of nodes and their containers
- **kube-proxy**: Ensures that the necessary (network) rules are in place on the worker nodes to allow the containers running in them to reach each other

What runs on what:
- Master: ETCD cluster, kube-apiserver, Kube Controller Manager, kube-scheduler
- Worker: kubelet, kube-proxy, Container Runtime Engine

### 10: ETCD for beginners
- ETCD is a distributed key-value store.
- Key value stores store information in the form of document, or pages
- Listens on port 2379 by default

Really easy to run it, you can download a release from github, uncompress it and launch ./etcd without any other dependencies (it's a go binary)
You can then use ./etcdctl to:
./etcdctl set key1 value1
./etcdctl get key1


### 11: ETCD in Kubernetes
- The ETCD datastore in Kubernetes stores: Nodes, Pods, Configs, Secrets, Accounts, Roles, Bindings & Others
- Every information that you get when you type kubectl get comes from ETCD
- Any change done to the cluster is updated to the ETCD server
- If you install k8s from scratch, you launch etcd by hand.
- If you use kube-adm to launch kubernetes, it'll launch etcd as a pod in the kubernetes namespace
- On manual set up, the most important flag is the --advertise-client-urls, which is the address on which etcd listens. This is the URL that should be configured on the kube-api server when it tries to reach etcd
- Kubeadm deploys etcd as a pod

### 12: Kube-API Server
When you run a kubectl command, it reaches for the kube-api server. The Kube api server authenticates the request and then validates it. It then communicates with the ETCD cluster and replies to the kubectl command line tool.
You don't need to use kubectl, you could invoke the API directly

#### Kube API server functionalities:
- Authenticates Users
- Validates requests
- Reads / writes from ETCD. kube-api is the only component that interacts directly with ETCD.
- The scheduler, kubelet and kubectl use the Kube API to perform updates in the cluster.
- kube-api is deployed as a pod if deployed by kubeadm

#### When creating a pod
This is a nice example put forward by Mumshad
1. curl -X POST createPod()
2. kube-api creates a pod object without assigning it to a node
3. Updates etcd
4. Updates user to notify the pod is being created
5. Scheduler monitoring the API server realizes that there is a new pod with no node assigned
6. Scheduler identifies the right node to put the new pod on and communicates that back to kube-api
7. kube-api then updates etcd 
8. kube-api then passes that information to the kubelet in the appropriate worker node
9. kubelet creates the pod on the node and instructs the container runtime engine to deploy the application image
10. Once done, the kubelet updates the status back to the api server
11. kube-api updates the data in the etcd cluster

A similar pattern is followed every time a change is requested. The kube-api is at the center of every single task that needs to be performed to make a change in the cluster

### 13: Kube controller manager:
As the name might indicate, manages other controllers inside the Kubernetes cluster
In Kubernetes terms: A controller is a process that continually monitors the state of various components within the system and works towards bringing the system to a desired state.

- Node controller: Responsible for monitoring the states of the nodes. It takes the necessary actions through the kube API server to make sure pods are allocated to healthy nodes, and removes nodes that do not respond to heartbeats under certain circumstances
- Replication controller: Responsible for monitoring the states of replicasets, ensuring that the set number of pods are available at all times in the set.
- There are many more other controllers. They're all packaged into a single process known as the Kubernetes controller manager.

If you set up your cluster with kubeadm, the kube-controller-manager will be a pod in the kube-system namespace

### 14: Kube scheduler
- Responsible for scheduling pods on nodes. Only responsible for deciding which pod goes on which node. But it doesn't actually place the pod on the node! :o
- The kubelet is the controller that creates the pod on the node. The scheduler only decides which pod goes where.
- The scheduler only decides on which nodes the pods are placed on. Pods can have different requirements, and nodes can have different ammounts of resources vailable.
- The scheduler looks at each pod and tries to find the best node for it. Pods might have certain requirements that only make them suitable to certain nodes (such as high cpu and memory requirements)
- Other factors that factor in that will be seen later:
    - Resource requirements and limits
    - Taints and tolerations 
    - Node selectors/affinity

### 15: Kubelet
- Runs on worker nodes. Registers the worker node in the cluster. When it receives a request to start a pod, it launches it by communicating with the Container runtime engine (Docker in most cases) to pull the required image and run an instance. Monitors the status of the pod and containers in it and reports it to the kube api server.
- Kubeadm does not deploy kubelets. You must always manually install Kubelet.

### 16: Kube proxy
- Inside a cluster, all pods can communicate among themselves. This is accomplished by deploying a pod networking solution to the cluster.
- A POD network is a virtual network that all the pods connect to and it spans across all the nodes in the cluster
- Kube proxy implements things like service IPs, which are virtual IPs. It uses IPTables rules to make sure that the traffic that his a virtual IP assigned to a service actually reaches a pod.

### 17: Recap - Pods
- Containers are incapsulated in a Kubernetes object known as Pod. A pod is a single instance of an application. A pod is the smallest object that you can create in Kubernetes.
- To scale up an application, you bring up new pods. Pods from the same app can be deployed on the same or different nodes.
- You can have multiple containers in a pod. You can have a helper container that needs to scale up with the app. This helper container will scale up and down together with your app, as it's on the same pod. Containers on the same pod can communicate with each other by referring to themselves by localhost. They share the same network space.

Commands in this lecture:
```
kubectl run nginx --image nginx
kubectl get pods
```

###18: Pods with yaml
Four top level fields are always present on a kubernetes definition file:
```yaml
apiVersion:
kind:
metadata:
spec:
```

See [19_redispod.yaml](./19_redispod.yaml)

How many pods exist in the system?
kubectl get pods
0

Create a new pod with the nginx image
kubectl run nginx --image nginx
OK

###19: Demo - PODs with Yaml
- Mumshad goes over yaml basics and pod specification basics

###20: Practice test introduction
- Mumshad goes over how the practice tests and interface are set up

###22: Practice test - Pods
Q: How many PODs exist on the system? in the current(default) namespace

A: `kubectl get pods` 0  

Q: Create a new pod with the NGINX image
A: `kubectl run nginx --image=nginx`

Q: What is the image used to create the new pods?
A: `kubectl describe pod  newpods-4hw5h`

Q: Which nodes are these pods placed on?
A: `kubectl get pods -owide`

Q: How many containers are on the web app pod?
A: `kubectl get pods -owide`

Q: What images are on the web app pod?
A: `kubectl get pods -owide`

Q: What is the state of the container 'agentx' in the pod 'webapp'?
A: `kubectl get pods -owide`

Q: Why do you think the container 'agentx' in pod 'webapp' is in error?
A: `kubectl describe pod  somepod` <- Image pull err, image does not exist

Q: What does the READY column in the output of the 'kubectl get pods' command indicate?
A: Running containers on a given pod

Q: Delete the 'webapp' Pod.
A: `kubectl delete pod webapp`

Q: Create a new pod with the name 'redis' and with the image 'redis123'. Use a pod-definition YAML file. And yes the image name is wrong!
A: `kubectl get pod nginx-6db489d4b7-4kxs2 -oyaml --export > redis.yaml` <- Edit the .yaml until you have the bare basics
A: You can also do it above with the etcd or coreDNS or any other pod running on the system
A: or: `kubectl run pod nginx --image=nginx -o yaml --dry-run`
A: or: `kubectl run nginx --image=nginx --restart=Never --dry-run -oyaml > pod.yaml`

Minimum viable pod:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: redis
spec:
  containers:
    - image: redis123
      name: redis
```

Q: Now fix the image on the pod to 'redis'.
A: `kubectl edit pod redis`

### 24: Recap - ReplicaSets
- A controller is a process that monitors Kubernetes objects and responds accordingly
- There is a controller named the replication controller
- The replication controller allows us to run multiple instances of the same pod in the cluster, providing high availability. (If you have a single pod and it dies, bad luck)
- Even if you have a single pod, a replication controller can be useful by bringing up a new pod of the current one fails
- It also allows you to scale up your number of pod instances up and down and to share the load across them
- A Replication Controller is not a Replica Set. The Replication Controller is older, but the concepts apply to both of these.
- We'll try to stick to replicasets

#### Main differences between ReplicaSets and ReplicationControllers:
- A selector is required for a ReplicaSet, but not for a ReplicationController.
- ReplicaSet uses: `apiVersion: apps/v1`
- ReplicationController uses: `apiVersion: v1`

- Commands
- `kubectl create -f replicaset.yaml`
- `kubectl get replicaset`
- `kubectl get pods`
- `kubectl replace -f replicaset-definition.yaml` <- to update number of replicas after modifying the .yaml definition
- `kubectl scale --replicas=6 -f replicaset-definition.yaml`

### 25: Practice test - ReplicaSets
Q: How many PODs exist on the system on the default namespace?
A: `kubectl get pods`

Q: How many ReplicaSets exist on the system?
A `kubectl get replicasets`

Q: How many PODs are DESIRED in the new-replica-set?
A `kubectl get replicasets`

Q: What is the image used to create the pods in the new-replica-set?
A: `kubectl describe replicaset new-replica-set`

Q: How many PODs are READY in the new-replica-set?
A: `kubectl get replicasets`

Q: Why do you think the PODs are not ready?
A: Image does not exist

Q: Delete any one of the 4 PODs
A: `kubectl delete pod new-replica-set-67v5c`

Q: How many pods exist now?
A: 4

Q: Why are there still 4 PODs, even after you deleted one?
A: A ReplicaSet ensures that the desired number of pods are always running

Q: Create a ReplicaSet using the 'replicaset-definition-1.yaml' file located at /root/
A:
```yaml
apiVersion: apps/v1 #<- Issue was here (originally was apiVersion: v1
kind: ReplicaSet
metadata:
  name: replicaset-1
spec:
  replicas: 2
  selector:
    matchLabels:
      tier: frontend
  template:
    metadata:
      labels:
        tier: frontend
    spec:
      containers:
      - name: nginx
        image: nginx
```

Q: Fix the issue in the replicaset-definition-2.yaml file and create a ReplicaSet using it.
A: 
```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: replicaset-2
spec:
  replicas: 2
  selector:
    matchLabels:
      tier: nginx # <- Selector and label must match, previously they did not
  template:
    metadata:
      labels:
        tier: nginx # <- Selector and label must match, previously they did not
    spec:
      containers:
      - name: nginx
        image: nginx
```

Q: Delete the two newly created ReplicaSets - replicaset-1 and replicaset-2
A: `kubectl delete replicaset-1` `kubectl delete replicaset-2`

Q: Fix the original replica set 'new-replica-set' to use the correct 'busybox' image. Either delete and re-create the ReplicaSet or Update the existing ReplicSet and then delete all PODs, so new ones with the correct image will be created.
A: `kubectl edit replicaset new-replica-set`
A: `kubectl delete pod new-replica-set-8ltws new-replica-set-cbh6q new-replica-set-sfv5j new-replica-set-x47nm`

Q: Scale the ReplicaSet to 5 PODs. Use 'kubectl scale' command or edit the replicaset using 'kubectl edit replicaset'
A: `kubectl edit replicaset new-replica-set` 

Q: Now scale the ReplicaSet down to 2 PODs. Use 'kubectl scale' command or edit the replicaset using 'kubectl edit replicaset'
A: To do it differently: `kubectl scale --replicas=2 replicaset/new-replica-set`

### 27: Deployments
- Deployments allow you to have several updates strategies to your pod images (rolling updates, undo changes, pause and resume changes as required)
- The syntax for a deployment type appears to be exactly the same as the ReplicaSet, the only difference being that it is of type "Deployment"
- A deployment actually creates a ReplicaSet, but it provides a level of abstraction more on top of it, allowing you to have strategies to update your images.
- To see all the objects related to a deployment: `kubectl get all`

### 28: Certification tip
- Using the `kubectl run` command can help you you can generating yaml files
- For example, is asked to create a pod or deployment with a specific name and image, you can simply run `kubectl run`
- Bookmark for the exam: https://kubernetes.io/docs/reference/kubectl/conventions/
Examples
```
kubectl run --generator=run-pod/v1 nginx --image=nginx
kubectl run --generator=run-pod/v1 nginx --image=nginx --dry-run -o yaml
kubectl run --generator=deployment/v1beta1 nginx --image=nginx
kubectl run --generator=deployment/v1beta1 nginx --image=nginx --dry-run -o yaml
kubectl run --generator=deployment/v1beta1 nginx --image=nginx --dry-run --replicas=4 -o yaml
kubectl run --generator=deployment/v1beta1 nginx --image=nginx --dry-run --replicas=4 -o yaml > nginx-deployment.yaml
```

### 29: Practice test - Deployments
Note: Some questions were too easy/covered previously (ex: How many pods/replicasets are in the cluster?) so I've skipped them.
Q: `Create a new Deployment with the below attributes using your own deployment definition file`
A: `kubectl run nginx --image=nginx --replicas=4`

### 30: Namespaces
- There is the default namespace, which is where I assume all your objects go to when you just get a cluster and start launching things. This namespace is created by default.
- There is also the kube-system namespace, which is where Kubernetes run it's own components.
- There is also the kube-public namespace, which is where resources managed by kubernetes but that should be acessible by the public are created
- If you environment is small, you don't really have to worry about namespaces. You can do everything on the default one.
- You can have "dev" and "prod" namespaces, which will isolate the resources between them.
- Namespaces can have policies that define who can do what, as well as resource quotas 
- The apps inside a namespace can reffer to themselves simply by their names. If you have a pod named "mydb" inside your namespace, you chould reach it by simply using the name "mydb" (ex: ping mydb)
- You apps can also reach apps in another namespaces. Imagine you have a dev123 namespace. To connect to mydb on it, you would do:
`ping db-service.dev.svc.cluster.local`

#### Namespaces & commands
- kubectl get pods will only list pods in the default namespace.
- If you wanted to list pods in another namespace, you have to specify it:
- `kubectl get pods --namespace=kube-system`
- To create a pod on another namespace, you can: `kubectl create -f pod-definition.yaml --namespace=dev`
- You can also move the namespace definition to the pod definition file, under the metadata section
- There are namespace definitions:
```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: dev
```
- Or: `kubectl create namespace dev`
- If you want to have `dev` as you default namespace, you can: `kubectl config set-context $(kubectl config current-context) --namespace=dev`
- `kubectl get pods` (will get the pods in the dev namespace)
- Also: `kubectl get pods --all-namespaces`

#### ResourceQuotas
- There are ResourceQuotas to limit resource usage on a given namespace!
- The sample ResourceQuota on the k8s documentation is very large (lots of lines, check it out there if you want to)

### 31: Practice test - Namespaces
Q: How many Namespaces exist on the system?
A: `kubectl get namespaces`

Q: Create a POD in the 'finance' namespace.
A: `kubectl run redis --image=redis --restart=Never --namespace=finance`

Q: Which namespace has the 'blue' pod in it?
A: `kubectl get pods --all-namespaces -owide`
   
   
### 31: Services
- Services are used to enable connectivity between pods and external data sources.
- Services are objects, just like pods and replicasets.

- A NodePort service will listen on a given port on the node, and forward that traffic to a pod.
- A ClusterIP service will create a virtual IP inside the cluster to enable communication between a set of services.
- A LoadBalancer service (???) was explained badly. Implies external connection apparently.

#### NodePort explanation
- The port on the pod in which the application is 80, and it's referred to as "TargetPort"
- The second port is the port on the service itself, and it's simply referred to as "Port"
- There is also the port on the node itself, used to access webserver internally, referred to as "NodePort". The valid port range for a NodePort is between 30000 to 32767
- For an example, see [31_nodeport.yaml](./31_nodeport.yaml)
- You could also manuallly reach your NodePort service by reaching a node on the NodePort, as in: `$NODE_IP:$nodePort`
- Out of these, the only mandatory field is "Port".
- If you do not provide a targetPort, it is assumed to be the same as Port.
- If you do not provide a nodePort, a free one in the valid range (30000 - 32767) will be automatically alocated.
- Note on the yaml definition that ports is an array. You can have multiple port mappings in the same service.

#### ClusterIP explanation
- The service create VirtualIP inside the cluster to enable communication between different services

####LoadBalancer explanation
- Provides a LoadBalancer for your application in supported cloud providers

### 33: Services: ClusterIP
- Pods have an IP adress assigned to them, but when you have a deployment handling those pods, these IPs change very often, as pods are ephemeral.
- You cannot rely on these IP addresses for internal communication.
- A K8s server can group the pods together with a Virtual IP and provide a single "interface" to reach these pods.
- ClusterIP is the default type on Kubernetes services (if you don't specify a service type on your yaml, you get a cluster IP)
- On the yaml spec: The target port where the backend is exposed in the pods. The port is which port in the service this ports gets exposed.
- See [32_clusterip.yaml](./32_clusterip.yaml)

#### 34: Practive test - Services 
Q: How many Services exist on the system? in the current(default) namespace
A: `kubectl get services`

Q: What is the 'targetPort' configured on the 'kubernetes' service?
A: `kubectl describe service kubernetes`

Q: How many labels are configured on the 'kubernetes' service?
A: `kubectl describe service kubernetes`

Q: How many Endpoints are attached on the 'kubernetes' service?
A: `kubectl describe service kubernetes`

Q: Create a new service to access the web application using the service-definition-1.yaml file
A:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: webapp-service
spec:
  type: NodePort
  ports:
    - targetPort: 8080
      port: 8080
      nodePort: 30080
  selector:
    name: simple-webapp
```
### 35: Certification tips: Imperative commands with kubectl
#### Pods
- Familiarize yourself with the `--dry-run` and `-o yaml` options
- Create an nginx pod: `kubectl run --generator=run-pod/v1 nginx --image=nginx`
- Generate a pod manifest, but don't launch it: `kubectl run --generator=run-pod/v1 nginx --image=nginx`
#### Deployments
- Create a deployment: `kubectl run --generator=deployment/v1beta1 nginx --image=nginx`
- Create a deployment the newer, recommended way: `kubectl create deployment --image=nginx nginx`
- Generate a deployment yaml file: `kubectl run --generator=deployment/v1beta1 nginx --image=nginx --dry-run -o yaml`
- Generate a deployment yaml file, the newer way: `kubectl create deployment --image=nginx nginx --dry-run -o yaml`
- Generate a deployment yaml file with 5 replicas: `kubectl run --generator=deployment/v1beta1 nginx --image=nginx --dry-run --replicas=4 -o yaml`
- `kubectl create deployment` does not have a replicas option, you have to create it and then scale up
#### Services
- Create a Service named redis-service of type ClusterIP to expose pod redis on port 6379: `kubectl expose pod redis --port=6379 --name redis-service --dry-run-yaml`
- The above will automatically use the pod's labels as selectors
- Or: `kubectl create service clusterip redis --tcp=6379:6379 --dry-run -o yaml`
- The above will assume `app=redis` as selectors. You cannot pass selectors as an option
- Create a Service named nginx of type NodePort to expose pod nginx's port 80 on port 30080 on the nodes: `kubectl expose pod nginx --port=80 --name nginx-service --dry-run -o yaml`

###36: Practive test - Imperative commands
Q: Deploy a pod named nginx-pod using the nginx:alpine image.
A: `kubectl run nginx-pod --restart=Never --image=nginx:alpine`
  
Q: Deploy a redis pod using the redis:alpine image with the labels set to tier=db.
A: `kubectl run redis --restart=Never --image=redis:alpine --labels=tier=db`

Q: Create a service redis-service to expose the redis application within the cluster on port 6379.
A: `kubectl create service clusterip redis-service --tcp=6379:6379`
A: `kubectl edit service redis-service` <- change the selector 

Q: Create a deployment named webapp using the image kodekloud/webapp-color with 3 replicas
A: `kubectl create deployment webapp --image=kodekloud/webapp-color`
A: `kubectl edit deployment webapp` <- change replicas to 3
A: or: `kubectl scale deployment/webapp --replicas=3`

Q: Expose the webapp as service webapp-service application on port 30082 on the nodes on the cluster
A: `kubectl expose deployment webapp --type=NodePort --port=8080 --name=webapp-service --dry-run -o yaml > webapp-service.yam` <- Then edit the file and change the NodePort

---
Old notes:

Note: When doing kubectl get services, the port that you will see in the PORT(S) column is the service port, NOT the targetPort.
You can see the targetPort by describing the service.


32: Certification Tip: Imperative commands with kubectl
--dry-run: This will not create a service for you, but will help you validate the .yaml file.
-o yaml: This will output the resource definition of your object as yaml in the screen

To generate a config for a pod:
kubectl run --generator=run-pod/v1 nginx --image=nginx --dry-run -o yaml

To create a deployment:
kubectl run --generator=deployment/v1beta1 nginx --image=nginx (old way)
kubectl create deployment --image=nginx nginx (newer way)

To generate yamls for deployments:
kubectl run --generator=deployment/v1beta1 nginx --image=nginx --dry-run -o yaml
kubectl create deployment --image=nginx nginx --dry-run -o yaml

kubectl run --generator=deployment/v1beta1 nginx --image=nginx --dry-run --replicas=4 -o yaml

You can then save the definition to a file:
kubectl run --generator=deployment/v1beta1 nginx --image=nginx --dry-run --replicas=4 -o yaml > nginx-deployment.yaml

To generate service definitions:
kubectl expose pod redis --port=6379 --name redis-service --dry-run -o yaml


kubectl create service clusterip redis --tcp=6379:6379 --dry-run -o yaml
kubectl create service nodeport nginx --tcp=80:80 --node-port=30080 --dry-run -o yaml

(Follow up from Dani: You should definitely follow up on the creation of yaml definitions from the command line. It could be very useful! Note that for some things you need to run run (pod) and for others you need to do expose (service)
follow up on imperative commands for all covered object types for the certification)


33: Practice test: Imperative commands
Very useful: https://kubernetes.io/docs/reference/kubectl/conventions/

Deploy a nginx pod using the nginx:alpine image
kubectl run nginx-pod --image=nginx:alpine --generator=run-pod/v1

Create a redis pod with the label tier = db
kubectl run redis --image=redis:alpine --generator=run-pod/v1 --labels=tier=db

Create a service redis-service to expose the redis application within the cluster on port 6379.
kubectl expose pod redis --port=6379 --name redis-service
The above creates a ClusterIP service. Nice! :D

Create a deployment named webapp using the image kodekloud/webapp-color with 3 replicas
kubectl create deployment webapp --image=kodekloud/webapp-color #But how do you scale it?
kubectl run --generator=deployment/v1beta1 webapp --image=kodekloud/webapp-color --replicas=3 #Worked this way!

Expose the webapp as service webapp-service application on port 30082 on the nodes on the cluster. The web application listens on port 8080
Meh this one sucked. It told me to create a .yaml with an imperative command and then edit it. That defeats the purpose of using imperative commands!