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
invalid_uri = URI("wee e e")
p my_uri


#I wonder if you can use the debugger to check how an url is parsed?