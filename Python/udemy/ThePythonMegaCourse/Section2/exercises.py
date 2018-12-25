# List indexing practice
# name = "John Smith"
# print(name[2:4])

# letters = 'abcdefghijklmnopqrstuvwxyz'
# print(letters[-2])

# letters = 'abcdefghijklmnopqrstuvwxyz'
# print(letters[-3:-1])

def celsius_to_fahrenheit(celsius_temp):
    return (celsius_temp * 9 / 5 + 32)

def string_length(mystring):
    if type(mystring) == str:
        return len(mystring)
    elif type(mystring) == int:
        print("Integers don't have lenght")
    elif type(mystring) == float:
        print("floats don't have lenght")

#print(string_length("zasd"))
#print(string_length(22.1))

#myfile = open("sample.txt")
#content = myfile.read()
#print(content)
#myfile.close()
#
#The .splitlines function removes newline and other (tabulation?) characters from file file objects, making them a list.
#Very cool!

#2.40
#myfile = open("fruits.txt")
#content = myfile.read()
#print(content)
#myfile.close()

# myfile = open("sample.txt")
# c = myfile.read()
# c = c.splitlines()
#
# for i in c:
#     print(i)
#

#2.43
#mylist = [1, 2, 3, 4, 5]

#for element in mylist:
#    print(element)

#2.45
# mylist = [1, 2, 3, 4, 5]
#
# for item in mylist:
#     if item > 2:
#         print(item)

#2.47
#Pro tip: If you don't slitlines, len() will count the newline at the end of each newline as part of the lenght of the line :o
# myfile = open("fruits.txt")
# c = myfile.read()
# c = c.splitlines()
#
# for i in c:
#     print(len(i))

#2.49
# temperatures = [10, -20, 100]
# for temperature in temperatures:
#     print(celsius_to_fahrenheit(temperature))

#a+ mode allows you read and write to a file, neat!

# myfile = open("fruits.txt", "a+")
# print(myfile.read())
# myfile.seek(0)
# print(myfile.read()) #Neat!
# myfile.write("grape")
# myfile.close()

#2.54
numbers = [1,2,3]
myfile = open("numbers.txt", "w")

for number in numbers:
    myfile.write(str(number) + "\n")
myfile.close()








