#allow invocation of a method to specify which parameters the arguments correspond to, uh-oh

class Person
    attr_reader :name, :age, :gender, :occupation
    def initialize(details)
      @name = details[:name]
      @age = details[:age]
      @gender = details[:gender]
      @occupation = [:detailsoccupation]
    end
  
    #def introduce
    #  "Hi, I'm #{@name}. I'm a #{@age}-year-old #{@gender} working as a #{@occupation}."
    #end
  end

  #per = Person.new("joe", 18, "male", "happy dude")
  #puts per.introduce

  #uh-oh, we can pass a hash as a single argument to the function

  #Hmm, a clearer approach. Also, argument position doesn't matter as much, interesting!
  #However, if you forget or mistype a hash field, you will have nil values introduced!
  #When a hash is the last argument to a method, you can remove the curly braces! :o
  per = Person.new(
    name: "joe",
    age: 18,
    gender: "male",
    occupation: "happy dude"
  )

  puts per.name

  #required keyword arguments -- allows the invocatin of a method to specify which parameters the arguments correspond to 

  def sum(a:, b:) #will make ruby look for keyword arguments as opposed to sequential arguments!
    a+b
  end

  p sum(a: 2, b: 3)
  p sum(b: 3, a: 3) #order doesnt matter when you have names!
  #p sum(b: 3) #argument error
  #p sum(b: 3, c: 3) #argument error
  #p sum(a: 3, b: 3, c: 3) #argument error: unknown keyword: :c (ArgumentError)

  #optional keyword arguments!
  def sum2(a:1, b:1) #the invocation can provide a, but if it doesnt, lets fall back to a value of 1
    a+b
  end

  p sum2(b:2)
  p sum2(a:2)
  p sum2(a:2,b:2)

  #positional arguments and keyword arguments!
  #the traditional way is to put your positional arguments first, followed by your keyword arguments!

  def sum3(a, b:1)
    a+b
  end

  p sum3(3, b:5)
  p sum3(4)
  #p sum3(4,9) #errors out -- label of b is needed!
  #p sum33(b: 4, 9)


  #Lets define the person class again with keyword arguments!
  class Person
    attr_reader :name, :age, :gender, :occupation
    def initialize(
      name:,
      age:,
      gender:,
      occupation: "Happy person" #Occupation has a default value!
      )
      @name = name
      @age = age
      @gender = gender
      @occupation = occupation
    end
  
    def introduce
      "Hi, I'm #{@name}. I'm a #{@age}-year-old #{@gender} working as a #{@occupation}."
    end
  end

  per = Person.new(name: "Joe", age: 18, gender: "Male")
  puts per.introduce