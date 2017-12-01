
print(open('script2.py').read())

print()

for line in open('script2.py'):
    print(line.upper(), end='')

#Using file.readlines() is considered less than ideal to read files, as it loads the entire file into memory!
#This may cause the program not to run if you load files that are too big into memory.

#You can also iterate through a file with a while loop!

f = open('script2.py')

while True:
    line = f.readline()
    if not line: break
    print(line.upper(), end='')

#The book says: Iterators (for loop) run at C language speed. The while loop runs Python byte code through the Python virtual machine, so it tends to be slower.

print('\n')
f = open('script2.py')
print(f.__next__(), end='')
print(f.__next__(), end='')
print(f.__next__(), end='')

#There's a bunch of theory in the book about iterable objects and iterator, hopefully it becomes clearer as we code it out:

A = [1,2,3]
B = iter(A) #Obtain an iterator object from an iterable
print(B.__next__()) #Call iterator's next to advance to the next item
print(B.__next__())
print(B.__next__())
#print(B.__next__()) #If I try to do this I break the program, huh?

print('hi?')

print('yes')

f = open('script2.py')
if __name__ == '__main__':
    if iter(f) is f:
        print('yes')
    else:
        print('no?')

#Lists and other built in objects however, are not their own iterators as they support multiple open interactions (there can be more than one iterator, in more than one place)
#File iterators as seen above can't seek backwards or support multiple active scans, so they are their own iterators

L = ['a','b','c']
if iter(L) is L:
    print ('yes')
else:
    print('no')

#Automatic iteration:
L = ['a','b','c']
for a in L:
    print(a, end=' ')
print()

#Manual interation:
L = ['a','b','c']
I = iter(L)

while True:
    try:
        X = next(I)
    except StopIteration:
        break
    print(X, end=' ')
print('\n')

D = {'a':1, 'b':2, 'c':3}

for key in D.keys():
    print(key, D[key])

I = iter(D)
print(next(I))
print(next(I))
print(next(I))

import os

#Other things are iterable too!
P = os.popen('dir')
print(P.__next__())
print(P.__next__())
print(P.__next__())
print('\n')

R = range(5)
print(R) #Interesting, doesn't iterate
I = iter(R)
print(I.__next__())
print(I.__next__())
print(I.__next__())
print(next(I)) #appears to be the same thing!
print('\n')

#R is not a list, but can I iterate?
for item in R:
    print('printing item:', item) #Looks like I can

#But we can make it a list!
R = list(range(5))
print(R)

for item in R:
    print('printing item:', item)

E = enumerate('dani')
I = iter(E)
print(next(I))
print(next(I))
print(next(I))

print(list(enumerate('dani')))
print(list('dani'))

#old school:
L =[1,2,3,4,5]

for i in range(len(L)):
    L[i] += 10
print(L)

L = [x + 10 for x in L]
print(L)

f = open('script2.py')
lines = f.readlines()
print(lines)

lines = [line.rstrip() for line in lines]
print(lines)

lines = [line.rstrip() for line in open('script2.py')]
print(lines)

#These are so neat!
print([line.upper() for line in open('script2.py')])
print([line.rstrip().upper() for line in open('script2.py')])
print([line.rstrip().upper().split() for line in open('script2.py')])
print([line.replace(' ', ',') for line in open('script2.py')])
print([('sys' in line, line[:5]) for line in open('script2.py')])

lines = [line.rstrip() for line in open('script2.py') if line[0] == 'p']
print(lines)

lines = [line.rstrip() for line in open('script2.py') if line.rstrip()[-1].isdigit()]
print(lines)

#Let's count lines
output = list(os.popen('dir'))
print(len(output))

#for line in os.popen('dir'): print(line.rstrip())

lines = [line.rstrip() for line in os.popen('dir')]
print(len(lines))

#This checks for blank lines. The amount of the word 'line' in here is getting confusin!
print(len([line for line in lines if line.rstrip() == '']))

#map is a built in that applies a function to each item in tbe passed-in iteratable object.
#map returns an iterable object, , so we must wrap it in a list to receive all it's values at once.
print(list(map(str.upper, open('script2.py'))))

for item in map(str.upper, open('script2.py')):
    print('Item on your map for:',item.rstrip())

print('\nsingle line witchcraft:\n')

[print(line.rstrip()) for line in sorted(open('script2.py'))]
[print(line) for line in zip(open('script2.py'), open('script2.py'))] #line is a tuple here, there's no rstrip on a tuple
[print(line) for line in enumerate(open('script2.py'))]

#Curiously enough, python's iteration  protocol is present in a lot of places!
print('\n')
print(list(open('script2.py')))
print(tuple(open('script2.py')))
print('lol' . join(open('script2.py')))

a, b, c, d = open('script2.py')
print(a, b, c, d)

a, *b = open('script2.py')
print(a,b)

print('x = 2\n' in open('script2.py'))

L = [1,2,3]
L.extend(open('script2.py'))



