
#A hash stores a pair of key and values

#Hash rules:
#Hash keys must be unique
#Hash values can contain duplicates
#Hash values are extracted by key, not by order

#Author says: Ruby does keep an order in the hashes behind the scenes but you're not supposed to rely on it!

empty_hash = {}
puts empty_hash
puts empty_hash.class
puts empty_hash.length
puts empty_hash.empty?

# => hash rocket! hah I like it

#people and their salaries
people = {
  "Joe" => 10,
  "Bob" => 1_000,
  "Alice" => 500
}

puts people.length

#people and their pets
pets = {
  "Joe" => ["Cat", "Dog"],
  "Alice" => ["Cat"],
  "Bob" => []
}

puts people["Joe"]
p people["Josh"] #If you don't have the key, you get nil

p people.fetch("Joe") #the difference is that this method will return an error if the key does not exist!
#You can provide a fallback value in case the key does not exist!
p people.fetch("Josh", 42)

#symbols! a lightweight, immutable ruby object that is used as an identifier!
#useful for when you need a name or label that does not need to change. Useful as a hash key!

puts :hello
puts :hello.class
puts :hello.methods.length #has a lot less methods, faster to create and lighter weight (according to author)
puts "heya!".methods.length

a = :hello
b = :hello

puts a.object_id
puts b.object_id
puts :hello.object_id #all symbols with the same name have the same object id, unlike a string!

#person = {
#  :name => "Joe", 
#  :age => 10,
#}

#shortcut! no hash rocket
person = {
  name: "Joe", 
  age: 10,
}

puts person[:name]
puts person[:age]

#ruby 3.1 shorthand hash syntax!
red = 230
green = 0
blue = 50

color = {red: red, green: green, blue: blue}
puts color[:red]

#when your key matches a variable in your program
color = { red:, green:, blue:}
puts color

#adding a new value to a hash
menu = {burger: 3.99, taco: 1.99, chips: 1.99} 
p menu.length

menu[:ham_sandwich] = 2.49 #do note that this will overwrite something if it exists
p menu

#there's a store method!
menu.store(:salmon, 5.99)
p menu

#iterating over a hash!
salaries = {
  director: 100,
  producer: 200,
  ceo: 300
}

#author says again: don't rely on the order! If you want to store an ordered list of things, use an array!
salaries.each {|position, salary| puts "#{position} earns #{salary}"}
salaries.each_key {|key| print "#{key} "}
puts
salaries.each_value {|value| print "#{value} "}

#ahh these return an array!
p salaries.keys
p salaries.values

#check for inclusion on a hash!
cars = {
  toyota: "aygo",
  kia: "soul",
  ford: "fiesta",
}

p cars.include?(:toyota) #true
p cars.include?("toyota") #false
p cars.key?(:kia) #key? and has_key? do the same thing, key? seems to be preferred
p cars.has_key?(:ford)

p cars.value?("aygo") #do note that this will only look on the values, not on the keys!
p cars.has_value?("aygo")
p cars.value?("soul")

#select and reject methods on a hash
cars = {
  toyota: "aygo",
  kia: "soul",
  ford: "fiesta",
}
p cars.select{|make, model| model.start_with?("a")}
p cars.select{|make, model| model.length == 4}
p cars.reject{|make, model| model.length == 4}

p cars.select{|make, model| make.to_s.include?("o")}

#how to convert a hash to an array and vice versa!

p cars.to_a #oooo each kv becomes an array -- this gives an array of arrays

calories = [
  [:salmon, 500],
  [:crackers, 150],
  [:rice, 100],
  [:butter, 900],
]

cal_hash = calories.to_h

#delete removes a kv pair by its key
p cal_hash.delete(:butter) #returns the key
p cal_hash #mutates the original hash

#merge method! lets us merge two hashes
market = {garlic: "3 cloves", milk: "10 gallons"}
kitchen = {bread: "2 slices", milk: "100 gallons"}

#does not mutate either hash -- just returns a new one
p market.merge(kitchen) #the slice that has preference in case of duplicates values is the one passed to the merge method

#this however will mutate the hash:
market.merge!(kitchen)
p market

#create a hash with a default value with Hash.new.
numbers = {}
p numbers[:pi] #You get nil if the key has no value

numbers = Hash.new(0) #default value to provide in the absence of a key
p numbers[:pi] #0

#Reference problems with hash.new!

#you'll get an empty array when the key does not exist. this is just what ruby will give you, doesn't mean the hash is populated by empty arrays!
team_members = Hash.new([])

team_members["My Soccer Team"] << "Bob"
team_members["My Soccer Team"] << "Joe"
p team_members["My Soccer Team"]
p team_members

p team_members["Other team"] #Holy bananas I accidentally modified the default value :o

#so hang on let me try this again
team_members = Hash.new([]) #reset the hash

#If what the author explained is correct, all non-existing hashes reference the same empty array
p team_members["a"].object_id
p team_members["b"].object_id

#So when you do this, what you are doing is actually appending bob to the default empty array, as the "Accounting" key does not exist on the hash!
team_members["Accounting"] << "Bob"
p team_members["Accounting"].object_id
p team_members["key that does not exist"].object_id #yup

#It seems you can change the behaviour of Hash.new for team members to return a separate empty array to each key that does not exist!
team_members = Hash.new do |hash, key|
  hash[key] = []
end
puts
p team_members["Accounting"].object_id
p team_members["key that does not exist"].object_id 


 