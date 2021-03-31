1. `cd terraform && terraform init && terraform apply`

2. `aws eks --region eu-west-1 update-kubeconfig --name sample-eks-cluster`
	- Or whatever the name of the cluster is

3. `istioctl install --set profile=demo -y`

4. `kubectl label namespace default istio-injection=enabled`

5. `kubectl apply -f k8s`