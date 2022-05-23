package main

import (
	"fmt"
)

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

func bullerSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {

	arr := []int{9, 4, 5, 1, 16, 3}
	// quickSort
	fmt.Println(quickSort(arr))

	//bullerSort
	fmt.Println(bullerSort(arr))

}
