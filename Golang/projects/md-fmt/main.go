package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {

	//This is currently a very dumb program that only indents some types of markdown lists

	filename := "/home/daniel/Projects/notes/INBOX.md"

	src, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	processedSrc := processFileContents(string(src))

	fmt.Println(processedSrc)

	tempFile, err := ioutil.TempFile("", "INBOX.md")

	if err != nil {
		panic(err)
	}

	_, err = tempFile.WriteString(processedSrc)
	if err != nil {
		panic(err)
	}

}

//This function is a big mess
func processFileContents(s string) string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var buf bytes.Buffer

	previousLine := ""
	currentLine := ""
	indentTier := 0
	counter := -1
	for scanner.Scan() {
		counter++
		currentLine = scanner.Text()

		if isList(currentLine) {
			if counter == 0 {
				buf.WriteString(plsIndent(currentLine, indentTier) + "\n")
				continue
			}

			currentLineIndent := GetNumberOfIndents(currentLine)
			previousLineIndent := GetNumberOfIndents(previousLine)

			if currentLineIndent > previousLineIndent {
				indentTier++
				buf.WriteString(plsIndent(currentLine, indentTier) + "\n")
			}
			if currentLineIndent == previousLineIndent {
				buf.WriteString(plsIndent(currentLine, indentTier) + "\n")
			}
			if currentLineIndent < previousLineIndent {
				if currentLineIndent == 0 {
					indentTier = 0
				} else {
					indentTier--
				}
				buf.WriteString(plsIndent(currentLine, indentTier) + "\n")
			}
			previousLine = currentLine
		} else {
			buf.WriteString(currentLine)
		}
	}
	return buf.String()
}

func ProcessList(s string) string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var buf bytes.Buffer

	previousLine := ""
	currentLine := ""
	indentTier := 0
	counter := -1
	for scanner.Scan() {
		counter++
		currentLine = scanner.Text()

		//I'm not particularly proud of this if statement, certainly there's a better way of doing this
		if counter == 0 {
			buf.WriteString(plsIndent(currentLine, indentTier) + "\n")
			continue
		}

		currentLineIndent := GetNumberOfIndents(currentLine)
		previousLineIndent := GetNumberOfIndents(previousLine)

		if currentLineIndent > previousLineIndent {
			indentTier++
			buf.WriteString(plsIndent(currentLine, indentTier) + "\n")
		}
		if currentLineIndent == previousLineIndent {
			buf.WriteString(plsIndent(currentLine, indentTier) + "\n")
		}
		if currentLineIndent < previousLineIndent {
			if currentLineIndent == 0 {
				indentTier = 0
			} else {
				indentTier--
			}
			buf.WriteString(plsIndent(currentLine, indentTier) + "\n")
		}
		previousLine = currentLine
	}
	return buf.String()

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
