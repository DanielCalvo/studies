
def mysum(L):
    if not L:
        return 0
    else:
        return L[0] + mysum(L[1:])

mylist = [4,5,6]

print(mysum(mylist))

emptylist = [1]

