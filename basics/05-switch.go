package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	switch a {
	case 1:
		fmt.Println("first")
	case 2:
		fmt.Println("second")
	default:
		fmt.Println("default")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noonday")
	default:
		fmt.Println("It's after noonday")
	}
}
