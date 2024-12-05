package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type page struct {
	before []int
	after  []int
}

func main() {
	input := getInput()

	fmt.Println("Part 1 answer:", part1(input))
	fmt.Println("Part 2 answer:", part2(input))
}

func part1(input string) string {
	theMap := make(map[int]page)
	parts := strings.Split(input, "\n\n")

	pages := strings.Split(parts[1], "\n")
	orderRules := strings.Split(parts[0], "\n")

	for _, rule := range orderRules {
		parts := strings.Split(rule, "|")

		left := atoi(parts[0])
		right := atoi(parts[1])

		val, ok := theMap[left]
		if ok {
			val.after = append(val.after, right)
			theMap[left] = val
		} else {
			theMap[left] = page{after: []int{right}}
		}

		val, ok = theMap[right]
		if ok {
			val.before = append(val.before, left)
			theMap[right] = val
		} else {
			theMap[right] = page{after: []int{left}}
		}
	}

	for _, line := range pages {
		for _, pageVal := range strings.Split(line, ",") {

			fmt.Println(pageVal, theMap[atoi(pageVal)])
		}
	}

	// inputLines := strings.Split(input, "\n")
	part1Answer := ""
	return part1Answer

}

func part2(input string) string {
	// inputLines := strings.Split(input, "\n")
	part2Answer := ""
	return part2Answer
}

func atoi(str string) int {

	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("yo, bad numbers", str, err)
		os.Exit(1)
	}
	return val
}

func getInput() string {
	fileName := "demo.txt"
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
