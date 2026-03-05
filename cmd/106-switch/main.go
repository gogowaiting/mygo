package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	switch a {
	case 1:
		fmt.Println("first")
	case 2:
		fmt.Println("second")
	default:
		fmt.Println("default")
	}
	// 同一个case中使用逗号分隔多个表达式
	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("have a good rest day")
	default:
		fmt.Println(" work day")
	}
	// 无表达式，是一种if/else 实现方式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noonday")
	default:
		fmt.Println("It's after noonday")
	}

	// 类型Switch 比较了类型的而非值
	whoami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm a number")
		case float64:
			fmt.Println("I'm a float64")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whoami(true)
	whoami(123)
	whoami("good")
}
