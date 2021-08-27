package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx, _ := context.WithCancel(context.Background())
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("aa")
		cancel()
	}()

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	//defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
