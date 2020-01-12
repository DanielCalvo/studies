### 201: Deploy a Kubernetes Cluster with Kubeadm 
- Reference doc:  https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
- Doc is pretty straightforward


###202: Practive test - Deploy a Kubernetes cluster with Kubeadm
Q: Install the kubeadm package on master and node01
A: Ambiguous question, but the author meant:
```
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

Q: What is the version of kubelet installed?
A:
`kubelet help | grep version`
`kubelet --version`

Q: How many nodes are part of kubernetes cluster currently? Are you able to run kubectl get nodes?
A: No, so 0.
   
Q: Initialize Control Plane Node (Master Node)
A:
```
kubeadm init
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```   
Q: Join node01 to the cluster using the join token
A:
```
kubeadm join 172.17.0.28:6443 --token bcv3g4.7tagm7mna2g965j7 \
    --discovery-token-ca-cert-hash sha256:da6d2f848ca977368493316f24194e9517dcbdb295471f7d05dbf1dbc414a69f
```

Q: Install a Network Plugin. As a default, we will go with weave- Refer to the official documentation for the procedure
A: `kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"`

