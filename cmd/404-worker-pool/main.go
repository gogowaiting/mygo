package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := make(chan int)
	results := make(chan int)

	const workerCount = 3
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i+1, jobs, results, &wg)
	}

	go func() {
		defer close(jobs)
		for j := 1; j <= 10; j++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- j:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("result:", r)
	}
}

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker-%d stop by context\n", id)
			return
		case j, ok := <-jobs:
			if !ok {
				fmt.Printf("worker-%d jobs closed\n", id)
				return
			}
			time.Sleep(100 * time.Millisecond)
			results <- j * j
		}
	}
}
