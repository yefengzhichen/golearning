package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	time.Sleep(time.Second)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		// if i == 0 {
		// 	close(c)
		// }
		fmt.Println("fibonacci ", x)
		x, y = y, x+y
	}
	//close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
		if i == 0 {
			fmt.Println("before close")
			close(c)
		}
	}
	fmt.Println("main end")
}
