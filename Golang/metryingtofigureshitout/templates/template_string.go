package main

import (
	"os"
	"text/template"
)

func main() {
	srcString := `Hello my name is {{.Name}}.
My favourite food is {{.Food}}.
Today is {{.DayOfWeek}}`

	type MyData struct {
		Name      string
		Food      string
		DayOfWeek string
	}

	myData := MyData{
		Name:      "Daniel",
		Food:      "Oats",
		DayOfWeek: "Sunday",
	}

	tmpl, err := template.New("test").Parse(srcString)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, myData)
	if err != nil {
		panic(err)
	}

}
