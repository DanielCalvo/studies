#how to define our own methods that can accept blocks, neat!

#the yield keyword -- lets define a method that accepts its own block!

def pass_control
  puts "I'm at the start!"
  yield #yields to the attached block -- ruby will require a block!
  puts "After yield!"
  yield #you can yield as many time as you want!
end

pass_control {puts "Now I'm inside the block!"} 

#You can also use the syntax with do:
pass_control do
  puts "line 1"
  puts "line 2"
end

#block return values
def who_am_i
  puts "Hello there! Here's an adjective about me"
  adjective = yield
  puts "Given adjective was #{adjective}"
end

#don't use the return key
who_am_i { "persistent!"}
who_am_i { "silly"}

who_am_i do 
  "confused"
  "bamboozled" #only the last one is returned!
end

#who_am_i { return "persistent!"} #Dont do this! Ruby will pass the return key as well!

#the block_given? method -- meant to be invoked inside a method that may or may not receive a block!

def pass_control_on_condition
  puts "inside the method"
  yield if block_given? #much more flexible!
  puts "still inside the method!"
end

puts
pass_control_on_condition {puts "Inside the blockerinoooooo" }
puts
pass_control_on_condition

#block parameters!


def speak_the_truth(name)
  yield(name)
end
puts

speak_the_truth("Joe"){|name| puts "#{name} is trying!"} #name here could be whatever
#you get the me from the yield, interesting!


def number_evaluation(num1, num2, num3)
 yield(num1, num2, num3)
end

p number_evaluation(5,10,15) {|a,b,c| a+b+c}
p number_evaluation(3,4,5) {|a,b,c| a*b*c}

[10, 20, 30].each {|number| puts "The square of #{number} is #{number * number}"}

#lets do our own each!
def custom_each(elements)
  i = 0
  while i < elements.length
    yield(elements[i]) #woaahhh
    i += 1
  end
end

custom_each([10, 20, 30]) {|number| puts "The square of #{number} is #{number * number}"}

#But you can also do anything with any array!
custom_each(["Joe", "McJoe", "McBobson"]) {|string| puts "The length of #{string} is #{string.length}"}

#proc -- an object representation of a block. proc is short for procedure

to_cubes = Proc.new {|number| number ** 3}

#you can also use do - end
to_cubes = Proc.new do |number|
  number ** 3
end

#This is also valid
to_cubes = proc {|number| number ** 3}

#and so is this
to_cubes = proc do |number|
  number ** 3
end

a = [1,2,3,4,5]
b = [6,7,8,9,10]
c = [11,12,13,14,15]

p a.map{|number| number ** 3}
p a.map(&to_cubes) #You pass procs with the ampersand

#more examples of procs!
us_dollars = [1,2,3,4,5]

to_euros = Proc.new {|currency | currency*0.93}
to_rupees = Proc.new {|currency | currency*82.28}

p us_dollars.map(&to_euros)

is_senior = Proc.new {|age| age > 65}

ages = [10, 58, 69, 68, 23, 42, 90]

p ages.select(&is_senior)
p ages.reject(&is_senior)

#methods with proc parameters!

#Both are valid syntax
def talk_about(name, &my_proc)
  puts "Let me tell you about #{name}"
  my_proc.call(name)
end

def talk_about2(name)
  puts "Let me tell you about #{name}"
  yield(name)
end

good_thing = Proc.new{ |name| puts "#{name} is a jolly good fellow"}
bad_thing = Proc.new{|name| puts "#{name} is not cool"}

talk_about("Joe", &good_thing)
talk_about2("Bob", &bad_thing)
talk_about("Joel") {|name| puts "I can still pass my own block and talk about #{name}"}

talk_about2("Dan", &bad_thing)
talk_about2("Bob") {|name| puts "I can talk about #{name}!"} #you can also call it with a block!

#lambda -- a nameless method
#object, has a call method! similar to a proc!

squares_proc = Proc.new {|number| number ** 2}
squares_lambda = lambda {|number| number ** 2}
squares_lambda_alternative = -> (number) {number ** 2} #neat!


p [1,2,3].map(&squares_proc)
p [1,2,3].map(&squares_lambda) #Seems likeø the same deal but with a different keyword?
p [1,2,3].map(&squares_lambda_alternative)

#differences between lambda's and procs!

#a lambda cares about the number of arguments it receives!

my_proc = Proc.new {|name, age| puts "Your name is #{name} and you are #{age}"}
my_lambda = lambda {|name, age| puts "Your name is #{name} and you are #{age}"}

def do_stuff(&code)
  code.call("Joe", 123)
end

def do_more_stuff(&code)
  code.call("Joe")
end

do_stuff(&my_proc)
do_stuff(&my_lambda)

do_more_stuff(&my_proc) #proc just prints blank for age
#do_more_stuff(&my_lambda) #a lambda however will throw an exception if you pass the wrong number of arguments

my_proc = Proc.new { return "PROC RETURN"}
my_lambda = lambda { return "LAMBDA RETURN"}

def execute(&logic)
  puts "Starting"
  puts logic.call
  puts "Ending" #if you call with a proc, this line does not run
end

execute(&my_lambda) #when a lambda returns, it passes control back to the calling method. There's no way to stop the method
execute(&my_proc) #when a proc returns, it triggers a return from the calling method