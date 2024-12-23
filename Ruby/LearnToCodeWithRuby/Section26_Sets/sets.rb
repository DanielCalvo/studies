#A set is an unordered collection of unique elements. It guarantees that all elements are unique.
#Sets solve the problem of duplicate elements in an array.
#Sets are useful when you want to know if an element is in a collection or not.
#Ruby 3.2 automatically includes the Set class. Before that you had to require it.

myset = Set.new([1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5]) #{1, 2, 3, 4, 5}
p myset

#If you try to add a duplicate element, it will be ignored, sets guarantee uniqueness
seasons = Set.new(["Spring", "Summer", "Fall", "Winter", "Spring"]) 
p seasons
p seasons.length
p seasons.include?("Spring")

#Order is not guaranteed in a set

#Add and delete elements
myset2 = Set.new([1, 2, 3, 4, 5])
myset2.add(6)
p myset2
myset2.add(6) #no difference
p myset2

myset2.delete(6)
p myset2

#explore sets in github
#https://github.com/ruby/ruby/blob/master/lib/set.rb
