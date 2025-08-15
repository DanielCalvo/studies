package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	data, err := os.ReadFile("file.yaml")
	if err != nil {
		panic(err)
	}

	var content map[string]interface{}
	if err := yaml.Unmarshal(data, &content); err != nil {
		panic(err)
	}

	// Print the full structure
	fmt.Printf("YAML contents:\n%v\n\n", content)

	// Access fields dynamically
	for key, value := range content {
		fmt.Printf("Key: %s, Value: %#v\n", key, value)
	}
}
