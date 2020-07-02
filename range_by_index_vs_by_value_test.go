package golang_benchmark

import (
	"fmt"
	"testing"
)

// range by index vs range by value
// Result: range by value is faster when value is simple
// range value is struct is equal speed

// Size usage 0 B time usage 75.072573ms
func TestLoopByIndex(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var counter int64
	p.Start()
	for i := 0; i < tenMillions; i++ {
		if counter > randomValues[i].Test1 {
			counter = counter - randomValues[i].Test1
		} else {
			counter = counter + randomValues[i].Test1
		}
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}

// Size usage 0 B time usage 77.020588ms
func TestLoopByValue(t *testing.T) {
	p := MemTimeProfiler{}
	randomValues := GenerateRandomTestStructures(tenMillions)

	var counter int64
	p.Start()
	for _, v := range randomValues {
		if counter > v.Test1 {
			counter = counter - v.Test1
		} else {
			counter = counter + v.Test1
		}
	}
	p.Finish()

	fmt.Println("Size usage", p.Size(), "time usage", p.Time())
}
