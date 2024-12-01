

i = 1

while i <= 30

    case
    when i % 3 == 0 && i % 5 == 0
      puts "FizzBuzz"
    when i % 5 == 0
      puts "Buzz"
    when i % 3 == 0
      puts "Fizz"
    else 
      puts i
    end

  i += 1
end
