#https://ruby-doc.org/3.4.1/Dir.html

dir = Dir.new("/home/daniel")

mdFiles = Dir.glob("/home/daniel/Projects/**/*.md") #lists all markdown files in the Projects directory recursively

#the above returns an array, so you can iterate over it like any other array 
mdFiles.each do |file|
  puts file
end