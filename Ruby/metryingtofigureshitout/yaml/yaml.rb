#Lets parse some yaml file

require 'yaml'

data = YAML.load_file('file2.yaml')
puts data.inspect
puts data.class #Yaml gets loaded into a cash!

puts data['someitem']
puts data['list'][1]
puts data['person']['age'] #neat!

data['person']['surname'] = 'mcbobson'

#OH WAIT THERE'S .to_yaml?
puts data.to_yaml #holy bananas

# File.write('/tmp/test.yml', data.to_yaml)

# Here's a challenge: Recursively go through all of the contents of a yaml file!
data = YAML.load_file('pod.yaml')

# YAML.load_file gets you a hash, so what you need here is to figure out how to go through all of a hash... which can contain plain entries and arrays right?
# A hash entry can contain: text, another hash, or an array
puts data.class
puts

#well hey at least you managed to iterate over all of it, now you just gotta handle it properly!
def recursive_yaml_print(yaml)
  puts yaml
  sleep(1)
  if yaml.class == String
    puts yaml
  end

  if yaml.class == Hash
    yaml.each do |key, value|
      recursive_yaml_print(value)
    end

  end
  if yaml.class == Array
    yaml.each do |element|
      recursive_yaml_print(element)
    end
  end
end

#so if the above is a string i want to print it and be done
#if its a hash, i want to iterate over all of it again
#if its an array, same!
#ah, this is a recursion problem!
data.each do |key, value|
  #  puts "#{key}: #{value} - #{value.class}"
  if value.class == String
    puts "#{key}: #{value}"
  elsif value.class == Hash || value.class == Array
    recursive_yaml_print(value)
  end
end

