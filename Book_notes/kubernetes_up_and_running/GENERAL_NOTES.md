
### CHAPTER 1: Introductions

#### How Kubernetes can improve things:

- **Velocity**: It's fast and you can deploy a bunch of things fast without breaking everything
- **Immutable**:
- **Declarative configuration**:
- **Self healing**:
- **Immutable Infrastructure**: 
- **Scaling**:
- **Abstracting your infrastructure**:
- **Efficiency**:

### CHAPTER 2: Creating and Running containers

- **Container Images** Docker container image is the de facto standard. Images are made of layers.
- **Building Application Images with Docker**: Docker build, etc etc
- **Image Security**: Never build images with passwords, even if you remove them. These passwords will still be present on the Docker image layers.
- **Optimizing image size**: Order your layers from least likely to change to most likely to change
- **Storing images in a Remote Registry**: docker tag & docker push
- **Docker Container Runtime**: docker run
- **Limiting resource usage**:
```
docker run -d --name kuard \
--publish 8080:8080 \
--memory 200m \
--memory-swap 1G \
--cpu-shares 1024 \
gcr.io/kuar-demo/kuard-amd64:1
```
Google has a container registry of it's own: GCR (Google Countainer Registry)  
Oh, docker-gc is a garbage collector for Docker!  
- **Summary**: Docker images are easy to use and pretty cool!

### CHAPTER 3: Deploying a Kubernetes Cluster

#### Installing Kubernetes on a Public Cloud Provider

I'll do it on Google Cloud because that's what all the cool kids are using.  
Login into Google Cloud, launch their shell and run:
```
gcloud config set compute/zone us-west1-a
gcloud container clusters create danitest-cluster
gcloud auth application-default login
```
#### Checking Cluster Status
```
kubectl version
kubectl get componentstatuses
```
- **Controller Manager**:
- **Scheduler**:
- **etcd**:

```
kubectl get nodes
kubectl describe nodes
```
The above will give you information about nodes / a given node.

#### Cluster Components
All these run in the kube-system namespace.

- **Kubernetes Proxy**:
- **Kubernetes DNS**:
- **Kubernetes UI**:

### CHAPTER 4: Common kubectl Commands
- **Namespaces**:
- **Contexts**:

#### Viewing Kubernetes API Objects

Each object in Kubernetes exists at a unique HTTP path! Like:  
https://your-k8s.com/api/v1/namespaces/default/pods/my-pod  

The most basic command for viewing Kubernetes objects is through `kubectl get`, you can pass the `-o wide`, `-o json` and `-o yaml`flags to it
You can also use JSONPath to extract certain atributes using the `kubectl get` command:
```
kubectl get pods my-pod -o jsonpath --template={.status.podIP}
```
#### Creating, Updating, and Destroying Kubernetes Objects
In a nutshell:
```
kubectl apply -f obj.yaml
kubectl delete -f obj.yaml
```
#### Labels and anotations
Well looks like you can label things. Nice. Skipping this one for now.

#### Debugging Commands
```
kubectl logs <pod-name> #Has a few interesting flags, like -c and -f
kubectl exec -it <pod-name> -- bash
kubectl cp <pod-name>:/path/to/remote/file /path/to/local/file
```
Oh and don't forget:
```
kubectl help
kubectl help <command-name>
```
### CHAPTER 5: Pods

Now before you go any further, this is the first part in which you actually run things. I'm using Google cloud to try things out, so you need that:
```
#From: https://cloud.google.com/sdk/docs/quickstarts
export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)"
echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
apt-get update && apt-get install google-cloud-sdk


    #Start here again tomorrow!


gcloud init
apt-get install kubectl
```
- **Pods**: The smallest unit of anything in Kubernetes. A group of containers.
All containers on a pod are always on the same node. They share the same IP Adress, hostname and port space.
Containers in different pods are strongly isolated.  
The question to ask your self when creating pods: "Will these containers work correctly if they land on different machines?" If the answer is no, put them together in a pod.

