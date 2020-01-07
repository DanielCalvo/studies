### 139: Storage in Docker
- Docker stores it's data by default under:
- /var/lib/docker
    - aufs
    - containers
    - image
    - volumes
- When docker builds an image, it builds it in a layered architecture
- Each line in a Dockerfile is a layer

- Docker uses layers from the cache when building images
- What is really cool is that docker can use layers from the cache of other images when building images that have the same layers 

121: Storage:
Section keeps scope to Kubernetes only
(Follow up from Dani: Can you follow up with a google search for "storage for kubernetes"?)

### 140: Volume Drivers Plugins in Docker
- By default, Docker uses the local volume driver
- But there are volume drivers for Azure File STorage, Digial Ocean Block Storage, gce-docker, GlusterFS, VMware vSphere Storage and many others
- You can specify the volume driver when running a docker container with `--volume-driver`

### 141: Container Storage Interface (CSI)
- The Container Runtime Interface specifies how a container orchestrator tool (like k8s) communicates with container runtimes (such as Docker)
- There is also the Container Networking Interface (CNI) - Which extends support to different networking solutions

### 143: Volumes
- We can attach a volume to a pod, and whatever data is generated in that volume, persists.
- You can specify a bunch of options we'll see later. You can specify among other things a hostPath.
- It is not recommended to use hostPath on volumes, as you'll be storing data on the host.
- Your pods would expect the data in there to be consistent across all hosts, which won't happen
- Kubernetes supports a bunch of storage solutions: NFS, GlusterFS, Flocker, Ceph, ScaleIO, and vendor ones from AWS and GCP.


### 144: Persistent volumes
- Regular volumes would have to configure manually on every single pod yaml on the cluster.
- Instead you can manage storage centrally. A persistent volume is a cluster wide pool of storage volumes configured by an administrator, to be used by users deploying applications on the cluster.
- The users can choose storage from this pool using persistent volume claims.

- Let's create a persistent volume! Take a look at [./144_pv.yaml](./144_pv.yaml)
- `kubectl apply -f 124_pv1.yaml`
- `kubectl get persistentvolume`


### 145: Persistent volume claims (PVCs)
- An admin creates a PV
- An user creates a PVC to use the storage
- Once a PVC is created, Kubernetes binds the PVC to a PV based on the properties of the request.
- Every PVC is bound to a single PV.
- You can still use labels and selectors in case you want to bind your PVC to a specific PV.
- If there are no volumes available, the PVC claim will remain pending until new volumes are made available.
- Have a look at `145_pvc.yaml`
- `kubectl create -f 145_pvc.yaml`
- `kubectl get persistentvolumeclaim`
- There is a persistentVolumeReclaimPolicy. By default it is retain, which means your PVC will stay there until it's manually deleted.

### 146: Using PVCs in PODs
- Once you create a PVC use it in a POD definition file by specifying the PVC Claim name under persistentVolumeClaim section in the volumes section like this:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
    - name: myfrontend
      image: nginx
      volumeMounts:
      - mountPath: "/var/www/html"
        name: mypd
  volumes:
    - name: mypd
      persistentVolumeClaim:
        claimName: myclaim
```

### 147: Practice Test - Persistent Volumes and Persistent Volume Claims
Q: The application stores logs at location /log/app.log. View the logs.
A: `kubectl exec webapp tail /log/app.log`

Q: Configure a volume to store these logs at /var/log/webapp on the host
A:
```yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: 2020-01-05T19:24:13Z
  name: webapp
spec:
  containers:
  - env:
    - name: LOG_HANDLERS
      value: file
    image: kodekloud/event-simulator
    imagePullPolicy: Always
    name: event-simulator
    volumeMounts:
      - mountPath: /log
        name: myvolume
  volumes:
    - name: myvolume
      hostPath:
        path: /var/log/webapp
        type: Directory
```

Q: Create a 'Persistent Volume' with the given specification.
A: 
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-log
spec:
  accessModes:
    - ReadWriteMany
  capacity:
    storage: 100Mi
  hostPath:
    path: /pv/log
```

Q: Let us claim some of that storage for our application. Create a 'Persistent Volume Claim' with the given specification.
A: 
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: claim-log-1
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 50Mi
```

Q: Update the Access Mode on the claim to bind it to the PVC. Delete and recreate the claim-log-1
A: Edited `ReadWriteOnce` to `ReadWriteMany` on the pvc! 
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: claim-log-1
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 50Mi
```

Q: Update the webapp pod to use the persistent volume claim as its storage. Replace hostPath configured earlier with the newly created PersistentVolumeClaim
A:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: webapp
spec:
  containers:
  - name: event-simulator
    image: kodekloud/event-simulator
    env:
    - name: LOG_HANDLERS
      value: file
    volumeMounts:
    - mountPath: /log
      name: log-volume
  volumes:
  - name: log-volume
    persistentVolumeClaim:
      claimName: claim-log-1
```
