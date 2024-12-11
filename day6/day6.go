package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	part2()
}

func part1() {
	grid, player, err := makeGridFromFile("puzzleInput.txt")
	if err != nil {
		panic(err)
	}
	processPath(grid, player)
	count := grid.visitedCount()
	fmt.Printf("Visited Count: %d", count)
}

func part2() {
	grid, player, err := makeGridFromFile("puzzleInput.txt")
	if err != nil {
		panic(err)
	}

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			duplicate := grid.copy()
			if duplicate[i][j].IsWall {
				continue
			}
			duplicate[i][j].IsWall = true
			isLoop := processPath(duplicate, player)
			if isLoop {
				count += 1
			}
		}
	}
	fmt.Printf("Loop count: %d", count)
}

type Grid [][]Pos

func (g Grid) copy() Grid {
	var duplicate Grid
	duplicate = make(Grid, len(g))
	for i := range g {
		duplicate[i] = make([]Pos, len(g[i]))
		copy(duplicate[i], g[i])
	}
	return duplicate
}

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
			cell := g[i][j]
			if cell.IsWall {
				line += "#"
			} else if cell.VisitedDown || cell.VisitedUp || cell.VisitedLeft || cell.VisitedRight {
				line += "X"
			} else {
				line += "."
			}
		}
		result += line + "\n"
	}
	return result
}

func (g Grid) visitedCount() int {
	count := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			cell := g[i][j]
			if cell.VisitedDown || cell.VisitedUp || cell.VisitedLeft || cell.VisitedRight {
				count++
			}
		}
	}
	return count
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

// Returns true if the guard is stuck in a loop
func processPath(grid Grid, player Player) bool {
	grid[player.Y][player.X].markVisited(player)

Finished:
	for {
		var nextPos Pos
		switch player.MovingDirection {
		case "UP":
			if player.Y < 1 {
				break Finished
			}
			nextPos = grid[player.Y-1][player.X]
		case "DOWN":
			if player.Y > len(grid)-2 {
				break Finished
			}
			nextPos = grid[player.Y+1][player.X]
		case "RIGHT":
			if player.X > len(grid[player.Y])-2 {
				break Finished
			}
			nextPos = grid[player.Y][player.X+1]
		case "LEFT":
			if player.X < 1 {
				break Finished
			}
			nextPos = grid[player.Y][player.X-1]
		}
		if !nextPos.IsWall {
			hasVisited, err := nextPos.hasVisited(player)
			if err != nil {
				panic(err)
			}
			if hasVisited {
				return true
			}
			player.move()
			grid[player.Y][player.X].markVisited(player)
		} else {
			player.turn()
			grid[player.Y][player.X].markVisited(player)
		}
	}
	return false
}

func (p *Pos) markVisited(player Player) {
	switch player.MovingDirection {
	case "UP":
		p.VisitedUp = true
	case "DOWN":
		p.VisitedDown = true
	case "RIGHT":
		p.VisitedRight = true
	case "LEFT":
		p.VisitedLeft = true
	}
}

func (p *Pos) hasVisited(player Player) (bool, error) {
	switch player.MovingDirection {
	case "UP":
		return p.VisitedUp, nil
	case "DOWN":
		return p.VisitedDown, nil
	case "RIGHT":
		return p.VisitedRight, nil
	case "LEFT":
		return p.VisitedLeft, nil
	}
	return false, errors.New("Invalid player MovingDirection")
}

func (p *Player) move() {
	switch p.MovingDirection {
	case "UP":
		p.Y -= 1
	case "DOWN":
		p.Y += 1
	case "RIGHT":
		p.X += 1
	case "LEFT":
		p.X -= 1
	}
}

func (p *Player) turn() {
	switch p.MovingDirection {
	case "UP":
		p.MovingDirection = "RIGHT"
	case "DOWN":
		p.MovingDirection = "LEFT"
	case "RIGHT":
		p.MovingDirection = "DOWN"
	case "LEFT":
		p.MovingDirection = "UP"
	}
}
