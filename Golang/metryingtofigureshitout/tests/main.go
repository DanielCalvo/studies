package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Flowers struct {
	Flowers []string `yaml:"flowers"`
}

func main() {

	f, err := GetYaml("./a.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Flowers)
}

func GetYaml(path string) (Flowers, error) {
	var f Flowers

	file, err := os.ReadFile(path)
	if err != nil {
		return f, err
	}

	err = yaml.Unmarshal(file, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}
