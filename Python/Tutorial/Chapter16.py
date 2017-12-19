
def times(x,y):
    return x * y

print(times(2,3))

x = times(9,9)
print(x)

print(times('dani', 3)) #oh boy
print(times('dani', len('dani'))) #so cool

def intersect(seq1, seq2):
    res = []
    for x in seq1:
        if x in seq2:
            res.append(x)
    return res


list1 = [1,2,4,5,6,6,6,6,6,7,7]
list2 = [2,7,5,7]

print(intersect(list1,list2))
print(intersect(list2,list1))
