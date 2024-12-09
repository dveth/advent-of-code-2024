package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestGridToString(t *testing.T) {
	grid, _, err := makeGridFromFile("testinput.txt")
	if err != nil {
		t.Fatalf("Received error from gridToString: %s", err.Error())
	}

	want := ""

	file, err := os.Open("testinput.txt")
	if err != nil {
		t.Fatalf("Received error from gridToString: %s", err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		want += scanner.Text() + "\n"
	}

	got := grid.gridToString()
	want = strings.ReplaceAll(want, "^", ".")
	want = strings.ReplaceAll(want, ">", ".")
	want = strings.ReplaceAll(want, "v", ".")
	want = strings.ReplaceAll(want, "<", ".")

	if got != want {
		t.Fatalf("In gridToString, wanted \n%s, got \n%s", want, got)
	}
}

func TestGetPlayerFromFile(t *testing.T) {
	_, player, err := makeGridFromFile("testinput.txt")
	if err != nil {
		t.Fatalf("Received error from gridToString: %s", err.Error())
	}

	want := Player{
		X:               4,
		Y:               6,
		MovingDirection: "UP",
	}

	if player.MovingDirection != want.MovingDirection || player.X != want.X || player.Y != want.Y {
		t.Fatalf("In getting player from file, wanted %v, got %v\n", want, player)

	}
}

func TestMarkVisited(t *testing.T) {
	player := Player{
		MovingDirection: "RIGHT",
	}
	pos := Pos{}
	pos.markVisited(player)

	want := Pos{
		VisitedLeft:  false,
		VisitedRight: true,
		VisitedDown:  false,
		VisitedUp:    false,
	}

	if pos.VisitedDown != want.VisitedDown || pos.VisitedUp != want.VisitedUp || pos.VisitedRight != want.VisitedRight || want.VisitedLeft != pos.VisitedLeft {
		t.Fatalf("In markVisited, wanted %v, got %v\n", want, pos)
	}
}

func TestVisitedCount(t *testing.T) {
	grid, player, err := makeGridFromFile("testinput.txt")
	if err != nil {
		t.Fatalf("Received error from visitedCount: %s", err.Error())
	}

	want := 41
	processPath(grid, player)
	got := grid.visitedCount()

	if want != got {
		t.Logf("Current grid: \n%s\n", grid.gridToString())
		t.Fatalf("In visitedCount, wanted %d, got %d\n", want, got)
	}
}
