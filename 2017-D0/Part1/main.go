package main

import "strings"

var input string = `YOUR_INPUT_HERE`

func main() {

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		for _, thing := range strings.Split(input, ",") {
			if thing == "" {
				continue
			}

		}
	}

}
