
passwd = File.open("sample_file.txt")
puts passwd.class

passwd.each {|line| puts line }

#writing text to a file

#by default ruby treats opening files as a ready only operation, by default ruby prohibits itself from opening files
#w says lets grant write priviledges to this file!
passwd2 = File.open("sample_file2.txt", "w") do |file|
    file.puts "I am creating this text file using ruby!"
    file.write "The line method does not add a line break at the end!"
end

#One thing to know: The "w" function will completely replace the contents of the file! (aka it'll truncate it)

passwd2 = File.open("sample_file2.txt", "a") do |file|
    file.puts "This apppends to the file with the \"a\" flag"
end

#rename and delete a file!

File.rename("sample_file2.txt", "sample_file3.txt")
File.delete("sample_file3.txt")

if File.exist?("banana.txt")
    puts "Deleting the banana file"
    File.delete("sample_file3.txt")
else
    puts "banana.txt does not exist!"
end

#command line arguments!