package main

import "fmt"

type American struct{}
type Chinese struct{}

type Speaker interface {
	sayHello()
	sayBye()
}

func (a American) sayHello() {
	fmt.Println("american say Hello")
}

func (a American) sayBye() {
	fmt.Println("american say Bye")
}

func (c Chinese) sayHello() {
	fmt.Println("chinese say Hello")
}

func (a Chinese) sayBye() {
	fmt.Println("chinese say Bye")
}

func main() {
	var s Speaker
	a := American{}
	c := Chinese{}
	s = a
	s.sayHello()
	a.sayBye()
	s = c
	c.sayHello()
}
