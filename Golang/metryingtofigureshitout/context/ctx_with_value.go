package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("yo")

	ctx := context.Background()
	key := "people"
	value := "bob alice"
	ctxWithValue := context.WithValue(ctx, key, value)
	fmt.Println(ctxWithValue.Value(key))

}
