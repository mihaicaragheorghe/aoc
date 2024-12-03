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

var mulRegex = regexp.MustCompile(`mul\(\d+,\d+\)`)

func SolveA() {
	result, duration, mem := utils.MeasureExecution(solveA)
	fmt.Printf("\n[D3P1] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveA() int {
	file, err := os.Open("./internal/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += findMultiplications(scanner.Text())
	}
	return sum
}

func findMultiplications(s string) int {
	sum := 0
	matches := mulRegex.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		sum += multiplyMatch(match)
	}
	return sum
}

func multiplyMatch(match []string) int {
	first, second := parseDigits(match)
	return first * second
}

func parseDigits(match []string) (int, int) {
	s := strings.Join(match, "")
	digits := strings.Split(s[4:len(s)-1], ",")
	return utils.S2i(digits[0]), utils.S2i(digits[1])
}
