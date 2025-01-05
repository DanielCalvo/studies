# https://github.com/ruby-git/ruby-git
require "git"
require 'uri'
require 'net/http'

#what exception do we put in here when you cant clone the repo?
#how about failed filesystem operations? (ex: no write permissions)

def clone_repo
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
  git
end


#Do you get directories? How about symlinks?
#parametrize with a named parameter
#
#Finds all files that end in .md recursively in the passed directory
def findMarkdownFiles(directory)
  Dir.glob(directory + "/**/*.md")
end

#p findMarkdownFiles("/tmp/boulder")

#clone git
#find md files
#find markdown links
#check them
#generate html with result
#push it s3
#create dockerfile
#put in k8s
#write tests!!!1!1one

#uri could be invalid
def checkLink
  uri = URI('https://www.twilio.com/en-us/blog/5-ways-make-http-requests-ruby')
  #then just do a head
end

checkLink

#I wonder if there's a gem to find markdown links in files in ruby?