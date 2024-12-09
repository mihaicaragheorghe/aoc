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
	_, path, err := walk(maze, Directions.Up, start, make(map[Point]Direction), 0)
	if err != nil {
		panic(err)
	}
	return addObstacles(maze, path, start)
}

func addObstacles(maze [][]byte, path map[Point]Direction, start Point) int {
	count := 0
	for step := range path {
		if step.X == start.X && step.Y == start.Y {
			continue
		}
		maze[step.Y][step.X] = '#'
		if _, _, err := walk(maze, Directions.Up, start, make(map[Point]Direction), 0); err != nil {
			count++
		}
		maze[step.Y][step.X] = unvisted
	}
	return count
}
