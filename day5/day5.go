package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, updates, err := getInputFromFile("puzzleInput.txt")
	if err != nil {
		panic(err)
	}

	invalidUpdates := []Update{}
	for _, update := range updates {
		if checkUpdateRulesBroken(rules, update) {
			for checkUpdateRulesBroken(rules, update) {
				update.fix(rules)
			}
			fmt.Printf("Update after fix: %v\n", update)
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	totalMiddles := 0
	for _, update := range invalidUpdates {
		totalMiddles += update.middle()
	}
	fmt.Printf("Total of middle page numbers: %d\n", totalMiddles)
}

type PageRule struct {
	FirstNum  int
	SecondNum int
}

func newRuleFromLine(line string) (PageRule, error) {
	lineSplit := strings.Split(line, "|")
	firstNum, err := strconv.Atoi(lineSplit[0])
	if err != nil {
		return PageRule{}, err
	}
	secondNum, err := strconv.Atoi(lineSplit[1])
	if err != nil {
		return PageRule{}, err
	}
	return PageRule{
		FirstNum:  firstNum,
		SecondNum: secondNum,
	}, err
}

type Update []int

func (u Update) contains(x int) bool {
	for _, value := range u {
		if value == x {
			return true
		}
	}
	return false
}

func (u Update) middle() int {
	return u[(len(u)-1)/2]
}

func (u Update) fix(rules []PageRule) {
	for i := 0; i < len(u); i++ {
		for _, rule := range rules {
			isBroken, i, j := checkPosRuleBroken(rule, u, i)
			if isBroken {
				u[i], u[j] = u[j], u[i]
			}
		}
	}
}

func newUpdateFromLine(line string) (update Update, err error) {
	lineSplit := strings.Split(line, ",")
	for _, value := range lineSplit {
		num, err := strconv.Atoi(value)
		if err != nil {
			return update, err
		}
		update = append(update, num)
	}
	return update, nil
}

func getInputFromFile(filename string) (rules []PageRule, updates []Update, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return rules, updates, err
	}

	scanner := bufio.NewScanner(file)
	doingRules := true
	for scanner.Scan() {
		if scanner.Text() == "" {
			doingRules = false
			continue
		}
		if doingRules {
			newRule, err := newRuleFromLine(scanner.Text())
			if err != nil {
				return rules, updates, err
			}
			rules = append(rules, newRule)
		} else {
			newUpdate, err := newUpdateFromLine(scanner.Text())
			if err != nil {
				return rules, updates, err
			}
			updates = append(updates, newUpdate)
		}
	}
	return rules, updates, nil
}

func checkPosRuleBroken(rule PageRule, update Update, pos int) (bool, int, int) {
	if rule.FirstNum != update[pos] {
		return false, -1, -1
	}
	for i := 0; i < pos; i++ {
		if update[i] == rule.SecondNum {
			return true, i, pos
		}
	}
	return false, -1, -1
}

func checkUpdateRulesBroken(rules []PageRule, update Update) bool {
	for i := 0; i < len(update); i++ {
		for _, rule := range rules {
			isBroken, _, _ := checkPosRuleBroken(rule, update, i)
			if isBroken {
				return true
			}
		}
	}
	return false
}
