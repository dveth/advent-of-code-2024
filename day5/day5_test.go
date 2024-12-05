package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestNewRuleFromLine(t *testing.T) {
	line := "47|53"
	want := PageRule{
		FirstNum:  47,
		SecondNum: 53,
	}
	got, err := newRuleFromLine(line)
	if err != nil {
		t.Fatalf("Got err from newRuleFromLine: %s\n", err.Error())
	}

	if want != got {
		t.Fatalf("In newRuleFromLine, wanted %v, got %v\n", want, got)
	}
}

func TestNewUpdateFromLine(t *testing.T) {
	line := "75,47,61,53,29"

	var got Update
	var want Update
	want = []int{75, 47, 61, 53, 29}
	got, err := newUpdateFromLine(line)

	if err != nil {
		t.Fatalf("Got err from newUpdateFromLine: %s\n", err.Error())
	}

	if !slices.Equal(want, got) {
		t.Fatalf("In newUpdateFromLine, wanted %v, got %v\n", want, got)
	}
}

func TestCheckUpdateRulesBroken(t *testing.T) {
	rules, updates, err := getInputFromFile("testinput.txt")
	if err != nil {
		t.Fatalf("Got err from checkUpdateRulesBroken: %s\n", err.Error())
	}
	got := []Update{}
	for _, update := range updates {
		if !checkUpdateRulesBroken(rules, update) {
			got = append(got, update)
		}
	}

	want := []Update{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("In checkUpdateRulesBroken, wanted %v, got %v", want, got)
	}

}

func TestMiddle(t *testing.T) {
	var update Update
	update = []int{75, 47, 61, 53, 29}

	want := 61
	got := update.middle()

	if want != got {
		t.Fatalf("In update.middle, wanted %d, got %d, from %v", want, got, update)
	}

	update = []int{75, 29, 13}

	want = 29
	got = update.middle()

	if want != got {
		t.Fatalf("In update.middle, wanted %d, got %d, from %v", want, got, update)
	}
}

func TestFix(t *testing.T) {
	var want Update
	want = []int{97, 75, 47, 61, 53}

	var brokenUpdate Update
	brokenUpdate = []int{75, 97, 47, 61, 53}

	rules, _, _ := getInputFromFile("testinput.txt")
	brokenUpdate.fix(rules)

	if !slices.Equal(want, brokenUpdate) {
		t.Fatalf("In update.fix, wanted %v, got %v\n", want, brokenUpdate)
	}
}
