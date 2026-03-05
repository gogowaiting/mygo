package main

import "fmt"

func main() {
	// append 基本用法
	var s []string
	s = append(s, "a")
	s = append(s, "b", "c")
	fmt.Println("s:", s)

	// 观察容量增长：当 len == cap 时 append 会触发扩容
	var arr []int
	for i := 0; i < 20; i++ {
		arr = append(arr, i)
		fmt.Printf("  i=%2d  len=%2d  cap=%2d\n", i, len(arr), cap(arr))
	}
	// 可以看到容量按 ~2 倍增长: 1,2,4,8,16,32...

	// 合并两个切片
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println("merged:", c)

	// append 不影响原 slice（容量足够时共享底层数组）
	base := make([]int, 3, 10)
	base[0], base[1], base[2] = 1, 2, 3
	extended := append(base, 4)
	fmt.Println("base:", base, "len:", len(base), "cap:", cap(base))
	fmt.Println("extended:", extended, "len:", len(extended), "cap:", cap(extended))
	// extended 和 base 共享底层数组
	extended[0] = 999
	fmt.Println("after modify extended[0]:")
	fmt.Println("  base:", base) // base[0] 也被改了

	// 独立拷贝：避免共享底层数组的副作用
	original := []int{1, 2, 3}
	independent := make([]int, len(original))
	copy(independent, original)
	independent[0] = 999
	fmt.Println("original:", original)       // 不受影响
	fmt.Println("independent:", independent) // [999 2 3]
}
