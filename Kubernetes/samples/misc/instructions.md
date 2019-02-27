

##### Step 1: Deploy a web application with a deployment:

```yaml
apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```

##### Step 2: Expose your Deployment as a Service internally
Service has to be NodePort or LoadBalancer type  
```
kubectl expose deployment nginx-deployment --target-port=80 --type=NodePort
```

Step 3: Create an ingress resource:
```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: nginx-deployment-ingress
spec:
  backend:
    serviceName: nginx-deployment
    servicePort: 80
```




kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user $(gcloud config get-value account)
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml























### Google's thing  
```
kubectl run web --image=gcr.io/google-samples/hello-app:1.0 --port=8080  
kubectl expose deployment web --target-port=8080 --type=NodePort
```
Ingress:
```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: basic-ingress
spec:
  backend:
    serviceName: web
    servicePort: 8080
```
```yaml
kubectl apply -f basic-ingress.yaml
```

Ingresses take about 5 minutes to take effect, so be a bit patient!

















