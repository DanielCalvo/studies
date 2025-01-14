
city_codes = {
  "Madrid" => "28001",
  "Valencia" => "46001",
  "Barcelona" => "08001"
}


loop do 
  puts "Do you want to lookup an area code based on a city name? (y/n)"
  input = gets.chomp.downcase #downcase string method, nice!
  if input != "y"
    break
  end

  puts "Which city do you want the area code for?"
  puts city_codes.keys
  input = gets.chomp

  if city_codes[input]
    puts "The city code for #{input} is #{city_codes[input]}"
  else
    puts "City #{input} was not found"
  end

end