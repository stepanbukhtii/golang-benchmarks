package golang_benchmark

import (
	"fmt"
	"sync"
	"testing"
)

// 1 - time 33.805399121s
// 2 - time 17.076792862s
// 4 - time 9.375797496s
// 8 - time 8.781864343s
// 20 - time 8.928432821s
// 40 - time 8.793372772s
func TestNumberOfGoroutines(t *testing.T) {
	var p MemTimeProfiler
	var wg sync.WaitGroup
	numberOfGoroutines := 15
	mustCount := 120

	p.Start()
	for i := 0; i < mustCount/numberOfGoroutines; i++ {
		wg.Add(numberOfGoroutines)
		for y := 0; y < numberOfGoroutines; y++ {
			go heavyWork(&wg)
		}
		wg.Wait()
	}
	p.Finish()

	fmt.Println("time", p.Time())
}

func heavyWork(wg *sync.WaitGroup) {
	x := 5
	for i := 0; i < million*1000; i++ {
		if i > x {
			x = x + i + i
			x = x - i - i
		}

		x = x + i + i
		x = x - i - i
	}
	wg.Done()
}
