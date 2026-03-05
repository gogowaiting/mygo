package main

import (
	"fmt"
	// "math"
)

// 接口interface 是方法申明的一个集合
// 任何对象实现了接口中的所有方法则表明该对象实现了该接口
// 作为一种数据类型，实现了该接口的对象都可以给对应接口的变量赋值
// 一个类型的指针类型实现了该接口，那么它的值类型也会隐式地实现该接口
// interface{} 0个方法表示空接口，任何类型都实现了空接口，空接口变量可以赋任何类型的值

type gepmetry interface {
	// 定义接口，里面需要实现两个方法
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

// 实现一个接口需要实现接口中定义的所有方法
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return r.width*2 + r.height*2
}

func measure(g gepmetry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

type animal interface {
	Run(name string) string
}

type cat struct{}

func (c cat) Run(name string) string {
	return fmt.Sprintf("my name is %s , i can run", name)
}

type Fixble interface{}

func main() {
	r := rect{width: 3, height: 4}
	measure(r)

	var fix Fixble
	fix = "18"
	fmt.Println(fix)
	fix = "this is a emtpy fixble interface"
	fmt.Println(fix)

	var dog animal
	dog = cat{}
	fmt.Println(dog.Run("dog"))

}
