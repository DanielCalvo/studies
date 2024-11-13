
puts true
puts 5 < 10
puts 10 > 12

puts a = true
puts a.class #TrueClass, interesting
puts false.class #FalseClass, huh

puts 10 == 10
puts 2 == 3
puts "2" == 2
puts "2".to_i == 2

puts 5 == 5.0 #true, interesting

puts 10 != 5

#predicated method, ah so that's what they're called!
puts 10.odd?
puts 10.even?
puts 10.positive?
puts 10.negative?
puts 10.zero?

#methods with arguments!
puts "Big Mac".include?("Big")
puts "Big Mac".include?"aa" #oh you can do it this way without parenthesis

#puts "Big Mac".include? # will raise an ArgumentError -- needs one arg
#puts "Big Mac".include?("aa", "cc") also raises argument error

#convention: if method takes no args, leave out parenthesis. If it does take args, use parenthesis
# for top level methods, like puts, parenthesis are usually avoided.

#methods with multiple arguments!
puts 20.between?(10, 30)
puts "b".between?("a","c") #true, ha! might be good to not rely on this too much though

#arithmetic methods
puts 1+3
#equivalenet to
puts 1.+(3) #woah, a method named "+"
puts 5.-3
puts 10./3
puts 10.div 3
puts 10.div (3) #so cool! these are not used in the real world much, but its nice to know about them

#float methods
puts 10.9.floor.class #returns an int, neat!
puts 10.9.ceil #returns an int, neat!
puts 10.01.ceil
puts 3.14159.round #rounds to closes int with if given no args
puts 3.14159.round(2)
puts 3.5.round #rounds to 4 if .5, interesting

puts 5.35.abs #prints the distance from 0
puts -3.1.abs #absolute value -- method also present on integers

#author says: if a method accepts arguments, even if they are optional, include the parenthesis