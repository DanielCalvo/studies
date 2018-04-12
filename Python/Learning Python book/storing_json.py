
import json

name = dict(first='Bob', last='Smith')

rec = dict (name=name, job=['dev', 'mgr'])

print (rec)

S = json.dumps(rec)

print (S)

F = open('data.txt', 'w')
json.dump(rec, fp=open('testjson.txt', 'w'), indent=4)

#Let's read that json!
P = json.load(open('testjson.txt'))

print ("Printing P:", P)

