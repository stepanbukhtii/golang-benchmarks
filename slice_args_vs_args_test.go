package golang_benchmark

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func SliceArgs(in []TestStructure) float64 {
	if in[0].Test2 > in[1].Test2 {
		return in[0].Test2 - in[1].Test2
	}
	return in[0].Test2 + in[1].Test2
}

func ManyArgs(firstIn, secondIn TestStructure) float64 {
	if firstIn.Test2 > secondIn.Test2 {
		return firstIn.Test2 - secondIn.Test2
	}
	return firstIn.Test2 + secondIn.Test2
}

// Total heap alloc: 0 Mb 0 bytes. Time: 11.37786ms
func TestSliceArgs(t *testing.T) {
	customProfiler := CustomProfiler{}

	rand.Seed(time.Now().UnixNano())
	randomNumber := float64(rand.Intn(5-1) + 1)
	var counter int64

	sliceStucts := []TestStructure{
		{Test1: 1, Test2: randomNumber, Test3: 1, Test4: 1, Test5: 1, Test6: 1},
		{Test1: 2, Test2: randomNumber, Test3: 2, Test4: 2, Test5: 2, Test6: 2},
	}

	customProfiler.Start()
	for i := 0; i < million; i++ {
		sliceStucts[0] = sliceStucts[1]
		sliceStucts[1] = TestStructure{
			Test1: int64(i),
			Test2: randomNumber + float64(i),
			Test3: float64(i),
			Test4: float64(i),
			Test5: float64(i),
			Test6: float64(i),
		}
		c := SliceArgs(sliceStucts)
		counter += int64(c)
	}
	customProfiler.Finish()

	fmt.Println(counter)
}

// Total heap alloc: 0 Mb 0 bytes. Time: 11.227266ms
func TestManyArgs(t *testing.T) {
	customProfiler := CustomProfiler{}

	rand.Seed(time.Now().UnixNano())
	randomNumber := float64(rand.Intn(5-1) + 1)
	var counter int64

	firstStruct := TestStructure{Test1: 1, Test2: randomNumber, Test3: 1, Test4: 1, Test5: 1, Test6: 1}
	secondStruct := TestStructure{Test1: 2, Test2: randomNumber, Test3: 2, Test4: 2, Test5: 2, Test6: 2}

	customProfiler.Start()
	for i := 0; i < million; i++ {
		firstStruct = secondStruct
		secondStruct = TestStructure{
			Test1: int64(i),
			Test2: randomNumber + float64(i),
			Test3: float64(i),
			Test4: float64(i),
			Test5: float64(i),
			Test6: float64(i),
		}
		c := ManyArgs(firstStruct, secondStruct)
		counter += int64(c)
	}
	customProfiler.Finish()

	fmt.Println(counter)
}
