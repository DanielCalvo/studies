#reverse
p [1,2,3,4].reverse #does not mutate the equivalent array

queue = [1,2,3,4]
queue.reverse! #this does though
p queue

#sort!
a = [5,22,34,5,55,11,4,0,-1,923]
p a.sort #ascending order
p a.sort.reverse #descending

b = ["asdf", "kjfddf", "ooofdf", "Zf"] #sorts alphabetically. Ruby considers capital letters to come first!
b.sort!
p b

#uniq
a = [1,11,1,2,2,33,33,3,123]
p a.uniq #returns an array without duplicates. There's also a ! version that will mutate the array

#compact -- removes all nil values, neat!
a = [1,2,3,44,2,nil,1,2,nil]
p a.compact

#inject and reduce -- aliases -- they derive a new value by combining all array elements!
#classic example: lets sum all the elements on an array

p [10, 20, 30].inject(0) {|sum, number| sum + number}

# first block variable: the aggregate value, the value being calculated
# second blog variable: the current array element
# block calculation: what to send to the next loop as the aggregate value!
# the argument to reduce is the starting value of the aggregate value
color_counts = ["Red", "Blue", "Red"].reduce({}) do |counts, color|
    if counts[color].nil?
        counts[color] = 1
    else
        counts[color] += 1
    end
    counts
end
p color_counts #dang this is noice

#the flatten method -- creates a single dimmensional array on an array!
attendees = ["joe", ["alice", "bianca"], "bob", [["jo"]]]
p attendees.flatten

#sample -- extract on eor more random elements from an array!
numbers = [1,2,3,4,5,6,7,8]
p numbers.sample #if no arguments, returns a single thing
p numbers.sample(2) #if you want 2 elements though, you get an array!
p numbers.sample(20) #you get all of the array if you ask for more samples than there are elements

#multiply an array with the asterisk!
p [1,2,3] * 3 #neat, duplicates the elements 3 times

#merging arrays and exclude duplicates!
p [1,2,3,4] | [3,4,5,6] # ooo the | is the union operator! union removes duplicates
p [1,2,3,4].|([3,4,5,6]) #goodness gracious the vertical pipe IS A METHOD, WOAH

#you can also use the minus sign to subtract elements in the first array, like so:
p [1,2,3,4] - [1,2,3] 
#You guessed it, its a method:
p [1,2,3,4].-([1,2,3])

#Array intersection with the ampersand method -- keep elements that are found on both arrays
p [1,2,3,4] & [3,4,5,6] 