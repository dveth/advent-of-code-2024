package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total, err := processInput("puzzleInput.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total: %d\n", total)
}

func removePrefix(line string) (string, error) {
	for {
		if len(line) < 4 || !strings.Contains(line, "mul(") {
			return "", errors.New("Line doesn't contain 'mul'")
		}
		line = strings.TrimSpace(line)
		if line[:4] != "mul(" {
			line = line[1:]
		} else {
			return line, nil
		}
	}
}

func getMulInstructionResult(line string) (int, error) {
	if line[:3] != "mul" {
		return 0, errors.New("Doesn't contain mul(")
	}

	line = line[4:]
	lineSplit := strings.Split(line, ")")
	if len(lineSplit) == 1 {
		return 0, errors.New("No closing parenthesis")
	}

	innerParens := strings.Split(lineSplit[0], ",")
	if len(innerParens) != 2 { // Fuck this line. used to be len(innerParens) == 1. So stupid of me.
		return 0, errors.New("No comma.")
	}

	firstNum, err := strconv.Atoi(innerParens[0])
	if err != nil {
		return 0, err
	}
	secondNum, err := strconv.Atoi(innerParens[1])
	if err != nil {
		return 0, err
	}

	if firstNum > 999 || secondNum > 999 || firstNum < 0 || secondNum < 0 {
		return 0, errors.New("Number not in range.")
	}

	return firstNum * secondNum, nil
}

func processInput(filename string) (int, error) {
	inputText := ""
	total := 0

	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputText += scanner.Text()
	}

	for len(inputText) > 3 {
		tempText, err := removePrefix(inputText)
		if err != nil {
			return total, nil
		}
		inputText = tempText
		result, err := getMulInstructionResult(inputText)
		if err != nil {
			inputText = inputText[4:]
			continue
		} else {
			total += result
			inputText = inputText[1:]
		}
	}
	return total, nil
}
