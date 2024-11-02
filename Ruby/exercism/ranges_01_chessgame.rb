module Chess
  # TODO: define the 'RANKS' constant
  # TODO: define the 'FILES' constant

  #Its (), not [] you dummy!
  RANKS = (1..8)
  FILES = ("A".."H")

  def self.valid_square?(rank, file)
    raise "Please implement the Chess.valid_square? method"
  end

  def self.nick_name(first_name, last_name)
    raise "Please implement the Chess.nick_name method"
  end

  def self.move_message(first_name, last_name, square)
    raise "Please implement the Chess.move_message method"
  end
end



puts 1..5
print (1..5).to_a #to array

#Oh there is also a range operator
puts Range.new(1,5).to_a

#When you want to call a method on a range in ruby wrap it in a parenthesis otherwise you'll call it on the 2nd argument of the operator

#You can use ranges to slice a string:
puts "Hello world"[0..4]
puts "Hello world"[6..10]

#Oooo ranges have methods!
puts (1..5).sum
puts (1..5).size
puts (1..5).include?(3)

#Endless range! Careful, can generate an infinite sequence if not used on a collection
puts "Hello world"[0..]
puts "Hello world"[..5]
puts "Hello world"[6..]

# You can also use ranges on strings but its behaviour can be unexpected, use with caution
puts ("aa".."ad").to_a

#How do I create a board with these?
RANKS = (1..8)
FILES = ("A".."H")

puts (RANKS)
ranks = (1..8)
puts ranks.to_a.map

a = [1..3]
puts a.to_a