package main

import "testing"

func TestGetEquationFromLine(t *testing.T) {
	line := "190: 10 19"
	want := Equation{
		TestValue: 190,
		Nums:      []int{10, 19},
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
		Nums:      []int{16, 10, 13},
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
