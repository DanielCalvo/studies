
#Variables inside functions are local to said functions:

x = 5

def useless_function():
    x = 9
    return x

print(useless_function(), x)

#The LEGB Rule:
#Local, Enclosing, Global, Built-in
#Assignments always determine the scope. Always!

#import builtins
#[print(x) for x in dir(builtins)] # That's a lot of output


def hi():
    global AAA
    AAA = 3
hi()
print(AAA)

#Pro tip: Do not change variables in other files, such as

#import first
#first.X = 99

#The first module probably has no clue that some other, arbitrary module is chaging it's variables at run time
#It is also the same as having global variables -- it can be very, very difficult to track down a variable's value

#factory function:

def maker(N):
    def action(X):
        return X ** N
    return action

f = maker(2)
print(f(4))

g = maker(3)
print(g(4))

def maker_lambda(N):
    return lambda X: X ** N

h = maker(3)
print(h(4))

#defs within defs can be confusing however. Flat is preffered, and more "pythonic" as claimed by the book!


def fu1():
    x = 88
    fu2(x)

def fu2(x):
    print(x)

fu1() #interedasting
fu2(2)


def func():
    x = 4
    action = (lambda a: x ** a) #where does a come from?
    return action

x = func()
print(x(3)) #what is going on here?! I

#Avoided most of the nested function stuff

#Non local only has meaning inside functions. This fails:
#nonlocal Z

