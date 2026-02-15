package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	root := "/home/daniel/Projects/studies/Golang/metryingtofigureshitout"
	// you write the logic inside the function call
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		// you get this error from walkdir calling the fn if fn(root, nil, err) if it errors out!
		// walkdir lets you handle the error for the root folder!
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(d.Name()) == ".go" {
			fmt.Println("Found go file:", path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, err)
	}
}
