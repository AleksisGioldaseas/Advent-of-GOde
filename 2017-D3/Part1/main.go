package main

import (
	"fmt"
	"math"
	"strconv"
)

// var demo string = ""

// var input string = `289326`
var input string = `82`

func main() {
	str := input
	// P1Answer := 0
	myInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("some kind of error")
	}
	ring := int(math.Ceil(math.Sqrt(float64(myInt)))) / 2 //the ring that myInt is in
	startingVal := (ring + ring + 1) * (ring + ring + 1)  //the value on the bottom right corner of the ring myInt is in
	fmt.Println("my num is:", myInt, " and my ring is:", ring, " start val: ", startingVal)

	// if myInt >= ring*ring && myInt >= (ring*ring)+
	// for y := ring; y < -ring; y--{ //from bottom right corner walking up to top right corner

	// }

}

//Solution explanation and how I derived it
//I created on excel the shape of the spiral in an attempt to notice a pattern I could use to calculate my answer efficiently.

// 65	64	63	62	61	60	59	58	57
// 66	37	36	35	34	33	32	31	56
// 67	38	17	16	15	14	13	30	55
// 68	39	18	5	4	3	12	29	54
// 69	40	19	6  (1)	2	11	28	53
// 70	41	20	7	8  (9)	10	27	52
// 71	42	21	22	23	24 (25)	26	51
// 72	43	44	45	46	47	48 (49)	50
// 73	74	75	76	77	78	79	80 (81) < If you notice this patterns goes like this: 3*3, 5*5, 7*7, 9*9.
//                                        Basically x is incrementing by 2 and is then squared.
//                                        So I can use this number to calculate how many "rings" my number is away from the origin
//                                        And from there, I will try to locate the number on the ring, and then calculate the manhattan distance
//
//ring number = input > quare root, rounded up, then divided by 2 (remainder rejected), now we have the x,y position of the bottom right corner
//																					  And the value of that corner is x quared
//
//Then I can just trace the ring from the bottom right corner counter clock wise,
//to end up on my number, and then I just add together the x and y to get our answer.
//There's definitely a faster way, but my solution is plenty fast.
