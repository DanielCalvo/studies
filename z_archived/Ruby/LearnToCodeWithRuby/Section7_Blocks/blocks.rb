#Blocks: a procedure, collection of code

3.times { puts "hi" }

#oooh you can do it with the do keyword on the times thing, how interesting!

value = 3.times do
  puts "woah, 3 times!"
  puts "value: #{value}" #not available in here this way, interesting
end

puts "value: #{value}" #3

3.times { |count| puts count } #zero indexed, noice

#up to and downto
5.upto(10) {|current| puts "The loop is now on #{current}"}
puts
10.downto(5) {|current| puts "The loop is now on #{current}"}

#the step method
1.step(100, 5) {|current| print "#{current} "}

0.step(100, 10) do |current|
  print "#{current} "
end

