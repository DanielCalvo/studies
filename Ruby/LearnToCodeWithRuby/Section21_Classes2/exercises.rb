# Define a Product class below.
#
# A Product object should initialize with a @name and @price
# These values should by set to arguments to the initialize method
# @name will be a string and @price will be a floating-point value.
#
# Define getter methods for the 2 instance variables.
# Define setter methods for the 2 instance variables.
#
# When overwriting the product's price, the new price should be
# greater than 0. If it's not greater than 0, do not overwrite 
# the old price value.
#
# When overwriting the product's name, the new name should have a
# length between 3 and 20 characters. If it doesn't fulfill that
# criteria, overwrite the name to "TBD" instead.
#
# SAMPLE CODE:
#
# book = Product.new("1984", 9.99)
# p book.name # "1984"
#
# book.name = "Harry Potter"
# p book.name # "Harry Potter"
#
# book.name = "OK"
# p book.name # "TBD"
#
# p book.price # 9.99
#
# book.price = 24.99
# p book.price # 24.99
#
# book.price = -100
# p book.price # 24.99
 
class Product
  attr_reader :name, :price
  def initialize(name, price)
    @name = validate_name(name)
    @price = validate_price(price) 
  end
  #you can't use attr_writer to set the price because you can't add validation logic when using those!
  def name=(name)
    @name = validate_name(name)
  end
  def price=(price)
    @price = validate_price(price)
  end

  private
  def validate_name(name)
    if name.length >= 3 && name.length <= 20
      return name
    else
      return "TBD"
    end
  end
  def validate_price(price)
    if price > 0.0
      return price
    else
      return @price
    end
  end
end

book = Product.new("A", 2)
book.name = "B"
p book
book.name = "BBBB"
p book
book.price = 99
p book
book.price = -1
p book

# Define a SushiLunchOrder class below.
#
# A SushiLunchOrder object should initialize with @salmon, @tuna,
# and @yellowtail instance variables from arguments to the
# initialize method. All values will be integers representing
# the number of that type of fish in the lunch order.
#
# Define getter methods for the 3 instance variables.
#
# Define a salmon_special class method that instantiates a
# SushiLunchOrder instance with 6 pieces of salmon, 3 pieces
# of tuna, and 3 pieces of yellowtail.
#
# Define a family_combo class method that instantiates a
# SushiLunchOrder instance with 12 pieces of salmon, 12 pieces
# of tuna, and 12 pieces of yellowtail.
#
# Define a total_pieces class variable that keeps track of
# the TOTAL number of pieces of fish that have been sold.
# This is not the number of SushiLunchOrder instances but rather
# the sum of all the parts of fish.
#
# Define a total_pieces class method that exposes the value of
# the total_pieces class variable.
#
# EXAMPLE
# order1 = SushiLunchOrder.salmon_special
# p order1.salmon     # 6
# p order1.tuna       # 3
# p order1.yellowtail # 3
# p SushiLunchOrder.total_pieces # 12
#
# order2 = SushiLunchOrder.family_combo
# p order2.salmon     # 12
# p order2.tuna       # 12
# p order2.yellowtail # 12
# p SushiLunchOrder.total_pieces # 48
#
# order3 = SushiLunchOrder.new(3, 4, 5)
# p order3.salmon     # 3
# p order3.tuna       # 4
# p order3.yellowtail # 5
# p SushiLunchOrder.total_pieces # 60

class SushiLunchOrder
  attr_reader :salmon, :tuna, :yellowtail
  @@total_pieces = 0
  def initialize(salmon, tuna, yellowtail)
    @salmon = salmon
    @tuna = tuna
    @yellowtail = yellowtail
    @@total_pieces += salmon + tuna + yellowtail
  end
  def self.total_pieces
    @@total_pieces
  end

    def self.salmon_special
      new(6,3,3)
    end
    def self.family_combo
      new(12,12,12)
    end
end

su = SushiLunchOrder.salmon_special
su1 = SushiLunchOrder.family_combo
p su.salmon

p SushiLunchOrder.total_pieces

# Monkey-patch the Array class to add a more_than_once? predicate method
# The method should accept an argument representing an element
# The method should return true if the element occurs more than once
# within the array.
# 
# Example
# my_array = [1, 2, 2, 3]
# my_array.more_than_once?(2)    #=> true
# my_array.more_than_once?(3)    #=> false
#
#
# Monkey-patch the Hash class to add a common_keys_and_values method
# The method should return an array consisting of the elements
# that can be found among BOTH the hash's keys and values.
#
# Example:
# my_hash = { a: "hello", b: "goodbye", "goodbye" => 5 }
# p my_hash.common_keys_and_values  #=> ["goodbye"]
#

class Array
  def more_than_once?(element)
    counter = 0
    self.each do |el|
      if element == el
        counter += 1
      end
      if counter == 2
        return true
      end
    end
    return false
  end
end

class Hash
  def common_keys_and_values(myhash)
    #any key that also occurs as as value
    #any value that also occurs as as key

    result = Hash.new
  #if element is in array1 and if element is in array 2
  #return true
  #else return false
  end
end

puts "counting elements"
p [1,2,3].more_than_once?(2)
p [1,2,3,2].more_than_once?(2)
