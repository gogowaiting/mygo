package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("allocate new buffer")
			return new(bytes.Buffer)
		},
	}

	for i := 0; i < 3; i++ {
		buf := pool.Get().(*bytes.Buffer)
		buf.Reset()
		buf.WriteString("sync.Pool demo ")
		buf.WriteString(fmt.Sprint(i))
		fmt.Println(buf.String())
		pool.Put(buf)
	}
}
