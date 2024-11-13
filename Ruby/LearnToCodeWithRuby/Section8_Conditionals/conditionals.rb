if 5 < 7
  puts "yeah!"
end

if 5 < 3
  puts "nope!"
end

password = "topsecret"

if password == "topsecret"
  puts "Congrats, you are logged in!"
end

word = "banana"

if word.length == 6
  puts "banana has 6 characters"
end

#conditionals with predicate method

if 5.odd?
  puts "five is indeed odd"
end

if word.include?("ana")
  puts "ana is in banana"
end

if word.include?("zebra")
  puts "uhhh"
end

#truthy and falsy, oh my
if 5
  puts "uhhh five is truthy"
end

#in ruby there are only 2 falsy values: false and nil
# any numeric value is an example of a truthy value, any string is an example of a truthy value
# everything except nil and false is truthy:
if -9 && "woah" && 0 && ""
  puts "all true!"
end

color = "Green"

if color == "Red"
  puts "its red!"
elsif color == "Yellow"
  puts "its yellow"
elsif color == "Green"
  puts "its green"
end

num = 55

if num < 25
  puts "smol number"
elsif num > 50
  puts "not so smol number"
elsif num > 75
  puts "bigger number number" #does not run even though it would evaluate to true, only the first statement that matches runs
end

#good old else if
grade = "D"
# An elsif is not guaranteed to execute as it may evaluate to false, but the else always if nothing matches
if grade == "A"
  puts "woah got an A!"
elsif grade == "B"
  puts "eh good enough"
else
  puts "uhhh not really that good"
end

#multiple conditions
username = "joe"
password = "secret"

if username == "joe" && password == "secret"
  puts "user and password match!"
end

food = "toast"
price = 1
if food == "bread" || price < 2
  puts "ok"
end

#hey look this is interesting
meal_is_affordable = price < 2
food_is_toast = food == "toast"

puts meal_is_affordable, food_is_toast #so you could have an if statement with those two

# divisible_by_three_and_four(24)  => true

#parenthesis precedence
agent = "joe"
name = "mcbobson"
title = "secret agent"

if (agent == "bo" && title == "agent") || name == "mcbobson"
  puts "if matched"
end

##oo the ternary operator

#is 1 less than two? if yes, run the first thing right next to it, otherwise run what's on the other side of the : symbol
puts 1 < 2 ? "yeah" : "nope"

#OH BUT WAIT
value = 1 < 2 ? "yeah" : "nope"
puts value

value = agent == "joe" ? true : false
puts value

#simple example of calling a method from another method
def add(a,b)
  a + b
end

def calculator(a, b, operation)
  if operation == "add"
    add(a,b)
  end
end

puts calculator(8, 8, "add")

#the case statement!

food = "beans"

#you could put these inside a function and just return the strings too!
case food
when "banana"
  puts "I like bananas"
when "apples"
  puts "I like apples too"
when "kiwi"
  puts "Not much of a kiwi guy"
when "steak", "sushi", "chicken"
  puts "those are fine too, look, you can have comma separated options on a when!"
else
  puts "This evaluates when none of these have matched"
end