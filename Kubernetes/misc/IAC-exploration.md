## What if I don't wanna use terraform to create k8s clusters?

### eksctl
Actually very cool, seemed one of the most mature tools. You can create clusters declaratively using yaml and manage them through gitops if you want. Docs were great!

### Crossplane
Moderately cool. Limited support for AWS resouces (ex: no Lambda, no ECS). Documentation does not seem to be the most generous, mostly an API reference for the yamls

### Pulumi
Actually very cool! Support for most resources and docs seems good. But then everyone's gotta program. I should get around creating a sample config sometime...

### Cluster API?
Very rough around the edges. Had to use `clusterawsadm` and then `clusterctl` to get a set of yamls for AWS. EKS support is still experimental. Certain autoscaling features are still experimental. Docs are... okay but not amazing. Maybe this is one for later! 
    - https://cluster-api-aws.sigs.k8s.io/
    - https://github.com/kubernetes-sigs/cluster-api-provider-aws

### Terraform
Clunky. It has no if statement. It can fail if data sources are not present. It's cumbersome to use sometimes, it's not the most elegant... but it works. The docs are good. Everyone knows it.

