package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	myJson := `{"Name":"Bob McBobson","Age":1}`
	fmt.Println("myJson:", myJson)

	enc := json.NewEncoder(os.Stdout)

	must(enc.Encode(myJson))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
