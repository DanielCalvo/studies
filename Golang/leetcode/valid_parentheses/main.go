package main

import (
	"fmt"
	"sort"
	"strings"
)

// Maybe not the best naming -- but lets roll with it for now
const openers = "([{"
const closers = ")]}"

func main() {

	//fmt.Println("isValid3:",isValid3("()"))
	//fmt.Println(isValid2("(("))
	//fmt.Println(isValid2("(())[[]{})()"))

	//----
	//fmt.Println(isValid3("(())[[]{})()"))
	//fmt.Println(isValid3("(())[[]{}]()"))
	//fmt.Println("empty string", isValid3(""))
	//fmt.Println(")", isValid3(")"))
	//fmt.Println("(", isValid3("("))
	//fmt.Println("[]", isValid3("[]"))
	//fmt.Println("(()", isValid4("(()")) //lets do parenthesis only first!
	fmt.Println("())", isValid4("())"))
	//fmt.Println("()", isValid4("()"))

	//fmt.Println("([)", isValid3("([)"))
	//fmt.Println("()", isValid3("()"))
	//fmt.Println("([]([])[])()", isValid3("([]([])[])()"))
	//fmt.Println("(())()())", isValid3("(())()())"))
	//fmt.Println("((((()))(())))", isValid3("((((()))(())))"))
}

func isValid5(s string) bool {

	return false
}

func isValid4(s string) bool {
	/*
			1. found an opener
		 	2. Is there a closer somewhere down the line?
			if yes: Remove closer and opener from string
			if a closer is not found: return false
			if len(stack) > 0 after all of this, return false
	*/
	//ah crap you probably can do this iterating only once -- but how? lets try improving this later, lets see if we can make it work first
	for k, v := range s {
		if v == '(' {
			for kk, vv := range s {
				if vv == ')' {
					fmt.Println("Found opener at:", k)
					fmt.Println("Found closer at:", kk)
				}
			}
		}
	}
	if len(s) > 0 {
		return false
	}
	return true

}

// Convert to generics!
func removeCharAtIndexes(s string, elements ...int) string {
	sort.Ints(elements)
	for i := len(elements) - 1; i >= 0; i-- { //Iterate through elements from end to beginning
		indexToRemove := elements[i] //for better clarity
		s = s[:indexToRemove] + s[indexToRemove+1:]
	}
	return s
}

// doesn't work, misses stray closers like (]), but otherwise functions
func isValid3(s string) bool {
	if len(s) == 0 || !strings.Contains(openers, string(s[0])) { //if empty or doesn't start with an opener
		return false
	}

	var sl_openers []string

	for _, v := range s {
		if strings.Contains(openers, string(v)) {
			sl_openers = append(sl_openers, string(v))
		}

		if len(sl_openers) == 0 {
			return false
		}

		if v == ')' && sl_openers[len(sl_openers)-1] == string('(') {
			fmt.Println("Got here )")
			sl_openers = sl_openers[:len(sl_openers)-1] //was using RemoveLastElement here but changed it so it all fits ina single function
		}
		if v == ']' && sl_openers[len(sl_openers)-1] == string('[') {
			fmt.Println("Got here ]")
			sl_openers = sl_openers[:len(sl_openers)-1]
		}
		if v == '}' && sl_openers[len(sl_openers)-1] == string('{') {
			fmt.Println("Got here }")
			sl_openers = sl_openers[:len(sl_openers)-1]
		}
	}

	if len(sl_openers) > 0 { //unclosed opener
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
