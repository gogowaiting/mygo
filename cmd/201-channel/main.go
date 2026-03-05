package main

import "fmt"

// channel 定义，空值为nil，申明配合make使用
//  var 名称 chan 类型

func main() {
	//无缓冲，同步

	ch0 := make(chan string)
	go func() {

		ch0 <- "hello, no cahce channel" // 阻塞，知道有goroutine接收
	}()
	msg := <-ch0 // 阻塞，直到有gorutine发送
	fmt.Println(msg)

	//有缓冲，异步
	ch1 := make(chan int, 3)
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch1)
	}

	// for-range
	ch := make(chan int, 10)
	// 生产者
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	// 消费者
	go func() {
		fmt.Println("sub goroutine staring")
		for num := range ch {
			fmt.Println(num)
		}
	}()
	fmt.Println("main goroutine over")

}
