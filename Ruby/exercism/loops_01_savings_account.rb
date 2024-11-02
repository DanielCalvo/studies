module SavingsAccount
  def self.interest_rate(balance) #Kind of a weird way of handling percentages from the exercise but okay
    return 0.5 if balance == 0.0 #Phew, apparently exercise wanted interest rate of 0.5 on a 0 balance? Hard to see how that even matters
    return 3.213 if balance < 0.0
    return 0.5 if balance > 0.0 && balance < 1000.0
    return 1.621 if balance >= 1000.0 && balance < 5000.0
    return 2.475 if balance >= 5000.0
  end

  def self.annual_balance_update(balance)
    return balance + balance * self.interest_rate(balance) / 100
  end

  def self.years_before_desired_balance(current_balance, desired_balance)
    counter = 0
    while current_balance < desired_balance
      current_balance = annual_balance_update(current_balance)
      counter += 1
    end
    return counter

  end
end

a = SavingsAccount
puts a.annual_balance_update(200.75)
puts a.years_before_desired_balance(200, 301)
puts a.annual_balance_update(0)

puts a.interest_rate(0.000001)

