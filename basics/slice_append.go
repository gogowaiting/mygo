package main

import "fmt"

func addElement() []string {
	var testSlice []string
	a, b := "a", "b"
	testSlice = append(testSlice, a)
	testSlice = append(testSlice, b)
	fmt.Println(testSlice)
	return testSlice
}

func main() {
	addElement()
}
