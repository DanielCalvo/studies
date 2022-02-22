package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	h := hmac.New(sha256.New, []byte("mykey"))
	_, err := io.WriteString(h, "banana")
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
}
