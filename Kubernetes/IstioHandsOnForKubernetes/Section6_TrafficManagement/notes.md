### 20. Introducing Canaries
- Canary release
  	- Deploy a new version of a software component
	- But only make that imave "live" for a percentage of the time
	- Most of the time, the old (known working) version is the one being used

### 21. Canaries with Replicas
- No effective way of doing canaries in vanilla k8s (I mean you can do it, but it ain't pretty)
- You can have services with different deployments/pods as targets but it's not very elegant 

### 22. Version Grouping
- Back to Kiali: Workloads graph: Usually deployments (sets of pods)
	- The app graph: Bounds all of that into something it calls an "app"
	- Kiali puts a special meaning on any label called "app"
	- Versioned app graph: If you also supply a label named "version" it'll show that on the graph as well

### 23. Elegant Canaries and Staged Releases
- You can do a canary through Kiali?
- Service > right click > Actions > Weighted routing
- It's not round robin, it's random

### 24. What is an Istio Virtual Service? 
- Messing around with the slider in the previous lecture created:
	- A virtual service
	- A destination rule

#### What is a VirtualService?
- A VirtualService enables us to configure custom routing rules to the service mesh
- Oh a service is just a bunch of DNS entries, interesting
- Kubernetes services are in no way replaced by istio's virtual services
- The virtual service reconfigures network routing on the envoy proxies so that requests go to certain containers based on the rules you set. This is managed by pilot (aka istiod)
- VirtualServices don't replace services -- they do different jobs
- You only need a virtualservice if you really _need_ one. They're not obligatory 

- FOR LATER: Read the istio doc on virtual services 

### 25. VirtualService Configuration in yaml
- spec.hosts: The service name you're applying routing rules to. You can have this be in another namespace, but then you need the FQDN for it

### 26. What is an Istio Destination rule?
- A destination rule is a configuration for a load balancer (virtual service?) for a particular service