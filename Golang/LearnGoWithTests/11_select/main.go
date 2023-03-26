package main

import (
	"fmt"
	"net/http"
	"time"
)

var oneSecondTimeout = 1 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, oneSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// a chan struct{} is the smallest data type available from a memory perspective
func ping(url string) chan struct{} {
	//Always make channels!
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
