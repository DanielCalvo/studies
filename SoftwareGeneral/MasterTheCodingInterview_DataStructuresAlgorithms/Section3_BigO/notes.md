
### 22. Setting up your Environment
- Other than your local environment, you can also use:
    - repl.it
    - glot.io

### 23. Section overview
- Big O is very popular in interviews!

### 25. What is good code?
- Good code can be described as:
    1. Readable
    2. Scalable (Big O notation allows us to measure code that is scalable)

### 26. Big O and Scalability
- Big O notation is the language we use for how long it takes for an algorithm to run
- Big O: When we grow bigger and bigger with our input, how much does the code (or function) slow down?
- The Big O complexity chart is a thing, it shows algorithmic efficiency

### 27. O(n)
- The find element in array function as described in the example has a Big O notation of O(n). It's linear. It takes linear time to find the element
- O(n) - Linear time. One of the most common

### 28. O(1)
-  O(1): Constant time
- Example: Function to print the name of the first element on an array
- No matter how much the input of your function increases, in terms of scalability, it doesn't matter how large the input is, we always do the same number of operations

### 29. Exercise: Big O Calculation
- I believe this function has a Big O of O(n), as the operations in the function increase linearly with input
- Big O of the function was actually Big O(3 + 4n). Interesting!

### 31. Exercise: Big O Calculation 2
- This looks linear too. I'm gonna go with O(n)

### 33. Simplifying Big O
- You don't need to go through every line of the function to calculate Big O, there are rules you can apply to simplify Big O

### 34. Big O Rule 1
- Rule 1: Worst case
- When calculating Big O, always think about the worse case

### 35. Big O Rule 2
- Rule 2: Remove constants
- As far as Big O is concerned, you're only concerned about the input
- With Big O you don't care how steep the line is on the graph, you care about how the line moves as the input increases

### 36. Big O Rule 3
- Rule 3: Different terms for inputs
- If your function takes, say 2 lists, and you loop over the two lists, your big O is O(a + b)
- If two loops are after the other: O(a + b). If two loops are nested: O(a * b)

### 37. O(n^2)
- When you see nested loops, you use multiplication for the Big O.
- Sample problem: `Log all pairs of arrays`
- Big O for this os O(n^2). Quadratic time!

### 38. Big O Rule 4
- Drop the non dominant terms
- If you function has a Big O of O(n + n^2), that means you care about the most important term! So you have: O(n + n^2)
- As the input increases, n^2 will increase a lot more than just n 

### 39. Big O Cheat sheet
#### Big Os
- O(1) Constant: no loops
- O(log N) Logarithmic: usually searching algorithms have log n if they are sorted (Binary Search)
- O(n) Linear: for loops, while loops through n items
- O(n log(n)) Log Linear: usually sorting operations
- O(n^2) Quadratic: every element in a collection needs to be compared to ever other element. Two nested loops
- O(2^n) Exponential: recursive algorithms that solves a problem of size N
- O(n!) Factorial: you are adding a loop for every element
- Iterating through half a collection is still O(n)
- Two separate collections: O(a * b)

##### What can cause time in a function?
- Operations: (+, -, *, /)
- Comparisons: (<, >, ==)
- Looping: (for, while)
- Outside: Function call (function())
#### Rule Book
- Rule 1: Always worst Case
- Rule 2: Remove Constants
- Rule 3: Different inputs should have different variables. O(a+b). A and B arrays nested would be:
    - O(a*b)
    - + for steps in order
    - * for nested steps
- Rule 4: Drop Non-dominant terms

#### What causes Space complexity?
- Variables
- Data Structures
- Function Call
- Allocations

### 40. What does this all mean?
- Noice: https://www.bigocheatsheet.com/
- Data structures + Algorithms = Programs!
- A good programmer picks the right data structure and the right algorithm for his program.
- Remember, good code: Readable and scalable

### 41. O(n!)
- :scared_face:
- Very expensive! Factorial time. Also called by the author "oh no" :joy: 
- With this notation, we're adding a nested loop for every element that we have

### 42. Pillars of programming
1. Readable
2. Scalable
    - Speed (how fast is it? how many operations does it cost? how long does it take to run?)
    - Memory (computers have limited memory)
- Which code is best?
    - Readable (clean code that others can read and is maintainable)
    - Speed (code that has a big O time complexity that is efficient, it scales well)
    - Memory (Low space complexity. We use the same big O notation for memory usage)
- There is usually a tradeoff between speed and memory
    - There is usually a tradeoff between saving time and saving space

### 43. Space Complexity
- The amount of memory used by a program or function. Big O is used to represent space complexity too

### 44. Exercise: Space Complexity
- Populating an array with every item on a list is of space complexity O(n)

### 45. Exercise: Twitter
- Calculating the length of a string has which complexity?
    - Depends on the language and how it works