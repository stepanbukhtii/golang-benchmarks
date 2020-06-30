package golang_benchmark

import (
	"math/rand"
	"time"
)

const million = 1000000
const tenMillions = 10 * million

type TestStructure struct {
	Test1  int64
	Test2  int64
	Test3  int64
	Test4  int64
	Test5  int64
	Test6  int64
	Test7  int64
	Test8  int64
	Test9  int64
	Test10 int64
}

type TestValue struct {
	Test1 int64
	Test2 int64
	Test3 int64
	Test4 int64
}

type Output struct {
	Value  int64
	value2 int64
}

func (o *Output) GetValue() int64 {
	return o.Value
}

func (o *Output) UpdateValue(value int64) {
	o.Value = value
}

func GenerateRandomTestStructures(quantity int) []TestStructure {
	rand.Seed(time.Now().UnixNano())

	randomValues := make([]TestStructure, quantity)
	for y := range randomValues {
		randomValues[y] = TestStructure{
			Test1:  rand.Int63n(10000),
			Test2:  rand.Int63n(10000),
			Test3:  rand.Int63n(10000),
			Test4:  rand.Int63n(10000),
			Test5:  rand.Int63n(10000),
			Test6:  rand.Int63n(10000),
			Test7:  rand.Int63n(10000),
			Test8:  rand.Int63n(10000),
			Test9:  rand.Int63n(10000),
			Test10: rand.Int63n(10000),
		}
	}

	return randomValues
}
