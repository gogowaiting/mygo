package main

import (
	"fmt"
	"sync"
)

// 生产者：发送数据后关闭 channel
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 关闭 channel，for-range 才能正常结束
}

func main() {
	// for-range 消费 channel
	// range 会持续接收直到 channel 被关闭
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go producer(ch, &wg)

	for elem := range ch {
		fmt.Printf("received: %d\n", elem)
	}
	wg.Wait()
	fmt.Println("channel closed, all done")

	// 实际场景：fan-out 工作分发
	tasks := make(chan string, 3)
	go func() {
		for _, task := range []string{"task-a", "task-b", "task-c"} {
			tasks <- task
		}
		close(tasks)
	}()

	for t := range tasks {
		fmt.Println("processing:", t)
	}
}
