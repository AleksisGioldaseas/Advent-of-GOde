package main

import (
	"fmt"
	"os"
	"strings"
)

type guard struct {
	x          int
	y          int
	dirIndex   int
	directions [4][2]int
}

func (G *guard) getDir() (int, int) {
	return G.x + G.directions[G.dirIndex%4][0], G.y + G.directions[G.dirIndex%4][1]
}

func (G *guard) rotate() {
	G.dirIndex += 1
}

func (G *guard) step() {
	G.x, G.y = G.getDir()
}

func main() {
	input := getInput()
	fmt.Println("Part 1 answer:", part1(input))
	fmt.Println("Part 2 answer:", part2(input))
}

func part1(input string) string {
	grid, theGuard := inputToGrid(input)
	grid[theGuard.y][theGuard.x] = 'X'
	for isInBounds(theGuard.x, theGuard.y, grid) {
		// fmt.Println(theGuard.directions[theGuard.dirIndex%4][0], theGuard.directions[theGuard.dirIndex%4][1], theGuard.dirIndex)
		// printGrid(grid)
		newX, newY := theGuard.getDir()
		if !isInBounds(newX, newY, grid) {
			theGuard.step()
		} else if grid[newY][newX] == '#' {
			theGuard.rotate()
		} else {
			theGuard.step()
			grid[theGuard.y][theGuard.x] = 'X'
		}
	}

	countX := 0
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == 'X' {
				countX++
			}
		}
	}

	return fmt.Sprint(countX)
}

func part2(input string) string {
	// inputLines := strings.Split(input, "\n")
	part2Answer := ""
	return part2Answer
}

func inputToGrid(input string) ([][]rune, guard) {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))

	theGuard := guard{}
	theGuard.directions = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for y, line := range lines {
		for x, tile := range line {
			switch tile {
			case '^':
				theGuard.x = x
				theGuard.y = y
				grid[y] = append(grid[y], '.')
			default:
				grid[y] = append(grid[y], tile)
			}
		}
	}
	return grid, theGuard
}

func printGrid(grid [][]rune) {
	for y := range grid {
		for x := range grid[0] {
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println()
	}
	fmt.Println("\n\n")
}

func isInBounds(x, y int, grid [][]rune) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func getInput() string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println(fileName, " file is empty")
		os.Exit(1)
	}
	input := strings.ReplaceAll(string(data), "\r\n", "\n") //doing this replace so it can handle both linux and window text format
	return strings.TrimSpace(input)                         //doing this cause usually there's an extra new line at the bottom of the input
}
