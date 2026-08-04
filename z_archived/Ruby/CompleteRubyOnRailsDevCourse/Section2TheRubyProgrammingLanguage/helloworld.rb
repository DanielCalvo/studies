#Ah, looks like you need to put this before main or before you call it
def say_hello(thing_to_say)
  puts thing_to_say
end

puts "Hello world"
p "Hello world" #Prints the quotes
print "Hello world" #No new line, neat

greeting = "Hello world"
puts greeting
say_hello("heya")

