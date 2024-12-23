#private methods! instance method that cannot be called outside of the object!

class SmartPhone
    attr_reader :username, :production_number

    def initialize(user,pass)
        @username = user
        @password = pass
        @production_number = generate_production_number
    end
    def password=(password)
        @password = password if valid_password?(password)
    end

    private #any methods that follow below this keyword will be private, you cannot call this outside of the object
    def generate_production_number
        random_number = rand(10_000..99_999)
        another_random_number = rand(10_000..99_999)
        "2024-#{random_number}-#{another_random_number}"
    end
    def valid_password?(password)
        password.length >= 6
    end
end

phone = SmartPhone.new("rubyfan", "secretpass")
#p phone.generate_production_number

#protected methods -- not as common in the real world, but author says its good knowing that they exist.
#can only be invoked by internal methods or objects of the same class

class Car
    def initialize(age, miles)
        base_value = 20_000
        age_deduction = age * 1000
        miles_deduction = miles / 10
        @value = base_value - age_deduction - miles_deduction
    end

    def compare_car_with(car)
        self.value > car.value ? "Your car is better" : "Your car is worse"
    end

    #a protected method is a method that can be invoked but only by objects of the same class!
    #any car can call the value method on itself or any other car, but other objects can't
    protected
    def value
        @value
    end
end

civic = Car.new(2, 20_000)
fiat = Car.new(10, 90_000)

p civic.compare_car_with(fiat)


#ooo derived code, an anti-pattern!
#the area of a rectalng is a derived value -- it has no value by itself, it is being constantly calculated!

class Rectangle
    attr_accessor :height, :width
    
    def initialize(height, width)
        @height = height
        @width = width
        #@area = height * width #Do not do this. Later you'll change height for instance and area will be wrong
    end

    #instead make it a method -- so that area is calculated every time you want to know it
    #do not store derived state, instead make it a method
    def area
        height * width
    end
end
r = Rectangle.new(3,5)
p r.area

#class methods -- a method that we can invoke on a class itself, as opposed to an object instance of that class!

class Vehicle
    class << self
        def car #instead of having to initalize a car with 4 wheels and 6 passengers every time, you can have the car class method that does this for you
            new(4,6)
        end
        def truck
            new(18,2)
        end
    end

    attr_reader :wheels, :passengers
    def initialize(wheels, passengers)
        @wheels = wheels
        passengers = passengers
    end
    #classname.methodname indicates that this is a method to be called on the class
    #But if you use self in the body if a class, it'll refer to the class itself!
    

    #alternate syntax to declaring class methods
    #since self is already in this first line, you can remove the self word in these

end

mycar =  Vehicle.car
mytruck = Vehicle.truck
puts mycar.wheels, mytruck.wheels

#class variables! data that lives on a class rather than an instance

class Bicycle
  #for class variables, you use @@
  @@count = 0 #count class variable -- counts how many bicycles we madew
  def self.count #count class method
    @@count
  end
  def initialize
    @@count += 1
  end
  def count #count instance method -- a bit of an odd thing to do, but still works
    @@count
  end

end

Bicycle.new
Bicycle.new
Bicycle.new

b1 = Bicycle
p b1.count

#building a class over time -- classes can be defined in parts in more than one file, ay dios mio

class Book
    attr_reader :title, :author, :pages
  def initialize(title, author, pages)
    @title = title
    @author = author
    @pages = pages
  end
end

mybook = Book.new("Atlas Shrugged", "Ayn Rand", 1300)

#BUT WAIT, THERE'S MORE
#you can't call mybook.read here

class Book
    def read
        1.step(pages, 20) {|page| puts "Reading #{page}"}
        puts "Done with #{title}"
    end
end

#but you can here!
mybook.read

#mokey patching -- adding functionality to an existing ruby class. Its fun (according to the author :joy:)
#monkey patching can be dangerous -- you can overwrite an existing method (like the push one in the array)
#monkey patching is advised against!

class String #You can define new instance methods that all strings will have
    def count_vowels
        self.downcase.count("aeiou")
    end
end
hello = "hello"
p hello.count_vowels
p "banana".count_vowels

class Array
    def sorted? #just be careful, you could accidentally define "sort" here and break the existing sort method!
        self == self.sort
    end
end

p [1,2,3].sorted?
p [1,2,4,3].sorted?