
#Apparently you can inherit from one of the parsers and define your on function: https://github.com/vmg/redcarpet?tab=readme-ov-file#and-you-can-even-cook-your-own
#Maybe you can even print the links in csv format or something!

require 'redcarpet'


markdown = Redcarpet::Markdown.new(Redcarpet::Render::HTML, autolink: true, tables: true)
puts markdown.render("This is *bongos*, indeed.")


