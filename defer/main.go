package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("defer main")
	var user = ""

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success . err:", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()
			if user == "" {
				fmt.Println("user is empty")
				panic("should set user env.")
			}
		}()
		fmt.Println("end of func")
	}()
	time.Sleep(time.Second)
	fmt.Println("end")

}
