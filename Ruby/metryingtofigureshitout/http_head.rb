
#https://ruby-doc.org/3.4.1/stdlibs/net/Net/HTTP.html
require 'net/http'

uri = URI('http://example.com')
hostname = uri.hostname # => "example.com"
req = Net::HTTP::Head.new(uri) # => #<Net::HTTP::Head HEAD>
res = Net::HTTP.start(hostname) do |http|
  http.request(req)
end

puts res.code

#how do I handle this though?
some_url = "github.com/torvalds"

#welp guess i'll append https if it doesnt have any preffix to it
if !some_url.start_with?("http") || !some_url.start_with?("https")
  some_url_with_https = "https://" + some_url
end
puts some_url_with_https


#If you will make only a few requests of all kinds, consider using the various singleton convenience methods in this class.
#Each of the following methods automatically starts and finishes a session that sends a single request:
# https://github.com/

#There is a get top level function, but not a head :(
uri = URI("https://github.com/torvalds/linux")
res = Net::HTTP.get(uri)
puts res.class #returns a string, so can't do much here to get the http.status

#Ah so this is how it also can go
hostname = 'github.com'
Net::HTTP.start(hostname) do |http| #http is a Net::HTTP class
  puts http.head('/torvalds/linux').code #aaaahm!
  puts http.head('/torvalds/linux').class #Net::HTTPMovedPermanently, interesting
end

# More notes of me reading the docs:
require 'net/http'

uri = URI("https://github.com/torvalds")

http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = (uri.scheme == "https")  # Enable SSL for HTTPS

request = Net::HTTP::Head.new(uri)  # Create a HEAD request
response = http.request(request)
puts response.code #huh, so this gives you a 200!

# -----
http = Net::HTTP.new("github.com")
other_res = http.head('/torvalds') # => #But you get a 301 if you do it like this though!
puts other_res.code  #huh, so this gives you a 301!

