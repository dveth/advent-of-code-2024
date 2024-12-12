package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
}

func part1() {
	grid, _ := makeGridFromFile("puzzleInput.txt")
	antennas := grid.getAntennas()
	antinodePositions := getAntinodePositions(antennas)
	validPositions := []Pos{}

	for _, antinode := range antinodePositions {
		if antinode.X >= 0 && antinode.X < len(grid[0]) && antinode.Y >= 0 && antinode.Y < len(grid) {
			validPositions = append(validPositions, antinode)
		}
	}
	validPositions = removeDuplicatePositions(validPositions)
	fmt.Printf("Total valid antinodes: %d", len(validPositions))
}

type Grid []string

func makeGridFromFile(filename string) (Grid, error) {
	file, err := os.Open(filename)
	var grid Grid
	grid = []string{}
	if err != nil {
		return grid, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, strings.TrimSpace(scanner.Text()))
	}
	return grid, nil
}

type Antenna struct {
	Pos
	Signal string
}

type Pos struct {
	X int
	Y int
}

func (g Grid) getAntennas() []Antenna {
	antennas := []Antenna{}
	for i := 0; i < len(g); i++ {
		line := g[i]
		for j := 0; j < len(line); j++ {
			if line[j] != '.' {
				newAntenna := Antenna{
					Pos: Pos{
						X: j,
						Y: i,
					},
					Signal: string(line[j]),
				}
				antennas = append(antennas, newAntenna)
			}
		}
	}
	return antennas
}

func getAntinodePositions(antennas []Antenna) []Pos {
	// For antenna 4,3 and antenna 5,5, antinode positions are at 3,1 and 6,7
	antinodePositions := []Pos{}
	for _, antenna1 := range antennas {
		for _, antenna2 := range antennas {
			if antenna1.X == antenna2.X && antenna1.Y == antenna2.Y { // Same antenna
				continue
			}
			if antenna1.Signal != antenna2.Signal {
				continue
			}
			xDifference := getDifference(antenna1.X, antenna2.X)
			yDifference := getDifference(antenna1.Y, antenna2.Y)

			newAntinode1 := Pos{}
			newAntinode2 := Pos{}

			if antenna1.X < antenna2.X {
				newAntinode1.X = antenna1.X - xDifference
				newAntinode2.X = antenna2.X + xDifference
			} else {
				newAntinode1.X = antenna1.X + xDifference
				newAntinode2.X = antenna2.X - xDifference
			}
			if antenna1.Y < antenna2.Y {
				newAntinode1.Y = antenna1.Y - yDifference
				newAntinode2.Y = antenna2.Y + yDifference
			} else {
				newAntinode1.Y = antenna1.Y + yDifference
				newAntinode2.Y = antenna2.Y - yDifference
			}
			antinodePositions = append(antinodePositions, newAntinode1)
			antinodePositions = append(antinodePositions, newAntinode2)
		}
	}
	return antinodePositions
}

func removeDuplicatePositions(positions []Pos) []Pos {
	checkedPositions := []Pos{}
	for _, position := range positions {
		if !checkPosExists(position, checkedPositions) {
			checkedPositions = append(checkedPositions, position)
		}
	}
	return checkedPositions
}

func checkPosExists(pos Pos, positions []Pos) bool {
	for _, value := range positions {
		if pos.X == value.X && pos.Y == value.Y {
			return true
		}
	}
	return false
}

func getDifference(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}
