package main

import (
	"sync"
	"testing"
)

// go test main.go main_test.go
func TestSafeCounter(t *testing.T) {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	numGoroutines := 1000
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()

	if got := counter.Value(); got != numGoroutines {
		t.Errorf("Expected counter to be %d, but got %d", numGoroutines, got)
	}
}
