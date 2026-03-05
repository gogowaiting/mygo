package main

import (
	"fmt"
	"runtime"
	"testing"
)

func cpuIntensiveTask() int {
	sum := 0
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			sum += i * j
		}

	}
	return sum
}

func BenchmarkCpuIntensiveTask(b *testing.B) {
	oldProc := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(oldProc)

	b.Run("GOMAXPROCS=1", func(b *testing.B) {
		runtime.GOMAXPROCS(1)
		for i := 0; i < b.N; i++ {
			cpuIntensiveTask()
		}
	})

	b.Run("GOMAXPROCS=NumCPU", func(b *testing.B) {
		runtime.GOMAXPROCS(runtime.NumCPU())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cpuIntensiveTask()
		}
	})
}

func main() {
	fmt.Println("cpuIntensiveTask:")
	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{},
		[]testing.InternalBenchmark{
			{"BenchmarkCpuIntensiveTask", BenchmarkCpuIntensiveTask},
		},
		nil)

}
