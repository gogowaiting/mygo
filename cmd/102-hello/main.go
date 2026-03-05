package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	fmt.Println("你好，Go！")

	// fmt.Printf 格式化输出
	name := "Go"
	version := 1.22
	fmt.Printf("语言: %s, 版本: %.2f\n", name, version)

	// fmt.Sprintf 返回字符串而不打印
	msg := fmt.Sprintf("I love %s", name)
	fmt.Println(msg)
}
