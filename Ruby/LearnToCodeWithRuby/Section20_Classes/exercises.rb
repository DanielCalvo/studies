# Define a Cookie class within the file.
# Declare a create_cookie method that returns a Cookie object/instance
# Declare a multiple_cookies method that returns an array of
# two separate Cookie objects

class Cookie
end

def create_cookie
  return Cookie.new
end

def multiple_cookies
  result = []
  result << Cookie.new
  result << Cookie.new
  return result
end

# Declare a Musical class that includes "name",
# "cast", and "duration" instance variables.
# Feel free to initialize the instance variables
# to whatever values you'd like.
class Musical
  def initialize
    @name = "The UFO"
    @cast = ["Alien 1", "Alien 2"]
    @duration = 30
  end
end

# Define a Cake class. It will have 3 instance methods.
#    - A bake method that returns the string "Baking the cake"
#    - A slice method that returns the string "Slicing the cake"
#    - A sell method that return the string "Sold the cake"

class Cake
  def bake
    return "Baking the cake"
  end
  def slice
    return "Slicing the cake"
  end
  def sell
    return "Sold the cake"
  end
end

# A Computer class is defined below.
#
# Define a to_s instance method to customize a Computer object's string representation.
#
# The method should return the following string:
#   'A powerful Intel i7 computer with 64GB memory and 2 TB of storage'
#
# The string should incorporate the 3 instance variables.

class Computer
  def initialize
    @cpu = 'Intel i7'
    @memory = 64
    @storage = '2 TB'
  end
  def to_s
    "A powerful #{@cpu} computer with #{@memory}GB memory and #{@storage} of storage"
  end
end

# Define an Airplane class below.
#
# The initialization should define 3 instance variables:
#  - @maker set to "Boeing"
#  - @model set to 757
#  - @seats set to 60
#
# Define 3 getter methods (maker, model, seats) to return
# the value of the respective instance variables

class Airplane
  def initialize
    @maker = "Boeing"
    @model = 757
    @seats = 60
  end
  def maker
    @maker
  end
  def model
    @model
  end
  def seats
    @seats
  end
end

a = Airplane.new
puts a.maker


# Define an Project class below.
# The instantiation should define 3 instance variables:
#  - @name set to "Q4 Tech updates"
#  - @budget set to 100000
#  - @team_members set to ["Piers", "Rob", "Jon"]
#
# Define 3 getter methods (name, budget, team_members) to return
# the value of the respective instance variables
#
# Define 1 setter method (budget=) to update the value 
# of the @budget instance variable
class Project
  def initialize
    @name = "Q4 Tech Updates"
    @budget = 100_000
    @team_members = ["Piers", "Rob", "Jon"]
  end
  def name
    @name
  end
  def budget
    @budget
  end
  def team_members
    @team_members
  end
  def budget=(budget)
    @budget=budget #all valid, do note that budget and @budget are different!
  end
end

pr = Project.new
p pr.name
pr.budget=500
p pr.budget

# Define an FinancialTransaction class below.
#
# The instantiation should define 4 instance variables.
# Arguments to initialize should provide initial values for all 4 variables.
#  - @to
#  - @from
#  - @amount
#  - @completed
#
# Define getter methods for @to, @from, and @amount
# Define getter + setter methods for @completed
#
# Sample use:
# my_rent = FinancialTransaction.new("Landlord", "Boris", 1000, false)
# p my_rent.to
# p my_rent.from
# p my_rent.amount
# p my_rent.completed
# my_rent.completed = true
# p my_rent.completed

class FinancialTransaction
  attr_accessor :completed
  attr_reader :to, :from, :amount
  def initialize(to, from, amount, completed)
    @to = to
    @from = from
    @amount = amount
    @completed = completed
  end
end

fi = FinancialTransaction.new("Joe", "Bob", 10, true)
p fi.amount