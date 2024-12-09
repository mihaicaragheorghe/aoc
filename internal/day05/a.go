package day05

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
	fmt.Printf("\n[D5P1] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveA() int {
	file, err := os.Open("./internal/day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	section := 0
	rules := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			section++
		} else if section == 0 {
			r := strings.Split(s, "|")
			rules[r[1]] = append(rules[r[1]], r[0])
		} else {
			pages := strings.Split(s, ",")
			if checkPages(pages, rules) {
				sum += utils.S2i(pages[len(pages)/2])
			}
		}
	}
	return sum
}

func checkPages(pages []string, rules map[string][]string) bool {
	for i, p := range pages {
		if val, ok := rules[p]; ok {
			for _, v := range val {
				if exists, _ := lookupPage(pages, v, i+1); exists {
					return false
				}
			}
		}
	}
	return true
}

func lookupPage(pages []string, page string, from int) (bool, int) {
	for i := from; i < len(pages); i++ {
		if pages[i] == page {
			return true, i
		}
	}
	return false, -1
}
