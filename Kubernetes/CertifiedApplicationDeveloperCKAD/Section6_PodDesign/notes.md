### 82. Labels, selectors and annotations
- Labels allow you to group things together and filter them based on your needs!
- metadata.labels
- kubectl get pods --selector app=myapp
- A replicaset will have 2 sets of labels: for itself and for the underlying pods
- annotations seem to be for miscelaneous metadata

### 85. Rolling Updates & Rollbacks in Deployments
- RollingUpdate is the default deployment strategy
- The alternative is Recreate, which recreates everything at once and results in downtime
- Meh I wanted to learn more!