package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "file.csv"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			fmt.Println("End of File")
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(row)
	}
}
