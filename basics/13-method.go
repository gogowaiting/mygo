package main

import "fmt"

type rect struct {
	height, width int
}

func (r *rect) area() int {
	return r.height * r.width
}
func (r rect) perim() int {
	return r.height*2 + r.width*2
}

func main() {
	// golang 中方法是实现结构体的函数集合
	r := rect{height: 10, width: 10}
	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())
	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())

}
