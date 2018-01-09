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

def tracer (func, *pargs, **kargs):
    print('arguments received:', func, *pargs, **kargs)
    print('calling:', func.__name__)
    return func(*pargs, **kargs)

def func(a,b,c,d):
    return a+b+c+d

print(tracer(func, 1,2,3,4))
print(func(1,2,2,3))

def kwonly1(a, *b, c):
    print(a,b,c)

kwonly1(1,2,33,c=3) #a may be passed by name or position, b collects any extra positional arguments, c must be passed by keyword

def kwonly2(a, *, b, c): #this forces b and c to be passed by keyword
    print(a,b,c)

kwonly2(1, b=44, c=55)

def kwonly3(a, *, b, c='spammm'):
    print(a,b,c)

kwonly3('a', b='stuff')

#Keyword arguments must be defined after a single star, not two star!

def fu7(a=0, *b, c=7, **d):
    print(a, b, c, d)
    for key in d:
        print(key, d[key])

fu7(1,2,3,4)
fu7(ccc=8, ddd=9, eee=10)

def fu8(a, *b, c=6, **d): print(a,b,c,d)

fu8(1, *(2,3), **dict(x=8,z=99))

def min1(*args):
    res = args[0]
    for arg in args[1:]:
        if arg < res:
            res = arg
    return res

def min2(first, *rest):
    for arg in rest:
        if arg < first:
            first = arg
    return first

def min3(*args):
    tmp = list(args)
    tmp.sort()
    return tmp[0]

def max1(*args):
    tmp = list(args)
    tmp.sort()
    return tmp[-1]

print(min1(45,6,6,6,6,777,4))
print(min2(45,6,6,6,6,777,4))
print(max1(55,66,1))

def minmax(function, *args):
    res = args[0]
    for arg in args[1:]:
        if function(arg, res):
            res = arg
    return res

def lessthen(x, y): return x < y
def greaterthen(x, y): return x > y

print(minmax(lessthen,6,4,3,32,22,66,2,34,4,8,888)) #I am so confused. EDIT: I am only slightly confused.

def intersect(*args):
    res = []
    for x in args[0]:
        if x in res: continue
        for other in args[1:]:
            if not x in other: break
            else:
                res.append(x)
    return res

def union(*args):
    res = []
    for seq in args:
        for x in seq:
            if not x in res:
                res.append(x)
    return res

s1 = 'SPAM'
s2 = 'SCAM'
s3 = 'BLAM'
dani1 = 'dani'
dani2 = 'dani'

list_one = [1,5,6]
list_two = [1,8,7,6,6,6,6,7,66]

print(intersect(list_one,list_two))
print(intersect(s1,s2,s3), union(s1,s2))
print(intersect(dani1,dani2))
#A google search reveals that intersections can be performed with a lot less fuss:

list_three = [val for val in list_one if val in list_two]
print(list_three)


def tester(func, items, trace=True):
    for i in range (len(items)):
        items = items[1:] + items[:1]
        if trace: print(items)
        print(sorted(func(*items)))

print('checking out the tester fucntion')
print(tester(union,(list_one,list_two),trace=False))


def myfunc1(a, b, c=3, d=4): print(a, b, c, d)

myfunc1(1, *(5, 6))

