# https://github.com/ruby-git/ruby-git
require "git"

git = nil

#what exception do we put in here when you cant clone the repo?
#how about failed filesystem operations? (ex: no write permissions)
#or how about a network timeout? or wrong repo url?
if ! File.directory?("/tmp/boulder")
  puts "Clonning boulder"
  git = Git.clone("https://github.com/letsencrypt/boulder.git", "/tmp/boulder")
else
  puts "Boulder already cloned! Continuing"
  #Open the git repo
  git = Git.open("/tmp/boulder")
end

puts "Programming is fun!"
puts git.remote.url


