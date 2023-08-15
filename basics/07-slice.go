package main

import "fmt"

func main() {
	// slice 切片
	s := make([]string, 3) // 初始化为非零长度空切片，长度为3的切片
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(s[2:5])
	fmt.Println(s[:5])
	fmt.Println(s[2:])

	good := []string{"g", "o"}
	fmt.Println(good)

}
