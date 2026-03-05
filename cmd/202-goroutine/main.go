package main

import (
	"fmt"
	"time"
)

// goroutine 是golang中轻量级线程，由关键字go调用
// 所有goroutine 会随着main函数的结束而结束
func work(msg string) {
	fmt.Printf("this %s work \n", msg)
}

func runnig() {
	var times int

	for {
		times++
		fmt.Println("tick: ", times)
		time.Sleep(time.Second)
	}
}

func main() {
	go runnig()

	// 接受通过命令行输入，不做任何操作
	var input string
	fmt.Scanln(&input)
	// go work(input)

}
