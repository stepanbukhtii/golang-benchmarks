package golang_benchmark

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const loopIterator = 100

func ArgumentFuncSlice(values []TestValue) int64 {
	counter := int64(0)
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
	return counter
}

// Size usage 358.4 kB 0 B time usage 1m28.237713695s // first loopIterator
// Size usage 358.4 kB 0 B time usage 1m0.479882931s // first loop outputs and use single argumentValue
func TestArgumentsSlice(t *testing.T) {
	p := MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	p.Start()
	outputs1 := make([][]int64, loopIterator)
	outputs2 := make([][]int64, loopIterator)
	outputs3 := make([][]int64, loopIterator)
	outputs4 := make([][]int64, loopIterator)
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

	var counter int64
	argumentValue := make([]TestValue, 3)

	p.Start()
	//for i := 2; i < loopIterator; i++ {
	//	for o1 := range outputs1 {
	//		for o2 := range outputs2 {
	//			for o3 := range outputs3 {
	//				for o4 := range outputs4 {
	//					values := []TestValue{{
	//						Test1: outputs1[o1][i],
	//						Test2: outputs2[o2][i],
	//						Test3: outputs3[o3][i],
	//						Test4: outputs4[o4][i],
	//					}, {
	//						Test1: outputs1[o1][i-1],
	//						Test2: outputs2[o2][i-1],
	//						Test3: outputs3[o3][i-1],
	//						Test4: outputs4[o4][i-1],
	//					}, {
	//						Test1: outputs1[o1][i-2],
	//						Test2: outputs2[o2][i-2],
	//						Test3: outputs3[o3][i-2],
	//						Test4: outputs4[o4][i-2],
	//					}}
	//					r := ArgumentFuncSlice(values)
	//					counter = counter + r
	//				}
	//			}
	//		}
	//	}
	//}
	for _, v1 := range outputs1 {
		for _, v2 := range outputs2 {
			for _, v3 := range outputs3 {
				for _, v4 := range outputs4 {
					for i := 0; i < loopIterator; i++ {
						argumentValue[2] = TestValue{
							Test1: v1[i],
							Test2: v2[i],
							Test3: v3[i],
							Test4: v4[i],
						}
						r := ArgumentFuncSlice(argumentValue)
						argumentValue[0] = argumentValue[1]
						argumentValue[1] = argumentValue[2]
						counter = counter + r
					}
				}
			}
		}
	}
	p.Finish()

	fmt.Println("counter", counter)

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}

type ArgumentSingle struct {
	Counter int64
}

func (a *ArgumentSingle) ArgumentFuncSingle(value TestValue) {
	if value.Test1 > value.Test4 {
		a.Counter++
	}
	if value.Test2 > value.Test1 {
		a.Counter++
	}
	if value.Test3 > value.Test2 {
		a.Counter++
	}
	if value.Test4 > value.Test3 {
		a.Counter++
	}
}

// Size usage 800.4 MB 0 B time usage 42.749906321s first loopIterator
// Size usage 800.4 MB 0 B time usage 1m56.730182636s first loop outputs
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
		counter := 0
		for o1 := range outputs1 {
			for o2 := range outputs2 {
				for o3 := range outputs3 {
					for o4 := range outputs4 {
						numberSubjects[counter].ArgumentFuncSingle(TestValue{
							Test1: outputs1[o1][i],
							Test2: outputs2[o2][i],
							Test3: outputs3[o3][i],
							Test4: outputs4[o4][i],
						})
						counter++
					}
				}
			}
		}
	}
	//counter := 0
	//for _, v1 := range outputs1 {
	//	for _, v2 := range outputs2 {
	//		for _, v3 := range outputs3 {
	//			for _, v4 := range outputs4 {
	//				for i := 0; i < loopIterator; i++ {
	//					numberSubjects[counter].ArgumentFuncSingle(TestValue{
	//						Test1: v1[i],
	//						Test2: v2[i],
	//						Test3: v3[i],
	//						Test4: v4[i],
	//					})
	//				}
	//				counter++
	//			}
	//		}
	//	}
	//}
	p.Finish()

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}

type ManyArguments struct {
	Counter int64
}

func (a *ManyArguments) ArgumentFuncSingle(value TestValue, value2 TestValue, value3 TestValue) {
	if value.Test1 > value2.Test1 {
		a.Counter++
	}
	if value.Test2 > value2.Test2 {
		a.Counter++
	}
	if value.Test3 > value2.Test3 {
		a.Counter++
	}
	if value.Test4 > value3.Test4 {
		a.Counter++
	}
}

