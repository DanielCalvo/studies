- Variables were reworked for the 0.12 release!

### Simple types
- String
- Number
- Boolean

### Complex types
- List(type)
- Set(type)
- Map(type)
- Object
- Tuple

### More info
- List: [1,2,5,6]
- Map: {"key" = "value"}
- A list is always ordered
- A set is like a list, but it does not keep the order the elements are instantiated and can only contain unique values
- A list that has [5,1,1,2] becomes [1,2,5] in a set
- An object is like a map, but each element can have a different type, ex:
```
firstname = "Meme"
housenumber = 10
```
- A tuple is like a list but each element can have a different type
- The most common types are list and map, the other ones are use only sporatically
- The other ones you only really need to know if you're doing module development or you're doing plugins
- Remember: String, Bool, number, list, map
- You can also let terraform decide the type of your variables (you don't need to assign it)