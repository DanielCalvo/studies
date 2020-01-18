### 99: OS Upgrades
- When a node goes down, the pods on that node become obviously unaccesible
- Depending on how you deploy those pods, your users can be affected or not
- If your pod was part of a replica set, there will be another instance of your pod in another node and everything will be ok

- If the node comes back almost immediately, the kubelet starts and the pods come back online
- If the node is down for more than 5 minutes, the pods are terminated from that node. Kubernetes considers them as dead
- If the pods were part of a replicaset, they will be recreated on other nodes.
- The time it waits for a pod to come back online is known as the --pod-eviction-timeout and is set on the on the kube-controller-manager with a default value of 5 minutes

So when a node goes offline, the master node waits up to 5 minutes before considering the node dead.

- When a node comes back online after the pod eviction timeout, it'll come back with 0 pods scheduled on it.
- Pods that were on this node that were not part of some sort of replicaset will be gone forever.

- You can use `kubectl drain mynode` to safely drain the node of all the workloads so that they're moved to other nodes.
- When you drain a node, the pods are gracefully terminated from that node and created on another. The node that was drained is then marked as unschedulable.

- After you do maintenance on the drained node, you need to run `kubectl uncordon mynode` so that's marked as schedulable.
- The pods that moved to other nodes do not immediately fall back to the old node.

There is also the `kubectl cordon mynode` command. It will mark a node as unschedulable, but it will not terminate any pods running on it.

### 100: Practice test - Cluster upgrades
Q: We need to take node01 out for maintenance. Empty the node of all applications and mark it unschedulable.
A: `kubectl drain node01 --ignore-daemonsets`

Q: The maintenance tasks have been completed. Configure the node to be schedulable again.
A: `kubectl uncordon node01`

Q: Why are there no pods on node01?
A: Only newly scheduled pods will be created there

Q: Can you drain node02 using the same command as node01? Try it
A: Nop, you get: `error: cannot delete Pods not managed by ReplicationController, ReplicaSet, Job, DaemonSet or StatefulSet (use --force to override): default/hr-app`. You have to:
- `kubectl drain node02 --ignore-daemonsets`
- `kubectl drain node02 --ignore-daemonsets --force`
- Pods not part of a ReplicationController, ReplicaSet, Job, DaemonSet or StatefulSet will be gone forever when you drain a node that contains them! 

Q: Node03 has our critical applications. We do not want to schedule any more apps on node03. Mark node03 as unschedulable but do not remove any apps currently running on it .
A: `kubectl cordon node03`

### 101: Kubernetes releases & versions
- When you run the kubectl get nodes you get the version of the Kubernetes cluster 
- The kubernetes versioning consists of 3 parts:
- v1.11.3 (major.minor.patch)

- There are also alpha and beta releases with bugfixes and features being tested

- The kubernetes release package in github, when downloaded, has all the control plane software in the same version
- (kube-apiserver, controller-manager, kube-scheduler, kubelet and kube-proxy all v1.13.4 for instance)
- There are other components in the control plane that do not have the same version numbers, such as ETCD and coreDNS

### 102: References:
- Further reading as suggested by the author:
- https://kubernetes.io/docs/concepts/overview/kubernetes-api/
- https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md
- https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api_changes.md

### 103: Cluster upgrade process (stopped here)
- No component in the cluster can be at a version higher than kube-apiserver
- The versions of certain services can be lower than others, but only in a specific way. Let's assume that kube-apiserver is at version X.

- kube-apiserver -> X (ex: v1.10)
- controller-manager, kubescheduler -> X-1 (can be up to 1 version older, ex: v1.9)
- kubelet, kube-proxy -> (can be up to 2 versions older, ex: v1.8)

- The exception to this is the kubectl command line tool, which can be the one version newer or one version older than kube-apiserver.
- Everything can also be the same version if you're not an adventurer.

- When should you upgrade? At any given time, Kubernetes supports only the most recent minor versions.
- If v1.12 is the lastest release, Kubernetes will support v1.12, v1.11 and v1.10
- When v1.13 is releases, v1.10 will go unsupported.

- The recommended update strategy is to update things one minor version at the time (v1.10 to v1.11 for example)

- The upgrade process depends on how your cluster is set up. If you're using GKE, you can upgrade your cluster with just a few clicks.
- If you deployed your cluster with kubeadmin, this can help you plan and upgrade your cluster with
- `kubectl upgrade plan`
- `kubectl upgrade apply`

- If you deployed the cluster from scratch (the hard way (TM)) then you manually have to upgrade things yourself.

- The following options relate to the upgrade strategy on kubeadm
- Upgrading a cluster has two major steps:
    - First, you upgrade your master node.
    - Then, you upgrade your worker nodes

- When the master is being updated, all the control plane processes go down briefly.
- While the master is down, all workloads on the worker nodes continue to serve users. But you cannot access the cluster using kubectl. You cannot deploy new things, delete or modify existing ones.
- The controler managers don't function either, if a pod fails, it won't be recreated.

