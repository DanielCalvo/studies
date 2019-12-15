package main

import (
	"fmt"
	//"gopkg.in/src-d/go-git.v4"
	"os"
)

//if file does not exist, clone
//if file is dir and exists, skip
//if dir exists, do a cheackout
//take branch as parameter?
//if no parameter assume master

func main() {

	//stopped here: https://github.com/src-d/go-git/blob/master/_examples/checkout/main.go
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		fmt.Println("path/to/whatever does not exist")
	}

	//_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
	//	URL:      "https://github.com/src-d/go-git",
	//	Progress: os.Stdout,
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err := filepath.Walk("/tmp/go-git", func(path string, info os.FileInfo, err error) error {
	//	if err != nil {
	//		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
	//		return err
	//	}
	//
	//	if strings.HasSuffix(info.Name(), ".md"){
	//
	//		getLinkFromMD(path)
	//	}
	//
	//	return nil
	//
	//})
	//if err != nil {
	//	fmt.Printf("error walking the path %q: %v\n", err)
	//	return
	//}

}
