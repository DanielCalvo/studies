markdown_content = <<~MD
# Sample Markdown File

This is a sample markdown file with some links.

- Here is a [Google link](https://www.google.com).
- Another link to [GitHub](https://github.com).
- A reference-style link: [Ruby][1].
- A link with special characters: [Example](https://example.com/path?query=value&another=true).
- A broken link: [Broken](not_a_url).

## More Links

You can also have inline links like this: [Stack Overflow](https://stackoverflow.com).

Or even multiple links on the same line: [First](https://first.com) and [Second](https://second.com).

[1]: https://www.ruby-lang.org/en/ "Ruby Programming Language"

A markdownfile can also link to a file, like so: [passwd](/etc/passwd). Links can be absolute or relative (ex: [myfile](/.myfile) is also a valid link)

## You can have a header and a link to it:

And a link to it

- [Associated Repositories](#associated-repositories)

## Associated Repositories

MD


# Lets just work with http links for now, you handle  mail, file and markdown header links later

regex = /\[([^\]]+)\]\(([^)]+)\)/

markdown_content.each_line do |line|
  line.scan(regex).each do |match|
    puts "Text: #{match[0]}, URL: #{match[1]}"
    #try to create uri from link, if it works, its an http link, if it doesnt, its something else
  end
end
