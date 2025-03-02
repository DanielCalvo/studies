# so here's a list of strings:

my_strings = [
  'google.com',
  'banana',
  'https://google.com',
  'mailto:someone@example.com',
  'https://github.com/torvalds',
  'github.com/torvalds',
  '#some-title',
  'http://example.com/invalid url',
  '\\\\',
  '🤩'
]

my_strings.each do |my_string|
  if my_string.include?('@')
    puts "probably an email: #{my_string}"
  elsif my_string.start_with?('http')
    puts "probably a website: #{my_string}"
  elsif my_string.start_with?('#')
    puts "probably a link to a heading in this file: #{my_string}"
  elsif my_string.start_with?('/')
    puts "probably a link to a file: #{my_string}"
  else
    puts "probably an URL, nothing else matched: #{my_string}"
  end
end