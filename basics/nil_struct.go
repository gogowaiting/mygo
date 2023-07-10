package main

import (
	"fmt"
)

// 通过map+空struct{} 模拟Set
type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}
func (s Set) Remove(key string) {
	delete(s, key)
}

// 空struct{}不发送数据的channel,控制协程执行或者并发
func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something")
	close(ch)
}

func main() {
	s := make(Set)
	s.Add("Tom")
	s.Add("Sam")
	fmt.Println(s.Has("Tom"))
	s.Add("Sam")
	fmt.Println(s)

	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}

}
