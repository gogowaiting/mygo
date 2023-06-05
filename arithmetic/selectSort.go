package main

import "fmt"

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// 选择排序思路： 从第一个元素开始，依次比较相邻的两个元素，如果前一个比后一个大，则交换位置，直到最后一个元素。
func main() {
	arr := []int{9, 4, 5, 1, 16, 3}
	fmt.Println(selectSort(arr))
}
