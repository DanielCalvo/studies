package main

import (
	"fmt"
	"strings"
)

const openers = "([{"
const closers = ")]}"

func main() {

	//fmt.Println("isValid3:",isValid3("()"))
	//fmt.Println(isValid2("(("))
	//fmt.Println(isValid2("(())[[]{})()"))

	//----
	//fmt.Println(isValid3("(())[[]{})()"))
	//fmt.Println(isValid3("(())[[]{}]()"))
	fmt.Println("empty string", isValid3(""))
	fmt.Println(")", isValid3(")"))
	fmt.Println("(", isValid3("("))
	fmt.Println("[]", isValid3("[]"))
	fmt.Println("(])", isValid3("(])"))
	fmt.Println("([)", isValid3("([)"))
	fmt.Println("()", isValid3("()"))
	fmt.Println("([]([])[])()", isValid3("([]([])[])()"))
	fmt.Println("(())()())", isValid3("(())()())"))
	fmt.Println("((((()))(())))", isValid3("((((()))(())))"))
}
func isValid3(s string) bool {
	if len(s) == 0 || !strings.Contains(openers, string(s[0])) { //if empty or doesn't start with an opener
		return false
	}

	var stackerino []string

	for _, v := range s {
		if strings.Contains(openers, string(v)) {
			stackerino = append(stackerino, string(v))
		}

		if len(stackerino) == 0 {
			return false
		}

		if v == ')' && stackerino[len(stackerino)-1] == string('(') {
			stackerino = stackerino[:len(stackerino)-1] //was using RemoveLastElement here but changed it so it all fits ina single function
		}
		if v == ']' && stackerino[len(stackerino)-1] == string('[') {
			stackerino = stackerino[:len(stackerino)-1]
		}
		if v == '}' && stackerino[len(stackerino)-1] == string('{') {
			stackerino = stackerino[:len(stackerino)-1]
		}
	}

	if len(stackerino) > 0 { //unclosed opener
		return false
	}

	return true
}

func RemoveLastElement[T any](slice []T) []T {
	t1 := slice[:len(slice)-1]
	return t1
}

// Second try: This works on more strings than the first try but still fails for something like: `(())[[]{})()
func isValid2(s string) bool {
	//has to start with an opener
	if !strings.Contains(openers, string(s[0])) {
		fmt.Println(string(s[0]), openers)
		return false
	}
	//has to finish with a closer
	if !strings.Contains(closers, string(s[len(s)-1])) {
		return false
	}

	/*
		an opener can be followed by any opener or a closer of the same type
		a closer can have before it an opener of the same type, or any closer
	*/

	for i := 1; i < len(s)-1; i++ { //skip first and last, never goes out of bounds! :D

		switch s[i] {
		case '{':
			if s[i+1] != '}' && !strings.Contains(openers, string(s[i+1])) {
				return false
			}
		case '[':
			if s[i+1] != ']' && !strings.Contains(openers, string(s[i+1])) {
				return false
			}
		case '(':
			if s[i+1] != ')' && !strings.Contains(openers, string(s[i+1])) {
				return false
			}
		case '}':
			if s[i-1] != '{' && !strings.Contains(closers, string(s[i-1])) {
				return false
			}
		case ']':
			if s[i-1] != '[' && !strings.Contains(closers, string(s[i-1])) {
				return false
			}
		case ')':
			if s[i-1] != '(' && !strings.Contains(closers, string(s[i-1])) {
				return false
			}
		default:
			fmt.Println("invalid character")
			return false
		}

	}
	return true
}

// First implementation -- I assumed the problem only asked about sequential open-closed symbols
// but it also wants things like "[{}]
func isValid(s string) bool {
	//If the string doesn't start with {, [ or (, its not valid to begin with
	if s[0] != '{' && s[0] != '[' && s[0] != '(' {
		return false
	}

	for i := 0; i < len(s); i += 2 {
		fmt.Println(i)
		if s[i] == '{' {
			if s[i+1] != '}' {
				return false
			}
		}
		if s[i] == '[' {
			if s[i+1] != ']' {
				return false
			}
		}
		if s[i] == '(' {
			if s[i+1] != ')' {
				return false
			}
		}
	}
	return true
}
