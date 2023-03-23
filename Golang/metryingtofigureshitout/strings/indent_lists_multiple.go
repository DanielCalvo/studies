package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func main() {

	md := `            - Heres a list that starts at zero
 - Then goes to four
  - Then to eight
 - Then to four
- Yo, another element at zero
asd
	asd
` //The above should be indented with 0, 4 and 8 whitespaces, respectively
	//processListItem(md)

}

/*
go through lines
counter the number of indents
if current line has more indented than previous line
	indentTier++
if current line has the same indents as previous line
	indentTier stays as is
if current line has less indents than previous line
	indentTier--
*/

//To indent properly in all occasions, seems like information about the next and previous lines is required
//type list struct{
//	line string
//	numberOfIndents int
//	indentTier int
//}

// This should've been called "indentButHackily"
func processListItem(s string) string {
	scanner := bufio.NewScanner(strings.NewReader(s))

	//indentTier := 0
	//var curretLine string
	//i := 0
	previousLine := ""
	currentLine := ""
	indentTier := 0
	counter := -1

	for scanner.Scan() {
		counter++
		currentLine = scanner.Text()

		//I'm gonna do something "unelegant" just to move forward for now
		if counter == 0 {
			//fmt.Println("i",GetNumberOfIndents(currentLine),"T", indentTier,"D", indentTier*4,"sAME:", currentLine)
			fmt.Println(plsIndent(currentLine, indentTier))
			continue
		}

		//So just indent things to indentTier * 4 and YOLO it?
		if GetNumberOfIndents(currentLine) > GetNumberOfIndents(previousLine) {
			indentTier++
			//fmt.Println("i",GetNumberOfIndents(currentLine),"T", indentTier, "D", indentTier*4, "MORE:", currentLine)
			fmt.Println(plsIndent(currentLine, indentTier))
		}
		if GetNumberOfIndents(currentLine) == GetNumberOfIndents(previousLine) {
			//fmt.Println("i",GetNumberOfIndents(currentLine),"T", indentTier,"D", indentTier*4,"SAME:", currentLine)
			fmt.Println(plsIndent(currentLine, indentTier))
		}
		if GetNumberOfIndents(currentLine) < GetNumberOfIndents(previousLine) {
			if GetNumberOfIndents(currentLine) == 0 {
				indentTier = 0
			} else {
				indentTier--
			}
			//fmt.Println("i",GetNumberOfIndents(currentLine),"T", indentTier,"D", indentTier*4,"LESS:", currentLine)
			fmt.Println(plsIndent(currentLine, indentTier))
		}

		previousLine = currentLine

	}
	return s
}

func plsIndent(s string, i int) string {
	if !isList(s) {
		return s
	}

	trimmedString := strings.TrimLeftFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	var indentPrefix string
	for a := 0; a < i; a++ {
		indentPrefix = "    " + indentPrefix
	}
	return indentPrefix + trimmedString
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

func roundup(numToRound int, multiple int) int {
	remainder := numToRound % multiple
	return numToRound + multiple - remainder
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
