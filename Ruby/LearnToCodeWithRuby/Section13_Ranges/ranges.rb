inclusive_nums = 1..5 #1 and 5 are included -- closed range
p inclusive_nums.class #Range!
p inclusive_nums

exclusive_nums = 1...5 #1 and 5 not included -- open range

puts inclusive_nums.first
puts inclusive_nums.last
puts exclusive_nums.first #both give you one, huh!

p inclusive_nums.first(3)
p exclusive_nums.first(3) #huh? prints the first element too

p inclusive_nums.last(3)
p exclusive_nums.last(3) #does not print the last one -- excluded

#author recomments wrapping the range in parenthesis if you're going to do things with it directly
p (2..10).last

#alphabetical ranges!
alphabet = "a".."z"
p alphabet.first(3)
p alphabet.last(3)

alphabet = "A".."z" #A-Z, a-z
p alphabet.first(40)

alphabet = "a".."z"
p alphabet.include?("j")
p alphabet.member?("j") #same idea, just a differenet name

# triple equal ===, checks for inclusion
p alphabet === "o"

#random numbers
puts rand
puts rand.round(2)
puts rand.round(2) * 30

puts rand(100) #100 exclusive, 0..99

puts rand(50..60) #Can receive a range, neat! Random number between 50 and 60

#extracrting elements from a string
a = "banana"
puts a[0..3]
puts a[0..99] #pulls as much as it can, does not return nils though
puts a[2..-1] #also works with negative values
a[2..5] = "baba" #Can also assign to the range in a string
puts a

puts a[2..5]
#puts a(2..5) #nope, works only on an array

#case with range!

#you can even have the case with on a single line!
def calculate_test_grade(grade)
    case grade
    when 90..100 then "A"
    when 80..89 then "B"
    when 70..79 then "C"
    when 60..69 then "D"
    else "F" #else doesn't need then!
    end
end

puts calculate_test_grade(72)

#converting a range to an array!
#to_a

p ("A".."Z").to_a #If you don't put the parenthesis, ruby gets confused!
arr = ("A".."Z").to_a
p arr[2]