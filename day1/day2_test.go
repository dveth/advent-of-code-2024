package main

import (
	"slices"
	"testing"
)

func TestMakeLists(t *testing.T) {
	wantList1 := []int{3, 4, 2, 1, 3, 3}
	wantList2 := []int{4, 3, 5, 3, 9, 3}

	gotList1, gotList2, err := makeLists("./testdata.txt")
	if err != nil {
		t.Fatalf("Got error in makeLists: %s", err.Error())
	}

	if !slices.Equal(wantList1, gotList1) {
		t.Fatalf("In makelists list1, wanted %v, got %v", wantList1, gotList1)
	}
	if !slices.Equal(wantList2, gotList2) {
		t.Fatalf("In makelists list2, wanted %v, got %v", wantList2, gotList2)
	}
}

func TestMakePairs(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	want := []Pair{
		{1, 3},
		{2, 3},
		{3, 3},
		{3, 4},
		{3, 5},
		{4, 9},
	}
	got, err := makePairs(list1, list2)
	if err != nil {
		t.Fatalf("In makePairs, got err: %s", err.Error())
	}

	if !slices.Equal(want, got) {
		t.Fatalf("In makePairs, wanted %v, got %v", want, got)
	}
}

func TestGetTotalDistance(t *testing.T) {
	pairs := []Pair{
		{1, 3},
		{2, 3},
		{3, 3},
		{3, 4},
		{3, 5},
		{4, 9},
	}
	want := 11
	got := getTotalDistance(pairs)

	if want != got {
		t.Fatalf("In getTotalDistance, wanted %d, got %d", want, got)
	}
}

func TestSimilarityScore(t *testing.T) {
	list := []int{4, 3, 5, 3, 9, 3}
	num := 3

	want := 9
	got := getSimilarityScore(num, list)

	if want != got {
		t.Fatalf("In getSimilarityScore, wanted %d, got %d", want, got)
	}
}

func TestTotalSimilarityScore(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	want := 31
	got, err := getTotalSimilarityScore(list1, list2)
	if err != nil {
		t.Fatalf("In getTotalSimilarityScore, received error: %s", err.Error())
	}

	if want != got {
		t.Fatalf("In getTotalSimilarityScore, wanted %d, got %d", want, got)
	}
}
