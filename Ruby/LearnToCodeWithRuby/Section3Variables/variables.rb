#variables are case sensitive!
#ruby preffers snake case

age = 31
age = age + 3
puts age

#trying to print a variable that does not exist will give you a NameError
#puts somevar

#parallel variable assignment, neat
a, b, c = 1, 2, 3

#swapping variable values!
a, b = b, a
print a,b,c,"\n"

#shortcuts
a += 1 #aka a++
puts a

#also works on strings
name = "Bob"
name += " McBobson"
puts name

#constants cannot be changed
#Constants only require a capital first letter, but the community convention is to have it in all caps
PI = 3.14
#you can change a constant, but ruby will give you a warning
PI = 6