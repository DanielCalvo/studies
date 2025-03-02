#Righty so I have these two ranges, how do I generate an array of chess board coordinates?
RANKS = (1..8)
FILES = ("A".."H")

#ChatGPT suggested:
# formatted_combinations = RANKS.product(FILES).map { |rank, file| "#{rank}#{file}" }
# puts formatted_combinations.inspect
# But what is even going on there? Let's find out!

#https://ruby-doc.org/core-3.0.1/Array.html#method-i-map

#Woah, product returns a combination of elements from all arrays!

my_array =  [1,2,3].product(["a", "b", "c"])

some_array = ["a", "b", "c"]

#not as idiomatic in ruby, the each method is more common
for element in some_array
  print element + '\n'
end

#the very common each method
some_array.each do |element|
  puts element
end

#The map method can be used to create a new array based on the original array, but with the values modified by the supplied block.
#Let me use that to convert an array of numbers to strings

arr1 = [1,2,10,20]
arr1.map! do |element|
  element.to_s
end
p arr1

#holy bananas look at this, you can do a loop during assignment
squared_numbers = [1, 2, 3].map do |num|
  num * num
end
puts squared_numbers.inspect #returns a string representation of the object, particularly useful for debugging!

upper_case = ["a", "b", "c"].map { |string| string.upcase }
puts upper_case.inspect #woah

#You can use the map method to transfor each element and return a new array!