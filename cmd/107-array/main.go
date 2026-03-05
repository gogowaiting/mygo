package main

import "fmt"

func main() {
	// 声明并初始化
	var a [5]int
	a[4] = 100
	fmt.Println("a:", a)
	fmt.Println("len:", len(a))

	// 字面量初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// 编译器自动推导长度
	c := [...]int{10, 20, 30}
	fmt.Println("c:", c, "len:", len(c))

	// 按索引初始化
	d := [5]int{1: 100, 3: 300}
	fmt.Println("d:", d)

	// 数组是值类型，赋值和传参会拷贝
	original := [3]int{1, 2, 3}
	copied := original
	copied[0] = 999
	fmt.Println("original:", original) // [1 2 3] 不受影响
	fmt.Println("copied:", copied)     // [999 2 3]

	// 多维数组
	var matrix [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}
	fmt.Println("matrix:", matrix)

	// 遍历数组
	for i, v := range b {
		fmt.Printf("  b[%d] = %d\n", i, v)
	}
}
