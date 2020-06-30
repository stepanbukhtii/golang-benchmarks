package golang_benchmark

import (
	"fmt"
	"runtime"
	"time"
)

type MemTimeProfiler struct {
	m1, m2    runtime.MemStats
	startTime time.Time
	elapsed   time.Duration
}

func (c *MemTimeProfiler) Start() {
	runtime.ReadMemStats(&c.m1)
	c.startTime = time.Now()
}

func (c *MemTimeProfiler) Finish() {
	c.elapsed = time.Since(c.startTime)
	runtime.ReadMemStats(&c.m2)
}

func (c MemTimeProfiler) Size() string {
	totalSize := int64(c.m2.TotalAlloc - c.m1.TotalAlloc)
	return BytesToSizeString(totalSize)
}

func (c MemTimeProfiler) Time() string {
	return c.elapsed.String()
}

func BytesToSizeString(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
