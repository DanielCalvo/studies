package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

/*
If you uncomment the lock, you allow multiple goroutines to modify the counter at the same time without any coordination
Multiple goroutines can read the same value before any of them write it back, so two routines could, say, try to add 1 to 42
Which is why the end result is sometimes less than 100 if you remove the lock
*/
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	counter := SafeCounter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter.Value())
}
