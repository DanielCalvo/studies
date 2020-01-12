
### 210: Application Failure
- Remember to have the traffic flow of your application mapped out before starting any troubleshooting effort
- Start with the frontend: ingress, service nodeport, then pods, etc
- Compare selectors yo!

- Remember to check if the pod is running. The number of restarts can be a good indicator of what's happening if not all is going smoothly
- Also helpful:
    - `kubectl describe pod web`
    - `kubectl logs web`
    - `kubectl logs -f web`
    - `kubectl logs -f --previous web`
    
- Also remember: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/

### 211: Practice test - Application Failure
My attention to detail is lacking. Pay special attention to the information that the question gives you, the answer might be related!

### 213: Control plane failures
- `kubectl get nodes`
- `kubectl get pods -n kube-system` 

If the components on the cluster were not deployed with kubeadm and are not running as containers, check their services:
- `service kube-apiserver status`
- `service kube-controller-manager status`
- `service kube-scheduler status`

On worker nodes:
- `service kubelet status`
- `service kube-proxy status`

To see logs, if the containers are pods:
- `kubectl logs kube-apiserver-master -n kube-system`
- `journalctl -u kube-apiserver`

Also check: https://kubernetes.io/docs/tasks/debug-application-cluster/debug-cluster/

### 216: Worker node failure
- `kubectl get nodes`
- `kubectl describe node worker-1`
- Node have conditions and flags!
- A bunch of unknown things can indicate the loss of a node
- Check the status of the kubelet service, the kubelet logs, check the age of the kubelet certificates