### 220: Pre-requisite: JSON PATH

#### YAML:

- Key value pairs
```yaml
Fruit: Apple
Vegetable: Carrot
Liquid: Water
Meat: Chicken
```

- Array/Lists:
```yaml
Fruits: 
  - Orange
  - Apple
  - Banana
Vegetables:
  - Carrot
  - Cauliflower
  - Tomato
```

- Dictionary/Map:
```yaml
Banana:
  Calories: 105
  Fat: 0.4g
  Carbs: 27g
Grapes:
  Calories: 62
  Fat: 0.3g
  Carbs: 16g
```
- Spacing determines which values are parts of which values

Key Value / Dictionary lists

```yaml
Fruits:
  - Banana:
      Calories: 105
  - Grape:
      Calories: 62
```

- Dictionaries are unordened, lists are ordered

### 221: Pre-requisites - JSON PATH
- I did a mini-course on this! Notes are at [./jsonpathcourse/course.md](./jsonpathcourse/course.md)

### 222: Advances kubectl commands
- This was covered on the mini-course, see lecture 221

### 223: Practice test - Advanced Kubectl Commands
Q: Get the list of nodes in JSON format and store it in a file at /opt/outputs/nodes.json
A: `kubectl get nodes -o json > /opt/outputs/nodes.json`

Q: Get the details of the node node01 in json format and store it in the file /opt/outputs/node01.json
A `kubectl get node node01 -o json > /opt/outputs/node01`

Q: Use JSON PATH query to fetch node names and store them in /opt/outputs/node_names.txt
A: `kubectl get nodes -o jsonpath={.items.*.metadata.name} > /opt/outputs/node_names.txt`

Q: What's the image of the current OS on nodes?
A: `kubectl get nodes -o jsonpath={.items.*.status.nodeInfo.osImage} > /opt/outputs/nodes_os.txt`

Q: A kube-config file is present at /root/my-kube-config. Get the user names from it and store it in a file /opt/outputs/users.txt 
Use the command kubectl config view --kubeconfig=/root/my-kube-config to view the custom kube-config
A: `kubectl config view --kubeconfig=/root/my-kube-config -o jsonpath={.users.*.name} > /opt/output/users.txt`

Q: A set of Persistent Volumes are available. Sort them based on their capacity and store the result in the file /opt/outputs/storage-capacity-sorted.txt
A: `kubectl get pv -o jsonpath={.items.*.spec.capacity.storage}` <- close, but wrong
`kubectl get pv --sort-by=.spec.capacity.storage > /opt/outputs/storage-capacity-sorted.txt`

Q: That was good, but we don't need all the extra details. Retrieve just the first 2 columns of output and store it in /opt/outputs/pv-and-capacity-sorted.txt
The columns should be named NAME and CAPACITY. Use the custom-columns option. And remember it should still be sorted as in the previous question.
A: You need to do a loop with custom output for this one. This one is tough! 

Q: Use a JSON PATH query to identify the context configured for the aws-user in the my-kube-config context file and store the result in /opt/outputs/aws-context-name.
A: This requires a query, aka `[?()]`. This one is tough! 

### Extra: JSON Path practice
- `kubectl config view -o jsonpath='{range .users[*]}{.name}{"\n"}{end}`
- `kubectl get pod kube-dns-79868f54c5-7jxxf -o jsonpath='{range .spec.containers[*]}{.name}{"\n"}{end}'`
- `kubectl get pod kube-dns-79868f54c5-7jxxf -o jsonpath='{.spec.containers[?(@.name == "kubedns")].image}'`

Q: Sort all containers on the cluster by alphabetic order
A:
`kubectl get pods -o jsonpath="{range .items[*]}{.metadata.name}{'\n'}{end}"`
`kubectl get pods -o jsonpath="{range .items[*]}{range .spec.containers[*]}{.name}{'\n'}{end}{end}" --sort-by=.metadata.name`

Q: Show me the name of containers (not pods) which have a resource requirement of 100m CPU
A:

Q: Show the me resource requirements of all pods
A:

Q: Show me all pods sorted by creation timestamp
A: `kubectl get pods -o custom-columns=NAME:.metadata.name,CREATED_AT:.metadata.creationTimestamp`

Q: List all the images of all the pods in the initialized state