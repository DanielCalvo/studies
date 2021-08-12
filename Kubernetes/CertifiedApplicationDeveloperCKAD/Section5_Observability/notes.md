### 73. Readiness and Liveness Probes
- A pod has a pod status and a pod condition
#### Pod status
- Pending: When it is first created, it is pending, the scheduler looks where to place it. If it can't find where to place it, the pod remains pending
    - To find out why it's pending, describe the pod
- ContainerCreating: Image is being pulled and pod is creating
- Running: All is gucci and it's running (until terminated)

### Pod Conditions
- Conditions complement pod status
- They're an array of boolean values 
    - PodScheduled
    - Initialized
    - ContainersReady
    - Ready 

- By default, k8s assumes that as soon as a container is created, it is ready to receive traffic, so it sets the value of ready to true

### Readiness Probes
- HTTP test: /api/ready
- DB: TCP test, 3306
- You can also just exec a command inside the container 

- When a readiness probe is being used, k8s does not set the ready condition immediately to true, instead, it performs a test!
- readinessProbe options: httpGet, tcpSocker and exec command
- There's also an additional delay setting available

### 74. Liveness Probe
- Liveness probes can be configured in the container to test if the application inside the container is actually healthy
- If the test fails, the container is considered unhealthy and it is recreated
- You have the same config options as the Readiness probe!

### 76. Container logging
- kubectl logs -f mypod 
- kubectl logs -f mypod mycontainer

### 79. Monitor and Debug Applications
- There's a metrics server named "metrics server"
- cAdvisor collects metrics from pods and makes them available through kubelet for the metrics server. Neat!
- `minikube start`
- `minikube addons enable metrics-server`
- `kubectl top node`
- `kubectl top pod`