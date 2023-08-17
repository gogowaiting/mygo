package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type student struct {
	person
	id int
}

func main() {
	// var testMap map[string]int
	m := make(map[string]int)
	m["one"] = 1
	fmt.Println(m)
	var sl []string
	s := make([]string, 0)

	sl = append(sl, "one")
	s = append(s, "two")
	fmt.Println(sl)
	fmt.Println(s, len(s))

	var s2 = student{person{"Tom", 18}, 101}
	fmt.Println(s2)

}
