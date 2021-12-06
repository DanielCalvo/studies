package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {

	myMarkdown := `Here's one line with a following list
- One good list
 - One badly indented list
- One well indented list
   - aaaa
        - But what about nested lists?
Not a list at all
 Not a list at all with some whitespace at the beginning
      	 hey wait a minute`
	scanner := bufio.NewScanner(strings.NewReader(myMarkdown))

	var b bytes.Buffer

	for scanner.Scan() {
		if !isList(scanner.Text()) {
			b.WriteString(scanner.Text() + "\n")
			continue
		}
		if !isIndented(scanner.Text()) {
			b.WriteString(plsIndent(scanner.Text() + "\n"))
		}
	}

	fmt.Print(b.String())
	fmt.Println("all gucci")

}

func isList(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' {
		return true
	}

	for _, v := range s {
		if unicode.IsSpace(rune(v)) {
			continue
		}
		if v == '-' {
			return true
		}
		break
	}
	return false
}

func GetNumberOfIndents(s string) int {
	counter := 0
	for _, v := range s {
		if v == ' ' {
			counter++
			continue
		}
		if v == '\t' {
			counter = counter + 4
			continue
		}
		break
	}
	return counter
}

func plsIndent(s string) string {
	trimmedString := strings.TrimLeftFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	//four whitespaces
	return "    " + trimmedString
}

func isIndented(s string) bool {
	switch {
	case len(s) == 0:
		return true
	case s[0] == '-':
		return true
	case s[0] == ' ' && s[1] == ' ' && s[2] == ' ' && s[3] == ' ' && s[4] == '-':
		return true
	case s[0] == '\t' && s[1] == '-':
		return true
	case isList(s) && GetNumberOfIndents(s) > 4: //I'm going with good enough for now, I'll deal with multiple levels later
		return true
	default:
		return false
	}
}
