package day05

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
	fmt.Printf("\n[D5P2] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

func solveB() int {
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
			if fixPages(pages, rules) {
				sum += utils.S2i(pages[len(pages)/2])
			}
		}
	}
	return sum
}

func fixPages(pages []string, rules map[string][]string) bool {
	ordered := checkPages(pages, rules)
	if ordered {
		return false
	}
	for !ordered {
		for i, p := range pages {
			if val, ok := rules[p]; ok {
				for _, v := range val {
					if exists, idx := lookupPage(pages, v, i+1); exists {
						swapPages(pages, i, idx)
						break
					}
				}
			}
		}
		ordered = checkPages(pages, rules)
	}
	return true
}

func swapPages(pages []string, i, j int) {
	tmp := pages[i]
	pages[i] = pages[j]
	pages[j] = tmp
}
