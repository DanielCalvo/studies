
print('Hello world! :D')

message = 'Bobby\'s world'
message = "Bobby's world"
message = 'Hello World'

print(len(message))
print(message.__len__())
print(message[0])

#First index is inclusive, last one is not :o
print(message[0:5])
print(message[:5])
print(message[6:]) #This is called slicing

#A method is just a function that belongs to an object!

print(message.lower())
print(message.upper())

print(message.count('Hello'))
print(message.count('l'))
print(message.find('world'))

message = message.replace('World', 'Universe') #strings are immutable
print(message)

greeting = 'Hello'
name = 'Dani'

message = greeting + ', ' +name + '. Welcome!'
print (message)

message  = '{}, {}. Welcome!'.format(greeting, name)
print(message)

message  = f'{greeting}, {name}. Welcome!'
message  = f'{greeting}, {name.upper()}. Welcome!'

print(message)

#print(help(str))
print(help(str.lower))
