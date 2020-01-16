
### 1: Introduction to YAML

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

### 2: Labs - YAML
```yaml
employee:
    name: john
    gender: male
    age: 24
    address:
      city: edison
      state: new jesery
      country: united states
```

### 3: Introduction to JSON PATH
- Yaml uses identation
- JSON uses braces
- They're both similar and can be easilily converted between each other using online tools
- JSON PATH is a query language than when applied to a certain json dataset, gets you results that are a subset of the rest of the data

Sample jsonpath queries:
    - `car.color`
    - `bus.price`
    - `vehicles.car.color`

- The top level element of a json document (which has no name) is known as the root element
- The root element is denoted by a dollar
    - `$.car` 
    - `$.car.color `

- All results of a json path query are encapsulated within an array
- To get a certain element of a list returned by jsonpath:
- `$[0]`
- `$.car.wheels[1].model`

- Any output of a JSON PATH will be available to you within a pair of square brackets
- You can also specify a criteria:
- `$[?(@ > 40)]`
- `$[?(@ in 40,41,42)]`
- `$[?(@ nin 42)]`
- `$[?(@.location == "rear-right").model]`

### 4: Labs - JSON PATH
1. `property1`
2. `$.bus`
3. `$.bus.price`
4. `$.vehicles.car.price`
5. `$.car.wheels`
6. `$.car.wheels[2]`
7. `$.car.wheels[2].model`
8. `$.employee.payslips`
9. `$.employee.payslips[2]`
10. `$.employee.payslips[2].amount`
11. `???`
12. `$.[0]`
13. `$.[0,3]`

### 5: JSON PATH - Part 2, Wildcard
- `$.*.color`
- `$.*.price`
- `$[0].model`
- `$[*].model`
- `$.car.wheels[0].model`
- `$.car.wheels[*].model`
- `$.*.wheels[*].model`

### 6: Labs - Wildcard
1. `$.*`
2. `$.*.color`
3. `$.*.*.price`
4. `$.*.model`
5. `*.*.*.model`
6. `*.*.*.model`
7. `$.employee.payslips.*.amount`
8. `$.*.*.*.*.firstname`
9. `???`

### 7: JSON PATH - Part 3, Lists
- `$[0]`
- `$[3]`
- `$[0,3]`
- `$[0:3]`
- `$[0:8:2]` (START:END:STEP)
- `$[-1]` or `$[-1:0]` (last element, depends on JSONPATH implementation)
- `$[-3:]` (Last 3 elements)

### 8: Labs - JSON PATH - Lists
1. `$[0]`
2. `$[0,4]`
3. `$[0:5]`
4. `$[2:7]`
5. `$[-1:]`
6. `$[6:]`
7. `$[3:9]`
8. `$[-1:]`
9. `$[-3:]`
10. `$[91:97]`
11. `$.[0:5].phone`
12. `$.[-5:].age`

### 9: JSON PATH Use case - Kubernetes
- JSONPATH can be useful when working with k8s to filter through large datasets (imagine thousands of pods)
- To get output in JSON: `kubectl get pods -o json`
- Look at the json
- Create the JSON PATH query based on that JSON to get the output you want, ex:
- `kubectl get pods -o=jsonpath=.items[0].spec.containers[0].image`
- When learning, author recommends getting the output from kubectl in json, then copying it to a jsonpath editor online and playing with the query until you get the output
- Sample query: `kubectl get nodes -o=jsonpath='{.items[*].status.capacity.cpu}'`
- You can combine queries: `kubectl get nodes -o=jsonpath='{.items[*].metadata.name}{"\n"}{.items[*].status.capacity.cpu}'`
- Custom columns: `kubectl get nodes -o=custom-columns=NODE:.metadata.name, CPU:.status.capacity.cpu`
- `kubectl get nodes --sort-by=.status.capacity.cpu`

### 10: Labs - JSON PATH - Kubernetes
6. `$.kind`
7. `$.metadata.name`
8. `$.spec.nodeName`
9.  `$.spec.containers.*`
10. `$.spec.containers.*.image`
11. `$.status.phase`
12. `$.status.containerStatuses[0].state.waiting.*.`
13. `$.status.containerStatuses[1].restartCount`
14. `$.status.containerStatuses[?(@.name == 'redis-container')].restartCount`

### 11: Labs - JSON PATH - Kubernetes - 2
3. `*.metadata.name`
4. `$.users.*.name`