
### 5. Getting istio running
- Autho skipped using istioctl in favour of manually applying yamls
- A CRD is an extension of the kubernetes API, with a CRD you can invent your own k8s objects

### 6. Enabling Sidecar Injection
- Did a kubectl apply -f of all the yamls for the course
- Then did: `kubectl label namespace default istio-injection=enabled` to enable sidecar injection on the default namespace

### 7. Visualizing the system with Kialo
- Let's go!