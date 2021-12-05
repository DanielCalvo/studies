package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"ToUpper": strings.ToUpper,
}

func main() {

	//So this works:
	tpl, err := template.ParseFiles("/tmp/myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, "banana")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("End of simple.go")

}
