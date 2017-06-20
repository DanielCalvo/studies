

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
