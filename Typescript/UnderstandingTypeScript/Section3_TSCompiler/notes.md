## 34. Using watch mode
- How to work with multiple files!
- `tsc app.ts --watch`

## 35. Compiling the Entire Project / Multiple Files
- `tsc --init`
- After this you can just run `tsc`
- You can exclude certain files from being compiled (ex: `*.dev.ts`)
- node_modules is excluded by default

## 36. Including & Excluding Files
- Went over include and excludes

## 37. Setting a Compilation Target
- Most options do not matter
- Important ones
    - target: target JS version to compile the coder

## 38. Understanding TypeScript Core Libs
- Lib: Something to do with defaults (you could need to change this if you're writing backend code)
- Neat, some options were shown!

## 39. More Configuration & Compilation Options
- allowJS: You can include JS files in the compilation!
- checkJS will not compile it but check the syntax on report errors
    - You can use this to check vanilla JS!

## 40. Working with Source Maps
- Helps with debug 
- If set to true, you can see your typescript files in the browser, it does some sort of bridging, you can work with TS on the browser?

## 41. rootDir and outDir
- OutDir, a directory for the output javascript (aka `dist`)
- Don't forget to adjust the imports in the html! 
- rootDir: Explicitly say where your TS files are to be compiled (excludes other dirs I think1)
- downlevelIteration is about some compatibility issue with older versions of javascript (something about for loops not working well in certain rare cases)

## 42. Stop Emitting Files on Compilation Errors
- noEmitOnError: TS can create JS files even on errors (like when looking for a button that does not exist)
- defaults to false, which will generate JS files even if you have an error. If set to true, problematic files will not be generated

## 43. Strict Compilation
- Typechecking options!
- `noImplicytAny`, neat. Allows you to not have any implicit types (like in a function argument without a type specified)
- for variables its okay, but for functions it is not!
- `strictNullChecks` anticipates for missing things that will result in null (like a missing button) and fails your compilation
- `strictFunctionTypes`: This checks in which function you're calling .bind() or .apply() and it checks if what you're setting makes sense (this is a bit advanced for my current level of knowledge)
- `alwaysStrict` ensures the JS being compiled always usses strict mode!

## 44. Code Quality Options
- `noUnusedLocals` and `noUnusedParameters` force you to not have unused instances of these
- Global variables remain allowed however
- `noImplicitReturns` makes that your function must always explicitly return something 

## 45. Debugging with Visual Studio Code
- Instructions on how to debug with VS code (dive in more deeply if you need to)