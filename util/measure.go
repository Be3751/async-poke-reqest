package util

import (
	"fmt"
	"time"
)

func measurer(fnc func()) time.Duration {
	fmt.Println("Let's get started to capture all pokemons!")
	start := time.Now()
	fnc()
	end := time.Now()
	return end.Sub(start)
}

func CalcAvgRuntime(target func(), N int) time.Duration {
	var avgRuntime time.Duration
	for i := 0; i < N; i++ {
		fmt.Printf("\n%s time\n", convNumNth(i+1))
		avgRuntime += measurer(target)
	}
	avgRuntime = avgRuntime / time.Duration(int64(N))
	return avgRuntime
}
