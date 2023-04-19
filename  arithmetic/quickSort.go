package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	low := make([]int, 0, 0)
	hight := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, pivot)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			low = append(low, arr[i])
		} else if arr[i] > pivot {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = quickSort(low), quickSort(hight)
	newArr := append(append(low, mid...), hight...)
	return newArr
}

//快排思路：通过一次循环将待排序记录分割成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。
func main() {
	arr := []int{9, 4, 5, 1, 16, 3}
	fmt.Println(quickSort(arr))
}
