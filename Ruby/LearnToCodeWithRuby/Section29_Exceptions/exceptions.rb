#the begin and rescue keywords

def sum(a, b)
  begin
    a+b
    rescue #the rescue case just by itself will catch any exception
      #rescue used this widely remembers me of go's if err != nil syntax
      "unknown"
  end
end


def sum2(a, b)
  begin
    a+b
    #rescue is a catch all though
    rescue => e #e is the exception object
      puts "Class name: #{e.class}"
      puts "Message: #{e.message}"
  end
end

def sum3(a, b)
  begin
    a+b
    #you can rescue a specific exception
    rescue TypeError => e #e is the exception object
      puts "Class name: #{e.class}"
      puts "Message: #{e.message}"
    rescue NoMethodError => e
      puts "We're going to do something different here"
      puts "Class name: #{e.class}"
      puts "Message: #{e.message}"
  end
end


puts sum(3,5)
puts sum(3,"5") #TypeError
puts sum2(3,"5") #TypeError
puts sum3(3,"5") #TypeError
puts sum3(nil,nil) #NoMethodError -- can't sum two nils.

#the retry keyword
#retry restarts the begin block from the beginning, neat!
#be careful not to get into an infinite loop with it
def sum4(a, b)
  begin
    a+b
    rescue TypeError => e
      a = a.to_i
      b = b.to_i
      retry
    rescue NoMethodError => e
      a = 0
      b = 0
      retry
    ensure
      puts "I'm always going to run!"
  end
end
puts sum4(3,11) #runs even if there is no exception :o
puts sum4(3,"5")
puts sum4(3,"9") 
puts sum4("9", nil) 

#the ensure keyword
#ensure is a piece of code that will always run at the end of the method, no matter what
#usually used for cleanup -- like closing a file or trying to read a db, so you close the connection

#alternative ways of using the begin and rescue keywords. You can use them at the top level!
#oooo the begin word is optional?

#You don't need the begin and rescue keywords, you can do stuff without them!
def sum5(a, b)
  a+b
  rescue TypeError => e
    a = a.to_i
    b = b.to_i
    retry
  rescue NoMethodError => e
    a = 0
    b = 0
    retry
  ensure
    puts "I'm always going to run!"
end
puts "sum5"
puts sum5(3,12)
puts sum5(3,"13")

#you can also do begin and rescue outside of the function
def sum6(a,b)
  a+b
end

begin
  sum6("3", 5)
rescue TypeError => e #you can't acccess a and b here, so you can't correct them!
  puts "There was a type error on sum6"
rescue NoMethodError => e
  puts "There was a NoMethodError on sum6"
ensure
  puts "wrapping it up!"
end


#custom exceptions -- remember to inherit from StandardError, otherwise rescue won't catch it
class OvenIsOffError < StandardError
end

#the raise keyword -- you can raise your own exceptions!
class Oven
  attr_accessor :state
  def initialize
    @state = "off"
  end
  
  def turn_on
    self.state = "on"
  end

  def bake(item) #in addition to giving raise a message, you can gave it a class
    #Then you can rescue your own custom error!
    raise OvenIsOffError, "You need to turn the oven on first!" if state == "off"
    puts "baking the #{item}"
  end
end

myoven = Oven.new
begin
  myoven.bake("cake")
rescue OvenIsOffError => e
  puts e.message
  puts "Turning the oven on and trying again!"
  myoven.turn_on
  retry
end

