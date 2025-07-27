// main.go
package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func main() {
	choices := []string{"1", "2", "3"}

	// Create a new prompt
	prompt := promptui.Select{
		Label: "Pick a number",
		Items: choices,
	}

	// Show the prompt and get the result
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	fmt.Printf("You selected: %s\n", result)
}
