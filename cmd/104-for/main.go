package main

import "fmt"
import "time"

func main() {
	i := 1
	// 死循环
	for {
		fmt.Println("loop...")
		time.Sleep(time.Second * 5)
		break
	}
	// 表达式循环
	for n := 0; n <= 8; n++ {
		if n%2 == 0 {
			continue
		}
	}

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

loop:
	for i := 0; i < 10; i++ {
		if i == 5 {
			break loop
		}
		fmt.Print(i)
	}
}
