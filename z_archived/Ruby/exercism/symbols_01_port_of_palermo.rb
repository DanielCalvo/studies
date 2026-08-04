module Port
  # The identifier are the first four letters of the name of the port, in uppercase.
  def self.get_identifier(city)
    return city[0..3].upcase.to_sym
  end

  def self.get_terminal(ship_identifier)
    if ship_identifier.to_s[0..2] == "OIL" || ship_identifier.to_s[0..2] == "GAS"
      :A
    else
      :B
    end
  end
end


#Symbols are named identifiers. They're used for keys in hashes, and to represend method and variable names.
:foo
puts :foo.object_id

puts "foo".object_id
puts "foo".object_id #Different identifiers

#Symbols are immutable, they cannot be modfied. When you modify a symbol, you're actually creating a new symbol.

p = Port
puts p.get_identifier("Madrid")

#Convers a string to symbol
puts "Banana".to_sym

puts p.get_terminal(:OIL123)
puts :OIL123.to_s[0..2]