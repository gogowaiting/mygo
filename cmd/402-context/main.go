package main

import (
	"context"
	"fmt"
	"time"
)

// 1. WithCancel：手动取消
func cancelDemo() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel() // 主动取消
	}()

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  finished")
	case <-ctx.Done():
		fmt.Println("  cancelled:", ctx.Err())
	}
}

// 2. WithTimeout：超时自动取消
func timeoutDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel() // 即使超时也要调用 cancel 释放资源

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("  finished")
	case <-ctx.Done():
		fmt.Println("  timeout:", ctx.Err())
	}
}

// 3. WithDeadline：指定截止时间
func deadlineDemo() {
	deadline := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("  finished")
	case <-ctx.Done():
		fmt.Println("  deadline exceeded:", ctx.Err())
	}
}

// 4. WithValue：传递请求级元数据
func valueDemo() {
	type contextKey string
	const reqIDKey contextKey = "requestID"

	ctx := context.WithValue(context.Background(), reqIDKey, "req-abc-123")

	// 模拟在下游函数中读取
	processRequest(ctx, reqIDKey)
}

func processRequest(ctx context.Context, key interface{}) {
	if reqID, ok := ctx.Value(key).(string); ok {
		fmt.Println("  processing request:", reqID)
	}
}

// 5. 实际场景：并发请求，任意一个完成即取消其余
func raceDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultCh := make(chan string, 2)

	go func() {
		time.Sleep(100 * time.Millisecond)
		select {
		case resultCh <- "fast-service":
		case <-ctx.Done():
		}
	}()
	go func() {
		time.Sleep(300 * time.Millisecond)
		select {
		case resultCh <- "slow-service":
		case <-ctx.Done():
		}
	}()

	select {
	case winner := <-resultCh:
		fmt.Println("  winner:", winner)
	case <-ctx.Done():
		fmt.Println("  all failed")
	}
	cancel() // 取消慢的那个
	time.Sleep(100 * time.Millisecond)
}

func main() {
	fmt.Println("=== WithCancel ===")
	cancelDemo()

	fmt.Println("=== WithTimeout ===")
	timeoutDemo()

	fmt.Println("=== WithDeadline ===")
	deadlineDemo()

	fmt.Println("=== WithValue ===")
	valueDemo()

	fmt.Println("=== Race (first win) ===")
	raceDemo()
}
