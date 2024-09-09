package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	lines, err := readLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	lines = replaceSpelledDigits(lines)
	for _, line := range lines {
		digits := findDigits(line)
		if len(digits) > 0 {
			num := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			sum += i
		}
	}
	fmt.Printf("The result is %d\n", sum)
}

func readLines(file string) ([]string, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}

func findDigits(s string) []int {
	var digits []int
	for _, r := range s {
		if unicode.IsDigit(r) {
			digits = append(digits, int(r-'0'))
		}
	}
	return digits
}

func replaceSpelledDigits(arr []string) []string {
	r := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)

	for i, s := range arr {
		arr[i] = r.Replace(s)
	}

	return arr
}
