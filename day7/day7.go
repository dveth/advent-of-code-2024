package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part2() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var total uint64
	count := 0
	for scanner.Scan() {
		equation, err := getEquationFromLine(scanner.Text())
		if err != nil {
			panic(err)
		}
		if evaluateAll(equation.Nums, equation.TestValue) {
			total += equation.TestValue
		}
		count++
	}
	fmt.Printf("Equation Count: %d\n", count)
	fmt.Printf("Total: %d\n", total)
}

func testPart2() {
	nums := []uint64{73, 91, 24, 78, 75, 93, 11, 36, 42, 86, 25, 24, 66, 26}
	want := 13
	got := getAllOperatorCombinations(len(nums))

	if want != len(got) {
		fmt.Printf("In evaluateAll for nums %v, wanted %v, got %v\n", nums, want, len(got[0]))
	}
}

type Equation struct {
	TestValue uint64
	Nums      []uint64
}

func (e Equation) equals(other Equation) bool {
	return e.TestValue == other.TestValue && slices.Equal(e.Nums, other.Nums)
}

func getEquationFromLine(line string) (Equation, error) {
	// Line Example: 3267: 81 40 27
	lineSplit := strings.Split(line, ":")
	if len(lineSplit) < 2 {
		return Equation{}, errors.New("Too few values in line")
	}
	testValue, err := strconv.Atoi(lineSplit[0])
	if err != nil {
		return Equation{}, err
	}

	numsSplit := strings.Split(strings.TrimSpace(lineSplit[1]), " ")
	nums := []uint64{}

	for _, str := range numsSplit {
		value, err := strconv.Atoi(str)
		if err != nil {
			return Equation{}, errors.New("Non-uint64eger value after : in line")
		}
		nums = append(nums, uint64(value))
	}
	return Equation{
		TestValue: uint64(testValue),
		Nums:      nums,
	}, nil
}

// Evaluate if any combination of operators could work to make the values in nums equal the testValue
// For example, if nums was {10, 19} then operators {"*"} would make the function return 'true'
func evaluateAll(nums []uint64, testValue uint64) bool {
	operatorCombinations := getAllOperatorCombinations(len(nums))
	for _, operators := range operatorCombinations {
		numsCopy := make([]uint64, len(nums))
		copy(numsCopy, nums)
		operatorsCopy := make([]string, len(operators))
		copy(operatorsCopy, operators)
		total, err := evaluate(numsCopy, operatorsCopy)
		if err != nil {
			continue
		}
		if total == testValue {
			return true
		}
	}
	return false
}

func evaluate(nums []uint64, operators []string) (uint64, error) {
	if len(nums) != len(operators)+1 {
		return 0, errors.New("Incorrect amount of operators for amount of nums.")
	}
	total := nums[0]
	for i := 1; i < len(nums); i++ {
		if operators[i-1] == "+" {
			total += nums[i]
		}
		if operators[i-1] == "*" {
			total *= nums[i]
		}
		if operators[i-1] == "||" {
			first := strconv.FormatUint(total, 10)
			second := strconv.FormatUint(nums[i], 10)
			var result uint64
			fullString := first + second
			result, err := strconv.ParseUint(fullString, 10, 0)
			if err != nil {
				return 0, err
			}
			total = result
		}
	}
	return total, nil
}

func getAllOperatorCombinations(length int) [][]string {
	if length == 0 || length == 1 {
		return [][]string{}
	}
	if length == 2 {
		return [][]string{
			{"+"},
			{"*"},
			{"||"},
		}
	}
	result := [][]string{}
	innerResult := getAllOperatorCombinations(length - 1)
	for _, value := range innerResult {
		valueCopy := make([]string, len(value))
		copy(valueCopy, value)
		result = append(result, append(valueCopy, "+"))
		valueCopy = make([]string, len(value))
		copy(valueCopy, value)
		result = append(result, append(valueCopy, "*"))
		valueCopy = make([]string, len(value))
		copy(valueCopy, value)
		result = append(result, append(valueCopy, "||"))
	}
	return result
}
