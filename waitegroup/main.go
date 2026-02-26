package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go func ()  {
		defer wg.Done()

		fmt.Println("hello from goroutine")
		time.Sleep(100 * time.Millisecond)
	}()

	wg.Wait()
}