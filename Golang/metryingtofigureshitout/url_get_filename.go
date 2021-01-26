package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	myUrl := "https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/alb-ingress.md"
	parsedUrl, err := url.ParseRequestURI(myUrl)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Path)
	fmt.Println(path.Base(parsedUrl.Path)) //alb-ingress.md

}
