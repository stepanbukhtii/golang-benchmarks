package golang_benchmark

import (
	"fmt"
	"testing"
)

// arguments pointer to struct vs arguments value struct
// Result: about the same
// Passing by value often is cheaper but if variable is a large struct and performance is an issue, it’s preferable to pass variable by pointer.

type ArgsValue struct {
	Result int64
}

func (a *ArgsValue) Handle(value TestStructure) {
	if a.Result > value.Test1 {
		a.Result = a.Result - value.Test1
		return
	}
	a.Result = a.Result + value.Test1
}

// Size usage 0 B time usage 81.744033ms
func TestArgumentValue(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var testObject ArgsValue
	p.Start()
	for i := range randomValues {
		testObject.Handle(randomValues[i])
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

type ArgsPointer struct {
	Result int64
}

func (a *ArgsPointer) Handle(value *TestStructure) {
	if a.Result > value.Test1 {
		a.Result = a.Result - value.Test1
		return
	}
	a.Result = a.Result + value.Test1
}

// Size usage 0 B time usage 92.327654ms
func TestArgumentPointer(t *testing.T) {
	p := MemTimeProfiler{}
	randomValuesValues := GenerateRandomTestStructures(tenMillions)
	randomValues := make([]*TestStructure, 0, tenMillions)
	for i := range randomValuesValues {
		randomValues = append(randomValues, &randomValuesValues[i])
	}

	var testObject ArgsPointer
	p.Start()
	for i := range randomValues {
		testObject.Handle(randomValues[i])
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}
