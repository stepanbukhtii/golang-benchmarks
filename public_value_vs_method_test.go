package golang_benchmark

import (
	"fmt"
	"testing"
)

// change by public value vs change by public method
// Result: read and change by public value is faster

type PublicVariable struct {
	Result int64
	Value  TestStructure
}

func (a *PublicVariable) Handle() {
	if a.Result > a.Value.Test1 {
		a.Result = a.Result - a.Value.Test1
	} else {
		a.Result = a.Result + a.Value.Test1
	}
}

// Size usage 0 B time usage 87.740992ms
func TestPublicVariable(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var testObject PublicVariable

	p.Start()
	for i := 0; i < tenMillions; i++ {
		testObject.Value = randomValues[i]
		testObject.Handle()
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

type PublicMethod struct {
	Result int64
	value  TestStructure
}

func (a *PublicMethod) Handle() {
	if a.Result > a.value.Test1 {
		a.Result = a.Result - a.value.Test1
	} else {
		a.Result = a.Result + a.value.Test1
	}
}

func (a *PublicMethod) SetValue(value TestStructure) {
	a.value = value
}

// Size usage 0 B time usage 108.307281ms
func TestPublicMethod(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var testObject PublicMethod

	p.Start()
	for i := 0; i < tenMillions; i++ {
		testObject.SetValue(randomValues[i])
		testObject.Handle()
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

type ReadByMethod struct {
	Result int64
	Value  TestStructure
}

func (a *ReadByMethod) Handle() {
	if a.Result > a.Value.GetValue() {
		a.Result = a.Result - a.Value.GetValue()
	} else {
		a.Result = a.Result + a.Value.GetValue()
	}
}

// Size usage 0 B time usage 146.301541ms
func TestReadByMethod(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var testObject ReadByMethod

	p.Start()
	for i := 0; i < tenMillions; i++ {
		testObject.Value = randomValues[i]
		testObject.Handle()
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}
