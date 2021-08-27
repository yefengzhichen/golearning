package main

import "fmt"

type People struct {
}

func (p People) Print(str string) {
	fmt.Println("people")
}

type Student struct {
	People
}

func (stu *Student) Print(str string) {
	if str == "boy" {
		fmt.Println("a boy")
	} else {
		fmt.Println("a people")
	}
}

type IAnimal interface {
	Print(str string)
}

type Animal struct {
}

func (ani *Animal) Print(str string) {
	fmt.Println(str)
}

func main() {
	p := Student{}
	p.Print("boy")

	var a IAnimal = Animal{}
	a.Print("aa")
}
