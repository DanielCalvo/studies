require "faker"

puts Faker.class
puts Faker::Name.class

puts Faker::Name.name #Gives you a random name
puts Faker::Name.first_name
puts Faker::Name.last_name

p Faker.constants

puts Faker::FunnyName.name

#These are funny, ha!
puts Faker::Quote.famous_last_words

puts Faker::Quote.fortune_cookie