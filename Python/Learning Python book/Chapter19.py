
def mysum(L):
    if not L:
        return 0
    else:
        return L[0] + mysum(L[1:])

mylist = [4,5,6]

print(mysum(mylist))

emptylist = [1]


def mysum1(L): return 0 if not L else L[0] + mysum1(L[1:])
print(mysum1(mylist))

def mysum2(L):
    first, *rest = L
    return first if not rest else first + mysum2(rest)

print(mysum2(mylist))

#Recursion is probably overkill in this context (small functions)

mylist1 = [1,2,3,4,5,6]
sum = 0

while mylist1:
    sum += mylist1[0]
    mylist1 = mylist1[1:]

print(sum)

def sumtree(L):
    tot = 0
    for x in L:
        if not isinstance(x, list):
            tot += x
        else:
            tot += sumtree(x)
    return tot

nestedlist = [1, [2, [3, 4], 5], 6, [7, 8]]
print(sumtree(nestedlist))

#The same example, without recursion:

def sumtree1(L):
    tot = 0
    items = list(L)
    while items:
        front = items.pop()
        if not isinstance(front, list):
            tot += front
        else:
            items.extend(front)
    return tot

print(sumtree1(nestedlist))

def echo(message):
    print(message)

echo('Direct call')

x = echo
x('Indirect call')

def indirect(func, arg):
    func(arg)

indirect(echo, 'Argument call!')

schedule = [ (echo, 'Spam!'), (echo, 'Ham!') ]
for (func, arg) in schedule:
    func(arg)

def make(label):
    print('Function created with label:',label)
    def echo(message):
        print(label + ':' + message)
    return echo

F = make('Spam')
F('Ham')

print('F\'s varnames: ', F.__code__.co_varnames)


def funcspam(a):
    b = 'spam'
    return b * a

print(funcspam(8))

#You can attribute user generasted attributes to a function:
#Though I'm not sure if that' recommended!

funcspam.count = 2
print(funcspam.count)

def func_annotate(a: 'spam', b: (1,10), c: float) -> int:
    return a+b+c

print(func_annotate(1,2,3))

print(func_annotate.__annotations__)

for arg in func_annotate.__annotations__:
    print(arg, '=>', func_annotate.__annotations__[arg])

#You can still use defaults, but maybe it gets a bit messy to read:

def func2(a: 'spam' = 5, b: (1,10) = 5, c: float = 6) -> int:
    return a+b+c

print(func2(5,6,7))

#Lambda things!

f = lambda x,y,z: x+y+z
print(f(6,7,8))

x = (lambda a='fee', b='fie', c='foe': a+b+c)

print(x('feeeeee'))

def knights():
    title = 'Sir'
    action = (lambda x: title +' '+ x)
    return action

act = knights()
print(act('Robin'))

L = [
    lambda x: x ** 2,
    lambda x: x ** 3,
    lambda x: x ** 4,
]

for f in L:
    print(f(2))

print(L[0](3))

key = 'got'

{'already': (lambda: 2+2),
'got': (lambda: print(2*4)),
'one': (lambda: 2**6)}[key]()

