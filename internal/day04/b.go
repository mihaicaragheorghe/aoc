package day04

import (
	"fmt"
	"slices"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

var (
	DiagonalDirections = []Direction{NW, NE, SW, SE}
)

func SolveB() {
	result, duration, mem := utils.MeasureExecution(solveB)
	fmt.Printf("\n[D4P2] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveB() int {
	count := 0
	matrix := buildMatrix()
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if matrix[i][j] == 'A' && searchDiagonally(matrix, i, j, "MAS") {
				count++
			}
		}
	}
	return count
}

func searchDiagonally(hay [][]rune, x, y int, needle string) bool {
	bd := make([]rune, 2)
	ac := make([]rune, 2)
	for i, d := range []Direction{NW, SE} {
		bd[i] = hay[x+d.X][y+d.Y]
	}
	for i, d := range []Direction{NE, SW} {
		ac[i] = hay[x+d.X][y+d.Y]
	}
	return slices.Contains(bd, rune(needle[0])) && slices.Contains(bd, rune(needle[2])) &&
		slices.Contains(ac, rune(needle[0])) && slices.Contains(ac, rune(needle[2]))
}
