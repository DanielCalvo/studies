def hash_from_arrays(arr1, arr2)
  result = {}
  arr1.each_with_index do |val,index|
    result[arr1[index]] = arr2[index]
  end
  return result
end

a1 = []
hasherino = hash_from_arrays(["a","b","c"],["1","2","3"])
p hasherino

# Define a word_frequency method that accepts a piece of text.
# Return a hash with a count of the number of times each word
# appears within the text. The hash keys should be the words
# and the values should be their counts. Assume the text will
# be in all lowercase.
#
# Examples:
# The => indicates the expected return value
# word_frequency("blue red blue green")  => {"blue"=>2, "red"=>1, "green"=>1}
# word_frequency("a land far far away")  => {"a"=>1, "land"=>1, "far"=>2, "away"=>1}
# word_frequency("")                     => {}
def word_frequency(words)
  result = Hash.new(0)
  words.split(" ").each do |word|
    result[word] += 1
  end
  return result
end

words = "There is no banana in the banana jungle but there are apples. There are also no grapes and there are no oranges. There might be a coconut"
p word_frequency(words)

#Cheatcode on the above however, there's a tally method on arrays that does exactly what we want!
p words.split.tally #dang, noice