
### For later
- Make a list of Istio features so you don't forget about the things that are possible, even if you don't implement them
- https://istio.io/latest/docs/concepts/traffic-management/

- Implement most things on this course on your own with your own yaml so you really learn what happened here
- What are the Istio objects you can use with kubectl again? https://istio.io/latest/docs/concepts/traffic-management/
- Can istio be used to limit outgoing bandwitdh for a given deployment?
	- I'm thinking about making mdscanner intentianlly slow just to see if it breaks!

- Remember to list all the features istio has, the ones you've seen on this couse but as well other ones listed on the documentation, so when you do your pet project you can pick which ones to choose.
- Having an overview of all the features is a really nice thing to have too! Knowldge of all the things you can do is POWAH
- Oh damn, multicluster istio meshes are a thing? That looks like an interesting thing to investigate!
	- Can you research k8s multicluster administration as a topic? Is there aything useful tere?
- The istio yamls have a very nice example of using the HPA and limits!
- For later on Istio: Explore the grafana dashboards!
- Try not to make guesses with requests: Look at grafana to see how much CPU and memory your pods are using

---

# Leftover notes from the docs to process:

## Concepts

### Traffic management
- Istio can help with:
- Circuit breakers (timeouts and retries)
- A/B testing
- Canary rollouts / staged rollouts with traffic split by percentages
- Failure recovery!
- You can apply rules to traffic going in and out of your mesh
- You can also add an external dependency to your mesh 

- The resources for Istio traffic management are:
	- Virtual services
	- Destination rules
	- Gateways
	- Service entries
	- Sidecars

### Virtual Services
- Virtual services along with destination rules are the building blocks of istio's traffic management
- A virtual service lets you configure how a requests are routed to a service within an istio service mesh

