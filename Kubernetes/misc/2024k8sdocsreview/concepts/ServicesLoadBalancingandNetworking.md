- Every Pod in a cluster gets its own unique cluster-wide IP address. 
- Pods can communicate with all other pods on any node without NAT
- Agents on a node can communicate with all pods on that node

On kubernetes networking:
- Containers within a pod communicate using loopback
- Cluster networking provides communication between different Pods
- There's a service API that lets you expose an application to the outside world using services and ingresses

# Service

A service is a way to expose a networked application that is running as one or more pods in your cluster. You use service to make a set of pods available on the network so that other clients can interact with it

Notes to clarify

- Services have target port and destination port (doesn't have to be the same)
- TCP, UDP and SCTP are supported protocols
- You can use a service without a selector mapped to an EndpointSlice to point to an IP outside the cluster 

There was then a brief explanation on endpoint and endpoint slices that I think is too specific for every day usage.

- Services can have multiple ports! Ports must be named
 
A service can be
- ClusterIP (default of no type is specified)
- NodePort (meh)
- LoadBalancer (exposes service externally if you have a load balancing component installed)
- ExternalName: Maps a service to an external DNS entry, interesting

## Headless service
- You can have a headless service with `.spec.clusterIP` set to `None`
- An IP is not allocated. No load balancing is done.
- With selectors: This type of service seems to be used for when you want to point to a single pod specifically and that's it (so no need for an extra service IP)
- Without selectors: A DNS entry is created. Used for external services it seems.

---

There are two ways for service discovery: 
- Environment variables on the node managet by kubelet
- DNS sentries (the preferred way as confirmed by the docs)
 
There is also some type of `.spec.externalIP` functionality but this is not very well documented in the service doc page

# Ingress Controllers
# Gateway API
# EndpointSlices

# Network Policies
Network policies allow you to control traffic at the IP or TCP/UDP/SCTP level. NetworkPolicies allow you to specify how a pod is allowed to communicate with various network entities.

Entities that a pod can communicate with are identified through a combination of the following three:
- Other pods
- Namespaces
- IP blocks

You can isolate a pod of egress and ingress. By default pods are **not** isolated for egress (all outbound traffic goes).
By default, pods are alot **not** isolated by ingress either.

## NetworkPolicu resource
Here's an example. Note the `namespaceSelector` and `podSelector` selectors as well as the `ipBlock` fields.

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: test-network-policy
  namespace: default
spec:
  podSelector:
    matchLabels:
      role: db
  policyTypes:
    - Ingress
    - Egress
  ingress:
    - from:
        - ipBlock:
            cidr: 172.17.0.0/16
            except:
              - 172.17.1.0/24
        - namespaceSelector:
            matchLabels:
              project: myproject
        - podSelector:
            matchLabels:
              role: frontend
      ports:
        - protocol: TCP
          port: 6379
  egress:
    - to:
        - ipBlock:
            cidr: 10.0.0.0/24
      ports:
        - protocol: TCP
          port: 5978
```

And that's about it for network policies, they're rather simple compared to other featires. Nice!

## Things you cannot do with NetworkPolicies
- Force internal traffic to go through a gateway
- TLS related anyhthing
- Node specific policies
- Targeting services by name
- Creating or managing policy requests (?)
- Default policies applied to all namespaces/pods
- Advanced policy querying (?) and reachability tooling (?)
- Network event logging (like refused connections)
- The ability to explicitly deny policies (? like, denying the application of policies themselves?)
- The ability to prevent loopback or incoming traffic access. Pods cannot blog localhost access.

# DNS for Services and Pods
# IPv4/IPv6 dual-stack
# Topology Aware Routing
# Networking on Windows
# Service ClusterIP allocation
# Service Internal Traffic Policy