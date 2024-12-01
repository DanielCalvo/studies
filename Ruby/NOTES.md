## The basics
- What operators does ruby have again? There are some I don't recognize
- How do I run tests again? exercism doesn't cover that, boo!
- AAAAA DO HASHES AGAIN
- AAAAA HOW2READ AND WRITE FROM FILES?
- How do I create my own gem?
- How do you modules again?
    - How do you do that hting you discovered in go once that when you import a module, it runs some code? 
- Create an example throwing a typeerror
- As an amusing thought experiment: Can you catch a syntax error? That would be fun, ha!
- Write a function that throws an argument error if you pass the wrong args!
- Google: What are the most common errors in ruby?
- Explore all the integer and string methods, there are so many cool ones!
- What can rubocop do for you?
- Any "awesome ruby" repos out there?
- Oh woah, write your own block method for your own object sometime
- Try out some recursion on your own sometime, you lack practice!
- Does the inclusion operator work on an array or a hash? what about a string? test it out!
- Save all your exercises from the course in a file somewhere!
## The docs
- https://ruby-doc.org/3.3.5/
- https://ruby-doc.org/3.3.5/File.html
- https://ruby-doc.org/3.3.5/stdlibs/yaml/YAML.html

## woah
What is going on here? From integer.rbs
```ruby
  def step: () { (Integer) -> void } -> void
          | (Numeric limit, ?Integer step) { (Integer) -> void } -> void
          | (Numeric limit, ?Numeric step) { (Numeric) -> void } -> void
          | (to: Numeric, ?by: Integer) { (Integer) -> void } -> void
          | (by: Numeric, ?to: Numeric) { (Numeric) -> void } -> void
          | () -> Enumerator[Integer, bot]
          | (Numeric limit, ?Integer step) -> Enumerator[Integer, void]
          | (Numeric limit, ?Numeric step) -> Enumerator[Numeric, void]
          | (to: Numeric, ?by: Integer) -> Enumerator[Integer, void]
          | (by: Numeric, ?to: Numeric) -> Enumerator[Numeric, void]
```