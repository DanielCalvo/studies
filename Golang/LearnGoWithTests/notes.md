## Chapter 1
Writing a test is just like writing a function, with a few rules: 
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing"
Testing package docs: https://pkg.go.dev/testing

## Chapter 2
Oooh examples are a thing: https://go.dev/blog/examples

## Noice!
- Continue here: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#refactor-2
    - (on the sum all tails part)
- AAAA you learned about functions with variable initialization on the header, what was the name of that again?
    - Go over each thing you learned on each chapter (the comments!) and leave a reference here!
- `go test -cover` can show you your test coverage 
- I learned about `reflect.DeepEqual` to compare slices and that was cool!