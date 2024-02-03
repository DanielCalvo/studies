## 4. Installing & Using TypeScript
Wrote an example and compiled it to Javascript. Neat!

## 5. TypeScript Advantages - Overview
- TS adds types! IDEs can also help!
- Also next gen JS can be compiled for older browsers, neat
- Adds features, like interfaces and generics
- Metaprogramming features, like decorators!
- Adds... rich configuration options?
- Tooling can help... in non typescript projects?

## 6. Course outline
- Basics
- Compiler & config deep dive
- Next gen JS code!
- Classes & Interfaces
- Advances types & features
- Generics 
- Decorators
- Full project (neat!)
- Namespaces & modules!
- Webpack & Typescript
- Third parti libraries & typescript
- React + typescript + nodejs + typescript!

## 7. How to get most of the course
- Watch the videos
- Code along
- Practice!
- Debug & search

## 8. Setting Up A Code Editor / IDE
VScode with a bunch of plugins (copy later if you want)


## 9. The Course Project Setup
- See index.html -- this allows you to see code changes on the screen in the browser!
```shell
npm init
npm install --save-dev lite-server
npm start
```
Neat!

## 11. Using Types
Core types
- number (no int or floats, just numbers)
- string (text! with 2 quote types and backticks for template literals)
- boolean

Remember, typescript's type support only helps you during compilation, the browser has no built-in type support
Typescript does not change your run time code

## 12. TypeScript Types vs JavaScript Types
- typeof
- Javascript is dynamically typed, you might need to check for types
- Typescript is statically typed! (checked during development -- but not at runtime)

## 13. Important: Type Casing
- Types in typescript are all lowercase, ex: `number` and `string`

## 14. Working with Numbers, Strings & Booleans
- All numbers are floats by defaults
- If you try to print (number + number + string) it'll concatenate everything as a string
    - but printing  (number + string) works, eh

## 15. Type Assignment & Type Inference
- Typescript is able to infer a type from an assignment
- That's one of the biggest tasks for typescript: Checking types and tell you if you're using them incorrectly

## 16. Object Types
- Ooooh there's an object type!
- Hmm, seems like a golang struct
- If you try to access a property of a object that doesn't exist, typescript will error out at compile time
- ooo typescript infers

## 18. Array types!
- Neat, regular arrays, not much to it

## 19. Tuples
- Fixed lenght and fixed type array
- Seems handy for arrays that must have exactly a certain number of elements
- ruh-roh a union type came up, neat. TS can infer a type, but you may want to overwrite that
- You can use a push on a tupple, but you can't assign the wrong type on a tuple directly if you try to

## 20. Enums
- Enumerated list of global constant identifiers
- Your first custom type
- Good for identifiers that are human readable and have some mapped value behind the scenes

## 21. The any type
- Avoid using any whenever possible -- it takes away the advantages typescript gives you
- You can use any as a fallback for when you really don't know what type of data you'll get

## 22. Union types
- Allows you to have assume something might be of certain types and then handle those types in a certain way

## 23. Literal Types
- Literal types are types in which you're very clear about that type something should hold (uh-oh)
- Literal types are based on your core types, but with a specific value

## 24. Type Aliases / Custom Types
- You can create a type to represent a union type
- Actually kinda cool!

## 26. Function return types & "void""
- Function return types, like in Go!
- Usually there's no need to set the return type
- void just means this function doesn't have a return statement
- If you try to print the return statement of a void function, you'll get undefined, although if you try to return undefined explicitly, TS errors out
- void is the type you want if you have no return statement

## 27. Functions as types
- You can assign a function to a variable
- There's a function type!
- ooo you can create a type that must match a given function signature (just like in go!)

## 28. Function types and callbacks
- You can pass a function as a parameter to another function and do a callback function thing!
- The function that you pass as parameter must meet the signature of the function that requires it
- It was a bit hard to interpret everything the autor meant here, left a note to review later!

## 29. The unknown type
- uh-oh, for when you don't know the type of something, like for user input!
- unknown is a better choice to any if you know what you want to do with it eventually -- add some if checks
- unknown

## 30. The "never" type
- never is a type functions can return
- if a function throws an error, it doesn't return anything, not even void