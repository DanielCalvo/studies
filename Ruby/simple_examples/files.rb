
#So easy!
puts File.read("/etc/passwd")

# How about line by line?

f = File.open("/etc/passwd", "r")
f.each_line do |line|
  puts line
  sleep(0.05) #neat!
end
f.close