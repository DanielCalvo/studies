#begin, rescue and what else?
#lets wget something that doesnt exist and then handle this faliure


begin
  file = File.open("doesntexist.txt", "r")
rescue # A rescue clause without an explicit Exception class will rescue all StandardErrors (and only those).
  puts "fail" #grabs everything
ensure
  file.close if file
end