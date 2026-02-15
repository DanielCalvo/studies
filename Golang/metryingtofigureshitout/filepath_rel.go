package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	base := "/repos/kubernetes/kubernetes"
	target := "/repos/kubernetes/kubernetes/docs/setup/install.md"

	rel, err := filepath.Rel(base, target)
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // docs/setup/install.md

	base = "/repos/kubernetes/kubernetes/docs"
	target = "/repos/kubernetes/kubernetes/README.md"

	rel, _ = filepath.Rel(base, target)
	fmt.Println(rel) // oooh nice, its smart enough to figure out other paths!: ../README.md

}
