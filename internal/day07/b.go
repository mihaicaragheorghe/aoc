package day07

import (
	"fmt"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

func SolveB() {
	result, duration, mem := utils.MeasureExecution(solveB)
	fmt.Printf("\n[D7P2] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveB() int {
	sum := 0
	eqs := parseInput()

	for _, eq := range eqs {
		if solve(eq.result, eq.nums, []byte{add, mul, cat}) {
			sum += int(eq.result)
		}
	}
	return sum
}
