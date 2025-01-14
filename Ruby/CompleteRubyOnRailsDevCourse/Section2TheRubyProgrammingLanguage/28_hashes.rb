
myhash = {}

somehash = {"a": 1, "b": 2, "c": 3}
puts somehash
puts somehash[:"c"]

#arrays of keys and values
puts somehash.keys
puts somehash.values

somehash.each do |key, value|
  puts "The key is #{key} and the value is #{value}"
end

otherhash = {a: 1, b: 2} #a and b are symbols, not strings

otherhash.each do |key, value|
  puts "The class of key is #{key.class} and the class of value is #{value.class}"
end

otherhash[:c] = 2
puts otherhash[:c]

otherhash.select {|k,v| puts k.is_a?(Symbol) }

#oh so this works
otherhash.select {
  |k,v| puts k.is_a?(Symbol)
}

#and this works too -- I wonder if its a matter of style or preference
otherhash.select do |k,v|
  puts k.is_a?(Symbol)
end

otherhash.delete(:a) #dont forget the colon!
puts otherhash