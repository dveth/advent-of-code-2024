package main

import (
	"slices"
	"testing"
)

func TestGetDifference(t *testing.T) {
	num1 := 5
	num2 := 7
	want := 2
	got := getDifference(num1, num2)

	if want != got {
		t.Fatalf("In getDifference, wanted %d, got %d", want, got)
	}

	num1 = 10
	num2 = 6
	want = 4
	got = getDifference(num1, num2)

	if want != got {
		t.Fatalf("In getDifference, wanted %d, got %d", want, got)
	}
}

func TestIsIncreasing(t *testing.T) {
	var report Report
	report = []int{1, 2, 7, 8, 9}
	want := true
	got := report.isIncreasing()

	if want != got {
		t.Fatalf("In isIncreasing, wanted %v, got %v, for %v", want, got, report)
	}

	report = []int{1, 3, 2, 4, 5}
	want = false
	got = report.isIncreasing()

	if want != got {
		t.Fatalf("In isIncreasing, wanted %v, got %v, for %v", want, got, report)
	}
}

func TestIsDecreasing(t *testing.T) {
	var report Report
	report = []int{1, 2, 7, 8, 9}
	want := false
	got := report.isDecreasing()

	if want != got {
		t.Fatalf("In isIncreasing, wanted %v, got %v, for %v", want, got, report)
	}

	report = []int{8, 6, 4, 4, 1}
	want = false
	got = report.isDecreasing()

	if want != got {
		t.Fatalf("In isIncreasing, wanted %v, got %v, for %v", want, got, report)
	}

	report = []int{9, 7, 6, 2, 1}
	want = true
	got = report.isDecreasing()

	if want != got {
		t.Fatalf("In isIncreasing, wanted %v, got %v, for %v", want, got, report)
	}
}

func TestIsGradual(t *testing.T) {
	var report Report
	report = []int{1, 2, 7, 8, 9}
	want := false
	got := report.isGradual()

	if want != got {
		t.Fatalf("In isGradual, wanted %v, got %v, for %v", want, got, report)
	}

	report = []int{7, 6, 4, 2, 1}
	want = true
	got = report.isGradual()

	if want != got {
		t.Fatalf("In isGradual, wanted %v, got %v, for %v", want, got, report)
	}
}

func TestIsSafe(t *testing.T) {
	var report Report
	report = []int{1, 2, 7, 8, 9}
	want := false
	got := report.isSafe()

	if want != got {
		t.Fatalf("In isSafe, wanted %v, got %v, for %v", want, got, report)
	}

	report = []int{7, 6, 4, 2, 1}
	want = true
	got = report.isSafe()

	if want != got {
		t.Fatalf("In isSafe, wanted %v, got %v, for %v", want, got, report)
	}
}

func TestNewReport(t *testing.T) {
	line := "7 6 4 2 1"
	var want Report
	want = []int{7, 6, 4, 2, 1}
	got, err := newReport(line)
	if err != nil {
		t.Fatalf("In newReport, received error: %s", err.Error())
	}

	if !slices.Equal(got, want) {
		t.Fatalf("In newReport, wanted %v, got %v, for %s", want, got, line)
	}
}

func TestDay2Part1(t *testing.T) {
	filename := "testdata.txt"
	want := 2
	got, err := day1Part1(filename)

	if err != nil {
		t.Fatalf("In day2Part1, received error: %s", err.Error())
	}

	if want != got {
		t.Fatalf("In day2Part1, wanted %d, got %d", want, got)
	}
}

func TestSliceWithRemoved(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4}
	got := sliceWithRemoved(slice, 4)

	if !slices.Equal(want, got) {
		t.Fatalf("In sliceWithRemoved, wanted %v, got %v", want, got)
	}

	want = []int{2, 3, 4, 5}
	got = sliceWithRemoved(slice, 0)

	if !slices.Equal(want, got) {
		t.Fatalf("In sliceWithRemoved, wanted %v, got %v", want, got)
	}

	want = []int{1, 2, 4, 5}
	got = sliceWithRemoved(slice, 2)

	if !slices.Equal(want, got) {
		t.Fatalf("In sliceWithRemoved, wanted %v, got %v", want, got)
	}
}

func TestSafeWithDampened(t *testing.T) {
	var report Report
	report = []int{1, 2, 7, 8, 9}
	want := false
	got := report.isSafeDampened()

	if want != got {
		t.Fatalf("In isSafe, wanted %v, got %v, for %v", want, got, report)
	}

	report = []int{7, 6, 4, 2, 1}
	want = true
	got = report.isSafeDampened()

	if want != got {
		t.Fatalf("In isSafe, wanted %v, got %v, for %v", want, got, report)
	}

	report = []int{1, 3, 2, 4, 5}
	want = true
	got = report.isSafeDampened()

	if want != got {
		t.Fatalf("In isSafe, wanted %v, got %v, for %v", want, got, report)
	}
}
