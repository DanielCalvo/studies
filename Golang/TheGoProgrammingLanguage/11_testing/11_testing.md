
## 11.1. The go test tool
The go test command runs test for Go packages. Files ending in _test.go are not built by go build, but built by go test.

Inside these test files you can have three kind of functions:
- tests
- benchmarks
- examples

---

- A test function is a function whose name starts with Test
- A benchmark function begins with Benchmark
- An example function begins with Example

---

- A test function exercises some logic to check for correct program behavior
- A benchmark function measures the performance of some operation
- An example function provides machine checked documentation. This looks really cool! I should try writing one of these

## 11.2. Test Functions
Alright listen up the naming of stuff is important:
- Each testing file must import the testing package

Test functions must:
- Begin with the word Test
- Have this following signature:

```go
func TestName(t *testing.T) {
// ...
}
```

The suffix name is optional, but it must begin with a capital letter (Name, not name).
The t parameter provides methods for reporting test failures and logging 


