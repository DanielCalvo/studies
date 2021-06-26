
### 52. In place upgrades
- Upgrading Istio!
- Author says: Upgrading istio is awkward and is risky! Istio is not just a feature or a pod, istio is infrastructure and it's present everywhere in the cluster!
- If something goes down, downtime can happen!
- According to the docs, there are two ways of upgrading istio: Canary upgrades and in-place upgrades. Author says: There is yet another method, which is the one the one he uses!
- In place upgrade: Easiest, single command. But it is also very risky! If things go poorly, you break everything!

### 53. Canary Upgrades (Rolling Upgrades)
- In place upgrades are dangerous! Author describes how you can have mismatching versions of the istio proxy and istiod
- Canary upgrades are the recommended way forward: They upgrade isti ofor a small percentage of workloads 
- `istioctl install --set profile=demo --set revision=1-7` <- This command runs with version 1.7 of the istio executable
- `istioctl proxy-status` <- Shows to which istiod the proxies are connected to
- `istioctl install --set profile=demo --set revision=1-8` <- This command runs with version 1.8 of the istio executable
- `kubectl get pods -n istio-system` <- Two istiods!
- Deployed pods will still be using the 1-7 sidecar though (check with proxy-status)
- If you kill a pod running 1-7 and it spawns again, it'll spawn using the 1-8 version of istio. Neat!
- To revert new pods to 1-7, just change the `istio.io/rev: 1-7` label to 1-7.
- To remove the older version: `istioctl x uninstall --revision=1.7`

### 54. Live Cluster Switchovers (Alternative to the official upgrade paths)
- Author provisions a new cluster with a new version of istio and has the domain point to that, or some other solution in front of it (like cloudflare) and then manages the traffic between the two clusters there. It's like a canary release, but for clusters!
- DNS switchovers are slow though :(

### 55. Goodbye
- Istio is cool! Circuit breaking, fault injection and mutual TLS are cool features, on top of the ingress and egress gateways, virtual services and tracing.
- Istio is complex though
- Official documentation needs to be improved a bit