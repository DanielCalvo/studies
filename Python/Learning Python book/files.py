
#f = open('data.txt', 'w')
#f.write("hey!")
#f.write("wassup?")
#f.close()

for line in open('data.txt'):
    print (line, end='')

print (open('data.txt').read())