// Size usage 800.4 MB 0 B time usage 45.117197977s first loopIterator
// Size usage 800.4 MB 0 B time usage 1m47.337565368s first loop outputs
func TestManyArguments(t *testing.T) {
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
	numberSubjects := make([]ManyArguments, len(outputs1)*len(outputs2)*len(outputs3)*len(outputs4))
	p.Finish()
	size := p.Size()

	p.Start()
	//for i := 2; i < loopIterator; i++ {
	//	counter := 0
	//	for o1 := range outputs1 {
	//		for o2 := range outputs2 {
	//			for o3 := range outputs3 {
	//				for o4 := range outputs4 {
	//					numberSubjects[counter].ArgumentFuncSingle(TestValue{
	//						Test1: outputs1[o1][i],
	//						Test2: outputs2[o2][i],
	//						Test3: outputs3[o3][i],
	//						Test4: outputs4[o4][i],
	//					}, TestValue{
	//						Test1: outputs1[o1][i-1],
	//						Test2: outputs2[o2][i-1],
	//						Test3: outputs3[o3][i-1],
	//						Test4: outputs4[o4][i-1],
	//					}, TestValue{
	//						Test1: outputs1[o1][i-2],
	//						Test2: outputs2[o2][i-2],
	//						Test3: outputs3[o3][i-2],
	//						Test4: outputs4[o4][i-2],
	//					})
	//					counter++
	//				}
	//			}
	//		}
	//	}
	//}
	counter := 0
	for _, v1 := range outputs1 {
		for _, v2 := range outputs2 {
			for _, v3 := range outputs3 {
				for _, v4 := range outputs4 {
					for i := 2; i < loopIterator; i++ {
						numberSubjects[counter].ArgumentFuncSingle(TestValue{
							Test1: v1[i],
							Test2: v2[i],
							Test3: v3[i],
							Test4: v4[i],
						}, TestValue{
							Test1: v1[i-1],
							Test2: v2[i-1],
							Test3: v3[i-1],
							Test4: v4[i-1],
						}, TestValue{
							Test1: v1[i-2],
							Test2: v2[i-2],
							Test3: v3[i-2],
							Test4: v4[i-2],
						})
					}
					counter++
				}
			}
		}
	}
	p.Finish()

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}

type PointerStructByMethod struct {
	Output1 *Output
	Output2 *Output
	Output3 *Output
	Output4 *Output
	Counter int64
}

func (p *PointerStructByMethod) Calc() {
	if p.Output1.GetValue() > p.Output4.GetValue() {
		p.Counter++
	}
	if p.Output2.GetValue() > p.Output1.GetValue() {
		p.Counter++
	}
	if p.Output3.GetValue() > p.Output2.GetValue() {
		p.Counter++
	}
	if p.Output4.GetValue() > p.Output3.GetValue() {
		p.Counter++
	}
}

// Size usage 4.0 GB 0 B time usage 55.412901425s
func TestPointerStructByMethod(t *testing.T) {
	p := MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	outputs1 := make([]Output, 100)
	outputs2 := make([]Output, 100)
	outputs3 := make([]Output, 100)
	outputs4 := make([]Output, 100)

	p.Start()
	numberSubjects := make([]PointerStructByMethod, 0, len(outputs1)*len(outputs2)*len(outputs3)*len(outputs4))
	for o1 := range outputs1 {
		for o2 := range outputs2 {
			for o3 := range outputs3 {
				for o4 := range outputs4 {
					numberSubjects = append(numberSubjects, PointerStructByMethod{
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
		for o := range outputs3 {
			outputs4[o].UpdateValue(randomValues[i-3])
		}

		for i := range numberSubjects {
			numberSubjects[i].Calc()
		}
	}
	p.Finish()

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}

type PointerStructByValue struct {
	Output1 *Output
	Output2 *Output
	Output3 *Output
	Output4 *Output
	Counter int64
}

func (p *PointerStructByValue) Calc() {
	if p.Output1.Value > p.Output4.Value {
		p.Counter++
	}
	if p.Output2.Value > p.Output1.Value {
		p.Counter++
	}
	if p.Output3.Value > p.Output2.Value {
		p.Counter++
	}
	if p.Output4.Value > p.Output3.Value {
		p.Counter++
	}
}

// Size usage 4.0 GB 0 B time usage 55.20273306s
func TestPointerStructByValue(t *testing.T) {
	p := MemTimeProfiler{}

	rand.Seed(time.Now().UnixNano())

	outputs1 := make([]Output, 100)
	outputs2 := make([]Output, 100)
	outputs3 := make([]Output, 100)
	outputs4 := make([]Output, 100)

	p.Start()
	numberSubjects := make([]PointerStructByValue, 0, len(outputs1)*len(outputs2)*len(outputs3)*len(outputs4))
	for o1 := range outputs1 {
		for o2 := range outputs2 {
			for o3 := range outputs3 {
				for o4 := range outputs4 {
					numberSubjects = append(numberSubjects, PointerStructByValue{
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
		for o := range outputs3 {
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

// Size usage 7.2 GB 0 B time usage 1m33.737323656s // for i := 3; i < loopIterator; i++ { for y := range numberSubjects
// Size usage 7.2 GB 0 B time usage 45.516164078s // for _, v := range numberSubjects { for i := 3; i < loopIterator; i++ {
func TestValueStruct(t *testing.T) {
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
	//for i := 3; i < loopIterator; i++ {
	//	for y := range numberSubjects {
	//		numberSubjects[y].Output1.UpdateValue(randomValues[i])
	//		numberSubjects[y].Output2.UpdateValue(randomValues[i-1])
	//		numberSubjects[y].Output3.UpdateValue(randomValues[i-2])
	//		numberSubjects[y].Output4.UpdateValue(randomValues[i-3])
	//		numberSubjects[y].Calc()
	//	}
	//}
	for _, subject := range numberSubjects {
		for i := 3; i < loopIterator; i++ {
			subject.Output1.UpdateValue(randomValues[i])
			subject.Output2.UpdateValue(randomValues[i-1])
			subject.Output3.UpdateValue(randomValues[i-2])
			subject.Output4.UpdateValue(randomValues[i-3])
			subject.Calc()
		}
	}
	p.Finish()

	fmt.Println("Size usage", size, p.Size(), "time usage", p.Time())
}
