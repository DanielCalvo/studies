#a module is a toolbox of classes, methods and/or constants
#It organizes functionality into "boxes"

#community convention is to use PascalCase for modules
module LengthConversions
  #puts self #prints LengthConversions
  #def LengthConversions.miles_to_feet(miles)
  #Its better to just use self
  #oh wait, if you remove self it'll by default look for the method within the module, so you can remove it too
  def self.miles_to_feet(miles)
    miles * 5280
  end

  def self.miles_to_inches(miles)
    feet = self.miles_to_feet(miles)
    feet * 12
  end

  def self.miles_to_centimeters(miles)
    inches = miles_to_inches(miles)
    inches * 2.54
  end
end

puts LengthConversions.miles_to_feet(42)
puts LengthConversions.miles_to_inches(42)
puts LengthConversions.miles_to_centimeters(42)

#modules with identical methods
module Square
  def self.area(side)
    side * side
  end
end

module Rectandle
  def self.area(length, width)
    return length * width
  end
end

puts Rectandle.area(2,4)
puts Square.area(4)

#import modules into the current file

require_relative "circle"
puts Circle.area(5)

#the math module!
puts Math
puts Math.class
p Math.methods
puts Math.sqrt(2)
#:: is for extracting the value of a constant 
puts Math::PI

#the uri and net/http module
require "uri"
require "net/http"

p URI.class

uri = URI.parse('https://www.google.com')
#p Net::HTTP.get(uri)

#a mixin is a module that we add to a class to add additional behaviour
class Refrigerator
  include Enumerable

  attr_reader :snacks, :drinks
  def initialize(snacks:, drinks:)
    @snacks = snacks
    @drinks = drinks
  end

  def items
    snacks + drinks
  end
  #To make the Enumerable mixin available for this class, we must define the each method, and then we get access to a bunch of methods in the enumerable mixin
  def each
    items.each { |item| yield item}
  end

end



myfridge = Refrigerator.new(
  snacks: ["Doritos", "Cheetos", "Potatoes"],
  drinks: ["Water", "Coke", "Pepsi"]
)

puts myfridge.items
p myfridge.sort
p myfridge.include?("Doritos")
p myfridge.any? {|item| item.length > 4}
p myfridge.all? {|item| item.length > 4}
#And so on, you can call all the enumerable methods on it https://ruby-doc.org/3.3.6/Enumerable.html

#The comparable mixin -- it enables you to compare your objects
class OlympicMedal
  include Comparable
  attr_reader :type

  def initialize(type:)
    @type = type
  end

  def <=>(other)
    medal_values = {gold: 3, silver: 2, bronze: 1}
    current_medal_value = medal_values[type] #number
    other_medal_value = medal_values[other.type] #number
    if current_medal_value < other_medal_value
      return -1
    elsif current_medal_value == other_medal_value
      return 0
    else
      return 1
    end
  end
end

#Ruby has no idea how to compare these
bronze = OlympicMedal.new(type: :bronze)
silver = OlympicMedal.new(type: :silver)
gold = OlympicMedal.new(type: :gold)

p bronze > gold
p gold > silver #woah
p gold == gold #woah

#mixin in your own module!
#protip
#use inheritance when the relationship of something is of type "is-a"
#ex: car is a type of vehicle

#use a mixin when the relationship is a "has-a" relationship
#ex car is towable, purchasable, crushable, etc
#you can mix in multiple modules but only inherit from 1 superclass 

module Purchasable
  def purchase(item)
    "#{item} has been purchased!"
  end
end

class Bookstore
  include Purchasable
end

class Supermarket
  include Purchasable
end

class Bodega < Supermarket
end

bookstore = Bookstore.new
supermarket = Supermarket.new
bookstore = Bookstore.new

puts bookstore.purchase("1984")

puts supermarket.purchase("Cheese")

#the ancestor method in depth
p Bookstore.ancestors #the ancestors array tells you the order in which ruby will look for methods
#so if a method runs its because ruby found it first in the ancestors array
p Object.class 
p Kernel.class
p BasicObject.class

#the prepend keyword -- to add something to the beginning of something else
#it allows you to add the mixins methods before the instance methods in the lookup order

module Purchasable2
  def purchase(item)
    "#{item} has been purchased!"
  end
end

class Bookstore2
  prepend Purchasable2
  def purchase(item)
    "#{item} has been purchased from the bookstore!"
  end
end

bookstore2 = Bookstore2.new
p bookstore2.purchase("Joe's biography")

#extend keyword -- add the mixins methods as class methods
#note: class, not instance methods!
module Announcer
  def who_am_i
    "The name of this class is #{self}"
  end
end

class Dog
  extend Announcer
end

class Cat
  extend Announcer
end

#Do note that this is a class method
p Dog.who_am_i
p Cat.who_am_i

#mixin in multiple modules
module A
  def some_method
    "Hello from A"
  end
  def whatever
    "Whatever from A"
  end
end

module B
  def some_method
    "Hello from B"
  end
end

#in the real world this might happen as you can aggregate a bunch of modules in the same class definition
class SomeClass
  include A
  include B
end

some_class = SomeClass.new
p some_class.some_method #ruby looks top down for the orders that the modules were mixed in, and the last module wins out
p some_class.whatever #for modules that don't have the same methods, this doesn't matter

#multiple declarations for the same module
#sometimes the folder name becomes the module name and all of the files in that folder will build the actual module!
require_relative "low_quality"
require_relative "high_quality"

class Song
  include Downloadable
end

mysong = Song.new
p mysong.download_low_quality
p mysong.download_high_quality

#modules within modules
module FileManagement
  module CSV
    class Reader
      #...
    end
  end

  module Excel
    class Reader
      #...
    end
  end
end

#The :: symbol is called the scope resolution operator. Its what you use to access modules inside modules
p FileManagement::CSV::Reader.new