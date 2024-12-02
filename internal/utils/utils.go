package utils

import (
	"log"
	"runtime"
	"strconv"
	"time"
)

func StringToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return i
}

func Abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

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
