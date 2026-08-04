A markdown link is a link found in a markdown file. These links can be of various types: Link to a file, http address, email, or a link to a section of the document, among others.

In this example, lets have a parent class for a markdown link, and then subclasses for a markdown link to a file and another one to a http address. Email and document section I'll leave unhandled for now.

# Alrite but what attributes does a link have?

## On a markdown file, a link has
- Link name
- Link destination 

## But you also want to know
- Link type
- Link status (ex: http 200 or 404, file exists or is not found)
- Link last checked time

Maybe you can have more attributes later but lets work with these for now

## Ok but wait
- I'm too much of a dummy, let me keep it simple and put everything on an http link and i'll make it pretty aftewards.