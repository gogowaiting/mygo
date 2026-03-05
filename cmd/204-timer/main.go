package main

import (
	"fmt"
	"time"
)

// Timer：单次定时触发
func timerDemo() {
	timer := time.NewTimer(100 * time.Millisecond)
	defer timer.Stop()

	fmt.Println("timer start:", time.Now())
	<-timer.C
	fmt.Println("timer fired:", time.Now())
}

// Ticker：周期性触发
func tickerDemo() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	count := 0
	for t := range ticker.C {
		count++
		fmt.Printf("  tick %d at %v\n", count, t.Format("15:04:05.000"))
		if count >= 3 {
			break
		}
	}
}

// select + timeout：超时控制
func timeoutDemo() {
	ch := make(chan string, 1)

	// 模拟耗时操作
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "result"
	}()

	// 100ms 超时
	select {
	case result := <-ch:
		fmt.Println("got result:", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout!")
	}

	// 500ms 超时 —— 足够等到结果
	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "result2"
	}()

	select {
	case result := <-ch2:
		fmt.Println("got result:", result)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout!")
	}
}

// AfterFunc：延迟执行回调
func afterFuncDemo() {
	fmt.Println("AfterFunc start:", time.Now())
	time.AfterFunc(100*time.Millisecond, func() {
		fmt.Println("AfterFunc callback:", time.Now())
	})
	time.Sleep(200 * time.Millisecond)
}

func main() {
	fmt.Println("=== Timer ===")
	timerDemo()

	fmt.Println("\n=== Ticker ===")
	tickerDemo()

	fmt.Println("\n=== Timeout ===")
	timeoutDemo()

	fmt.Println("\n=== AfterFunc ===")
	afterFuncDemo()
}
