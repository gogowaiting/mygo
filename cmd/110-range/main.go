package main

import "fmt"

func main() {
	// range 可以遍历各种数据结构中的元素

	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	// 遍历切片，index ，values
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}
	// 遍历map， key, value

	m := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 只通过key遍历

	for k := range m {
		fmt.Println("key", k)
	}
	// 遍历字符串
	for i, c := range "google" {
		fmt.Println(i, c)
	}
	// range 直接修改返回的数据，不会生效，返回数据并不是原始数据
	type account struct {
		price float32
	}
	accounts := []account{
		{price: 100},
		{price: 200},
	}

	for _, a := range accounts {
		a.price += 10

	}

	fmt.Println(accounts)
	for i := range accounts {
		accounts[i].price += 10
	}
	fmt.Println(accounts)
}
