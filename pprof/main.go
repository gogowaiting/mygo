package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	// CPU profile
	f, err := os.Create("./cpu_profile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// Memory profile
	fm, err := os.Create("./menu_profile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fm.Close()
	pprof.WriteHeapProfile(fm)
	for i := 0; i < 100; i++ {
		fmt.Println("this is a test")
	}
}
