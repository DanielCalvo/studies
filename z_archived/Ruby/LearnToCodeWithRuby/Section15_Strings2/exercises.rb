require "debug"

def longest_word(string)
  #didn't initialize words in here though! 
  result = ""
  words = string.split(" ")
  words.each {|word| result = word if word.length >= result.length }
  return result
end

 p longest_word("There is no spoon and also no banana")

#splitting  a string -- chars and each_char methods -- very similar to .each on an array


#keed adding characters to a string until you reach the divisor, then append that string to an array
# and reset it
def custom_split(string, divisor)
  result = []
  word = ""
  string.each_char do |char|
    if char != divisor
      word << char
    else 
      result << word if word != ""
      word = ""
    end
  end
  if word != "" #gets last word
    result << word
  end
  return result
end

p custom_split("The dawg is not lazy", "a")
p custom_split("coma,separated,values,yay!", ",")
p custom_split(" hello", " ")

# Define a custom_join method that accepts an array of strings
# and a delimiter. The method should merge/join the array elements
# together into a single string. It should insert the delimiter
# in between every two subsequent elements. Do not use the
# built-in join method in your solution.
#
# Examples:
# The => indicates the expected return value
# custom_join(["red", "green", "blue"], "!") => "red!green!blue"
# custom_join(["Big", "Mac"], "$$")          => "Big$$Mac"
# custom_join([], "$$$")                     => ""

def custom_join(array, delimiter)
  result = ""
  array.each do |element|
    result << element 
    if element != array.last
      result << delimiter
    end
  end
  return result
end

p custom_join(["red", "green", "blue"],"!")

# Define a custom_count method that accepts a string and a string
# of search characters. The method should count how many times the
# search characters appear in the original string. Do not use the
# built-in count method in your solution.
#
# Examples:
# The => indicates the expected return value
# custom_count("Hello World", "l")     => 3
# custom_count("Hello World", "O")     => 0
# custom_count("Hello World", "z")     => 0
# custom_count("Hello World", "lo")    => 5
# custom_count("Hello World", "ol")    => 5

#iterate over string
#does character match one of the searchers? if so, counter++!
def custom_count(string, search)
  count = 0
  string.each_char do |string_char| 
    search.each_char do |search_char| 
      count += 1 if string_char == search_char
    end
  end
  return count
end

puts custom_count("banana", "na")

# Define a custom_index method that accepts a string and a search term.
# The method should return the first index position of the 
# search term within the string. If the search term does not exist,
# return nil. Do not use the built-in index method in your solution.
#
# Examples:
# The => indicates the expected return value
# custom_index("I am very handsome", "I")     => 0
# custom_index("I am very handsome", "e")     => 6
# custom_index("I am very handsome", "Z")     => nil
# custom_index("I am very handsome", "am")    => 2
# custom_index("I am very handsome", "ma")    => nil

#This is how you can do it with just 1 element, but how about 2?
def custom_index1(string, search)
  str_i = 0
  src_i = 0
  while str_i < string.length
    return str_i if search == string[str_i]
    str_i += 1
  end
  return nil
end

#use string indexing you dummy!
def custom_index(text, search)
  search_length = search.length
  text.chars.each_with_index do |char, index|
    if text[index, search_length] == search
      return index
    end
  end
  nil
end

puts "Custom index1:" #Works for 1 character:
puts custom_index1("joe","o") 

p custom_index("joe","oe")
p custom_index("banana","na")
p custom_index("McDonalds","M") 
p custom_index("McDonalds","ds") 

#I think you could've used contains? in here

def custom_delete(text, search)
  result = ""
  text.chars do |text_char|
    match = false
    search.chars do |search_char|
      if search_char == text_char
        match = true
        break
      end
    end
    if match == false
      result << text_char
    end
  end
  return result
end
#wew the authors solution was way better than mine!
#he iterated over text and used and used include? on characters to delete
#you thought this was a many-to-many thing but he saw it as a one-to-many, much easier to solve!

#iterate over text
#iterate over search

p custom_delete("mcdonalds", "ma")