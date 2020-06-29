package golang_benchmark

import (
	"fmt"
	"testing"
)

type ValuePointer struct {
	Result float64
}

func (t *ValuePointer) HandlerArgumentValueLoopIndex(values []TestStructure) {
	for i := range values {
		t.Result = values[i].Test2 + values[i].Test6 + 1
	}
}

func (t *ValuePointer) HandlerArgumentValueLoopValue(values []TestStructure) {
	for _, v := range values {
		t.Result = v.Test2 + v.Test6 + 1
	}
}

func (t *ValuePointer) HandlerArgumentPointerLoopIndex(values []*TestStructure) {
	for i := range values {
		t.Result = values[i].Test2 + values[i].Test6 + 1
	}
}

func (t *ValuePointer) HandlerArgumentPointerLoopValue(values []*TestStructure) {
	for _, v := range values {
		t.Result = v.Test2 + v.Test6 + 1
	}
}

// Total heap alloc: 0 Mb 0 bytes. Time: 639.5818ms
func TestArgumentValueLoopIndex(t *testing.T) {
	customProfiler := MemTimeProfiler{}

	var target ValuePointer
	customProfiler.Start()

	values := make([]TestStructure, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = TestStructure{
			Test1: int64(i),
			Test2: float64(i) + 1,
			Test3: float64(i) + 2,
			Test4: float64(i) + 3,
			Test5: float64(i) + 4,
			Test6: float64(i) + 5,
		}
	}

	for i := 0; i < million; i++ {
		target.HandlerArgumentValueLoopIndex(values)
	}
	customProfiler.Finish()

	fmt.Println(target.Result)
}

// Total heap alloc: 0 Mb 0 bytes. Time: 2.140671s
func TestArgumentValueLoopValue(t *testing.T) {
	customProfiler := MemTimeProfiler{}

	var target ValuePointer
	customProfiler.Start()

	values := make([]TestStructure, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = TestStructure{
			Test1: int64(i),
			Test2: float64(i) + 1,
			Test3: float64(i) + 2,
			Test4: float64(i) + 3,
			Test5: float64(i) + 4,
			Test6: float64(i) + 5,
		}
	}

	for i := 0; i < million; i++ {
		target.HandlerArgumentValueLoopValue(values)
	}
	customProfiler.Finish()

	fmt.Println(target.Result)
}

// Total heap alloc: 0 Mb 48000 bytes. Time: 747.8482ms
func TestArgumentPointerLoopIndex(t *testing.T) {
	customProfiler := MemTimeProfiler{}

	var target ValuePointer
	customProfiler.Start()

	values := make([]*TestStructure, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = &TestStructure{
			Test1: int64(i),
			Test2: float64(i) + 1,
			Test3: float64(i) + 2,
			Test4: float64(i) + 3,
			Test5: float64(i) + 4,
			Test6: float64(i) + 5,
		}
	}

	for i := 0; i < million; i++ {
		target.HandlerArgumentPointerLoopIndex(values)
	}
	customProfiler.Finish()

	fmt.Println(target.Result)
}

// Total heap alloc: 0 Mb 48000 bytes. Time: 746.6788ms
func TestArgumentPointerLoopValue(t *testing.T) {
	customProfiler := MemTimeProfiler{}

	var target ValuePointer
	customProfiler.Start()

	values := make([]*TestStructure, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = &TestStructure{
			Test1: int64(i),
			Test2: float64(i) + 1,
			Test3: float64(i) + 2,
			Test4: float64(i) + 3,
			Test5: float64(i) + 4,
			Test6: float64(i) + 5,
		}
	}

	for i := 0; i < million; i++ {
		target.HandlerArgumentPointerLoopValue(values)
	}
	customProfiler.Finish()

	fmt.Println(target.Result)
}
