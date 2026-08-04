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

