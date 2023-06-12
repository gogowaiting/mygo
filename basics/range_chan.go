package main

import (
	"fmt"
)

func addElement(chanName chan int) {
	for i := 0; i < 10; i++ {
		chanName <- i
	}
}

// 遍历 chan
func main() {
	var a = make(chan int, 10)
	addElement(a)
	for elem := range a {
		fmt.Printf("get element for chan: %d\n", elem)
	}

}
