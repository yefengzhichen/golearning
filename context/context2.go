package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		fmt.Println("before close")
		// close(ch)
	}()

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	//defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ch:
		fmt.Println("ch recieve")
	}
}
