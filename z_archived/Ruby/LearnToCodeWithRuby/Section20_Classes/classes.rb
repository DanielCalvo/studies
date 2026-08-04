#a class is a blueprint for creating objects!

#class method -- returns the class from which the object was made

puts 5.class
puts 5.class == 1.class
puts 5.4.class

puts "".class
puts [].class
puts ({}).class
puts true.class
puts false.class
puts nil.class
puts ().class
puts (0..9).class
puts Proc.new {}.class

#lets define our new custom ruby class! UpperCamelCase

#classes in ruby follow the UpperCamelCase naming conventnion
class Guitar
end

acoustic = Guitar.new #new creates a new object, all classes have it
puts acoustic #uh-oh, gets a class and a position in memory for the object
puts acoustic.object_id

p Hash.new
p String.new("hi")
p Array.new(1)

#when you call things like [], "" or {}, they are a shortcut for Array.new, String.new and Hash.new!
#But the ruby team did that so that i'd be convenient for you, neat!

#instance variables and the initialize method!

#@ sign saves the value inside the type variable in the object
class Guitar
  def initialize #initialize is a special one -- ruby looks for initialize methods. Ruby runs this method whenever this object is instantiated
    @type = "Acoustic" #these are private to other objects and the rest of the program!
    @wood = "Alder"
    @strings = 6
    @colors = ["black", "gold"]
  end
end

guitar_1 = Guitar.new
guitar_2 = Guitar.new
p guitar_1
p guitar_2

#instance methods
class Guitar
  
  #the self keyword!
  puts "Inside the guitar class: #{self}"
  def initialize #initialize is a special one -- ruby looks for initialize methods. Ruby runs this method whenever this object is instantiated
    @type = "Acoustic" #these are private to other objects and the rest of the program!
    @wood = "Alder"
    @strings = 6
    @colors = ["black", "gold"]
    @price = 1500
  end
  # do note that `information` will be available on instances for the Guitar class, but not class itself
  # you can't call Guitar.information
  def information #this one can be whatever we want
    "An #{@type} #{@wood} guitar with #{@strings} strings"
  end

  def to_s #overwrite the string representation!
    "An #{@type} #{@wood} guitar with #{@strings} strings"
  end
  def details
    #self is this instance of guitar
    #self then looks how to convert this instance of guitar to a string, and then ends up calling the to_s method
    #if you comment out to_s you'll get the hex values again related to this instance of guitar!
    puts "inside the details instance method: #{self}"
    #self allows you to query the object itself! you can also use it to call other instance methods
    puts "is it nil? #{self.nil?}. It is made from the #{self.class} class"
  end
  def nil_details
    puts "is it nil? #{nil?}" #you can also remove the self here
  end
  def class_details
    puts "It is made from the #{self.class} class" #you can't remove the self here as class is a ruby built in!'
  end
  def tell_details
    #if you ommit the self word, and you just call a method, ruby is going to try to by default to locate it in the current object!
    #self.nil_details
    #self.class_details
    nil_details
    class_details

    def type #have a getter with the same name of the variable! a good convention!
      @type
    end
    def wood
      @wood
    end
    def strings
      @strings
    end
    def price
      @price
    end

    #setters
    def price=(new_price) #setters are set by convention to `variable=`
      @price = new_price
    end
  end

  #getter method that reads/gets/retrieves the value of an instance variable

end

p guitar_1.information

#override the to_s method
puts guitar_1.class
puts guitar_1
puts guitar_1.to_s #ruby has no clue how to represent your object as a string
#we're free to customize it though!

guitar_1.details
p guitar_1.nil?
puts
guitar_1.tell_details

p guitar_1.strings
guitar_1.price=(1999)
p guitar_1.price

guitar_1.price = 2110 #shortcut for when you use the price= function name!
p guitar_1.price

#the attr_reader and attr_writer methods -- shortcuts for getters and setters

class Person
  #one symbol for each instance variable you want to create a reader for 
 # attr_reader :name, :age
  #same for create a setter (aka writer)
 # attr_writer :name, :age
#however, if you want readers and writers for variables you can use accessor:
  attr_accessor :name, :age
  def initialize(name, age) #initialize with arguments!
    @name = name
    @age = age
  end
  def to_s
    [@name, @age]
  end
end

per = Person.new("Bob", 99)
p per.name
per.name = "Joeson"
p per.name
puts
per = Person.new("Bob", 99)
p per.name, per.age
p per