package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

func SolveA() {
	result, duration, mem := utils.MeasureExecution(solveA)
	fmt.Printf("\n[D2P1] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveA() int {
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
		if safe, _ := checkSafety(nums); safe {
			safeCount++
		}
	}
	return safeCount
}

func checkSafety(nums []string) (isSafe bool, unsafeLevel int) {
	var prev int
	decreasing := false

	for i, v := range nums {
		curr := utils.S2i(v)
		if i == 0 {
			prev = curr
			continue
		}
		if prev == curr {
			return false, i
		}
		if i == 1 {
			decreasing = curr < prev
		}
		if decreasing {
			if curr > prev {
				return false, i
			}
			if prev-curr > 3 {
				return false, i
			}
		} else {
			if curr < prev {
				return false, i
			}
			if curr-prev > 3 {
				return false, i
			}
		}
		prev = curr
	}
	return true, -1
}
