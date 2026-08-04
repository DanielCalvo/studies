class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze
  class UnsupportedOperation < StandardError #Oh this is what the test wanted, took me forever to figure out wew
  end
  def self.calculate(num1, num2, operator)
    raise ArgumentError unless num1.is_a?(Numeric) && num2.is_a?(Numeric)

    # raise UnsupportedOperation if operator != '+' && operator != '*' && operator != '/'

    # Oh wait it is much better to create an constant array of allowed operators and check if what was passed is there!
    raise UnsupportedOperation if ALLOWED_OPERATIONS.include?(operator) #raise exception is the operator is not in the array of allowed operators

    case operator
    when '+'
      "#{num1} + #{num2} = #{num1 + num2}"
    when '*'
      "#{num1} * #{num2} = #{num1 * num2}"
    when '/'
      return "Division by zero is not allowed." if num2 == 0
      "#{num1} / #{num2} = #{num1 / num2}"
    end
  end
end


a = 1
b = 2
c = '+'

calc = SimpleCalculator
puts calc.calculate(a, b, c)
 puts calc.calculate(3, 3, '*')


puts calc.calculate(a, b, '**')

# puts calc.calculate(4, 0, '/')
