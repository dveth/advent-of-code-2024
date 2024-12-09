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
