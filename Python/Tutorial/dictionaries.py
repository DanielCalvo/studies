import math

D = {'food' : 'Tuna', 'quantity' : 42, 'color' : 'brown'}

print (D['food'])

D['quantity'] += 1

print (D['quantity'])

Person = {}

Person['name'] = 'Bob'
Person['job'] = 'rich, ain\'t got no need for no jahb'
Person['age'] = '25'

print (Person['job'])

Person = dict(name='Bahb', job='wandering rich man', age=26)

#This overwrites the previous content of the dictionary
Person = dict(address = 'whatevs')
print (Person)

bob2 = dict(zip(['name', 'job', 'age',], ['bob', 'dev', 21]))
print (bob2)

bob3 = {'name': {'first': 'Bob', 'second': 'Bobson'},
        'jobs': ['dev', 'ops guy']}

print (bob3['name']['second'])
print (bob3['name'])
bob3['jobs'].append('boss of everything')
print (bob3['jobs'])

D = {'a': 1, 'ddd': 90, 'b': 2, 'c': 3}
D['e'] = 123
print (D['e'])

if not 'f' in D:
    print ('f in D is missing!')

value = D['x'] if 'x' in D else 0
print (value)

K = list(D.keys())
K.sort()
print (K)

for key in K:
    print (key, '=>', D[key])

print ("This output is getting confusing!")

for key in sorted(D):
    for char in key:
        print (char.upper())

x = 5
while x > 0:
    print (x, "weeee")
    x -= 1

#Wew look at this nexting:

square_roots = [math.sqrt(x) for x in [1, 9, 16, 25, 33, 43] ]
for square_root in square_roots: print (square_root)

square_roots = [math.sqrt(x) for x in [1,3,4,5,6]]
print (square_roots)
