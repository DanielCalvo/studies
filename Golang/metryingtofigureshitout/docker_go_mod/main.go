package main

import (
	"fmt"
	"github.com/DanielCalvo/markdownscanner/pkg/config"
	"github.com/DanielCalvo/markdownscanner/pkg/mdscanner"
	"os"
	//"github.com/DanielCalvo/markdownscanner/pkg/mdscanner"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("repo:", os.Getenv("repo"))

	c := config.Config{} //Set sane values for config if you find no file? (aka running from container in cmdline mode?)
	repo, err := mdscanner.NewRepository(&c, "https://github.com/DanielCalvo/markdownscanner")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(repo)

}
