
x = 'dani'

while x:
    print(x, end=' ')
    x = x[1:]

print() #empty line to keep things less confusing

a = 0; b = 10
while a < b:
    print(a, end=' ')
    a += 1

print()

a = 0; b = 10
while a < b:
    print(a, end=' ')
    if a == 5:
        print("a equals five, breaking!")
        break
    a += 1

print()

x = 10
while x:
    x = x - 1
    if x % 2 == 0:
        print (x, end=' ')

print()
x = 10
while x:
    x = x - 1
    if x % 2 == 0:
        continue
    else:
        print(x, end=' ')

#Commenting this out as it's too annoying
#while True:
#    name = input('Enter name ')
#    if name == 'stop':
#        break

print()
y = 20
x = y // 2

while x > 1:
    if y % x == 0:
        print(y, 'has factor ', x)
        break
    x -= 1
else:
    print(y, 'is prime')

print()
x = 'daniel'

while x:
    print(x)
    if x == 'l':
        print('ni!!!')
    x = x[1:]
else:
    print('match condition never met')

print()
for x in ['one', 'two', 'three']:
    print(x, end=' ')

print()
sum = 0
for x in [1,2,3,4,5]:
    sum = sum+x
print('sum is: ', sum)

T = [(1,2), (3,4), (5,6)]
sum = 0
for (a,b) in T:
    sum = sum+a+b
print('sum of tuples is: ', sum)
print()

D = {'a':1, 'b':2, 'c':3}

for x in D:
    print(x, D[x])

print()
#the above, in other words:

for key in D:
    print(key, D[key])

for (a,b,c) in (1,2,3), (3,4,5), ('aaa','b','ca'), (('aaa','bbb'), 'ddd', 'fff'):
    print('Element one: ', a)

for (a,*b, c) in (1,2,3,4,5,6,7), (1,2,3,4,5,6,7,8):
    print(a,b,c)

items = ['aaa', 111, (4,5), 2.1]
tests = ['aaa', 2.1]

for key in tests:
    for item in items:
        if item == key:
            print(key, 'was found')
            break
        else:
            print(key, 'was not found')


seq1 = "spam"
seq2 = "scam"
res = []

for x in seq1:
    for y in seq2:
        if x == y:
            res.append(y)

print(res)
res = []

for x in seq1:
    if x in seq2:
        res.append(x)
print(x)

print([x for x in seq1 if x in seq2])

file = open('Chapter13_file.txt', 'r')
print (file.read())

file = open('Chapter13_file.txt', 'r')
while True:
    char = file.read(1)
    if not char: break
    print(char)

for char in open('Chapter13_file.txt', 'r'):
    print(char)

#Read into memory line by line:
file = open('Chapter13_file.txt', 'r')
while True:
    line = file.readline()
    if not line: break
    print(line.rstrip()) #line already has a newline at the end of it!

#Probably the easiest bit:

for line in open('Chapter13_file.txt', 'r'):
    print(line.rstrip())
    if not line: break

file = open('Chapter13_file.txt', 'rb')
while True:
    chunk = file.read(10)
    if not chunk: break
    print(chunk) #oh wow, weird output


for line in reversed(open('Chapter13_file.txt', 'r').readlines()):
    print(line)

print(list(range(10)))
print(list(range(5, 20, 2)))

for a in range(10):
    print (a)


L = [1,2,3,4,6,7]

for i in range(len(L)):
    X = L[i:] + L[:i]
    print(X, end=' ')

print()
S = 'abcdefghijklmno'
for i in range(0, len(S), 2):
    print(S[i], end=' ')

print()
print("L contains:", L)

print()
for a in range(len(L)):
    print(L[a], end=' ')

L1 = [1,2,3,4]
L2 = [5,6,7,8]
L3 = list(zip(L1, L2))

print()
print(L3)

for a in L3:
    print(a)

print()
S = ' daniel'

for (offset, item) in enumerate(S):
    print(offset, item)

print('Look, cat -n!', '\n')
for (i, l) in enumerate(open('Chapter13_file.txt')):
    print(i, l, end='')

import os

#F = os.popen('dir')
#print(F.readline())
#print(F.read(50))

#for line in os.popen('dir'): print(line.rstrip())

#for line in os.popen('systeminfo'):
#    parts = line.split(':')
#    if parts and parts[0].lower() == 'system type':
#        print(parts[1].strip())

#Hey, this is pretty useful for you crawler later on!
from urllib.request import urlopen
for line in urlopen('https://www.google.com'):
    print(line)
