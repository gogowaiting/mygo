package main

import "fmt"

func main() {
	// if 条件语句，用于条件处理多情况场景，可搭配 else ，else if 一起使用
	if 10%2 == 0 {
		fmt.Println("10 is even and 10 is divisible by 2")
	} else {
		fmt.Println("10 is odd")
	}

	var num int
	fmt.Println("input you age: ")
	fmt.Scanln(&num)
	if num < 18 && num > 0 {
		fmt.Println(num, "you is nonage")
	} else if num >= 18 {
		fmt.Println(num, "you is ault")
	} else {
		fmt.Println("you input age error,mabey you are a baby")
	}

}
