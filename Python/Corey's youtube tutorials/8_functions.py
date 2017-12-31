
def hello_func():
    print('Hello world!')

#print(hello_func()) #None, no return value

hello_func()

#Don't repeat yourself!
#For beginners: Don't get caught up on understanding every detail of what every function does. Just focus on the input and what's returned.

def hello_func2():
    return ('Hello world 2!')

print(hello_func2().upper())


def hello_func3(greeting, name='You'):
    return '{}, {}'.format(greeting, name)

print(hello_func3('Hi'))
print(hello_func3('Hi','Dani'))

def studend_info(*args, **kwargs): #Allowing us to accept an arbitrary number of positional or keyword arguments
    print(args) #A tuple!
    print(kwargs) #A dictionary with all our keyword values

courses = ['Math', 'Art']
info = {'name': 'Dani', 'age': '333'}

studend_info(*courses, **info)


# Number of days per month. First value placeholder for indexing purposes.
month_days = [0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

def is_leap(year):
    """Return True for leap years, False for non-leap years."""

    return year % 4 == 0 and (year % 100 != 0 or year % 400 == 0)


def days_in_month(year, month):
    """Return number of days in that month in that year."""

    if not 1 <= month <= 12:
        return 'Invalid Month'

    if month == 2 and is_leap(year):
        return 29

    return month_days[month]

print(is_leap(2017))
print(is_leap(2020))

print(days_in_month(2017, 2))
print(days_in_month(2017, 12))