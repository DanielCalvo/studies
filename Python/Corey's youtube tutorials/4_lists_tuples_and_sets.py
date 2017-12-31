

courses = ['History', 'Math', 'Physics', 'CompSci']

print(courses)
print(len(courses))

print(courses[0])
print(courses[3])
print(courses[-1])
#print(courses[4]) #List of index out of range

print(courses[0:2]) #First index is inclusive, second one is not
print(courses[:2])
print(courses[2:]) #This is named slicing
print(courses[0])

courses.append('Art')
print(courses)

courses.insert(0, 'Art')
print(courses)

courses2 = ['Geography', 'Education']

#courses.append(courses2) #Puts the list inside the other list

courses.extend(courses2) #Puts the elements of a list inside the other list (does not nest lists!)

courses.remove('Art') #Removes the first instance of 'Art', not all of them apparently
print(courses)

courses.pop() #Useful if you want to use your list like a stack or a queue, removes last item
print(courses)

#pop() function returns the value it removed,

popped = courses.pop()
courses.append(popped) #Ha! I just did nothing!
print(courses)

courses.reverse()
print(courses)

courses.sort()
print(courses)

numbers = [5,6,888,1,33,5]
numbers.sort(reverse=True) #Descending order, neat!
print(numbers)

courses.sort(reverse=True)
print(courses)

print(sorted(courses)) #Does not sort the list in place. Returns a sorted version of the list!
print(courses)

print(min(numbers), max(numbers))

print(sum(numbers))
#print(sum(courses)) #Gives an error if you run it on a list with strings

print(courses.index('Math'))

print('Art' in courses) #True
print('Medicine' in courses) #False

for course in courses:
    #print(course)
    pass

for index, course in enumerate(courses): #enumerate returns two values: The index that we're on and the element
    print(index, course, type(index), type(course))
    break

for course in enumerate(courses):
    print(course, type(index))
    break

for course in enumerate(courses, start=1):
    print(course, type(index))
    break


#Turn our list of courses into a string of comma separated values (a csv!)

course_str = ','.join(courses)
print(course_str) #Neat!

new_list = course_str.split(',')
print(new_list) #Oh!

#Tuples and sets!

#Tuples are similar to lists, but you can't modify them. They're immutable. Sorta like a read only list.

#Sets:
#Unlike lists or tuples, sets don't care about orders!
#Sets throw away duplicates, so they're unique values only apparently

cs_courses = {'History', 'CompSci', 'Art', 'Math', 'Math'}
art_courses = {'History', 'Design', 'Painting', 'Math', 'Math'}
print(cs_courses)

print('Math' in cs_courses)
print(type(cs_courses))
print(cs_courses.intersection(art_courses)) #Prints elements in common
print(cs_courses.difference(art_courses)) #Courses that are in cs_courses but not on art_courses

print(cs_courses.union(art_courses)) #Prints all unique elements on both sets

empty_set = {} #Wrong! Creates a dictionary
empty_set = set() #Correct!