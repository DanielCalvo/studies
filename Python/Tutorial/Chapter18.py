#Arguments!

#The theory on this is a bit heavy, let's get to some coding, maybe that will clear things up!


def f(a):
    a = 99 #Changes lo cal variable only!

b = 88
f(b)
print(b)

def changer(a, b):
    a = 2
    b[0] = 'spam' #Changes shared object!

X = 1
L = [1, 2]

changer(X, L)
print(X, L) #Interesting, the list changes.

X = 1
a = X
a = 2
print(X)

L = [1,2]
b = L
b[0] = 'spam!'
print(L)

#Changing a mutable object in place can impact other references of that object

mylist = [1,2,3,4]

changer(4,list.copy(mylist))
print(mylist) #aha!

changer(3, mylist[:])
print(mylist)

#changer(3, tuple(mylist)) #fails

#Functions might update mutable objects like lists and dictionaries passed into them. This is a feature!

def multiple(x,y):
    x = 2
    y = [3,4]
    return x,y

X = 1
L = [1,2]
X, L = multiple(X, L)
print(X, L)

#Simple function with 3 arguments

def fu(a,b,c): print(a,b,c)
fu(6,6,6)

#Oh, we can match by name!

fu(c=8,b=7,a=6)

#You can mix and mash the two:
fu(1,c=3,b=3)

#Creating a function with defaults arguments if not enough ones are passed:

def fu1(a, b=2, c=3): print(a,b,c)

fu1(7)
fu1(7,1)
fu1(4,c=7)

def fu2(spam, eggs, toast=0, ham=0): print(spam, eggs, toast, ham)
fu2(1,2)

#Take as many arguments as you want!

def fu3(*args): print(args) #Tuple!
fu3(235,2345,435,6577)

def fu4(**args): print(args) #Only works for keyword arguments, makes a dictionary
a = fu4(a=4,aaa=666)

def fu5(a, *pargs, **kargs): print (a, pargs, kargs)

fu5(1,2,3, x=1, y=2)

def fu6(a,b,c,d): print(a,b,c,d)
myargs = (1,2)
myargs += (3,4) #It's a tuple!

fu6(*myargs)

myargs = {'a': 1, 'b': 2, 'c': 3}
myargs['d'] = 4

fu6(**myargs)

fu6(*(1,2), **{'d': 44, 'c': 33})


#That tracer function looks cool!