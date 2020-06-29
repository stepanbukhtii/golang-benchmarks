package golang_benchmark

const million = 1000000

type TestStructure struct {
	Test1 int64
	Test2 float64
	Test3 float64
	Test4 float64
	Test5 float64
	Test6 float64
}

type TestValue struct {
	Test1 int64
	Test2 int64
	Test3 int64
	Test4 int64
}

type Output struct {
	Value int64
	value2 int64
}

func (o *Output) GetValue() int64 {
	return o.Value
}

func (o *Output) UpdateValue(value int64) {
	o.Value = value
}
