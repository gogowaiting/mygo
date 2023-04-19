package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// WaitGroup 内部维护着一个计数器，初始值为0
func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()

}
