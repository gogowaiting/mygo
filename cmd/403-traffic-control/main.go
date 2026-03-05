package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

type TrafficControl struct {
	source     []int
	queryCount uint32
	base       int
	ratio      int
}

/*
确定比例，并根据比例得到一个基数base，例如比例是3:7，那么基数就是10；
生成长度为基数base的数组source，并填充数据0、1、2、3、4、5...；
打乱数组source中元素顺序；
创建全局计数器queryCount，每次有请求时加1（确保原子性）；
计算计数器queryCount与base取余后的值rate，并得到数组中对应位置的值source[rate]；
判断source[rate]落在哪个区间。
*/
func NewTrafficControl(base int, ratio int) *TrafficControl {
	source := make([]int, base)
	for i := 0; i < base; i++ {
		source[i] = i
	}
	rand.Shuffle(base, func(i, j int) {
		source[i], source[j] = source[j], source[i]
	})

	return &TrafficControl{
		source:     source,
		queryCount: 0,
		base:       base,
		ratio:      ratio,
	}
}

func (t *TrafficControl) Allow() bool {
	rate := t.source[int(atomic.AddUint32(&t.queryCount, 1))%t.base]
	if rate < t.ratio {
		return true
	} else {
		return false
	}
}

func main() {
	trafficCtl := NewTrafficControl(2, 1)
	cnt := 1000
	serviceAQueryCnt := 0
	serviceBQueryCnt := 0
	for cnt > 0 {
		if trafficCtl.Allow() {
			serviceAQueryCnt++
		} else {
			serviceBQueryCnt++
		}
		cnt--
	}
	fmt.Printf("service A query conut:%v, service B query conut:%v\n", serviceAQueryCnt, serviceBQueryCnt)
}
