## Kubernetes patterns! Let's do this.
- This book was read on May/June 2021.
- Good book so far! Has some useful information on it.
- On page 31 there's a very handy list of objects you can explore as part of your "exporing k8d objects" idea!

## Chapter 2: Predictable Demands
- Set resource requests and limits on your pods. Depending on how you set these pods can be Best Effort, Burstable or Guaranteed.
- `kind: PriorityClass` lets you define which pods are most important relative to each other. More important pods will be placed first and other pods might be removed to make space for them.
- Note that if a volume is only available on one node, and pods require that volume, the pods will be scheduled on that node. If the volume is unavailable, pods will not be placed at all :o
- If certain pod requires a configmap, the absence of this config map will cause the pod to not start :o

## Chapter 3: Declarative deployment
- For deployments to work properly, containers must react appropriately to SIGTERM
- Out of the box deployment strategies: Rolling and recreate
- Not much else here other than an overview of vanilla k8s deployments

## Chapter 4: Health probe
- Process health check: Comes for free, k8s restarts your container if it's process is not running
- Liveness probe: K8s checkds your container to confirm it's still health. This can be: An HTTP probe, a TCP connection, or an exec command
- Readiness probe: It's the same as a liveness probe, but when failed, instead of restarting the container, it is removed from the service endpoint and does not receive any traffic. Used to allow pods to start up and remove them from traffic when overloaded
- Book suggests: Different API endpoints for liveness and readiness checks

- TODO: Not on the book maybe, start up probe?

## Chapter 5: Managed Lifecycle
- Poststart hook:
- PreStop hook:

## Chapter 6: Automated Placement
- Oh I forgot you can run multiple schedulers
- You can use .spec.nodeSelector to select a given node for your pods, like one with an SSD or GPU (remember to label your nodes!)
- Remember that nodeAffinity is a thing! You can tell nodes to be scheduled in nodes with more than 3 cpus. There are "preffered" and "required" requirements. There are also antiaffinity rules!
- You can taint a node and you can have a pod tolerating that taint.
- There is a kubernetes descheduler that can tryto re-schedule pods

# Part II: Behavioral Pattenrs

## Chapter 7: Batch job
- There's a `kind: Job` for batch jobs. These are jobs that run _right now_, one or more times. You can specify many runs (it'll queue) with a parallelism setting

## Chapter 8: Periodic job
- Extends the batch job. Essentially a cronjob! The "startingDeadlineSeconds" parameter is interesting! It means that if a job can't start in X seconds, it is not scheduled

## Chapter 9: Daemon Service
- DeamonSet: Pods that run on the cluster nodes and provide background capabilities for the rest of the cluster. It's main purpose is to run a single pod on every node (or specific nodes). You can have a nodeSelector. Main use cases: log collectors, metric exporters, and even kube-proxy
- Pods spawned by a Daemonset don't require a scheduler. Neat!

## Chapter 10: Singleton Service
- That whole singleton stuff -- kinda weird to think about it. It seems like a really peculiar use case. Some more examples would've helped!
- K8s doesn't support singletons out of the box, there might be more than one pod running at at time for a singleton. You might need to work around that

### Pod Disruption Budget
- Ensures that a number or percentage of pods will not be evicted at a given time
- Good to make sure that you always have a number of replicas up during a deployment, for instance!

## Chapter 11: Stateful Service
- Stopped here!

## Chapter 22: Controler
- A controler monitors a set of k8s resources in a given state
- Kubernetes is a distributed state manager. You give it a desired state for a given component, and it will try to reconciliate that state
- Controller loop: Observe, analyze, act
- Operators: More sofisticated controllers
- Controllers: Simple reconciliation loop that monitors and actos on standard Kubernetes resources
- Operators: A sophisticated reconciliation process that interactis with CustomResourceDefinitions (CRDs) which are at the heard of the operator pattern. Neat!
- The controller done in shellscript to delete pods based on what's on a configmap was really cool!

## Chapter 23: Operator
- An operator is a Controller that uses a CRD to encapsulate operational knowledge for an application
- Custom Resource Definitions (CRD) allow extensions of the Kubernetes API, by adding custom resources to your Kubernetes clusters and using them as if they were native resources.
- There are Installation CRDs and Application CRDs
- Hmm RBAC can also be used for security
- Hmm there seems to be a difference between regular CRDs and API CRDs
- Toolkits
    - CoreOS operator framework
    - Kubebuilder
    - Metacontroller
- Topics you can google: kubebuilder, awesome operators, k8s sample-apiserver, k8s apiserver-builder, openapi v3, kubernetes documentation operators
