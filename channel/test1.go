package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
	}()
	fmt.Println(<-c)
	close(c)
	time.Sleep(time.Second)
	fmt.Println("main end")
}
