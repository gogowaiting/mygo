package main

import "fmt"

func main() {
	// 基本 if/else
	if 10%2 == 0 {
		fmt.Println("10 is even")
	} else {
		fmt.Println("10 is odd")
	}

	// if 带初始化语句，变量作用域限于 if/else 块内
	if num := 42; num > 0 {
		fmt.Println("num is positive:", num)
	} else if num < 0 {
		fmt.Println("num is negative:", num)
	} else {
		fmt.Println("num is zero")
	}

	// 常见模式：检查 error
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 / 3 =", result)
	}

	// 除零错误
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	}

	// 检查 map 中 key 是否存在
	m := map[string]int{"a": 1, "b": 2}
	if v, ok := m["a"]; ok {
		fmt.Println("found key a:", v)
	}
	if _, ok := m["c"]; !ok {
		fmt.Println("key c not found")
	}
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a) / float64(b), nil
}
