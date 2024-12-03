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
		if len(line) < 4 {
			return "", errors.New("Line doesn't contain 'mul'")
		}
		line = strings.TrimSpace(line)
		if line[:4] != "mul(" && (len(line) > 3 && line[:4] != "do()") && (len(line) > 6 && line[:7] != "don't()") {
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
	isProcessing := true

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
		fmt.Printf("After prefix removal: %s\n", inputText)
		fmt.Printf("Processing: %v\n", isProcessing)

		// Check for do/don't
		if len(inputText) > 3 && inputText[:4] == "do()" {
			isProcessing = true
			inputText = inputText[1:]
			continue
		}
		if len(inputText) > 6 && inputText[:7] == "don't()" {
			isProcessing = false
			inputText = inputText[1:]
			continue
		}

		result, err := getMulInstructionResult(inputText)
		if err != nil {
			inputText = inputText[4:]
			continue
		} else {
			if isProcessing {
				total += result
			}
			inputText = inputText[1:]
		}
	}
	return total, nil
}
