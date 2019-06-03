package main

import (
	"fmt"
	"strings"
)

func main() {

	mystring1 := "This is is a string that contains dbt"
	mystring2 := "This another string also contains the word I'm looking for: dbt"

	if strings.Contains(mystring1, "dbt") || strings.Contains(mystring2, "dbt") {
		fmt.Println("it matched!")
	} else {
		fmt.Println("No match")
	}

}
