package main

import (
	"fmt"
	"strconv"
	"strings"
)

//TODO, problem is IMPOSSIBLE to brute force:
// things to try:
//		find the longest state cycle and modulo it so that skip() isn't called too often
// 		skip delay's that are multiples of known state that cannot work,
//		probably the best one: use the above idea is immediately find the smallest number without doing no manual calculation

var input string = `0: 3
1: 2
2: 4
4: 6
6: 4
8: 6
10: 5
12: 8
14: 8
16: 6
18: 8
20: 6
22: 10
24: 8
26: 12
28: 12
30: 8
32: 12
34: 8
36: 14
38: 12
40: 18
42: 12
44: 12
46: 9
48: 14
50: 18
52: 10
54: 14
56: 12
58: 12
60: 14
64: 14
68: 12
70: 17
72: 14
74: 12
76: 14
78: 14
82: 14
84: 14
94: 14
96: 14
`

var layerMap map[int]*layerT

type layerT struct {
	ScannerPos int
	Length     int
	Dir        int //1 means down, -1 means up
}

func main() {
	firewallLen := 0
	layerMap = make(map[int]*layerT)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		slice := strings.Split(line, ": ")

		layer, ok := strconv.Atoi(slice[0])
		if ok != nil {
			fmt.Println("Atoi error!")
		}

		len, ok := strconv.Atoi(slice[1])
		if ok != nil {
			fmt.Println("Atoi error!")
		}

		firewallLen = max(firewallLen, layer)
		l := &layerT{}
		l.Length = len
		l.Dir = 1

		layerMap[layer] = l

	}
	p1answer := 0
	delay := 0
	for {
		for key := range layerMap { //resetting the layer values
			layerMap[key].Dir = 1
			layerMap[key].ScannerPos = 0
		}

		for i := 0; i < delay; i++ {
			step()
		}
		caught := false
		for i := 0; i <= firewallLen; i++ {
			layer, ok := layerMap[i]

			if ok {
				if layer.ScannerPos == 0 {
					p1answer += i * layer.Length
					delay++
					caught = true
					break
				}
			}
			step()
		}
		if caught == false {
			fmt.Println("Part 2:", delay)
			return
		}

	}

}

func step() {

	for key := range layerMap {

		if layerMap[key].ScannerPos+layerMap[key].Dir >= layerMap[key].Length || layerMap[key].ScannerPos+layerMap[key].Dir < 0 {
			layerMap[key].Dir *= -1
		}
		layerMap[key].ScannerPos += layerMap[key].Dir
	}
}
