package main

import "testing"

//go:noinline
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Res int

func BenchmarkMax(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = Max(-1, i)

	}
	Res = r
}
