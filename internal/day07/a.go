package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mihaicaragheorghe/aoc/internal/utils"
)

func SolveA() {
	result, duration, mem := utils.MeasureExecution(solveA)
	fmt.Printf("\n[D7P1] Result: %d | Memory: %s | Duration: %v\n\n", result, utils.FormatBytes(mem), duration)
}

type Equation struct {
	result int64
	nums   []int64
}

const (
	add = '+'
	mul = '*'
	cat = '|'
)

func solveA() int {
	sum := 0
	eqs := parseInput()

	for _, eq := range eqs {
		if solve(eq.result, eq.nums, []byte{add, mul}) {
			sum += int(eq.result)
		}
	}
	return sum
}

func solve(total int64, nums []int64, ops []byte) bool {
	if len(nums) == 1 {
		return nums[0] == total
	}

	if any(ops, mul) && total%nums[len(nums)-1] == 0 && solve(total/nums[len(nums)-1], nums[:len(nums)-1], ops) {
		return true
	}
	if rem := total - nums[len(nums)-1]; any(ops, add) && rem >= 0 && solve(rem, nums[:len(nums)-1], ops) {
		return true
	}
	if any(ops, cat) {
		totalS := strconv.Itoa(int(total))
		lastNumS := strconv.Itoa(int(nums[len(nums)-1]))
		lenDiff := len(totalS) - len(lastNumS)
		if len(totalS) > len(lastNumS) &&
			totalS[lenDiff:] == lastNumS &&
			solve(utils.Atoi64(totalS[:lenDiff]), nums[:len(nums)-1], ops) {

			return true
		}
	}
	return false
}

func any(hay []byte, needle byte) bool {
	for _, b := range hay {
		if b == needle {
			return true
		}
	}
	return false
}

func parseInput() []Equation {
	file, err := os.Open("./internal/day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var eqs []Equation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		for i := 0; i < len(s); i++ {
			if s[i] == ':' {
				val := utils.Atoi64(s[:i])
				nums := strings.Split(s[i+2:], " ")
				eqs = append(eqs, Equation{val, utils.SliceAtoi64(nums)})
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return eqs
}
