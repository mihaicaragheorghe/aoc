package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

var instructionsRegex = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

func SolveB() {
	result, duration, mem := utils.MeasureExecution(solveB)
	fmt.Printf("\n[D3P2] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveB() int {
	file, err := os.Open("./internal/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text()
	}
	return findMultiplicationsWithInstructions(text)
}

func findMultiplicationsWithInstructions(s string) int {
	sum := 0
	mulEnabled := true
	matches := instructionsRegex.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		m := strings.Join(match, "")
		if m == "do()" {
			mulEnabled = true
		} else if m == "don't()" {
			mulEnabled = false
		} else if mulEnabled {
			sum += multiplyMatch(match)
		}
	}
	return sum
}
