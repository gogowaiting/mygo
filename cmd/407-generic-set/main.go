package main

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Values() []T {
	out := make([]T, 0, len(s))
	for v := range s {
		out = append(out, v)
	}
	return out
}

func main() {
	nums := NewSet(1, 2, 2, 3)
	nums.Add(4)
	fmt.Println("has 3:", nums.Has(3))
	fmt.Println("values:", nums.Values())

	langs := NewSet("go", "java", "go")
	fmt.Println("lang values:", langs.Values())
}
