
import sys
print(dir(sys)) #I wonder what's in here!

#for thing in dir(sys): #Ok it works but that's a lot of output
#    print(thing)

print(len([x for x in dir(sys) if not x.startswith('__')]))


print(len(dir([])), len([x for x in dir([]) if not x.startswith('_')]))

def dir1(x): return [a for a in dir(x) if not a.startswith('_')]

print([x for x in dir([]) if not x.startswith('_')])
print([x for x in dir('') if not x.startswith('_')])

print(dir1([]))# oooOoooh a function# !

print(dir(str) == dir(''))
print(dir(list) == dir([])) #Same deal

import docstrings
# print(docstrings.square(4))
#print(docstrings.square.__doc__)

#print(sys.__doc__)

print(str.upper.__doc__)
print(int.__doc__)

help('')
print('Help on an empty string wouldn\'t work according to the book, but it did!')


help(docstrings.square)
help(docstrings)