require "debug"

candy = "sour patch!"

puts "I love eating #{candy}"

binding.break #works the same as debugger

beverage = "cola"
puts "I like #{beverage}"

binding.break #works the same as debugger

puts "This is the end of the program"


#continue navigates to the next breakpoint
#step seems to only step into the next part of the program
#info prints all variable values! (shortcut: i)

3.times do |count|
  puts "On loop number #{count}"
  puts "blabla"
  debugger
end