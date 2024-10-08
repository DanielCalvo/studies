Review Chapter 6 before wrapping it up!

Continue here: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors

## Chapter 1
Writing a test is just like writing a function, with a few rules: 
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing"
Testing package docs: https://pkg.go.dev/testing

## Chapter 2
Oooh examples are a thing: https://go.dev/blog/examples

## Chapter 5:
Table driven tests are a thing: https://github.com/golang/go/wiki/TableDrivenTests

## Chapter 6:
The `assertError` function to check which error you're getting is really cool for some reason! 

## Chapter 7: maps
- The key type of a map must be of a comparable type
- https://go.dev/ref/spec#Comparison_operators
- The value type can be whatever you want though, even another map
- Remember that map lookups can return 2 values!
- A map value is a pointer to a runtime.hmap structure

## Chapter 8: Dependency injection
- Brief chapter

## Chapter 9: Mocking
- Cool chapter, goes into detail on testing a print/sleep function
- Has good advice about mocking in the end!

## Chapter 11: Select
- Has an example of `how you would write a real HTTP server in Go` 

## Chapter 12: Reflect(ion)
- This chapter is mega confusing! 
- I'm not too familiar with reflection (that doesn't help) and the author doesn't clarify very well what we're setting out to accomplish here
- I think I should explore the topic of reflection in go independently and then come back to this more prepared
- Here's where you stopped:
    - https://pkg.go.dev/reflect#ValueOf
    - https://go.dev/blog/laws-of-reflection

## Chapter 14: Context
- In this chapter we'll use the package context to help us manage long-running processes.
- Hey an exercise idea that you had: Read lines from a file and after x seconds pass, time it out, like an http request but to a file!


## Noice!
- AAAA you learned about functions with variable initialization on the header, what was the name of that again?
    - Go over each thing you learned on each chapter (the comments!) and leave a reference here!
- `go test -cover` can show you your test coverage 
- I learned about `reflect.DeepEqual` to compare slices and that was cool!