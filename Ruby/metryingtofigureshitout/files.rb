
#So easy!
puts File.read("/etc/passwd")

# How about line by line?

f = File.open("/etc/passwd", "r")
f.each_line do |line|
  puts line
  #sleep(0.05) #neat!
end
f.close

puts File.exist?("/etc/passwd")
puts File.directory?("/etc/passwd")

#I wonder if file or fileutils allow us to split by path and filename. Yup!
#https://ruby-doc.org/3.4.1/File.html#method-c-basename
#https://ruby-doc.org/3.4.1/File.html#method-c-dirname

file = "/home/daniel/Downloads/finding_nemo.mp4"
puts File::dirname(file)
puts File::basename(file)
puts File::SEPARATOR