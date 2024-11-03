#Lets parse some yaml file

require 'yaml'

data = YAML.load_file('yamlfile.yaml')
puts data.inspect
puts data.class #Yaml gets loaded into a cash!

puts data['someitem']
puts data['list'][1]
puts data['person']['age'] #neat!

data['person']['surname'] = 'mcbobson'

#OH WAIT THERE'S .to_yaml?
puts data.to_yaml #holy bananas
puts data.to_json #noice

# File.write('/tmp/test.yml', data.to_yaml)