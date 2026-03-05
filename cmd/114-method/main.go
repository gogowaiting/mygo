package main

import (
	"fmt"
	"math"
)

// 基本结构体
type rect struct {
	height, width int
}

// 值接收者：不修改原始数据
func (r rect) perim() int {
	return 2 * (r.height + r.width)
}

// 指针接收者：可以修改原始数据
func (r *rect) scale(factor int) {
	r.height *= factor
	r.width *= factor
}

func (r rect) area() int {
	return r.height * r.width
}

// 接口定义
type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

// 接收者方法实现接口，打印任意 shape
func printShape(s shape) {
	fmt.Printf("  %T: area=%.2f, perimeter=%.2f\n", s, s.area(), s.perimeter())
}

func main() {
	// 值接收者 vs 指针接收者
	r := rect{height: 10, width: 5}
	fmt.Println("rect:", r)
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	r.scale(2)
	fmt.Println("after scale(2):", r)
	fmt.Println("area:", r.area())

	// 指针也可以调用值接收者方法
	rp := &r
	fmt.Println("pointer perim:", rp.perim())

	// 方法值：绑定到具体实例
	areaFunc := r.area
	fmt.Println("method value:", areaFunc())

	// 通过接口实现多态
	shapes := []shape{
		circle{radius: 5},
		square{side: 4},
		circle{radius: 3},
	}
	for _, s := range shapes {
		printShape(s)
	}
}
