package golang_benchmark

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type PublicVariable struct {
	Result TestStructure
}

type PublicMethod struct {
	result TestStructure
}

func (t *PublicMethod) Result() TestStructure {
	return t.result
}

func (t *PublicMethod) SetResult(value float64) {
	t.result.Test2 = value
}

// Total heap alloc: 0 Mb 0 bytes. Time: 892.232µs
func TestPublicVariable(t *testing.T) {
	customProfiler := MemTimeProfiler{}

	target := PublicVariable{
		Result: TestStructure{
			Test1: 1,
			Test2: 2,
			Test3: 3,
			Test4: 4,
			Test5: 5,
			Test6: 6,
		},
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := float64(rand.Intn(5-1) + 1)

	customProfiler.Start()
	for i := 0; i < million; i++ {
		counter := target.Result.Test2 + randomNumber
		target.Result.Test2 = counter
	}
	customProfiler.Finish()

	fmt.Println(target.Result.Test2)
}

// Total heap alloc: 0 Mb 0 bytes. Time: 2.730752ms
func TestPublicMethod(t *testing.T) {
	customProfiler := MemTimeProfiler{}

	target := PublicMethod{
		result: TestStructure{
			Test1: 1,
			Test2: 2,
			Test3: 3,
			Test4: 4,
			Test5: 5,
			Test6: 6,
		},
	}
	rand.Seed(time.Now().UnixNano())
	randomNumber := float64(rand.Intn(5-1) + 1)

	customProfiler.Start()
	for i := 0; i < million; i++ {
		counter := target.Result().Test2 + randomNumber
		target.SetResult(counter)
	}
	customProfiler.Finish()

	fmt.Println(target.Result().Test2)
}
