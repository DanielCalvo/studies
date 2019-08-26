package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `huehuehue`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
