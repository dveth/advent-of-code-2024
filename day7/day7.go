package main

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

func main() {

}

type Equation struct {
	TestValue int
	Nums      []int
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
	nums := []int{}

	for _, str := range numsSplit {
		value, err := strconv.Atoi(str)
		if err != nil {
			return Equation{}, errors.New("Non-integer value after : in line")
		}
		nums = append(nums, value)
	}
	return Equation{
		TestValue: testValue,
		Nums:      nums,
	}, nil
}
