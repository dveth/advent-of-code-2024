package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// [row][column]
type Grid [][]string

func (g Grid) getFlippedHorizontal() Grid {
	var newGrid Grid
	for i := 0; i < len(g); i++ {
		reversedSlice := make([]string, len(g[i]))
		copy(reversedSlice, g[i])
		slices.Reverse(reversedSlice)
		newGrid = append(newGrid, reversedSlice)
	}
	return newGrid
}

func (g Grid) getXMASCount() int {
	return g.getHorizontalXMAS() + g.getVerticalXMAS() + g.getDiagonalXMAS()
}

func (g Grid) getHorizontalXMAS() int {
	count := 0
	for _, line := range g {
		lineString := strings.Join(line, "")
		count += strings.Count(lineString, "XMAS")
		count += strings.Count(lineString, "SAMX")
	}
	return count
}

func (g Grid) getVerticalXMAS() int {
	count := 0
	for i := 0; i < len(g); i++ {
		var line string
		for j := 0; j < len(g[i]); j++ {
			line += g[j][i]
		}
		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")
	}
	return count
}

func (g Grid) getDiagonalXMAS() int {
	return g.getDiagonalRightToLeftXMAS() + g.getDiagonalLeftToRightXMAS()
}

func (g Grid) getDiagonalLeftToRightXMAS() int {
	count := 0
	// Get first row's diagonals
	for i := 0; i < len(g); i++ {
		var line string
		for j := 0; j < len(g); j++ {
			if i+j >= len(g) {
				break
			}
			line += g[j][j+i]
		}
		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")
	}
	for i := 1; i < len(g); i++ {
		var line string
		for j := 0; j < len(g); j++ {
			if i+j >= len(g) {
				break
			}
			line += g[i+j][j]
		}
		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")
	}
	return count
}

func (g Grid) getDiagonalRightToLeftXMAS() int {
	reversedGrid := g.getFlippedHorizontal()
	return reversedGrid.getDiagonalLeftToRightXMAS()
}

func getGridFromFile(filename string) (Grid, error) {
	var grid Grid
	file, err := os.Open(filename)
	if err != nil {
		return grid, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var newSlice []string
		for _, value := range scanner.Text() {
			newSlice = append(newSlice, string(value))
		}
		grid = append(grid, newSlice)
	}
	return grid, nil
}

func main() {
	grid, err := getGridFromFile("puzzleInput.txt")
	if err != nil {
		panic(err)
	}
	count := grid.getXMASCount()
	fmt.Printf("Full count: %d\n", count)
}

func reverse(s string) string {
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
