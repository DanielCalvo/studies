str = "a string"
p str

# single quotes do not allow string interpolation
str2 = 'other'
p str2

#concat
p str + ' ' + str2

#interpolation
puts "Hey look your string: #{str}"

#irb is a thing!

#huh neat -- yo can call methods on primitive types -- everything is an object in ruby
p str2.class
p "a".class
p 1.class
p 1.1.class

#prints all the string methods, noice!
p str.methods

puts "banana".empty?
puts str2.length
puts "banana".reverse
p "".empty?
p "".nil? #even though its empty, its not nil
p nil.nil? #true

sentence = "Welcome to the jungle"
p sentence.sub("the jungle", "my house") #ohh this returns the new string, but does not change it in place

#"Welcome to the jungle"
p sentence

sentence = sentence.sub("the jungle", "my house")
p sentence

str3 = 'Escaped \'quotes\'' #if you don't escape them, it breaks
p str3

#p "Getting input:"
#str4 = gets.chomp
#p "input: #{str4}"