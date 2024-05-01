package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6
`

var demo string = `0	2	7	0
`

func main() {

	intSlice := []int{}
	if strings.Contains(input, "\n") { //checking if the input contains a new line just in case
		input = strings.Split(input, "\n")[0]
	}
	//turning the string input to a slice of ints
	for _, inst := range strings.Split(input, "\t") {
		if inst != "" {

			val, err := strconv.Atoi(inst)
			if err != nil {
				fmt.Println("Atoi error!")
			}
			intSlice = append(intSlice, val)
		}
	}

	p1answer := 0
	pointer := 0
	coolMap := make(map[string]int)

	for {
		p1answer++

		winnerIndex := 0
		winnerVal := -1
		//finding which bank has the most blocks
		for i, val := range intSlice {
			if val > winnerVal {
				winnerIndex = i
				winnerVal = val
			}
		}

		pointer = winnerIndex
		intSlice[winnerIndex] = 0
		SliceSize := len(intSlice)
		//iterating over every entry after the biggest one and adding one until we've run out of blocks
		for i := 0; i < winnerVal; i++ {
			pointer++
			intSlice[pointer%SliceSize]++
		}

		//turning the intslice into a word we can use as key for a map for easy identification
		word := ""
		for _, num := range intSlice {
			word += string(num)
			word += string(",")
		}

		val, ok := coolMap[word] //attempt to read assuming the key already exists
		if ok == true {          //<- this means the word has been found before, therefore we've run into a repeated pattern, so we're done
			fmt.Println("Part 2:", p1answer-val)
			return
		}

		//this incrementation also happens to add a key+val entry if there isn't one already and then increments it,
		//pretty odd if you ask me...
		coolMap[word] = p1answer //assign p1answer that is basically counts the cycles, so that next time
		//we see a repeated state we can use this to count how many steps it took

	}

}
