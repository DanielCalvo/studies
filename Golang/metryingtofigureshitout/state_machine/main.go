package main

import (
	"fmt"
	"strings"
)

const (
	red   = "\033[31m"
	reset = "\033[0m"
)

// the idea for this function comes from parsing the response headers on github api calls, i wanted to split the string in the link header by comma, but only for commas outside of the <> diamonds on the string, so something like this: "<hello, world>, <heya>" would yield ["<hello, world>", "<heya>"] -- i dont want the commas inside the <> to be counted
// it seems a state machine is a way to accomplish this -- i just want strings.split() with extra steps
func main() {

	s := "<hel>, <uh, oh, we>,  <no commas here!>; <we just had a semicolon, that shouldn't cause these to split!>"

	var finalSplit []string
	var res string
	var in bool

	//let me try a different approach
	//ah, much cleaner than my first one!
	for _, v := range s {
		if in {
			res += string(v)
			if v == '>' {
				in = false
			}
		} else {
			if v == ',' {
				finalSplit = append(finalSplit, res)
				res = ""
				continue
			}
			if v == '<' {
				in = true
			}
			res += string(v)
		}
	}
	finalSplit = append(finalSplit, res)

	for _, v := range finalSplit {
		fmt.Println("---", strings.TrimSpace(v))
	}

	/*
				this was what codex suggested in the end.
		    	it shared an interesting insight:
			  A useful rule of thumb:
			  - if a character has special meaning, giving it its own branch often makes parser code easier to read

				  if in {
				  	res += string(v)
				  	if v == '>' {
				  		in = false
				  	}
				  } else {
				  	if v == ',' {
				  		finalSplit = append(finalSplit, res)
				  		res = ""
				  	} else if v == '<' {
				  		res += string(v)
				  		in = true
				  	} else {
				  		res += string(v)
				  	}
				  }


	*/

	//this was the first implementation in which i wrote too many if statements
	//for _, v := range s {
	//	//we're not interested in whitespace outside <>
	//	if v == '<' {
	//		in = true
	//	}
	//	if v == '>' {
	//		in = false
	//	}
	//	//if you're in <> you always append no matter what
	//	if in {
	//		res += string(v)
	//	}
	//	//if youre out and youre not a comma, we want to append that too!
	//	if !in && v != ',' {
	//		res += string(v)
	//	}
	//	//if you're outside and you stumble upon a comma, that is your split condition!
	//	//this fails to append the last element
	//	if !in && v == ',' {
	//		finalSplit = append(finalSplit, res)
	//		res = ""
	//	}
	//}
	////append the last element as you dont stumble upon a comma on the last iteration of the loop to append it
	//finalSplit = append(finalSplit, res)

}
