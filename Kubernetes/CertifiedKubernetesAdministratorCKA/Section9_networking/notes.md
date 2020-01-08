 ### 152: Prerequisite - Switching routing
 - Linux network recap & review
 - `ip link`, `ip addr`
 - A router routes between networks
 - `route`
 - `ip route add ...`
 - `/proc/sys/net/ipv4/ip_forward`

### 153: Prerequisite - DNS
- `/etc/hosts`
- `/etc/resolv.conf`
- Domain names
- Subdomains
- Record types (A, AAAA, CNAME)
- nslookup, dig

### 154: CoreDNS
- Read more about CoreDNS here:
- https://github.com/kubernetes/dns/blob/master/docs/specification.md
- https://coredns.io/plugins/kubernetes/

### 155: Prerequisite - Network Namespaces
- Network namespaces are used by certain services (such as Docker) to implement network isolation
- To create a network namespace: `ip netns add $NAME`
- To list: `ip netns`
- To run commands inside a network namespace: `ip netns exec red ip link`
- The arp command will return different entries if ran on the host or on the container
- You can set up intra namespace communication by performing operations with the `netns` subsection of `ip`. There are a lot of commands on this lecture
- You can create a linux bridge network to enable many components to talk to each other over a virtual switch

### 156: FAQ
- While testing the Network Namespaces, if you come across issue where you can't ping one namespace from the other, make sure you set the NETMASK while setting IP Address. ie: 192.168.1.10/24

### 157: Prerequisite - Docker Networking
- Docker container network types:
    - None: No network
    - Host network: Container is attached to the host network, no isolation. If you deploy a app on port 80, it'll be available on port 80 on the host
    - Bridge: An internal network shared between Docker and the host

### 158: Prerequisite CNI
- The container network interface (CNI) is a set of standards that defines how programs should be developed to solve networking challenges in a container runtime environment.
- Docker does not implement CNI. Docker has it's own set of standards named CNM (Container network model)
- What k8s does is it launches docker containers on the none network, and then uses it's plugins to attach an interface to it

### 159: Cluster networking:
- Might go without saying, but all nodes in the cluster need to have unique hostnames
- Careful when clonning VMs (hostnames will also be clonned)
- Ports that need to be open:

- Master node:
    - ETCD: 2379
    - kube-api: 6443
    - kubelet: 10250
    - kube-scheduler: 10251
    - kube-controller-manager: 10252
    - If you have a multimaster setup, you need port 2380 open so that ETCD clients can talk to each other.

- Worker node:
    - Services: 30000 - 32767
    - kubelet: 10250

- Commands that can be handy:
    - `ip (link|addr|route), arp, netstat`

### 160: Practice test -  Explore Kubernetes environment:
Q: What is the network interface configured for cluster connectivity on the master node?
A: `ip link` - ens3

Q: What is the IP address assigned to the master node on this interface?
A: `ip addr` - 172.17.0.8/16

Q: What is the MAC address of the interface on the master node?
A: `ip addr`

Q: What is the IP address assigned to node02?
A: kubectl get nodes -owide 

Q: What is the MAC address assigned to node02?
A: `arp node02`

Q: We use Docker as our container runtime. What is the interface/bridge created by Docker on this host?
A: `docker network inspect bridge` && `ip addr` (docker0) 

Q: If you were to ping google from the master node, which route does it take? What is the IP address of the Default Gateway?
A: `ip route` (default via 172.17.0.1 dev ens3)

Q: What is the port the kube-scheduler is listening on in the master node?
A: `netstat -anlp`

Q: Notice that ETCD is listening on two ports. Which of these have more client connections established?
A: `master $ netstat -anlp | grep 2379 | wc -l`

### 161: Pod networking
- Requirements for k8s networking: 
    - Every POD should have an IP address
    - Every POD should be able to communicate with every other POD in the same node.
    - Every POD should be able to communicate with every other POD on other nodes without NAT.

- The first thing that is done in the video tutorial is to set up a bridge interface internal to every single node that every pod on this node connects to.
- This way these pods on the same node can communicate, but not with pods on other nodes yet.

- You then need to add routes on every node.
- A given node will have routes to all the private internal bridge networks on all the other nodes.

- To conform to CNI, this script needs to have an ADD and a DEL section.

- The kubelet on each node is responsible for creating containers. Whenever a container is created, kubelet looks into:
- `--cni-conf-dir=/etc/cni/net.d`
- `--cni-conf-dir=/etc/cni/bin`
- To run net-script.sh with arguments: `add, <container>, <namespace>`

