
import sys
x = 1
y = 2

if y > x:
    print ("y is bigger than x")

if (y > x):
    print ("y is bigger than x")

if (x == 1 and
    y == 2):
    print ("yeah")

a, b = "hello", "world"

print (a)
print (b)

[c, d] = ["hello", "world"]

print (c)
print (d)

a, b, c, d = "four" #Assigns one char to each variable!
print (a, b, c, d)

a = b = c = d = "everybody gets the same thing!"
print (a, b, c, d)

string = "dani"
a, b, c = string[0], string[1], string[2:]
print (a, b, c)

for a,b,c in [(1,2,3),(4,5,6)]:
    print (a,b,c)

red, green, blue = range(3)
print (red, green, blue)

mylist = ("one", 2, 3)
a, b, c = mylist

a, *b, c = range(4)

for a, *b, c in (1,3,4,5,6,7,8), (1,2,3,4), ("asd", "ddddd", "huehueh", "ay"):
    print (a,b,c)

#for all in (1,3,4,5,6,7,8), (1,2,3,4), ("asd", "ddddd", "huehueh", "ay"):
#    for a in all:
#        print (a)

print ("Be careful with this:")
a = b = []
a.append(42)
print (a,b)

a, b = [], []
#a.append(42)
print ("aha!")
print (a,b)

x = print ("lalala")
print (x)

print ("does print add a white space", "between my sentences?", "I think it does!", sep=' HUE ', file=sys.stdout)

print ('a', 'b', 'c', end=''); print('','e', 'f')

sys.stdout.write("hello world! :D")

print ("hue!", file=sys.stdout)