package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

func SolveA() {
	result, duration, mem := utils.MeasureExecution(solvea)
	fmt.Printf("\n[D1P1] Result: %d | Memory: %v bytes | Duration: %v\n\n", result, mem, duration)
}

func solvea() int {
	file, err := os.Open("./internal/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		left = append(left, utils.StringToInt(split[0]))
		right = append(right, utils.StringToInt(split[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += utils.Abs(left[i], right[i])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return sum
}
