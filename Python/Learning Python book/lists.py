
L = [123, 'spam', 1.23]

print (len(L))

L.append('wooo')

print (L)

#You can't sort L. A list with strings and integers is unorderable
#L.sort()

L = ['a', 'b', 'd', 'ffff', 'asdddd']
L.sort()
print (L)

#Note: If you declate L with parenthesis like this:
L = ('a', 'b', 'c')
#Python assumets it's a tuple. Lists are declared like this:
L = ['a', 'b', 'd', 'ffff', 'asdddd']

#print (L[99]) <- out of range!

M = [ [1,2,3], [4,5,6], [7,8,9] ]
print (M[1])

col2 = [row[1] for row in M]
print ("This is row 2:", col2)

print ([row[1] +1 for row in M])

#filter out odd items, this is some really crazy shit
print ([row[1] for row in M if row[1] % 2 == 0])

print (list(range(4)), list(range(-6 ,7, 2)))

G = (sum(row) for row in M)

#I don't understand what I'm doing here!
print (next (G))
print (next (G))
print (next (G))

print (list(map(sum, M)))

number_list = [123, 444, 111, 2323, 333, -999]

#aha, got it!
print (sum(number_list))

string_list = ['asd', 'basd', 'dddd', 'cccc', 'zzz', 'xxx', 'eee']

mystring = 'aaaaaw yeah'
print ([ord(x) for x in mystring ])
