package main

// init 函数并不一定是最先执行的，比如定义了一个变量，那么这个变量的初始化就会在 init 函数之前执行。
import "fmt"

const num = 123

var a = func() {
	fmt.Println("in variable a func")
}

func init() {
	fmt.Println("in init func")
}

type BB struct {
	CC
}

type CC struct {
	dd int
}

func main() {
	fmt.Println(&a)
	fmt.Println("const:", num)
	fmt.Println("in main progress")

	var b BB
	b.dd = 123
	fmt.Println(b)

	b.CC.dd = 456
	fmt.Println(b)
	// 预期110，实际108；在go中，以0开头数字代表8进制，以0x开头代表16进制
	// 为了可读性使用8进制，最好使用0o开头表示
	sum := 100 + 010
	fmt.Println("sum: ", sum)
}
