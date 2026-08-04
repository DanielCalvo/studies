def introduce_myself
  puts "Hello!"
end

a = "asd"

#the variable a will not exist inside the method, interesting, it is not global by default, it seems to be scoped to the main function
def praise_person(name, compliment) #name is a local variable here
  puts "#{name} is doing #{compliment}!"
  a = "#{name} and #{compliment}" #local variable -- variables in methods are local by default
  puts a
end

praise_person("Bob", "awesome")
puts a
# praise_person calling it with no arguments gives an ArgumentError -- same deal if you call with too many arguments

#function that returns a value
# every method in ruby must have a return value, in the absence of a return value, ruby just returns nil
def add_two_numbers(num1, num2)
  # return num1 + num2
  num1 + num2 #implicit return
  #best practice: if you're going to return something on the last line of a method, you can ommit the return keyword
end

#default value for suffix if not provided, so that's what that is!
def title_assigner(name, suffix = "the happy")
  "#{name} #{suffix}"
end

puts title_assigner("joe")
puts title_assigner("joe", "the wise")