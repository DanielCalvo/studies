
#import my_module as mm
#from my_module import find_index as fi, test #Only gives access to the find index, not everything else in the module
from my_module import find_index, test
#from my_module import * #You can't tell what came from the module and what didn't!
#Careful with renaming, don't make your code unreadable!
import sys
#sys.path.append() #Can be used to append a directory for Python to look for modules

courses = ['History', 'Math', 'Physics', 'CompSci']

#print(find_index(courses, 'Math'))
#print(test) #Neat!

print(sys.path) #Where python looks for modules

import random

random_course = random.choice(courses)

print(random_course)

import math

rads = math.radians(90)
print(rads)
print(math.sin(rads))

import datetime, calendar

today = datetime.date.today()
print(today)

print(calendar.isleap(2017))

import os

print(os.getcwd())

print(os.__file__)