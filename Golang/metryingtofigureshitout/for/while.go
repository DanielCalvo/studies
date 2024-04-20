package main

import "fmt"

func main() {
	//Go doesn't have a while statement, but you can use a for loop like one. For will not "iterate" over anything in this case, it'll just keep repeating -- its up to you to change the condition (ex: i < 10) so that the loop stops
	var i int
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
