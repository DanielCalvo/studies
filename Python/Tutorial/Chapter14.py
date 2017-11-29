
print(open('script2.py').read())

print()

for line in open('script2.py'):
    print(line.upper(), end='')

#Using file.readlines() is considered less than ideal to read files, as it loads the entire file into memory!
#This may cause the program not to run if you load files that are too big into memory.

#You can also iterate through a file with a while loop!

f = open('script2.py')

while True:
    line = f.readline()
    if not line: break
    print(line.upper(), end='')

#The book says: Iterators (for loop) run at C language speed. The while loop runs Python byte code through the Python virtual machine, so it tends to be slower.

