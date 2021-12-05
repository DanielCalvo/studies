package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Fileyaml struct {
	One   string `yaml:"one"`
	Two   string `yaml:"two"`
	Three string `yaml:"three"`
}

func main() {
	fmt.Println("hello world")

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	yamlFile, err := ioutil.ReadFile(pwd + string(os.PathSeparator) + "file.yaml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var fileyaml Fileyaml
	err = yaml.Unmarshal(yamlFile, &fileyaml)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Println(fileyaml)
}
