package main

import (
	"fmt"
	"time"
)

func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-conn:
		timer.Stop()
		a := <-conn
		fmt.Println(a)
		return true
	case <-timer.C:
		fmt.Println("WaitChannel timeout")
		return false

	}
}

func AfterFuncDemo() {
	fmt.Println("AfterFuncDemo start:", time.Now())
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("AfterFuncDemo end", time.Now())
	})

	time.Sleep(2 * time.Second)

}

func main() {
	var newStr = make(chan string)
	test := "hello"
	for i := range test {
		fmt.Println(i, test[i])
	}
	WaitChannel(newStr)

	AfterFuncDemo()
}
