package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := getInput()

	fmt.Println("Part 1 answer:", part1(input))
	fmt.Println("Part 2 answer:", part2(input))
}

func part1(input string) string {
	inputLines := strings.Split(input, "\n")
	part1Answer := ""

}

func part2(input string) string {
	inputLines := strings.Split(input, "\n")
	part2Answer := ""

}

func getInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println("Input.txt file is empty")
		os.Exit(1)
	}
	input := strings.ReplaceAll(string(data), "\r\n", "\n") //doing this replace so it can handle both linux and window text format

	return input
}
