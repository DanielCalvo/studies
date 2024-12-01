numbers = [1,2,3,4,5,6]
p numbers #prints everything in the same line, better than puts!
puts numbers.length

#ruby allows this, but its not a good practice
stuff = ["a", 1, false, 3.14]

#nested arrays!

spreadsheet = [
    ["Student", "Class", "Grade"],
    ["Joe", "First", "B"],
    ["Bob", "Second", "A"],
    ["Charlie", "Second", "A"]
]

p spreadsheet

first_row, second_row, third_row = spreadsheet #holy bananas, assigning the nested arrays
p first_row
p second_row
p third_row

#the %w shortcut
names = ["Jack", "Joe", "Jill"]
p names
names = %w[Jack Joe Jill Jones Jamie Jose]
p names #woah, noice!

#accessing and overwriting arrays by index position
p names[0]
if names[99] == nil
    puts "empty element gives you nil"
end

p names[-1] #last element
p names.slice(-1) #same deal
p names.[](0) #this works too, huh!
names[0] = "Bob"
p names[0]

names[10] = "Joseph" #everything until position 10 gets filled with nil
p names

#fetch!
#puts names.fetch(100) #this will raise an exception on purpose if you try to get something out of bounds
puts names.fetch(100, "Bobson") #if the index position does not exist, returns "Bobson"

w = %w[Joe Jack Ma]
p w[-1][0]

#Readding it here just for clarity

#accessing and overwriting multiple elements
names = ["Bob", "Joe", "Jill", "Jones", "Jamie", "Jose", "Joseph"]
p names[0,3] #starts at 3, pulls 3 elements
p names[3,10] #gets as much as it cans and ignores the rest

p names.slice(0,3)
p names[3,2] = ["aa","bb"] #start and index position 3 and replace the next 2

p names[3,2] = ["0","0"] #You can replace with less or more
p names

#values at method
p names.values_at(0, 4) #returns an array, elements at 0 and 4
p names.values_at(1,2,3) #you can ask for any elements.
p names.values_at(99, 98) #nil for positions that don't exist

#helper methods
p names.first #shortcut for 0
p names.first(2) #first 2 elements

p names.last(2) #last 2 in ascending order

#lenght, size and count
p "Length", names.length, names.size #These are equivalent

p names.count #counts everythingif without argument

p names.count("Bob") #But how many times is bob present in the array? neat

p [1,2,3,4,4,4,5,6].count(4)

puts
def split_in_two(a)
    if a.length.even?
        return a.first(a.length / 2), a.last(a.length / 2) 
    elsif a.length.odd?
        return a.first((a.length+1) / 2), a.last(a.length / 2) 
    end
end

b = [1,2,3,4,5]
p split_in_two(b)

#empty method
c = []
puts c.empty? #equivalent to comparing the length of the array to zero
puts c.nil? #false -- the nil method is available in every single object in ruby, useful for when you need to check if something is nul!
puts nil.nil?

#equality and inequality operators in array objects
a = ["a", "b", "c"]
b = ["a", "b", "c"]

puts a == b

#two arrays are considered equal if they have the same lenght, the same elements and in the same order

# the spaceship operator!
#can return -1, 0, 1 or nil! woah!
p 5 <=> 5 #they're equal so we get 0
p ["a", "b", "c"] <=> ["a", "b", "c"]

#you get -1 if the value on the left is smaller!
p 5 <=> 6
p [1,2,3] <=> [1,2,4] #left one is sorted earlier... eh dunno how i feel about this
p [3,2,3] <=> [1,2,4]

#you get -1 if the value on the right is smaller!
p 15 <=> 6

#you get nil if you can't compare both objects
p 15 <=> "15"

#push and shovel
a = ["a", "b", "c"]
a.push("d")
a.push("e", "f", "g")
a << "h" #shovel operator!
p a

#insert method -- inserts at a specified position
a.insert(0, "a")
a.insert(0, "1", "2", "3")
p a

#if you do an insert past the limit of the array, ruby is going to fill the gaps with nil!
a = [1,2,3]
a.insert(6, 6)
p a

#removes and returns one element from the end of the array
e = a.pop
p e, a

#but you can just pop
a.pop

#or you can pass an argument saying how many arguments you want from the end, and it'll give you an array with those
b = a.pop(2)
p b

#shift and unshift methods -- its like pop, but targets the beginning or the array

a = [1,2,3,4,5]
e = a.shift(2)
p a, e
a.shift
p a

#unshift adds an element(s) to the beginning of the array!
a.unshift(1,2,3)
p a #neat!