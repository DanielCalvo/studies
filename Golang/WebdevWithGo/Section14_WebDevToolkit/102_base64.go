package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	s := "Hey so this is my string and I have characters in my string!"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

}
