person = {
  name: "Alice",
  age: 30,
  city: "New York"
}

puts person
puts person.inspect

a="aaa"

person.each do |key, value|
  puts "#{key}: #{value}"
end

people = [
  { name: "Alice", age: 30, city: "New York" },
  { name: "Bob", age: 25, city: "San Francisco" },
  { name: "Charlie", age: 35, city: "Los Angeles" }
]

people.each do |person|
  puts "#{person[:name]} lives in #{person[:city]}."
end