

if True:
    print("Well I guess it's true")
if 1:
    print("I guess that's true too!")

x = "dog"

if x == "wolf":
    print("wolf!")
elif x == "rabbit":
    print("it's a rabbit!")
elif x == "mouse":
    print("is it a mouse?")
else:
    print("I've got no idea what x is")

choice = 'ham'

print ({'spam': 1.25,
        'ham': 1.50 }[choice])

branch = {'spam' : 10, 'bacon': 20, 'tuna': 30}

print(branch['tuna'])
choice = 'bacon'

if choice in branch:
    print(choice, branch[choice])
else:
    print("choice not available!")
try:
    print(branch['banana'])
except KeyError:
    print("bad choice")

#some sample identatoin

x = 1
if x:
    y = 2
    if y:
        print('block2')
    print('block1')
print('block0')

x = 'SPAM'
if 'laughter' in 'slaughter':
    print('lulz')
    x += 'NI'
    if x.endswith('NI'):
        print(x)

#awkward multiline example

if (1 and 2
    and 3
    and 4):
    print('yeah')

#This type of multiline work with backlashes is best avoided:
x = 1 + 2 + \
    4
print(x)

S = """
aaaa
bbbb
ccc
"""

print(S) #Adds a bunch of newlines

S = ('aaa'
     'bbb'
     'ccc')

print(S)

#Useful for single statements
if 1: print('printing on one!')

if 0:
    print('Naaaaah')
else:
    print('I\'m pretty sure zero is false')

X = 1
Y = 0 # means false

if X and Y:
    print('X AAAAND Y are true')
if X or Y:
    print('X OOOOR y are true')

if Y == False:
    print('Yeah Y is false')
if Y == 0:
    print('Yeah Y is false')
if 0 == False:
    print('Who would\'ve thought! 0 is false too!')

if 0 and 2 and 2 and 3:
    print('no way all of this is going to be true')
else:
    print('Python evaluates the operands from left to right and stops of the left operand is a false object as it determines the result (if one is false in this if, all is false')

A = 'yes' if 1+1 != 3 else 'no'
print(A)

A = 'yes' if '' else 'no'
print(A)

X = 1
Y = 2
Z = 3

A = ((X and Y) or Z)
print(A)
A = Y if X else Z
print(A)

print(['False', 'True'][bool('')])
print(['False', 'True'][bool('not empty')])

X = 'A' or default
print(X)

L = [1, 2, 3, 0 , 1 ,'', 'asd']

#Interesting ways to check for values that are True on a list:
print([x for x in L if x])
print(list(filter(bool, L)))
print(any(L))
print(all(L))

A = 1 and 0 and 2
print(A)

print('hello'
      'world')

A = 1 + 2 \
    + 3