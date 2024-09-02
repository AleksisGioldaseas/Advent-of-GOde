package main

// In a boid system, each bird (or boid) follows three basic rules:

//     Separation: Avoid crowding nearby boids. Move away from other boids that are too close.
//     Alignment: Align velocity with nearby boids. Move in the same direction as nearby boids.
//     Cohesion: Move towards the average position of nearby boids. Stay close to nearby boids.

// These rules can be implemented with simple vector calculations. For example, to achieve separation,
// calculate a vector pointing away from nearby boids and adjust the bird's velocity accordingly.
// Similar calculations can be done for alignment and cohesion. By combining these rules,
//  you can simulate the flocking behavior of birds in your project.

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const FPS int = 30

var baseGrid [90][300]rune
var screenGrid [90][300]rune

var birds []bird

type vec2 struct {
	x float64
	y float64
}

type bird struct {
	position  vec2
	direction vec2
	velocity  vec2
}

func main() {

	//init delay
	timeDelay := time.Microsecond * time.Duration(1000000/FPS)

	//init grid
	for y := 0; y < len(baseGrid); y++ {
		for x := 0; x < len(baseGrid[0]); x++ {
			if y == 0 || y == len(baseGrid)-1 || x == 0 || x == len(baseGrid[0])-1 {
				baseGrid[y][x] = rune('#')
			} else {
				baseGrid[y][x] = rune(' ')
			}
		}
	}

	initTerminal()

	birds = []bird{}

	for i := 0; i < 20; i++ {
		b := bird{}
		b.position.x, b.position.y = float64(rand.Int()%90), float64(rand.Int()%90)
		b.direction.x, b.direction.y = float64(rand.Int()%100), float64(rand.Int()%100)
		b.direction = normalizeVec2(b.direction)
	}

	counter := 0
	for {
		counter += 1

		resetScreen()
		resetCadet()

		for _, p := range birds {
			drawSquare(int(p.position.x), int(p.position.y), 2, '#', '#')
		}

		printScreen()
		time.Sleep(timeDelay)
	}

}

func drawSquare(x_offset, y_offset, size int, outsymb, insymb rune) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if posIsValid(x+x_offset, y+y_offset) {
				if x == 0 || x == size-1 || y == 0 || y == size-1 {
					screenGrid[y+y_offset][x+x_offset] = outsymb
				} else {
					screenGrid[y+y_offset][x+x_offset] = insymb
				}
			}
		}
	}
}

func birdScan(myb bird, rang float64) []bird { //returns which birds are within the range given
	rang = rang * rang

	diffx := 0.0
	diffy := 0.0

	newbirds := []bird{}

	for _, b := range birds {
		diffx = b.position.x - myb.direction.x
		diffx *= diffx
		diffy = b.position.y - myb.direction.y
		diffy *= diffy
		if diffx+diffy < rang {
			newbirds = append(newbirds, b)
		}
	}

	return newbirds
}

func posIsValid(x, y int) bool {
	return x >= 0 && x < len(screenGrid[0]) && y >= 0 && y < len(screenGrid)
}

func posValidEnough(x, y int) bool {
	return x >= -100 && x < len(screenGrid[0])+100 && y >= -100 && y < len(screenGrid)+100
}

func initTerminal() {
	fmt.Print("\033[?25l")
	fmt.Print("\033[2J")
}

func printScreen() {
	for y := 0; y < len(screenGrid); y++ {
		fmt.Println(string(screenGrid[y][:]))
	}
}

func resetScreen() {
	screenGrid = baseGrid
}

func resetCadet() {
	fmt.Print("\033[1;1H")
}

func stepPoints(slice *[]bird) {
	newslice := &[]bird{}
	for _, p := range *slice {

		p.position.x += p.velocity.x / 100
		p.position.y += p.velocity.y / 100
		if posValidEnough(int(p.position.x), int(p.position.y)) { //if point outside remove
			*newslice = append(*newslice, p)
		}
	}
	*slice = *newslice

}

func normalizeVec2(p vec2) vec2 {
	len := math.Sqrt(p.x*p.x + p.y*p.y) //find vector magnitude
	p.x /= len
	p.y /= len
	return p
}
