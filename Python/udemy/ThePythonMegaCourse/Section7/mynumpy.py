import numpy

n = numpy.arange(27)
#print(n)

n = n.reshape(3,9)
#print(n)

n = n.reshape(3,3,3)
#print(n)

mylist = [[123,12,123,12,33],[],[]]

my_numpylist = numpy.asarray([[123,12,123,12,33],[],[]])

print(my_numpylist)
print(type(my_numpylist))