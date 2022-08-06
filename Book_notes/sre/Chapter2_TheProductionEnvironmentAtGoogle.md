## Chapter 2 - The Production Environment at Google, from the Viewpoint of an SRE
- Google uses Borg, which was a precursor to Kubernetes
- They use fancy networking switching
- They have a cluster wide filesystem 
- They also use Borgmon, a precursor to Prometheus
- Google uses monorepo (although you could just create PRs on other repos)
- They use global load balancers and have servers worldwide