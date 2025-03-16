
#what attributes did the link have again?

=begin
attributes:
repo
file in github
link text
link address
link type
link status
last checked
=end

=begin
methods:
check link
=end


#Don't foger to documenet the class later
class MarkdownLink
  def initialize(text:, address:, local_file_path:, http_file_path:)
    #Skipping filename as on my go program
    @text = text
    @address = address
    @local_file_path = local_file_path
    @http_file_path = http_file_path
    @type = "UNKNOWN" #Figure this out by calling a private method to get the link type! According to chatgpt calling a method inside initialize is idiomatic and all!
    @status = "UNKNOWN"
  end
end

#maybe you can create a function to get the file name to replace the "filename" class variable you had in go