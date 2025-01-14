#Summary of Keywords used when handling exceptions, brought by our friend gepeto:
#Core: begin, rescue, ensure, else, raise.
#Advanced: retry, catch, throw, fail

#https://ruby-doc.org/3.4.1/Exception.html

#lets get something that doesnt exist and then handle this faliure
require 'tmpdir'

begin
  file = File.open("doesntexist.txt", "r")
rescue # A rescue clause without an explicit Exception class will rescue all StandardErrors (and only those).
  puts "fail" #grabs everything, not good
ensure
  file.close if file
end

#lets do the above again but rescue the specific doest not exist error:
begin
  file = File.open("doesntexist.txt", "r")
rescue Errno::ENOENT => e
  puts e.message
ensure
  file.close if file
end

#now lets try to create a file in a place where we lack permissions (/root) and after that create that file in the /tmp folder (where we do have permissions!)
#but wait, what if the file already exists?
begin
  file = File.open("/root/singlefile.txt", "w")
rescue Errno::EACCES => e
  file = File.open("/tmp/singlefile.txt", "w")
ensure
  file.close if file
end

#lets try something 3 times
attemps = 0
begin
  attemps += 1
  puts "Doing thing! Attempt number #{attemps}"
  raise "Thing failed!"
rescue
  retry if attemps < 3
end

#Ok so lets try the above but with a twist:
#We still have our file list. We try to create the files.
#If we fail to create a given file for some issue, like /root/123 for insufficient permissions, lets create this file under tmp (in this case, /tmp/123)
#Phase two: Try to do it onlye one more time! Who knows, maybe the disk is full... 
#eh its looping too many times for some reason
#
#no wait: only nhandle permission denied errors, for the rest have an "unhandled error" and just skip the file! (you don't know what other errors you can get, like an empty disk space error!)

#Will this enter an infinite loop if you don't have permissions to write to /tmp?


file_list = ["/etc/passwd", "/usr/aaa", "/tmp/aaaa", "/tmp/banana", "/tmp/😂", "/root/123"]
file_list.each do |file|
  begin
    unless File.exist?(file)
      File.open(file, "w") {}
      puts "File #{file} created successfully!"
    end
  rescue Errno::EACCES => e
    temp_file = File.join(Dir.tmpdir, File.basename(file))
    unless File.exist?(temp_file)
      File.open(temp_file, "w") {}
      puts "File #{file} created in temporary location: #{temp_file}"
    end
  end
end

#wait, can we open /etc/passwd for writing, even if we cant' write to it?
begin
  File.open("/etc/passwd", "w") {}
  puts "/etc/passwd opened successfully!"
rescue Errno::EACCES => e
  puts "trying to open /etc/passwd threw #{e.class}" #nope, throws exception. neat!
end
