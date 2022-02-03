package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "/etc"

	var folders []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			folders = append(folders, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, folder := range folders {
		fmt.Println(folder)
	}

}
