package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

func SolveB() {
	result, duration, mem := utils.MeasureExecution(solveB)
	fmt.Printf("\n[D2P2] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveB() int {
	file, err := os.Open("./internal/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		nums := strings.Split(ln, " ")
		if safe, unsafeLvl := checkSafety(nums); !safe {
			if !retryWithoutLevel(unsafeLvl, nums) && !retryWithoutLevel(unsafeLvl-1, nums) && !retryWithoutLevel(0, nums) {
				continue
			}
		}
		safeCount++
	}
	return safeCount
}

func retryWithoutLevel(level int, nums []string) bool {
	cp := make([]string, len(nums))
	copy(cp, nums)

	if level == len(nums) {
		cp = cp[:level]
	} else {
		cp = append(cp[:level], cp[level+1:]...)
	}
	safe, _ := checkSafety(cp)
	return safe
}
