package main

import "fmt"

// 初始化函数，模块调用就执行, 无需单独调用
func init() {
	fmt.Println("initializing")
	fmt.Println(plus(2, 3))
}

// 函数是实现某一功能能的集合，接受参数，返回预期内容
// 函数名首字母大写表示模块外内外都可以调用，首字母小写只能在模块内部调用
func plus(a, b int) int {
	return a + b
}

// 函数支持多返回值
func one() (int, int) {
	return 1, 2
}

// 函数支持接受不定参数个数（类型一样）...type
func two(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

// 匿名函数
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 函数递归,斐波那契
func fact(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fact(n-2) + fact(n-1)

}

func main() {
	sum := plus(1, 2)
	fmt.Println(sum)

	fmt.Println(one())
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(two(nums...)) // slice... 表示打散传入

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(fact(10))
}
