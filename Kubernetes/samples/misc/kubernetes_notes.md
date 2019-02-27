

#### Resources:
https://github.com/kubernetes-up-and-running/examples  
https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/

Goal 1: Launch a pod

Goal 2: Get a pod to connect to a database

Goal 3: Get an ingress controller 


This launches the nginx hello world in a pod:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-hello
spec:
  containers:
  - image: nginxdemos/hello:latest
    name: kuard
    ports:
    - containerPort: 8080
      name: http
      protocol: TCP
```

Let's put this pod in a ReplicaSet:

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: frontend
  labels:
    app: guestbook
    tier: frontend
spec:
  replicas: 10
  selector:
    matchLabels:
      tier: frontend
  template:
    metadata:
      labels:
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: nginxdemos/hello:latest
```

Let's put this pod in a Deployment:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 10
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx-from-deployment
        image: nginxdemos/hello:latest
        ports:
        - containerPort: 80
```

Let's configure an Ingress to reach that deployment:

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: test-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - http:
      paths:
      - path: /testpath
        backend:
          serviceName: test
          servicePort: 80
```



I would like to have an ingress service that allows my pod to be reached from the internet.

