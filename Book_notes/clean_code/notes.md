### Chapter 2: Meaningful names
- Use intention revealing names. Think about how to name things properly
- Avoid disinformation. Don't call something accountList if it's not an actual list
- Careful with 1 and l and O and 0
- Avoid things named "a1" and "a2"
- If you have a product class, having variables named "ProductInfo" and "ProductData" are noise
- Use pronounceable names!
- Use searchable names. Single letter variables are difficult to search. Numerical constants too!
- Dont be afraid of long variable names if they make sense
- Class should be nouns, without verbs
- Functions should have verbs!
- Don't be cute. Don't put cultural jokes in the code
- Pick one word per concept and stick with it. Fetch, retrieve, get, etc: Pick one and stick to it

### Chapter 3: Functions
- The first rule of functions is that they should be small!
- The second rule is that they should be smaller than that!
- Functions should do one thing and do it well!
- Try to keep functions within one level of abstraction
- "You know you are working on clean code when each routine turns out to be pretty much what you expected."
- Careful with switch/cases, you might end up having to copy/paste them all over the place
- Have as less arguments as possible? :o
- Booleans as arguments are red flags. Does one thing if true, other thing if false?
- When a function requires too many arguments, maybe wrap them in a class/type
- Have no side effects! Don't modify global things on functions
- Output arguments should be avoided. If your function must change something, have it change the state of it's owning object
- Try not to repeat yourself
- Functions should have one entry and one exit (oof)
- In a small function, a return, break or continue statement can expressive
- When wrting a function, it might be ugly when first written. Look at it and improve it until it is good!

### Chapter 4: Comments
- The proper use of comments is to compensate for our failure to express ourself in code. Try writing clearer code first
- Any comment that forces you to look into another module to understand it, is a bad comment
- Don't use a comment when you can name a function or variable properly
- Careful commenting code out (not recommended)

### Chapter 5: Formatting
- Agree on formatting rules and follow them. `go fmt` solves this, ha!
- Small files are usually easier to understand than large files
- You want source files to read like a newspaper. Summarized up top, richer in details as you progress (more abstraction to less abstraction)
- Jumping lines between functions and concepts is encouraged as it helps with clarity
- Lines of code that are tightly related should appear vertically dense
- On the same file, concepts that closely relate to each other should be kept vertically close
- Separating closely related concept across (many) different files is not encouraged
- Declare variables as close to their usage as possible, but in small functions it's ok to declare them at the top of the function
- Dependent functions should be close, and the caller should be above the callee if possible
- The more affinity to parts of code have for each other, the closer they should be
- Function dependencies are encouraged to point in the downward direction. A function that is called should be below the one that does the calling
- Beware of lines too long. Anything longer than 120 characters might be too long
- Whitespaces can be used to indicate preference (as in math, for instance)
- Identation: Having things on multiple lines can be easier to understand than having many clever things on one line
- Guess I'll stick to using my `fmt` commands, ha!

### Chapter 6: Objects and data structures
- Author appears to favour abstractions (on a Java interface getPercentFuelRemaining() is favourable compared to getFuelTankCapacityInGallons() and getGallonsOfGasoline())
- Objects: Hide their data behind abstractions and expose functions that operate on that data
- Data structures: Expose their data and have no meaningful functions

Stopped on page 126. This chapter is heavy on OO theory.

### Chapter 7: Error handling
- Error handling is important, but if it obscures code, it's wrong
- Books talks about try-catch-finally, but I don't use those in golang or shell (if you do python again, re-read it)
- Try to write tests that force exceptions, and add that to your code
- Each exception that you throw should provide enough information to determine the source and location of an error
- If you're going to be handling the same error multiple times, you can write a wrapper for it
- Wrapping a third party API is considered best practice (?)
- You can have your object have a special case pattern so you don't have to put extra logic inside your try-catch. Interesting
- Don't return null! Or at least in the context of an error in OO languages. Consider returning a special case object instead
- Don't pass null into methods either

### Chapter 8: Boundaries
- Seems to focus on generics and interfaces a bit too much
- It may be on your best interest to write code for the third party code you use
- It is encouraged to write tests to explore the behaviour of third party code (when first using/discovering/exploring it)
- Exploratory tests cost nothing: You had to learn the API/external code anyway, and that was a way to learn their isolated behaviour. And now you have tests!
- When coding against an api or external component that you don't have yet, you can design it as you _would behave_ it would be and develop against that

### Chapter 9: Unit tests
- The three laws of TDD (which I don't understand well)
    1. You may not write production code until you have written a failing unit test.
    2. You may not write more of a unit test than is sufficient to fail, and not compiling is failing
    3. You may not write more production code than is sufficient to pass the currently failing test
- Having dirty and half-assed tests is the same as having no tests
- Tests must change as the production code changes. Tests are just as important as production code
- Tests keep you code flexible, as with tests you don't have to be afraid to change the code
- It is important for your tests to be readable!
- You can use the BUILD-OPERATE-CHECK pattern to write tests
- F.I.R.S.T: Clean tests should follow these five rules:
    - Fast: Tests should be fast, they should run quickly
    - Independent: Tests should not depend on each other
    - Repeatable: Tests should be repeatable in any environment
    - Self validating: Tests should have boolean output
    - Timely: Tests need to be written at the proper time, preferably just before the production code to make them pass

### Chapter 10: Classes
- In Java, a class should begin with a list of variables
- Utility functions and variables are encouraged to be private
- You can allow tests inside the package to reach private variables if it helps test development
- Classes should be small! (just like functions)
- The name of a class should describe the responsability if fulfills
- Classes can have too much responsability! (a bad thing)
- Certain words like super, manager or processor can hit at too much responsability
- Classes should have one responsability only
- Author argue that smaller classes > larger classes, as it's easier to organize smaller clases, and complex systems will be complex anyway
- DIP (Dependency Inversion Principle) says that our classes should depend upon abstractions, not on concrete details (?)

### Chapter 11: Systems
- One way to separate construction from use is to move all aspects of construction to main, or modules called by main (?) and to use the rest of the system assuming those objects have been constructed appropriately
- A lot about is talked about abstract factories in here. Kinda a bit over my level for what I usually do
- Implement today's stories, then refactor and expand the system to implement new stories tomorrow

You stopped on page 188. Dependency injection seems interesting