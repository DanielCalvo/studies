
## 2.1 names
This is basic but I suppose it doesn't hurt writing it down:

### Entity scope
- If an entity is declared within a function it is local to that function
    - If declared outside a function it is visible in all files of the package 

The case of a first letter of a name determines its visibility:
- If a name begins with an uppercase letter, it is exported, meaning it is accessible outside of its own package
- If it starts with lower case, it's only visible inside the package

Go programs lean towards short names. (ex: i instead of loopIndex)
Camel case is preferred for names, ex: convertToString 


## 2.2. declarations
Here's the basic go program structure
1. package declaration
2. imports
3. package level declarations of types variables constants and functions. The book says in any order but have usually seen types then constants/variabels then functions


## 2.3. Variables
this one is hard to forget but it doesn't hurt to write it down, but you can declare a variable like this and it's the most expressive form:
- `var name type = expression`

the type or the expression can be omitted, but not both. ex:
- `var s string`
- `var i = 2`

If the type is omitted it is determined by the initializer expression
If the expression is omitted the initial value is the zero value for that type
this is zero for numbers, false for booleans, "" for strings, and nil for interfaces in reference types (like slices pointers maps channels and functions)

you can also declare multiple variables in one go:
- `var i, j, k int`

you can even declare variables of different types together:
- `var b, f, s = true, 2.3, "four` <- not sure I like this though, it's not particularly clear

oh wow I didn't even know you could do this:
- `var f, err = os.Open(name)`

### 2.3.1. Short Variable Declarations
within a function you can use short variable declarations, which you're very familiar with, it uses `name := expression`:
- `freq := rand.Float64() * 3.0`
- `i, j := 0, 1`

oh I didn't know this: `:= is a declaration, whereas = is an assignment.`
(I mean maybe I knew but not with the correct terms)

another thing you already knew: a short variable declaration does not declare all the variables on the left side if some of them were already declared
if some variable was already declared, then the short variable declaration acts like an assignment:

```go
in, err := os.Open(infile)
// ...
out, err := os.Create(outfile)
```


Are some other things that are covered, but I am somewhat familiar with them so I didn't go into much depth:
- 2.3.2. Pointers
- 2.3.3. The new Function
- 2.3.4. Lifetime of Variables
- 2.4. Assignments
- 2.4.1. Tuple Assignment
- 2.4.2. Assignability
- 2.5. Type Declarations
- 2.6. Packages and Files
- 2.6.1. Imports
- 2.6.2. Package Initialization (there is an interesting exercise you could do here to initialize stuff in a package as you import the package,I think you saw this in Prometheus)
- 2.7. Scope


## assorted notes
To make your program work with the variable defined in the other file: go run main.go otherfile.go 

this is kind of a not question but is it a way is there a way to make the initialized value of something completely different? can I have a custom type of underlying type integer and have the initialized value be like forty two?
