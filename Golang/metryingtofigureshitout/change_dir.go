package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	_ = CheckTmp()

	//The directory change on another function does not influence the directory of the main function
	fmt.Println(dir)

}

func CheckTmp() error {
	err := os.Chdir("/tmp")
	if err != nil {
		return err
	}
	return nil
}
