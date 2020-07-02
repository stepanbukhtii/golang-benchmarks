package golang_benchmark

import (
	"fmt"
	"testing"
)

// change in million values in 200 struct by
// public value vs pass value through argument vs change value by pointer
// Result: change by pointer is faster

const countStructs = 2000

type ByPublicValue struct {
	Value  TestStructure
	Result int64
}

func (a *ByPublicValue) Handle() {
	if a.Result > a.Value.Test1 {
		a.Result = a.Result - a.Value.Test1
		return
	}
	a.Result = a.Result + a.Value.Test1
}

// Size usage 0 B time usage 7.425657755s
func TestByPublicValue(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(million)

	testObjects := make([]ByPublicValue, countStructs)

	p.Start()
	for _, v := range randomValues {
		for i := range testObjects {
			testObjects[i].Value = v
			testObjects[i].Handle()
		}
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

type ByArgumentValue struct {
	Result int64
}

func (a *ByArgumentValue) Handle(value TestStructure) {
	if a.Result > value.Test1 {
		a.Result = a.Result - value.Test1
		return
	}
	a.Result = a.Result + value.Test1
}

// Size usage 0 B time usage 5.954588607s
func TestByArgumentValue(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(million)

	testObjects := make([]ByArgumentValue, countStructs)

	p.Start()
	for _, v := range randomValues {
		for i := range testObjects {
			testObjects[i].Handle(v)
		}
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

type ByPointerValue struct {
	Value  *TestStructure
	Result int64
}

func (a *ByPointerValue) Handle() {
	if a.Result > a.Value.Test1 {
		a.Result = a.Result - a.Value.Test1
		return
	}
	a.Result = a.Result + a.Value.Test1
}

// Size usage 0 B time usage 1.754330667s
func TestByPointerValue(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(million)

	testObjects := make([]ByPointerValue, countStructs)
	var value TestStructure
	for i := range testObjects {
		testObjects[i].Value = &value
	}

	p.Start()
	for _, v := range randomValues {
		value = v
		for i := range testObjects {
			testObjects[i].Handle()
		}
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}
