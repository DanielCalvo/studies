package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"image"
	"io/ioutil"
	"os"
)

type Img struct {
	Img      image.Image
	Filename string
	SrcDir   string
	DstDir   string
	Ratio    int
}

type Yamelsito struct {
	SrcDir string `yaml:"SrcDir"`
	DstDir string `yaml:"DstDir"`
	Ratio  []int  `yaml:"Ratio"`
}

func main() {

	//Gives current dir or program
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	yamlFile, err := ioutil.ReadFile(pwd + "/" + "parameters.yml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var yamelsito Yamelsito
	err = yaml.Unmarshal(yamlFile, &yamelsito)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Println(yamelsito.Ratio)
	fmt.Println(yamelsito.DstDir)
	fmt.Println(yamelsito.SrcDir)

}
