package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

func SolveA() {
	result, duration, mem := utils.MeasureExecution(solveA)
	fmt.Printf("\n[D6P1] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

type Point struct {
	X, Y int
}

type Direction struct {
	X, Y int
}

var Directions = struct {
	Up, Down, Left, Right Direction
}{
	Up:    Direction{0, -1},
	Down:  Direction{0, 1},
	Left:  Direction{-1, 0},
	Right: Direction{1, 0},
}

const (
	obstacle = '#'
	unvisted = '.'
	start    = '^'
)

func solveA() int {
	maze := parseInput()
	start := getStart(maze)
	path := make(map[Point]Direction)
	k, _, err := walk(maze, Directions.Up, start, path, 0)
	if err != nil {
		panic(err)
	}
	return k
}

func parseInput() [][]byte {
	file, err := os.Open("./internal/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var maze [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		maze = append(maze, make([]byte, len(s)))
		for i := 0; i < len(s); i++ {
			maze[len(maze)-1][i] = s[i]
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return maze
}

func getStart(maze [][]byte) Point {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == start {
				return Point{j, i}
			}
		}
	}
	return Point{-1, -1}
}

func walk(maze [][]byte, dir Direction, curr Point, path map[Point]Direction, k int) (int, map[Point]Direction, error) {
	if !inBounds(maze, curr) {
		return k, path, nil
	}

	val, ok := path[curr]
	if ok && val == dir {
		return k, path, fmt.Errorf("infinite loop at count %d", k)
	} else if !ok {
		k++
	}
	path[curr] = dir

	n := next(dir, curr)
	for inBounds(maze, n) && maze[n.Y][n.X] == obstacle {
		dir = right(dir)
		n = next(dir, curr)
	}

	return walk(maze, dir, n, path, k)
}

func next(d Direction, p Point) Point {
	next := Point{p.X + d.X, p.Y + d.Y}
	return next
}

func right(curr Direction) Direction {
	switch curr {
	case Directions.Down:
		return Directions.Left
	case Directions.Left:
		return Directions.Up
	case Directions.Right:
		return Directions.Down
	default:
		return Directions.Right
	}
}

func inBounds(maze [][]byte, p Point) bool {
	return p.Y >= 0 && p.Y < len(maze) && p.X >= 0 && p.X < len(maze[p.Y])
}
