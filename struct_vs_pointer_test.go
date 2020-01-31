package golang_benchmark

import (
	"testing"
)

type StructValue struct {
	Result TestStructure
}

type StructPointer struct {
	Result *TestStructure
}

// Total heap alloc: 45 Mb 48012128 bytes. Time: 26.180074ms
func TestStructValue(t *testing.T) {
	customProfiler := CustomProfiler{}

	counter := float64(0)

	customProfiler.Start()
	values := make([]StructValue, million)
	for i := 0; i < million; i++ {
		target := StructValue{
			Result: TestStructure{
				Test1: 1,
				Test2: 2,
				Test3: 3,
				Test4: 4,
				Test5: 5,
				Test6: 6,
			},
		}
		values[i] = target
		counter += target.Result.Test2
	}
	customProfiler.Finish()
}

// Total heap alloc: 53 Mb 56012288 bytes. Time: 64.915662ms
func TestStructPointer(t *testing.T) {
	customProfiler := CustomProfiler{}

	counter := float64(0)

	customProfiler.Start()
	values := make([]StructPointer, million)
	for i := 0; i < million; i++ {
		target := StructPointer{
			Result: &TestStructure{
				Test1: 1,
				Test2: 2,
				Test3: 3,
				Test4: 4,
				Test5: 5,
				Test6: 6,
			},
		}
		values[i] = target
		counter += target.Result.Test2
	}
	customProfiler.Finish()
}
