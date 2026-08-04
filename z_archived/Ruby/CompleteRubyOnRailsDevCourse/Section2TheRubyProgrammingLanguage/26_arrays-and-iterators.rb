a = [0,1,2,3,4,5]
puts a #iterates and prints a newline after every element
puts "Here is a: #{a}" #does not iterate, interesting
puts a.first, a.last

x = 1..10
puts x
puts x.class #Range, interesting!
x = x.to_a.shuffle
puts x

x = 1..10
x = x.to_a
#The bang (!) actually mutates the caller
x.reverse!
puts "Here is x: #{x}"

x = "a".."z"
x = x.to_a #can't do x.to_a! directly, interesting
x.shuffle!
puts "Here is x: #{x}"

a = [0,1,2]
a << 3 #Appends 3, awesome!
puts a 
a.append(4)
puts a 
a.unshift(0) #Prepends
a.unshift(0)
puts 
puts a 
a.uniq! #Removes duplicates
puts 
puts a 

if a.include?(0) #This is really cool!
  puts "a contains 0, neat!" 
end

b = []

if b.empty?
  puts "b is empty!"
end

b = "a".."e"
b = b.to_a
b.pop
b.pop #Changes b in place, interesting
b.push("0")
b.push("0")
puts b

s = %w("banana split is my favourite ice cream")
puts s

#prefered way of iterating through an array in ruby
s.each do |word|
  print word + " - "
end
puts
#but author says they really like doing it in a single line in ruby
s.each {|word| print word, " "}

z = (1..20).to_a.shuffle

#prints odd numbers, woah, this is really different from what I'm used to 
puts z.select {|number| number.odd?}