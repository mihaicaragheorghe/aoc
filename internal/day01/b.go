package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

func SolveB() {
	result, duration, mem := utils.MeasureExecution(solveb)
	fmt.Printf("\n[D1P2] Result: %d | Memory: %v bytes | Duration: %v\n\n", result, mem, duration)
}

func solveb() int {
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

	m := make(map[int]int)
	for _, v := range right {
		m[v]++
	}
	sum := 0
	for _, v := range left {
		sum += v * m[v]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return sum
}
