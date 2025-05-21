package main

import (
	"context"
	"fmt"
	"time"
)

//C.A.T.V  (contexts always tracks values)
// cancelations, atomatic timeouts, timeout and dead line and Values

func doSomething(ctx context.Context) {
	for {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("doSomething")
		case <-ctx.Done():
			fmt.Println("timeout done")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("1 second")
		}
	}

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

	defer cancel()

	doSomething(ctx)
}
