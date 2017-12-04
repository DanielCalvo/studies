

#1a)

S = 'dani'
[print(ord(a)) for a in S]

#1b)


#Let's overcomplicate it!
sum1 = 0
for b in [ord(a) for a in S]:
    sum1 += b
print('sum is:',sum1)

#Let's keep it simple:
sum2 = 0
for a in S:
    sum2 += ord(a)
print(sum2)

#1c)
charlist = []
[charlist.append(ord(a)) for a in S]
print(charlist)

#2) That one looks kinda boring!

#3)

#Here's our dictionary:
D = {'a': 10, 'ddd': 90, 'b': 20, 'c': 3}
D['e'] = 123

#The easy way:
for k in sorted(D):
    print(k, D[k], end=' ')
print('\n')

#Now, let's make it needlessly complicated:

#Wait, this part is actually wrong!
keylist = []
valuelist = []

for key in D:
    keylist.append(key)
    valuelist.append(D[key])

#LOOK AWAY NOW CHILD
keylist.sort()

sorted_D = dict(zip(keylist, valuelist))

for key in sorted_D:
    print(key, sorted_D[key])