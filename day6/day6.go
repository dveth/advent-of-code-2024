package main

import (
	"bufio"
	"os"
)

func main() {

}

type Grid [][]Pos

type Pos struct {
	IsWall       bool
	VisitedLeft  bool
	VisitedRight bool
	VisitedUp    bool
	VisitedDown  bool
}

type Player struct {
	X               int
	Y               int
	MovingDirection string
}

func (g Grid) gridToString() string {
	result := ""
	for i := 0; i < len(g); i++ {
		line := ""
		for j := 0; j < len(g[i]); j++ {
			if g[i][j].IsWall {
				line += "#"
			} else {
				line += "."
			}
		}
		result += line + "\n"
	}
	return result
}

func makeGridFromFile(filename string) (Grid, Player, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Grid{}, Player{}, err
	}

	scanner := bufio.NewScanner(file)
	var grid Grid
	grid = [][]Pos{}
	var player Player

	lineCount := 0
	for scanner.Scan() {
		positions := []Pos{}
		columnCount := 0
		for _, value := range scanner.Text() {
			newPos := Pos{}
			if value == '#' {
				newPos.IsWall = true
			}
			positions = append(positions, newPos)

			if value == '^' {
				player = Player{
					X:               columnCount,
					Y:               lineCount,
					MovingDirection: "UP",
				}
			}
			if value == '>' {
				player = Player{
					X:               columnCount,
					Y:               lineCount,
					MovingDirection: "RIGHT",
				}
			}
			if value == '<' {
				player = Player{
					X:               columnCount,
					Y:               lineCount,
					MovingDirection: "LEFT",
				}
			}
			if value == 'v' {
				player = Player{
					X:               columnCount,
					Y:               lineCount,
					MovingDirection: "DOWN",
				}
			}
			columnCount++
		}
		grid = append(grid, positions)
		lineCount++
	}

	return grid, player, nil
}
