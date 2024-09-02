package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `R1, L3, R5, R5, R5, L4, R5, R1, R2, L1, L1, R5, R1, L3, L5, L2, R4, L1, R4, R5, L3, R5, L1, R3, L5, R1, L2, R1, L5, L1, R1, R4, R1, L1, L3, R3, R5, L3, R4, L4, R5, L5, L1, L2, R4, R3, R3, L185, R3, R4, L5, L4, R48, R1, R2, L1, R1, L4, L4, R77, R5, L2, R192, R2, R5, L4, L5, L3, R2, L4, R1, L5, R5, R4, R1, R2, L3, R4, R4, L2, L4, L3, R5, R4, L2, L1, L3, R1, R5, R5, R2, L5, L2, L3, L4, R2, R1, L4, L1, R1, R5, R3, R3, R4, L1, L4, R1, L2, R3, L3, L2, L1, L2, L2, L1, L2, R3, R1, L4, R1, L1, L4, R1, L2, L5, R3, L5, L2, L2, L3, R1, L4, R1, R1, R2, L1, L4, L4, R2, R2, R2, R2, R5, R1, L1, L4, L5, R2, R4, L3, L5, R2, R3, L4, L1, R2, R3, R5, L2, L3, R3, R1, R3`

type pos struct {
	x int
	y int
}

func main() {
	mymap := make(map[pos]int)

	dirs := []pos{}
	dirs = append(dirs, pos{x: 0, y: -1}) //up
	dirs = append(dirs, pos{x: 1, y: 0})  //right
	dirs = append(dirs, pos{x: 0, y: +1}) //down
	dirs = append(dirs, pos{x: -1, y: 0}) //left

	me := pos{x: 0, y: 0}

	dir := 0 // up 0, right 1, down 2, left 3

	for _, chunk := range strings.Split(input, ", ") {
		if chunk == "" {
			continue
		}

		switch chunk[0] {
		case 'R':
			dir++

		case 'L':
			dir--
		}

		dir = dir % 4
		if dir < 0 {
			dir = 3
		}

		mult, err := strconv.Atoi(chunk[1:])
		if err != nil {
			fmt.Println("yo! there's an error here dawg")
		}
		for i := 0; i < mult; i++ {
			me.x += dirs[dir].x
			me.y += dirs[dir].y
			fmt.Println(me)
			_, ok := mymap[me]
			if ok {
				fmt.Println("Part 2: ", max(me.x, -me.x)+max(me.y, -me.y))
				return
			}
			mymap[me] = 1
		}

	}

}