#### Creating a pod imperatively: 


Containers running on different pods are strongly separated
What 


### CHAPTER 6: Labels and Annotations

Labels: Labels are pretty cool, they allow you to define characteristics to a given object.
Much like grains in salt stack.

"Labels are used to identify and optionally group objects in a Kubernetes cluster"
(environment, app version, deploy date, anything really)

"Anotations provide object scoped key/value storage of metadata"


### CHAPTER 7: Service Discovery

Chapter 7: Service Discovery
Confusing as fuck

### CHAPTER 8: ReplicaSets
ReplicaSet: A clusterwide pod manager, ensuring that the right types and numbers of pods are running at all times.
ReplicaSets use labels to identify the set of pods they should be managing
It's a pod manager, man.
Some people even default to replicasets instead of pods.
Good for stateless applications.

### CHAPTER 9: DaemonSets
Chapter 9: DaemonSets  
ReplicaSets are for creating a service with multiple replicas for redundancy.  
DaemonSets are used to deploy system daemons such as log collects and monitoring agents, that must be present in every node.  
ReplicaSets should be used when your app is completely decoupled from the node and you can run multiple copies on a given node without further consideration.  
DaemonSets should be used when a single copy of your application must run on all or a subset of nodes in the cluster (think prometheus exporters and filebeat)  
DaemonSets will create pods on every node by default, unless a node selector is used.  
Deleting a DaemonSet will also delete all the pods managed by that DaemonSet

### CHAPTER 10: Jobs
Jobs: Made to handle short lived, one off tasks. Ruh-roh, I wonder if these are anything like cron jobs.  
A job creates Pods that run until the job finishes  
By default a job runs on a single pod until termination. Jobs support parallelism well.  
Pods spawned by Jobs and stay in "CrashLoopBackOff" status (?) it the job fails.


### CHAPTER 11: ConfigMaps and Secrets
You should strive to make images as reusable as possible (same image for dev, test and prod).  
Testing can get riskier and more complicated if images are different for each environment.
ConfigMaps are used to provide configuration information.
Secrets are similar to ConfigMaps but focused on making sensitive information available to the workload.  
ConfigMap: A set of variables that can be used when defining the environment for your containers.  

#### Secrets
Secrets enable container images to be created without bundling sensitive data.  
Careful: Anyone with root access to a node can see the secrets.  
Secrets are created using the kubectl command line tool.  
You can also expose secrets to pods using the secret volume type.  

### CHAPTER 12: Deployments
The deployment object exists to manage the release of new applications  
Deployments represent deployed applications  
Deployments allow you to carefully move from one version to another of your code.  
The Deployment yaml structure is very similar to that of the replicaset, but there's a strategy part which describes how to roll out new software.  

### CHAPTER 13: Integrating Storage Solutions and Kubernetes

**Good chapter, follow up thoroughly later**

At some point, you need to have data stored somewhere. Integrating this data with container orchestration is oftentimes the most difficult aspect of building a distributed system.  
Author advises: Moving a database is not wise as a first step.
The author also echoes what I think: For live, you can remain using the legacy database, but for continous testing, you may use a test database as a transient container.  

You can create a service named ExternalName, which makes the cluster add a name entry pointing to something outside the cluster (handy to connect to a legacy db for instance)  
You may also run just a single pod that runs the database or other storage solution. It is not less reliable than running it on a VM or physical machine.  

The chapter then describes how to run a MySQL singleton (reliably!) **investigate later more**

Persistent Volume: Has a lifetime independent of any pod or container.  
Book then decouples storage definition from pod definition (write the yamls here)  
Author then uses a ReplicaSet with a single pod, to ensure the pod is rescheduled somewhere else in case of node failure.


### CHAPTER 14: Deploying Real-World Applications











General feel of the book:  
A bit heavy on the Jargon and the theory so far.  
Examples are not comprehensive, just a "to do this, run this"  
Page 30 edit: I think the authors meant to keep the book short by not using too many words.  
Page: 30  