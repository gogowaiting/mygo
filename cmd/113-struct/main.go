package main

import "fmt"

// struct 结构体是字段类型的集合
// struct{} 空结构体不占用内存空间，只占位
type person struct {
	name string
	age  int
}

// struct 结构体嵌套,
type student struct {
	person
	stuId  int
	school string
}

// 通过 stuct{} 实现golang中的Set
type Set[K comparable] map[K]struct{}

func (s Set[K]) Add(val K) {
	s[val] = struct{}{}
}
func (s Set[K]) Remove(val K) {
	delete(s, val)
}

func (s Set[K]) Contains(val K) bool {
	_, ok := s[val]
	return ok
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

	stu_1 := student{
		person: person{
			"小王", 18,
		},
		stuId:  110,
		school: "育才学校",
	}
	fmt.Println("王同学：", stu_1)

	students := []student{
		{
			person: person{
				"小张", 18,
			},
			stuId:  111,
			school: "育才学校",
		},
		{
			person: person{
				"小李", 16,
			},
			stuId:  112,
			school: "育才学校",
		},
	}

	for i := range students {
		fmt.Println(students[i].stuId)
	}

	set := Set[string]{}
	set.Add("a")
	set.Add("b")
	fmt.Println(set)
	fmt.Println(set.Contains("a"))

}
