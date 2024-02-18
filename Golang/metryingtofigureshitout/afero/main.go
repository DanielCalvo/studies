package main

import (
	"github.com/spf13/afero"
	"log"
)

func main() {

	var AppFs = afero.NewMemMapFs()
	f, err := AppFs.Create("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString("banana")
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	log.Println("Finished!")

}