### 162: CNI in Kubernetes
- CNI is configured/invoked on/by the Kubelet
- `--network-plugin=cni`
- `--cnd-bin-dir=/opt/cni/bin`
- `--cni-conf-dir=/etc/cni/net.d`

- Also viewable by doing ps aux | grep kubelet and checking the args for kubelet
- The bin directory has all the CNI plugins as executables:
- `ls /opt/cni/bin`

- The configuration dir:
- `ls /etc/cni/net.d`
- Has as set of config files, which is where the kubelet sees which plugin is to be used. If there are multiple files there, it will chose one in alfabetical order
- Under `/etc/cni/net.d/10-bridge.conf`

That conf file will have info if an interface is gateway, the ip range and name, and a bunch of other things.

### 163: Practice test - Explore CNI Weave
Q: Inspect the kubelet service and identify the network plugin configured for Kubernetes.
A: `ps -aux | grep kubelet` (--network-plugin=cni)

Q: What is the path configured with all binaries of CNI supported plugins?
A: `ps -aux | grep kubelet` (--cni-bin-dir=/opt/cni/bin)

Q: Identify which of the below plugins is not available in the list of available CNI plugins on this host
A: `/opt/cni/bin` contains a bunch of executables. Answer: Cisco

Q: What is the CNI plugin configured to be used on this kubernetes cluster?
A: Under /etc/cni/net.d/10-weave.conf is the only config available, so I'm gonna go with that.

### 164: CNI weave, (aka weaveworks)
- As you expand your cluster, it is very difficult to keep track of all the routes.
- Weaveworks deploys an agent on each node. These agents communicate with each other to exchange information on pods and networks on every node.
- Each agent stores a topology of the entire set up.
- Weave can be deployed on the node system manually (as in kubernetes the hardway) or it can run in a pod in k8s.

- You can see the weave pods with:
- `kubectl get pods -n kube-system`
- `kubectl logs weave-net-1234 weave -n kube-system`

### 165: Practice test - Explore CNI Weave 2
Q: What is the Networking Solution used by this cluster?
A: `ps aux | grep kubelet` (weave, see previous practive test)

Q: How many weave agents/peers are deployed in this cluster?
A: `kubectl get daemonset --all-namespaces`

Q: On which nodes are the weave peers present?
A: On all nodes (see number of members on daemonset and number of nodes)  

Q: Identify the name of the bridge network/interface created by weave on each node
A: `ip link` or 

Q: What is the POD IP address range configured by weave?
A: `ip link`

Q: What is the default gateway configured on the PODs scheduled on node03?
A:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox
  namespace: default
spec:
  nodeName: node03
  containers:
  - name: busybox
    image: busybox:1.28
    command:
      - sleep
      - "3600"
    imagePullPolicy: IfNotPresent
  restartPolicy: Always
