=begin
#How does the uri module handle an invalid host exception?
def URI(uri) calls
URI.parse(uri)
This in turn calls DEFAULT_PARSER.parse(uri)
DEFAULT_PARSER is an instance of RFC3986_PARSER, defined in rfc3986_parser.rb

RFC3986_PARSER.parse(uri) calls:
URI.for(*self.split(uri), self)
the split method then has the bulk of the logic, it checks if the url can be converted to string
then checks if its ascii only, and then checks if matches the RFC3986_URI regex
if that doesnt work, tries to match RFC3986_relative_ref
and if that doesnt work either you get a InvalidURIError, ha!
=end

#The tests for uri validating what an invalid URL is are here: 

require "uri"

#handle an invalid host exception when going through a list of hosts, and continue!
begin
  invalid_uri = URI("wee e e")
rescue URI::InvalidURIError
  puts "Invalid URI"
end

#I wonder if you can use the debugger to check how an url is parsed?

#can I have a uri with something like github.com/torvalds?
myuri = URI("github.com/torvalds")
puts myuri #apparently, it does instantiate yeah
puts myuri.host #but apparently it does not get populated
puts myuri.path
puts myuri.port

myuri2 = URI("https://github.com/torvalds")
puts myuri2
puts myuri2.host #this does get me all the fields
puts myuri2.path
puts myuri2.port

#so apparently you need to handle adding the http(s) suffix yourself to an url if you don't have it

#anyways lets check for what is a valid uri and what isnt:

urls = [
"google.com",
"banana",
"https://google.com",
"mailto:someone@example.com",
"https://github.com/torvalds",
"github.com/torvalds",
"http://example.com/invalid url",
"\\\\",
"🤩",
]

puts "starting url check!"
urls.each do |url|
  begin
    URI.parse(url)
    puts "Valid URL: #{url}"
  rescue => e
    puts "Invalid URI: #{url}, error: #{e.class}"
  end
end