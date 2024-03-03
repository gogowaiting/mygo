package main

import "fmt"

// channel 定义，空值为nil，申明配合make使用
//  var 名称 chan 类型

func main() {

	ch := make(chan int, 10)

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	go func() {
		// fmt.Println("sub goroutine staring")
		for num := range ch {
			fmt.Println(num)
		}
	}()
	fmt.Println("main goroutine over")

}