```
- `kubectl exec busybox -it -- sh`
```
/ # ip route
default via 10.32.0.1 dev eth0
10.32.0.0/12 dev eth0 scope link  src 10.32.0.3
```

### 166: Practice test - Deploy Network Solution
- Just deployed weavenet


### 144: IP Adress Management - Weave (IPAM)
- CNI says it's the responsibility of the network plugin to assign IPs to the containers (and it's important that we don't have network collisions!)
- There's a plugin called "host local plugin" that manages free IP addresses for pods inside the host apparently
- Inside /etc/cni/net.d/net-script-conf there is an IPAM section in which we specify the type of plugin to be used, the subnet and route info
- Weaveworks by default allocates the range 10.32.0.0/12 for the entire network
- This range is split among nodes

### 168: Service networking
- A service is just a virtual object.
- kube-proxy gets the ip of a given service and creates forwarding rules on each node saying that any traffic coming to the ip of the service needs to go to the IP of the pod.
- So whenever you try to access a service, you get forward to the pod's ip adress.
- But remember: It's not just the IP, it's an IP:port combination.
- Whenever these services are created or deleted, the kubeproxy creates or deletes these rules.
- Ways these rules can be created: userspace, ipvs, iptables. iptables is the default option.
- The range for the IP of services is set by the --service-cluster-ip option of kube-apiserver
- You can actually check the rules for your service by doing:
- `iptables -L -t net | grep myservice`
- You can also see these entries being created in /var/log/kube-proxy.log

### 169: Practice test - Service networking
Q: What network range are the nodes in the cluster part of?
A: `kubectl get nodes -owide` or `ip addr`

Q: What is the range of IP addresses configured for PODs on this cluster?
A: The cluster is configured with weave, you could: `ip addr` <- and look for weave
or: `ps aux | grep weave` - weaver has the --ipalloc-range configured as an argument
or: `kubectl exec weave-net-jtddl -n kube-system -it -- sh` && `ps aux | grep weaver`

Q: What is the IP Range configured for the services within the cluster?
A: `ps aux | grep kube-api`

Q: What type of proxy is the kube-proxy configured to use?
A:
```
master $ kubectl logs kube-proxy-wptzf -n kube-system
W0106 20:34:57.106692       1 server_others.go:287] Flag proxy-mode="" unknown, assuming iptables proxy
I0106 20:34:57.108242       1 server_others.go:140] Using iptables Proxier.
```
- Turns out the logs have a lot of useful network and assorted info if you're looking for a given spec!

### 170: DNS in Kubernetes
- Whenever a service is created, the k8s DNS service create a record for the service
- So if you have mywebservice you can ping mywebservice and it'll resolve
- If the web service was in a different namespace named apps, you would
- `ping webservice.apps`
- or
- `ping webservice.apps.svc`
- or
- `ping webservice.apps.svc.cluster.local` (FQDN)
- Records for pods are not enabled by default, but they can be enabled.

### 171: CoreDNS for Kubernetes
- Before v1.12, Kubernetes used kube-dns
- CoreDNS is deployed as a ReplicaSet... within a deployment. But all we have to do is look at that pod.
- CoreDNS is configured at `/etc/coredns/Corefile` by default
- The plugin that makes CoreDNS work with Kubernetes is the Kubernetes plugin. You can see cluster.local configured there.
- This is passed as a configmap to coredns
- `kubectl get configmap -n kubesystem`
- If you can to edit the config, you can edit the configmap
- CoreDNS watched the cluster for pods and services and each time one of those is created, it adds a record for it.
- CoreDNS creates a service named kube-dns available to the cluster.
- The IP of this service is configured as the nameserver (`/etc/resolv.conf`) on pods
- The DNS configuration on the pods is done automatically when they're created by Kubelet.
- If you check: `/var/lib/kubelet/config.yaml` you'll find the IPs of the DNS servers on your cluster.

### 172: Practice test - Explore DNS
Q: Identify the DNS solution implemented in this cluster.
A: `kubectl get pods --all-namespaces`

Q: Where is the configuration file located for configuring the CoreDNS service?
A: `kubectl describe pod coredns-78fcdf6894-9n6kw -n kube-system`
```
    Args:
      -conf
      /etc/coredns/Corefile
```

Q: How is the Corefile passed in to the CoreDNS POD?
A: Configmap

Q: What is the root domain/zone configured for this kubernetes cluster?
A: `kubectl describe configmap coredns -n kube-system`

- The rest of the questions were about naming ($SERVICE.$NAMESPACE.svc.$CLUSTER_BASE)
- Ex: mysql.payroll.svc.cluster.local


### 173: Ingress
- When you create a load balancer in gcp, k8s requests a load balancer for google (?)
- This LB has an external IP that can be used to access the app.
- Remember that you need to pay for each LB as they have a public IP!
- The ingress allows users to access your applications based on what HTTP url or HTTP name they reach.
- And at the same time, implement SSL security!
- Think of an ingress as a layer 7 load balancer for HTTP.
- k8s does not come with an ingress controller by default.
- Ingresses available: nginx, contour, haproxy, traefix, istio and GCE (from google)
- The solution you deploy is called an ingress controller.
- The set of rules you configure for this ingress are named ingress resources

- Have a look at 173_nginxingress.yaml and 173_nginxnodeport.yaml
- Please note that missing here is also a configmap required for the nginx config and a ServiceAccount for nginx

- You can route users based on if they reach things like
- http://myapp/url1
- http://myapp/url2

Or if they visit http://myapp.com or http://myotherapp.com

- Don't forget you can:
- `kubectl get ingress`

- You can have multiple rules for a single ingress, regardless of name or url path.

### 174: Practice test - Ingress 1
Q: We have deployed Ingress Controller, resources and applications. Explore the setup.
A: `kubectl get all --all-namespaces`
`kubectl get ingress --all-namespaces`

Q: Which namespace is the Ingress Controller deployed in?
A: `kubectl get all --all-namespaces`
Note that Ingress controller != Ingress resource

Q: What is the name of the Ingress Controller Deployment?
A: `kubectl get deployments`
(nginx-ingress-controller)

Q:
A:

Q:
A:

Q:
A:

Q:
A:

Q:
A:

Q:
A:

### 175: Ingress - Annotations and rewrite target

### 176: Practive test - Ingress 2 
