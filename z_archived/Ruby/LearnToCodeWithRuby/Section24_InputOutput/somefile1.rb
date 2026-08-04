
puts "welcome to the program!"
#by default ruby looks for the file in the current directory!
#load "somefile2.rb"

#Require will only load the file once -- running multiple times does nothing
#You need ./ here, otherwise ruby looks for this in the ruby install!
require "./somefile2" #require automatically looks for a file with the .rb extension
require "./somefile2"
require "./somefile2"
require_relative "somefile2" #looks to load a file relative to the currnet directory, so you can omit the ./

some_method
table = Table.new
puts table.class

puts "This program is run top to bottom!"

#You can call load multiple times if you need for some reason:
#load "somefile2.rb"


#require only opens the file one and it caches you
#require starts by looking at the ruby installation directory though, not your local dir