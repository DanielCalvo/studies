### 34. Header based routing
- You can route based on HTTP headers. Neat!

### 35. Dark releases for all microservices
- You can deploy a new version of your service in live, and the only people who can access it are the ones who pass a header. Neat!
	- That's called a "dark release"
- It's private, but it's in the production cluster!
- All you need to do to release is switch the routings around and increment the number of replicas on dark release version of the microservice
	 - Uuuuh probably
- You can use Kiali to generate the istio yaml/rules for you if that helps

- Remember to enable header propagation! If the service you want to do a dark release for is deep in your service mesh, the header has to get to it
- Istio does not do this propagation for you. So if you have a header named "X-experimental" your app needs to propagate that all the way