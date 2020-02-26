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
- Small files are usually easier to understand than large files

You stopped on page 108
