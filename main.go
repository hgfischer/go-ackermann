package main

import (
	"log"
	"runtime/debug"
	"time"
)

func main() {
	debug.SetMaxStack(16000000000)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			start := time.Now()
			ans := ackermann(i, j)
			elapsed := time.Since(start)
			log.Printf("Calculated `ackermann(%d, %d) = %d` in %s", i, j, ans, elapsed)
		}
	}
}

type pair struct {
	a, b int
}

var (
	lookupTable = make(map[pair]int)
)

func ackermann(m, n int) (ans int) {
	if ans, ok := lookupTable[pair{m, n}]; ok {
		return ans
	}
	if m == 0 {
		ans = n + 1
	} else if n == 0 {
		ans = ackermann(m-1, 1)
	} else {
		ans = ackermann(m-1, ackermann(m, n-1))
	}
	lookupTable[pair{m, n}] = ans
	return
}
