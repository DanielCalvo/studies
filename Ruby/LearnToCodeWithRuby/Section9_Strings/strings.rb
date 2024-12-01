#Double vs single quotes
puts "hello\n"
puts 'hello\n' #inteprets /n literally

#String interpolation only works with double quotes!
var = "joe"
puts "I have a friend named #{var}"
puts 'I have a friend named #{var}' #no interpolation

#multiline strings!

# TEXT is an identifier -- this be anything
mytext = <<TEXT #Ruby interprets this string literally
Roses are red
Violets are blue
All of my base, 
Are belong to you
TEXT

puts mytext

#different ways to compare strings in ruby
a = "Hello"
b = "hello"
c = "Hello"

puts a == b
puts a == c
puts b == 'hello'
puts b != 'hello'

#You can use the < and > signs to compare two characters based on their order in the alphabet
puts "a" > "b" #false -- looks like later in the alphabet means higher value
puts "6" > "5" #true 
puts "Banana" > "Apple" #true apparently

puts "Dog" > "Cat" #Checks for an alphabetical sort, and since dog is further later in the dictionary than the cat word, this evaluates to true.

#concatenating strings
puts "a"+"b"

somevar = "ab"
somevar += "c"
puts somevar #abc -- neat, its like the numeric increase thing but for strings it concatenates
#But there's also a contact method on strings!
somevar.concat("d") #This however mutates the string that calls the method!
puts somevar
puts somevar.prepend("0") #both concat and prepend modify the original string AND return the modified string, interesting
puts somevar

#Shovel operator <<
somevar << "e" #appends while mutating the original string
puts somevar

#YOu can call the shovel operator multiple times in sequence!
somevar << "f" << "g"
puts somevar

#lenght and size. Size is an alias for length -- they accomplish the exact same thing
puts "banana".length
puts "banana".size

#index positions
str = "Hey here is my string"
puts str[0]
puts str[0..2]
puts str[-1]
puts str[-2]
puts str[-3..-1]
puts str[4..-2] #interesting way to range
p str[1999] #nil

#theres also the slice method that does the same deal
puts str.slice(0..2)
puts str.slice(-1)

puts str[4,4] #start at index 4 and extract 4 characters, neat!
puts str[0, str.length] #the whole thing!
puts str[-6, str.length] #neat, works with negative starts too

#in ruby, strings are mutable!
thing = "banana"
thing[-1] = "e"
puts thing #woah!

food = "Meat"
food[0, 4] = "Fish"
puts food #WOAH -- you can target a segment!

car = "Blue car"
car[0,4] = "Red" #You can replace an element slice with less elements as seen here
puts car

#The insert method -- does not overwrite, only inserts!
thing = "banana"
thing.insert(-1, " ice cream") #neat, essentially appends if done like this
puts thing

#empty and nil methods
puts "".empty?
puts "".nil?

abc = nil
puts abc.nil? #the nil method is available on every single object

#case methods -- none of these mutate the string, they all return new strings!
s = "Hello there!"
puts s.upcase
puts s.downcase
puts s.swapcase
puts s.capitalize #capitalizes first letter and everything else is lowercased
puts s.reverse

#Bang methods -- they always end with an exclamation point
#They perform some kind of mutation -- changes the original objects
word = "banana"
word.capitalize! #oOoooOoo
puts word
word.upcase!
puts word
word.downcase!
puts word

#And so on...