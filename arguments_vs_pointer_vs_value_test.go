package golang_benchmark

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const loopIterator = 100

func ArgumentFuncSlice(values []TestValue) {
	counter := 0
	if values[0].Test1 > values[1].Test1 {
		counter++
	}
	if values[0].Test2 > values[1].Test2 {
		counter++
	}
	if values[0].Test3 > values[1].Test3 {
		counter++
	}
	if values[0].Test4 > values[1].Test4 {
		counter++
	}
}

// Size usage 358.4 kB 0 B time usage 1m22.9417546s
func TestArgumentsSlice(t *testing.T) {
	p := MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	p.Start()
	outputs1 := make([][]int64, 100)
	outputs2 := make([][]int64, 100)
	outputs3 := make([][]int64, 100)
	outputs4 := make([][]int64, 100)
	for i := range outputs1 {
		randomValues := make([]int64, loopIterator)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs1[i] = randomValues
	}
	for i := range outputs2 {
		randomValues := make([]int64, loopIterator)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs2[i] = randomValues
	}
	for i := range outputs3 {
		randomValues := make([]int64, loopIterator)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs3[i] = randomValues
	}
	for i := range outputs4 {
		randomValues := make([]int64, loopIterator)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs4[i] = randomValues
	}
	p.Finish()
	size := p.Size()

	p.Start()
	for i := 2; i < loopIterator; i++ {
		for o1 := range outputs1 {
			for o2 := range outputs2 {
				for o3 := range outputs3 {
					for o4 := range outputs4 {
						values := []TestValue{{
							Test1: outputs1[o1][i],
							Test2: outputs2[o2][i],
							Test3: outputs3[o3][i],
							Test4: outputs4[o4][i],
						}, {
							Test1: outputs1[o1][i-1],
							Test2: outputs2[o2][i-1],
							Test3: outputs3[o3][i-1],
							Test4: outputs4[o4][i-1],
						}, {
							Test1: outputs1[o1][i-2],
							Test2: outputs2[o2][i-2],
							Test3: outputs3[o3][i-2],
							Test4: outputs4[o4][i-2],
						}}
						ArgumentFuncSlice(values)
					}
				}
			}
		}
	}
	p.Finish()

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}

type ArgumentSingle struct {
	internal1 bool
	internal2 bool
	internal3 bool
	internal4 bool
	internal5 bool
	internal6 bool
	internal7 bool
	internal8 bool
}

func (a *ArgumentSingle) ArgumentFuncSingle(values TestValue) {
	counter := 0
	if values.Test1 > values.Test4 {
		a.internal1 = true
		counter++
	}
	if values.Test2 > values.Test1 {
		a.internal2 = true
		counter++
	}
	if values.Test3 > values.Test2 {
		a.internal3 = true
		counter++
	}
	if values.Test4 > values.Test3 {
		a.internal4 = true
		counter++
	}
}

// Size usage 800.4 MB 0 B time usage 1m23.4842087s
func TestArgumentsSingle(t *testing.T) {
	p := MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	p.Start()
	outputs1 := make([][]int64, 100)
	outputs2 := make([][]int64, 100)
	outputs3 := make([][]int64, 100)
	outputs4 := make([][]int64, 100)
	for i := range outputs1 {
		randomValues := make([]int64, loopIterator)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs1[i] = randomValues
	}
	for i := range outputs2 {
		randomValues := make([]int64, loopIterator)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs2[i] = randomValues
	}
	for i := range outputs3 {
		randomValues := make([]int64, loopIterator)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs3[i] = randomValues
	}
	for i := range outputs4 {
		randomValues := make([]int64, loopIterator)
		for y := range randomValues {
			randomValues[y] = rand.Int63n(10000)
		}
		outputs4[i] = randomValues
	}
	numberSubjects := make([]ArgumentSingle, len(outputs1)*len(outputs2)*len(outputs3)*len(outputs4))
	p.Finish()
	size := p.Size()

	p.Start()
	for i := 0; i < loopIterator; i++ {
		for y := range numberSubjects {
			index2 := y/100
			index3 := index2/100
			index4 := index3/100
			numberSubjects[y].ArgumentFuncSingle(TestValue{
				Test1: outputs1[y%100][i],
				Test2: outputs2[index2%100][i],
				Test3: outputs3[index3%100][i],
				Test4: outputs4[index4%100][i],
			})
		}
	}
	p.Finish()

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}

