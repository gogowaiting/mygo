package main

import "fmt"

func main() {
	// 创建一个空map
	m := make(map[string]int)
	m["k1"] = 123
	m["k2"] = 456
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println(len(m))

	delete(m, "k1")
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]string{"k1": "one", "k2": "two"}
	fmt.Println("n:", n)

}
