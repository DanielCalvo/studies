
num = 1
str = 'aaa'
floatnum = 3.14
nilsy = nil

puts num.instance_of? Integer
puts str.instance_of? String

puts num.class

#Ruby isn't concerned with types as much as it is with classes:
if num.class == Integer then puts ("#{num} is an integer") end
puts Integer.class

# Numeric is a parent class of both Integer and Float. Neat!
num.is_a?(Numeric)

