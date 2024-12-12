package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestGetEquationFromLine(t *testing.T) {
	line := "190: 10 19"
	want := Equation{
		TestValue: 190,
		Nums:      []uint64{10, 19},
	}
	got, err := getEquationFromLine(line)
	if err != nil {
		t.Fatalf("Received error in getEquationFromLine: %s\n", err.Error())
	}

	if !got.equals(want) {
		t.Fatalf("In getEquationFromLine, wanted %v, got %v from %s\n", want, got, line)
	}

	line = "161011: 16 10 13"
	want = Equation{
		TestValue: 161011,
		Nums:      []uint64{16, 10, 13},
	}
	got, err = getEquationFromLine(line)
	if err != nil {
		t.Fatalf("Received error in getEquationFromLine: %s\n", err.Error())
	}

	if !got.equals(want) {
		t.Fatalf("In getEquationFromLine, wanted %v, got %v from %s\n", want, got, line)
	}

	line = "18: f 6"

	_, err = getEquationFromLine(line)
	if err == nil {
		t.Fatalf("Expected error in getEquationFromLine for input: %s\n", line)
	}
}

func TestGetAllOperatorCombinations(t *testing.T) {
	length := 2
	want := [][]string{
		{"+"},
		{"*"},
		{"||"},
	}
	got := getAllOperatorCombinations(length)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("In getAllOperatorCombinations for length %d, wanted %v, got %v\n", length, want, got)
	}

	length = 3
	want = [][]string{
		{"+", "+"},
		{"+", "*"},
		{"+", "||"},
		{"*", "+"},
		{"*", "*"},
		{"*", "||"},
		{"||", "+"},
		{"||", "*"},
		{"||", "||"},
	}
	got = getAllOperatorCombinations(length)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("In getAllOperatorCombinations for length %d, wanted %v, got %v\n", length, want, got)
	}

}

func TestEvaluate(t *testing.T) {
	nums := []uint64{10, 19}
	operators := []string{"+"}
	got, err := evaluate(nums, operators)
	var want uint64
	want = 29
	if err != nil {
		t.Fatalf("Received error in evaluate for nums %v and operators %v: %s\n", nums, operators, err.Error())
	}

	if want != got {
		t.Fatalf("In evaluate for nums %v and operators %v, wanted %d, got %d\n", nums, operators, want, got)
	}

	nums = []uint64{10, 19}
	operators = []string{"*"}
	got, err = evaluate(nums, operators)
	want = 190
	if err != nil {
		t.Fatalf("Received error in evaluate for nums %v and operators %v: %s\n", nums, operators, err.Error())
	}

	if want != got {
		t.Fatalf("In evaluate for nums %v and operators %v, wanted %d, got %d\n", nums, operators, want, got)
	}

	nums = []uint64{81, 40, 27}
	operators = []string{"+", "*"}
	got, err = evaluate(nums, operators)
	want = 3267
	if err != nil {
		t.Fatalf("Received error in evaluate for nums %v and operators %v: %s\n", nums, operators, err.Error())
	}

	if want != got {
		t.Fatalf("In evaluate for nums %v and operators %v, wanted %d, got %d\n", nums, operators, want, got)
	}

	nums = []uint64{81, 40, 27}
	operators = []string{"*", "+"}
	got, err = evaluate(nums, operators)
	want = 3267
	if err != nil {
		t.Fatalf("Received error in evaluate for nums %v and operators %v: %s\n", nums, operators, err.Error())
	}

	if want != got {
		t.Fatalf("In evaluate for nums %v and operators %v, wanted %d, got %d\n", nums, operators, want, got)
	}
}

func TestEvaluateAll(t *testing.T) {
	testValue := uint64(190)
	nums := []uint64{10, 19}
	want := true
	got := evaluateAll(nums, testValue)

	if want != got {
		t.Fatalf("In evaluateAll for nums %v and testValue %d, wanted %v, got %v\n", nums, testValue, want, got)
	}

	testValue = 3267
	nums = []uint64{81, 40, 27}
	want = true
	got = evaluateAll(nums, testValue)

	if want != got {
		t.Fatalf("In evaluateAll for nums %v and testValue %d, wanted %v, got %v\n", nums, testValue, want, got)
	}

	testValue = 21037
	nums = []uint64{9, 7, 18, 13}
	want = false
	got = evaluateAll(nums, testValue)

	if want != got {
		t.Fatalf("In evaluateAll for nums %v and testValue %d, wanted %v, got %v\n", nums, testValue, want, got)
	}

	testValue = 292
	nums = []uint64{11, 6, 16, 20}
	want = true
	got = evaluateAll(nums, testValue)

	if want != got {
		t.Fatalf("In evaluateAll for nums %v and testValue %d, wanted %v, got %v\n", nums, testValue, want, got)
	}

	testValue = 11364228
	nums = []uint64{46, 26, 2, 2, 9, 527}
	want = true
	got = evaluateAll(nums, testValue)

	if want != got {
		t.Fatalf("In evaluateAll for nums %v and testValue %d, wanted %v, got %v\n", nums, testValue, want, got)
	}

	testValue = 536008016
	nums = []uint64{73, 91, 24, 78, 75, 93, 11, 36, 42, 86, 25, 24, 66, 26}
	want = true
	got = evaluateAll(nums, testValue)

	if want != got {
		t.Fatalf("In evaluateAll for nums %v and testValue %d, wanted %v, got %v\n", nums, testValue, want, got)
	}

	testValue = 156
	nums = []uint64{15, 6}
	want = true
	got = evaluateAll(nums, testValue)

	if want != got {
		t.Fatalf("In evaluateAll for nums %v and testValue %d, wanted %v, got %v\n", nums, testValue, want, got)
	}
}

func TestPart1(t *testing.T) {
	file, err := os.Open("testinput.txt")
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

	wantCount := 9
	if count != wantCount {
		t.Fatalf("In part1, count should be %d, got %d\n", wantCount, count)
	}

	var wantTotal uint64
	wantTotal = 11387
	if wantTotal != total {
		t.Fatalf("In part1, total should be %d, got %d\n", wantTotal, total)
	}

}
