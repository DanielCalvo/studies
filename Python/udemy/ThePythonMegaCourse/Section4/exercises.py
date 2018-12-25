
#4.77

def divide(a,b):
    try:
        return a/b
    except ZeroDivisionError:
        return("Can't divide by zero")


print(divide(1,2))
print(divide(1,0))