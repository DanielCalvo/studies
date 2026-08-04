if true
    puts "hi"
end

one = true
two = true

if one && one
    puts "both are true"
end

if true || false
    puts "true of false evaluates to true"
end

user = "joe"

if user == "joe"
    puts "welcome joe"
elsif user == "bob"
    puts "welcome bob"
else
    puts "welcome user"
end