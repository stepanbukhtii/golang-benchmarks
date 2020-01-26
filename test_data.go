package golang_benchmark

import (
	"fmt"
	"runtime"
	"time"
)

const million = 1000000

type TestStructure struct {
	Test1 int64
	Test2 float64
	Test3 float64
	Test4 float64
	Test5 float64
	Test6 float64
}

type CustomProfiler struct {
	m1, m2    runtime.MemStats
	startTime time.Time
}

func (c *CustomProfiler) Start() {
	c.startTime = time.Now()
	runtime.ReadMemStats(&c.m1)
}

func (c *CustomProfiler) Finish() {
	runtime.ReadMemStats(&c.m2)

	elapsed := time.Since(c.startTime)
	fmt.Printf("Total heap alloc: %d Mb %d bytes. Time: %s\n",
		(c.m2.TotalAlloc-c.m1.TotalAlloc)/1024/1024, c.m2.TotalAlloc-c.m1.TotalAlloc, elapsed.String())
}
