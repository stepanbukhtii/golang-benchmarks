package golang_benchmark

import (
	"fmt"
	"testing"
)

// change public value vs get from argument
// Result:

type ByPublicValue struct {
	Value  TestStructure
	Result int64
}

func (a *ByPublicValue) Handle() {
	if a.Value.Test1 > a.Result {
		a.Result = a.Result - a.Value.Test1
		return
	}
	a.Result = a.Result + a.Value.Test1
}

// Size usage 0 B time usage 10.367554ms
func TestByPublicValue(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(10*million)

	var testObject ByPublicValue

	p.Start()
	for _, v := range randomValues {
		testObject.Value = v
		testObject.Handle()
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

type ByArgumentValue struct {
	Result int64
}

func (a *ByArgumentValue) Handle(value TestStructure) {
	if value.Test1 > a.Result {
		a.Result = a.Result - value.Test1
		return
	}
	a.Result = a.Result + value.Test1
}

// Size usage 0 B time usage 10.342643ms
func TestByArgumentValue(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(10*million)

	var testObject ByArgumentValue

	p.Start()
	for _, v := range randomValues {
		testObject.Handle(v)
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}
