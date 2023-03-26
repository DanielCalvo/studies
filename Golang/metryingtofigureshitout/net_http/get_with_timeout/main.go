package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("heya")
	fmt.Println(HeadUrlWithTimeout("https://google.com", time.Second))
	fmt.Println(HeadUrlWithTimeout("https://localhost:9999", time.Second))
}

func HeadUrlWithTimeout(url string, timeout time.Duration) (int, error) {
	select {
	case httpStatusCode := <-HTTPHeadUrl(url):
		return httpStatusCode, nil
	case <-time.After(timeout):
		return 404, fmt.Errorf("timeout on http.Head() for %s", url)
	}
}

func HTTPHeadUrl(url string) chan int {
	ch := make(chan int)
	go func() {
		head, _ := http.Head(url)
		ch <- head.StatusCode
		close(ch)
	}()
	return ch
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
