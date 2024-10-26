#Ruby uses snake cases for variables. It is not as strongly typed as go:

my_first_variable = 1
my_first_variable = "Some string"

#Constants however can only be assigned once. They start with a capital letter
MY_CONST = 10
Aconst = 11

#Ruby is organized into classes. Classes are defined by the class keyword and generaly created by instantiating them with the .new method
class Calculator
  def add (num1, num2)
    return num1 + num2
  end
  def multiply (num1, num2)
    return num1 * num2
  end

end

mycalc = Calculator.new
mycalc.add(1, 5)

#Functionality is encapsulated in methods -- like functions in other languages


#Exercise. const suggestion for the layer time was pretty cool!
class Lasagna
  EXPECTED_MINUTES_IN_OVEN=40
  LAYER_TIME=2
  def remaining_minutes_in_oven(actual_minutes_in_oven)
    return EXPECTED_MINUTES_IN_OVEN - actual_minutes_in_oven
  end

  def preparation_time_in_minutes(layers)
    return layers * LAYER_TIME
  end

  def total_time_in_minutes(number_of_layers:, actual_minutes_in_oven:)
    return number_of_layers * LAYER_TIME + actual_minutes_in_oven
  end
end