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

type Pair struct {
	num1 int
	num2 int
}

func (p Pair) getDistance() int {
	if p.num1 > p.num2 {
		return p.num1 - p.num2
	} else {
		return p.num2 - p.num1
	}
}

func main() {
	list1, list2, err := makeLists("./puzzleInput.txt")
	if err != nil {
		panic(err)
	}
	score, err := getTotalSimilarityScore(list1, list2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total score: %d\n", score)
}

func makeLists(filename string) (list1 []int, list2 []int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return list1, list2, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "   ")
		num1, err := strconv.Atoi(splitLine[0])
		if err != nil {
			return list1, list2, err
		}
		num2, err := strconv.Atoi(splitLine[1])
		if err != nil {
			return list1, list2, err
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	return list1, list2, err
}

func makePairs(list1 []int, list2 []int) (pairs []Pair, err error) {
	if len(list1) != len(list2) {
		return pairs, errors.New("Lengths are not equal.")
	}
	slices.Sort(list1)
	slices.Sort(list2)
	for i := 0; i < len(list1); i++ {
		newPair := Pair{
			num1: list1[i],
			num2: list2[i],
		}
		pairs = append(pairs, newPair)
	}
	return pairs, nil
}

func getTotalDistance(pairs []Pair) (total int) {
	for _, pair := range pairs {
		total += pair.getDistance()
	}
	return total
}

func getSimilarityScore(num int, list []int) (score int) {
	for _, value := range list {
		if value == num {
			score += 1
		}
	}
	return score * num
}

func getTotalSimilarityScore(list1 []int, list2 []int) (score int, err error) {
	if len(list1) != len(list2) {
		return score, err
	}

	for _, value := range list1 {
		score += getSimilarityScore(value, list2)
	}

	return score, nil
}
