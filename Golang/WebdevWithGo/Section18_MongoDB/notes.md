
## 158
```shell
db.dogs.insert({"name":"toby"})
show dbs
show collections
db.dogs.find()
db.dogs.find().pretty()
db.dogs.drop()
db.dropDatabase() 
```

## 159
- Collections can be created implicitly or explicitly

## 160
```shell
use playroom
db
show dbs

db.crayons.insert([
                    {
                        "hex": "#EFDECD", 
                        "name": "Almond", 
                        "rgb": "(239, 222, 205)"
                    }, 
                    {
                        "hex": "#CD9575", 
                        "name": "Antique Brass", 
                        "rgb": "(205, 149, 117)"
                    }, 
                    {
                        "hex": "#FDD9B5", 
                        "name": "Apricot", 
                        "rgb": "(253, 217, 181)"
                    }, 
                    {
                        "hex": "#78DBE2", 
                        "name": "Aquamarine", 
                        "rgb": "(120, 219, 226)"
                    }, 
                    {
                        "hex": "#87A96B", 
                        "name": "Asparagus", 
                        "rgb": "(135, 169, 107)"
                    }, 
                    {
                        "hex": "#FFA474", 
                        "name": "Atomic Tangerine", 
                        "rgb": "(255, 164, 116)"
                    }, 
                    {
                        "hex": "#FAE7B5", 
                        "name": "Banana Mania", 
                        "rgb": "(250, 231, 181)"
                    }
                ])
```

```shell
db.crayons.find()
db.crayons.find().pretty()
db.crayons.drop()
db.dropDatabase()
```

## 161 Find (aka query)
```shell
use store
db
show dbs
db.customers.insert([{"role":"double-zero","name": "Bond","age": 32},{"role":"citizen","name": "Moneypenny","age":32},{"role":"citizen","name": "Q","age":67},{"role":"citizen","name": "M","age":57},{"role":"citizen","name": "Dr. No","age":52}])
db.customers.find()
db.customers.findOne()

db.customers.find({"name":"Bond"})
db.customers.find({name:"Bond"})
db.customers.find({$and: [{name:"Bond"}, {age:32}]})
db.customers.find({$and: [{name:"Bond"}, {age:{$lt:20}}]})
db.customers.find({$and: [{name:"Bond"}, {age:{$gt:20}}]})


db.customers.find({role:"citizen"})
db.customers.find({age:52})
db.customers.find({$and: [{role:"citizen"}, {age:52}]})
db.customers.find({$or: [{role:"citizen"}, {age:52}]})
db.customers.find({$or: [{role:"citizen"}, {age:52}, {name:"Bond"}]})

db.customers.find({name: {$regex: '^M'}})
```

## 162 Update 
```shell
db.<collection name>.update(<selection criteria>, <update data>, <optional options>)
db.customers.find()
db.customers.update({_id:ObjectId("5891221756867ebff44cc886")},{$set:{role:"double-zero"}})
db.customers.update({name:"Moneypenny"},{$set:{role:"double-zero"}})
db.customers.update({name:"Moneypenny"},{$set:{role:"citizen", name: "Miss Moneypenny"}})
db.customers.update({age:{$gt:35}},{$set:{role:"double-zero"}})
db.customers.update({age:{$gt:35}},{$set:{role:"double-zero"}}, {multi:true})

#Save either inserts or overwrites an existing record
db.customers.save({"role":"villain","name":"Jaws","age":43})
db.customers.save({"_id" : ObjectId("6234f37493eb9f3f81c9001a"),"role":"villain","name":"Goldfinger","age":717})
```

## 163 Remove
```shell
db.<collection name>.remove(<selection criteria>)
db.customers.remove({role:"double-zero"})
db.customers.remove({role:"villain"})
db.customers.remove({role:"citizen"},1) #Removes only one entry
db.customers.remove({}) #Removes all
```