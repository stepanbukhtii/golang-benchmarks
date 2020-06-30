package golang_benchmark

import (
	"fmt"
	"testing"
)

// arguments is slice of values vs arguments is slice of pointers
// Result: about the same

type SliceOfValues struct {
	Result int64
}

func (a *SliceOfValues) Handle(values []TestStructure) {
	if a.Result > values[0].Test1 {
		a.Result = a.Result - values[0].Test1
	} else {
		a.Result = a.Result + values[0].Test1
	}
	if a.Result > values[1].Test1 {
		a.Result = a.Result - values[1].Test1
	} else {
		a.Result = a.Result + values[1].Test1
	}
}

// Size usage 0 B time usage 76.093157ms
func TestSliceOfValues(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var testObject SliceOfValues
	p.Start()
	for i := 0; i < tenMillions-19; i++ {
		testObject.Handle(randomValues[i : i+20])
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

type SliceOfPointers struct {
	Result int64
}

func (a *SliceOfPointers) Handle(values []*TestStructure) {
	if a.Result > values[0].Test1 {
		a.Result = a.Result - values[0].Test1
	} else {
		a.Result = a.Result + values[0].Test1
	}
	if a.Result > values[1].Test1 {
		a.Result = a.Result - values[1].Test1
	} else {
		a.Result = a.Result + values[1].Test1
	}
}

// Size usage 0 B time usage 82.431759ms
func TestSliceOfPointers(t *testing.T) {
	p := MemTimeProfiler{}

	randomValuesValues := GenerateRandomTestStructures(tenMillions)
	randomValues := make([]*TestStructure, 0, tenMillions)
	for i := range randomValuesValues {
		randomValues = append(randomValues, &randomValuesValues[i])
	}

	var testObject SliceOfPointers
	p.Start()
	for i := 0; i < tenMillions-19; i++ {
		testObject.Handle(randomValues[i : i+20])
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}
