/*
Hey lets get a context with a timeout and just do nothing and let it time out!
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	//We're not doing anything, so this will block for 100ms and then print
	fmt.Println("Context timed out: ", <-ctx.Done())

	ctx, cancel = context.WithCancel(context.Background())
	cancel()
	fmt.Println("uh it was cancelled")

}
