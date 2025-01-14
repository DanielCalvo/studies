
#This is from the course. Oof, wouldn't it have been better to just have a hash, and not hashes inside an array?
users = [
  {username: "joe", password: "a"},
  {username: "bob", password: "b"},
  {username: "alice", password: "c"} #it does work with a stray last coma in here or not, interesting
]

puts "Welcome to the authenticator"
25.times {print "-"}
puts

counter = 0

#You could also have done something like: while couter <= 3
loop do
  print "Username: "
  user = gets.chomp
  print "Password: "
  pass = gets.chomp

  #nil is returned if an index is not found. In Ruby, nil is false.
  index = users.index({username: user, password: pass})

  if index
    puts users[index]
  else 
    puts "Credentials were not correct"
    counter+=1
  end

  if counter >= 3
    puts "You have exceeded the number of attempts"
    exit
  end

  puts "Press n to quit or any other key to continue"
  key = gets.chomp
  if key == "n"
    puts "Exiting the program"
    exit
  end

end