package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	//_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
	//	URL:      "https://github.com/src-d/go-git",
	//	Progress: os.Stdout,
	//})
	//
	//if err != nil {
	//	fmt.Println("Couldn't clone")
	//}

	err := filepath.Walk("/tmp/foo", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", err)
		return
	}

}
