package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Machines []struct {
	Name            string   `yaml:"name"`
	DownloadsFolder string   `yaml:"downloads_folder"`
	UsedPrograms    []string `yaml:"used_programs"`
}

func main() {

	m := Machines{}

	yamlFile, err := ioutil.ReadFile("/home/daniel/PycharmProjects/studies/Golang/WebdevWithGo/Section2_Templates/11_datastructuresintotemplates/homedir.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	tpl, err := template.ParseFiles("/home/daniel/PycharmProjects/studies/Golang/WebdevWithGo/Section2_Templates/11_datastructuresintotemplates/homedir.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatal(err)
	}

}
