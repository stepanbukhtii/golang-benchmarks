package golang_benchmark

import (
	"fmt"
	"testing"
)

// arguments is slice of struct vs argument is many struct
// Result: slice is faster ???

type ArgsSlice struct {
	Result int64
}

func (a *ArgsSlice) Handle(values []TestStructure) {
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

// Size usage 0 B time usage 118.091721ms
func TestSliceArgs(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var testObject ArgsSlice
	argumentsSlice := []TestStructure{randomValues[0], randomValues[1]}

	p.Start()
	for _, v := range randomValues {
		argumentsSlice[1] = v
		testObject.Handle(argumentsSlice)
		argumentsSlice[0] = v
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

type ManyArgs struct {
	Result int64
}

func (a *ManyArgs) Handle(value1, value2 TestStructure) {
	if a.Result > value1.Test1 {
		a.Result = a.Result - value1.Test1
	} else {
		a.Result = a.Result + value1.Test1
	}
	if a.Result > value2.Test1 {
		a.Result = a.Result - value2.Test1
	} else {
		a.Result = a.Result + value2.Test1
	}
}

// Size usage 0 B time usage 148.651818ms
func TestManyArgs(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var testObject ManyArgs
	firstArgument := randomValues[0]

	p.Start()
	for _, v := range randomValues {
		testObject.Handle(firstArgument, v)
		firstArgument = v
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}
