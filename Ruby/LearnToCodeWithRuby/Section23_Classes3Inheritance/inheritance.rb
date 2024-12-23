#In ruby, a class can only inherit from one class, but the inheritance can go many levels, ex:
# vehicle > car > sportscar

#define a subclass that inherits from a superclass

class Employee
  attr_reader :name
  attr_accessor :age
  def initialize(name, age)
    @name = name
    @age = age
  end
  def introduce
    "Hey I'm #{name} and I'm #{age} years old"
  end
  def to_s

  end
end

em = Employee.new("Joe", 18)
p em.introduce


#subclass!
#manager class inherits from employee class!
#they're going inherit the initialize method, the variables and the introduce method!

class Manager <  Employee
  attr_reader :rank
  def initialize(name, age, rank)
    super(name, age)
    @rank = rank
  end

  def yell
    "I'm the boss!!!1!one"
  end
  def introduce
    result = super #calls introduce on employee and stores its result on result!
    result + " I am also a manager!"
  end
end

class Worker < Employee
  def clock_in(time)
    "Starting my shift at #{time}"
  end
end

bob = Manager.new("Bob", 100, "SVP")
dan = Worker.new("Dan", 10)

p dan.class
p bob.class

#superclass and ancestor class methods

#p 5.superclass #there's no superclass on this one!
p 5.class.superclass #ooooo integer inherits from numeric!
p 5.class.superclass.superclass #numeric inherits from Object!
p 5.class.superclass.superclass.superclass #BasicObject
p 5.class.superclass.superclass.superclass.superclass #nil -- basicobject has no superclass!
p 5.class.ancestors #returns all the superclasses and some modules apparently

p 3.14.class.ancestors
p "hello".class.ancestors
p [].class.ancestors

#Check inheritance hierarchy of our subclasses -- these also work on our classes!
puts
#puts Manager.superclass
puts Manager.ancestors

puts Manager < Employee #evaluates if manager inherits from employee, which it does!
puts Employee < Worker #false, worker doesn't inherit anything!
puts Manager < Object

#The is_a? and instance_of? methods
puts bob.instance_of?(Manager) #Only true if it matches the class that the object is made from. Does not match superclasses.
puts bob.instance_of?(Employee) #false

puts
puts bob.is_a?(Manager) #is_a? matches the class that the object is made from, and any other superclasses it might belong to!
puts bob.is_a?(Employee)
puts bob.is_a?(Object)
puts bob.is_a?(BasicObject)
puts bob.is_a?(Worker) #nope

#the methods method -- gives you an array of all the object's available methods
integer_methods = 5.methods.sort
float_methods = 3.14.methods.sort
p float_methods & integer_methods #methods present in both classes
p float_methods - integer_methods #all of the methods that exist on a float but not on an integer!
p integer_methods - float_methods

#exclusive instance methods in subclasses, uh-oh!
p bob.yell
p dan.clock_in("10:00")

#override methods in a subclass!
#if you define a method with the same name in a subclass as in the superclass, the one in the subclass will win over!
#So ruby will look for a method in a bottom-to-top approach: First in the manager, then in an employee, then in an object and so on
p bob.introduce
puts
#the super keyword -- invokes the method with the same name in the super class, uh-oh
p bob.introduce

#there are 3 ways to use the super keyword
#1. Without parentheses, super passes all subclass method's arguments to the superclass methods
#2. With parentheses and no arguments, super passes no arguments to the superclass method
#3. With parentheses and arguments, super passes those arguments to the superclass method

class Car
  attr_reader :make

  def initialize(maker)
    @maker = maker
  end
end

class Firetruck < Car
  attr_reader :sirens
  def initialize(maker, sirens)
    super(maker)
    @sirens = sirens
  end
end

ft = Firetruck.new("Ford", 5)
p ft.sirens

#defining equality == the instance method to define object equality!

class IceCream
  attr_reader :flavor, :calories, :price
  def initialize(flavor:, calories:, price:)
    @flavor = flavor
    @calories = calories
    @price = price
  end
  #Our own custom logic for comparison! icecreams are to be equal if they have the same price and calories
  def ==(other)
    calories == other.calories && price == other.price
  end
end

class Candy
  attr_reader :calories, :price
  def initialize(calories:, price:)
    @calories = calories
    @price = price
  end
end

cookies_and_cream = IceCream.new(flavor: "Cookies and cream", calories: 300, price: 3.99)
rum_raisin = IceCream.new(flavor: "Rum Raisin", calories: 300, price: 3.99)

p cookies_and_cream == rum_raisin

#ruby has no way of establishing equality, it'll only consider the same object as an equal to itself, even objects with the same fields will not be identical
cookies_and_cream2 = IceCream.new(flavor: "Cookies and cream", calories: 300, price: 3.99)
p cookies_and_cream2 == cookies_and_cream #now its true

#duck typing -- author says that if the object has the methods you want, its probably good enough!
#Ex: If you are eating sushi, you can use a fork or chopsticks. Both are eating utensils that you can eat sushi with
#with the == comparison above on icecream, we're just comparing calories and price. It doesn't even have to be an icecream object! It can be, but doesn't have to.

sour_patch = Candy.new(calories: 300, price: 3.99)
puts 
p cookies_and_cream == sour_patch