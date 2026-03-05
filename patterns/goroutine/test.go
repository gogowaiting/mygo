package main

import "fmt"

func testSelect() {
	fmt.Println("test select")
	ch1 := make(chan int, 1)
	select {
	case <-ch1:
		fmt.Println("Received from ch1")
	default:
		fmt.Println("this is default")

	}
	fmt.Println("test select end")

}
