
3.times {print "hi"}
puts
count = 1

while count < 5
    puts count
    count +=  1
end

letters = "a"

while letters.length < 5
    letters << "a"
end
puts letters

#a while loop executes while a condition is true
#an until loop executes until a condition is true (executes as long as the condition evaluates to false)

i = 1
until i == 10
    print i
    i += 1
end

# Next keyword, similar to the C/Go Continue
puts
s = "banana"
i = 0

while i < s.length
    if s[i] == "a"
        i += 1
        next
    end
    puts s[i]
    i += 1
end

str = "This is a string with $, a dollar sign!"

i = 0
while i < str.length
    if str[i] == "$"
        puts "Found a dollar sign at position #{i}"
        break
    end
    i += 1
end

#Recursion, uh-oh!