package main

import "fmt"

func main() {
	// slice 切片
	// 切片默认零值为nil
	// 切片的本质（运行时结构）
	// type slice struct {
	//     array unsafe.Pointer  // 指向底层数组
	//     len   int             // 长度
	//     cap   int             // 容量
	// }

	s := make([]string, 3) // 初始化为非零长度空切片，长度为3的切片
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(s[2:5])
	fmt.Println(s[:5])
	fmt.Println(s[2:])

	good := []string{"g", "o"}
	fmt.Println(good)

	test := []int{1, 2, 3, 4, 5}
	b := test[1:3] //切片与原数组共享底层数组
	fmt.Println(b)
	b[0] = 100 // 切片是对底层数组的引用，修改切片会修改底层数组，因此test[1]也被修改了
	fmt.Println(test)

	// 切片扩容
	// 切片扩容机制（Go 1.18+）：
	// 新容量 < 256：新容量 = 旧容量 × 2
	// 新容量 ≥ 256：新容量 = 旧容量 × 1.25 + 192（过渡阶段）
	// 最终容量会进行内存对齐

	test = make([]int, 3, 5)
	test = append(test, 1)
	fmt.Println(test)
	test = append(test, 2, 3) // 触发底层数组扩容
	fmt.Println(test)

}