- As far as upgrading the worker nodes, you can upgrade all of them at the same time, but then all of your pods go down.
- Once the worker nodes are back up, new pods are started and everything is back to normal

- You can also only upgrade one node at a time. After upgrading the master, upgrade the first node, and the workloads move elsewhere.
- Once the first node is upgraded and back up, you bring the second node down and upgrade that one.

- A third way of doing this, is simply adding nodes to the cluster with a newer software version already installed. Very convenient if you're in a cloud environment where you can easily create new machines.
- Just gradually add new nodes and remove the old ones, untill all nodes are new and on the new version.

- Running the kubeadm upgrade plan will give you a lot of useful information (current cluster version, kube adm tool version and a lot of other things)
- After upgrading the cluster, you must manually upgrade the kubelet version on each node.

- To upgrade the master:
    - `apt-get upgrade -y kubeadm=1.12.0-00`
    - `kubeadmin upgrade apply v1.12.0`

- The next step is to upgrade the kubelets. You may or may not have kubelet installed on your master node.
    - `apt-get upgrade -y kubelet=1.12.0-00`
    - `systemctl restart kubelet`

- To upgrade a node:
    - `kubectl drain node-1`
    - `apt-get upgrade -y kubeadm=1.12.0-00`
    - `apt-get upgrade -y kubelet=1.12.0-00`
    - `kubeadm upgarde node config --kubelet-version v1.12.0`
    - `systemctl restart kubelet`
    - `kubectl uncordon node-1`

##104 : Practice test cluster upgrade
Q: What's the cluster version?
A: `kubectl version`

Q: What is the latest stable version available for upgrade?
A: `kubectl upgrade plan`

Q: We will be upgrading the master node first. Drain the master node of workloads and mark it UnSchedulable
A: `kubectl drain master --ignore-daemonsets`

Q: Upgrade the master components to v1.12.0 
A: 
- `apt-get install -y kubeadm=1.12.0-00`
- `kubeadm upgrade apply v1.12.0`
- `apt-get install kubelet=1.12.0-00`

Q: Mark the master node as "Schedulable" again
A: `kubectl uncordon master`

Q: Next is the worker node. Drain the worker node of the workloads and mark it UnSchedulable
A: `kubectl drain node01 --ignore-daemonsets`

Q: Upgrade the worker node to v1.12.0
A:
- `apt install kubeadm=1.12.0-00`
- `apt install kubelet=1.12.0-00`
- `kubeadm upgrade node config --kubelet-version 1.12`
- `systemctl restart kubelet`
- `kubectl uncordon node01`

### 106: Backup and Restore methods 
- On the ETCD cluster is there all the cluster related information is stored!
- Persistent volumes on the cluster are also a candidate for backups
- A good practice is to store your yaml definitions on a source code repository
- Possible backup of everything on the cluster: `kubectl get all --all-namespaces -o yaml > all-deploy-services.yaml`
- Velero is good for backups

#### ETCD Backup
- ETCD has a datadir configured by the `--data-dir` parameter
- `etcdctl` also has a snapshot feature: `etcdctl snapshot save snapshot.db`
- `etcdctl snapshot status status snapshot.db` 
- To restore your cluster from this backup: 
    - `service kube-apiserver stop`
    - `etcdctl snapshot restore snapshot.db \
    --data-dir xxx \
    --initial-cluster xxx
    --initial-cluster-token xxx
    --initial-advertise-peer-urls xxx`
- When you run the above command, a new data directory will be created
- Then remember to change the following settings on etcd:
    - `--initial-cluster-token xxxx`
    - `--data-dir`
- Then:
    - `systemctl daemon-reload`
    - `service etcd restart`
    - `service kube-apiserver start`

- With all the etcd commands, remember to specify `--endpoints, --cacert, --cert, --key`
- If you're using a managed k8s solution, just get the yamls, you may not even have access to the etcd cluster

### 107: Practice test - Backup and Restore Methods
Q: What's the version of etcd on the cluster?
A: `kubectl exec -it etcd-master -n kube-system sh`
`etcd --version`

Q: At what address do you reach the ETCD cluster from your master node?
A: `kubectl describe pod etcd-master -n kube-system`
`--listen-client-urls=https://127.0.0.1:2379,https://172.17.0.48:2379`

Q: Where is the ETCD server certificate file located?
Q: Where is the ETCD CA Certificate file located?
A: `kubectl describe pod etcd-master -n kube-system`

Q: Backup etcd!
A: `etcdctl backup --data-dir=/var/lib/etcd --backup-dir=./bkp`

Q: Luckily we took a backup. Restore the original state of the cluster using the backup file.
A: https://github.com/mmumshad/kubernetes-the-hard-way/blob/master/practice-questions-answers/cluster-maintenance/backup-etcd/etcd-backup-and-restore.md

### 108: Certification Exam Tip!
- The exam won't tell you if what you did is correct or not
- If the questions asks you to create a pod with a certain image, validate it yourself by doing `kubectl describe pod`

### 109: References
- https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/#backing-up-an-etcd-cluster
- https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/recovery.md