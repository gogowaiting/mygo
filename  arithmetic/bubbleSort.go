
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 冒泡排序
func main() {
	arr := []int{9, 4, 5, 1, 16, 3}
	fmt.Println(bubbleSort(arr))
}