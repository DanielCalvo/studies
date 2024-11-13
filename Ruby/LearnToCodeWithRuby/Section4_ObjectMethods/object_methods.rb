
#the string in here is called the receiver
puts "hello world".length
puts "hello world".upcase
puts "TOTALLY DIFFERENT".downcase

a = "TOTALLY DIFFERENT"
puts a.downcase

#you can also:
puts a.downcase() #but parenthesis are not used in ruby if they aren't necessary

puts 10.next #WOAH
puts 10.succ #successor. A SUCC METHOD AY LMAO
puts 10.pred

#puts a.downcas # raises NoMethodError

#method chaining
puts "hi there".upcase.length.succ
puts 10.succ.succ.succ

#inspect method -- converts an object into a string representation for debugging
#oooh the p method actally calls the inspect method on whatever you pass to it!
puts "hello world\n"
puts "hello world\n".inspect
p "hello world\n"

#huhhh not sure the author is right, he claims p calls inspect, but p 11 and p 11.inspect pring different clæsses apparently
p 11
p 11.inspect

#in ruby, nil is a object that represents emptiness, or the absence of a value
puts nil
p nil

#string interpolation -- injecting dynamic content inside a string
name = "Bob"
puts "Hello #{name}"

age = 2
puts "I am #{age} years old"
puts "In 5 years I'll be #{age + 5} years old" #ooooo you can run logic in there
x = 6
y = 5
puts "The sum of #{x} and #{y} is #{x+y}" #neat!

#oh you can't do this, its like trying to concat an integer into a string, interesting:
#puts "i am" + age+ " years old"

#gets
# puts "hey whats up?"
# input = gets.chomp #enter adds a newline, so chomp removes it
# puts "Whats up: #{input} -- awesome!"

#class method!
puts 3.class,  "aa".class

#methods to convert objects
a = "5".to_i
puts a.class, "10".to_i.class

b = "15 apples"
puts b.to_i #woah, 15!
puts b.to_f

c = "oranges 15"
puts c.to_i #gives you zero if it can't convert, but doesnt throw an error

d = 11
put d.to_s

#you also can:
puts "asd".to_s #which doesnt make any sense in this case, but can be handy when you have some dynamic value in some method that you don't know what class it is and really want it to be a string. So if its already a string, no harm in running this.
#this is related to polymorphism, author says ruby may not be concerned that much with the class or something, but more with the methods it can execute