package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string = `YOUR_INPUT_HERE`

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
	for i := 0; i <= firewallLen; i++ {
		layer, ok := layerMap[i]

		if ok {
			if layer.ScannerPos == 0 {
				p1answer += i * layer.Length

			}
		}
		step()
	}
	fmt.Println("Part 1:", p1answer)

}

func step() {

	for key := range layerMap {

		if layerMap[key].ScannerPos+layerMap[key].Dir >= layerMap[key].Length || layerMap[key].ScannerPos+layerMap[key].Dir < 0 {
			layerMap[key].Dir *= -1
		}
		layerMap[key].ScannerPos += layerMap[key].Dir
	}
}