type PointerStruct struct {
	Output1 *Output
	Output2 *Output
	Output3 *Output
	Output4 *Output
	Counter int64
}

func (p *PointerStruct) Calc() {
	if p.Output1.GetValue() > p.Output2.GetValue() {
		p.Counter++
	}
	if p.Output2.GetValue() > p.Output3.GetValue() {
		p.Counter++
	}
	if p.Output3.GetValue() > p.Output4.GetValue() {
		p.Counter++
	}
	if p.Output4.GetValue() > p.Output1.GetValue() {
		p.Counter++
	}
}

// Size usage 4.0 GB 0 B time usage 52.9345274s
func TestPointer(t *testing.T) {
	p := MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	outputs1 := make([]Output, 100)
	outputs2 := make([]Output, 100)
	outputs3 := make([]Output, 100)
	outputs4 := make([]Output, 100)

	p.Start()
	numberSubjects := make([]PointerStruct, 0, len(outputs1)*len(outputs2)*len(outputs3)*len(outputs4))
	for o1 := range outputs1 {
		for o2 := range outputs2 {
			for o3 := range outputs3 {
				for o4 := range outputs4 {
					numberSubjects = append(numberSubjects, PointerStruct{
						Output1: &outputs1[o1],
						Output2: &outputs2[o2],
						Output3: &outputs3[o3],
						Output4: &outputs4[o4],
					})
				}
			}
		}
	}
	p.Finish()
	size := p.Size()

	fmt.Println("numberSubjects", len(numberSubjects), p.Size())

	randomValues := make([]int64, loopIterator)
	for i := range randomValues {
		randomValues[i] = rand.Int63n(10000)
	}

	p.Start()
	for i := 3; i < loopIterator; i++ {
		for o := range outputs1 {
			outputs1[o].UpdateValue(randomValues[i])
		}
		for o := range outputs2 {
			outputs2[o].UpdateValue(randomValues[i-1])
		}
		for o := range outputs3 {
			outputs3[o].UpdateValue(randomValues[i-2])
		}
		for o := range outputs4 {
			outputs4[o].UpdateValue(randomValues[i-3])
		}

		for i := range numberSubjects {
			numberSubjects[i].Calc()
		}
	}
	p.Finish()

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}

type ValueStruct struct {
	Output1 Output
	Output2 Output
	Output3 Output
	Output4 Output
	Counter int64
}

func (v *ValueStruct) Calc() {
	if v.Output1.GetValue() > v.Output2.GetValue() {
		v.Counter++
	}
	if v.Output2.GetValue() > v.Output3.GetValue() {
		v.Counter++
	}
	if v.Output3.GetValue() > v.Output4.GetValue() {
		v.Counter++
	}
	if v.Output4.GetValue() > v.Output1.GetValue() {
		v.Counter++
	}
}

// Size usage 4.0 GB 0 B time usage 1m8.9718179s
func TestValue1(t *testing.T) {
	p := MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	p.Start()
	numberSubjects := make([]ValueStruct, 100*100*100*100)
	for i := range numberSubjects {
		numberSubjects[i] = ValueStruct{
			Output1: Output{},
			Output2: Output{},
			Output3: Output{},
			Output4: Output{},
		}
	}
	p.Finish()
	size := p.Size()

	randomValues := make([]int64, loopIterator)
	for i := range randomValues {
		randomValues[i] = rand.Int63n(10000)
	}

	p.Start()
	for i := 3; i < loopIterator; i++ {
		for y := range numberSubjects {
			numberSubjects[y].Output1.UpdateValue(randomValues[i])
			numberSubjects[y].Output2.UpdateValue(randomValues[i-1])
			numberSubjects[y].Output3.UpdateValue(randomValues[i-2])
			numberSubjects[y].Output4.UpdateValue(randomValues[i-3])
			numberSubjects[y].Calc()
		}
	}
	p.Finish()

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}
