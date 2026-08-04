From having a look at: https://github.com/ruby/ruby/blob/master/lib/uri/

- Each file seems to be home to a class (http, https, file, etc)
- All classes seem to be inside the URI module
- Most files have a require_relative to the generic file
- Most classes inherit from the generic class
