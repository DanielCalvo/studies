
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
- The top level element of a json document (which has no name) is known as the root element
- The root element is denoted by a dollar
- All results of a json path query are encapsulated within an array
- Any output of a JSON PATH will be available to you within a pair of square brackets
- You can also specify a criteria:
- `$[?(@ > 40)]`
- `$[?(@ in 40,41,42)]`
- `$[?(@ nin 42)]`
- `$[?(@.location == "rear-right").model]`