
#Dictionaries in other languages may be called hash maps (or hashes) or associative arrays

#Dictionary: Two linked values where they key is a unique identifier where we can find our data, and the value is that data.
#Almost like a real physical dictionary!

student = {'name': 'Daniel', 'age': '28', 'courses': ['Math', 'CompSci']}

#Prints everything:
print(student)

#Getting a single key:
print(student['name'])
print(student['courses'])

#Keys can be any immutable data type

#print(student['phone']#This will raise an exception

print(student.get('age'))


print(student.get('phone')) #This returns none instead of error!

print(student.get('phone', 'Not found! :o')) #Default value for keys that don't exist!

student['phone'] = '12345-5555'
print(student.get('phone', 'Not found! :o')) #Default value for keys that don't exist!

student.update({'name': 'Jane', 'age': 23, 'phone': '4444444'}) #Neat!

del student['age']

student.update({'age': '33'})

age = student.pop('age')
print(age)

print(len(student)) #How many keys and values pairs do we have?
print(student.keys())
print(student.values())
print(student.items()) #Keys and values

#Look through keys and values!

for key, value in student.items():
    print(key, value)

