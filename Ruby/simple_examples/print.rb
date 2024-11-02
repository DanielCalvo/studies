#How do I format output again?
# Can I do c-style %s? #what print functions do I have?

puts "This automatically adds a newline at the end of an output."
print "This does not!\n"

#The p method displays an object raw value, interesting! Including any quotes or escape characters
name = "Alice"
age = 30
array = [1, 2, 3]
p name, age, array

#Printf baby!
printf "Name: %s, Age: %d\n", name, age

#inserting variable into a string:
puts "Name: #{name}, Age: #{age}"

#Heredocs, also called here documents
# << is called the shovel operator. Oh you don't even need quotes!
message = <<-BODY
"multi
line
string"
BODY

message2 = <<-BODY
holy
moly
no quotes!
BODY

message3 = %(also the percentage thing
works too. Hi, #{name}!)

puts message
puts message2
puts message3

#Squiggly heredoc, removes leading indentation
message4 = <<~BODY
  This is a multiline snippet.
  def print text
    puts text
  end
BODY

puts message4

#Using single quotes disables interpolation in heredoc messages. But ~ still removes leading indendation
content = <<~'CONTENT'
  An example of disabled variable interpolation:

  #{name}
CONTENT

puts content