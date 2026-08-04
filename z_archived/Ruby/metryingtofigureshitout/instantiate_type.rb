#Well the correct terminology is "class" but let me do my thing

a = Float
a = 3.14

puts a
puts a.class

a = 1 #If you do this it is no longer a float -- it becomes an integer
puts a.class

b = 1.0 #If I give it a decimal value, is it a float?
puts "b is #{b.class}" #Yes! Neat