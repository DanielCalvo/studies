require 'bcrypt'


#Wew maybe I'm too much of a beginner but I found Go's modules were in general a bit better documented.
#But maybe I am a beginner.
my_password = BCrypt::Password::create("banana")

puts my_password.version
puts my_password.cost
puts my_password == "banana" #wew, how can this be true?
puts my_password == "not banana"