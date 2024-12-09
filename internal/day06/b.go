package day06

import (
	"fmt"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

func SolveB() {
	result, duration, mem := utils.MeasureExecution(solveB)
	fmt.Printf("\n[D6P2] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveB() int {
	maze := parseInput()
	start := getStart(maze)
	path := make(map[Point]Direction)
	_, err := walk(maze, Directions.Up, start, path, 0)
	if err != nil {
		panic(err)
	}
	return addObstacles(maze, path, start)
}

func addObstacles(maze [][]byte, path map[Point]Direction, start Point) int {
	count := 0
	pathTmp := make(map[Point]Direction)
	for step := range path {
		if step.X == start.X && step.Y == start.Y {
			continue
		}

		maze[step.Y][step.X] = '#'
		if _, err := walk(maze, Directions.Up, start, pathTmp, 0); err != nil {
			count++
		}
		maze[step.Y][step.X] = unvisted

		for k := range pathTmp {
			delete(pathTmp, k)
		}
	}
	return count
}
