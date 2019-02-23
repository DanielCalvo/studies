
###CHAPTER 1: Introductions

####How Kubernetes can improve things:

- **Velocity**: It's fast and you can deploy a bunch of things fast without breaking everything
- **Immutable**:
- **Declarative configuration**:
- **Self healing**:
- **Immutable Infrastructure**: 
- **Scaling**:
- **Abstracting your infrastructure**:
- **Efficiency**:

###CHAPTER 2: Creating and Running containers

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

###CHAPTER 3: Deploying a Kubernetes Cluster

####Installing Kubernetes on a Public Cloud Provider

I'll do it on Google Cloud because that's what all the cool kids are using.  
Login into Google Cloud, launch their shell and run:
```
gcloud config set compute/zone us-west1-a
gcloud container clusters create danitest-cluster
gcloud auth application-default login
```
####Checking Cluster Status
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

####Cluster Components
All these run in the kube-system namespace.

- **Kubernetes Proxy**:
- **Kubernetes DNS**:
- **Kubernetes UI**:

###CHAPTER 4: Common kubectl Commands
- **Namespaces**:
- **Contexts**:

####Viewing Kubernetes API Objects

Each object in Kubernetes exists at a unique HTTP path! Like:  
https://your-k8s.com/api/v1/namespaces/default/pods/my-pod  

The most basic command for viewing Kubernetes objects is through `kubectl get`, you can pass the `-o wide`, `-o json` and `-o yaml`flags to it
You can also use JSONPath to extract certain atributes using the `kubectl get` command:
```
kubectl get pods my-pod -o jsonpath --template={.status.podIP}
```
####Creating, Updating, and Destroying Kubernetes Objects
In a nutshell:
```
kubectl apply -f obj.yaml
kubectl delete -f obj.yaml
```
####Labels and anotations
Well looks like you can label things. Nice. Skipping this one for now.

####Debugging Commands
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
###CHAPTER 5: Pods


###CHAPTER 6: Labels and Annotations
###CHAPTER 7: Service Discovery
###CHAPTER 8: ReplicaSets
###CHAPTER 9: DaemonSets
###CHAPTER 10: Jobs
###CHAPTER 11: ConfigMaps and Secrets
###CHAPTER 12: Deployments
###CHAPTER 13: Integrating Storage Solutions and Kubernetes
###CHAPTER 14: Deploying Real-World Applications











General feel of the book:  
A bit heavy on the Jargon and the theory so far.  
Examples are not comprehensive, just a "to do this, run this"  
Page 30 edit: I think the authors meant to keep the book short by not using too many words.  
Page: 30  