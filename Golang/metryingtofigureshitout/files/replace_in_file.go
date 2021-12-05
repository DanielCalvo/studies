package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*
		You can't remove a line in a file without eithe reading the file to memory or iterating through the file, you googled this a lot!
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
