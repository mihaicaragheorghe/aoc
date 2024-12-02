package utils

import (
	"runtime"
	"time"
)

func MeasureExecution[T any](f func() T) (result T, duration time.Duration, memAllocated uint64) {
	var memStatsStart, memStatsEnd runtime.MemStats
	runtime.ReadMemStats(&memStatsStart)
	start := time.Now()

	result = f()

	duration = time.Since(start)
	runtime.ReadMemStats(&memStatsEnd)
	memAllocated = memStatsEnd.TotalAlloc - memStatsStart.TotalAlloc
	return
}
