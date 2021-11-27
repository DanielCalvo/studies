package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	//Switch statements express conditionals across many branches
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) { //Seems like type can only be used in a switch?
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Println("Not too sure what type that is")
		}
	}
	whatAmI(1)
	whatAmI(true)

}
