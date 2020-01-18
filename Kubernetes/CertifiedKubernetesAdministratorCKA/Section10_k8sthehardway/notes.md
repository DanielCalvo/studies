
### 180: Designing a Kubernetes cluster
- Ask yourself:
    - What is the purpose of this cluster?
    - What is the cloud your org uses?
    - What workloads and applications are you going to host on this cluster? 
    - What type of traffic?

- For studying:
    - Minikube
    - Single node cluster with kubeadm/GCP/AWs
- For dev/testing
    - Multi-node cluster with a single master and multiple workers
    - Use kubeadm or a managed solution on AWS/GCP 
- For production level clusters
    - Multi node, multi worker
    - Kubeadm, GCP or Kops
    - Up to 5000 nodes, 150000 pods, 300000 containers and up to 100 pods per node
    - GCP/AWS do/may automatically choose the node size for you according to your cluster requirements

#### Depending on your local/cloud environment
- Use kubeadm for on-prem
- GKE for GCP
- Kops for AWS
- AKS for Azure

#### Storage requirements and pruporse
- For high performance: SSD backed storage
- For multiple concurrent connections: Network based storage
- Use persistent shared volumes for data shared accross pods
- Label nodes with specific disk types
- Use node selectors to assign applications to nodes with specific disk types

- On very large k8s deployments, you can make the etcd cluster be separated from the other control plane components
    

It appears that on this section there's nothing certification _specific_.
Though it would certainly help to remember what you did here, as it builds familiarity with k8s.

### 181: Choosing Kubernetes infrastructure.

- Ways to deploy on prem for studying:
    - Minikube: Deploys VM, single node.
    - Kubeadm: Requires VM to be ready, single/multi node

For production use, there are two types of solutions:
- Turnkey solutions:
    - You provision VMs
    - You configure VMs
    - You use scripts to deploy the cluster
    - You maintain VMs yourself
    - Ex. k8s on aws with kops
    - Popular turnkey solutions: Openshift, Cloud Foundry Container Runtime, VMware Cloud PKS, Vagrant

- Hosted solutions:
    - Kubernetes as a service
    - Provider provisions VMs
    - Provider install Kubernetes
    - Provider maintains VMs
    - Ex: GKE
    - Popular hosted solutions: GKE, Openshift online, AKS, EKS

### 182: Choosing a network solution
- Supported network solutions: https://kubernetes.io/docs/concepts/cluster-administration/networking/#how-to-implement-the-kubernetes-networking-model
- Good read: https://www.objectif-libre.com/en/blog/2018/07/05/k8s-network-solutions-comparison/
- When choosing a network solution, consider it's support for Network Policies.
- We're gonna use weavenet!
 

### 183: Configure high availability
- The master hosts the control plane components on your cluster. If it goes down, bad deal!
- Consider master HA!

- kubectl tool points to kube-apiserver running at a given IP.
- Tutorial recommends having load balancing in front of a multi master set up

- What about the scheduler and the controller manager?
- These are controllers that watch the state of the cluster and take actions.
- If you have two of them active simultaneously, actions might be duplicated.

- As such, these components do not run in parallel. They run on an active-standby mode.
- There is a leader election process.
- Both kube-controller-manager and scheduler have a --leader-elect options and lease settings for the leader.
- They'll constantly try to obtain a lock on a kubernetes object, as to always have service in standby and another one active.

- ETCD: Two topologies available.
- Stacked control plane nodes topology: ETCD running in every master node?
- External ETCD Topology: ETCD runs on it's own set of servers.

- When you set up kube-api server, there's a parameter indicating the address of the etcd servers.
- Ultimately, you just need to make sure that the api server can reach it.

### 184: ETCD in HA
- On an ETCD cluster, you can read/write from any instance. ETCD ensures that data is consistent on all instances at the same time.
- ETCD has a leader node. If a write operation arrives to the leader node, it happens there and then is propagated to all other ETCD nodes.
- If a write comes to follower node, then they forward the write to the leader internally, the leader processes it, writes it, and then the change is propagated to the follower nodes.

- How is the leader elected, and how do we make sure that a write is propagated through all instances?
    - ETCD implements distrubuted consensus using the RAFT protocol.
    - RAFT implements random timers upon startup. The first to send a request to the other nodes to be a leader, becomes a leader.
    - The leader sends notifications to other etcd instances informing that it is the leader.

- If the nodes do not receive messages from the leader after some time, the nodes initiate a re-election process and a new leader is set up.
- When a write to etcd comes in, it is only considered complete when it is replicated to other instances in the cluster.

- A write is considered to be complete if it can be written in the majority of the nodes on the etcd cluster.
- Quorum = Minimum number of nodes that must be available for the cluster to function properly or make a successful write.
- Quorum = N/2 + 1
- Having 2 instances is the same as having 1 instance, which is why it is recommended to have a minimum of 3 instances.
- In case of network segmentation, there are better chances for your cluster to stay alive if you're running an odd number of ETCD instances.

### All else
- All further lectures on this section are practical lectures from Mumshad's k8s the hardway on virtualbox
- https://github.com/mmumshad/kubernetes-the-hard-way