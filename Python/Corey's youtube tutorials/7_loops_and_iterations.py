
nums = [1,2,3,4,5]

for num in nums:
    print(num)

#break keyword will completely break out of a loop
#continue keyword will move on to the next iteration of a loop

for num in nums:
    if num == 2:
        print('Found 2!')
        break
    print(num)

print()

for num in nums:
    if num == 2:
        print('Found',num)
        continue
    print(num)

print()

mylist = [['banana', 'apple', 'orange'], ['red', 'green', 'blue'], 1, 400000]

for item in mylist:
    if type(item) is list:
        for subitem in item:
            print(subitem)
    else:
        print(item, 'is not iterable')

for num in nums:
    for letter in 'abc':
        print(num, letter)


#run through a loop 10 times

for i in range(1,11):
    print(i)

print()
x = 0
while x < 10:
    if x == 5:
        break
    print(x)
    x += 1