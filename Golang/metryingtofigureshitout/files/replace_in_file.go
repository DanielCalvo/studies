package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*
		I got in a conceptual kerfuffle, this doesn't make a lot of sense.
		You can't remove a line in a file without either reading the file to memory or iterating through the file
	*/

	//amagaaaaahd this turned out ugly
	file, err := os.Open("/home/daniel/Projects/studies/Golang/metryingtofigureshitout/files/file.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var modifiedLines []string

	for scanner.Scan() {
		currentLine := scanner.Text()
		if strings.Contains(currentLine, "banana") {
			currentLine = strings.Replace(currentLine, "banana", "apple", -1)
			modifiedLines = append(modifiedLines, currentLine)
		}
	}
	file.Close()

	//Not my proudest moment holding everything on memory but oh well

	if len(modifiedLines) > 0 {
		file, err = os.Create("/home/daniel/Projects/studies/Golang/metryingtofigureshitout/files/file.txt")
		if err != nil {
			panic(err)
		}
		for _, line := range modifiedLines {
			_, err = file.WriteString(line + "\n")
			if err != nil {
				continue
			}
		}
		file.Close()
	} else {
		fmt.Println("Nothing to write")
	}

}
