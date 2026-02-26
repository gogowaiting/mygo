package main

import (
	"fmt"
)

// func unique(nums []int) []int {
// 	seen := make(map[int]struct{})
// 	res := make([]int, 0, len(nums))

// 	for _, v := range nums {
// 		if _, ok := seen[v]; !ok {
// 			seen[v] = struct{}{}
// 			res = append(res, v)
// 		}
// 	}

// 	return res
// }

func Unique[T comparable](data []T) []T {
	seen := make(map[T]struct{}, len(data))
	res := make([]T, 0, len(data))

	for _, v := range data {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			res = append(res, v)
		}
	}

	return res
}

func main() {
	nums := []int{1, 2, 2, 4, 4, 6, 7, 4, 9, 10}
	new_nums := Unique(nums)
	fmt.Println(nums)
	fmt.Println(new_nums)

	new_string := Unique([]string{"a", "b", "b", "c", "c", "d", "e", "f", "g"})
	fmt.Println(new_string)
}
