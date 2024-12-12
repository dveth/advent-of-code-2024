package main

import (
	"slices"
	"testing"
)

func TestMakeGridFromFile(t *testing.T) {
	filename := "testInput.txt"
	want := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}
	got, err := makeGridFromFile(filename)

	if err != nil {
		t.Fatalf("In makeGridFromFile, received error: %s\n", err.Error())
	}

	if !slices.Equal(got, want) {
		t.Fatalf("In makeGridFromFile, wanted %v\n, got %v\n", want, got)
	}
}

func TestGetAntennas(t *testing.T) {
	grid, _ := makeGridFromFile("testInput.txt")
	want := []Antenna{
		{
			Pos:    Pos{X: 8, Y: 1},
			Signal: "0",
		}, {
			Pos:    Pos{X: 5, Y: 2},
			Signal: "0",
		}, {
			Pos:    Pos{X: 7, Y: 3},
			Signal: "0",
		}, {
			Pos:    Pos{X: 4, Y: 4},
			Signal: "0",
		}, {
			Pos:    Pos{X: 6, Y: 5},
			Signal: "A",
		}, {
			Pos:    Pos{X: 8, Y: 8},
			Signal: "A",
		}, {
			Pos:    Pos{X: 9, Y: 9},
			Signal: "A",
		}}

	got := grid.getAntennas()

	if !slices.Equal(want, got) {
		t.Fatalf("In getAntennas, wanted %v, got %v\n", want, got)
	}

}

func TestGetAntinodePositions(t *testing.T) {
	grid, _ := makeGridFromFile("testInput.txt")
	antennas := grid.getAntennas()
	antinodePositions := getAntinodePositions(antennas)
	validPositions := []Pos{}

	for _, antinode := range antinodePositions {
		if antinode.X >= 0 && antinode.X < len(grid[0]) && antinode.Y >= 0 && antinode.Y < len(grid) {
			validPositions = append(validPositions, antinode)
		}
	}
	validPositions = removeDuplicatePositions(validPositions)
	want := 14
	got := len(validPositions)

	if got != want {
		t.Fatalf("In length of getAntinodePositions, wanted %d, got %d\n", want, got)
	}
}
