package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 2 {
		fmt.Println("hello, world!")
	} else {
		fmt.Println("please input two args")
	}

	fmt.Println("This is a test program...")
}
