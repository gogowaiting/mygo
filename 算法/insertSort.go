package main

import "fmt"

func insertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

// 插入排序思路：先选择一个基准数，然后一次比较其余数字与基准数字大小，将小于基准数的数字放在基准数的左边，将大于基准数的数字放在基准数的右边，直到排序完成。
func main() {
	arr := []int{9, 4, 5, 1, 16, 3}
	fmt.Println(insertSort(arr))
}
