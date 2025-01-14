puts 1+2
puts 10/4 #uhhh integer division, drops the decimal
puts 10.0/4
puts 10/4.to_f
puts (10/4).to_f #not gonna work as intended

puts "-"*20 #repeats the string 20 times
20.times { print "-"} #same deal as above, neat
puts
10.times { print " #{rand(10)}"} #kinda ugly but I get it
puts
puts "10".to_i
puts "banana".to_i #0, interesting

#Keep the data in its raw form and convert it when you need to work with it