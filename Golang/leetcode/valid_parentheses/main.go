package main

import "fmt"

// Maybe not the best naming -- but lets roll with it for now
const openers = "([{"
const closers = ")]}"

func main() {

	fmt.Println(AreParenthesisValid("(())"))
	fmt.Println(AreParenthesisValid("()"))
	fmt.Println(AreParenthesisValid("(()())"))
	fmt.Println(AreParenthesisValid("((()())())"))
	fmt.Println(AreParenthesisValid("(()()())((()))"))

}

func AreParenthesisValid(s string) bool {
	//if len(s) == 0 {
	//	return false
	//}
	////s must contain an even amount of parenthesis to be valid
	//if len(s)%2 != 0 {
	//	return false
	//}
	////Must start with opener
	//if s[0] != '(' {
	//	return false
	//}
	////Must end with closer
	//if s[len(s)-1] != ')' {
	//	return false
	//}
	//if string contains anything other than parenthesis {
	//return false
	//}

	openerPos := -1
	closerPos := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			openerPos = i
		}
		if s[i] == ')' {
			closerPos = i
		}
		if openerPos != -1 && closerPos != -1 {
			s = s[:closerPos] + s[closerPos+1:] //rm closer
			s = s[:openerPos] + s[openerPos+1:] //rm opener
			openerPos = -1                      //set position to -1 so it has to match again
			closerPos = -1                      //set position to -1 so it has to match again
			i = -1                              //restart counter so it starts from the beginning of the string again -- this -1 is odd, I'm probably doing something wrong here
		}
	}

	if len(s) == 0 {
		return true
	}

	return false
}
