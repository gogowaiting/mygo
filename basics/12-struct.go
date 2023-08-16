package main

import "fmt"

// struct 是字段类型的集合
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20}) // 不用指明key,直接初始化
	fmt.Println(person{name: "Bob", age: 16})
	fmt.Println(person{name: "Tom"}) // 字段缺失默认为初始化值

	fmt.Println(&person{name: "tony", age: 18}) // & 指定结构体指针

	demo := person{name: "Bob", age: 30}
	de := &demo
	de.name = "Hank"
	fmt.Println(de)
}
