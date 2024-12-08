#split on a string -- splits based on a delimiter
sentence = "There is no spoon"
p sentence.split(" ") #assumes space if no argument is passed

my_csv = "1,2,3,4,5"
p sentence.split(",")

vehicle = "spaceship"
characters = vehicle.split("")
p characters

characters = vehicle.chars #same deal
p characters

characters.each { |character| puts character} 

vehicle.each_char { |char| print char }  #an iteration specifically for string characters!

#the join method on an array!
puts
registrants = ["bob", "joe", "john"]
p registrants.join(" ")

p registrants.join("-").split("-")
p registrants.join(", ")

#count method on a string -- count the number of occurences of a character in a string
puts "Hello world".count("l")
puts "Hello world".count("lo") #not a sequential search, counts the ammount of time both characters appear

#index and rindex methods
str = "There are bananas in the kitchen"
p str.index("k")
p str.index("n")
p str.index("e") #You get the first match in the string
p str.index("here") #You get the starting index of the first occurrence

p str.index("e",10) #You can start your search 3 index positions in

str.rindex("e") #starts looking from the end of the string