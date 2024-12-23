# The customers.txt file includes a list of all customers who
# visited out shop in the last month along with their phone numbers.
# Each customer name and phone number is separated by a comma.
# I'd like to call each phone number once to ask about them their experience.
# Some customers visited the store multiple times, so I want to make 
# sure there are no duplicate phone numbers in the final collection.
#
# Define a generate_unique_phone_numbers method.
# The method should parse the attached customers.txt file
# Return a Set with the phone numbers of the customers.
# Ignore the customer's names.
#
# Parse the file and return a Set of all unique phone numbers.

def generate_unique_phone_numbers
  customers = File.open("customers.txt")
  result = Set.new
  customers.each do |customer|
      phone_number = customer.split(",")[1]
      result.add(phone_number.strip)
  end
  return result
end

puts generate_unique_phone_numbers

