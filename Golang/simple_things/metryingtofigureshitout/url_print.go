package main

import (
	"fmt"
	"net/url"
)

func main() {

	dest := "/docs/using-client-keystone-auth.md/"

	url, err := url.ParseRequestURI(dest)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}

	fmt.Println(url.Path)

}
