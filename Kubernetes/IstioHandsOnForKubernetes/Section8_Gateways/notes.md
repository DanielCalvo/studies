### 29. Why do I need an Ingress Gateway?
- Author says: Imagine now there's a canary requirement that will be deployed to handle 10% of traffic for your app
- Doing it with virtual services didn't work in his demo, hmm

### 30. Edge Proxies and Gateways
- You don't go through any proxies (envoy) on the ingress! :o
- An edge proxy, according to the envoy docs: An instance of envoy that sits on the edge of your cluster, configured to listen to all incoming requests from the outside
- Istio Gateway: Allows you to configure an edge proxy (called Istio Gateway)
- Note from Dani: You need to implement all of this yourself, the explanation the author gives is very lenghty

### 31. Prefix Based Routing
- You can have one set of pods available at `/` and another one at `/experimental` and route 
	- This is routing based on the prefixes instead of the weighted canaries
- Oh you can also route based on the scheme! (if it's HTTP or HTTPS) ans based on the http request (if GET or POST)
- If the app you're using is sensible to paths, they may not work very well, or you might need to configure it around that

### 32. Subdomain Routing
- You can have various VirtualService rules with many domains and/or wildcards
- Using subdomains for deploying is better is it's easier to adjust certain applications to that, and other applications may not might that entirely