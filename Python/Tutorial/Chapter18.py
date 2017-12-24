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






