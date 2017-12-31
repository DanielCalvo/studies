
if True:
    print('Condition was :D')

language = 'Java'

if language == 'Python':
    print('Language is Python!')
elif language == 'Java':
    print('Language is Java!')
elif language == 'Javascript':
    print('Language is Javascript!')
else:
    print('No match')

#Python doesn't have a switch case! Just keep adding elif statements.

# Object identity: is
# Evalues if objects are the same in memory

#Booleans:
#and
#or
#not

user = 'admin'
logged_in = False

if user == 'admin' or logged_in == True:
    print('admin_page')
else:
    print('Bad credentials')

if not logged_in:
    print('Please log in')
else:
    print('Welcome!')

a = [1,2,3]
b = [1,2,3]

print(a == b)
print(id(a), id(b))

print(a is b) #Returns false! These are two different objects. Checks if IDs are the same as above

b = a

print(id(a), id(b))
print(a is b) #Same object in memory
print(a == b)
print(id(a) == id(b)) #What the "is" comparison does behind the scenes


# False Values:
    # False
    # None
    # Zero of any numeric type
    # Any empty sequence. For example, '', (), [].
    # Any empty mapping. For example, {}.

#Everything else evaluates to true!

condition = 0

if condition:
    print('Evaluated to True')
else:
    print('Evaluated to False')

