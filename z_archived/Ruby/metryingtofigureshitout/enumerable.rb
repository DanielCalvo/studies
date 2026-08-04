=begin These Ruby core classes include (or extend) Enumerable:
ARGF
Array
Dir
Enumerator
ENV (extends)
Hash
IO
Range
Struct

These Ruby standard library classes include Enumerable:
CSV
CSV::Table
CSV::Row
Set
=end

#Lets check an example using dir
dir = Dir.new("/home/daniel") #returns an array

dir.each_with_index do |entry, index| #so this is the same as iterating over any other array!
  puts "#{index} - #{entry}"
end