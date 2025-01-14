my_array = [1, 2, 3, 4, 5]

#How can we loop over an array again?
#For loop:
for element in my_array
  print element
end

my_array.each do |element| #vertical bars are block parameters
  print element
end

(1..3).each do |element|
  print element
end

#Oh theres also each with index
my_array.each_with_index do |element, i|
  print i, element*2
end

puts my_array.each.class #This returns an enumerator!
puts (1..3).each.class #same deal, neat

for i in (1..3)
  print i
end

for i in (1..3) do print i end

#Ooh you can use next like a continue
for i in 1..5
  next if i == 3
  puts i
end








#Hey can we do a c-style loop? well, not really, ex: for (int i = 0; i < n; i++) isn't available.