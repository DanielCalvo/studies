require 'date'
require 'net/http'

#This represents a markdown link with an http(s) address on it
class HTTPLink
  #can i have a reader on something arbitrary, not in the initalize function?
  attr_reader :text, :address, :http_status, :date_checked
  #lets add everything on initialize for clarity for now
  def initialize(text:, address:)
    @text = text
    @address = address
    @http_status = nil
    @date_checked = nil
  end
  def check_link
    uri = URI(@address)
    hostname = uri.hostname # => "example.com"
    #hmm but wait, what if I want to get something like github.com/torvalds?
    req = Net::HTTP::Head.new(uri) # => #<Net::HTTP::Head HEAD>
    res = Net::HTTP.start(hostname) do |http|
      http.request(req)
    end

    @http_status = res.code
    @date_checked = Date.today
  end

end

li = HTTPLink.new(text: "google.com", address: "https://google.com")

puts li.check_link
puts li.date_checked
puts li.http_status


li = HTTPLink.new(text: "google.com", address: "http://google.com")

li2 = HTTPLink.new(text: "torvalds", address: "github.com/torvalds") #fails, says not a uri, guess we gotta handle that too!
li2.check_link
puts li.http_status