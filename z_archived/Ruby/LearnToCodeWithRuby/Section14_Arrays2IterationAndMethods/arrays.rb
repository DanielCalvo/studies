
3.times {|number| puts "Iterating: #{number}"}


4.times do |number| #block variable
  puts number
end  

names = ["joe", "john", "jay"]
#Ruby needs to be told what to do for each array element
names.each {|name| puts name.upcase} #the each method accepts a block, neat!

[1,2,4,5,6].each do |current_number|
  square = current_number * current_number
  puts square
end

#filtering with the each method
fives = [5, 10, 15, 20, 25, 30, 35, 40]
evens = []
fives.each {|value| evens.push(value) if value.even? }
p evens

#nested each loops
a1 = [1,2,3]
a2 = [4,5,6]

a1.each do |el_a1|
  a2.each do |el_a2|
    puts "#{el_a1} #{el_a2}"
  end
end

#for loop! less preferred in ruby apparently

for number in [1,2,3]
  puts number
end
#ooo, one disadvantage is that number will persist after the execution
puts number #not cool

#.each seems a bit cleaner, I like it

a1 = [1,2,3]

def double_elements(a)
  a.each do |el|
      el = el * 2 #nope, doesnt work in multiplying the elements in place
  end
end
puts a1

#each with index

[10,20,30,40].each_with_index do |number, index|
  puts "#{number} #{index}"
end

#iterating over an array with while or until
animals = ["cat", "dog", "parrot"]

i = 0
while i < animals.length
  puts animals[i]
  i += 1
end

i = 0
until i == animals.length #nice one -- remember, until only runs if this evaluates to false
  puts animals[i]
  i += 1
end

#map and collect -- they are aliases and accomplish the same thing!
#useful to do a consistent operation to each element in the array it seems
animals = ["cat", "dog", "parrot"]

lenghts = animals.map { |animal| animal.length } #returns an array
p lenghts #woah

#select and reject methods
words = ["racecar", "selfless", "sentences", "level"]
palindromes = words.select {|word| word == word.reverse} #some boolean evaluation which tells ruby if it should keep or not the element in the array
puts palindromes #dang, noice

#reject does the opposite
animals = ["cat", "dog", "parrot", "cow"] 
#lets exclude animals with the letter c
puts animals.reject {|animal| animal.include?("c")}

#patition: splits an array into two arrays, based on matching/not matching a condition
numbers = [1,2,3,4,5,6]
evens, odds = numbers.partition {|number| number.even? }
p evens
p odds

#any? and all? methods
#checks if any method fits a condition, or check if all methods fit a condition! always returns a boolean!
sports = ["soccer", "tenis", "baseball", "tennis", "golf"]
p sports.any? { |sport| sport.length == 8 }
p sports.all? { |sport| sport.length < 100 }

#find and detect -- they're both aliases. Finds first array that matches condition
animals = ["cat", "dog", "parrot", "cow", "cat"] 

p animals.find { |animal| animal.include?("o")}

#index and find_index -- also aliases! 
#Locates the index position of a given element
p animals.index("cat")

#include method -- also exists on an array!
p "action".include?("act")
p animals.include?("dog")

#max and min methods -- extract the largest numeric value or the longest string
p numbers.max
p animals.max
p numbers.min
p animals.min

def custom_max(numbers)
  max = 0
  numbers.each {|number| max = number if number > max}
  return max
end

p custom_max(numbers) 

#unlimited method arguments
def adder(*nums) #asterisk here is called "sponge" or "splat" -- means it groups any number of arguments in an array
  sum = 0
  nums.each {|num| sum += num}
  return sum
end

p adder()
p adder(1)
p adder(1,2)
p adder(1,2,3)

#you can however guarantee a first arugment, and then sponge everything on the second one
def adder_name(name, *numbers)
  sum = 0
  puts "Doing operation for #{name}"
  numbers.each {|num| sum += num}
  return sum
end

p adder_name("daniel",5,5,5)