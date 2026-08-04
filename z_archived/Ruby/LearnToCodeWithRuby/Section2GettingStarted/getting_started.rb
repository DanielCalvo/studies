#invoking the puts method
puts "hi!"
#you can add parenthesis to a method invocation
puts ("hi")
#but as a general people don't use them if they are not needed in ruby
puts 5, 3, 5 #adds a newline, neat

puts "text\nafter a newline!"
puts "Someone said: \"Goodbye\""

print "hi\n"
print "print does not add a line break\n"

p "the p method! -- can give you more details, like printing the doublequotes"
p "text\nafter a newline!" #prints the escape character, neat!

=begin
multi
line
comment! neat
=end

puts 1_000 #optional underscore to separate numbers for readability
p 1_000
puts 3.14 #dot makes ruby interpret this number as a float
puts 0.14

#PEMPAS
# Parentheses, exponents, multiplication, division, addition, subtraction
puts (2+5) * 10

#By default, ruby does floor division, always gives you and int
puts 12 / 5
puts 12.0 / 5 #gives float result
puts 12 / 5.0 #same deal
puts 5 % 2

#concatenation
puts "race" + "car"

#puts "a" + 2 #typeerror