package waitgroup

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

// WaitGroup 内部维护着一个计数器，初始值为0
func Hello(i int) {
	defer wg.Done()
	fmt.Println("Hello", i)
}

func TestHello(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Hello(i)
	}
	wg.Wait()

}
