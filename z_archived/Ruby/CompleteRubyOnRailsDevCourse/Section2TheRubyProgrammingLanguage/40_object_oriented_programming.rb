#First lstter has to be capitalized
class Student
  attr_accessor :first_name, :second_name, :age #does a getter and setter for you

  def initialize(first_name, second_name, age)
    @first_name = first_name
    @second_name = second_name
    @age = age
  end

  def to_s
    return "#{@first_name}, #{@second_name}, #{@age}" #apparently you can't return 2 variables? if I remove the quotes and the templating it doesn't work
  end
end

joe = Student.new("joe","mcjoeson", 18)
puts joe