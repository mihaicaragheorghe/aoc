package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

var (
	NW = Direction{-1, -1}
	N  = Direction{-1, 0}
	NE = Direction{-1, 1}
	W  = Direction{0, -1}
	E  = Direction{0, 1}
	SW = Direction{1, -1}
	S  = Direction{1, 0}
	SE = Direction{1, 1}

	Directions = []Direction{NW, N, NE, W, E, SW, S, SE}
)

func SolveA() {
	result, duration, mem := utils.MeasureExecution(solveA)
	fmt.Printf("\n[D4P1] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveA() int {
	found := 0
	matrix := buildMatrix()
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'X' {
				found += find(matrix, i, j, "XMAS")
			}
		}
	}
	return found
}

func find(matrix [][]rune, i, j int, needle string) int {
	matches := 0
	for _, d := range Directions {
		if backtrack(matrix, needle, d, i, j, 1) {
			matches++
		}
	}
	return matches
}

func backtrack(hay [][]rune, needle string, direction Direction, i, j, n int) bool {
	curr := Direction{i + direction.X, j + direction.Y}

	if n == len(needle) {
		return true
	}
	if curr.X < 0 || curr.X >= len(hay) || curr.Y < 0 || curr.Y >= len(hay[i]) {
		return false
	}
	if hay[curr.X][curr.Y] != rune(needle[n]) {
		return false
	}

	return backtrack(hay, needle, direction, curr.X, curr.Y, n+1)
}

func buildMatrix() [][]rune {
	file, err := os.Open("./internal/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		row := make([]rune, len(ln))
		for col, r := range ln {
			row[col] = r
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

type Direction struct {
	X, Y int
}
