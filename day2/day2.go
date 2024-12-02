package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report []int

func (r Report) isIncreasing() bool {
	increasing := true
	for i := 1; i < len(r); i++ {
		if r[i] <= r[i-1] {
			increasing = false
		}
	}
	return increasing
}

func (r Report) isDecreasing() bool {
	decreasing := true
	for i := 1; i < len(r); i++ {
		if r[i] >= r[i-1] {
			decreasing = false
		}
	}
	return decreasing
}

func (r Report) isGradual() bool {
	gradual := true
	for i := 1; i < len(r); i++ {
		difference := getDifference(r[i], r[i-1])
		if difference > 3 || difference < 1 {
			gradual = false
		}
	}
	return gradual
}

func (r Report) isSafe() bool {
	return (r.isIncreasing() || r.isDecreasing()) && r.isGradual()
}

func (r Report) isSafeDampened() bool {
	if r.isSafe() {
		return true
	}

	for i := 0; i < len(r); i++ {
		var newSlice Report
		newSlice = sliceWithRemoved(r, i)
		if newSlice.isSafe() {
			return true
		}
	}
	return false
}

func newReport(line string) (Report, error) {
	lineSplit := strings.Split(line, " ")
	var report Report
	for _, value := range lineSplit {
		num, err := strconv.Atoi(value)
		if err != nil {
			return report, err
		}
		report = append(report, num)
	}
	return report, nil
}

func main() {
	total, err := day1Part2("puzzleInput.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total: %d\n", total)
}

func day1Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		report, err := newReport(scanner.Text())
		if err != nil {
			return total, err
		}
		if report.isSafe() {
			total += 1
		}
	}
	return total, nil
}

func day1Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		report, err := newReport(scanner.Text())
		if err != nil {
			return total, err
		}
		if report.isSafeDampened() {
			total += 1
		}
	}
	return total, nil
}

// Gets the difference between two integers, or the absolute value of subtracting one number from the other
func getDifference(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

func sliceWithRemoved(slice []int, s int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:s], newSlice[s+1:]...)
}
