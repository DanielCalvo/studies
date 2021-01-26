package main

import (
	"fmt"
	"strings"
)

func main() {

	s := `Referer: http://localhost:8080/banana/asdasdasd/asd`
	fmt.Println(s)

	sp := strings.Split(s, "/")
	fmt.Println(sp)

	fmt.Println(sp[3:])

	for _, v := range sp {
		fmt.Println("v is:", v)
	}
}
