
### 76: Rolling updates and rollbacks
- When you first create a deployment, it triggers a rollout.
- A new rollout creates a new deployment revision. Let's call it rev1.
- When the container version is updated (new image?) a new rollout is triggered, generating a new revision, which we can call rev2.
- This allows us to keep track of the changes made to our deployments and enables us to rollback if necessary.

- You can see the rollout status of a given deployment by doing:
- `kubectl rollout status deployment/myapp-deployment`

- To see the revisions and roll out history:
- `kubectl rollout history deployment/myapp-deployment`

There are two type of deployment strategies.
- Recreate strategy: All the pods in the deployment are destroyed, and then recreated.
- This is disadvantageous as it causes downtime in between the time all the containers are destroyed and the new ones are spawning up.

- Rolling update: Imagine 5 pods. Rolling update takes down 1 pod with the old version, and spawns up a new pod with the new version.
- It keeps repeating this process until all the pods are running the new version. This is the default strategy.

- To change the image version of your deployment, you would simply change the it on the deployment definition yaml and do a:
- `kubectl apply -f mydeployment.yaml`

- But you can also do it imperatively:
- `kubectl set image deployment/myapp-deployment nginx=nginx:1.9.1`

- When using kubectl describe deployment mydeployment, if using the Recreate strategy, you will see the underlying ReplicaSet being scaled down to 0, and then scaled up to whatever number it had previously.
- When using the RollingUpdate strategy, the old replicaset will be scaled down by 1, while the new replicaset will be scaled up by 1 over time.
- This process will repeat itself until all pods are on the new replicaset.

- To rollout a deployment, you can:
- kubectl rollout undo deployment/myapp-deployment

- A deployment runs a ReplicaSet under the hood
- You can see the new and old replicasets both change the number of replicas to reflect the undo change (pods will be gradually destroyed on the new and started on the old)

- To create a deployment imperatively: `kubectl run nginx --image=nginx`

- Command summary:
- `kubectl -f create mydeployment.yaml`
- `kubectl get deployments`
- `kubectl apply -f mydeployment.yaml #With a new image version`
- `kubectl set image deployment/mydeployment nginx=nginx:1.9.1`
- `kubectl rollout status deployment/myapp-deployment`
- `kubectl rollout history deployment/myapp-deployment`
- `kubectl rollout undo deployment/myapp-deployment`


### 77: Practice Test: Rolling Updates and rollbacks
- Pretty easy, it was just editing existing objects, changing their update strategy and images and seeing how the pods come and go


### 80: Commands:
- Not a required topic, but important to know!
- Commands and arguments on a pod definition file. Not strictly part of the certification, but author believes it's important!
- In case of ENTRYPOINT, whatever you specify as an argument, will get appended to the entrypoint.
- In case of CMD, the command parameters passed will get replaced entirely.
- You can also use ENTRYPOINT and CMD like:
```
FROM Ubuntu
ENTRYPOINT ["sleep"]
CMD["5"]
```

- The default argument for sleep will be 5, but you can overwrite it by passing an argument to the container at launch time.
You can also do something like: `docker run --entrypoint sleep_v2 ubuntu-sleeper 10`

### 81: Commands and Arguments
- In a pod definition:
    - Command is analogous to ENTRYPOINT
    - args is analogous to CMD


67: Configure Environment Variables in applications
You can set environment variables as a key-value under env on your deployment/pod/etc yaml definition
You can also set them on a configmap or a secret


68: Configuring configmaps in applications
Having too many environment variables set up on a pod definition file will become unwieldy very quickly (duplication over all files & a lot of repeating yourself)
We can take that out of the pod definition and put it on a ConfigMap! :D

There are two phases when creating a configmap:
1. Create the configmap and define variables you want
2. Inject these variables into the pod

Just like any Kubernetes object, there are two ways of creating a configmap.
Imperative: kubectl create configmap
Declarative: kubectl create/apply -f

Creating a configMap imperatively:
kubectl create configmap myconfigname --from-literal=mykey=myvalue
kubectl create configmap myconfigname --from-file=pathtomyfile

kubectl get configmaps
kubectl describe configmaps
kubectl describe configmap myconfigmap


69: Practice test: Environment variables
You can't edit environment variables of a running pod :(
(Follow up from Dani: Creating a pod loading the configs from a configmap was really tough. You need to work on this!


70: Configure secrets in applications
Secrets are very similar fom configmaps, except they're stored on a encoded format

You can also create a secret imperatively and declaratively:
kubectl create secret generic mysecret --from-literal=mysecret=myvalue --from-literal=mysecret1=myvalue1
kubectl create -f

To generate a secret to put on a secret file, do:
echo "mysecret" | base64

To decode the secret:
echo "bXlzZWNyZXQK" | base64 --decode

To see secrets:
kubectl get secrets

To see the actual values of the secrets:
kubectl get secret mysecret -o yaml


71: A note about secrets
Author explains that the way secrets are set up in Kubernetes isn't the safest or most coherent option.
(Follow up from Dani: This page is suggested as further reading: https://kubernetes.io/docs/concepts/configuration/secret


72: Practice test - Secrets
(Follow up from Dani: Why did the test accept a secret created imperatively, but not descriptively? What did you mess up? :thinking:)
kubectl create secret generic db-secret --from-literal=DB_Host=sql01 --from-literal=DB_User=root --from-literal=DB_Password=password123

(Follow up from Dani: Try adding secrets to other objects other than pods, such as Deployments and DaemonSets)


74: Multi container pods
You can have your microservice container and a side car container (such as a log shipper) inside the same pod.
This is because both containers need to scale up and down together.
They share the same network space, which means they can refer to each other as local host.
They both have access to the same volumes as defined in the pod definition.


75: Practice test - Multiple container pods
(Follow up from Dani: How do you change the default namespace so you can do "kubectl get pods" from a different namespace without having to specify a namespace?)


76: Multi-container pods design patterns
There are three multi container POD design patterns: Sidecar, Adaptor and Ambassador


77: initContainers
When running a multi container pod, all pods are expected to stay alive at all times.
If any container fails, the entire pod restarts.

You can specify a type of container inside a pod named initContainer. Inside the spec, instead of being inside "containers" it's inside "initContainers". Check 77_initcontainer.yaml

The initContainer container will run first, before the regular containers start. This can be useful if you need to do some sort of set up before your main container starts (such as cloning a repo)
You can have multiple containers listed under initContainer. They will run sequentially.
You can use initContainers to wait for

If any of the initContainers fail to succeed, Kubernetes will restart the pod repeatedly until the initContainer(s) succeed.

(Follow up from Dani: Can you do this with deployments as well? Hey maybe you should give all the concepts below a read, maybe they cover more things that you don't know!)
(Follow up from Dani: Doublecheck/redo initContainers, the katakoda lab was not recognizing them for some reason!)
More reading material is available here: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/

79: Self healing applications:
Kubernetes by default restarts pods when they crash.
But I remember liveliness probes and readiness probes and other things
(Follow up from Dani: Can you take a look at this later?)