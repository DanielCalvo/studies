import b
b.spam('banana')

import sys
print(sys.path)

from b import spam
spam('more bananas')

from module1 import *
printer("Hello world")

from small import x,y
x = 42
y[0] = 42
print(x, y)

#This chapter is a bit heavy on the theory.
#I am unlikely to require knowledge beyond the basics in this area for the simple programs I write