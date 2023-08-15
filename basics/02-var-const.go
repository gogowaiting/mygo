package main

import (
	"fmt"
	"math"
)

func main() {
	// var 定义变量关键字， var 名称 类型
	var a = "testString"
	var b, c = 1, 2
	var d = false
	var e float64
	f := float32(0)

	g := a + " fake"
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)

	// const 定义常量关键字， const 名称类型
	// 常量支持 字符、字符串、布尔值、数值
	const s string = "ERROR: unexpected"
	const h = 500000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
