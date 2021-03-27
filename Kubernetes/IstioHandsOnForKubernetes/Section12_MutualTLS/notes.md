### 40. Why is encryption needed inside a cluster
- Mutual TLS!
- Any project deploying to multiple nodes could benefit from multiple TLS!
- A few lines of yaml that you can apply and author claims a massive win! (uh-oh)
- Author goes over https briefly
- All pods communicate with each other using HTTP
- Q: Do you care that the traffic between your pods is unencrypted? 
- Author says: You should care!
- Author argues that traffic between pods could be intercepted by a bad actor in your datacenter
- If you nodes in multiple AZs, the traffic between your pods might "leave buildings" and might be "snoopable"
- Even if you think that your pod to pod traffic will neaver leave a building, assume it will!

### 41. How Istio can upgrade traffic to TLS
- Remember that the proxy on istio is part of the same pod!
- Remember that two containers on the same pod will always be on the same node and share the same network space
- Encrypted communication happens in between the envoy proxies
- istio-citadel was a standalone thing but now it's part of istiod, it managed the certificates for the inbetween pod TLS
- Author says: istiod/citadel appears to not be exclusive to HTTP traffic, it could also work with gRPC or TCP, but tnat seems more advanced
- We can enforce a policy:
	- Tell istio to block any non TLS traffic (between the proxies)
	- Automatically configure any proxy to proxy communication to use mTLS

### 42. Enabling mTLS - it's Automatic
- Oh wait! As of the current version of Istio, mTLS is enabled by default! There's nothing that needs to be done! Awesome!
- If you toggle on Kiali to display a "security" setting, it will show little padlocks in between the apps, that tells us that TLS traffic is going through the microservices!

### 43. Sctrict vs Permissive TLS
- By default, istio uses "Permissive mTLS", so if a pod that is on a namespace that does not use istio (therefore no proxy and no mTLS), istio will accept this connection 
- Istio will not upgrade this connection. The call to the service in the istio namespace with mTLS will succeed, but this call will be insecure

- There is a setting we can apply to the istio mesh called *strict* mTLS
- With strict mTLS, if we have an incomming connection from a client that is unencrypted and cannot be upgraded, then the connection will be rejected. Setting mTLS to strict blocks TLS traffic
- To enable strict TLS you just apply a k8s object of type "PeerAuthentication"

### 44. Strict mTLS works in both directions
- If you initiate an HTTP call to a pod that is not on an istio'd namespace and does not have mTLS enable, and you have strict mTLS enabled, this request would fail
- So strict mTLS forces both incoming requests from other pods to use HTTPS, but also forces outgoing requests to other pods to also use HTTPS!